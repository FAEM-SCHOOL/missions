package main

import (
	"encoding/csv"
	"fmt"
	"os"
)
type Exam struct{
	Question string `json:"qustion"`
	Answer   string `json:"answer"`

}
var command string
var J bool
//Command list output function.
func help(){
	fmt.Println("Help - Справачная информация по командам.")
	fmt.Println("Quiz - начать прохождение викторины.")
	fmt.Println("Exit - закрыть программу.")
	command =""
}
//The function that conducts the quiz counts and displays the results.
func quiz(){
	score:= 0
	num:= 0
	csvFile, err := os.Open("problem.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	cnt := 0
	for _, line := range csvLines {
		cnt++
		emp := Exam{
			Question: line[0],
			Answer:   line[1],
		}
		fmt.Printf("Question %d: %s", cnt, emp.Question)
		command =""
		fmt.Print("\nAnswer: ")
		fmt.Scan(&command)
		num+=1
		if emp.Answer== command {
			score +=1
		}
	}
	fmt.Println(" Ваши баллы = " , score , "Количество вопросов = " , num)
	command =""
	return
}
func main(){
	fmt.Println("Для получения информации по командам <h>, E выйти из программы.")
	for{
	fmt.Scan(&command)
	if command =="Exit"{
		command =""
	break}
	if command !=""{
		fcomm()}
	}
	}
//A pool of commands that perform the transition against a fool.
func fcomm() {
	protection := [4]string{"H", "h", "Help", "help"}
	for i := 0; i < len(protection)-1; i++ {
		if command == protection[i] {
			help()
			break}}

	quizi := [4]string{"Q", "q", "Quiz", "quiz"}
	for i := 0; i < len(quizi)-1; i++ {
		if command == quizi[i] {
			quiz()
			break}}


}