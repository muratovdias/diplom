package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/muratovdias/diplom/internal/models"
)

type Comment interface {
	Create(comment models.Comment) error
	Delete(authorID, trainerID int) error
	GetCommentsByTrainerID(id int) ([]models.Comment, error)
}

type CommentRepo struct {
	db *sqlx.DB
}

func NewCommentRepo(db *sqlx.DB) *CommentRepo {
	return &CommentRepo{
		db: db,
	}
}

func (c *CommentRepo) Create(comment models.Comment) error {
	query := `INSERT INTO comments(author_id, trainer_id, text, date)
				VALUES($1, $2, $3, $4)`
	_, err := c.db.Exec(query, comment.AuthorID, comment.TrainerID, comment.Text, comment.Date)
	if err != nil {
		fmt.Printf("repo: comment: Create: %v", err)
		return err
	}
	return nil
}

func (c *CommentRepo) Delete(authorID, trainerID int) error {
	query := `DELETE 
			FROM comments
			WHERE author_id=$1 AND trainer_id=$2`
	_, err := c.db.Exec(query, authorID, trainerID)
	if err != nil {
		return err
	}
	return nil
}

func (c *CommentRepo) GetCommentsByTrainerID(id int) ([]models.Comment, error) {
	query := `SELECT author_id, trainer_id, full_name, text, date
			FROM comments c
			INNER JOIN users u
			ON c.author_id = u.user_id
			WHERE trainer_id=$1
			ORDER BY date DESC`
	comments := []models.Comment{}
	rows, err := c.db.Query(query, id)
	if err != nil {
		return comments, err
	}
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.AuthorID, &comment.TrainerID, &comment.Author, &comment.Text, &comment.Date); err != nil {
			return comments, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
