package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/Shanmuganthan/go-lang-mongo/common"
	"github.com/Shanmuganthan/go-lang-mongo/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateAdminUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	log.Print("user | CreateAdminUser | Started")

	userID, ok := r.Context().Value("user_details").(interface{})

	if !ok {
		log.Println("Someting went wrong")
		json.NewEncoder(w).Encode(map[string]string{"message": "Sowthing Went Wong"})
		return
	}

	log.Println(userID)

	body := r.Body
	var userModel models.UserModel

	json.NewDecoder(body).Decode(&userModel)

	log.Println("Checking is email already exists")

	filter := bson.M{"email": userModel.Email}

	var result models.UserModel
	emailCheckError := common.GetDB().Collection(common.USER_COLLECTION).FindOne(context.Background(), filter).Decode(&result)
	log.Println(emailCheckError)
	if emailCheckError != nil {
		if emailCheckError == mongo.ErrNoDocuments {
			log.Println("Email does not exists!")
		}

	} else {
		log.Println("Email already exist", result.Email)
		json.NewEncoder(w).Encode(map[string]string{"message": "Email already Exists"})
		return
	}

	fmt.Println(result)

	userModel.CreatedAt = time.Now()
	userModel.UpdatedAt = time.Now()
	hashedPwd, err := models.HashPassword(userModel.Password)

	if err != nil {
		log.Println("Password Hashing Issue", err)
	}

	userModel.Password = hashedPwd

	fmt.Printf("Request body %+v", userModel)

	insertedRecord, err := common.GetDB().Collection(common.USER_COLLECTION).InsertOne(context.Background(), userModel)

	if err != nil {
		fmt.Println("Error while crating user ", err)
	}

	json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("user has been create successfully.with id %s", insertedRecord.InsertedID)})
}

func structToBsonD(data interface{}) bson.D {
	var result bson.D

	// Use reflection to iterate over the struct fields
	v := reflect.ValueOf(data)

	for i := 0; i < v.NumField(); i++ {
		// Get field type
		field := v.Type().Field(i)

		// Get the bson tag
		bsonTag := field.Tag.Get("bson")
		if bsonTag == "-" || bsonTag == "" || bsonTag == "_id" {
			// Skip fields with "-" or empty bson tag
			continue
		}
		// Append the field to bson.D if it is not empty or the zero value
		result = append(result, bson.E{Key: bsonTag, Value: v.Field(i).Interface()})
	}

	return result
}

func UpdateAdminUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	data := mux.Vars(r)

	id, exists := data["id"]

	body := r.Body

	var userModel models.UserModel

	json.NewDecoder(body).Decode(&userModel)

	if !exists {
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid request"})
		return
	}

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid object id"})
		return
	}

	filter := bson.M{"_id": objectID}

	var result models.UserModel

	filterError := common.GetDB().Collection(common.USER_COLLECTION).FindOne(context.TODO(), filter).Decode(&result)

	if filterError != nil {
		json.NewEncoder(w).Encode(map[string]string{"message": "No result found"})
		return
	}

	update := bson.M{"$set": structToBsonD(userModel)}

	res, err := common.GetDB().Collection(common.USER_COLLECTION).Database().Collection(common.USER_COLLECTION).UpdateOne(context.TODO(), filter, update)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprint("Error while updating", err)})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"message": "User Details Updated Successfully.", "data": res})
}

func DeleteAdminUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	id, idExists := vars["id"]

	if !idExists {
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid Request"})
		return
	}
	newObjectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid ID"})
		return
	}
	filter := bson.M{"_id": newObjectId}

	deletedRes := common.GetDB().Collection(common.USER_COLLECTION).FindOneAndDelete(context.TODO(), filter)

	if deletedRes.Err() != nil {
		if deletedRes.Err() == mongo.ErrNoDocuments {
			json.NewEncoder(w).Encode(map[string]string{"message": "No Record Found. Please try again."})
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"message": "There was an issue deleting the record. Please try again."})
		return
	}

	log.Println("User has beend deleted successfully")

	json.NewEncoder(w).Encode(map[string]string{"message": "user has been deleted successfully."})
}

func GetByIdAdminUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	id, idExists := vars["id"]

	if !idExists {
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid Request"})
		return
	}
	newObjectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid ID"})
		return
	}
	filter := bson.M{"_id": newObjectId}

	var result models.UserModel

	err = common.GetDB().Collection(common.USER_COLLECTION).FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		fmt.Println("Something Wentwrong", err)
		json.NewEncoder(w).Encode(map[string]string{"message": "Something went wrong"})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"message": "User Details found Successfully", "data": result})

}

func GetAllAdminUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var result []models.UserModel

	filter := bson.M{}

	cursor, err := common.GetDB().Collection(common.USER_COLLECTION).Find(context.TODO(), filter)

	defer cursor.Close(context.TODO())

	if err != nil {
		fmt.Println("Something Wentwrong", err)
		json.NewEncoder(w).Encode(map[string]string{"message": "Something went wrong"})
		return
	}

	cursorError := cursor.All(context.TODO(), &result)

	if cursorError != nil {
		fmt.Println("Something Wentwrong", cursorError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Something went wrong"})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"message": "User List fetched successfully.", "data": result})
}
