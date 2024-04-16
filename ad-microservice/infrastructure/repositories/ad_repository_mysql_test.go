package repositories_test

import (
	"ad-microservice/domain/models"
	"ad-microservice/infrastructure/repositories"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// func TestConnectDBError(t *testing.T) {
// 	// Configurar el mock de SQL
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("Error al crear mock de base de datos: %v", err)
// 	}
// 	defer db.Close()

// 	// Configurar el mock para manejar la consulta SELECT VERSION()
// 	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("5.7.31"))

// 	// Inicializar el GORM con el mock de base de datos
// 	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
// 	if err != nil {
// 		t.Fatalf("Error al abrir GORM: %v", err)
// 	}

// 	// Actuar
// 	imple := repositories.NewAdRepository(gormDB)
// 	err = imple.ConnectDB()

// 	// Verificar
// 	assert.NoError(t, err)
// }

func TestCreateAd(t *testing.T) {
	// Configurar el mock de SQL
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error al crear mock de base de datos: %v", err)
	}
	defer db.Close()

	// Configurar el mock para manejar la consulta SELECT VERSION()
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("5.7.31"))

	// Inicializar el GORM con el mock de base de datos
	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error al abrir GORM: %v", err)
	}

	// Crear instancia de MySQLAdRepository con la configuración falsa
	implObj := repositories.NewAdRepository(gormDB)

	// Crear un nuevo anuncio
	newAd := models.Add{
		ProductID: 1,
		Price:     10.5,
		Time:      60,
		Date:      time.Now().AddDate(0, 0, 1),
		UserID:    1,
		View:      100,
	}

	t.Run("Create ad success", func(t *testing.T) {
		// Configurar el mock para esperar la inserción
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `adds` (.+)").
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), newAd.ProductID, newAd.Price, newAd.Time, newAd.Date, newAd.UserID, newAd.View).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		// Act
		err = implObj.CreateAd(newAd)

		// Assert
		assert.NoError(t, err)

		// Asegurarse de que todas las expectativas se cumplan
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectativas no cumplidas: %s", err)
		}
	})

	t.Run("Create ad with error", func(t *testing.T) {
		// Configurar el mock para simular un error en la base de datos
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `adds` (.+)").
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), newAd.ProductID, newAd.Price, newAd.Time, newAd.Date, newAd.UserID, newAd.View).
			WillReturnError(errors.New("error de base de datos"))
		mock.ExpectRollback()

		// Act
		err = implObj.CreateAd(newAd)

		// Assert
		assert.Error(t, err)
		assert.EqualError(t, err, "error de base de datos")

		// Asegurarse de que todas las expectativas se cumplan
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectativas no cumplidas: %s", err)
		}
	})
}

func TestGetAllAd(t *testing.T) {
	// Configurar el mock de SQL
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error al crear mock de base de datos: %v", err)
	}
	defer db.Close()

	// Configurar el mock para manejar la consulta SELECT VERSION()
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("5.7.31"))

	// Inicializar el GORM con el mock de base de datos
	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error al abrir GORM: %v", err)
	}

	// Crear instancia de MySQLAdRepository con la configuración falsa
	implObj := repositories.NewAdRepository(gormDB)

	t.Run("Get all ads success", func(t *testing.T) {
		// Configurar el mock para esperar la consulta
		rows := sqlmock.NewRows([]string{"product_id", "price", "time", "date", "user_id", "view"}).
			AddRow(1, 10.5, 60, time.Now().AddDate(0, 0, 1), 1, 100).
			AddRow(2, 20.5, 30, time.Now().AddDate(0, 0, 1), 2, 200)
		mock.ExpectQuery("SELECT (.+) FROM `adds`").WillReturnRows(rows)

		// Act
		ads, err := implObj.GetAllAd()

		// Assert
		assert.NoError(t, err)
		assert.Len(t, *ads, 2)

		// Verificar que se hayan cumplido las expectativas del mock
		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("Get all ads with error", func(t *testing.T) {
		// Configurar el mock para simular un error en la consulta
		mock.ExpectQuery("SELECT (.+) FROM `adds`").WillReturnError(errors.New("error de base de datos"))

		// Act
		ads, err := implObj.GetAllAd()

		// Assert
		assert.Error(t, err)
		assert.Nil(t, ads)
		assert.EqualError(t, err, "error de base de datos")

		// Verificar que se hayan cumplido las expectativas del mock
		assert.Nil(t, mock.ExpectationsWereMet())
	})
}

func TestGetAddByIDProduct(t *testing.T) {
	// Configurar el mock de SQL
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error al crear mock de base de datos: %v", err)
	}
	defer db.Close()

	// Configurar el mock para manejar la consulta SELECT VERSION()
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("5.7.31"))

	// Inicializar el GORM con el mock de base de datos
	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error al abrir GORM: %v", err)
	}

	// Crear instancia de MySQLAdRepository con la configuración falsa
	implObj := repositories.NewAdRepository(gormDB)

	t.Run("Get ad by ID success", func(t *testing.T) {
		// Configurar el mock para esperar la consulta
		rows := sqlmock.NewRows([]string{"product_id", "price", "time", "date", "user_id", "view"}).
			AddRow(1, 10.5, 60, time.Now().AddDate(0, 0, 1), 1, 100)
		query := "SELECT * FROM `adds` WHERE product_id = ? AND `adds`.`deleted_at` IS NULL"
		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

		// Act
		ad, err := implObj.GetAddByIDProduct("1")

		// Assert
		assert.Nil(t, err)
		assert.NotNil(t, *ad)

		// Verificar que se hayan cumplido las expectativas del mock
		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("Get ad by ID with error", func(t *testing.T) {
		// Configurar el mock para simular un error en la consulta
		query := "SELECT * FROM `adds` WHERE product_id = ? AND `adds`.`deleted_at` IS NULL"
		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("error de base de datos"))

		// Act
		ad, err := implObj.GetAddByIDProduct("1")

		// Assert
		assert.Error(t, err)
		assert.Nil(t, ad)
		assert.EqualError(t, err, "error de base de datos")

		// Verificar que se hayan cumplido las expectativas del mock
		assert.Nil(t, mock.ExpectationsWereMet())
	})
}

