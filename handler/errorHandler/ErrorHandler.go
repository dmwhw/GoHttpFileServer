package errorHandler

import (
	"fmt"
	"net/http"
)

func Handle(w http.ResponseWriter, req *http.Request, err error,
	code int,
	msg string,
	isHanlded bool) {
	fmt.Println("-------------main error------------------")
}
