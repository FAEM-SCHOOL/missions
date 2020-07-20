package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

const quizFilename = "quiz.csv"

var start string

func main() {
	f, err := os.Open(quizFilename)
	if err != nil {
		fmt.Printf("Не удалось открыть файл : %v", err)
		return
	}
	defer f.Close()
	read := csv.NewReader(f)
	records, err := read.ReadAll()

	if err != nil {
		fmt.Printf("Не удалось прочитать csv файл: %v\n", err)
		return
	}

	fmt.Println("Сейчас будут появляться задания, вы должны будете их решить")
	fmt.Println("При готовности введите любой символ")
	fmt.Scan(&start)

	var correctAnswers int

	t0 := time.Now()

	for i, record := range records {

		question, correctAnswer := record[0], record[1]
		fmt.Printf("%d. %s?\n", i+1, question)

		var answer string

		if _, err := fmt.Scan(&answer); err != nil {
			fmt.Printf("Не удалось счесть ответ :%v\n", err)
			return
		}
		fmt.Printf("Ваш ответ: %s\n", answer)


		if answer == correctAnswer {
			correctAnswers++
		}
	}

	t1 := time.Now()


	fmt.Println("Вы решили все задания за: ",t1.Sub(t0)," секунд")
	fmt.Println("Результат : ", correctAnswers,"/", len(records))

}
