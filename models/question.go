package models

type QuestionType string

const (
	MultipleChoice QuestionType = "multiple_choice"
	TrueFalse      QuestionType = "true_false"
	ShortAnswer    QuestionType = "short_answer"
)

type Question struct {
	Text    string       `json:"text"`
	Image   string       `json:"image,omitempty"`
	Choices []string     `json:"choices"`
	Answer  interface{}  `json:"answer"`
	Type    QuestionType `json:"type"`
}
