package repositoryMySql

import (
	"database/sql"
	"fmt"

	// "time"
	"gapai-skor-api/domain"

	"github.com/google/uuid"
)

type repoAttempt struct {
	DB *sql.DB
}

func CreateRepoAttempt(db *sql.DB) domain.AttemptRepository {
	return &repoAttempt{DB: db}
}

func (repo *repoAttempt) GetAll() ([]domain.Attempt, error) {

	rows, err := repo.DB.Query("SELECT * FROM attempts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []domain.Attempt

	for rows.Next() {
		var attempt domain.Attempt
		err := rows.Scan(&attempt.ID, &attempt.UserID, &attempt.TestID, &attempt.Score,
			&attempt.AttemptDate, &attempt.CreatedAt, &attempt.UpdatedAt)
		fmt.Println(err)
		if err != nil {
			return data, err
		}
		// attempt.CreatedAt = time.Now().Add(24 * time.Hour)
		// attempt.UpdatedAt = time.Now().Add(24 * time.Hour)
		data = append(data, attempt)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data, err
}

func (repo *repoAttempt) GetAttemptWithAttemptAnswer(id string) (domain.Attempt, error) {

	query := `
	SELECT 
		att.id AS attempt_id,
		att.test_id AS test_id,
		att.user_id AS user_id,
		att.score AS score,
		att.attempt_date AS attempt_date,
		att.updated_at AS updated_at,
		aa.id AS attempt_answer_id,
		aa.attempt_id AS aa_attempt_id,
		aa.question_id AS aa_question_id,
		aa.selected_answer_option_id AS selected_answer_option_id,
		ao.id AS answer_option_id,
		ao.is_correct AS is_correct,
		q.id AS question_id,
		q.question_type AS question_type
	FROM 
		attempts att
	LEFT JOIN 
		attempt_answers aa
	ON 
		att.id = aa.attempt_id
	LEFT JOIN 
		answer_options ao
	ON
		ao.id = aa.selected_answer_option_id
	LEFT JOIN 
		questions q
	ON
		q.id = aa.question_id
	WHERE 
		att.id = ?;
		`

	var data domain.Attempt
	var attemptAnswers []domain.AttemptAnswer
	rows, err := repo.DB.Query(query, id)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		var attemptAnswer domain.AttemptAnswer
		var answerOption domain.AnswerOption
		var question domain.Question

		err := rows.Scan(&data.ID, &data.TestID, &data.UserID, &data.Score,
			&data.AttemptDate, &data.UpdatedAt, &attemptAnswer.ID, &attemptAnswer.AttemptID,
			&answerOption.QuestionID, &attemptAnswer.SelectedAnswerOptionID, &answerOption.ID,
			&answerOption.IsCorrect, &question.ID, &question.QuestionType)
		fmt.Println(err)
		if err != nil {
			return data, err
		}
		attemptAnswer.IsCorrect = answerOption.IsCorrect
		attemptAnswer.QuestionType = question.QuestionType
		fmt.Println(attemptAnswer)

		// attempt.CreatedAt = time.Now().Add(24 * time.Hour)
		// attempt.UpdatedAt = time.Now().Add(24 * time.Hour)
		attemptAnswers = append(attemptAnswers, attemptAnswer)
	}

	data.AttemptAnswers = attemptAnswers
	if err := rows.Err(); err != nil {
		return data, err
	}
	return data, err
}

func (repo *repoAttempt) GetAttemptHistory(id string) ([]domain.Attempt, error) {

	query := `
	SELECT 
		att.id AS attempt_id,
		att.test_id AS test_id,
		att.user_id AS user_id,
		att.score AS score,
		att.attempt_date AS attempt_date,
		att.updated_at AS updated_at,
		t.id AS test_id,
		t.test_title AS test_title
	FROM 
		attempts att
	LEFT JOIN 
		tests t
	ON 
		t.id = att.test_id
	WHERE 
		att.user_id = ?;
		`

	var data []domain.Attempt
	rows, err := repo.DB.Query(query, id)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		var attempt domain.Attempt
		var test domain.Test

		err := rows.Scan(&attempt.ID, &attempt.TestID, &attempt.UserID, &attempt.Score,
			&attempt.AttemptDate, &attempt.UpdatedAt, &test.ID, &attempt.TestTitle)
		fmt.Println(err)
		if err != nil {
			return data, err
		}

		// attempt.TestTitle = test.TestTitle
		data = append(data, attempt)
	}

	if err := rows.Err(); err != nil {
		return data, err
	}
	return data, err
}

