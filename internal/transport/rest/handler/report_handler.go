package handler

import (
	"backend/internal/database"
	"backend/internal/database/repository/report"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetFieldHarvestingReport(c *fiber.Ctx) error {
	startPeriod, err := time.Parse("2006-01-02", c.Query("start_period"))
	if err != nil {
		c.Status(500).SendString("Ошибка запроса")
	}
	endPeriod, err := time.Parse("2006-01-02", c.Query("end_period"))
	if err != nil {
		c.Status(500).SendString("Ошибка запроса")
	}
	customerId, err := strconv.ParseUint(c.Query("customer_id"), 10, 64)
	if err != nil {
		c.Status(500).SendString("Ошибка запроса")
	}
	report, err := report.GetFieldHarvestingReport(database.DB, startPeriod, endPeriod, customerId)
	if err != nil {
		c.Status(500).SendString("Ошибка базы данных")
	}
	return c.JSON(report)
}

func GetPeopleExperienceReport(c *fiber.Ctx) error {
	report, err := report.GetPeopleExperienceReport(database.DB)
	if err != nil {
		c.Status(500).SendString("Ошибка базы данных")
	}
	return c.JSON(report)
}

func GetMachinesReport(c *fiber.Ctx) error {
	report, err := report.GetMachineReport(database.DB)
	if err != nil {
		c.Status(500).SendString("Ошибка базы данных")
	}
	return c.JSON(report)
}
