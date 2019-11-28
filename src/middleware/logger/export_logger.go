package logger

import (
	"fmt"
	"net/http"
	"time"
)

type Logger struct{}

func (*Logger) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("The logger middleware is excuting!")
	t := time.Now()
	next.ServeHTTP(w, r)
	fmt.Println("Execution time: %s \n", time.Now().Sub(t).String())
}
