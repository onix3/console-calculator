package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Найти количество знаков в дробной части
func precition(f float64) int {
	f = f - math.Trunc(f) // извлечь дробную часть
	s := strconv.FormatFloat(f, 'f',10,64)
	s = strings.TrimRight(s,"0")
	return len(s)-2
}

// Извлечь два числа, разделённые оператором
// "a*b" → a (float64) и b (float64)
func getNumbers(re *regexp.Regexp, s string) (numbers []float64) {
	for _,numberString := range re.FindAllStringSubmatch(s,-1)[0][1:] {
		n,err := strconv.ParseFloat(numberString, 64)
		if err != nil {
			log_err(err)
		}
		numbers = append(numbers,n)
	}
	return
}

func doOperations(re *regexp.Regexp, s string, operator string) string {
	for re.MatchString(s) {
		oldPart := re.FindString(s) // исходная подстрока "a*b"
		numbers := getNumbers(re,s) // извлечь числа
		fl := float64(0)
		switch operator { 			// вычислить
			case "*" : fl = numbers[0]*numbers[1]
			case "/" : fl = numbers[0]/numbers[1]
			case "+" : fl = numbers[0]+numbers[1]
		}

		// форматировать
		newPart := strconv.FormatFloat(fl, 'f',9,64)

		// подстроку "a*b" заменить на вычисленное
		s = strings.Replace(s, oldPart, newPart,1)
	}
	return s
}

// Вычисление выражения
func compute() string {
	// не вычислять некорректное выражение
	if dE == "" || dE == "-" || strings.HasSuffix(dE,"0.") {
		return ""
	}

	s := dE

	// "-1-2" → "-1+-2"   |   нет разности, есть отрицательные числа и сумма
	re := regexp.MustCompile(`(\d+)-`)
	s = re.ReplaceAllStringFunc(s,func (old string) string {
		return strings.Replace(old,"-", "+-", 1)
	})

	// паттерны   `(дробное или целое)оператор(дробное или целое)`
	re_m := regexp.MustCompile(`(-?\d+\.\d+|-?\d+)\*(-?\d+\.\d+|-?\d+)`)
	re_d := regexp.MustCompile(`(-?\d+\.\d+|-?\d+)/(-?\d+\.\d+|-?\d+)`)
	re_a := regexp.MustCompile(`(-?\d+\.\d+|-?\d+)\+(-?\d+\.\d+|-?\d+)`)

	// приоритет математических операций: * / + -
	s = doOperations(re_m, s,"*")
	s = doOperations(re_d, s,"/")
	s = doOperations(re_a, s,"+")

	// убрать оставшиеся не-цифры справа
	s = strings.TrimRight(s, "/*-+.")

	// вычисленное значение и его форматированный вывод
	fl,err := strconv.ParseFloat(s,64)
	if err != nil {
		log_err(err)
		return ""
	}

	// есть проблема: 10.0^11 * 10.0^11 даст 10000000000000000000000
	//              а 10.0^11 * 10.0^12 даст 99999999999999991611392
	// поэтому для больших чисел будет научная запись
	if -math.Pow(10,15) < fl && fl < math.Pow(10,15) {
		prec := precition(fl)
		if fl == math.Trunc(fl) {
			prec = 0 // если целое, то показывать без дробной части
		}
		s = strconv.FormatFloat(fl, 'f', prec,64)
	} else {
		s = strconv.FormatFloat(fl, 'e',3,64)
	}

	return s
}