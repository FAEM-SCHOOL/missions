package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

type questionAnswer struct {
	Question string
	Answer   string
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var fileName string
	fmt.Print("Enter file name: ")
	_, _ = fmt.Scan(&fileName)
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File does not exist")
		return
	}
	var quiz []questionAnswer
	_ = json.Unmarshal(file, &quiz)
	var answer string
	var count int
	var limS string
	var lim int
	for {
		fmt.Printf("Enter number of questions (max %d): ", len(quiz))
		_, _ = fmt.Scan(&limS)
		lim, err = strconv.Atoi(limS)
		if err != nil || lim > len(quiz) || lim < 1 {
			fmt.Println("Wrong number")
		} else {
			break
		}
	}
	a := rand.Perm(len(quiz))
	a = a[:lim]
	c := make(chan string)
	go timer(c, 9)
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d. %s=", i+1, quiz[a[i]].Question)
		//fmt.Printf("Answer(): ")
		go input(c)
		answer = <-c
		if answer == "end" {
			fmt.Println("\nTime is up")
			break
		}
		if answer == quiz[a[i]].Answer {
			count++
		}
	}
	fmt.Printf("Right answers: %d out of %d\n", count, len(a))
}

func timer(c chan string, t int) {
	fmt.Printf(" (%d)", t)
	time.Sleep(time.Millisecond * 100)
	for i := 10 * (t - 1); i >= 0; i-- {
		fmt.Print("\b\b\b\b")
		time.Sleep(time.Millisecond * 100)
		fmt.Printf(" (%d)", i/10+1)
	}
	c <- "end"
}

func input(c chan string) {
	var s string
	_, _ = fmt.Scan(&s)
	if len(c) != 0 {
		return
	}
	c <- s
}
