package repositoryMySql

import (
	"database/sql"
	"fmt"

	// "time"
	"gapai-skor-api/domain"

	"github.com/google/uuid"
)

type repoAnswerOption struct {
	DB *sql.DB
}

func CreateRepoAnswerOption(db *sql.DB) domain.AnswerOptionRepository {
	return &repoAnswerOption{DB: db}
}

func (repo *repoAnswerOption) GetAll() ([]domain.AnswerOption, error) {

	rows, err := repo.DB.Query("SELECT * FROM answer_options")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []domain.AnswerOption

	for rows.Next() {
		var answerOption domain.AnswerOption
		err := rows.Scan(&answerOption.ID, &answerOption.QuestionID, &answerOption.ContentAnswer, &answerOption.ImageURL,
			&answerOption.AudioURL, &answerOption.IsCorrect, &answerOption.CreatedAt, &answerOption.UpdatedAt)
		fmt.Println(err)
		if err != nil {
			return data, err
		}
		// answerOption.CreatedAt = time.Now().Add(24 * time.Hour)
		// answerOption.UpdatedAt = time.Now().Add(24 * time.Hour)
		data = append(data, answerOption)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data, err
}

func (repo *repoAnswerOption) GetByID(id string) (domain.AnswerOption, error) {
	row := repo.DB.QueryRow("SELECT * FROM answer_options where id=?", id)
	fmt.Println(id)

	var data domain.AnswerOption

	err := row.Scan(&data.ID, &data.QuestionID, &data.ContentAnswer, &data.ImageURL,
		&data.AudioURL, &data.IsCorrect, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		return data, err
	}
	// data.CreatedAt = time.Now().Add(24 * time.Hour)
	// data.UpdatedAt = time.Now().Add(24 * time.Hour)

	if err := row.Err(); err != nil {
		return domain.AnswerOption{}, err
	}
	// fmt.Println(data)
	return data, err
}

func (repo *repoAnswerOption) Create(tx *sql.Tx, answerOption *domain.AnswerOption) (err error) {
	newUUID, _ := uuid.NewRandom()
	// newUUID, _ := uuid.NewUUID()
	id := newUUID.String()
	query := "INSERT INTO answer_options (id, question_id, content_answer, image_url, audio_url, is_correct) values (?, ?, ?, ?, ?, ?)"
	if tx != nil {
		_, err = tx.Exec(query, id, answerOption.QuestionID, answerOption.ContentAnswer, answerOption.ImageURL,
			answerOption.AudioURL, answerOption.IsCorrect)
	} else {
		_, err = repo.DB.Exec(query, id, answerOption.QuestionID, answerOption.ContentAnswer, answerOption.ImageURL,
			answerOption.AudioURL, answerOption.IsCorrect)
	}
	if err != nil {
		return err
	}
	return err
}
