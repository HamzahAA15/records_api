package records

import (
	"errors"
	"fmt"

	"github.com/HamzahAA15/records_api/utils/errorMap"
	"github.com/gofiber/fiber/v2"
)

type RecordsHandler struct {
	recordsService RecordsService
}

func NewRecordsHandler(recordService RecordsService) *RecordsHandler {
	return &RecordsHandler{recordsService: recordService}
}

func (rh *RecordsHandler) GetAllRecords(c *fiber.Ctx) error {
	requestPayload := RequestPayload{}
	err := c.QueryParser(&requestPayload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponsePayload{
			Code: fiber.StatusBadRequest,
			Msg:  fmt.Sprintf("failed to parse query params: %s", err.Error()),
		})
	}

	err = validateReqeust(requestPayload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponsePayload{
			Code: fiber.StatusBadRequest,
			Msg:  fmt.Sprintf("check the client payload: %s", err.Error()),
		})
	}

	records, err := rh.recordsService.GetFilteredRecords(c.Context(), requestPayload)
	if err != nil {
		if errors.Is(err, errorMap.RecordsNotFound) {
			return c.Status(fiber.ErrNotFound.Code).JSON(ResponsePayload{
				Code: fiber.ErrNotFound.Code,
				Msg:  fmt.Sprintf("error [GetFilteredRecords]: %s", err.Error()),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(ResponsePayload{
			Code: fiber.StatusInternalServerError,
			Msg:  fmt.Sprintf("failed to get records: %s", err.Error()),
		})
	}

	return c.Status(fiber.StatusOK).JSON(ResponsePayload{
		Code:    0,
		Msg:     "success",
		Records: records,
	})
}

func validateReqeust(req RequestPayload) error {
	if req.EndDate == req.StartDate || req.StartDate > req.EndDate || req.EndDate < req.StartDate || req.StartDate == "" || req.EndDate == "" {
		return errors.New("input date is not valid")
	}

	if req.MinCount == req.MaxCount || req.MinCount == 0 || req.MaxCount == 0 || req.MinCount > req.MaxCount || req.MaxCount < req.MinCount {
		return errors.New("input count is not valid")
	}

	return nil
}

// setup routes for records
func (rh *RecordsHandler) SetupRoutes(app *fiber.App) {
	recordsGroup := app.Group("/api/records")
	recordsGroup.Get("/", rh.GetAllRecords)
}
