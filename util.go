package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func defaultUsers() {
	admin := User{
		Id:        "0",
		Role:      "ADMIN",
		FirstName: "Admin",
		LastName:  "Temporary",
		Email:     "admin@temp.com",
		Password:  "admin",
		Active:    true,
	}

	db.Set("user", "0", admin)

	fmt.Printf("\nTemporary admin credentials:\n\n\tEmail:\t\t%s\n\tPassword:\t%s\n\n", admin.Email, admin.Password)
}

func genId() string {
	return strconv.Itoa(int(time.Now().UnixNano()))
}

func toBase64Json(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(b)
}
