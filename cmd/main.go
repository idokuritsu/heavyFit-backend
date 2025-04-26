package main

import (
	"log"
	"net/http"

	"github.com/idokuritsu/heavyFit-backend/internals/db"
	"github.com/uptrace/bunrouter"
)

func main(){
	db.InitDB("postgres://user:pass@localhost:5432/gym_db?sslmode=disable")
	router:=bunrouter.New()
	router.POST("/api/auth/register",RegisterHandler)
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", router)
}