package router

import (
	inventory_info_controller "inventory_management/controller/inventory_info_controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/inventory_info", inventory_info_controller.Post_Oneinventory_info).Methods("POST")
	router.HandleFunc("/api/inventory_info", inventory_info_controller.Get_Allinventory_info).Methods("GET")
	router.HandleFunc("/api/inventory_info/{unique_item_id}", inventory_info_controller.GetOneinventory_info).Methods("GET")
	router.HandleFunc("/api/inventory_info/{unique_item_id}", inventory_info_controller.Update_Oneinventory_info).Methods("PUT")
	router.HandleFunc("/api/inventory_info/{unique_item_id}", inventory_info_controller.Delete_Oneinventory_info).Methods("DELETE")

	return router
}
