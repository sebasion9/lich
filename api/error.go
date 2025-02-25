package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func QueryErr(err error) (bool, int, map[string]any) {
	status := http.StatusOK
	obj := gin.H {}
	exit := false
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true, http.StatusNotFound, nil
	}
	if err != nil {
		return true, http.StatusInternalServerError, gin.H { "msg": "internal server error", "err": err.Error() }
	}

	return exit, status, obj
}

func InsertErr(err error) (bool, int, map[string]any) {
	status := http.StatusOK
	obj := gin.H {}
	exit := false
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return true, http.StatusBadRequest, gin.H {"msg": "resource with this name already exist"}
	}
	if err != nil {
		return true, http.StatusInternalServerError, gin.H { "msg": "internal server error", "err": err.Error() }
	}

	return exit, status, obj

}

