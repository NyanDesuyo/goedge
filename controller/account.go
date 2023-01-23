package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/edgedb/edgedb-go"
	"github.com/gin-gonic/gin"

	"localhost/config"
)

type Account struct {
	No       edgedb.OptionalInt64    `edgedb:"no" json:"no"`
	ID       edgedb.UUID             `edgedb:"id" json:"id"`
	CreateAt edgedb.OptionalDateTime `edgedb:"create_at" json:"create_at"`
	UpdateAt edgedb.OptionalDateTime `edgedb:"update_at" json:"update_at"`
	DeleteAt edgedb.OptionalDateTime `edgedb:"delete_at" json:"delete_at"`
	Name     string                  `edgedb:"name" json:"name"`
	Currency string                  `edgedb:"currency" json:"currency"`
	Balance  edgedb.OptionalFloat64  `edgedb:"balance" json:"balance"`
	Status   bool                    `edgedb:"status" json:"status"`
}

func AccountCreate(c *gin.Context) {
	var payload Account

	err := json.NewDecoder(c.Request.Body).Decode(&payload)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error",
		})
		return
	}

	result := Account{}

	err = config.Pool.QuerySingle(
		c,
		`With inserted := (
			Insert Account {
        create_at := <datetime>$0,
        name := <str>$1,
        currency := <str>$2,
        balance := <float64>$3,
        status := <bool>$4,
			}
		)
		Select inserted {
		  no,
		  id,
		  create_at,
		  update_at,
		  delete_at,
		  name,
		  currency,
		  balance,
		  status,
		}`,
		&result,
		time.Now(),
		payload.Name,
		payload.Currency,
		payload.Balance,
		payload.Status,
	)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message:": "OK!",
		"data":     result,
	})
}

func AccountFindMany(c *gin.Context) {
	result := []Account{}

	err := config.Pool.Query(
		c,
		`Select Account {
			no,
			id,
			create_at,
			update_at,
			delete_at,
			name,
			currency,
			balance,
			status,
		}
		Order By .no`,
		&result,
	)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ok!",
		"data":    result,
	})
}
