package handler

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/auth_service/config"
	"github.com/auth_service/models"
	"github.com/auth_service/pkg/helper"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @ID login
// @Router /login [POST]
// @Summary Create Login
// @Description Create Login
// @Tags Login
// @Accept json
// @Produce json
// @Param Login body models.Login true "LoginRequestBody"
// @Success 201 {object} models.LoginResponse "GetLoginBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) Login(c *gin.Context) {
	var login models.Login

	err := c.ShouldBindJSON(&login)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.storage.User().GetByPKey(
		context.Background(),
		&models.UserPrimarKey{Login: login.Login},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	if login.Password != resp.Password {
		c.JSON(http.StatusInternalServerError, errors.New("error password is not correct").Error())
		return
	}

	data := map[string]interface{}{
		"id": resp.Id,
	}

	token, err := helper.GenerateJWT(data, config.TimeExpiredAt, h.cfg.AuthSecretKey)
	if err != nil {
		log.Printf("error whiling GenerateJWT: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GenerateJWT").Error())
		return
	}

	c.JSON(http.StatusCreated, models.LoginResponse{AccessToken: token})
}

// Register godoc
// @ID register
// @Router /register [POST]
// @Summary Create Register
// @Description Create Register
// @Tags Register
// @Accept json
// @Produce json
// @Param Regester body models.Register true "RegisterRequestBody"
// @Success 201 {object} models.RegisterResponse "GetRegisterBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) Register(c *gin.Context) {
	var register models.Register

	err := c.ShouldBindJSON(&register)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.User().Create(
		context.Background(),
		&models.CreateUser{
			FirstName:   register.FirstName,
			LastName:    register.LastName,
			Login:       register.Login,
			Password:    register.Password,
			PhoneNumber: register.PhoneNumber,
		},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	data := map[string]interface{}{
		"id": id,
	}

	token, err := helper.GenerateJWT(data, config.TimeExpiredAt, h.cfg.AuthSecretKey)
	if err != nil {
		log.Printf("error whiling GenerateJWT: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GenerateJWT").Error())
		return
	}

	c.JSON(http.StatusCreated, models.RegisterResponse{AccessToken: token})
}
