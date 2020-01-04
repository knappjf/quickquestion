package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gocraft/dbr/v2"
	"github.com/gocraft/dbr/v2/dialect"
	"github.com/knappjf/quickquestion/internal/models"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateThing(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	defer db.Close()

	mock.ExpectExec(`INSERT INTO "things" \("name","description","enabled"\) VALUES \('foo','foo''s description',1\)`).
		WillReturnResult(sqlmock.NewResult(1, 1))

	conn := &dbr.Connection{
		DB:            db,
		Dialect:       dialect.SQLite3,
		EventReceiver: &dbr.NullEventReceiver{},
	}

	thing := models.Thing{
		Name:        "foo",
		Description: "foo's description",
		Enabled:     true,
	}
	repo := NewThingRepository(conn.NewSession(nil))
	assert.Nil(t, repo.CreateThing(thing))

	mock.ExpectClose()
	assert.Nil(t, conn.Close())
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestGetThing(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	defer db.Close()

	mock.ExpectQuery(`SELECT name, description, enabled FROM things WHERE \(thing_id = 123\)`).
		WillReturnRows(
			sqlmock.NewRows([]string{"name", "description", "enabled"}).
				AddRow("foo", "foo's description", true))

	conn := &dbr.Connection{
		DB:            db,
		Dialect:       dialect.SQLite3,
		EventReceiver: &dbr.NullEventReceiver{},
	}

	repo := NewThingRepository(conn.NewSession(nil))
	thing, err := repo.GetThing(123)

	assert.Nil(t, err)

	expected := models.Thing{
		Name:        "foo",
		Description: "foo's description",
		Enabled:     true,
	}
	assert.Equal(t, expected, thing)
	mock.ExpectClose()
	assert.Nil(t, conn.Close())
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDeleteThing(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	defer db.Close()

	mock.ExpectExec(`DELETE FROM "things" WHERE \(thing_id = 123\)`).
		WillReturnResult(sqlmock.NewResult(1, 1))

	conn := &dbr.Connection{
		DB:            db,
		Dialect:       dialect.SQLite3,
		EventReceiver: &dbr.NullEventReceiver{},
	}

	repo := NewThingRepository(conn.NewSession(nil))

	assert.Nil(t, repo.DeleteThing(123))
	mock.ExpectClose()
	assert.Nil(t, conn.Close())
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUpdateThing(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	defer db.Close()

	conn := &dbr.Connection{
		DB:            db,
		Dialect:       dialect.SQLite3,
		EventReceiver: &dbr.NullEventReceiver{},
	}

	//TODO: Figure out this regex
	mock.ExpectExec(`UPDATE "things" SET .* WHERE \(thing_id = 123\)`).
		WillReturnResult(sqlmock.NewResult(1, 1))

	thing := models.Thing{
		Name:        "foo",
		Description: "foo description",
		Enabled:     false,
	}

	repo := NewThingRepository(conn.NewSession(nil))

	assert.Nil(t, repo.UpdateThing(123, thing))

	mock.ExpectClose()
	assert.Nil(t, conn.Close())
	assert.Nil(t, mock.ExpectationsWereMet())
}
