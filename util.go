package main

import (
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
