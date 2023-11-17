package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/palp1tate/brevinect/api/global"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
)

type Response struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Token interface{} `json:"token,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func HandleHttpResponse(c *gin.Context, code int, msg interface{}, token interface{}, data interface{}) {
	response := Response{
		Code:  code,
		Msg:   msg,
		Token: token,
		Data:  data,
	}
	c.JSON(code, response)
	return
}

func HandleGrpcErrorToHttp(c *gin.Context, err error) {
	if e, ok := status.FromError(err); ok {
		HandleHttpResponse(c, http.StatusBadRequest, e.Message(), nil, nil)
		return
	}
}

func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleValidatorError(c *gin.Context, err error) {
	var errs validator.ValidationErrors
	if ok := errors.As(err, &errs); !ok {
		zap.S().Errorw("验证器参数校验失败", "err", err.Error())
		HandleHttpResponse(c, http.StatusInternalServerError, err.Error(), nil, nil)
	} else {
		zap.S().Info("参数格式错误")
		HandleHttpResponse(c, http.StatusBadRequest, removeTopStruct(errs.Translate(global.Translator)), nil, nil)
	}
	return
}
