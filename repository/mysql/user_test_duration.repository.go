package repositoryMySql

import (
	"database/sql"
	"fmt"

	// "time"
	"gapai-skor-api/domain"
)

type repoUserTestDuration struct {
	DB *sql.DB
}

func CreateRepoUserTestDuration(db *sql.DB) domain.UserTestDurationRepository {
	return &repoUserTestDuration{DB: db}
}

func (repo *repoUserTestDuration) GetAll() ([]domain.UserTestDuration, error) {

	rows, err := repo.DB.Query("SELECT * FROM user_test_durations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []domain.UserTestDuration

	for rows.Next() {
		var test domain.UserTestDuration
		err := rows.Scan(&test.ID, &test.UserID, &test.TestID, &test.StartTime,
			&test.EndTime, &test.Duration, &test.Status, &test.CreatedAt, &test.UpdatedAt)
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

func (repo *repoUserTestDuration) GetByID(id string) (domain.UserTestDuration, error) {
	row := repo.DB.QueryRow("SELECT * FROM user_test_durations where id=?", id)
	fmt.Println(id)

	var data domain.UserTestDuration

	err := row.Scan(&data.ID, &data.UserID, &data.TestID, &data.StartTime,
		&data.EndTime, &data.Duration, &data.Status, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		return data, err
	}
	// data.CreatedAt = time.Now().Add(24 * time.Hour)
	// data.UpdatedAt = time.Now().Add(24 * time.Hour)

	if err := row.Err(); err != nil {
		return domain.UserTestDuration{}, err
	}
	// fmt.Println(data)
	return data, err
}

func (repo *repoUserTestDuration) Create(test *domain.UserTestDuration) error {
	_, err := repo.DB.Exec("INSERT INTO user_test_durations (user_id, test_id, start_time, end_time, duration, status) values (?, ?, ?, ?, ?, ?)",
		test.UserID, test.TestID, test.StartTime, test.EndTime, test.Duration, test.Status)
	return err
}
