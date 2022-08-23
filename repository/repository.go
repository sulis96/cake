package repository

import (
	"CAKE-STORE/entity"
	"context"
	"database/sql"
	"errors"
	"fmt"

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

func (cr *cakeRepository) GetDetailCake(ctx context.Context, id int) (entity.Cake, error) {
	var cake entity.Cake
	db, err := cr.DB.Conn(ctx)
	if err != nil {
		err = errors.New("GET DETAIL CAKE -> CAN'T CONNECT TO MY SQL DB : " + err.Error())
		return cake, err
	}
	defer db.Close()

	query := `SELECT 
	id, 
	title, 
	description, 
	COALESCE(rating,0.00),
	image,
	created_at, 
	COALESCE(updated_at, created_at)
	FROM cake 
	WHERE id = ?
	ORDER BY rating DESC, title ASC`
	row, err := db.QueryContext(ctx, query, id)
	if err != nil {
		err = errors.New("GET DETAIL CAKE -> CAN'T SELECT DATA FROM MY SQL DB :" + err.Error())
		return cake, err
	}
	defer row.Close()

	for row.Next() {
		err = row.Scan(
			&cake.Id,
			&cake.Title,
			&cake.Description,
			&cake.Rating,
			&cake.Image,
			&cake.CreatedAt,
			&cake.UpdatedAt)
		if err != nil {
			err = errors.New("GET DETAIL CAKE -> CAN'T SCAN DATA FROM MY SQL DB : " + err.Error())
			return cake, err
		}
	}

	if cake.Title == "" {
		err = errors.New("GET DETAIL CAKE -> NO DATA WITH ID = " + fmt.Sprintf("%v", id))
		return cake, err
	}

	return cake, nil
}

func (cr *cakeRepository) GetListCake(ctx context.Context, title string) ([]entity.ListCake, error) {
	var (
		cakes []entity.ListCake
		cake  = entity.ListCake{}
	)
	db, err := cr.DB.Conn(ctx)
	if err != nil {
		err = errors.New("GET LIST CAKE -> CAN'T CONNECT TO MY SQL DB : " + err.Error())
		return nil, err
	}
	defer db.Close()

	query := `SELECT 
	id, 
	title
	FROM cake 
	WHERE title LIKE ?
	ORDER BY rating DESC, title ASC`
	row, err := db.QueryContext(ctx, query, "%"+title+"%")
	if err != nil {
		err = errors.New("GET LIST CAKE -> CAN'T SELECT DATA FROM MY SQL DB :" + err.Error())
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(
			&cake.Id,
			&cake.Title)
		if err != nil {
			err = errors.New("GET LIST CAKE -> CAN'T SCAN DATA FROM MY SQL DB : " + err.Error())
			return nil, err
		}
		cakes = append(cakes, cake)
	}

	return cakes, nil
}

func (cr *cakeRepository) InsertCake(ctx context.Context, cake entity.Cake) error {
	db, err := cr.DB.Conn(ctx)
	if err != nil {
		err = errors.New("INSERT CAKE -> CAN'T CONNECT TO MY SQL DB : " + err.Error())
		return err
	}
	defer db.Close()

	query := fmt.Sprintf("INSERT INTO `%v` (`title`, `description`, `rating`, `image`, `created_at`) VALUES('%v', '%v', '%v', '%v', current_timestamp())", "cake", cake.Title, cake.Description, cake.Rating, cake.Image)

	row, err := db.QueryContext(ctx, query)
	if err != nil {
		err = errors.New("INSERT CAKE -> FAILED TO INSERT DATA :" + err.Error())
		return err
	}
	defer row.Close()

	return nil
}

func (cr *cakeRepository) UpdateCake(ctx context.Context, id int, cake entity.Cake) error {
	db, err := cr.DB.Conn(ctx)
	if err != nil {
		err = errors.New("UPDATE CAKE -> CAN'T CONNECT TO MY SQL DB : " + err.Error())
		return err
	}
	defer db.Close()

	var existCake = entity.Cake{}
	err = db.QueryRowContext(ctx, "SELECT title, description, rating, image from cake WHERE id = ?", id).Scan(&existCake.Title, &existCake.Description, &existCake.Rating, &existCake.Image)

	if existCake.Title == "" {
		err = errors.New("UPDATE CAKE -> NO DATA WITH ID =" + fmt.Sprintf("%v", id))
		return err
	}
	if err != nil {
		err = errors.New("UPDATE CAKE -> ERROR WHEN SCAN DATA: " + err.Error())
		return err
	}

	var query = fmt.Sprintf("UPDATE `%v` SET `updated_at` = current_timestamp(), `title`='%v', `description`='%v', `rating`=%v, `image`='%v' WHERE `id`=%v", "cake", cake.Title, cake.Description, cake.Rating, cake.Image, id)

	row, err := db.QueryContext(ctx, query)
	if err != nil {
		err = errors.New("UPDATE CAKE -> FAILED TO UPDATE DATA :" + err.Error())
		return err
	}
	defer row.Close()
	return nil
}

func (cr *cakeRepository) DeleteCake(ctx context.Context, id int) error {
	db, err := cr.DB.Conn(ctx)
	if err != nil {
		err = errors.New("DELETE CAKE -> CAN'T CONNECT TO MY SQL DB : " + err.Error())
		return err
	}
	defer db.Close()

	var title string
	err = db.QueryRowContext(ctx, "SELECT title from cake WHERE id = ?", id).Scan(&title)
	if title == "" {
		err = errors.New("DELETE CAKE -> NO DATA WITH ID =" + fmt.Sprintf("%v", id))
		return err
	}
	if err != nil {
		err = errors.New("DELETE CAKE -> ERROR WHEN SCAN DATA: " + err.Error())
		return err
	}

	query := fmt.Sprintf("DELETE FROM `%v` WHERE `id`=%v", "cake", id)

	_, err = cr.DB.QueryContext(ctx, query)
	if err != nil {
		err = errors.New("DELETE CAKE -> FAILED TO DELETE DATA :" + err.Error())
		return err
	}

	return nil
}
