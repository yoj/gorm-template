package main

import (
	"fmt"

	"github.com/dbtry/dbmodules"
)

type admin struct {
	ID             int    `json:"id"`
	Email          string `json:"emial"`
	PasswordDigest string `json:"paaswordDigest"`
	LastName       string `json:"lastName"`
	FirstName      string `json:"firstName"`
	LastNameKana   string `json:"lastNameKana"`
	FirstNameKana  string `json:"firstNameKane"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
}

func main() {
	db := dbmodules.GormConnect()
	defer db.Close()

	admin := []admin{}

	db.Find(&admin)

	for _, r := range admin {
		fmt.Println(r.Email)
	}
}
