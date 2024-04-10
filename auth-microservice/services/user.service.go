package services

import (
	"auth-microservice/models"
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"
)

type Repository interface {
	Close()
	GetAllUsers() ([]models.User, int)
	GetUser(id string) (*models.User, int)
	CreateUser(user *models.User) (*models.User, int)
	UpdateUser(id string, updatedUser *models.User) (*models.User, int)
	ChangeUserRole(id string, newRole models.Role) (*models.User, int)
	GetUserByIdAuth(idAuth string) (*models.User, int)
}

// repository represent the repository model
type repository struct {
	db *sql.DB
}

// NewRepository will create a variable that represent the Repository struct/
func NewRepository(dialect, dsn string, idleConn, maxConn int) (Repository, error) {
	db, err := sql.Open(dialect, dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(idleConn)
	db.SetMaxOpenConns(maxConn)

	return &repository{db}, nil
}

func (r *repository) Close() {
	r.db.Close()
}
func (r *repository) GetAllUsers() ([]models.User, int) {
	var users []models.User

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, "SELECT id_auth, name, email, image FROM users")
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.IdAuth, &user.Name, &user.Email, &user.Image)
		if err != nil {
			return nil, http.StatusInternalServerError
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, http.StatusInternalServerError
	}

	return users, http.StatusOK
}

func (r *repository) GetUser(id string) (*models.User, int) {
	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := r.db.QueryRowContext(ctx, "SELECT * FROM users WHERE id_auth = ?", id)
	err := row.Scan(&user.IdAuth, &user.Name, &user.Email, &user.Image)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, http.StatusNotFound
		}
		return nil, http.StatusInternalServerError
	}
	return &user, http.StatusOK
}

func (r *repository) CreateUser(user *models.User) (*models.User, int) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "INSERT INTO users (id_auth, name, email, password, image, ubication, role) VALUES (?, ?, ?, ?, ?, ?, ?)"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, user.IdAuth, user.Name, user.Email, user.Password, user.Image, user.Ubication, user.Role)
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	return user, http.StatusCreated
}

func (r *repository) UpdateUser(id string, updatedUser *models.User) (*models.User, int) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "UPDATE users SET name = ?, email = ?, password = ?, image = ?, ubication = ?, role = ? WHERE id_auth = ?"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, updatedUser.Name, updatedUser.Email, updatedUser.Password, updatedUser.Image, updatedUser.Ubication, updatedUser.Role, id)
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	return updatedUser, http.StatusOK
}

func (r *repository) ChangeUserRole(id string, newRole models.Role) (*models.User, int) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, "UPDATE users SET role = ? WHERE id_auth = ?", string(newRole), id)
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	return &models.User{IdAuth: id, Role: newRole}, http.StatusOK
}

func (r *repository) GetUserByIdAuth(idAuth string) (*models.User, int) {
	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := r.db.QueryRowContext(ctx, "SELECT id_auth, name, email, image FROM users WHERE id_auth = ?", idAuth)
	err := row.Scan(&user.IdAuth, &user.Name, &user.Email, &user.Image)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, http.StatusNotFound
		}
		return nil, http.StatusInternalServerError
	}
	return &user, http.StatusOK
}

// func GetAllUsers() ([]models.User, int) {
// 	var users []models.User
// 	if err := db.DB.Find(&users).Error; err != nil {
// 		return nil, http.StatusInternalServerError
// 	}
// 	return users, http.StatusOK // Cambiado 200 a http.StatusOK
// }

// func CreateUser(user *models.User) (*models.User, int) {

// 	if err := db.DB.Create(user).Error; err != nil {
// 		return nil, http.StatusInternalServerError
// 	}
// 	return user, http.StatusOK
// }

// func UpdateUser(id string, updatedUser *models.User) (*models.User, int) {

// 	existingUser, err := GetUserByIdAuth(id)
// 	if err != http.StatusOK {
// 		return nil, http.StatusInternalServerError
// 	}

// 	existingUser.Name = func() string {
// 		if updatedUser.Name != "" {
// 			return updatedUser.Name
// 		}
// 		return existingUser.Name
// 	}()

// 	existingUser.Email = func() string {
// 		if updatedUser.Email != "" {
// 			return updatedUser.Email
// 		}
// 		return existingUser.Email
// 	}()

// 	existingUser.Password = func() string {
// 		if updatedUser.Password != "" {
// 			return updatedUser.Password
// 		}
// 		return existingUser.Password
// 	}()

// 	existingUser.Image = func() string {
// 		if updatedUser.Image != "" {
// 			return updatedUser.Image
// 		}
// 		return existingUser.Image
// 	}()

