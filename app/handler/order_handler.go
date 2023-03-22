package handler

import (
	"strconv"
	"vandyahmad24/maxsol/app/domain/entity"
	uc "vandyahmad24/maxsol/app/usecase/order_usecase"
	"vandyahmad24/maxsol/app/util"

	"github.com/gofiber/fiber/v2"
	"github.com/opentracing/opentracing-go"
)

type OrderService struct {
	svc uc.OrderUsecasePort
}

// NewIbService new Ib service
func NewOrderServiceService(svc uc.OrderUsecasePort) *OrderService {
	return &OrderService{svc: svc}
}

func (p *OrderService) CreateOrderService(c *fiber.Ctx) error {

	sp, ctx := opentracing.StartSpanFromContext(c.Context(), "CreateOrderService")
	defer sp.Finish()

	var orderInput entity.OrderInput
	if err := c.BodyParser(&orderInput); err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusBadRequest).JSON(util.ApiErrorResponse("Bad Request", err.Error()))
	}
	util.LogRequest(sp, orderInput)

	errValidate := entity.ValidateInputOrder(orderInput)
	if errValidate != nil {
		util.LogObject(sp, "ErrorValidates", errValidate)
		return c.Status(fiber.StatusBadRequest).JSON(errValidate)
	}

	res, err := p.svc.CreateOrder(ctx, orderInput)
	if err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusBadRequest).JSON(util.ApiErrorResponse("Bad Request", err.Error()))

	}

	util.LogResponse(sp, res)

	return c.Status(fiber.StatusOK).JSON(util.ApiResponse("Success Create Order", res))
}

func (p *OrderService) GetAllOrderService(c *fiber.Ctx) error {

	sp, ctx := opentracing.StartSpanFromContext(c.Context(), "GetAllOrderService")
	defer sp.Finish()

	res, err := p.svc.GetAllOrder(ctx)
	if err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusInternalServerError).JSON(util.ApiErrorResponse("Bad Request", err.Error()))
	}

	util.LogResponse(sp, res)

	return c.Status(fiber.StatusOK).JSON(util.ApiResponse("Success Get All Order", res))
}

func (p *OrderService) GetOrderService(c *fiber.Ctx) error {

	sp, ctx := opentracing.StartSpanFromContext(c.Context(), "GetOrderService")
	defer sp.Finish()

	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusInternalServerError).JSON(util.ApiErrorResponse("Param error", err.Error()))
	}
	res, err := p.svc.GetOrder(ctx, idInt)
	if err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusNotFound).JSON(util.ApiErrorResponse("Order Not Found", err.Error()))
	}

	util.LogResponse(sp, res)

	return c.Status(fiber.StatusOK).JSON(util.ApiResponse("Success Get Order", res))
}

func (p *OrderService) DeleteOrderService(c *fiber.Ctx) error {

	sp, ctx := opentracing.StartSpanFromContext(c.Context(), "DeleteOrderService")
	defer sp.Finish()

	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusInternalServerError).JSON(util.ApiErrorResponse("Param error", err.Error()))
	}

	err = p.svc.DeleteOrder(ctx, idInt)
	if err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusNotFound).JSON(util.ApiErrorResponse("Order Not Found", err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(util.ApiResponse("Success Delete Order", nil))
}

func (p *OrderService) UpdateOrderService(c *fiber.Ctx) error {

	sp, ctx := opentracing.StartSpanFromContext(c.Context(), "UpdateOrderService")
	defer sp.Finish()

	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusInternalServerError).JSON(util.ApiErrorResponse("Param error", err.Error()))
	}

	var inpuOrder entity.OrderInput
	if err := c.BodyParser(&inpuOrder); err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusBadRequest).JSON(util.ApiErrorResponse("Bad Request", err.Error()))
	}
	util.LogRequest(sp, inpuOrder)

	errValidate := entity.ValidateInputOrder(inpuOrder)
	if errValidate != nil {
		util.LogObject(sp, "ErrorValidates", errValidate)
		return c.Status(fiber.StatusBadRequest).JSON(errValidate)
	}

	res, err := p.svc.UpdateOrder(ctx, idInt, inpuOrder)
	if err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusBadRequest).JSON(util.ApiErrorResponse("Failed Update Order", err.Error()))
	}
	util.LogResponse(sp, res)

	return c.Status(fiber.StatusOK).JSON(util.ApiResponse("Success Update Order", res))
}

func (p *OrderService) CreateBulkOrderService(c *fiber.Ctx) error {

	sp, ctx := opentracing.StartSpanFromContext(c.Context(), "CreateOrderService")
	defer sp.Finish()

	var orderInput entity.OrderInputBulk
	if err := c.BodyParser(&orderInput); err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusBadRequest).JSON(util.ApiErrorResponse("Bad Request", err.Error()))
	}
	util.LogRequest(sp, orderInput)

	errValidate := entity.ValidateInputBulk(orderInput)
	if errValidate != nil {
		util.LogObject(sp, "ErrorValidates", errValidate)
		return c.Status(fiber.StatusBadRequest).JSON(errValidate)
	}
	//
	res, err := p.svc.CreateOrderBulk(ctx, orderInput)
	if err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusBadRequest).JSON(util.ApiErrorResponse("Bad Request", err.Error()))

	}
	//
	util.LogResponse(sp, res)

	return c.Status(fiber.StatusOK).JSON(util.ApiResponse("Success Create Order", res))
}
