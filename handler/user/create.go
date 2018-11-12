package user

import (
	"github.com/gin-gonic/gin"
)

// @Summary 创建用户
// @Description 创建一个用户
// @Accept json
// @produce json
// @Param body body user.CreateRequest true "user info"
// @Success 200 {object} user.CreateResponse
// @Router /v1/user/ [post]
func Create(c *gin.Context) {

}

type CreateRequest struct {
	Username string `json:"username" example:"string"`
	Password string `json:"password" example:"string"`
}

type CreateResponse struct {
	Username string `json:"username" example:"string"`
}
