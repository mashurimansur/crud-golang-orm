package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name  string
	Email string
}

func InitialMigration() {
	db, err := gorm.Open("mysql", "root:@/db_learn_golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func AllUseres(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "root:@/db_learn_golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "root:@/db_learn_golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	name := r.FormValue("name")
	email := r.FormValue("email")

	db.Create(&User{Name: name, Email: email})
	fmt.Fprint(w, "New User Succesfully Added")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete Users Endpoint Hit")

	db, err := gorm.Open("mysql", "root:@/db_learn_golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to database")
	}

	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	var user User
	db.Where("id = ?", id).First(&user)
	db.Unscoped().Delete(&user)

	fmt.Fprint(w, "Successfully Deleted Data")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Update Uers Endpoint Hit")

	db, err := gorm.Open("mysql", "root:@/db_learn_golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to database")
	}

	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	var user User
	db.Where("id = ?", id).Find(&user)

	user.Name = r.FormValue("name")
	user.Email = r.FormValue("email")

	db.Save(&user)
	fmt.Fprint(w, "Successfully Updated Data")

}
