package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Contact struct {
	Id string `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Company string `json:"company"`
	Project string `json:"project"`
	Notes string `json:"notes"`
}

var contacts = []Contact{
	Contact{Id: "1", FirstName: "Darcy", LastName: "Christiansen", Phone: "01202345678", Email: "darcy@tilma.com", Company: "Tillman Brothers", Project: "Project Tilman", Notes: "."},
	Contact{Id: "2", FirstName: "Jackson", LastName: "Keeling", Phone: "01201234567", Email:"jackson@rosewhite.com", Company: "Rose White Ltd", Project: "", Notes: ""},
	Contact{Id: "3", FirstName: "Kim", LastName: "Morrison", Phone: "012340000000", Email: "kim@lycos.com", Company: "Morrison’s Ltd", Project: "", Notes: "Manager"},
	Contact{Id: "4", FirstName: "John", LastName: "Doe", Phone: "01234111111", Email: "john@doe.com", Company: "Doe Ltd ", Project: "", Notes: ""},
	Contact{Id: "5", FirstName: "Levi", LastName: "Abbott", Phone: "0134085092", Email: "levi@gmail.com", Company: "Lankin,Heller and Parker Ltd", Project: "Project L", Notes: "Sales manager"},
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/contact", getContacts).Methods("GET")
	router.HandleFunc("/contact", createContact).Methods("POST")
	router.HandleFunc("/contact/{id}", getContact).Methods("GET")

	http.ListenAndServe(":8000", router)
}

func getContacts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

func getContact(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	for _, item := range contacts {
		if item.Id == id {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func createContact(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var contact Contact
	_ = r.ParseForm()

	contact.Id = strconv.FormatInt(int64(len(contacts) + 1), 10)
	contact.FirstName = r.FormValue("firstname")
	contact.LastName = r.FormValue("lastname")
	contact.Phone = r.FormValue("phone")
	contact.Email = r.FormValue("email")
	contact.Company = r.FormValue("company")
	contact.Project = r.FormValue("project")
	contact.Notes = r.FormValue("notes")

	contacts = append(contacts, contact)
	json.NewEncoder(w).Encode(&contact)
}
