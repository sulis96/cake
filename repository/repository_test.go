package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetDetailCake(t *testing.T) {
	mockDB, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	cr := cakeRepository{DB: mockDB}

}
