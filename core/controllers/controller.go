package controllers

import (
	"errors"
	"github.com/NC-Cj/gen-gin-norm/core"
	"github.com/NC-Cj/gen-gin-norm/core/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type publicResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func CatchController(fn func(*gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := fn(c)

		if err != nil {
			var customErr error

			for _, ex := range core.ExceptionsToCatch {
				println(ex)
				if errors.As(err, &ex) {
					customErr = ex.(core.Exception)
					break
				}
			}

			if customErr != nil {
				c.JSON(http.StatusInternalServerError, publicResponse{
					Code: core.FAILURE,
					Data: nil,
					Msg:  customErr.Error(),
				})
			} else {
				logger.Warn("Caught exception: %s", customErr.Error())

				if core.OUTPUT_INTERNAL_ERROR {
					c.JSON(http.StatusOK, publicResponse{
						Code: core.InternalError,
						Data: nil,
						Msg:  customErr.Error(),
					})
				} else {
					c.JSON(http.StatusOK, publicResponse{
						Code: core.InternalError,
						Data: nil,
						Msg:  "Internal server error",
					})
				}
			}
			return
		}

		c.JSON(http.StatusOK, publicResponse{
			Code: core.SUCCESS,
			Data: result,
			Msg:  "success",
		})
	}
}
