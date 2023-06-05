package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phanlop12321/Dev_GO/db"
	"github.com/phanlop12321/Dev_GO/util"
)

type Authincome struct {
	IDUser      string
	Description string
	Money       uint
}

func Createincome(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(Authincome)
		if err := c.BindJSON(req); err != nil {
			util.Error(c, http.StatusBadRequest, err)
			return
		}

	}

}
func Getincome(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		income, err := db.GetIncome()
		if err != nil {
			util.Error(c, http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, income)
	}

}
