package quizer

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// QuizData represents a single quiz question and its corresponding answer.
type QuizData struct {
	Question string
	Answer   string
}

// IsAnswerCorrect checks whether the provided input matches the correct answer for the QuizData.
// It returns true if the input is correct, and false otherwise.
func (qd *QuizData) IsAnswerCorrect(input string) bool {
	return qd.Answer == input
}

// Quizer represents a data structure for conducting quizzes and tallying scores.
type Quizer struct {
	QuizData     []*QuizData
	QuizFilePath string
	Total        int
	Correct      int
	Out          io.Writer
}

// New creates a new Quizer object with the given quiz data file path.
// It initializes the Quizer object with an empty QuizData slice, sets the QuizFilePath to the given path,
// and sets the Total and Correct fields to zero.
// If loading the quiz data fails, it logs a fatal error.
// It returns the created Quizer object.
func New(path string) *Quizer {
	q := &Quizer{
		QuizData:     make([]*QuizData, 0),
		QuizFilePath: path,
		Total:        0,
		Correct:      0,
		Out:          os.Stdout,
	}
	if err := q.LoadQuizData(); err != nil {
		log.Fatal(err)
	}
	if q.QuizData == nil {
		log.Fatal("Your quiz file doesn't contain any questions.")
	}
	return q
}

// ExecuteQuiz loads the quiz data from the specified file path,
// conducts the quiz by asking questions and collecting answers,
// and outputs the number of correct answers.
func (q *Quizer) ExecuteQuiz(duration int) {
	t := time.NewTimer(time.Duration(duration) * time.Second)
	done := make(chan bool)
	go q.ConductQuiz(done)
	for {
		select {
		case <-done:
			if !t.Stop() {
				<-t.C
			}
			return
		case <-t.C:
			fmt.Fprintf(q.Out, "\nTime is up!")
			return
		}
	}
}

// GatherInput reads input from the user and returns it after removing leading/trailing whitespaces and newlines.
func (q *Quizer) GatherInput() string {
	r := bufio.NewReader(os.Stdin)
	input, _ := r.ReadString('\n')
	return strings.TrimSpace(strings.Replace(input, "\n", "", -1))
}

// ConductQuiz conducts a quiz by iterating through the QuizData slice in the Quizer struct.
// It prints each question and gathers user input to check if the answer is correct.
// If the answer is correct, it increments the Correct counter in the Quizer struct.
// Once all questions have been answered, it signals that the quiz is done by sending a true value to the done channel.
func (q *Quizer) ConductQuiz(done chan<- bool) {
	for _, quiz := range q.QuizData {
		fmt.Fprintf(q.Out, "What is %s?\n", quiz.Question)
		if quiz.IsAnswerCorrect(q.GatherInput()) {
			q.Correct++
		}
	}
	done <- true

}

// DisplayResults prints the number of questions answered correctly out of the total number of questions in the quiz.
func (q *Quizer) DisplayResults() {
	if _, err := fmt.Fprintf(q.Out, "You answered correctly to %d out of %d questions.\n", q.Correct, q.Total); err != nil {
		log.Fatal("failed to print results..")
	}
}

// parseQuizData parses the given byte data as CSV and creates an array of QuizData objects.
// It returns the array of QuizData objects and an error if parsing fails.
func parseQuizData(data []byte) ([]*QuizData, error) {
	// parse quiz data
	r := csv.NewReader(bytes.NewReader(data))
	r.Comma = ','
	quizRecords, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	total := len(quizRecords)
	if total > 0 {
		qd := make([]*QuizData, total)
		for i, v := range quizRecords {
			qd[i] = &QuizData{Question: v[0], Answer: v[1]}
		}
		return qd, nil
	}
	return nil, nil
}

// LoadQuizData reads the quiz data from the specified file path and populates the QuizData field of the Quizer instance.
// It returns an error if the file cannot be read or if there is an error parsing the quiz data.
// The function also sets the Total field to the number of quiz questions loaded.
func (q *Quizer) LoadQuizData() error {
	data, err := os.ReadFile(q.QuizFilePath)
	if err != nil {
		return err
	}

	q.QuizData, err = parseQuizData(data)
	if err != nil {
		return err
	}

	q.Total = len(q.QuizData)
	return nil
}
