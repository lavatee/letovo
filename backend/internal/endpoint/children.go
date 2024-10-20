package endpoint

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TakeChildInput struct {
	ChildId         string `json:"child_id"`
	Code            string `json:"code"`
	UserEmail       string `json:"email"`
	UserFirstName   string `json:"first_name"`
	UserLastName    string `json:"last_name"`
	UserClass       string `json:"class"`
	UserPhoneNumber string `json:"phone_number"`
	UserTelegram    string `json:"telegram"`
}

func (e *Endpoint) TakeChild(c *gin.Context) {
	var input TakeChildInput
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		log.Print(err.Error())
		return
	}
	intId, err := strconv.Atoi(input.ChildId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		log.Print(err.Error())
		return
	}
	fullname, gift, err := e.Services.Children.TakeChild(intId, input.UserEmail, input.UserFirstName, input.UserLastName, input.UserPhoneNumber, input.UserTelegram, input.UserClass, input.Code)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		log.Print(err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"child_fullname": fullname,
		"child_gift":     gift,
	})
}

func (e *Endpoint) GetChildren(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"children": e.Services.Children.GetChildren(),
	})
}

type SendCodeInput struct {
	UserEmail string `json:"email"`
}

func (e *Endpoint) SendCode(c *gin.Context) {
	var input SendCodeInput
	if err := c.BindJSON(&input); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println(input.UserEmail)
	if err := e.Services.Children.SendCode(input.UserEmail); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

type NewAdminInput struct {
	Email string `json:"email"`
}

func (e *Endpoint) NewAdmin(c *gin.Context) {
	var input NewAdminInput
	if err := c.BindJSON(&input); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	if err := e.Services.Children.NewAdmin(input.Email); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

type GetChildrenInfoInput struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func (e *Endpoint) GetChildrenInfo(c *gin.Context) {
	var input GetChildrenInfoInput
	if err := c.BindJSON(&input); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	children, err := e.Services.Children.GetChildrenInfo(input.Email, input.Code)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"children": children,
	})
}
