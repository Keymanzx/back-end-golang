package response

import (
    "github.com/gin-gonic/gin"
)

func Success(c *gin.Context, code int, message string, data interface{}) {
    c.JSON(code, gin.H{
        "status":  "success",
        "message": message,
        "data":    data,
    })
}

func Error(c *gin.Context, code int, err string) {
    c.JSON(code, gin.H{
        "status":  "error",
        "message": err,
    })
}


type ErrorResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}