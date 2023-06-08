package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phanlop12321/Dev_GO/db"
	"github.com/phanlop12321/Dev_GO/model"
	"github.com/phanlop12321/Dev_GO/util"
	"golang.org/x/crypto/bcrypt"
)

const (
	cost = 12
)

type Authincome struct {
	ID          uint
	IDUser      string
	Description string
	Money       uint
}

type AuthReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(AuthReq)
		if err := c.BindJSON(req); err != nil {
			util.Error(c, http.StatusBadRequest, err)
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), cost)
		if err != nil {
			util.Error(c, http.StatusInternalServerError, err)
			return
		}
		user := model.User{
			Username: req.Username,
			Password: string(hash),
		}
		if err := db.CreateUser(&user); err != nil {
			util.Error(c, http.StatusInternalServerError, err)
			return
		}
		token, err := generateToken(user.ID)
		if err != nil {
			util.Error(c, http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

func Login(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(AuthReq)
		if err := c.BindJSON(req); err != nil {
			util.Error(c, http.StatusBadRequest, err)
			return
		}
		found, err := db.GetUserByUsername(req.Username)
		if found == nil || err != nil {
			util.Error(c, http.StatusUnauthorized, err)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(req.Password))
		if err != nil {
			util.Error(c, http.StatusUnauthorized, err)
			return
		}
		token, err := generateToken(found.ID)
		if err != nil {
			util.Error(c, http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

func Delete(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(Authincome)
		if err := c.BindJSON(req); err != nil {
			util.Error(c, http.StatusBadRequest, err)
			return
		}
		found, err := db.GetSaveByID(req.ID)
		if found == nil || err != nil {
			util.Error(c, http.StatusUnauthorized, err)
			return
		}
		if del := db.DeletSaveByID(req.ID); del != nil {
			util.Error(c, http.StatusUnauthorized, err)
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{
			"deleted": "deleted",
		})
	}
}

func Createincome(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(Authincome)
		if err := c.BindJSON(req); err != nil {
			util.Error(c, http.StatusBadRequest, err)
			return
		}
		income := model.Income{
			IDUser:      req.IDUser,
			Description: req.Description,
			Money:       req.Money,
		}
		if err := db.CreateIncome(&income); err != nil {
			util.Error(c, http.StatusInternalServerError, err)
			return
		}
	}

}
func Createexpenses(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(Authincome)
		if err := c.BindJSON(req); err != nil {
			util.Error(c, http.StatusBadRequest, err)
			return
		}
		expenses := model.Income{
			IDUser:      req.IDUser,
			Description: req.Description,
			Money:       req.Money,
		}
		if err := db.CreateExpenses(&expenses); err != nil {
			util.Error(c, http.StatusInternalServerError, err)
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
func Getexpenses(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		expenses, err := db.GetExpenses()
		if err != nil {
			util.Error(c, http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, expenses)
	}

}
