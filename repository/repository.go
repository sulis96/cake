package repository

import (
	"CAKE-STORE/config"
	"CAKE-STORE/entity"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type (
	cakeRepository struct {
		DB *sql.DB
	}

	CakeRepository interface {
		GetListCake(ctx context.Context, title string) ([]entity.ListCake, error)
		GetDetailCake(ctx context.Context, id int) (entity.Cake, error)
		InsertCake(ctx context.Context, cake entity.Cake) error
		UpdateCake(ctx context.Context, id int, cake entity.Cake) error
		DeleteCake(ctx context.Context, id int) error
	}
)

func NewCakeRepository(database *sql.DB) CakeRepository {
	return &cakeRepository{
		DB: database,
	}
}

func (mcr *cakeRepository) GetDetailCake(ctx context.Context, id int) (entity.Cake, error) {
	var cake = entity.Cake{}
	db, err := config.MySqlDatabase()
	if err != nil {
		err = errors.New("CAN'T CONNECT TO MY SQL DB : " + err.Error())
		return cake, err
	}

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
	rowQuery, err := db.QueryContext(ctx, query, id)
	if err != nil {
		err = errors.New("CAN'T SELECT DATA FROM MY SQL DB :" + err.Error())
		return cake, err
	}

	for rowQuery.Next() {
		err = rowQuery.Scan(
			&cake.Id,
			&cake.Title,
			&cake.Description,
			&cake.Rating,
			&cake.CreatedAt,
			&cake.UpdatedAt)
		if err != nil {
			err = errors.New("CAN'T SCAN DATA FROM MY SQL DB : " + err.Error())
			return cake, err
		}
	}

	return cake, nil
}

func (mcr *cakeRepository) GetListCake(ctx context.Context, title string) ([]entity.ListCake, error) {
	var (
		cakes []entity.ListCake
		cake  = entity.ListCake{}
	)
	db, err := config.MySqlDatabase()
	if err != nil {
		err = errors.New("CAN'T CONNECT TO MY SQL DB : " + err.Error())
		return nil, err
	}

	query := `SELECT 
	id, 
	title
	FROM cake 
	WHERE title LIKE ?
	ORDER BY rating DESC, title ASC`
	rowQuery, err := db.QueryContext(ctx, query, "%"+title+"%")
	if err != nil {
		err = errors.New("CAN'T SELECT DATA FROM MY SQL DB :" + err.Error())
		return nil, err
	}

	for rowQuery.Next() {
		err = rowQuery.Scan(
			&cake.Id,
			&cake.Title)
		if err != nil {
			err = errors.New("CAN'T SCAN DATA FROM MY SQL DB : " + err.Error())
			return nil, err
		}
		cakes = append(cakes, cake)
	}

	return cakes, nil
}

func (cr *cakeRepository) InsertCake(ctx context.Context, cake entity.Cake) error {
	db, err := config.MySqlDatabase()
	if err != nil {
		err = errors.New("CAN'T CONNECT TO MY SQL DB : " + err.Error())
		return err
	}

	query := fmt.Sprintf("INSERT INTO `%v` (`title`, `description`, `rating`, `image`, `created_at`) VALUES('%v', '%v', '%v', '%v', current_timestamp())", "cake", cake.Title, cake.Description, cake.Rating, cake.Image)

	_, err = db.QueryContext(ctx, query)
	if err != nil {
		err = errors.New("FAILED TO INSERT DATA :" + err.Error())
		return err
	}

	return nil
}

func (cr *cakeRepository) UpdateCake(ctx context.Context, id int, cake entity.Cake) error {
	db, err := config.MySqlDatabase()
	if err != nil {
		err = errors.New("CAN'T CONNECT TO MY SQL DB : " + err.Error())
		return err
	}

	query := "UPDATE `cake` SET"
	if cake.Title != "" {
		query = query + "`title`= '" + cake.Title + "'"
	}
	if cake.Description != "" {
		query = query + ", `description`= '" + cake.Description + "'"
	}
	if cake.Image != "" {
		query = query + ", `image`='" + cake.Image + "'"
	}
	if cake.Rating != 0 {
		query = query + ", `rating`=" + fmt.Sprintf(`%v`, cake.Rating)
	}
	query = query + ", `updated_at`=current_timestamp() WHERE `id` =" + strconv.Itoa(id) + ";"

	_, err = db.QueryContext(ctx, query)
	if err != nil {
		err = errors.New("FAILED TO UPDATE DATA :" + err.Error())
		return err
	}

	return nil
}

func (cr *cakeRepository) DeleteCake(ctx context.Context, id int) error {
	db, err := config.MySqlDatabase()
	if err != nil {
		err = errors.New("CAN'T CONNECT TO MY SQL DB : " + err.Error())
		return err
	}

	query := fmt.Sprintf("DELETE FROM `%v` WHERE `id`=%v", "cake", id)

	_, err = db.QueryContext(ctx, query)
	if err != nil {
		err = errors.New("FAILED TO DELETE DATA :" + err.Error())
		return err
	}

	return nil
}
