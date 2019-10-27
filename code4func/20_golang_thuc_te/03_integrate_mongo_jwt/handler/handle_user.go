package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"github.com/dgrijalva/jwt-go"

	config "go-jwt/config"
	driver "go-jwt/driver"
	models "go-jwt/model"
	repoImpl "go-jwt/repository/repoimpl"
)

var jwtKey = []byte("abcdefghjiklmnopq")

type Claims struct {
	Email string `json:"email"`
	DisplayName string `json:"displayName"`
	jwt.StandardClaims
}

func Register(w http.ResponseWriter, r *http.Request) {
	var regData models.RegistrationData
	err := json.NewDecoder(r.Body).Decode(&regData)
	if err != nil {
		ResponseErr(w, http.StatusBadRequest)
		return
	}
	findUserExist, err := repoImpl.NewUserRepo(driver.Mongo.Client.Database(config.DB_NAME)).FindUserByEmail(regData.Email)
	fmt.Println("findUserExist =", findUserExist.Email)
	fmt.Println("err = ", err)
	if findUserExist.Email == regData.Email {
		ResponseErr(w, http.StatusConflict)
		return
	}

	user := models.User{
		Email: regData.Email,
		Password: regData.Password,
		DisplayName: regData.DisplayName,
	}
	err = repoImpl.NewUserRepo(driver.Mongo.Client.Database(config.DB_NAME)).Insert(user)
	if err != nil {
		ResponseErr(w, http.StatusInternalServerError)
		return
	}

	var tokenString string
	tokenString, err = GenToken(user)
	if err != nil {
		ResponseErr(w, http.StatusInternalServerError)
	}


	ResponseOk(w, models.RegisterResponse{
		Token: tokenString,
		Status: http.StatusOK,
	})
}

func Login(w http.ResponseWriter, r *http.Request){
	var loginData models.LoginData
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		ResponseErr(w, http.StatusBadRequest)
		return
	}

	var user models.User
	user, err = repoImpl.NewUserRepo(driver.Mongo.Client.Database(config.DB_NAME)).CheckLoginInfo(loginData.Email, loginData.Password)
	if err != nil {
		ResponseErr(w, http.StatusUnauthorized)
		return
	}

	var tokenString string
	tokenString, err = GenToken(user)
	if err != nil {
		ResponseErr(w, http.StatusInternalServerError)
		return
	}

	ResponseOk(w, models.RegisterResponse{
		Token: tokenString,
		Status: http.StatusOK,
	})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	tokenHeader := r.Header.Get("Authorization")

	if tokenHeader == "" {
		ResponseErr(w, http.StatusForbidden)
		return
	}
	splitted := strings.Split(tokenHeader, " ") // Bearer jwt_token
	if len(splitted) != 2 {
		ResponseErr(w, http.StatusForbidden)
		return
	}

	tokenPart := splitted[1]
	tk := &Claims{}

	token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	//fmt.Println("token = ", token)
	if err != nil {
		fmt.Println("err = ", err)
		ResponseErr(w, http.StatusInternalServerError)
		return
	}

	if token.Valid {
		ResponseOk(w, token.Claims)
	}
}

func GenToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(120 * time.Second)
	claims := &Claims{
		Email:          user.Email,
		DisplayName:   	user.DisplayName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

func ResponseErr(w http.ResponseWriter, statusCode int)  {
	jData, err := json.Marshal(models.Error{
		Status: statusCode,
		Message: http.StatusText(statusCode),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func ResponseOk(w http.ResponseWriter, data interface{})  {
	if data == nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	jData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}