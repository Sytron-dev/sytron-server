package main

import (
	// local imports
	"fmt"
	"log"
	"net/http"
	"os"
	"sytron-server/routers"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
)

func main() {
	// oldMain()

	r := routers.InitRouters()
	r.Run() // listen and serve on 0.0.0.0:8080
}

/* ----------------------------------------------------------------------------------------------------------------------
|                                                                                                                        |
|                                                                                                                        |
|                                                                                                                        |
|                      Old main method runs on 8081                                                                      |
|                                                                                                                        |
|                                                                                                                        |
|                                                                                                                        |
| ----------------------------------------------------------------------------------------------------------------------
*/

func oldMain() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/createuser", createUser)
	http.HandleFunc("/updateuser", updateUser)
	http.HandleFunc("/user", user)
	http.HandleFunc("/deleteuser", deleteUser)
	http.ListenAndServe(":8080", nil)
	fmt.Printf("Starting server at port 8081\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}

}

type User struct {
	ID         uuid.UUID `json:"id"`
	First_name string    `json:"first_name"`
	Last_name  string    `json:"last_name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if r.URL.Path != "/updateuser" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	supabaseUrl := os.Getenv("supabaseUrl")
	supabaseKey := os.Getenv("supabaseKey")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)
	row := User{
		First_name: "Vincent",
		Last_name:  "Kamemia",
		Email:      "vincentkamemia@gmail.com",
		Password:   "12345",
	}
	var results map[string]interface{}
	err = supabase.DB.From("users").Update(row).Eq("ID", "00000000-0000-0000-0000-000000000000").Execute(&results)
	if err != nil {
		panic(err)
	}

}
func createUser(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if r.URL.Path != "/createuser" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	supabaseUrl := os.Getenv("supabaseUrl")
	supabaseKey := os.Getenv("supabaseKey")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)
	// Generate a new UUID

	// Create a new user
	row := User{
		ID:         uuid.New(),
		First_name: "Vincent",
		Last_name:  "Kamemia",
		Email:      "johndoe@mail.com",
		Password:   "password",
	}

	var results []User
	err = supabase.DB.From("users").Insert(row).Execute(&results)
	if err != nil {
		panic(err)
	}

}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if r.URL.Path != "/deleteuser" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	supabaseUrl := os.Getenv("supabaseUrl")
	supabaseKey := os.Getenv("supabaseKey")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	var results map[string]interface{}

	err = supabase.DB.From("users").Delete().Eq("ID", "00000000-0000-0000-0000-000000000000").Execute(&results)
	if err != nil {
		panic(err)
	}
}
func user(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if r.URL.Path != "/user" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	supabaseUrl := os.Getenv("supabaseUrl")
	supabaseKey := os.Getenv("supabaseKey")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	var results map[string]interface{}
	err = supabase.DB.From("users").Select("*").Single().Execute(&results)
	if err != nil {
		panic(err)
	}

	fmt.Println(results)
}
