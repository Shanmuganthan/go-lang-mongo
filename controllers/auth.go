package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Shanmuganthan/go-lang-mongo/common"
	"github.com/Shanmuganthan/go-lang-mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(w http.ResponseWriter, r *http.Request) {

	log.Println("Login Function Called")

	var loginBody models.LoginModel

	body := r.Body

	err := json.NewDecoder(body).Decode(&loginBody)

	if err != nil {
		log.Println("Error While decoding the data", err)
		json.NewEncoder(w).Encode(map[string]string{"message": "Something went wrong while decoding the body"})
		return
	}

	fmt.Printf("Values %#v", loginBody)

	filter := bson.M{"email": loginBody.Email}

	var userModel models.LoginModel

	findOneError := common.GetDB().Collection(common.USER_COLLECTION).FindOne(context.TODO(), filter).Decode(&userModel)

	if findOneError != nil {
		if findOneError == mongo.ErrNoDocuments {
			log.Printf("%s Username doesnt exists", loginBody.Email)
			json.NewEncoder(w).Encode(map[string]string{"message": "Username doesnt exists"})
			return
		}
	}

	isNotSame := models.VerifyPassword(userModel.Password, loginBody.Password)

	if isNotSame {
		log.Println("Incorrect Password!", isNotSame)
		json.NewEncoder(w).Encode(map[string]string{"message": "Incorrect Username or Password"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Success"})

}
