package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Jupita-SA/cbt/utils"
	"strconv"
)

func GetResult(c *fiber.Ctx) error {
	exam := c.Query("exam")
	userAnswers := make(map[string]string)
	for i := range exams[exam] {
		userAnswers["question"+strconv.Itoa(i)] = c.Cookies(exam + "question" + strconv.Itoa(i))
	}
	score := utils.CalculateScore(userAnswers, exams[exam])
	return c.SendString("Your score for " + exam + " is: " + strconv.FormatFloat(score, 'f', 2, 64) + "%")
}
