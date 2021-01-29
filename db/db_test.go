package db

import (
	"mycart/models"
	"testing"
)

func TestCreateConnection(t *testing.T) {
	var db Database
	if db.sql != nil {
		t.Error("SQL connection should be nil")
	}
	db = CreateConnection("test_ecommerce")
	if db.sql == nil {
		t.Error("SQL connection should not be nil")
	}
}

func TestInsertRow(t *testing.T) {
	db := CreateConnection("test_ecommerce")
	category := models.Category{}

	err := db.InsertRow("product", &category)

	if err == nil {
		t.Errorf("Row should not be inserted Required input: %T Got: %T ", &models.Product{}, &models.Category{})
	}
}
