package middle_ware

import (
	"github.com/gin-gonic/gin"

	"oasis/ready/biz/const_def"
	"oasis/ready/utils"
)

func CheckRequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		rId := c.Request.Header.Get(const_def.XRequestId)

		if rId == "" {
			rId = utils.GenXid()
		}

		c.Request.Header.Set(const_def.XRequestId, rId)
		c.Header(const_def.XRequestId, rId)
	}
}
