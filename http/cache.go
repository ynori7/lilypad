package http

import (
	"fmt"
	"net/http"
	"time"
)

const (
	cacheControlHeader = "Cache-Control"
	lastModifiedHeader = "Last-Modified"
	expiresHeader      = "Expires"
	maxAge             = "max-age"
	publicCache        = "public"
	noCache            = "no-cache"
)

func getCacheControlHeaders(cacheDuration int64) Headers {
	headers := make(map[string]string)

	if cacheDuration == 0 {
		headers[cacheControlHeader] = noCache
		return headers
	}

	cacheSince := time.Now().In(time.UTC).Format(http.TimeFormat)
	cacheUntil := time.Now().In(time.UTC).Add(time.Duration(cacheDuration) * time.Second).Format(http.TimeFormat)

	headers[cacheControlHeader] = fmt.Sprintf("%s %d, %s", maxAge, cacheDuration, publicCache)
	headers[lastModifiedHeader] = cacheSince
	headers[expiresHeader] = cacheUntil

	return headers
}
