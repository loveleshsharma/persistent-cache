package persistantcache

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

type APIHandler struct {
	cache *cache.Cache
}

func NewApiHandler(cache *cache.Cache) APIHandler {
	return APIHandler{
		cache: cache,
	}
}

func (h APIHandler) Get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	value, err := h.cache.Get(key)
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

func (h APIHandler) Set(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("error occurred while reading request: ", err)
		return
	}

	var setRequest request.SetRequest
	json.Unmarshal(bytes, &setRequest)

	if setRequest.Expiry != 0 {
		h.cache.SetWithExpiry(setRequest.Key, setRequest.Value, time.Duration(time.Minute*time.Duration(setRequest.Expiry)))
	} else {
		h.cache.Set(setRequest.Key, setRequest.Value)
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
