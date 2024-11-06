package repositoryMySql

import (
	"database/sql"
	"fmt"

	// "time"
	"gapai-skor-api/domain"
)

type repoTest struct{
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

func (repo *repoTest) Create(test *domain.Test) error {
    _, err := repo.DB.Exec("INSERT INTO tests (test_code, test_title, description, created_by, duration) values (?, ?, ?, ?, ?)", 
    test.TestCode, test.TestTitle, test.Description, test.CreatedBy, test.Duration )
    return err
}