// 	existingUser.Ubication = func() string {
// 		if updatedUser.Ubication != "" {
// 			return updatedUser.Ubication
// 		}
// 		return existingUser.Ubication
// 	}()

// 	if err := db.DB.Save(&existingUser).Error; err != nil {
// 		return nil, http.StatusInternalServerError
// 	}

// 	return existingUser, http.StatusOK
// }

// func ChangeUserRole(id string, newRole models.Role) (*models.User, int) {
// 	user, err := GetUserByIdAuth(id)
// 	if err != http.StatusOK {
// 		return nil, http.StatusInternalServerError
// 	}

// 	switch newRole {
// 	case models.RoleAdmin, models.RoleUser, models.RoleEnterprise:
// 	default:
// 		return nil, http.StatusBadRequest
// 	}
// 	user.Role = newRole
// 	if err := db.DB.Save(&user).Error; err != nil {
// 		return nil, http.StatusInternalServerError
// 	}

// 	return user, http.StatusOK
// }

// func GetUserByIdAuth(idAuth string) (*models.User, int) {
// 	var user models.User
// 	if err := db.DB.Where("id_auth = ?", idAuth).First(&user).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			// No se encontró ningún usuario con el ID de autenticación proporcionado, devolver error 404
// 			return nil, http.StatusNotFound
// 		}
// 		// Ocurrió un error diferente, devolver error 500
// 		return nil, http.StatusInternalServerError
// 	}
// 	// Usuario encontrado, devolver el usuario y el código de estado 200
// 	return &user, http.StatusOK
// }

// type ServicesMocked struct {
// 	userMap map[string]*models.User // Mapa para simular la base de datos en memoria
// }

// func NewServicesMocked() *ServicesMocked {
// 	return &ServicesMocked{
// 		userMap: make(map[string]*models.User),
// 	}
// }

// func (rm *ServicesMocked) GetAllUsers() ([]models.User, int) {
// 	// Crear un slice de models.User con diferentes usuarios
// 	users := []models.User{
// 		{IdAuth: "1", Name: "User1"},
// 		{IdAuth: "2", Name: "User2"},
// 	}
// 	if len(users) == 0 {
// 		return nil, http.StatusNoContent
// 	}

// 	// Devolver el slice de users y el status code 200
// 	return users, http.StatusOK
// }

// func (rm *ServicesMocked) CreateUser(user *models.User) (*models.User, int) {
// 	// Simular la creación de un usuario sin acceder a la base de datos
// 	return user, http.StatusOK
// }

// func (rm *ServicesMocked) UpdateUser(id string, updatedUser *models.User) (*models.User, int) {
// 	// Verificar si el usuario existe en el mapa
// 	existingUser, ok := rm.userMap[id]
// 	if !ok {
// 		return nil, http.StatusNotFound
// 	}

// 	// Actualizar los campos del usuario existente
// 	if updatedUser.Name != "" {
// 		existingUser.Name = updatedUser.Name
// 	}
// 	if updatedUser.Email != "" {
// 		existingUser.Email = updatedUser.Email
// 	}
// 	if updatedUser.Password != "" {
// 		existingUser.Password = updatedUser.Password
// 	}
// 	if updatedUser.Image != "" {
// 		existingUser.Image = updatedUser.Image
// 	}
// 	if updatedUser.Ubication != "" {
// 		existingUser.Ubication = updatedUser.Ubication
// 	}

// 	// Actualizar el usuario en el mapa
// 	rm.userMap[id] = existingUser

// 	return existingUser, http.StatusOK
// }

// func (rm *ServicesMocked) GetUserByIdAuth(idAuth string) (*models.User, int) {
// 	// Simular la obtención de un usuario por ID de autenticación sin acceder a la base de datos
// 	user, ok := rm.userMap[idAuth]
// 	if !ok {
// 		return nil, http.StatusNotFound
// 	}
// 	return user, http.StatusOK
// }

// func (rm *ServicesMocked) ChangeUserRole(id string, newRole models.Role) (*models.User, int) {
// 	// Simular el cambio de rol de un usuario sin acceder a la base de datos
// 	user := &models.User{IdAuth: id, Role: newRole}
// 	return user, http.StatusOK
// }

// func (rm *ServicesMocked) GetUser(id string) (*models.User, int) {
// 	// Simular la obtención de un usuario sin acceder a la base de datos
// 	user := &models.User{IdAuth: "123", Name: "Test User"}
// 	if id == user.IdAuth {
// 		return user, http.StatusOK

// 	}
// 	return user, http.StatusNotFound

// }
