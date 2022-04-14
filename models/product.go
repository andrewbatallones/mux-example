package models

import (
	"database/sql"
	"errors"
)

type product struct {
	ID    int     `json:"id`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *product) getProducts(db *sql.DB) error {
	return errors.New("Not Implmemented")
}

func (p *product) updateProduct(db *sql.DB) error {
	return errors.New("Not Implmemented")
}

func (p *product) deleteProduct(db *sql.DB) error {
	return errors.New("Not Implmemented")
}

func (p *product) createProduct(db *sql.DB) error {
	return errors.New("Not Implmemented")
}

func getProducts(db *sql.DB, start, count int) ([]product, error) {
	return nil, errors.New("Not Implmemented")
}
