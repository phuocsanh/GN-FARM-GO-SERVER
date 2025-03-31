package account

import (
	"go_ecommerce/global"
	"go_ecommerce/internal/model"
	"go_ecommerce/internal/service"
	"go_ecommerce/internal/utils/auth"
	"go_ecommerce/pkg/response"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// management controller login user

var Login = new(cUserLogin)

type cUserLogin struct {
}

func (c *cUserLogin) RefreshToken(ctx *gin.Context) {
	// Lấy Authorization header
	refreshToken, _ := auth.ExtractBearerToken(ctx)
	log.Println("refreshToken", refreshToken)

	// Kiểm tra nếu token rỗng
	if refreshToken == "" {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, "Refresh token is empty")
		return
	}

	params := model.RefreshTokenInput{
		RefreshToken: refreshToken,
	}

	// Gọi service xử lý Refresh Token
	codeRs, dataRs, err := service.UserLogin().RefreshToken(ctx, &params)
	log.Println("errr 39 %v", err)
	log.Println("dataRs %v", dataRs.AccessToken)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	response.SuccessResponse(ctx, codeRs, dataRs)
}

// User Login
// @Summary      User Login
// @Description  User Login
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body model.LoginInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/login [post]
func (c *cUserLogin) Login(ctx *gin.Context) {
	// Implement logic for login
	var params model.LoginInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	codeRs, dataRs, err := service.UserLogin().Login(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeRs, dataRs)
}

// UpdatePasswordRegister
// @Summary      UpdatePasswordRegister
// @Description  UpdatePasswordRegister
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body model.UpdatePasswordRegisterInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/update_pass_register [post]

func (c *cUserLogin) UpdatePasswordRegister(ctx *gin.Context) {
	var params model.UpdatePasswordRegisterInput
	log.Println("params  UpdatePasswordRegister", params.UserPassword)
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	result, err := service.UserLogin().UpdatePasswordRegister(ctx, params.UserToken, params.UserPassword)
	if err != nil {
		response.ErrorResponse(ctx, result, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.CodeSuccess, result)
}

// Verify OTP Login By User
// @Summary      Verify OTP Login By User
// @Description  Verify OTP Login By User
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body model.VerifyInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/verify_account [post]
func (c *cUserLogin) VerifyOTP(ctx *gin.Context) {

	var params model.VerifyInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	result, err := service.UserLogin().VerifyOTP(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrInvalidOtp, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.CodeSuccess, result)
}

// User Registration documentation
// @Summary      User Registration
// @Description  When user is registered send otp to email
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body model.RegisterInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/register [post]
func (c *cUserLogin) Register(ctx *gin.Context) {
	print("Register", ctx)
	var params *model.RegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {

		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	codeStatus, err := service.UserLogin().Register(ctx, params)
	if err != nil {
		global.Logger.Error("Error registering user OTP", zap.Error(err))
		response.ErrorResponse(ctx, codeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.CodeSuccess, nil)
}
