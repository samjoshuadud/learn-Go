
package handlers

import (
	"net/http"
	"time"
) 

func PanicTest(w http.ResponseWriter, r *http.Request) {
	panic("This is a test panic")
}

func SlowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Second)
	w.Write([]byte("This was a slow response"))
}
