package main

import (
	
	"fmt"
	"log"
)

func main(){

fmt.Print("Введите оценку от 0 до 100: ")
 var input  int
 fmt.Scan(&input)


res, err := raitingSyst(input)
if err != nil {
	log.Fatalf("raitingSyst error: %s", err)
}
fmt.Println(res)

}

func raitingSyst(input int) (string,  error){

	if input < 0 || input > 100 {
		return "", fmt.Errorf("raitingSyst: введенная оценка %d некорректна.", input)
	}
	if input >= 90 || input <=100{
		return "A", nil
	}
	if input >= 80 || input <= 89{
		return "B", nil	
	}
	if input >= 70 || input <= 79{
		return "C", nil
	}
	if input >= 60 || input <= 69{
		return "D", nil
	}
		return "F", nil
	}



