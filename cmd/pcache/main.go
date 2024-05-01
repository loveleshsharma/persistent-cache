package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/loveleshsharma/persistent-cache/cache"
	"github.com/loveleshsharma/persistent-cache/request"
	"github.com/loveleshsharma/persistent-cache/response"
)

func main() {
	err := initObjects()
	if err != nil {
		log.Fatalln("error occurred", err)
		return
	}

	http.HandleFunc("/set", Set)
	http.HandleFunc("/get", Get)

	http.ListenAndServe(":8080", nil)
}

func Get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	value, err := persistentCache.Get(key)
	if err != nil {
		w.Write([]byte("Key not found"))
		return
	}

	getResp := response.GetResponse{
		Value: value.(cache.Value).GetValue(),
	}

	bytes, _ := json.Marshal(getResp)
	w.Write(bytes)
}

func Set(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("error occurred while reading request: ", err)
		return
	}

	var setRequest request.SetRequest
	json.Unmarshal(bytes, &setRequest)

	if setRequest.Expiry != 0 {
		persistentCache.SetWithExpiry(setRequest.Key, setRequest.Value, time.Duration(time.Minute*time.Duration(setRequest.Expiry)))
	} else {
		persistentCache.Set(setRequest.Key, setRequest.Value)
	}
	setResponse := response.SetResponse{
		Status: "Success",
	}

	resBytes, err := json.Marshal(setResponse)
	if err != nil {
		fmt.Println("error occurres while marshalling response", err)
		return
	}

	w.Write(resBytes)
}
