package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Quiz() {
	file, err := os.Open("hello.txt")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	count := 0

	scanner := bufio.NewScanner(file)
	text := ""
	for scanner.Scan(){

		text = string(scanner.Text())
		QuestionAndAnswers := strings.Split(text, ",")

		if len(QuestionAndAnswers) < 2{ break }

		fmt.Println(QuestionAndAnswers[0])

		var answer string
		fmt.Print("Напишите Ваш ответ: " )
		fmt.Scan(&answer)

		if answer == QuestionAndAnswers[1] {
			count++
			fmt.Println("Поздравляю! Ответ правильный. Ваши баллы: " , count)
			fmt.Println()
		}else {
			fmt.Println("Ответ не правильный :(")
			fmt.Println()
		}
	}

	fmt.Println("Конец квиза. Ваш итоговый балл: ", count)

}

func Help()  {
	fmt.Println("Справка:")
	fmt.Println("q - начало квиза")
	fmt.Println("qt - начало квиза на время")
	fmt.Println("h - справка")
	fmt.Println("e - выход")
}

func main()  {
	Help()

	for{
		var command string
		fmt.Scan(&command)
		if command == "q"{
			Quiz()
			Help()
		}else {
			if command == "qt"{
				t1 := time.Now()
				Quiz()
				t2 := time.Now()

				fmt.Println("Время:", t2.Sub(t1))

				Help()
			}else {
				if command == "e"{
					fmt.Println("Закрытие квиза....")
					break
				}else {Help()}

			}


		}


	}


}