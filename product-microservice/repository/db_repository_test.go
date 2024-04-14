package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const FirstTestName = "Test Name 1"
const SecondTestName = "Test Name 2"
const SearchName = "name = ?"

type YourModel struct {
    ID            uint
    Name          string
    RelatedModels []RelatedModel
}

type RelatedModel struct {
    ID         uint
    YourModelID uint
    Detail     string
}

func TestAutoMigrate(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()
    dialector := mysql.New(mysql.Config{
        Conn:                      db,
        SkipInitializeWithVersion: true,
    })
    gormDB, err := gorm.Open(dialector, &gorm.Config{})
    require.NoError(t, err)
    mock.ExpectExec("CREATE TABLE").WillReturnResult(sqlmock.NewResult(0, 0))
    dbClient := NewGormDBClient(gormDB)
    err = dbClient.AutoMigrate(&YourModel{})
    assert.NoError(t, err)
    require.NoError(t, mock.ExpectationsWereMet())
}

func TestFirst(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()
    dialector := mysql.New(mysql.Config{
        Conn:                      db,
        SkipInitializeWithVersion: true,
    })
    gormDB, err := gorm.Open(dialector, &gorm.Config{})
    require.NoError(t, err)
    rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "test")
    mock.ExpectQuery("SELECT").WillReturnRows(rows)
    dbClient := NewGormDBClient(gormDB)
    var yourModel YourModel
    err = dbClient.First(&yourModel)
    assert.NoError(t, err)
    require.NoError(t, mock.ExpectationsWereMet())
}

func TestSaveUpdate(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()
    dialector := mysql.New(mysql.Config{
        Conn: db,
        SkipInitializeWithVersion: true,
    })
    gormDB, err := gorm.Open(dialector, &gorm.Config{})
    require.NoError(t, err)
    mock.ExpectBegin()
    mock.ExpectExec("INSERT INTO `your_models`").WillReturnResult(sqlmock.NewResult(1, 1))
    mock.ExpectCommit()
    mock.ExpectBegin()
    mock.ExpectExec("UPDATE `your_models`").WillReturnResult(sqlmock.NewResult(0, 1))
    mock.ExpectCommit()
    dbClient := NewGormDBClient(gormDB)
    err = dbClient.Save(&YourModel{Name: "New Name"})
    require.NoError(t, err)
    err = dbClient.Save(&YourModel{ID: 1, Name: "Updated Name"})
    require.NoError(t, err)
    require.NoError(t, mock.ExpectationsWereMet())
}

func TestCreate(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()
    dialector := mysql.New(mysql.Config{
        Conn: db,
        SkipInitializeWithVersion: true,
    })
    gormDB, err := gorm.Open(dialector, &gorm.Config{})
    require.NoError(t, err)
    mock.ExpectBegin()
    mock.ExpectExec("INSERT INTO `your_models`").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
    mock.ExpectCommit()
    dbClient := NewGormDBClient(gormDB)
    err = dbClient.Create(&YourModel{ID: uint(1), Name: "Test Name"})
    require.NoError(t, err)
    require.NoError(t, mock.ExpectationsWereMet())
}


func TestFindPreloaded(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()

    dialector := mysql.New(mysql.Config{
        Conn: db,
        SkipInitializeWithVersion: true,
    })
    gormDB, err := gorm.Open(dialector, &gorm.Config{})
    require.NoError(t, err)
    mainRows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Test Name")
    mock.ExpectQuery("^SELECT \\* FROM `your_models`").WillReturnRows(mainRows)
    relatedRows := sqlmock.NewRows([]string{"id", "your_model_id", "detail"}).AddRow(1, 1, "Related Data")
    mock.ExpectQuery("^SELECT \\* FROM `related_models`").WillReturnRows(relatedRows)
    dbClient := NewGormDBClient(gormDB)
    var yourModels []YourModel
    err = dbClient.FindPreloaded("RelatedModels", &yourModels)
    require.NoError(t, err)
    require.NoError(t, mock.ExpectationsWereMet())
    require.Len(t, yourModels, 1, "should retrieve one YourModel")
    require.Len(t, yourModels[0].RelatedModels, 1, "YourModel should have one related model")
}

