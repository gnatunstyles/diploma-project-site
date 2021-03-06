package handlers

import (
	"diploma-project-site/internal/models"
	"diploma-project-site/internal/service"

	"github.com/gofiber/fiber/v2"
)

func RandomProcessingHandler(c *fiber.Ctx) error {
	req := &models.ProcessingRandRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "error. wrong request format",
		})
	}

	ans, err := service.ConvertProcRand(req.ProjectName, req.FilePath, uint(req.UserId), int(req.Factor))
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "error. something went wrong with the server.",
			"error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "OK",
		"body":    ans,
	})
}

func GridCandidateProcessingHandler(c *fiber.Ctx) error {
	req := &models.ProcessingGridRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "error. wrong request format",
		})
	}

	ans, err := service.ConvertProcCandidate(req.ProjectName, req.FilePath, uint(req.UserId), int(req.Voxel))
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "error. something went wrong with the server.",
			"error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "OK",
		"body":    ans,
	})
}

func GridBarycenterProcessingHandler(c *fiber.Ctx) error {
	req := &models.ProcessingGridRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  400,
			"message": "error. wrong request format",
		})
	}

	ans, err := service.ConvertProcBarycenter(req.ProjectName, req.FilePath, uint(req.UserId), int(req.Voxel))
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "error. something went wrong with the server.",
			"error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "OK",
		"body":    ans,
	})
}
