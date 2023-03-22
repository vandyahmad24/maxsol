package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opentracing/opentracing-go"
	"time"
	"vandyahmad24/maxsol/app/domain/entity"
	"vandyahmad24/maxsol/app/model"
	jwtUc "vandyahmad24/maxsol/app/usecase/jwt_usecase"
	uc "vandyahmad24/maxsol/app/usecase/user_usecase"
	"vandyahmad24/maxsol/app/util"
)

type UserService struct {
	svc     uc.UserUsecasePort
	authSvc jwtUc.AuthServicePort
}

// NewIbService new Ib service
func NewUserServiceService(svc uc.UserUsecasePort, authSvc jwtUc.AuthServicePort) *UserService {
	return &UserService{svc: svc, authSvc: authSvc}
}

func (p *UserService) RegisterService(c *fiber.Ctx) error {

	sp, ctx := opentracing.StartSpanFromContext(c.Context(), "RegisterService")
	defer sp.Finish()

	var userInput entity.UserInput
	if err := c.BodyParser(&userInput); err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusBadRequest).JSON(util.ApiErrorResponse("Bad Request", err.Error()))
	}
	util.LogRequest(sp, userInput)

	errValidate := entity.ValidateUserInput(userInput)
	if errValidate != nil {
		util.LogObject(sp, "ErrorValidates", errValidate)
		return c.Status(fiber.StatusBadRequest).JSON(errValidate)
	}

	res, err := p.svc.RegisterUser(ctx, model.User{
		Name:     userInput.Name,
		Password: userInput.Password,
	})
	if err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusInternalServerError).JSON(util.ApiErrorResponse("Bad Request", err.Error()))
	}

	util.LogResponse(sp, res)

	response := entity.UserResponse{
		Id:        res.Id,
		Name:      res.Name,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	return c.Status(fiber.StatusOK).JSON(util.ApiResponse("Success Register", response))
}

func (p *UserService) LoginService(c *fiber.Ctx) error {

	sp, ctx := opentracing.StartSpanFromContext(c.Context(), "RegisterService")
	defer sp.Finish()

	var userInput entity.UserInput
	if err := c.BodyParser(&userInput); err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusBadRequest).JSON(util.ApiErrorResponse("Bad Request", err.Error()))
	}
	util.LogRequest(sp, userInput)

	errValidate := entity.ValidateUserInput(userInput)
	if errValidate != nil {
		util.LogObject(sp, "ErrorValidates", errValidate)
		return c.Status(fiber.StatusBadRequest).JSON(errValidate)
	}

	res, err := p.svc.LoginUsernamePass(ctx, model.User{
		Name:     userInput.Name,
		Password: userInput.Password,
	})
	if err != nil {
		util.LogError(sp, err)
		return c.Status(fiber.StatusInternalServerError).JSON(util.ApiErrorResponse("Bad Request", err.Error()))
	}

	util.LogResponse(sp, res)

	token, err := p.authSvc.GenerateToken(res)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(util.ApiErrorResponse("Bad Request", err.Error()))
	}

	response := entity.UserResponse{
		Id:        res.Id,
		Name:      res.Name,
		Token:     token,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(util.ApiResponse("Success Register", response))
}
