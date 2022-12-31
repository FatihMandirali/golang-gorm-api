package services

import (
	. "appword-api/models"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)
import . "appword-api/utils"

// NOT: SERVİS İSİMLERİ BÜYÜK HARF İLE BAŞLAMALI YOKSA MUX GÖREMİYOR HANDLEFUNC DERLER.. GetUsers gibi
// GET - USER
func GetUsers(w http.ResponseWriter, r *http.Request) {
	c := r.Header.Get("Authorization")
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(c, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	CheckError(err)

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

/*func GetUsers(w http.ResponseWriter, r *http.Request) {
	c := r.Header.Get("Authorization")
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(c, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	rows, err := Db.Query(`select * from "Users"`)
	CheckError(err)
	defer rows.Close()

	var users []*User
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.Id, &user.Role, &user.UserName, &user.Email, &user.Password, &user.CreateDate, &user.ModifiedDate, &user.IsActive, &user.IsDeleted)
		CheckError(err)
		users = append(users, user)
	}
	for _, usr := range users {
		fmt.Printf("%d - %s - %s - %s \n", usr.Id, usr.UserName, usr.Email, usr.Password)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}*/
