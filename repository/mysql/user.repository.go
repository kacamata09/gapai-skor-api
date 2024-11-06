package repositoryMySql

import (
	"database/sql"
	"fmt"

	// "time"
	"gapai-skor-api/domain"
)

type repoUser struct{
    DB *sql.DB
}

func CreateRepoUser(db *sql.DB) domain.UserRepository {
    return &repoUser{DB: db}
}

func (repo *repoUser) GetAll() ([]domain.User, error) {

	rows, err := repo.DB.Query("SELECT * FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var data []domain.User
    
    for rows.Next() {
        var user domain.User
        err := rows.Scan(&user.ID, &user.Fullname, &user.Username, &user.Email, 
            &user.Password, &user.Last_login, &user.Role_id, &user.CreatedAt, &user.UpdatedAt)
        fmt.Println(err)
        if err != nil {
            return data, err
        }
        // user.CreatedAt = time.Now().Add(24 * time.Hour)
        // user.UpdatedAt = time.Now().Add(24 * time.Hour)
        data = append(data, user)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }
    return data, err
}

func (repo *repoUser) GetByID(id string) (domain.User, error) {
    row := repo.DB.QueryRow("SELECT * FROM users where id=?", id)
    fmt.Println(id)
    
    var data domain.User
    
    err := row.Scan(&data.ID, &data.Fullname, &data.Username, &data.Email, 
        &data.Password, &data.Last_login, &data.Role_id, &data.CreatedAt, &data.UpdatedAt)
    if err != nil {
        return data, err
    }
    // data.CreatedAt = time.Now().Add(24 * time.Hour)
    // data.UpdatedAt = time.Now().Add(24 * time.Hour)
    
    if err := row.Err(); err != nil {
        return domain.User{}, err
    }
    // fmt.Println(data)
    return data, err
}

func (repo *repoUser) Create(user *domain.User) error {
    _, err := repo.DB.Exec("INSERT INTO users (fullname, username, email, password, role_id) values (?, ?, ?, ?, ?)", 
    user.Fullname, user.Username, user.Email, user.Password, user.Role_id)
    return err
}