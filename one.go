package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:mukunth@tcp(127.0.0.1:3306)/application"

type content struct {
	gorm.Model
	NAME        string `json: "name"`
	COCOA       int    `json: "cocoa"`
	CALORIE     int    `json: "calorie"`
	INGREDIENTS string `json: "ingredients"`
	MFD         int    `json: "-"`
	EXPDT       int    `json: "-"`
}

type data struct {
	l *log.Logger
}

func maindata(l *log.Logger) *data {
	return &data{l}

}

func Intailthing() {

	DB, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("cannot conect to database")

	}
	DB.AutoMigrate(&content{})
}
func (f *data) Getusers(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")
	var contents []content
	DB.Find(&contents)
	json.NewEncoder(rw).Encode(contents)
}
func (f *data) Getuser(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var contents content
	DB.First(&contents, params["id"])
}
func (f *data) Createuser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var contents content
	json.NewDecoder(r.Body).Decode(&contents)
	DB.Create(&contents)
	json.NewEncoder(rw).Encode(contents)
}
func (f *data) Updateuser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	var contents content
	params := mux.Vars(r)
	DB.First(&contents, params["id"])
	json.NewDecoder(r.Body).Decode(&contents)
	DB.Save(&contents)
	json.NewEncoder(rw).Encode(contents)
}
func (f *data) Deleteuser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("content-type", "application/json")
	var contents content
	params := mux.Vars(r)
	DB.Delete(&contents, params["id"])
	json.NewEncoder(rw).Encode("the all deleted")
}
