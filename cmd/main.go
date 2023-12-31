package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"github.com/tnaucoin/gophercise-1/quizer"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	quizPath := flag.String("quizpath", "problems.csv", "path to quiz csv datafile")
	quizDuration := flag.Int("duration", 10, "time limit for quiz in seconds")
	flag.Parse()
	q := quizer.NewQuizer(*quizPath)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*quizDuration)*time.Second)
	defer cancel()

	fmt.Println("Press [Enter] to begin the quiz.")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	if strings.Contains(input, "\n") {
		go q.ExecuteQuiz(ctx)

	}

	<-ctx.Done()
	q.DisplayResults()
}
