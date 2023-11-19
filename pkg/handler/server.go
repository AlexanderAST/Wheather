package handler

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
)

func sayHello(hello string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(hello))
		if err != nil {
			return
		}
	}
}

func takeWheather() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &City{}
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

		var data map[string]interface{}

		err = json.Unmarshal(responceBody, &data)

		if err != nil {
			http.Error(w, "ошибка при парсинге json", http.StatusInternalServerError)
			return
		}

		cityName := data["name"].(string)

		dataMain, ok := data["main"].(map[string]interface{})
		if !ok {
			fmt.Println("Поле 'main' не найдено или некорректный формат")
			return
		}

		temperature := dataMain["temp"].(float64)

		weatherArr, ok := data["weather"].([]interface{})
		if !ok {
			fmt.Println("Поле 'weather' не найдено или некорректный формат")
			return
		}

		weatherObj, ok := weatherArr[0].(map[string]interface{})
		if !ok {
			fmt.Println("Некорректный формат объекта погоды")
			return
		}

		mainWeather, ok := weatherObj["main"].(string)
		if !ok {
			fmt.Println("Поле 'main' погоды не найдено или некорректный формат")
			return
		}

		mainDescription, ok := weatherObj["description"].(string)
		if !ok {
			fmt.Println("Поле 'main' погоды не найдено или некорректный формат")
			return
		}

		dataWind, ok := data["wind"].(map[string]interface{})
		if !ok {
			fmt.Println("Поле 'main' не найдено или некорректный формат")
			return
		}

		windSpeed := dataWind["speed"].(float64)

		result := make(map[string]interface{})

		result["city"] = cityName
		result["temperature"] = int(temperature - 273)
		result["whether"] = mainWeather
		result["description whether"] = mainDescription
		result["wind speed"] = windSpeed

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			return
		}
	}
}
