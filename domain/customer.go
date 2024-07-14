package domain

import (
	"database/sql"
	"log"
	"modulo/errors"
	"modulo/logger"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	Id   string `json:"id" xml:"id"`
	Name string `json:"full_name" xml:"full_name"`
	City string `json:"city" xml:"city"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(id string) (*Customer, *errors.AppErrors)
	DeleteById(id string) *errors.AppErrors
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	log.Printf("querying all customers")
	findAllSql := "select id, name, city from customers"
	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("error while querying customer table: " + err.Error())
		return nil, err
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City)
		if err != nil {
			log.Println("error while scanning customers: " + err.Error())
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errors.AppErrors) {
	// logger.Info("searching for customer with id: %s", id)
	findByIdSql := "select id, name, city from customers where id = ?"
	row := d.client.QueryRow(findByIdSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NotFoundError("customer not found")
		} else {
			log.Println("error while scanning customer: " + err.Error())
			return nil, errors.InternalServerError("unexpected database error")
		}
	}
	return &c, nil
}

func (d CustomerRepositoryDb) DeleteById(id string) *errors.AppErrors {
	deleteByIdSql := "delete from customers where id = ?"
	result, err := d.client.Exec(deleteByIdSql, id)
	if err != nil {
		log.Println("error while deleting customer: " + err.Error())
		return errors.InternalServerError("unexpected database error")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("error while getting rows affected: " + err.Error())
		return errors.InternalServerError("unexpected database error")
	}

	if rowsAffected == 0 {
		return errors.NotFoundError("customer not found")
	}

	return nil
}

type CustomerRepositoryDb struct {
	client *sql.DB
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	db, err := sql.Open("mysql", "root:my-secret@/estantevirtual")
	if err != nil {
		logger.Info(err.Error())
		panic(err)
	}

	// connection settings
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return CustomerRepositoryDb{db}
}
