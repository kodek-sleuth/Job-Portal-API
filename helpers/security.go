package helpers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

func GenerateJWT() (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = "userid"
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("captainjacksparrowsayshi"))
	if err != nil {
		return "", err
	}
	return token, nil
}

func IsAuthorized(res http.ResponseWriter, req *http.Request, next http.HandlerFunc){

}

//func IsAuthorized(endpoint func(res http.ResponseWriter, req *http.Request)) http.Handler {
//	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
//		// check for token
//		if req.Header["Token"] != nil {
//			token, err := jwt.Parse(req.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
//				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//					ErrorResponse(res, http.StatusBadRequest, fmt.Sprintf("auth failed"))
//					return nil, nil
//				}
//				return mySigningKey, nil
//			})
//
//			if err != nil {
//				ErrorResponse(res, http.StatusBadRequest, fmt.Sprintf("wrong username or password %+v", err.Error()))
//			}
//
//			if token.Valid {
//				endpoint(res, req)
//			}
//
//		} else {
//			ErrorResponse(res, http.StatusBadRequest, fmt.Sprintf("wrong username or password"))
//		}
//	})
//}


