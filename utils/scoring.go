package utils

import (
	"log"
	"github.com/Jupita-SA/cbt/models"
	"strconv"
	"strings"
)

const questionWeight = 1.0

func scoreMultipleChoice(userAnswer string, correctAnswer string) float64 {
	if userAnswer == "" {
		return 0
	}
	if userAnswer == correctAnswer {
		return questionWeight
	}
	return 0
}

func scoreTrueFalse(userAnswer string, correctAnswer string) float64 {
	if userAnswer == "" {
		return 0
	}
	if userAnswer == correctAnswer {
		return questionWeight
	}
	return 0
}

func scoreShortAnswer(userAnswer string, correctAnswer string) float64 {
	userAns := strings.TrimSpace(strings.ToLower(userAnswer))
	correctAns := strings.TrimSpace(strings.ToLower(correctAnswer))
	if userAns == correctAns {
		return questionWeight
	}
	return 0
}

func CalculateScore(userAnswers map[string]string, questions []models.Question) float64 {
	totalScore := 0.0
	maxPossibleScore := float64(len(questions)) * questionWeight
	for i, question := range questions {
		userAnswer := userAnswers["question"+strconv.Itoa(i)]
		log.Printf("Question %d: userAnswer=%s, correctAnswer=%v, type=%s", i, userAnswer, question.Answer, question.Type)
		switch question.Type {
		case models.MultipleChoice:
			totalScore += scoreMultipleChoice(userAnswer, question.Answer.(string))
		case models.TrueFalse:
			totalScore += scoreTrueFalse(userAnswer, question.Answer.(string))
		case models.ShortAnswer:
			totalScore += scoreShortAnswer(userAnswer, question.Answer.(string))
		}
	}
	log.Printf("Total Score: %f, Max Possible Score: %f", totalScore, maxPossibleScore)
	return (totalScore / maxPossibleScore) * 70 // Scale score to 0-100 range
}
