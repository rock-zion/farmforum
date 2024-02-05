package routes

import (
	"github.com/farmforum/controllers"
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {

	router := mux.NewRouter()
	router.Handle("/crops", controllers.AddCrop()).Methods("POST")
	router.Handle("/crops", controllers.FetchAllCrops()).Methods("GET")
	router.Handle("/crops/{id}", controllers.FetchCropById()).Methods("GET")
	router.Handle("/crops/many", controllers.InsertManyCrops()).Methods("POST")
	router.Handle("/crops/{id}", controllers.DeleteCrop()).Methods("DELETE")
	router.Handle("/crops/{id}", controllers.EditCrop()).Methods("PUT")
	router.Handle("/crops/questions/{id}", controllers.FetchQuestionByCropId()).Methods("GET")

	//
	router.Handle("/question/{id}", controllers.AddQuestion()).Methods("POST")
	router.Handle("/questions/{id}", controllers.UpdateQuestion()).Methods("GET")

	// routers to handle user actions
	router.Handle("/user", controllers.HandleAddUser()).Methods(("POST"))
	router.Handle("/user/login", controllers.HandleLogUserIn()).Methods(("POST"))

	return router
}
