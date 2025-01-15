package cache

import (
	"net/http"
	"time"
)

type CacheObject struct {
	Response     *http.Response
	ResponseBody []byte
	Created      time.Time
}

func NewCacheObject(response *http.Response, responseBody []byte) *CacheObject {
	return &CacheObject{
		Response:     response,
		ResponseBody: responseBody,
		Created:      time.Now(),
	}
}
