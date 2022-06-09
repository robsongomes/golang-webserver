package main

import (
	"fmt"
	"net/http"
	"time"
)

type User struct {
	Username string
	LoggedAt time.Time
}

var user User

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	} else {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		if password == "" {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user = User{Username: username, LoggedAt: time.Now()}

		w.Write([]byte("Username: " + username + "\n"))
	}

}

func handleMe(w http.ResponseWriter, r *http.Request) {
	if user.Username == "" {
		http.Error(w, "no user logged in", http.StatusNotFound)
		return
	}
	w.Write([]byte("Username: " + user.Username + "\n"))
	w.Write([]byte("Logged at: " + user.LoggedAt.Format("02/01/2006 15:04:05") + "\n"))
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	user = User{}
	w.Write([]byte("Successfully logged out"))
}

func main() {
	// make the files in the static folder acessible by name
	static := http.FileServer(http.Dir("./static"))

	http.Handle("/", static)
	http.HandleFunc("/form", handleForm)
	http.HandleFunc("/logout", handleLogout)
	http.HandleFunc("/me", handleMe)

	fmt.Println("Listening on port 8000")
	http.ListenAndServe(":8000", nil)
}
