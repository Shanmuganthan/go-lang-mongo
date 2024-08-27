package router

import (
	"net/http"

	"github.com/Shanmuganthan/go-lang-mongo/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	//User Router
	r.HandleFunc("/api/user/create", controllers.CreateAdminUser).Methods(http.MethodPost)
	r.HandleFunc("/api/user/{id}", controllers.UpdateAdminUser).Methods(http.MethodPut)
	r.HandleFunc("/api/user/{id}", controllers.DeleteAdminUser).Methods(http.MethodDelete)
	r.HandleFunc("/api/user/all", controllers.GetAllAdminUser).Methods(http.MethodGet)
	r.HandleFunc("/api/user/{id}", controllers.GetByIdAdminUser).Methods(http.MethodGet)

	//Auth Router
	r.HandleFunc("/api/auth/login", controllers.Login).Methods((http.MethodPost))
	return r
}
