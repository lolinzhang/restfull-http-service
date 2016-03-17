package handler

import (
	"fmt"
	"net/http"
	"encoding/json"
	
	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"

	"orm"
)

func UsersGetHandler(w http.ResponseWriter, r *http.Request){
	users := orm.FindUsers()

	outJson, _ := json.Marshal(&users)
    
	w.Write(outJson)
	w.Write([]byte("\n"))
}

func UsersPOSTHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()  

	query := ""

    for k, _ := range r.Form {
    	query = k
    	break
    }

    queryByte := []byte(query)

    js, err := simplejson.NewJson(queryByte)

    if err != nil {
    	fmt.Println("UserPostError:", err)
    }

    userName, _ := js.Get("name").String()

    newUser := orm.CreateUser(userName)

    outJson, _ := json.Marshal(&newUser)
    
	w.Write(outJson)
	w.Write([]byte("\n"))
}

func UsersRelationshipsGetHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userId := vars["user_id"]
	relationships := orm.FindRelationships(userId)

	outJson, _ := json.Marshal(&relationships)
    
	w.Write(outJson)
	w.Write([]byte("\n"))
}

func UsersRelationshipsPutHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userId := vars["user_id"]
	otherUserId := vars["other_user_id"]

	r.ParseForm()
	query := ""

    for k, _ := range r.Form {
    	query = k
    	break
    }

    queryByte := []byte(query)

    js, err := simplejson.NewJson(queryByte)

    if err != nil {
    	fmt.Println("UsersRelationshipsPutHandlerError:", err)
    }

    state, _ := js.Get("state").String()

	relationship := orm.UpdateRelationship(userId, otherUserId, state)
	fmt.Println(relationship)

    outJson, _ := json.Marshal(&relationship)
    
	w.Write(outJson)
	w.Write([]byte("\n"))
}