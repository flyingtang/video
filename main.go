package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
	cl *ConnLimiter
}

func newMiddleWareHandler (r *httprouter.Router, cc int) middleWareHandler{
	m := middleWareHandler{}
	m.r = r
	m.cl = newConnLimiter(cc)
	return  m
}

func (mh middleWareHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	if !mh.cl.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "too many connection")
		return
	}
	mh.r.ServeHTTP(w, r)
	defer mh.cl.ReleaseConn()
}

func registerHandlers() *httprouter.Router {

	router := httprouter.New()
	router.GET("/videos/:vid-id", streamHandle)
	router.POST("/upload/:vid-id", uploadHandle)

	router.GET("/testPage", testHanle)
	return router
}

func main() {

	r := registerHandlers()
	mh := newMiddleWareHandler(r, 2)
	http.ListenAndServe(":4001", mh)
}