func (repo *repoAttempt) GetAttemptTestUser(id string) ([]domain.Attempt, error) {

	query := `
	SELECT 
		att.id,
		att.test_id,
		att.user_id,
		att.score,
		att.attempt_date,
		att.updated_at,
		u.fullname,
		u.phone,
		u.email
	FROM 
		attempts att
	LEFT JOIN 
		users u
	ON 
		u.id = att.user_id
	WHERE 
		att.test_id = ?;
		`

	var data []domain.Attempt
	rows, err := repo.DB.Query(query, id)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		var attempt domain.Attempt

		err := rows.Scan(&attempt.ID, &attempt.TestID, &attempt.UserID, &attempt.Score,
			&attempt.AttemptDate, &attempt.UpdatedAt, &attempt.FUllname, &attempt.Phone, &attempt.Email)
		fmt.Println(err)
		if err != nil {
			return data, err
		}

		// attempt.TestTitle = test.TestTitle
		data = append(data, attempt)
	}

	if err := rows.Err(); err != nil {
		return data, err
	}
	return data, err
}

func (repo *repoAttempt) GetByID(id string) (domain.Attempt, error) {
	row := repo.DB.QueryRow("SELECT * FROM attempts where id=?", id)
	fmt.Println(id)

	var data domain.Attempt

	err := row.Scan(&data.ID, &data.UserID, &data.TestID, &data.Score,
		&data.AttemptDate, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		return data, err
	}
	// data.CreatedAt = time.Now().Add(24 * time.Hour)
	// data.UpdatedAt = time.Now().Add(24 * time.Hour)

	if err := row.Err(); err != nil {
		return domain.Attempt{}, err
	}
	// fmt.Println(data)
	return data, err
}

func (repo *repoAttempt) VerifAttemptIsThere(attempt *domain.Attempt) (id int, err error) {

	// query := "SELECT id, user_id, test_id FROM attempts WHERE user_id = ? AND test_id = ?"
	// row := repo.DB.QueryRow(query, attempt.UserID, attempt.TestID)

	// var data domain.Attempt
	// err = row.Scan(&data.ID, &data.UserID, &data.TestID)

	// fmt.Println(data, "attempt id")

	// if err != nil {
	// 	return "", err
	// }

	// id = data.ID

	// return id, err

	query := "SELECT id, user_id, test_id FROM attempts WHERE user_id = ? AND test_id = ?"
	rows, err := repo.DB.Query(query, attempt.UserID, attempt.TestID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	lenRow := 0
	for rows.Next() {
		lenRow++
	}

	if err = rows.Err(); err != nil {
		return 0, err
	}

	return lenRow, nil

}

func (repo *repoAttempt) Create(tx *sql.Tx, attempt *domain.Attempt) (id string, err error) {
	newUUID, _ := uuid.NewRandom()
	// newUUID, _ := uuid.NewUUID()
	id = newUUID.String()

	query := "INSERT INTO attempts (id, user_id, test_id, score) values (?, ?, ?, ?)"
	if tx != nil {
		_, err = tx.Exec(query, id, attempt.UserID, attempt.TestID, attempt.Score)
	} else {
		_, err = repo.DB.Exec(query, id, attempt.UserID, attempt.TestID, attempt.Score)
	}
	if err != nil {
		return "", err
	}
	return id, err
}

func (repo *repoAttempt) Update(attempt *domain.Attempt) (err error) {
	query := "UPDATE attempts SET score = ? WHERE id = ?"

	_, err = repo.DB.Exec(query, attempt.Score, attempt.ID)

	if err != nil {
		return err
	}
	return err
}

func (repo *repoAttempt) Delete(id string) (err error) {

	query := `
		DELETE FROM attempts
		WHERE id = ?;
		`

	_, err = repo.DB.Exec(query, id)

	if err != nil {
		return err
	}
	return err
}
