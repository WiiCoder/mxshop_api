package api

import (
	"context"
	"mxshop_api/user-web/forms"
	"mxshop_api/user-web/global"
	"mxshop_api/user-web/global/response"
	"mxshop_api/user-web/proto"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg:": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": e.Code(),
				})
			}
			return
		}
	}
}

func HandleValidatorError(ctx *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"msg": err.Error()})
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Translator)),
	})

	return
}

func GetUserList(ctx *gin.Context) {
	pageNum := ctx.DefaultQuery("pn", "0")
	pageInt, _ := strconv.Atoi(pageNum)
	pageSize := ctx.DefaultQuery("ps", "10")
	sizeInt, _ := strconv.Atoi(pageSize)
	userListResponse, err := global.UserSrvClient.GetUserList(ctx, &proto.PageInfo{
		PageNum:  uint32(pageInt),
		PageSize: uint32(sizeInt),
	})

	if err != nil {
		zap.S().Errorw("[GetUserList] 查询【用户列表】失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	result := make([]interface{}, 0)
	for _, value := range userListResponse.Data {
		data := response.UserResponse{
			Id:       value.Id,
			NickName: value.NickName,
			Birthday: response.JsonTime(time.Unix(int64(value.BirthDay), 0)),
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		}

		result = append(result, data)
	}

	reMap := gin.H{
		"total": userListResponse.Total,
	}
	reMap["data"] = result
	ctx.JSON(http.StatusOK, reMap)
}

func PassWordLogin(ctx *gin.Context) {
	passWordLoginForm := forms.PassWordLoginForm{}
	if err := ctx.ShouldBind(&passWordLoginForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	rsp, err := global.UserSrvClient.GetUserMobile(context.Background(), &proto.MobileRequest{
		Mobile: passWordLoginForm.Mobile,
	})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusBadRequest, map[string]string{
					"mobile": "用户不存在",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]string{
					"mobile": "登录失败",
				})
			}
			return
		}
	}

	checkPassword, err := global.UserSrvClient.CheckPassword(context.Background(), &proto.CheckPasswordInfo{
		Password:          passWordLoginForm.Password,
		EncryptedPassword: rsp.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"password": "登录失败",
		})
	}

	if checkPassword.Success {

	}
}

func Register(ctx *gin.Context) {}

func GetUserDetail(ctx *gin.Context) {}

func UpdateUser(ctx *gin.Context) {}
