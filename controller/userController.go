package controller

import (
	"strconv"

	helper "github.com/Allifiando/go-gin-boilerplate/helper"
	"github.com/Allifiando/go-gin-boilerplate/middleware"
	"github.com/Allifiando/go-gin-boilerplate/model"
	"github.com/Allifiando/go-gin-boilerplate/model/request"
	responses "github.com/Allifiando/go-gin-boilerplate/model/response"
	Error "github.com/Allifiando/go-gin-boilerplate/pkg/error"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct{}

// ListUser ...
func (U *User) ListUser(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200
	limit, err := strconv.Atoi(c.Query("limit"))
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		Error.Error(err)
	}
	if page == 0 {
		page = helper.DefaultPage
	}
	limit, offset := helper.PaginationPageOffset(page, limit)
	userModel := model.UserModel{}
	data, count, err := userModel.GetAll(offset, limit)
	if err != nil {
		statusCode = 400
		Error.Error(err)
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helper.OutputAPIResponseWithPayload(errorParams))
		return
	}
	pagination := helper.PaginationRes(page, count, limit)
	params := map[string]interface{}{
		"payload": data,
		"meta":    pagination,
	}
	c.JSON(statusCode, helper.OutputAPIResponseWithPayload(params))
	return
}

// FindOneByUUID ...
func (U *User) FindOneByUUID(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200
	uuid := c.Params.ByName("uuid")
	userModel := model.UserModel{}
	data, err := userModel.FindByUUID(uuid)
	if err != nil {
		statusCode = 400
		Error.Error(err)
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helper.OutputAPIResponseWithPayload(errorParams))
		return
	}
	params := map[string]interface{}{
		"meta":    "success",
		"payload": data,
	}
	c.JSON(statusCode, helper.OutputAPIResponseWithPayload(params))
	return
}

// Login ...
func (U *User) Login(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200

	var body request.Login

	// Validation req body
	err := c.ShouldBindJSON(&body)
	if err != nil {
		statusCode = 406
		Error.Error(err)
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helper.OutputAPIResponseWithPayload(errorParams))
		return
	}

	userModel := model.UserModel{}
	data, err := userModel.FindByEmail(body.Email)

	if err != nil {
		statusCode = 400
		Error.Error(err)
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helper.OutputAPIResponseWithPayload(errorParams))
		return
	}

	// Compare Password
	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(body.Password))
	if err != nil {
		statusCode = 400
		Error.Error(err)
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helper.OutputAPIResponseWithPayload(errorParams))
		return
	}

	// Generate Token
	token, err := middleware.CreateToken(data)
	if err != nil {
		statusCode = 400
		Error.Error(err)
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helper.OutputAPIResponseWithPayload(errorParams))
		return
	}

	// data.Password = ""

	params := map[string]interface{}{
		"meta": map[string]interface{}{
			"token": token,
		},
		"payload": data,
	}
	c.JSON(statusCode, helper.OutputAPIResponseWithPayload(params))
	return
}

// Register ...
func (U *User) Register(c *gin.Context) {
	errorParams := map[string]interface{}{}
	statusCode := 200

	var body request.Register

	// Validation req body
	err := c.ShouldBindJSON(&body)
	if err != nil {
		statusCode = 406
		Error.Error(err)
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helper.OutputAPIResponseWithPayload(errorParams))
		return
	}

	// Set Request
	req := request.Register{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	}

	userModel := model.UserModel{}
	userData, _ := userModel.FindByEmail(body.Email)

	if userData.Email != "" {
		statusCode = 400
		Error.Error(err)
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": "User already exist",
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helper.OutputAPIResponseWithPayload(errorParams))
		return
	}

	data, uid, err := userModel.Create(req)
	if err != nil {
		statusCode = 400
		Error.Error(err)
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": "Error creating user",
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helper.OutputAPIResponseWithPayload(errorParams))
		return
	}

	// Hash Password
	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		statusCode = 406
		Error.Error(err)
		errorParams["meta"] = map[string]interface{}{
			"status":  statusCode,
			"message": "Error hashing password",
		}
		errorParams["code"] = statusCode
		c.JSON(statusCode, helper.OutputAPIResponseWithPayload(errorParams))
		return
	}

	// Set Response
	res := responses.UserModel{
		ID:       data,
		UUID:     uid,
		Name:     body.Name,
		Email:    body.Email,
		Password: string(bytes),
	}

	params := map[string]interface{}{
		"meta":    "success",
		"payload": res,
	}
	c.JSON(statusCode, helper.OutputAPIResponseWithPayload(params))
	return
}
