package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type HiGin struct {

}

func (*HiGin) Hello(c *gin.Context) {

	x := c.Request.URL.Query()["hello"]
	if len(x) > 0 {
		fmt.Println(x[0])
	}
}
