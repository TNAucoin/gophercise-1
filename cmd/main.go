package main

import (
	"flag"
	"github.com/tnaucoin/gophercise-1/quizer"
)

func main() {
	quizPath := flag.String("quizpath", "problems.csv", "path to quiz csv datafile")
	flag.Parse()

	q := quizer.NewQuizer(*quizPath)
	q.ExecuteQuiz()

}
