package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"../db"
	"../model"
	"io/ioutil"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

var loaded = false

// Registers a user. Takes in a vault key, email, and empty accounts array
// Checks if the email already exists, if so, send back error code
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	var res model.ResponseResult

	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	collection, err := db.GetDBCollection()

	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	var result model.User
	err = collection.FindOne(context.TODO(), bson.D{{"email", user.Email}}).Decode(&result)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			// insert user
			_, err = collection.InsertOne(context.TODO(), user)
			if err != nil {
				res.Error = "Error While Creating User, Try Again"
				json.NewEncoder(w).Encode(res)
				return
			}

			res.Result = "Registration Successful"
			w.WriteHeader(201)
			json.NewEncoder(w).Encode(res)
			return
		}

		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	// Runs if user with the specified email already exists
	res.Error = "Email Already Exists in Database"
	w.WriteHeader(422)
	json.NewEncoder(w).Encode(res)
	return
}

// Login method; checks id and returns auth token
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var result model.ResponseResult
	var res model.ResponseResult

	// Check if user is already authed
	isAuth, message, _ := isLoggedIn(r)
	if isAuth {
		res.Result = message
		json.NewEncoder(w).Encode(res)
		return
	}

	// Ok, user it not authed. Continue with login process...
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	// Get Database
	collection, err := db.GetDBCollection()

	if err != nil {
		log.Fatal(err)
	}

	// Check to see if user is in the database
	err = collection.FindOne(context.TODO(), bson.D{{"id", user.Id}}).Decode(&result)

	if err != nil {
		res.Error = "Invalid id"
		json.NewEncoder(w).Encode(res)
		return
	}

	// Create Auth Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["id"] = user.Id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(getSecret()))
	if err != nil {
		res.Error = "Auth Failed"
		json.NewEncoder(w).Encode(res)
	}

	result.Token = token
	json.NewEncoder(w).Encode(result)
}

// TODO: METHOD RETURNS A VAULT TO A USER

func UpdateHandler(w http.ResponseWriter, r *http.Request) { // TODO: UPDATE - MUST HAVE VALID AUTH -> UPDATES VAULT
	w.Header().Set("Content-Type", "application/json")

	// is user authed?
	isAuth, message, id := isLoggedIn(r)

	fmt.Println("IsAuth: ", isAuth)
	fmt.Println("Message: " + message)
	fmt.Println("ID: " + id)

	var result model.ResponseResult

	// User is not authed, send error and return
	if !isAuth {
		result.Error = message
		json.NewEncoder(w).Encode(result)
		return
	}

	// Get JSON Message

	// Convert to model

	// Update MongoDB

	// Return success
}

// Checks to see if a user is authenticated
// Returns <isAuthed, Error/Success message, id (if authed)>
func isLoggedIn(r *http.Request) (bool, string, string) {
	// Get auth code
	tokenString := r.Header.Get("Authorization")

	// Parse code to ensure its the required format
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(getSecret()), nil
	})

	var result model.User
	var res model.ResponseResult

	// If token is not valid
	if token == nil {
		res.Error = "Invalid Authentication Key"
		//json.NewEncoder(w).Encode(res)
		return false, "Invalid Authentication Key", ""
	}

	// If token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		result.Id = claims["id"].(string)

		return true, "User is Authenticated", result.Id
	}

	return false, err.Error(), ""
}

func getSecret() string {
	if !loaded {
		// Load Config File
		err := godotenv.Load("config.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		loaded = !loaded
	}

	return os.Getenv("ACCESS_SECRET")
}

// TODO: LOGOUT

/*
{
	"id": "asngasg",
	"email": "legitzx@gmail.com",
	"accounts": [
		{
			"url": "google.com",
			"name": "Google",
			"username": "coolUserName",
			"password": "coolPassword"
		},
		{
			"url": "google.com",
			"name": "Google",
			"username": "coolUserName",
			"password": "coolPassword"
		}
	]
}
*/
