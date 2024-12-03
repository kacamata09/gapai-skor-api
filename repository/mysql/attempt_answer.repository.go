package repositoryMySql

import (
	"database/sql"
	"fmt"

	// "time"
	"gapai-skor-api/domain"

	"github.com/google/uuid"
)

type repoAttemptAnswer struct {
	DB *sql.DB
}

func CreateRepoAttemptAnswer(db *sql.DB) domain.AttemptAnswerRepository {
	return &repoAttemptAnswer{DB: db}
}

func (repo *repoAttemptAnswer) GetAll() ([]domain.AttemptAnswer, error) {

	rows, err := repo.DB.Query("SELECT * FROM attempt_answers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []domain.AttemptAnswer

	for rows.Next() {
		var attemptAnswer domain.AttemptAnswer
		err := rows.Scan(&attemptAnswer.ID, &attemptAnswer.AttemptID, &attemptAnswer.QuestionID,
			&attemptAnswer.SelectedAnswerOptionID, &attemptAnswer.CreatedAt, &attemptAnswer.UpdatedAt)
		fmt.Println(err)
		if err != nil {
			return data, err
		}
		// attemptAnswer.CreatedAt = time.Now().Add(24 * time.Hour)
		// attemptAnswer.UpdatedAt = time.Now().Add(24 * time.Hour)
		data = append(data, attemptAnswer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data, err
}

func (repo *repoAttemptAnswer) GetByID(id string) (domain.AttemptAnswer, error) {
	row := repo.DB.QueryRow("SELECT * FROM attempt_answers where id=?", id)
	fmt.Println(id)

	var data domain.AttemptAnswer

	err := row.Scan(&data.ID, &data.AttemptID, &data.QuestionID,
		&data.SelectedAnswerOptionID, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		return data, err
	}
	// data.CreatedAt = time.Now().Add(24 * time.Hour)
	// data.UpdatedAt = time.Now().Add(24 * time.Hour)

	if err := row.Err(); err != nil {
		return domain.AttemptAnswer{}, err
	}
	// fmt.Println(data)
	return data, err
}

func (repo *repoAttemptAnswer) Create(tx *sql.Tx, attemptAnswer *domain.AttemptAnswer) (err error) {
	newUUID, _ := uuid.NewRandom()
	// newUUID, _ := uuid.NewUUID()
	id := newUUID.String()

	query := "INSERT INTO attempt_answers (id, attempt_id, question_id, selected_answer_option_id) values (?, ?, ?, ?)"
	if tx != nil {
		_, err = tx.Exec(query, id, attemptAnswer.AttemptID, attemptAnswer.QuestionID,
			attemptAnswer.SelectedAnswerOptionID)
	} else {
		_, err = repo.DB.Exec(query, attemptAnswer.AttemptID, attemptAnswer.QuestionID,
			attemptAnswer.SelectedAnswerOptionID)
	}
	if err != nil {
		return err
	}
	return err
}
