package caching

import (
	"crypto/md5"
	"fmt"
	stderr "log"
	"math/rand"
	"net/http"
	"time"
)

type cached struct {
	next    http.Handler
	options []string
}

//
func (c cached) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	now := time.Now()
	utc := now.UTC()
	tomorrow := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, 1)

	nows := fmt.Sprintf("%8x", utc.UnixNano())
	rnds := fmt.Sprintf("%8x", rand.Int63())
	url5 := md5.Sum([]byte(req.URL.String()))
	pth5 := md5.Sum([]byte(req.URL.Path))
	qry5 := md5.Sum([]byte(fmt.Sprintf("%+v", req.URL.Query())))
	rnd5 := md5.Sum([]byte(fmt.Sprintf("%v.%v", nows, rnds)))

	// ETags are currently just generated daily based on the time and the URL.
	etag := fmt.Sprintf(`W/"%x"`, md5.Sum([]byte(fmt.Sprintf("%v.%v", tomorrow, url5))))

	ifNoneMatch := req.Header.Get("If-None-Match")
	if etag != "" && etag == ifNoneMatch {
		stderr.Printf("Not Modified")
		http.Error(w, "Not Modified", http.StatusNotModified)
		return
	}

	hs := w.Header()

	hs.Set("Etag", etag)

	requestId := fmt.Sprintf("%x%x%x%x", pth5[:4], qry5[:4], url5[:4], rnd5[:4])

	hs.Set("Last-Modified", utc.Format(time.RFC850))
	hs.Set("X-Request-Id", requestId)

	if len(c.options) > 0 && c.options[0] == "development" {
		stderr.Printf("WARNING: Development mode enabled! Caching disabled.")
		hs.Set("Cache-Control", "no-cache, no-store, must-revalidate")
		hs.Set("Pragma", "no-cache")
		hs.Set("Expires", "0")
	} else {
		hs.Set("Cache-Control", fmt.Sprintf("public, max-age=%v", int(tomorrow.Sub(now).Seconds())))
	}

	c.next.ServeHTTP(w, req)
}

//
func New(server http.Handler, options ...string) http.Handler {
	return cached{server, options}
}
