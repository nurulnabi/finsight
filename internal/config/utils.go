package config

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func FormatResponse(response any) map[string]any {
	formatted := map[string]any{
		"status": struct {
			code       int
			error_code string
			message    string
		}{
			code:       0,
			error_code: "",
			message:    "",
		},
		"data": map[string]any{},
	}
	if err, ok := response.(error); ok {
		formatted["status"] = err
	} else {
		formatted["data"] = response
	}

	return formatted
}

func SendResponse(ctx *gin.Context, data any) {
	var formatted any = FormatResponse(data)
	fmt.Println("Utils.SendResponse", formatted)
	ctx.JSON(200, formatted)
}
