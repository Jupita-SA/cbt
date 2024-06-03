package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Jupita-SA/cbt/models"
	"strconv"
)

var exams = map[string][]Question{
	"Mathematics": {
		{
			Text:    "What is the capital of France?",
			Choices: []string{"Berlin", "Madrid", "Paris", "Rome"},
			Answer:  "Paris",
			Type:    MultipleChoice,
		},
		{
			Text:    "Identify this flag:",
			Image:   "/static/images/franceflag.jpg",
			Choices: []string{"Germany", "Italy", "France", "Spain"},
			Answer:  "France",
			Type:    MultipleChoice,
		},
		{
			Text:    "What is 2 + 2?",
			Choices: []string{"3", "4", "5", "6"},
			Answer:  "4",
			Type:    MultipleChoice,
		},
		{
			Text:    "The Earth is flat. (true/false)",
			Choices: []string{"true", "false"},
			Answer:  "False",
			Type:    TrueFalse,
		},
	},
	"GST101(ECO)": {
		{
			Text:    "What is 4 + 5?",
			Choices: []string{"3", "4", "11", "6"},
			Answer:  "11",
			Type:    MultipleChoice,
		},
		{
			Text:    "An egg is spherical. (true/false)",
			Choices: []string{"true", "false"},
			Answer:  "True",
			Type:    TrueFalse,
		},
		{
			Text:    "What is the chemical symbol for water?",
			Choices: []string{},
			Answer:  "H2O",
			Type:    ShortAnswer,
		},
	},
}

func GetQuestion(c *fiber.Ctx) error {
	exam := c.Query("exam")
	index, err := strconv.Atoi(c.Query("index"))
	if err != nil || index < 0 || index >= len(exams[exam]) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid question index")
	}

	var questionNumbers []int
	for i := 0; i < len(exams[exam]); i++ {
		questionNumbers = append(questionNumbers, i)
	}

	if c.Get("Accept") == "application/json" {
		return c.JSON(fiber.Map{
			"index":           index,
			"total":           len(exams[exam]),
			"question":        exams[exam][index],
			"questionNumbers": questionNumbers,
			"exam":            exam,
			"exams":           exams,
		})
	}

	return c.Render("question", fiber.Map{
		"Index":           index,
		"Total":           len(exams[exam]),
		"Question":        exams[exam][index],
		"QuestionNumbers": questionNumbers,
		"Exam":            exam,
		"Exams":           exams,
	})
}

func PostQuestion(c *fiber.Ctx) error {
	exam := c.FormValue("exam")
	index, err := strconv.Atoi(c.FormValue("index"))
	if err != nil || index < 0 || index >= len(exams[exam]) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid question index")
	}

	var questionNumbers []int
	for i := 0; i < len(exams[exam]); i++ {
		questionNumbers = append(questionNumbers, i)
	}

	userAnswer := c.FormValue("answer")
	c.Cookie(&fiber.Cookie{
		Name:  exam + "question" + strconv.Itoa(index),
		Value: userAnswer,
	})

	navigation := c.FormValue("navigation")
	if navigation == "prev" && index > 0 {
		return c.JSON(fiber.Map{
			"index":           index - 1,
			"question":        exams[exam][index-1],
			"total":           len(exams[exam]),
			"questionNumbers": questionNumbers,
			"exam":            exam,
		})
	}
	if navigation == "next" && index+1 < len(exams[exam]) {
		return c.JSON(fiber.Map{
			"index":           index + 1,
			"question":        exams[exam][index+1],
			"total":           len(exams[exam]),
			"questionNumbers": questionNumbers,
			"exam":            exam,
		})
	}
	if navigation == "submit" {
		if index == len(exams[exam])-1 {
			// Ensure the last answer is saved
			c.Cookie(&fiber.Cookie{
				Name:  exam + "question" + strconv.Itoa(index),
				Value: userAnswer,
			})
			// Redirect to the result page
			return c.Redirect("/result?exam=" + exam)
		}
		return c.JSON(fiber.Map{
			"index":           index + 1,
			"question":        exams[exam][index+1],
			"total":           len(exams[exam]),
			"questionNumbers": questionNumbers,
			"exam":            exam,
		})
	}
	return c.Status(fiber.StatusBadRequest).SendString("Invalid navigation action")
}
