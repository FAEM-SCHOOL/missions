package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// getUserInput читает пользовательский ввод из потока stdin и передает его в строковый канал in
func getUserInput(in chan<- string) {
	var input string
	fmt.Fscanln(os.Stdin, &input)
	input = strings.TrimSpace(input)
	in <- input
}

// getQuizData читает csv файл вопросов по пути filename
// Формат файла: вопрос,ответ
func getQuizData(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	quiz, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	return quiz
}

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "problems.csv", "Путь к csv файлу вопросов. "+
		"Формат файла: вопрос,ответ")
	var timeout time.Duration
	flag.DurationVar(&timeout, "timeout", 30*time.Second, "Общее время на ответы")
	var doShuffle bool
	flag.BoolVar(&doShuffle, "shuffle", false, "Перемешать вопросы")
	var questionsCount int
	flag.IntVar(&questionsCount, "count", -1, "Ограничение числа вопросов")
	flag.Parse()

	// загружаем вопросы из файла
	quiz := getQuizData(filename)
	if len(quiz) < questionsCount {
		fmt.Println("Значение флага count превышает число вопросов в файле!")
		return
	}

	// перемешиваем вопросы, если стоит флаг shuffle
	if doShuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(quiz), func(i, j int) { quiz[i], quiz[j] = quiz[j], quiz[i] })
	}

	// пауза перед началом таймера
	fmt.Println("Нажмите Enter, чтобы начать викторину...")
	fmt.Scanln()

	score := 0
	timer := time.NewTimer(timeout)
	input := make(chan string) // канал в который будут попадать ответы пользователя
quizLoop:
	for index, question := range quiz {
		fmt.Printf("%d) %s\n", index+1, question[0])
		go getUserInput(input)
		select {
		case userAnswer := <-input: // получен ответ пользователя
			{
				if userAnswer == question[1] {
					score += 1
				}
			}
		case <-timer.C: // таймер закончился
			{
				fmt.Println("\nВремя вышло!")
				break quizLoop
			}
		}
		// преждевременный выход из цикла если число вопросов было ограничено флагом count
		if index+1 == questionsCount {
			break quizLoop
		}
	}

	if questionsCount == -1 {
		questionsCount = len(quiz)
	}
	fmt.Printf("Счет: %d/%d", score, questionsCount)
}
