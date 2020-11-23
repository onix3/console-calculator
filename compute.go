package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

var ops []string

// найти количество знаков в дробной части
func precition(f float64) int {
	f = f - math.Trunc(f) // извлечь дробную часть
	s := strconv.FormatFloat(f, 'f',10,64)
	s = strings.TrimRight(s,"0")
	return len(s)-2
}

// извлечь два числа, разделённые оператором
func getNumbers(re *regexp.Regexp, s string) (numbers []float64) {
	for _,numberString := range re.FindAllStringSubmatch(s,-1)[0][1:] {
		n,err := strconv.ParseFloat(numberString, 64)
		if err != nil {
			log_err(err)
		}
		numbers = append(numbers,n)
	}
	//fmt.Println("Числа: ", numbers[0], numbers[1])

	return
}

func doOperations(re *regexp.Regexp, s string, operator string) string {
	for re.MatchString(s) {
		oldPart := re.FindString(s) // "a*b"
		numbers := getNumbers(re,s) // извлечь a и b
		//fmt.Println("a = ", numbers[0], "   b = ", numbers[1])
		fl := float64(0)
		switch operator {
			case "*" : fl = numbers[0]*numbers[1]
			case "/" : fl = numbers[0]/numbers[1]
			case "+" : fl = numbers[0]+numbers[1]
		}

		newPart := strconv.FormatFloat(fl, 'f',9,64)

		// в строке s "a*b" заменить на вычисленный результат
		s = strings.Replace(s, oldPart, newPart,1)
		ops = append(ops, s)
	}
	return s
}

func compute() string {
	if dE == "" || dE == "-" {
		return ""
	}

	// порядок математических операций: * / + -
	s := dE
	ops = []string{s}

	// если в конце "0.", то убрать
	s = strings.TrimSuffix(s,"0.")
	ops = append(ops, s)

	// "-1-2" → "-1+-2"   |   любая разность есть сумма
	re := regexp.MustCompile(`(\d+)-`)
	s = re.ReplaceAllStringFunc(s,func (old string) string {
		return strings.Replace(old,"-", "+-", 1)
	})

	re_m := regexp.MustCompile(`(-?\d+\.\d+|-?\d+)\*(-?\d+\.\d+|-?\d+)`)
	re_d := regexp.MustCompile(`(-?\d+\.\d+|-?\d+)/(-?\d+\.\d+|-?\d+)`)
	re_a := regexp.MustCompile(`(-?\d+\.\d+|-?\d+)\+(-?\d+\.\d+|-?\d+)`)

	s = doOperations(re_m, s,"*")
	s = doOperations(re_d, s,"/")
	s = doOperations(re_a, s,"+")

	s = strings.TrimRight(s, "/*-+.")

	// вычисленное число и его форматированный вывод
	fl,err := strconv.ParseFloat(s,64)
	if err != nil {
		log_err(err)
	}
	// проблема в том, что 10.0^11 * 10.0^11 даст 10000000000000000000000
	//                   а 10.0^11 * 10.0^12 даст 99999999999999991611392
	// поэтому для такого большого числа будет научная запись
	if fl < math.Pow(10,20) {
		prec := precition(fl)
		if fl == math.Trunc(fl) {
			prec = 0 // если целое, то нет дробной части вовсе
		}
		s = strconv.FormatFloat(fl, 'f', prec,64)
	} else {
		s = strconv.FormatFloat(fl, 'e',3,64)
	}

	if dE == "" {
		s = ""
	}

	//fmt.Println(strings.Join(ops," → "))

	return strings.TrimRight(s,"/*-+")
}