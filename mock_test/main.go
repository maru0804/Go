package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/maru0804/Go.git/mock_test/service"
	"net/http"
)

// ここでサーバの起動処理を記述

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", healthCheck)
	// クエリ文字列の取得
	r.HandleFunc("/create", createUserHandler)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("faild to run server : %s\n", err)
	}
}
func healthCheck(w http.ResponseWriter, r *http.Request) {
	status, err := w.Write([]byte("Gorilla!\n"))
	if err != nil {
		fmt.Printf("faild to write response : %s\n", err)
	} else {
		w.WriteHeader(status)
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	name := q.Get("name")
	age := q.Get("age")
	err := service.CreateUser(name, age)
	if err != nil {
		fmt.Printf("failed to create user : %s\n", err)
	}
	status, err := w.Write([]byte("create user !\n"))
	if err != nil {
		fmt.Printf("faild to write response : %s\n", err)
	} else {
		w.WriteHeader(status)
	}
}
