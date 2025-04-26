package controller

import (
	"promotion/internal/content"
	"promotion/pkg/failure"
	"promotion/pkg/logger"
	"promotion/pkg/response"

	"github.com/gin-gonic/gin"
)

type ContentController struct {
	log    *logger.Logger
	module *content.Module
}

func NewContentController(
	log *logger.Logger, m *content.Module,
) *ContentController {
	return &ContentController{log, m}
}

func (c *ContentController) GenerateLessonFromCode(ctx *gin.Context) {
	body, errBinding := BindJSON[content.Lesson](ctx)
	if errBinding != nil {
		response.ErrBinding(ctx, &failure.BindJSONErr{
			Code:        failure.ErrReusableCodeGetByCodeBinding,
			OriginalErr: failure.ErrWithTrace(errBinding.OriginalErr),
			Model:       errBinding.Model,
		})
		return
	}

	rc, err := c.module.Service.GenerateLessonFromCode(ctx.Request.Context(), body.Code)
	if err != nil {
		var errCode failure.ErrCode
		if failure.IsSQLRecordNotFound(err) {
			errCode = failure.ErrReusableCodeNotFound
		} else {
			errCode = failure.ErrReusableCodeFailed
		}
		response.ErrApp(ctx, &failure.AppErr{
			Code:        errCode,
			OriginalErr: failure.ErrWithTrace(err),
		})
		return
	}

	response.Success(ctx, rc)
}
