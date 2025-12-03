package controller

import (
	"apirest/database"
	"apirest/domain/feed/model"

	"github.com/gofiber/fiber/v2"
)

type feedRequest struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
}

func GetFeeds(c *fiber.Ctx) error {
	var feeds []model.Feed

	database.DB.Find(&feeds)

	return c.JSON(feeds)
}

func GetFeed(c *fiber.Ctx) error {
	id := c.Params("id")

	var feed model.Feed
	database.DB.Find(&feed, id)

	return c.JSON(feed)
}

func CreateFeed(c *fiber.Ctx) error {
	var feed feedRequest
	if err := c.BodyParser(&feed); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	feedModel := model.Feed{
		Title:  feed.Title,
		Author: feed.Author,
	}

	resp := database.DB.Create(&feedModel)
	if resp.Error != nil {
		return c.Status(400).SendString(resp.Error.Error())
	}

	return c.JSON(feed)
}

func UpdateFeed(c *fiber.Ctx) error {
	id := c.Params("id")

	var feed feedRequest
	if err := c.BodyParser(&feed); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	var oldFeed model.Feed
	database.DB.Find(&oldFeed, id)
	if oldFeed.Title == "" {
		return c.Status(404).SendString("No feed found with ID")
	}

	database.DB.Model(&oldFeed).Updates(feed)

	return c.JSON(oldFeed)
}

func DeleteFeed(c *fiber.Ctx) error {
	id := c.Params("id")

	var feed model.Feed

	database.DB.Find(&feed, id)
	if feed.Title == "" {
		return c.Status(404).SendString("No feed found with ID")
	}

	database.DB.Delete(&feed)

	return c.SendString("Feed successfully deleted")
}
