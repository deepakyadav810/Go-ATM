package controller

import (
	"net/http"
	"strconv"
	"strings"

	"atm-machine.com/atm-apis/models"
	"atm-machine.com/atm-apis/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

func inttostring(x int) string {
	var num string = strconv.Itoa(x)
	return num
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	acc, err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Generated Account Number": acc})
}

func (uc *UserController) DepositWithdraw(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	strlist := []string{inttostring(user.AccountNo), user.Balance}
	err := uc.UserService.DepositWithdraw(strlist)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) Transfer(ctx *gin.Context) {
	username := ctx.Param("num")
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	strlist1 := []string{inttostring(user.AccountNo), user.Balance}
	strlist2 := []string{username, strings.ReplaceAll(user.Balance, "-", "+")}
	err1 := uc.UserService.DepositWithdraw(strlist1)
	err2 := uc.UserService.DepositWithdraw(strlist2)
	if err1 != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err1.Error()})
		return
	}
	if err2 != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err2.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (uc *UserController) ChangePin(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	strlist := []string{inttostring(user.AccountNo), inttostring(user.Pin)}
	err := uc.UserService.ChangePin(strlist)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "pin changed successfully"})
}

func (uc *UserController) GetTransacion(ctx *gin.Context) {
	username := ctx.Param("num")
	user, err := uc.UserService.GetTransacion(username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user.Statement)
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/create", uc.CreateUser)
	userroute.PATCH("/update", uc.DepositWithdraw)
	userroute.PATCH("/transfer/:num", uc.Transfer)
	userroute.PATCH("/updatepin", uc.ChangePin)
	userroute.GET("/get/:num", uc.GetTransacion)
}
