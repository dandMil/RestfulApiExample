package main

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"io/ioutil"
	"log"
	_ "log"
	"net/http"
	_ "net/http"
)

type User struct {
	UserName string `json:"UserName"`
	PlayerID string `json:"PlayerId"`
}

//Mocking DB..
var userArray []User

func getHello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello!")
	fmt.Println("Get endpoint called")
}

func getUsers(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(userArray)

	}
	//POST
func createUser(w http.ResponseWriter, r*http.Request)  {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)
	userArray = append(userArray, user)
	json.NewEncoder(w).Encode(user)
}
	//GET
func getUser(w http.ResponseWriter, r*http.Request)  {
	//got path from router
	vars:= mux.Vars(r)
	//vars is a map containing all path variables
	key := vars["PlayerId"]
	fmt.Fprintf(w,"Key: "+key)
	fmt.Println(len(userArray))
	for _, user := range userArray {
		if user.UserName == key{
			fmt.Println(user.UserName)

			json.NewEncoder(w).Encode(user)
		}
	}
}

func deleteUser(w http.ResponseWriter, r*http.Request){
	vars := mux.Vars(r)
	key := vars["PlayerId"]
	
	for index, user := range userArray {
		if user.PlayerID == key{
			userArray = append(userArray[:index],userArray[:index+1]...)
		}
	}
}

func handleRequests(){
	 myRouter := mux.NewRouter().StrictSlash(true)
	 myRouter.HandleFunc("/",getHello)
	 myRouter.HandleFunc("/all",getUsers)
	 myRouter.HandleFunc("/user/{PlayerId}",getUser)
	 myRouter.HandleFunc("/createUser/",createUser).Methods("POST")
	 myRouter.HandleFunc("/deleteUser/{PlayerId}",deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080",myRouter))
}


func main ()  {
	userArray = []User{
		{UserName: "SuperWinDude", PlayerID: "ABC123"},
		{UserName: "ProPlayer2", PlayerID: "DEF456"},
	}
	handleRequests()
}