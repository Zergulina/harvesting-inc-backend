package repository

import (
	"backend/internal/models"
	"database/sql"
	"errors"
)

func GetAllPosts(db *sql.DB) ([]models.Post, error) {
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	posts := []models.Post{}

	for rows.Next() {
		p := models.Post{}
		err := rows.Scan(&p.Id, &p.Name)
		if err != nil {
			continue
		}
		posts = append(posts, p)
	}

	return posts, nil
}

func GetPostById(db *sql.DB, id uint64) (*models.Post, error) {
	post := new(models.Post)
	row := db.QueryRow("SELECT * FROM posts WHERE id = $1", id)
	err := row.Scan(&post.Id, &post.Name)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func GetPostByName(db *sql.DB, name string) (*models.Post, error) {
	post := new(models.Post)
	row := db.QueryRow("SELECT * FROM posts WHERE name = $1", name)
	err := row.Scan(&post.Id, &post.Name)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func CreatePost(db *sql.DB, post *models.Post) (*models.Post, error) {
	row := db.QueryRow("INSERT INTO posts (name) VALUES ($1) RETURNING id", post.Name)
	err := row.Scan(&post.Id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func DeletePost(db *sql.DB, id uint64) error {
	_, err := db.Exec("DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePost(db *sql.DB, post *models.Post) (*models.Post, error) {
	result, err := db.Exec("UPDATE posts SET name = $1 WHERE id = $2", post.Name, post.Id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, errors.New("")
	}

	return post, nil
}

func ExistsPostByName(db *sql.DB, name string) (bool, error) {
	var isExist bool
	row := db.QueryRow("SELECT (EXISTS (SELECT FROM posts WHERE name = $1))", name)
	err := row.Scan(&isExist)
	if err != nil {
		return false, err
	}
	return isExist, nil
}
