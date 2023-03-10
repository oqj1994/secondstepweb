package helpers

import (
	"fmt"
	"github.com/vitaLemoTea/secondstepweb/internal/config"
	"net/http"
	"runtime/debug"
)

var appConfig *config.Config

func NewHelpers(a *config.Config) {
	appConfig = a
}

func ClientError(w http.ResponseWriter, status int) {
	//send enough msg to User and to youeself
	appConfig.InfoLog.Printf("Client error with code:%d, text:%s", status, http.StatusText(status))
	http.Error(w, http.StatusText(status), status)
}
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	appConfig.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
