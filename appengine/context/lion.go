package go18test

import (
	"net/http"

	"github.com/celrenheit/lion"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {
	l := lion.New()
	l.GetFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := appengine.NewContext(r)
		log.Infof(c, "hello")
	})
	http.Handle("/", l)
}
