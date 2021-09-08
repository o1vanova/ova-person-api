package repo

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"

	models "github.com/ozonva/ova-person-api/internal/models"
)

// PersonRepo - интерфейс хранилища для сущности Entity
type PersonRepo interface {
	AddPerson(context context.Context, person models.Person) (uint64, error)
	ListPersons(context context.Context, limit, offset uint64) ([]models.Person, error)
	DescribePerson(context context.Context, personId uint64) (*models.Person, error)
	RemovePerson(context context.Context, personId uint64) (uint64, error)
}

type repository struct {
	db sqlx.DB
}

func (repo repository) AddPerson(context context.Context, person models.Person) (uint64, error) {
	query := `INSERT INTO persons (id, userId, firstName, lastName, middleName, createdAt, updatedAt) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	result, err := repo.db.ExecContext(context, query, person.Id, person.UserId, person.FirstName, person.LastName, person.MiddleName, person.CreatedAt, person.CreatedAt)
	if err != nil {
		log.Fatalf("Insert to DataBase is failed: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Last inserted id: %v", err)
	}
	return uint64(id), err
}

func (repo repository) ListPersons(context context.Context, limit, offset uint64) ([]models.Person, error) {
	query := `SELECT * FROM persons LIMIT $1 OFFSET $2`
	rows, err := repo.db.QueryContext(context, query, limit, offset)
	if err != nil {
		log.Fatalf("Get persons from DataBase is failed: %v", err)
	}
	defer rows.Close()

	persons := make([]models.Person, 0, limit)
	for rows.Next() {
		person := models.Person{}
		if err := rows.Scan(&person.Id,
			&person.UserId,
			&person.FirstName,
			&person.LastName,
			&person.MiddleName,
			&person.CreatedAt,
			&person.UpdatedAt,
			&person.DeletedAt); err != nil {
			log.Fatalf("Get persons request has scan error: %v", err)
			return nil, err
		}
		persons = append(persons, person)
	}
	return persons, nil
}

func (repo repository) DescribePerson(context context.Context, personId uint64) (*models.Person, error) {
	query := `SELECT * FROM persons WHERE id = $1`
	rows, err := repo.db.QueryContext(context, query, personId)
	if err != nil {
		log.Fatalf("Describe person %v from DataBase is failed: %v", personId, err)
	}
	defer rows.Close()

	person := models.Person{}
	if err := rows.Scan(&person.Id,
		&person.UserId,
		&person.FirstName,
		&person.LastName,
		&person.MiddleName,
		&person.CreatedAt,
		&person.UpdatedAt,
		&person.DeletedAt); err != nil {
		log.Fatalf("Describe person request has scan error: %v", err)
		return nil, err
	}
	return &person, nil
}

func (repo repository) RemovePerson(context context.Context, personId uint64) (uint64, error) {
	query := `DELETE FROM persons WHERE id = $1`
	result, err := repo.db.ExecContext(context, query, personId)
	if err != nil {
		log.Fatalf("Remove person %v from DataBase is failed: %v", personId, err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("The number of rows affected: %v", err)
	}
	return uint64(id), err
}

func NewContext(db sqlx.DB) PersonRepo {
	return &repository{db: db}
}
