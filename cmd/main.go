package main

import (
	"fmt"
	"os"
	// "log"
	"net/http"
	// "github.com/newrelic/go-agent"
	// "servicetest/config"
)

// func init() {
// 	conf := config.GetConfig()
// 	// set up logging. TODO use log interface
// 	servicename:= conf.GetString("service_name")

// }

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Kathy.")
}

func users(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get users success.")
}

func address(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get address success")
}

func useremail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get user email success.")
}

func usercompany(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get user company success")
}

func text(w http.ResponseWriter, r *http.Request) {
	// conf := config.GetConfig()
	// set up logging. TODO use log interface
	// textstring:= conf.GetString("text_string")
	textstring := os.Getenv("TEXT")
	if textstring == "" {
		textstring = "testenv_text"
	}
	fmt.Fprintf(w, textstring)
}

func encrypt(w http.ResponseWriter, r *http.Request) {
	// conf := config.GetConfig()
	// set up logging. TODO use log interface
	encryptstring := os.Getenv("ENCRYPT")
	if encryptstring == "" {
		encryptstring = "testenv_encrypt"
	}
	fmt.Fprintf(w, encryptstring)
}

func file(w http.ResponseWriter, r *http.Request) {
	// conf := config.GetConfig()
	// set up logging. TODO use log interface
	// filestring:= conf.GetString("file_string")
	filestring := os.Getenv("FILE")
	if filestring == "" {
		filestring = "testenv_file"
	}
	fmt.Fprintf(w, filestring)
}

func main() {
	// config := newrelic.NewConfig("servicetest", "94f09291122b5df5595f5d82b0d4b5c4fc9ee2d1")
	// config.CrossApplicationTracer.Enabled = true
	// config.DistributedTracer.Enabled = true
	// app, err := newrelic.NewApplication(config)
	

	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
	
	path := os.Getenv("BASE_PATH")
	if path == ""{
		path = "/servicetest/v1"
	}

	http.HandleFunc(path + "/health", health)
	
	
	http.HandleFunc(path + "/user", health)
	http.HandleFunc(path + "/text", text)
	http.HandleFunc(path + "/encrypt", encrypt)
	http.HandleFunc(path + "/file", file)

	http.HandleFunc(path + "/users", users)
	http.HandleFunc(path + "/users/id/address", address)
	http.HandleFunc(path + "/users/email", useremail)
	http.HandleFunc(path + "/users/company", usercompany)

	// http.HandleFunc(newrelic.WrapHandleFunc(app, "/servicetest/v1/user", health))
	// http.HandleFunc(newrelic.WrapHandleFunc(app, "/servicetest/v1/text", text))
	// http.HandleFunc(newrelic.WrapHandleFunc(app, "/servicetest/v1/encrypt", encrypt))
	// http.HandleFunc(newrelic.WrapHandleFunc(app, "/servicetest/v1/file", file))
	// txn := app.StartTransaction("myTxn", nil, nil)
	// defer txn.End()
	http.ListenAndServe(":8080", nil)
}
