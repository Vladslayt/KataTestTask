package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		symbols := strings.Fields(text)
		if len(symbols) > 3 {
			err := errors.New("incorrect format")
			panic(err)
		}
		firstNumberRoman, secondNumberRoman := false, false
		firstNumber, err := strconv.Atoi(symbols[0])
		if err != nil {
			firstNumberRoman = true
		}

		secondNumber, err := strconv.Atoi(symbols[2])
		if err != nil {
			secondNumberRoman = true
		}

		if firstNumberRoman && secondNumberRoman {
			firstNumber, err := romanToNumber(symbols[0])
			if err != nil {
				panic(err)
			}

			secondNumber, err := romanToNumber(symbols[2])
			if err != nil {
				panic(err)
			}

			outputNumberArabian, err := operationsRoman(symbols[1], firstNumber, secondNumber)
			if err != nil {
				panic(err)
			}
			outputNumberRoman := numberToRoman(outputNumberArabian)

			fmt.Println(outputNumberRoman)

		} else if !firstNumberRoman && !secondNumberRoman {
			if firstNumber <= 10 && secondNumber <= 10 {
				outputNumber := operationsArabian(symbols[1], firstNumber, secondNumber)
				fmt.Println(outputNumber)
			} else {
				err := errors.New("incorrect numbers")
				panic(err)
			}
		} else {
			err := errors.New("type of numbers not equals")
			panic(err)
		}

	}
}

func operationsArabian(oper string, firstNumber, secondNumber int) (output int) {
	switch oper {
	case "+":
		output = firstNumber + secondNumber
	case "-":
		output = firstNumber - secondNumber
	case "*":
		output = firstNumber * secondNumber
	case "/":
		output = firstNumber / secondNumber
	}
	return output
}

func operationsRoman(oper string, firstNumber, secondNumber int) (output int, err error) {
	switch oper {
	case "+":
		output = firstNumber + secondNumber
	case "-":
		output = firstNumber - secondNumber
	case "*":
		output = firstNumber * secondNumber
	case "/":
		output = firstNumber / secondNumber
	}
	if output <= 0 {
		err := errors.New("roman number <= 0")
		return 0, err
	}
	return output, err
}

func romanToNumber(str string) (number int, err error) {
	switch str {
	case "I":
		number = 1
	case "II":
		number = 2
	case "III":
		number = 3
	case "IV":
		number = 4
	case "V":
		number = 5
	case "VI":
		number = 6
	case "VII":
		number = 7
	case "VIII":
		number = 8
	case "IX":
		number = 9
	case "X":
		number = 10
	default:
		err := errors.New("error number: number more 10")
		return 0, err
	}
	return number, err
}

func numberToRoman(outputNumberArabian int) (number string) {
	outputNumberRoman := ""
	tens := outputNumberArabian / 10
	if tens == 10 {
		outputNumberRoman = "C"
	} else {
		if tens >= 5 {
			tens -= 5
			outputNumberRoman = "L"
		}
		for tens != 0 {
			outputNumberRoman += "X"
			tens--
		}
		if outputNumberArabian%10 != 0 {
			units, _ := unitNumberToRoman(outputNumberArabian % 10)
			outputNumberRoman += units
		}
	}
	return outputNumberRoman
}

func unitNumberToRoman(str int) (number string, err error) {
	switch str {
	case 1:
		number = "I"
	case 2:
		number = "II"
	case 3:
		number = "III"
	case 4:
		number = "IV"
	case 5:
		number = "V"
	case 6:
		number = "VI"
	case 7:
		number = "VII"
	case 8:
		number = "VIII"
	case 9:
		number = "IX"
	case 10:
		number = "X"
	default:
		err := errors.New("error number: number more 10")
		return "", err
	}
	return number, err
}
