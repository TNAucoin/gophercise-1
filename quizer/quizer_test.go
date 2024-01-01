package quizer

import (
	"bytes"
	"reflect"
	"testing"
)

func Test_parseQuizData(t *testing.T) {
	// Define table driven tests
	tests := []struct {
		name    string
		data    []byte
		want    []*QuizData
		wantErr bool
	}{
		{
			"Valid data",
			[]byte("question1,answer1\nquestion2,answer2"),
			[]*QuizData{
				&QuizData{Question: "question1", Answer: "answer1"},
				&QuizData{Question: "question2", Answer: "answer2"},
			},
			false,
		},
		{
			"Empty data",
			[]byte(""),
			nil,
			false,
		},
		{
			"Invalid CSV",
			[]byte("question1, answer1 \nquestion2"),
			nil,
			true,
		},
	}

	// Run each test case from the tests slice
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseQuizData(tt.data)

			if (err != nil) != tt.wantErr {
				t.Errorf("parseQuizData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseQuizData() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_DisplayResults(t *testing.T) {

	// Define test cases
	cases := []struct {
		name    string
		total   int
		correct int
		want    string
	}{
		{
			name:    "no correct answers",
			total:   5,
			correct: 0,
			want:    "You answered correctly to 0 out of 5 questions.\n",
		},
		{
			name:    "all correct answers",
			total:   5,
			correct: 5,
			want:    "You answered correctly to 5 out of 5 questions.\n",
		},
		{
			name:    "some correct answers",
			total:   5,
			correct: 3,
			want:    "You answered correctly to 3 out of 5 questions.\n",
		},
	}

	// Run the test cases
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			quiz := &Quizer{
				Total:   tc.total,
				Correct: tc.correct,
			}

			var b bytes.Buffer
			quiz.Out = &b
			quiz.DisplayResults()

			got := b.String()

			if got != tc.want {
				t.Errorf("DisplayResults() = %v, want %v", got, tc.want)
			}
		})
	}
}
