package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/tnaucoin/gophercise-1/quizer"
	"log"
	"os"
	"strings"
)

func main() {
	quizPath := flag.String("quizpath", "problems.csv", "path to quiz csv datafile")
	quizDuration := flag.Int("duration", 30, "time limit for quiz in seconds")
	flag.Parse()
	q := quizer.New(*quizPath)

	fmt.Println("Press [Enter] to begin the quiz.")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	if strings.Contains(input, "\n") {
		q.ExecuteQuiz(*quizDuration)
	}
	q.DisplayResults()
}
