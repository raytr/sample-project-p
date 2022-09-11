package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raytr/sample-project-p/config"
	"github.com/raytr/sample-project-p/domain/entity"
	"github.com/raytr/sample-project-p/domain/repository"
	"github.com/raytr/sample-project-p/security"
)

type userHandler struct {
	repo repository.UserRepository
	cfg  *config.Config
}

func NewUserHandler(userRepo repository.UserRepository, cfg *config.Config) userHandler {
	return userHandler{
		repo: userRepo,
	}
}

func (h *userHandler) GetUsers(c echo.Context) error {
	log.Println("GET TALENTS")
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
	log.Println(dto)

	//validate
	if err := dto.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	user := dto.FromDTO()
	user.HashPassword = security.Hash(dto.Password, h.cfg.Secret.Pepper)

	userID, err := h.repo.SaveUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	fmt.Println(userID)

	c.JSON(http.StatusOK, struct {
		ID entity.ID `json:"id"`
	}{
		//ID: userID,
	})

	return nil
}
