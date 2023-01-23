package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/edgedb/edgedb-go"
	"github.com/gin-gonic/gin"

	"localhost/config"
)

type Todo struct {
	No     edgedb.OptionalInt64 `edgedb:"no" json:"no"`
	ID     edgedb.UUID          `edgedb:"id" json:"id"`
	Title  string               `edgedb:"title" json:"title"`
	Body   edgedb.OptionalStr   `edgedb:"body" json:"body"`
	Status bool                 `edgedb:"status" json:"status"`
	Tag    []edgedb.OptionalStr `edgedb:"tag" json:"tag"`
}

func TodoCreate(c *gin.Context) {
	var payload Todo

	err := json.NewDecoder(c.Request.Body).Decode(&payload)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error",
		})
		return
	}

	result := []Todo{}

	err = config.Pool.Query(
		c,
		`With inserted := (
		  Insert Todo {
		    title := <str>$0,
		    body := <str>$1,
		    status := <bool>$2,
		    tag := <array<str>>$3,
		  }
		)
    Select inserted {
    	no,
      id,
      title,
      body,
      status,
      tag,
    }
		`,
		&result,
		payload.Title,
		payload.Body,
		payload.Status,
		payload.Tag,
	)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Ok!",
		"data":    result,
	})

}

func TodoFindMany(c *gin.Context) {
	result := []Todo{}

	err := config.Pool.Query(
		c,
		`Select Todo {
      no,
      id,
      title,
      body,
      status,
      tag,
    }
    Order by .no`,
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

func TodoFindOne(c *gin.Context) {
	result := Todo{}

	err := config.Pool.QuerySingle(
		c,
		`Select Todo {
      no,
      id,
      title,
      body,
      status,
      tag,
    }
    Filter .id = <uuid><str>$0
    Limit 1`,
		&result,
		c.Param("id"),
	)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok!",
			"data":    result,
		})
	}
}

func TodoUpdate(c *gin.Context) {
	var payload Todo

	err := json.NewDecoder(c.Request.Body).Decode(&payload)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error",
		})
		return
	}

	result := Todo{}

	parsedUUID, parseUUIDError := edgedb.ParseUUID(c.Param("id"))

	if parseUUIDError != nil {
		log.Print(parseUUIDError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error",
		})
		return
	}

	err = config.Pool.QuerySingle(
		c,
		`With updated := (
      Update Todo 
      filter .id = <uuid>$0
      Set {
        title := <str>$1,
        body := <str>$2,
        status := <bool>$3,
        tag := <array<str>>$4,
      }
    )
    Select updated {
      no,
      id,
      title,
      body,
      status,
      tag,
    }
    Limit 1`,
		&result,
		parsedUUID,
		payload.Title,
		payload.Body,
		payload.Status,
		payload.Tag,
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

func TodoDelete(c *gin.Context) {
	parseUUID, parseUUIDError := edgedb.ParseUUID(c.Param("id"))

	if parseUUIDError != nil {
		log.Print(parseUUIDError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error",
		})
		return
	}

	result := Todo{}

	err := config.Pool.QuerySingle(
		c,
		`With deleted := (
  		Delete Todo
  		filter .id = <uuid>$0
  	)
  	Select deleted {
  		no,
  		id,
  		title,
  		body,
  		status,
  		tag,
  	}`,
		&result,
		parseUUID,
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
