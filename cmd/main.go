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
	quizPath := flag.String("quizpath", "problems.csv", "path to quiz csv datafile.")
	quizDuration := flag.Int("duration", 30, "time limit for quiz in seconds.")
	shuffleQuiz := flag.Bool("shuffle", true, "shuffle the questions in quiz data file.")
	flag.Parse()
	q := quizer.New(*quizPath, *shuffleQuiz)

	fmt.Println("Press [Enter] to begin the quiz.")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	if strings.Contains(input, "\n") {
		q.ExecuteQuiz(*quizDuration)
	}
}
