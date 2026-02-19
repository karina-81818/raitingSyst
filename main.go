package main

import (
	"strconv"
	"fmt"
	"log"
)

func main(){

input := inputNumber("Введите оценку от 0 до 100: ")

res, err := raitingSyst(input)
if err != nil {
	log.Fatalf("raitingSyst error: %s", err)
}
fmt.Println(res)

}

func inputNumber(str string) int{
	var inputStr string
	for {
		fmt.Print(str)
		fmt.Scan(&inputStr)

		num, err := strconv.Atoi(inputStr)
		if err == nil {
			return num
		}
		fmt.Println("Ошибка: введите целое число!")
	}
}

func raitingSyst(input int) (string,  error){

	if input < 0 || input > 100 {
		return "", fmt.Errorf("введенная оценка %d некорректна.", input)
	}
	switch {
	case input >= 90:
		return "A", nil
	case input >= 80:
		return "B", nil
	case input >= 70:
		return "C", nil
	case input >= 60:
		return "D", nil
	default:
		return "F", nil
	}
}

