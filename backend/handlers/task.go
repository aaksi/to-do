package handlers

import "github.com/gofiber/fiber/v3"

type Task struct {
	Title         string `json:"title"`
	IsTimed       bool   `json:"is_timed"`
	AllocatedTime int    `json:"allocated_time"`
	ParentID      *uint  `json:"parent_id"`
}

func TaskHandler(c fiber.Ctx) {

}
