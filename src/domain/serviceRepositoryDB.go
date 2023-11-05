package domain

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type ServiceRepositoryDB struct {
	client *sql.DB
}

func (s ServiceRepositoryDB) FindAll() ([]Service, error) {
	log.Output(1, fmt.Sprintf("FindAll() called"))

	sqlQuery := "SELECT service_name, service_id, security_level, service_type, name FROM services"

	log.Output(1, fmt.Sprintf("Starting Query: %p", s.client.Query))
	rows, err := s.client.Query(sqlQuery)
	log.Output(1, fmt.Sprintf("Query complete: %p", s.client.Query))
	if err != nil {
		log.Output(1, fmt.Sprintf("Query failed: %p", s.client.Query))
		return nil, err
	}
	services := make([]Service, 0)
	for rows.Next() {
		var service Service
		err = rows.Scan(&service.ServiceName, &service.ServiceID, &service.SecurityLevel, &service.ServiceType, &service.Name)
		if err != nil {
			log.Output(1, fmt.Sprintf("Scan failed: %p", s.client.Query))
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}

func NewServiceRepositoryDB() ServiceRepositoryDB {
	client, err := sql.Open("mysql", "root:NigelTest@tcp(localhost:3306)/thsdb")
	if err != nil {
		log.Fatal(1, fmt.Sprintf("Failed to connect to the database. Error: %p", err))
		panic(err)
	}
	log.Output(1, fmt.Sprintf("client: %p", client))
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return ServiceRepositoryDB{client: client}
}
