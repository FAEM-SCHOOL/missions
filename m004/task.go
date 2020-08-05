package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

type task struct {
	Name           string `json:"name"`
	CompletionDate string `json:"completionDate,omitempty"`
}

var arr []task
var filename = "problem.json"
var commands = ""

func main() {
	fmt.Println("1-Помощь.")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(data, &arr); err != nil {
	}
	for {
		var a int64
		fmt.Println("Вы на главном экране", a)
		Scan1()
		a = +a
		switch commands {
		case "1":
			{
				help()
			}
		case "2":
			{
				fmt.Println("Введите наименование задания.")
				Scan1()
				add(commands)
				fmt.Println("Задание добавлено.")
				break
			}
		case "3":
			{
				var n int
				for {
					fmt.Println("Введите номер задания для вывода на экран.")
					var id string
					fmt.Scan(&id)
					z, err := strconv.ParseInt(id, 10, 32)
					if err != nil {
						fmt.Println("Неверный ввод номера задания")
					} else {
						n = int(z)
						break
					}
				}
				do(n)
				break
			}
		case "4":
			{
				listComplete()
				break
			}
		case "5":
			{
				listIncomplete()
				break
			}
		case "6":
			{
				var n int
				for {

					var id string
					q := int64(len(arr))
					fmt.Scan(&id)
					h, err := strconv.ParseInt(id, 10, 32)
					z, err := strconv.ParseInt(id, 10, 32)
					if err != nil {
						fmt.Println("Неверный ввод номера задания")
					} else if q < h {
						fmt.Println("Неверный номер задания")
					} else {
						n = int(z)
						fmt.Println("Задание удалено.")
						break
					}
				}
				remove(n - 1)
				break
			}
		case "7":
			{
				return
			}

		case "":
			{
				break
			}
		}
		commands = ""
	}
}

//Изменение статуса задания на "Выполненое" и запись даты выполнения.
func do(id int) {
	arr[id].CompletionDate = time.Now().Format("2006-01-02T15:04:05")
	updateFile()
}

//Обновление файла.
func updateFile() {
	data, err := json.Marshal(arr)
	ioutil.WriteFile(filename, data, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

//Добавление нового задания.
func add(name string) {
	arr = append(arr, task{Name: name})
	updateFile()
}

//Вывод незавершённых заданий.
func listIncomplete() {
	for i := 0; i < len(arr); i++ {
		if arr[i].CompletionDate == "" {
			fmt.Print("№", i+1, ":", arr[i].Name, "\n")
		}
	}
}

//Вывод завершённых заданий.
func listComplete() {
	for i := 0; i < len(arr); i++ {
		if arr[i].CompletionDate != "" {
			fmt.Print("№", i+1, ":")
			fmt.Println(arr[i].Name, " Дата выполнения:", arr[i].CompletionDate)
		}
	}
}

//Вывод справки по командам.
func help() {
	fmt.Println("1-Помощь.\n2-Добавить задачу.\n3-Задача выполнена.\n4-Вывести выполненые задачаи.\n" +
		"5-Вывести невыполненые задачи.\n6-Удалить задачу")
}

//Удаление задания из списка.
func remove(index int) {
	arr = append(arr[:index], arr[index+1:]...)
	updateFile()
}

//Функция скана командной строки.
func Scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	commands = in.Text()
	return in.Text()
}
