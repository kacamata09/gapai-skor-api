package repositoryMySql

import (
	"database/sql"
	"fmt"

	// "time"
	"gapai-skor-api/domain"
)

type repoQuestion struct {
	DB *sql.DB
}

func CreateRepoQuestion(db *sql.DB) domain.QuestionRepository {
	return &repoQuestion{DB: db}
}

func (repo *repoQuestion) GetAll() ([]domain.Question, error) {

	rows, err := repo.DB.Query("SELECT * FROM questions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []domain.Question

	for rows.Next() {
		var question domain.Question
		err := rows.Scan(&question.ID, &question.TestID, &question.ContentQuestion, &question.ImageURL,
			&question.AudioURL, &question.QuestionType, &question.Points, &question.CreatedAt, &question.UpdatedAt)
		fmt.Println(err)
		if err != nil {
			return data, err
		}
		// question.CreatedAt = time.Now().Add(24 * time.Hour)
		// question.UpdatedAt = time.Now().Add(24 * time.Hour)
		data = append(data, question)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data, err
}

func (repo *repoQuestion) GetByID(id string) (domain.Question, error) {
	row := repo.DB.QueryRow("SELECT * FROM questions where id=?", id)
	fmt.Println(id)

	var data domain.Question

	err := row.Scan(&data.ID, &data.TestID, &data.ContentQuestion, &data.ImageURL,
		&data.AudioURL, &data.QuestionType, &data.Points, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		return data, err
	}
	// data.CreatedAt = time.Now().Add(24 * time.Hour)
	// data.UpdatedAt = time.Now().Add(24 * time.Hour)

	if err := row.Err(); err != nil {
		return domain.Question{}, err
	}
	// fmt.Println(data)
	return data, err
}

func (repo *repoQuestion) Create(tx *sql.Tx, question *domain.Question) (err error) {
	query := "INSERT INTO questions (test_id, content_question, image_url, audio_url, question_type, points) values (?, ?, ?, ?, ?, ?)"
	if tx != nil {
		_, err = tx.Exec(query, question.TestID, question.ContentQuestion, question.ImageURL,
			question.AudioURL, question.QuestionType, question.Points)
	} else {
		_, err = repo.DB.Exec(query, question.TestID, question.ContentQuestion, question.ImageURL,
			question.AudioURL, question.QuestionType, question.Points)
	}
	if err != nil {
		return err
	}
	return err
}
