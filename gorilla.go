package main
import(
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"encoding/json"
)

type User struct{
	ID  	  string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

var users []User

func GetUserById(w http.ResponseWriter,req *http.Request){

	params:=mux.Vars(req)
	for _,item := range users{

		if item.ID==params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func CreateUser(w http.ResponseWriter,req *http.Request){
		params :=mux.Vars(req)
		var user User
		_= json.NewDecoder(req.Body).Decode(&user)
		user.ID=params["id"]
		users=append(users,user)

		json.NewEncoder(w).Encode(users)

}
func GetUsers(w http.ResponseWriter,req *http.Request){
		json.NewEncoder(w).Encode(users)
}
func DeleteUser(w http.ResponseWriter,req *http.Request){

	params := mux.Vars(req)
	for index,item :=range users{
		if item.ID==params["id"]{
			users =append(users[:index],users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}
func main(){
	fmt.Println("Welcom")
	
	users=append(users,User{ID:"1",Firstname:"Mythili",Lastname:"P"})
	users=append(users,User{ID:"2",Firstname:"Sharu",Lastname:"S"})

	router:=mux.NewRouter()
	router.HandleFunc("/users",GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}",GetUserById).Methods("GET")
	router.HandleFunc("/users/{id}",CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}",DeleteUser).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":8080",router))


}