package handler

import (
	"backend_github_trending/log"
	"backend_github_trending/model"
	"backend_github_trending/repository"
	"backend_github_trending/model/req"
	"backend_github_trending/security"
	"github.com/labstack/echo"
	uuid "github.com/google/uuid"
	"net/http"
	validator "github.com/go-playground/validator/v10"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}

func (U *UserHandler) HandleSignUp(c echo.Context) error {
	req := req.ReqSignup{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	hash := security.HashAndSalt([]byte(req.Password))
	role := model.MEMBER.String()
	userId, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	dataCollection := model.User{
		UserId:    userId.String(),
		FullName:  req.FullName,
		Email:     req.Email,
		Password:  hash,
		Role:      role,
		Token:     "",
	}

	user, err := U.UserRepo.SaveUser(c.Request().Context(), dataCollection)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user.Password = ""
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Signup success",
		Data:       user,
	})
}