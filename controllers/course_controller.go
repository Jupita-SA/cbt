package controllers

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/Jupita-SA/cbt/utils"
)

func GetCourses(c *fiber.Ctx) error {
	rows, err := utils.DB.Query("SELECT id, name, description FROM courses")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	defer rows.Close()

	var courses []map[string]interface{}
	for rows.Next() {
		var id int64
		var name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		courses = append(courses, map[string]interface{}{
			"id":          id,
			"name":        name,
			"description": description,
		})
	}

	return c.Render("admin/pages/manage_course", fiber.Map{
		"Courses": courses,
	})
}

func AddCourse(c *fiber.Ctx) error {
	name := c.FormValue("name")
	description := c.FormValue("description")

	_, err := utils.DB.Exec("INSERT INTO courses (name, description) VALUES ($1, $2)", name, description)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Redirect("/manage_course")
}

func EditCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	name := c.FormValue("name")
	description := c.FormValue("description")

	log.Printf("Received updated data for course ID %s: Name=%s, Description=%s", id, name, description)

	_, err := utils.DB.Exec("UPDATE courses SET name=$1, description=$2 WHERE id=$3", name, description, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Redirect("/manage_course")
}

func DeleteCourse(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := utils.DB.Exec("DELETE FROM courses WHERE id=$1", id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Redirect("/manage_course")
}
