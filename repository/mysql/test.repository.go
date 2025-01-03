package repositoryMySql

import (
	"database/sql"
	"fmt"

	// "time"
	"gapai-skor-api/domain"

	"github.com/google/uuid"
)

type repoTest struct {
	DB *sql.DB
}

func CreateRepoTest(db *sql.DB) domain.TestRepository {
	return &repoTest{DB: db}
}

func (repo *repoTest) GetAll() ([]domain.Test, error) {

	rows, err := repo.DB.Query("SELECT * FROM tests")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []domain.Test

	for rows.Next() {
		var test domain.Test
		err := rows.Scan(&test.ID, &test.TestTitle, &test.TestCode, &test.Description,
			&test.CreatedBy, &test.Duration, &test.CreatedAt, &test.UpdatedAt)
		fmt.Println(err)
		if err != nil {
			return data, err
		}
		// test.CreatedAt = time.Now().Add(24 * time.Hour)
		// test.UpdatedAt = time.Now().Add(24 * time.Hour)
		data = append(data, test)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data, err
}

func (repo *repoTest) GetByID(id string) (domain.Test, error) {
	row := repo.DB.QueryRow("SELECT * FROM tests where id=?", id)
	fmt.Println(id)

	var data domain.Test

	err := row.Scan(&data.ID, &data.TestTitle, &data.TestCode, &data.Description,
		&data.CreatedBy, &data.Duration, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		return data, err
	}
	// data.CreatedAt = time.Now().Add(24 * time.Hour)
	// data.UpdatedAt = time.Now().Add(24 * time.Hour)

	if err := row.Err(); err != nil {
		return domain.Test{}, err
	}
	// fmt.Println(data)
	return data, err
}

func (repo *repoTest) GetByTestCode(testCode string) (domain.Test, error) {
	row := repo.DB.QueryRow("SELECT * FROM tests where test_code=?", testCode)
	fmt.Println(testCode)

	var data domain.Test

	err := row.Scan(&data.ID, &data.TestTitle, &data.TestCode, &data.Description,
		&data.CreatedBy, &data.Duration, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		return data, err
	}
	// data.CreatedAt = time.Now().Add(24 * time.Hour)
	// data.UpdatedAt = time.Now().Add(24 * time.Hour)

	if err := row.Err(); err != nil {
		return domain.Test{}, err
	}
	// fmt.Println(data)
	return data, err
}

func (repo *repoTest) GetByTestCodeWithQuestions(testCode string) (domain.Test, error) {
	var data domain.Test
	var questionsMap = make(map[string]domain.Question)
	var answerOptions []domain.AnswerOption

	query := `
	SELECT 
		t.id AS test_id,
		t.test_code AS test_code,
		t.test_title AS test_title,
		t.description AS description,
		t.duration AS duration,
		t.created_at AS created_at,
		q.id AS question_id,
		q.test_id AS question_test_id,
		q.content_question AS content_question,
		q.image_url AS image_url,
		q.audio_url AS audio_url,
		q.question_type AS question_type,
		q.points AS points,
		q.question_number AS question_number,
		q.created_at AS question_created_at,
		ao.id AS answer_option_id,
		ao.question_id AS answer_question_id,
		ao.content_answer AS content_answer_option
	FROM 
		tests t
	INNER JOIN 
		questions q
	ON 
		t.id = q.test_id
	LEFT JOIN 
		answer_options ao
	ON
		q.id = ao.question_id
	WHERE 
		t.test_code = ?
	ORDER BY q.question_number DESC;
	`

	rows, err := repo.DB.Query(query, testCode)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		var question domain.Question
		var answerOption domain.AnswerOption

		if err := rows.Scan(
			&data.ID,
			&data.TestCode,
			&data.TestTitle,
			&data.Description,
			&data.Duration,
			&data.CreatedAt,
			&question.ID,
			&question.TestID,
			&question.ContentQuestion,
			&question.ImageURL,
			&question.AudioURL,
			&question.QuestionType,
			&question.Points,
			&question.QuestionNumber,
			&question.CreatedAt,
			&answerOption.ID,
			&answerOption.QuestionID,
			&answerOption.ContentAnswer,
		); err != nil {
			return data, err
		}

		if _, exists := questionsMap[question.ID]; !exists {
			questionsMap[question.ID] = question
		}

		if answerOption.ID != "" {
			answerOptions = append(answerOptions, answerOption)
		}
	}

	fmt.Println(answerOptions)

	if err := rows.Err(); err != nil {
		return data, err
	}

	for _, answerOption := range answerOptions {
		question := questionsMap[answerOption.QuestionID]
		question.AnswerOptions = append(question.AnswerOptions, answerOption)
		questionsMap[answerOption.QuestionID] = question
	}

	for _, question := range questionsMap {
		data.Questions = append(data.Questions, question)
	}

	return data, nil
}

func (repo *repoTest) Create(test *domain.Test) error {
	newUUID, _ := uuid.NewRandom()
	// newUUID, _ := uuid.NewUUID()
	id := newUUID.String()

	_, err := repo.DB.Exec("INSERT INTO tests (id, test_code, test_title, description, created_by, duration) values (?, ?, ?, ?, ?, ?)",
		id, test.TestCode, test.TestTitle, test.Description, test.CreatedBy, test.Duration)
	return err
}

func (repo *repoTest) Update(id string, test *domain.Test) error {

	_, err := repo.DB.Exec("UPDATE tests SET test_code = ?, test_title = ?, description = ?, created_by = ?, duration = ? WHERE id = ?",
		test.TestCode, test.TestTitle, test.Description, test.CreatedBy, test.Duration, id)
	return err
}

func (repo *repoTest) Delete(id string) error {

	_, err := repo.DB.Exec("DELETE FROM tests WHERE id = ?", id)

	return err
}
