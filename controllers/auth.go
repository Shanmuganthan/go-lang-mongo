package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Shanmuganthan/go-lang-mongo/common"
	"github.com/Shanmuganthan/go-lang-mongo/models"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

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

	userClaims := models.UserClaims{
		Id:    userModel.Id.Hex(),
		Email: userModel.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 7, 0)),
		},
	}

	jwtTOken, err := common.GenerateJWTToken(userClaims)

	if err != nil {
		log.Println("Something went wrong in token genration!", err)
		json.NewEncoder(w).Encode(map[string]string{"message": "Something went wrong"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Success", "token": jwtTOken})

}