func TestUpdateAddData(t *testing.T) {
	// Configurar el mock de SQL
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error al crear mock de base de datos: %v", err)
	}
	defer db.Close()

	// Configurar el mock para manejar la consulta SELECT VERSION()
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("5.7.31"))

	// Inicializar el GORM con el mock de base de datos
	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error al abrir GORM: %v", err)
	}

	// Crear instancia de MySQLAdRepository con la configuración falsa
	implObj := repositories.NewAdRepository(gormDB)

	t.Run("Update ad data success", func(t *testing.T) {
		// Crear un anuncio a actualizar
		adToUpdate := models.Add{
			ProductID: 1,
			Price:     15.5,
			Time:      90,
			Date:      time.Now().AddDate(0, 0, 2),
			UserID:    1,
			View:      150,
		}

		// Configurar el mock para esperar la actualización
		mock.ExpectBegin()
		query := "UPDATE `adds` SET `updated_at`=?,`product_id`=?,`price`=?,`time`=?,`date`=?,`user_id`=?,`view`=? WHERE product_id = ? AND `adds`.`deleted_at` IS NULL"
		mock.ExpectExec(regexp.QuoteMeta(query)).
			WithArgs(sqlmock.AnyArg(), adToUpdate.ProductID, adToUpdate.Price, adToUpdate.Time, adToUpdate.Date, adToUpdate.UserID, adToUpdate.View, adToUpdate.ProductID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		// Act
		err = implObj.UpdateAddData(adToUpdate)

		// Assert
		assert.Nil(t, err)

		// Verificar que se hayan cumplido las expectativas del mock
		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("Update ad data with error", func(t *testing.T) {
		// Crear un anuncio a actualizar
		adToUpdate := models.Add{
			ProductID: 1,
			Price:     15.5,
			Time:      90,
			Date:      time.Now().AddDate(0, 0, 2),
			UserID:    1,
			View:      150,
		}

		// Configurar el mock para simular un error en la actualización
		mock.ExpectBegin()
		query := "UPDATE `adds` SET `updated_at`=?,`product_id`=?,`price`=?,`time`=?,`date`=?,`user_id`=?,`view`=? WHERE product_id = ? AND `adds`.`deleted_at` IS NULL"
		mock.ExpectExec(regexp.QuoteMeta(query)).
			WithArgs(sqlmock.AnyArg(), adToUpdate.ProductID, adToUpdate.Price, adToUpdate.Time, adToUpdate.Date, adToUpdate.UserID, adToUpdate.View, adToUpdate.ProductID).
			WillReturnError(errors.New("error de base de datos"))
		mock.ExpectRollback()

		// Act
		err = implObj.UpdateAddData(adToUpdate)

		// Assert
		assert.Error(t, err)
		assert.EqualError(t, err, "error de base de datos")

		// Verificar que se hayan cumplido las expectativas del mock
		assert.Nil(t, mock.ExpectationsWereMet())
	})
}

func TestDeleteAddByProductID(t *testing.T) {
	// Configurar el mock de SQL
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error al crear mock de base de datos: %v", err)
	}
	defer db.Close()

	// Configurar el mock para manejar la consulta SELECT VERSION()
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("5.7.31"))

	// Inicializar el GORM con el mock de base de datos
	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error al abrir GORM: %v", err)
	}

	// Crear instancia de MySQLAdRepository con la configuración falsa
	implObj := repositories.NewAdRepository(gormDB)

	t.Run("Delete ad by product ID success", func(t *testing.T) {
		// Configurar el mock para esperar la eliminación
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE `adds` SET `deleted_at`=? WHERE product_id = ? AND `adds`.`deleted_at` IS NULL")).
			WithArgs(sqlmock.AnyArg(), "1").
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		// Act
		err = implObj.DeleteAddByProductID("1")

		// Assert
		assert.Nil(t, err)

		// Verificar que se hayan cumplido las expectativas del mock
		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("Delete ad by product ID with error", func(t *testing.T) {
		// Configurar el mock para simular un error en la eliminación
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE `adds` SET `deleted_at`=? WHERE product_id = ? AND `adds`.`deleted_at` IS NULL")).
			WithArgs(sqlmock.AnyArg(), "1").
			WillReturnError(errors.New("error de base de datos"))
		mock.ExpectRollback()

		// Act
		err = implObj.DeleteAddByProductID("1")

		// Assert
		assert.Error(t, err)
		assert.EqualError(t, err, "error de base de datos")

		// Verificar que se hayan cumplido las expectativas del mock
		assert.Nil(t, mock.ExpectationsWereMet())
	})
}
