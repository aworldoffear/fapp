package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	graphicText("Добро пожаловать!\n\n")
	graphicText("Пожалуйста, введите математическое выражение (пробелы и табуляции не имеют значения): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var expression string = strings.ReplaceAll(strings.ReplaceAll(scanner.Text(), " ", ""), "\t", "")
	var (
		symbol string
		expv2 []string
	)
	for count, symofExp := range expression {
		if count == 0 {
			continue
		}
		switch symofExp {
		case '+': symbol = "+"
		case '-': symbol = "-"
		case '/': symbol = "/"
		case '*': symbol = "*"
		}
		if symbol != "" {
			break
		}
	}
	expv2 = strings.Split(expression, symbol)
	var length int = len(expv2)
	switch {
		case length == 1:
			graphicText("Cтрока не является математической операцией.")
			os.Exit(0)
		case length > 2:
			graphicText("Необходимо два операнда и один оператор.")
			os.Exit(0)
	}
	fmt.Println(mainCalc(expv2[0], expv2[1], symbol))
}

func mainCalc(fNumber, sNumber, symbol string) string {
	switch {
		case tryConvert(fNumber) && tryConvert(sNumber):
			return strconv.Itoa(arabicCalc(fNumber, sNumber, symbol))
		case !tryConvert(fNumber) && !tryConvert(sNumber):
			return romanCalc(fNumber, sNumber, symbol)
		default:
			graphicText("Ошибка! Обнаружены разные системы счисления в одном выражении.")
			os.Exit(0)
	}
	return "0"
}

func tryConvert(strNumber string) bool {
	_, err := strconv.Atoi(strNumber)
	if err != nil {
		return false
	}
	return true
}

func arabicCalc(fNumber, sNumber, symbol string) int {
	intfNumber, _ := strconv.Atoi(fNumber)
	intsNumber, _ := strconv.Atoi(sNumber)
	switch {
		case intfNumber <= 0:
			graphicText("Ошибка! Введённое арабское число должно быть в диапазоне от 1 до 10 включительно.")
			os.Exit(0)
		case intsNumber <= 0:
			graphicText("Ошибка! Введённое арабское число должно быть в диапазоне от 1 до 10 включительно.")
			os.Exit(0)
		case intfNumber > 10:
			graphicText("Ошибка! Введённое арабское число должно быть в диапазоне от 1 до 10 включительно.")
			os.Exit(0)
		case intsNumber > 10:
			graphicText("Ошибка! Введённое арабское число должно быть в диапазоне от 1 до 10 включительно.")
			os.Exit(0)
	}
	switch symbol {
		case "+": return intfNumber + intsNumber
		case "-": return intfNumber - intsNumber
		case "/": return intfNumber / intsNumber
		case "*": return intfNumber * intsNumber
		default: 
		graphicText("Ошибка! Неверный символ. Возможно использование следующих символов: +, -, /, *")
		panic(symbol)
	}
}

func romanCalc(fNumber, sNumber, symbol string) string {
	var intfNumber int = fromRomanNumbertoInt(fNumber)
	var intsNumber int = fromRomanNumbertoInt(sNumber)
	var result int
	switch symbol {
		case "+": result = intfNumber + intsNumber
		case "-":
			result = intfNumber - intsNumber
			if result <= 0 {
				graphicText("Ошибка! Результат вычислений римских чисел был равен 0 или меньше.")
				panic(result)
			}
		case "/": result = intfNumber / intsNumber
		case "*": result = intfNumber * intsNumber
		default: 
		graphicText("Ошибка! Неверный символ. Возможно использование следующих символов: +, -, /, *")
		panic(symbol)
	}
	return fromIntNumbertoRoman(result)
}

func fromRomanNumbertoInt(number string) int {
	romanNumber := map[string]int {
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	for a, b := range romanNumber {
		if a == number {
			return b
		}
	}
	graphicText("Ввод римских чисел только от I до X.")
	panic(number)
}

func fromIntNumbertoRoman(number int) string {
	var result string
	romanNumber := map[string]int {
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
		"XX": 20, "XXX": 30, "XL": 40, "L": 50, "LX": 60, "LXX": 70, "LXXX": 80, "XC": 90, "C": 100,
	}
	for a, b := range romanNumber {
		if b == number {
			return a
		}
	}
	for a, b := range romanNumber {
		if b == number - (number % 10) {
			result += a
		}
	}
	for a, b := range romanNumber {
		if b == number % 10 {
			result += a
		}
	}
	return result
}

func graphicText(text string) {
	for _, b := range text {
		fmt.Print(string(b))
		time.Sleep(25 * time.Millisecond)
		switch string(b) {
		case ",": time.Sleep(150 * time.Millisecond)
		case "!": time.Sleep(150 * time.Millisecond)
		}
	}
	time.Sleep(500 * time.Millisecond)
}