package auth

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/idokuritsu/heavyFit-backend/internals/db"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request){
	var input RegisterInput
	if err:=json.NewDecoder(r.Body).Decode(&input);err!=nil{
		http.Error(w,"Invalid input",http.StatusBadRequest)
		return
	}
//trim nd validate
	input.Email=strings.TrimSpace(input.Email)
	input.Name=strings.TrimSpace(input.Name)
	if input.Email==""||input.Password==""||input.Name==""{
		http.Error(w,"All fields are required",http.StatusBadRequest)
		return
	}
//check if user exists already
   existing:=new(User)
  err:=db.DB.NewSelect().Model(existing).Where("email=?",input.Email).Scan(r.Context()) 
  if err==nil{
	http.Error(w, "Email already registered", http.StatusBadRequest)
		return
  } 
  //hash pwd
  
 hashed,err:=bcrypt.GenerateFromPassword([]byte(input.Password),bcrypt.DefaultCost)
 if err!=nil{
	http.Error(w, "Failed to encrypt password", http.StatusInternalServerError)
	return
 }

 user:=&User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashed),
	}
	_,err=db.DB.NewInsert().Model(user).Exec(r.Context())
	if err!=nil{
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{

		"message": "User registered successfully",
	})

}