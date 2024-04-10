package services

// import (
// 	"auth-microservice/models"
// 	"context"
// 	"database/sql"
// 	"time"
// )

// type Repository interface {
// 	Close()
// 	FindByID(id string) (*models.User, error)
// 	Find() ([]*models.User, error)
// 	Create(user *models.User) error
// 	Update(user *models.User) error
// 	Delete(id string) error
// }

// // repository represent the repository model
// type repository struct {
// 	db *sql.DB
// }

// // NewRepository will create a variable that represent the Repository struct
// func NewRepository(dialect, dsn string, idleConn, maxConn int) (Repository, error) {
// 	db, err := sql.Open(dialect, dsn)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		return nil, err
// 	}

// 	db.SetMaxIdleConns(idleConn)
// 	db.SetMaxOpenConns(maxConn)

// 	return &repository{db}, nil
// }

// // Close attaches the provider and close the connection
// func (r *repository) Close() {
// 	r.db.Close()
// }

// // FindByID attaches the user repository and find data based on id
// func (r *repository) FindByID(id string) (*models.User, error) {
// 	user := new(models.User)

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	err := r.db.QueryRowContext(ctx, "SELECT id_auth, name, email, image FROM users WHERE id_auth = ?", id).Scan(&user.IdAuth, &user.Name, &user.Email, &user.Image)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// // Find attaches the user repository and find all data
// func (r *repository) Find() ([]*models.User, error) {
// 	users := make([]*models.User, 0)

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	rows, err := r.db.QueryContext(ctx, "SELECT id_auth, name, email, image FROM users")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		user := new(models.User)
// 		err = rows.Scan(
// 			&user.IdAuth,
// 			&user.Name,
// 			&user.Email,
// 			&user.Image,
// 		)

// 		if err != nil {
// 			return nil, err
// 		}
// 		users = append(users, user)
// 	}

// 	return users, nil
// }

// // Create attaches the user repository and creating the data
// func (r *repository) Create(user *models.User) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	query := "INSERT INTO users (id, name, email, image) VALUES (?, ?, ?, ?)"
// 	stmt, err := r.db.PrepareContext(ctx, query)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.ExecContext(ctx, user.IdAuth, user.Name, user.Email, user.Image)
// 	return err
// }

// // Update attaches the user repository and update data based on id
// func (r *repository) Update(user *models.User) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	query := "UPDATE users SET name = ?, image = ?, email = ? WHERE id_auth = ?"
// 	stmt, err := r.db.PrepareContext(ctx, query)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.ExecContext(ctx, user.Name, user.Email, user.Image, user.IdAuth)
// 	return err
// }

// // Delete attaches the user repository and delete data based on id
// func (r *repository) Delete(id string) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	query := "DELETE FROM users WHERE id_auth= ?"
// 	stmt, err := r.db.PrepareContext(ctx, query)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.ExecContext(ctx, id)
// 	return err
// }
