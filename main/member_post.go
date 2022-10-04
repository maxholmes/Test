package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

// PostMember adds an Member from JSON received in the request body.
func (env Env) PostMember(c *gin.Context) {
	// Call BindJSON to bind the received JSON to
	// newMember
	var newMember Member
	if err := c.BindJSON(&newMember); err != nil {
		log.Printf("invalid JSON body: %v", err)
		makeGinResponse(c, http.StatusNotFound, err.Error())
		return
	}

	q := `INSERT INTO Member(first_name,last_name,email) VALUES($1,$2,$3) ON CONFLICT DO NOTHING`
	result, err := env.DB.Exec(q, newMember.FirstName, newMember.Last_Name, newMember.Email)
	if err != nil {
		log.Printf("error occurred while inserting new record into member table: %v", err)
		makeGinResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// checking the number of rows affected
	n, err := result.RowsAffected()
	if err != nil {
		log.Printf("error occurred while checking the returned result from database after insertion: %v", err)
		makeGinResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// if no record was inserted, let us say client has failed
	if n == 0 {
		e := "could not insert the record, please try again after sometime"
		log.Println(e)
		makeGinResponse(c, http.StatusInternalServerError, e)
		return
	}

	// NOTE:
	//
	// Here I wanted to return the location for newly created Member but this
	// 'pq' library does not support, LastInsertionID functionality.
	m := "successfully created the record"
	log.Println(m)
	makeGinResponse(c, http.StatusOK, m)
}
