package repositories

import (
	"database/sql"
	"fmt"
	"gapi/config"
	"gapi/pkg/models"
	"log"
	"sync"
)

type UserRepository struct {
	mu     sync.Mutex
	users  []models.User
	nextID int
	db     *sql.DB
}

func NewUserRepository() *UserRepository {
	conf, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		conf.Database.Username,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.DBName,
		conf.Database.SSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &UserRepository{
		users:  []models.User{},
		nextID: 1,
		db:     db,
	}
}

func (r *UserRepository) Create(user models.User) (models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.nextID
	r.nextID++
	r.users = append(r.users, user)

	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetByID(id int) (models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var user models.User

	query := `SELECT * FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var users []models.User

	query := `SELECT * FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return []models.User{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Update(id int, user models.User) (models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	updateQuery :=
		`UPDATE users 
			SET name = COALESCE(NULLIF($1, ''), name),
			email = COALESCE(NULLIF($2, ''), email)
			WHERE id = $3
			RETURNING name, email`

	var updatedUser models.User
	err := r.db.QueryRow(updateQuery, user.Name, user.Email, id).Scan(&updatedUser.Name, &updatedUser.Email)
	if err != nil {
		return models.User{}, err
	}

	return updatedUser, nil
}

func (r *UserRepository) setAdminStatus(id int, isAdmin bool) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	query := "UPDATE users SET is_admin = $1 WHERE id = $2"
	_, err := r.db.Exec(query, isAdmin, id)
	return err
}

func (r *UserRepository) AddAdmin(id int) error {
	return r.setAdminStatus(id, true)
}

func (r *UserRepository) DelAdmin(id int) error {
	return r.setAdminStatus(id, false)
}
