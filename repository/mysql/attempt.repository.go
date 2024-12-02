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
