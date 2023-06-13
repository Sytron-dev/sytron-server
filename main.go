package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/google/uuid"
	supa "github.com/nedpals/supabase-go"
) 


type User struct {
	ID uuid.UUID `json:"id"`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
}
func updateUser(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/updateuser" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	supabaseUrl := "https://bnzcbbpmekiavacefqfr.supabase.co"
	supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImJuemNiYnBtZWtpYXZhY2VmcWZyIiwicm9sZSI6ImFub24iLCJpYXQiOjE2ODY0Nzc1MjYsImV4cCI6MjAwMjA1MzUyNn0.MYAFfpmn2xo6OVoYgMDYyQuNBzpwfEvHaU1w8eF8VJg"
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)
	row := User{
		First_name: "Vincent",
		Last_name: "Kamemia",
		Email: "vincentkamemia@gmail.com",
		Password: "12345",
	}
	var results map[string]interface{}
	err := supabase.DB.From("users").Update(row).Eq("id", "c0b0b0a0-0a0a-0a0a-0a0a-0a0a0a0a0a0a").Execute(&results)
	if err != nil {
		panic(err)
	  }

}
func createUser(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/createuser" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	supabaseUrl := "https://bnzcbbpmekiavacefqfr.supabase.co"
	supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImJuemNiYnBtZWtpYXZhY2VmcWZyIiwicm9sZSI6ImFub24iLCJpYXQiOjE2ODY0Nzc1MjYsImV4cCI6MjAwMjA1MzUyNn0.MYAFfpmn2xo6OVoYgMDYyQuNBzpwfEvHaU1w8eF8VJg"
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)
	// Generate a new UUID
	
	// Create a new user
	row := User{
		ID : uuid.New(),
		First_name: "Vincent",
		Last_name: "Kamemia",
		Email: "johndoe@mail.com",
		Password: "password",
		}
		var results []User
		err := supabase.DB.From("users").Insert(row).Execute(&results)
		if err != nil {
			panic(err)
		  }

}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/createuser", createUser)
	http.HandleFunc("/updateuser", updateUser)
	http.ListenAndServe(":8080", nil)
	fmt.Printf("Starting server at port 8080\n")
	if err:= http.ListenAndServe(":8080", nil); err !=nil {
		log.Fatal(err)
	}

}
