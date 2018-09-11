package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

//LoginHandler
func LoginHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(`{"message":"Login"}`))
}

//LandingHandler
func LandingHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(`{"message":"Landing Page Data"}`))
}

//AuthHandler
func AuthHandler(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	fmt.Println("Authenticating....")
	// Set request header with key-value Secret and password respectively
	header := req.Header.Get("Secret")
	if header != "password" {
		fmt.Println("Invalid Password")
		res.WriteHeader(401)
		return
	} else {
		fmt.Println("Success")
		next(res, req)
	}
}

func main() {

	fmt.Println("Connection Established")

	router := mux.NewRouter()

	//Home path
	router.Path("/").HandlerFunc(LoginHandler)

	//App path
	authenticatedUser := mux.NewRouter()
	router.PathPrefix("/app").Handler(negroni.New(
		//This is only applicable to /admin routes
		negroni.HandlerFunc(AuthHandler),
		//Add your handlers here which is only applicable to /admin routes
		negroni.Wrap(authenticatedUser),
	))

	adminRoutes := authenticatedUser.PathPrefix("/app").Subrouter()
	adminRoutes.HandleFunc("/landing", LandingHandler).Methods("GET")
	http.ListenAndServe(":8000", router)
}
