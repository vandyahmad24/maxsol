package handler

import (
	"strconv"
	"vandyahmad24/maxsol/app/domain/entity"
	uc "vandyahmad24/maxsol/app/usecase/cake_usecase"
	"vandyahmad24/maxsol/app/util"

	"github.com/gofiber/fiber/v2"
	"github.com/opentracing/opentracing-go"
)

type CakeService struct {
	svc uc.CakeUsecasePort
}

// NewIbService new Ib service
func NewCakeServiceService(svc uc.CakeUsecasePort) *CakeService {
	return &CakeService{svc: svc}
}

func (p *CakeService) CreateCakeService(c *fiber.Ctx) error {

	sp, ctx := opentracing.StartSpanFromContext(c.Context(), "CreateCakeService")
	defer sp.Finish()

	var inpuCake entity.CakeInput
	if err := c.BodyParser(&inpuCake); err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusBadRequest).JSON(util.ApiErrorResponse("Bad Request", err.Error()))
	}
	util.LogRequest(sp, inpuCake)

	errValidate := entity.ValidateInputCake(inpuCake)
	if errValidate != nil {
		util.LogObject(sp, "ErrorValidates", errValidate)
		return c.Status(fiber.StatusBadRequest).JSON(errValidate)
	}

	res, err := p.svc.CreateCake(ctx, inpuCake)
	if err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusBadRequest).JSON(util.ApiErrorResponse("Bad Request", err.Error()))

	}

	util.LogResponse(sp, res)

	return c.Status(fiber.StatusOK).JSON(util.ApiResponse("Success Create Cake", res))
}

func (p *CakeService) GetAllCakeService(c *fiber.Ctx) error {

	sp, ctx := opentracing.StartSpanFromContext(c.Context(), "GetAllCakeService")
	defer sp.Finish()

	res, err := p.svc.GetAllCake(ctx)
	if err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusInternalServerError).JSON(util.ApiErrorResponse("Bad Request", err.Error()))
	}

	util.LogResponse(sp, res)

	return c.Status(fiber.StatusOK).JSON(util.ApiResponse("Success Get All Cake", res))
}

func (p *CakeService) GetCakeService(c *fiber.Ctx) error {

	sp, ctx := opentracing.StartSpanFromContext(c.Context(), "GetCakeService")
	defer sp.Finish()

	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusInternalServerError).JSON(util.ApiErrorResponse("Param error", err.Error()))
	}
	res, err := p.svc.GetCake(ctx, idInt)
	if err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusNotFound).JSON(util.ApiErrorResponse("Cake Not Found", err.Error()))
	}

	util.LogResponse(sp, res)

	return c.Status(fiber.StatusOK).JSON(util.ApiResponse("Success Get Cake", res))
}

func (p *CakeService) DeleteCakeService(c *fiber.Ctx) error {

	sp, ctx := opentracing.StartSpanFromContext(c.Context(), "DeleteCakeService")
	defer sp.Finish()

	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusInternalServerError).JSON(util.ApiErrorResponse("Param error", err.Error()))
	}

	err = p.svc.DeleteCake(ctx, idInt)
	if err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusNotFound).JSON(util.ApiErrorResponse("Cake Not Found", err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(util.ApiResponse("Success Delete Cake", nil))
}

func (p *CakeService) UpdateCakeService(c *fiber.Ctx) error {

	sp, ctx := opentracing.StartSpanFromContext(c.Context(), "UpdateCakeService")
	defer sp.Finish()

	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusInternalServerError).JSON(util.ApiErrorResponse("Param error", err.Error()))
	}

	var inpuCake entity.CakeInput
	if err := c.BodyParser(&inpuCake); err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusBadRequest).JSON(util.ApiErrorResponse("Bad Request", err.Error()))
	}
	util.LogRequest(sp, inpuCake)

	errValidate := entity.ValidateInputCake(inpuCake)
	if errValidate != nil {
		util.LogObject(sp, "ErrorValidates", errValidate)
		return c.Status(fiber.StatusBadRequest).JSON(errValidate)
	}

	res, err := p.svc.UpdateCake(ctx, idInt, inpuCake)
	if err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusBadRequest).JSON(util.ApiErrorResponse("Failed Update Cake", err.Error()))
	}
	util.LogResponse(sp, res)

	return c.Status(fiber.StatusOK).JSON(util.ApiResponse("Success Update Cake", res))
}
