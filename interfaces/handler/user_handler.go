package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raytr/sample-project-p/config"
	"github.com/raytr/sample-project-p/domain/entity"
	"github.com/raytr/sample-project-p/domain/repository"
	"github.com/raytr/sample-project-p/pkg/jwtoken"
	"github.com/raytr/sample-project-p/security"
)

type userHandler struct {
	repo      repository.UserRepository
	jwtEncode *jwtoken.JwtokenEncode
	cfg       *config.Config
}

func NewUserHandler(userRepo repository.UserRepository, cfg *config.Config) userHandler {
	return userHandler{
		repo: userRepo,
	}
}

func (h *userHandler) GetUsers(c echo.Context) error {
	dtos, err := h.repo.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, dtos)

	return nil
}

func (h *userHandler) SignUp(c echo.Context) error {
	dto := new(entity.AuthUserDTO)
	if err := c.Bind(dto); err != nil {
		return err
	}

	//validate
	if err := dto.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user := dto.FromDTO()
	user.HashPassword = security.Hash(dto.Password, h.cfg.Secret.Pepper)

	userID, err := h.repo.SaveUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	fmt.Println(userID)

	c.JSON(http.StatusOK, struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	})

	return nil
}

func (h *userHandler) Login(c echo.Context) error {
	req := new(entity.LoginReq)
	if err := c.Bind(req); err != nil {
		return err
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	//gererate

}
