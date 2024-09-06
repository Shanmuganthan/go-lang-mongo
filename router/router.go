package router

import (
	"net/http"

	"github.com/Shanmuganthan/go-lang-mongo/controllers"
	"github.com/Shanmuganthan/go-lang-mongo/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	//User Router
	r.Handle("/api/user/create", middleware.JWTMiddleware(http.HandlerFunc(controllers.CreateAdminUser))).Methods(http.MethodPost)
	r.Handle("/api/user/{id}", middleware.JWTMiddleware(http.HandlerFunc(controllers.UpdateAdminUser))).Methods(http.MethodPut)
	r.Handle("/api/user/{id}", middleware.JWTMiddleware(http.HandlerFunc(controllers.DeleteAdminUser))).Methods(http.MethodDelete)
	r.Handle("/api/user/all", middleware.JWTMiddleware(http.HandlerFunc(controllers.GetAllAdminUser))).Methods(http.MethodGet)
	r.Handle("/api/user/{id}", middleware.JWTMiddleware(http.HandlerFunc(controllers.GetByIdAdminUser))).Methods(http.MethodGet)

	//Auth Router
	r.HandleFunc("/api/auth/login", controllers.Login).Methods((http.MethodPost))
	return r
}
