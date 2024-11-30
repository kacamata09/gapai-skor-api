package repositoryMySql

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

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

	rows, err := repo.DB.Query(`
		SELECT 
			q.id AS question_id,
			q.test_id AS test_id,
			q.content_question AS content_question,
			q.image_url AS image_url,
			q.audio_url AS audio_url,
			q.question_type AS question_type,
			q.points AS points,
			q.question_number AS question_number,
			q.created_at AS question_created_at,
			ao.id AS answer_option_id,
			ao.question_id AS question_ao_id,
			ao.content_answer AS content_answer_option 
		FROM questions q 
		LEFT JOIN answer_options ao 
		ON ao.question_id = q.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	questionsMap := make(map[string]domain.Question)

	for rows.Next() {
		var question domain.Question
		var answerOption domain.AnswerOption

		err := rows.Scan(
			&question.ID, &question.TestID, &question.ContentQuestion, &question.ImageURL,
			&question.AudioURL, &question.QuestionType, &question.Points, &question.QuestionNumber,
			&question.CreatedAt, &answerOption.ID, &answerOption.QuestionID, &answerOption.ContentAnswer,
		)
		if err != nil {
			return nil, err
		}

		q, exists := questionsMap[question.ID]
		if !exists {
			question.AnswerOptions = []domain.AnswerOption{}
			questionsMap[question.ID] = question
		} else {
			question = q
		}

		if answerOption.ID != "" {
			question.AnswerOptions = append(question.AnswerOptions, answerOption)
			questionsMap[question.ID] = question
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	var data []domain.Question
	for _, question := range questionsMap {
		data = append(data, question)
	}

	return data, nil
}

func (repo *repoQuestion) GetByID(id string) (domain.Question, error) {

	// newUUID, _ := uuid.NewRandom()
	// // newUUID, _ := uuid.NewUUID()
	// id = newUUID.String()
	row := repo.DB.QueryRow("SELECT * FROM questions where id=?", id)
	fmt.Println(id)

	var data domain.Question

	err := row.Scan(&data.ID, &data.TestID, &data.ContentQuestion, &data.ImageURL,
		&data.AudioURL, &data.QuestionType, &data.Points, &data.QuestionNumber, &data.CreatedAt, &data.UpdatedAt)
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

func (repo *repoQuestion) Create(tx *sql.Tx, question *domain.Question) (id string, err error) {
	newUUID, _ := uuid.NewRandom()
	// newUUID, _ := uuid.NewUUID()
	id = newUUID.String()

	query := "INSERT INTO questions (id, test_id, content_question, image_url, audio_url, question_type, question_number, points) values (?, ?, ?, ?, ?, ?, ?, ?)"
	if tx != nil {
		_, err = tx.Exec(query, id, question.TestID, question.ContentQuestion, question.ImageURL,
			question.AudioURL, question.QuestionType, question.QuestionNumber, question.Points)
	} else {
		_, err = repo.DB.Exec(query, id, question.TestID, question.ContentQuestion, question.ImageURL,
			question.AudioURL, question.QuestionType, question.QuestionNumber, question.Points)
	}
	if err != nil {
		return "", err
	}
	return id, err
}
