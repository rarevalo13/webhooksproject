package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/fiber/v2/middleware/logger"
)


type ValidationEvent struct {
	PlainToken string `json:"plainToken,omitempty"`
	Encrypted string `json:"encryptedToken,omitempty"`
}

type Response struct {
	Message ValidationEvent `json:"message,omitempty"`
	Status int `json:"status"`
}

type Event struct {
	
	Event string `json:"event"`
	EventTS int  `json:"event_ts"`
	Payload Payload `json:"payload"`
	
	
}

type Payload struct {
	AccountID string `json:"account_id,omitempty"`
	Operator string `json:"operator,omitempty"`
	OperatorID string `json:"operator_id,omitempty"`
	PlainToken string `json:"plainToken,omitempty"`
	Object Object `json:"object,omitempty"`


 
}

type Object struct {
	UUID string `json:"uuid,omitempty"`
	Id int `json:"id,omitempty"`
	HostId string `json:"host_id,omitempty"`
	Topic string `json:"topic,omitempty"`
	Type int `json:"type,omitempty"`
	TimeStamp time.Time `json:"event_ts,omitempty"`
	Duration int `json:"duration,omitempty"`
	Timezone string `json:"timezone,omitempty"`
	JoinURL string `json:"join_url,omitempty"`
	
}

func Server() {
	app := fiber.New()

// app.Use(logger.New(logger.Config{
// 	Format: " ${status}\n",
// }))

	app.Static("/", "./frontend/build")

	app.Get("/home", func(ctx *fiber.Ctx) error {

		
		return ctx.SendFile("./frontend/build/index.html")
	})

	SetupRoutes(app)

	app.Listen(":4000")
}

func SetupRoutes(app *fiber.App) {
	app.Post("/webhook", Webhooks)
	app.Get("/token", TokenGeneration)
	app.Post("/newmeeting", CreateMeeting)
	
} 