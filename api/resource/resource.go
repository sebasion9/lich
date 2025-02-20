package resource

import "github.com/gin-gonic/gin"

// define whole CRUD for resource
// extend it with sync functionalities (pending?, fetch)

func New(dbOp func()) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

