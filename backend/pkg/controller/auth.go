package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"recipinder/pkg/model/dao"
	"recipinder/pkg/model/dto"
	"recipinder/pkg/view"
)

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {

		var sur dto.SignUpRequest
		if err := c.BindJSON(&sur); err != nil {
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"Wrong request body",
			)
			return
		}

		client := dao.MakeSignUpClient()
		_, err := client.Request(sur)
		if err != nil {
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				err.Error(),
			)
			return
		}

		c.JSON(http.StatusOK, view.ReturnSignResponse())
	}
}
