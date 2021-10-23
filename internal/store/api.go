package store

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	util "github.com/fatihsimsek/go-case-study/pkg/util"
)

func Init(service Service) {
	res := resource{service}

	registerHandlers(res)
	res.service.Init()
}

func registerHandlers(res resource) {
	http.Handle("/get", loggingMiddleware(http.HandlerFunc(res.Get)))
	http.Handle("/set", loggingMiddleware(http.HandlerFunc(res.Set)))
	http.Handle("/flush", loggingMiddleware(http.HandlerFunc(res.Flush)))
}

type resource struct {
	service Service
}

type setRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type response struct {
	IsSuccess bool        `json:"isSuccess"`
	Data      interface{} `json:"data"`
}

func (res resource) Get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		util.JSON(w, http.StatusBadRequest, response{IsSuccess: false, Data: "key is required"})
		return
	}
	value, found := res.service.Get(key)
	if found {
		util.JSON(w, http.StatusOK, response{IsSuccess: true, Data: value})
		return
	} else {
		util.JSON(w, http.StatusNotFound, response{IsSuccess: false, Data: "key not found"})
	}
}

func (res resource) Set(w http.ResponseWriter, r *http.Request) {
	var req setRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.JSON(w, http.StatusBadRequest, response{IsSuccess: false, Data: err.Error()})
		return
	}
	res.service.Put(req.Key, req.Value)
	util.JSON(w, http.StatusOK, response{IsSuccess: true})
}

func (res resource) Flush(w http.ResponseWriter, r *http.Request) {
	res.service.Flush()
	util.JSON(w, http.StatusOK, response{IsSuccess: true})
}

func loggingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		uri := r.RequestURI
		method := r.Method

		handler.ServeHTTP(w, r)

		duration := time.Since(start)

		log.Printf("Uri:%s Method:%s Duration:%s", uri, method, duration)
	})
}
