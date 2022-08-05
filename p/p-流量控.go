package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		defer func() { recover() }()
		log.Panic("too many requests")
		return
	}
	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func registerHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/ce", ce)
	return router
}

func ce(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	time.Sleep(time.Second * 100)
	t, _ := template.ParseFiles("ce.html")
	t.Execute(w, nil)
}

type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}
func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		fmt.Println("reached out")
		return false
	}
	cl.bucket <- 1
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	fmt.Printf("new connect coming %d", c)
}

func main() {
	r := registerHandlers()
	mh := NewMiddleWareHandler(r, 60)
	http.ListenAndServe(":9000", mh)
}
