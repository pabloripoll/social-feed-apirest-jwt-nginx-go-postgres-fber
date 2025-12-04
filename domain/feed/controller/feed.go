package controller

import (
	"apirest/database"
	feedPostModel "apirest/domain/feed/model/feed_post"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type feedRequest struct {
	Title string `json:"title" validate:"required"`
}

func GetFeeds(c *fiber.Ctx) error {
	var feeds []feedPostModel.FeedPost

	database.DB.Find(&feeds)

	return c.JSON(feeds)
}

func GetFeed(c *fiber.Ctx) error {
	id := c.Params("id")

	var feed feedPostModel.FeedPost
	resp := database.DB.First(&feed, id)
	if resp.Error != nil {
		if resp.Error == gorm.ErrRecordNotFound {
			return c.Status(404).SendString("No feed found with ID")
		}
		return c.Status(500).SendString(resp.Error.Error())
	}

	return c.JSON(feed)
}

func CreateFeed(c *fiber.Ctx) error {
	var req feedRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// FeedPost.Title is *string in the model, so take a pointer to the local value.
	title := req.Title
	feedModel := feedPostModel.FeedPost{
		Title: &title,
	}

	resp := database.DB.Create(&feedModel)
	if resp.Error != nil {
		return c.Status(400).SendString(resp.Error.Error())
	}

	// Return the created resource (including ID)
	return c.Status(201).JSON(feedModel)
}

func UpdateFeed(c *fiber.Ctx) error {
	id := c.Params("id")

	var req feedRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	var oldFeed feedPostModel.FeedPost
	resp := database.DB.First(&oldFeed, id)
	if resp.Error != nil {
		if resp.Error == gorm.ErrRecordNotFound {
			return c.Status(404).SendString("No feed found with ID")
		}
		return c.Status(500).SendString(resp.Error.Error())
	}

	// Prepare pointer for the Title field
	title := req.Title
	if err := database.DB.Model(&oldFeed).Updates(feedPostModel.FeedPost{Title: &title}).Error; err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Reload updated record
	database.DB.First(&oldFeed, id)

	return c.JSON(oldFeed)
}

func DeleteFeed(c *fiber.Ctx) error {
	id := c.Params("id")

	var feed feedPostModel.FeedPost

	resp := database.DB.First(&feed, id)
	if resp.Error != nil {
		if resp.Error == gorm.ErrRecordNotFound {
			return c.Status(404).SendString("No feed found with ID")
		}
		return c.Status(500).SendString(resp.Error.Error())
	}

	if err := database.DB.Delete(&feed).Error; err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendString("Feed successfully deleted")
}
