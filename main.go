// Parse the csv file into a quiz that can be used on a terminal

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file := "problems.csv"
	// if different file given:
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	// open file
	csvFile, _ := os.Open(file)
	// read file as csv
	reader := csv.NewReader(bufio.NewReader(csvFile))
	// initialize dictionary of question: answer
	var quiz = make(map[string]string)
	// while loop, until there isn't more to read
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		}
		// question, solution
		quiz[line[0]] = line[1]
	}

	timed(quiz)
	// readInput := bufio.NewReader(os.Stdin)
	// total := len(quiz) // how many total questions
	// var score int64    // answers correct

	//
	// for k := range quiz {

	// 	// print question
	// 	fmt.Printf("%s: ", k)
	// 	// prompt user for answer
	// 	text, _ := readInput.ReadString('\n')
	// 	// remove new line to compare strings
	// 	t := strings.Trim(text, "\n")
	// 	if t == quiz[k] {
	// 		score++
	// 	}

	// }
	// fmt.Printf("Result: %d/%d\n", score, total)
}

func getInput(readInput *bufio.Reader, c chan string, quiz map[string]string, total int, score *int64) {

	for k := range quiz {

		// print question
		fmt.Printf("%s: ", k)
		// prompt user for answer
		text, _ := readInput.ReadString('\n')
		// remove new line to compare strings
		t := strings.Trim(text, "\n")
		if t == quiz[k] {
			*score++
		}

	}
	result := "Result: " + strconv.Itoa(int(*score)) + "/" + strconv.Itoa(int(total))
	c <- result
}

// 30 seconds per quiz
func timed(quiz map[string]string) {
	readInput := bufio.NewReader(os.Stdin)
	total := len(quiz) // how many total questions
	var score *int64   // answers correct

	// creating a channel
	ch := make(chan string)
	// press enter to start quiz
	fmt.Println("Press Enter to Start:")
	_, _ = readInput.ReadString('\n')
	// start timer
	// start := time.Now()
	go getInput(readInput, ch, quiz, total, score)
	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(time.Second * 30):

	}

	fmt.Printf("Result: %d/%d\n", score, total)
	return
}

func getInputQ(readInput *bufio.Reader, c chan string) {
	text, _ := readInput.ReadString('\n')
	c <- text
}

// 30 seconds per question
func timedQ(quiz map[string]string) {
	readInput := bufio.NewReader(os.Stdin)
	total := len(quiz) // how many total questions
	var score int64    // answers correct

	// creating a channel
	ch := make(chan string)
	// press enter to start quiz
	fmt.Println("Press Enter to Start:")
	_, _ = readInput.ReadString('\n')
	// start timer
	// start := time.Now()

	for k := range quiz {
		// print question
		fmt.Printf("%s: ", k)

		// TODO: Need to channel this so we can test for timer
		// prompt user for answer
		go getInputQ(readInput, ch)
		select {
		case text := <-ch:
			fmt.Println("Message 1", text)
			// remove new line to compare strings
			t := strings.Trim(text, "\n")
			if t == quiz[k] {
				score++
			}
		case <-time.After(time.Second * 30):
			continue
		}
		// ch <- fmt.Sprintf("%v, %v", readInput.ReadString('\n'))

	}
	fmt.Printf("Result: %d/%d\n", score, total)
	return
}

// msg := <- c
// fmt.Println(msg)
// time.Sleep(time.Second * 1)

// select {
// case msg1 := <- c1:
//   fmt.Println("Message 1", msg1)
// case msg2 := <- c2:
//   fmt.Println("Message 2", msg2)
// case <- time.After(time.Second):
//   fmt.Println("timeout")
// }
