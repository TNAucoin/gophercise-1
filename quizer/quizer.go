package quizer

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/tnaucoin/gophercise-1/interaction"
	"log"
	"os"
)

// QuizData represents a single quiz question and its corresponding answer.
type QuizData struct {
	Question string
	Answer   string
}

// IsAnswerCorrect checks whether the provided input matches the correct answer for the QuizData.
// It returns true if the input is correct, and false otherwise.
// Example usage:
//
//	qd := &QuizData{
//		Question: "What is the capital of France?",
//		Answer:   "Paris",
//	}
//	fmt.Println(qd.IsAnswerCorrect("Paris")) // Output: true
//	fmt.Println(qd.IsAnswerCorrect("London")) // Output: false
func (qd *QuizData) IsAnswerCorrect(input string) bool {
	return qd.Answer == input
}

// Quizer represents a data structure for conducting quizzes and tallying scores.
type Quizer struct {
	QuizData     []*QuizData
	QuizFilePath string
	Total        int
	Correct      int
	IntrAction   *interaction.InputReader
}

// NewQuizer creates a new instance of the Quizer struct with initialized fields.
// It returns a pointer to the Quizer struct.
func NewQuizer(path string) *Quizer {
	return &Quizer{
		QuizData:     make([]*QuizData, 0),
		QuizFilePath: path,
		Total:        0,
		Correct:      0,
		IntrAction:   interaction.New(),
	}
}

// ExecuteQuiz loads the quiz data from the specified file path,
// conducts the quiz by asking questions and collecting answers,
// and outputs the number of correct answers.
func (q *Quizer) ExecuteQuiz() {
	err := q.LoadQuizData(q.QuizFilePath)
	if err != nil {
		log.Fatal(err)
	}
	q.ConductQuiz()
}

// ConductQuiz conducts the quiz by iterating over each quiz in the Quizer's QuizData.
// It prompts the user with the question and gathers their input using the provided InputReader.
// If the user's input matches the correct answer for the quiz, the Quizer's Correct count is incremented.
// After all the quizzes have been answered, the function prints the user's score as a percentage
// of correctly answered questions out of the total number of questions in the quiz.
// Example usage:
//
//	q := NewQuizer()
//	err := q.LoadQuizData("problems.csv")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	ir := &interaction.InputReader{reader: bufio.NewReader(os.Stdin)}
//	q.ConductQuiz(ir)
func (q *Quizer) ConductQuiz() {
	for _, quiz := range q.QuizData {
		fmt.Printf("What is %s?\n", quiz.Question)
		if quiz.IsAnswerCorrect(q.IntrAction.GatherInput()) {
			q.Correct++
		}
	}
	fmt.Printf("You answered correctly to %d out of %d questions.\n", q.Correct, q.Total)
}

// parseQuizData parses the given byte data as CSV and creates an array of QuizData objects.
// It returns the array of QuizData objects and an error if parsing fails.
//
// Example usage:
//
//	data, err := os.ReadFile(path)
//	if err != nil {
//	    return err
//	}
//	quizData, err := parseQuizData(data)
//	if err != nil {
//	    return err
//	}
//	q.QuizData = quizData
//
// QuizData is a struct containing question, answer, and correct fields.
//
// Quizer is a struct that holds the array of QuizData and provides methods to load quiz data and tally scores.
func parseQuizData(data []byte) ([]*QuizData, error) {
	// parse quiz data
	r := csv.NewReader(bytes.NewReader(data))
	r.Comma = ','
	quizRecords, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	total := len(quizRecords)
	qd := make([]*QuizData, total)
	for i, v := range quizRecords {
		qd[i] = &QuizData{Question: v[0], Answer: v[1]}
	}
	return qd, nil
}

// LoadQuizData loads the quiz data from the specified file path into the Quizer object.
// It reads the file content, parses it as CSV, and creates an array of QuizData objects.
// The loaded quiz data is assigned to the QuizData field of the Quizer object.
// If any error occurs during file reading or parsing, it is returned.
//
// Example usage:
//
//	q := NewQuizer()
//	err := q.LoadQuizData("problems.csv")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (q *Quizer) LoadQuizData(path string) error {
	data, err := os.ReadFile(path)
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
