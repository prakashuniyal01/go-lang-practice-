package handlers

import (
	"fmt"
	"net/http"
	"trademarkia/common"
	"trademarkia/models"
	"trademarkia/repository"
	"trademarkia/services"

	"github.com/gin-gonic/gin"
)

func UserSignup(c *gin.Context) {
	var input models.UserData
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error parsing json": err.Error()})
		return
	}
	if input.FirstName == "" || input.LastName == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error parsing json": "Please provide a valid firstname and lastname"})
		return
	}
	if input.Age <= 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error parsing json": "Please provide a valid age"})
		return
	}
	if input.Email == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error parsing json": "Please provide a valid email"})
		return
	}
	userData := repository.UserData{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Age:       input.Age,
		Married:   input.Married,
	}
	if err := services.UserSignup(userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error in signup": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "User data added succesfully"})
	return
}

func UserLogin(c *gin.Context) {
	var input models.LoginData
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error parsing json": err.Error()})
		return
	}
	response, err := services.UserLogin(input.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error in login": err.Error()})
		return
	}
	userData := models.UserData{
		FirstName: response.FirstName,
		LastName:  response.LastName,
		Email:     input.Email,
		Age:       response.Age,
		Married:   response.Married,
	}
	common.CreateJwtToken(userData, c)
	c.JSON(http.StatusAccepted, gin.H{"message": "User loged in succesfully"})
	return
}

func GetUserData(c *gin.Context) {
	email := c.GetString("userEmail")
	fmt.Println(email, "mail")

	response, err := services.GetUserData(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusFound, gin.H{"User Data": response})
}

func EditUserData(c *gin.Context) {
	email := c.GetString("userEmail")
	var input models.UserEditData
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error parsing json": err.Error()})
		return
	}
	userData := repository.UserData{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     email,
		Age:       input.Age,
		Married:   input.Married,
	}
	if err := services.EditUserData(userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error in edit data": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "User data updated succesfully"})
	return
}

func DeleteUserData(c *gin.Context) {
	email := c.GetString("userEmail")

	if err := services.DeleteUserData(email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error in delete data": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "User data deleted succesfully"})
	return
}

func Logout(c *gin.Context) {
	common.DeleteCookie(c, "Authorise")
	c.JSON(http.StatusAccepted, gin.H{"message": "User logedout succesfully"})
	return
}
