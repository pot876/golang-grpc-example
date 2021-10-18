package api

import (
	"fibo-prj/internal/fibo"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateRestGin() *gin.Engine {
	e := gin.Default()
	addFiboHandler(&e.RouterGroup, fibo.FromTo)

	return e
}

func addFiboHandler(r *gin.RouterGroup, fiboGetter func(fr, to uint64) ([]string, error)) {
	fiboHandler := func(c *gin.Context) {
		boundsString := c.Param("bounds")

		var bounds [2]uint64
		if !extractBoundHelper(boundsString, &bounds) {
			c.String(http.StatusBadRequest, "bad bounds format, expected url template: /fibo/{uint64}-{uint64}")
			return
		}

		var result []string
		var err error
		if result, err = fiboGetter(bounds[0], bounds[1]); err != nil {
			logrus.Errorf("fibo getter error, bounds: %v, err: [%v]", bounds, err)
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusAccepted, result)
	}
	r.GET("/fibo/:bounds", fiboHandler)
}

func extractBoundHelper(boundsString string, result *[2]uint64) bool {
	var err error

	splitted := strings.Split(boundsString, "-")
	if len(splitted) != 2 {
		return false
	}

	for i := range splitted {
		if result[i], err = strconv.ParseUint(splitted[i], 10, 0); err != nil {
			return false
		}
	}
	return true
}
