package main

import (
	"log"
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

type Login struct {
	Email string `json:"email"`
	Password string `json:"pass"`
}

func index(w http.ResponseWriter, r *http.Request) {	
	fmt.Fprintf(w, "Reyonapi Web Service v0.8.0")
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		json.NewEncoder(w).Encode("HATA: POST metodu disinda erisim yapilamaz -- Reyonapi Web Service v0.8.0")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print("Body read error -> %v", err)
		w.WriteHeader(500)
		return
	}
	var login Login
	if err := json.Unmarshal(body, &login); err != nil {
		log.Print("Body read error -> %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// JWT Token olusturmaca
	token, err := generateJWT(login.Email)
	if err != nil {
		log.Printf("Token olusturma esnasinda hata olustu. Hata -> %v", err)
	}

	//w.WriteHeader(200) // Successfully logged in.
	json.NewEncoder(w).Encode(`{"token": `+ token +`}`)
}

func list(w http.ResponseWriter, r *http.Request) {
	log.Println("==== REYON BOT TEST PROJECT ======")

	topicResponse, err := GetTopicList()
	if err != nil {
		log.Printf("JSON datasi hatali geliyor -> %s", err)
		json.NewEncoder(w).Encode("HATA: JSON datasi hatali geliyor -- Reyonapi Web Service v0.8.0")
		return
	}

	json.NewEncoder(w).Encode(topicResponse)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/list", list)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	setLogFile()

	log.Println("API server runing on 4002...")
	http.ListenAndServe(":4002", nil)
}

