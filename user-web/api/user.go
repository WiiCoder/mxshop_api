package api

import (
	"fmt"
	"mxshop_api/user-web/global/response"
	"mxshop_api/user-web/proto"
	"net/http"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

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

func GetUserList(ctx *gin.Context) {
	ip := "0.0.0.0"
	port := 50051
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GRPC Server] 连接【用户服务】失败")
		return
	}

	UserSrvClient := proto.NewUserClient(userConn)

	userListResponse, err := UserSrvClient.GetUserList(ctx, &proto.PageInfo{
		PageNum:  1,
		PageSize: 1,
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

func PassWordLogin(ctx *gin.Context) {}

func Register(ctx *gin.Context) {}

func GetUserDetail(ctx *gin.Context) {}

func UpdateUser(ctx *gin.Context) {}
