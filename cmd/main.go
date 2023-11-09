package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

type city struct {
	City string `json:"city" binding:"required"`
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/hello", sayHello("АОАООАОАОАОАОАО")).Methods(http.MethodGet)
	r.HandleFunc("/wheather", takeWheather()).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", r))

}

func sayHello(hello string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(hello))
	}
}

func takeWheather() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &city{}
		if err := godotenv.Load(); err != nil {
			fmt.Errorf("error loading env %v", err.Error())
		}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			error.Error(err)
			return
		}

		apiUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", req.City, os.Getenv("APIKEY"))
		responce, err := http.Get(apiUrl)
		if err != nil {
			http.Error(w, "Ошибка при отправке запроса", http.StatusInternalServerError)
			return
		}
		defer responce.Body.Close()

		responceBody, err := io.ReadAll(responce.Body)

		if err != nil {
			http.Error(w, "ошибка при чтении ответа", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(responceBody)
	}
}
