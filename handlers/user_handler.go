package handlers

import (
	"errors"
	"go-gin/config"
	"go-gin/models"
	"go-gin/repositories"
	"go-gin/request"
	"go-gin/response"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type user struct {
	userRepo repositories.UserRepo
}

func NewUserHandler(userRepo repositories.UserRepo) *user {
	return &user{
		userRepo: userRepo,
	}
}

func (h *user) GetUsers(c *gin.Context) {
	users, err := h.userRepo.GetAll()
	if err != nil {
		response.InternalServerError(c)
		log.Printf("error retrieve user: %v", err)
		return
	}

	if len(users) == 0 {
		res := []response.UserResponse{}
		response.Ok(c, res, "user data")
		return
	}

	res := []response.UserResponse{}
	for _, user := range users {
		res = append(res, response.UserResponse{
			Id:   user.Id,
			Name: user.Name,
		})
	}

	response.Ok(c, res, "user data")
}

func (h *user) GetById(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	result := config.DB.Model(&models.User{}).Select("id", "name").Where("id", id).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			response.NotFound(c, "user not found")
			return
		}

		response.InternalServerError(c)
		log.Printf("error get user by id: %v", result.Error)
		return
	}

	res := response.UserResponse{
		Id:   user.Id,
		Name: user.Name,
	}

	response.Ok(c, res, "user by id")
}

func (h *user) Save(c *gin.Context) {
	var req request.UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid payload")
		log.Printf("error binding: %v", err)
		return
	}

	user := models.User{Name: req.Name}
	result := config.DB.Create(&user)

	if result.Error != nil {
		response.InternalServerError(c)
		log.Printf("error insert new user: %v", result.Error)
		return
	}

	res := response.UserResponse{
		Id:   user.Id,
		Name: user.Name,
	}

	response.Ok(c, res)
}
