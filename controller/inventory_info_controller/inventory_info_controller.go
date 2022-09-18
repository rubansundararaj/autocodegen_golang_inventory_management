package inventory_infoController

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	connectwithdb "inventory_management/dbconnection"
	"inventory_management/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const dbName = "inventory_management"
const colName = "inventory_info"

var collection *mongo.Collection

func init() {
	client := connectwithdb.InitializeAndReturnConnection()

	collection = client.Database(dbName).Collection(colName)
}

func postOneinventory_info(inventory_info model.InventoryInfo) {
	fmt.Println("Before inserting")
	fmt.Println(inventory_info)
	inserted, err := collection.InsertOne(context.Background(), inventory_info)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 inventory_info", inserted.InsertedID)
}

func updateOneinventory_info(uniqueitemid string,inventory_info model.InventoryInfo) {
	filter := bson.M{"uniqueitemid": uniqueitemid }
	update := bson.M{"$set":inventory_info}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

func deleteOneinventory_info(uniqueitemid string) {
	filter := bson.M{"uniqueitemid": uniqueitemid}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted one count : ", deleteCount)

}

func getAllinventory_info() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var InventoryInfos []primitive.M
	for cursor.Next(context.Background()) {
		var InventoryInfo bson.M
		err := cursor.Decode(&InventoryInfo)
		if err != nil {
			log.Fatal(err)
		}
		InventoryInfos = append(InventoryInfos, InventoryInfo)
	}

	defer cursor.Close(context.Background())

	return InventoryInfos
}

func getOneinventory_info(uniqueitemid string) primitive.M {
	var result bson.M
	filter := bson.M{"uniqueitemid": uniqueitemid}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
	}
	fmt.Println(result)
	return result
}

func GetOneinventory_info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	oneinventory_info := getOneinventory_info(params["unique_item_id"])

	json.NewEncoder(w).Encode(oneinventory_info)
}


func Get_Allinventory_info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allinventory_info := getAllinventory_info()

	json.NewEncoder(w).Encode(allinventory_info)
}

func Post_Oneinventory_info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var inventory_info model.InventoryInfo

	_ = json.NewDecoder(r.Body).Decode(&inventory_info)
	fmt.Println(inventory_info)
	postOneinventory_info(inventory_info)
	json.NewEncoder(w).Encode(inventory_info)
}

func Update_Oneinventory_info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	var inventory_info model.InventoryInfo

	_ = json.NewDecoder(r.Body).Decode(&inventory_info)

	updateOneinventory_info(params["unique_item_id"],inventory_info)

	json.NewEncoder(w).Encode(params["unique_item_id"])

}

func Delete_Oneinventory_info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	deleteOneinventory_info(params["unique_item_id"])

	json.NewEncoder(w).Encode(params["unique_item_id"])
}

