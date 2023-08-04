package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/piotreknow02/6obcy-people-monitor/server/config"
	"github.com/piotreknow02/6obcy-people-monitor/server/controllers"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) *fiber.App {
	a := fiber.New()

	// Middleware
	a.Use(logger.New())
	a.Use(recover.New())
	a.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	a.Use(cors.New(cors.Config{
		AllowOrigins: config.Conf.AllowOrigins,
		AllowMethods: "GET, OPTIONS",
		AllowHeaders: "Origin, Host, Content-Type, Accept",
	}))
	a.Get("/monitor", monitor.New(monitor.Config{
		Title: "6obcy people monitor stats",
	}))
	
	// Controller
	api := controllers.Controller{DB: db}

	// Routing
	log := a.Group("/log")
	{
		log.Get("/day", api.GetDay)
		log.Get("/week", api.GetWeek)
		log.Get("/month", api.GetMonth)
		log.Get("/all", api.GetAll)
		// TODO: Get Range
		// log.Post("/range", )
	}
	
	return a
}
