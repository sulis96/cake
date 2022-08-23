package repository

import (
	"CAKE-STORE/entity"
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

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

	mockDB.ExpectQuery(regexp.QuoteMeta(`SELECT id, title FROM cake WHERE title LIKE ? ORDER BY rating DESC, title ASC`)).WillReturnRows(sqlmock.NewRows([]string{"1", "title"}).AddRow(1, "cake"))
	r, e := cr.GetListCake(ctx, "cake")
	assert.Nil(t, e)
	assert.NotNil(t, r)
}

func TestGetDetailCake(t *testing.T) {
	db, mockDB, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cr := cakeRepository{DB: db}

	query := `SELECT 
	id, 
	title, 
	description, 
	COALESCE(rating,0.00),
	created_at, 
	COALESCE(updated_at, created_at)
	FROM cake 
	WHERE id = ?
	ORDER BY rating DESC, title ASC`

	fmt.Println(time.Now())

	mockDB.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(sqlmock.NewRows([]string{"1", "cake", "desc", "2.0", "2022-08-23 17:31:30.7193691", "2022-08-23 17:31:30.7193691"}).AddRow(1, "cake", "desc", 2.0, time.Now(), time.Now()))
	r, e := cr.GetDetailCake(ctx, 1)
	assert.Nil(t, e)
	assert.NotNil(t, r)
}

func TestAddNewCake(t *testing.T) {
	db, mockDB, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cake := entity.Cake{Title: "cake", Description: "desc", Rating: 2.0, Image: "img"}

	cr := cakeRepository{DB: db}

	query := fmt.Sprintf("INSERT INTO `%v` (`title`, `description`, `rating`, `image`, `created_at`) VALUES('%v', '%v', '%v', '%v', current_timestamp())", "cake", cake.Title, cake.Description, cake.Rating, cake.Image)

	mockDB.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(mockDB.NewRows([]string{}))
	e := cr.InsertCake(ctx, cake)
	assert.Nil(t, e)
}
