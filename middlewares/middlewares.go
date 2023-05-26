package middlewares

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func FilterContenType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Inside filterContentType middleware, before the request is handled")

		if r.Header.Get("Content-type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported Media Type. Json type expected\n"))
			return
		}

		handler.ServeHTTP(w, r)

		log.Println("Inside filterContentType middleware, after the request was handled")
	})
}

func SetServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Inside setServerTimeCookie middleware, before the request is handled")

		handler.ServeHTTP(w, r)

		cookie := http.Cookie{Name: "Server Time(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}

		http.SetCookie(w, &cookie)

		log.Println("Inside setServerTimeCookie middleware, after the request is handled")

	})
}
