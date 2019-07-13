package admin

import "net/http"

type Admin struct {
}

var Default = Admin{}

func (a Admin) Register(writer http.ResponseWriter, request *http.Request) {}

func (a Admin) Login(writer http.ResponseWriter, request *http.Request) {}

func (a Admin) Logout(writer http.ResponseWriter, request *http.Request) {}