func TestFind(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()
    dialector := mysql.New(mysql.Config{
        Conn: db,
        SkipInitializeWithVersion: true,
    })
    gormDB, err := gorm.Open(dialector, &gorm.Config{})
    require.NoError(t, err)
    rows := sqlmock.NewRows([]string{"id", "name"}).
        AddRow(1, FirstTestName).
        AddRow(2, SecondTestName)
    mock.ExpectQuery("^SELECT \\* FROM `your_models`").WithArgs(FirstTestName).WillReturnRows(rows)
    dbClient := NewGormDBClient(gormDB)
    var yourModels []YourModel
    err = dbClient.Find(&yourModels, SearchName, FirstTestName)
    require.NoError(t, err)
    require.NoError(t, mock.ExpectationsWereMet())
    require.Len(t, yourModels, 2, "should retrieve matching records")
    require.Equal(t, FirstTestName, yourModels[0].Name, "the name of the first model should match")
}

func TestFindWithCondition(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()
    dialector := mysql.New(mysql.Config{
        Conn: db,
        SkipInitializeWithVersion: true,
    })
    gormDB, err := gorm.Open(dialector, &gorm.Config{})
    require.NoError(t, err)
    rows := sqlmock.NewRows([]string{"id", "name"}).
        AddRow(1, FirstTestName).
        AddRow(2, SecondTestName)
    mock.ExpectQuery("^SELECT \\* FROM `your_models` WHERE name =").
        WithArgs(FirstTestName).
        WillReturnRows(rows)
    dbClient := NewGormDBClient(gormDB)
    var yourModels []YourModel
    err = dbClient.FindWithCondition(&yourModels, SearchName, FirstTestName)
    require.NoError(t, err)
    require.NoError(t, mock.ExpectationsWereMet())
    require.Len(t, yourModels, 2, "debería recuperar el número esperado de registros")
    require.Equal(t, FirstTestName, yourModels[0].Name, "el nombre del primer registro debería coincidir con la condición")
}

func TestDelete(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()
    dialector := mysql.New(mysql.Config{
        Conn: db,
        SkipInitializeWithVersion: true,
    })
    gormDB, err := gorm.Open(dialector, &gorm.Config{})
    require.NoError(t, err)
    mock.ExpectBegin()
    mock.ExpectExec("DELETE FROM `your_models` WHERE `your_models`.`id` = ?").
        WithArgs(1).
        WillReturnResult(sqlmock.NewResult(0, 1))
    mock.ExpectCommit()
    dbClient := NewGormDBClient(gormDB)
    err = dbClient.Delete(&YourModel{ID: 1})
    require.NoError(t, err)
    require.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteWithCondition(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()

    dialector := mysql.New(mysql.Config{
        Conn: db,
        SkipInitializeWithVersion: true,
    })
    gormDB, err := gorm.Open(dialector, &gorm.Config{})
    require.NoError(t, err)
    mock.ExpectBegin()
    mock.ExpectExec("^DELETE FROM `your_models` WHERE name =").
        WithArgs(FirstTestName).
        WillReturnResult(sqlmock.NewResult(0, 1))
    mock.ExpectCommit()
    dbClient := NewGormDBClient(gormDB)
    err = dbClient.DeleteWithCondition(&YourModel{}, SearchName, FirstTestName)
    require.NoError(t, err)
    require.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteByID(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()
    dialector := mysql.New(mysql.Config{
        Conn: db,
        SkipInitializeWithVersion: true,
    })
    gormDB, err := gorm.Open(dialector, &gorm.Config{})
    require.NoError(t, err)
    mock.ExpectBegin()
    mock.ExpectExec("DELETE FROM `your_models` WHERE `your_models`.`id` = ?").
        WithArgs(1).
        WillReturnResult(sqlmock.NewResult(0, 1))
    mock.ExpectCommit()
    dbClient := NewGormDBClient(gormDB)
    err = dbClient.DeleteByID(&YourModel{}, 1)
    require.NoError(t, err)
    require.NoError(t, mock.ExpectationsWereMet())
}