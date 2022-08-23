package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetListCake(t *testing.T) {
	db, mockDB, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cr := cakeRepository{DB: db}
	var list = []string{"cake"}

	mockDB.ExpectBegin()
	mockDB.ExpectQuery("SELECT id, title FROM cake WHERE title LIKE ? ORDER BY rating DESC, title ASC").WithArgs("cake").WillReturnRows(sqlmock.NewRows(list))
	r, e := cr.GetListCake(ctx, "cake")
	assert.Nil(t, e)
	assert.NotNil(t, r)
}
