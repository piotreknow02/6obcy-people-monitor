package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/piotreknow02/6obcy-people-monitor/server/models"
)

func (base *Controller) GetDay(c *fiber.Ctx) error {
	var logs []models.Log
	res := base.DB.Where("created_at > datetime('now', '-1 days')").Find(&logs)
	if res.Error != nil {
		return res.Error
	}
	c.JSON(logs)
	return nil
}

func (base *Controller) GetWeek(c *fiber.Ctx) error {
	var logs []models.Log
	res := base.DB.Where("created_at > datetime('now', '-7 days')").Find(&logs)
	if res.Error != nil {
		return res.Error
	}
	c.JSON(logs)
	return nil
}

func (base *Controller) GetMonth(c *fiber.Ctx) error {
	var logs []models.Log
	res := base.DB.Where("created_at > datetime('now', '-1 months')").Find(&logs)
	if res.Error != nil {
		return res.Error
	}
	c.JSON(logs)
	return nil
}

func (base *Controller) GetAll(c *fiber.Ctx) error {
	var logs []models.Log
	res := base.DB.Find(&logs)
	if res.Error != nil {
		return res.Error
	}
	c.JSON(logs)
	return nil
}
