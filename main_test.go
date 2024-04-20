package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	languages := []Language{
		{Code: "ZH", Name: "Chinese"},
		{Code: "EN", Name: "English"},
		{Code: "DE", Name: "German"},
	}
	DB.Create(&languages)

	user := User{Name: "jinzhu", Languages: []Language{
		{Code: "ZH"},
		{Code: "EN"},
	}}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
