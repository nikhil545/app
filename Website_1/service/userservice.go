package service

import (
	"Website_1/util"
	"fmt"
	"net/http"
)

type Users struct {
	Username      string
	Password      string
	Email         string
	Id            string
	Authenticated bool
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, err := util.Store.Get(r, "web-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options.MaxAge = -1
	session.Save(r, w)
	err = util.Tpl.ExecuteTemplate(w, "index.gohtml", nil)
	fmt.Println(err)
}

func Login(w http.ResponseWriter, r *http.Request) {
	user, err := GetUserSession(r)
	if err != nil || user.Authenticated {
		http.Redirect(w, r, "/dashboard", 301)
	}

	err = util.Tpl.ExecuteTemplate(w, "login.gohtml", nil)
	fmt.Println(err)
}

func Loginprocess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	session, err := util.Store.Get(r, "web-session") //create or get an existing cookie
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	var databaseUsername string
	var databasePassword string
	var userId string

	err = util.DB.QueryRow("SELECT id, username, password FROM users WHERE username=$1 AND password=$2", username, password).Scan(&userId, &databaseUsername, &databasePassword)
	if err != nil {
		http.Redirect(w, r, "/login", 301)
		return
	}

	//saving user data from db for cookie
	user := Users{
		Username:      databaseUsername,
		Id:            userId,
		Authenticated: true, // to check and validate whether user is logged it
	}

	session.Options.Path = "/"
	session.Values["user"] = user //storing the cookie value with key name called user
	err = session.Save(r, w)      //save the cookie data into session to the browser
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/dashboard", 301)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	err := util.Tpl.ExecuteTemplate(w, "signup.gohtml", nil)
	fmt.Println(err)

}

func Signupprocess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	session, err := util.Store.Get(r, "web-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users := Users{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
		Email:    r.FormValue("email"),
	}

	var userId string
	err = util.DB.QueryRow("INSERT INTO users(username, password,email) VALUES($1,$2,$3) RETURNING ID", users.Username, users.Password, users.Email).Scan(&userId)
	if err != nil {
		http.Error(w, "Server error, unable to create your account.", 500)
		return
	}

	user := Users{
		Username:      users.Username,
		Id:            userId,
		Authenticated: true,
	}

	session.Options.Path = "/"
	session.Values["user"] = user
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/dashboard", 301)
	w.Write([]byte("User created!"))
	return

}

func GetUserSession(r *http.Request) (Users, error) {

	var user = Users{}
	session, err := util.Store.Get(r, "web-session")
	if err != nil {
		return user, err
	}
	val := session.Values["user"]

	user, ok := val.(Users)
	if !ok {
		return Users{Authenticated: false}, nil
	}
	return user, nil
}
