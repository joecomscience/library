package postgres

import (
	"database/sql"
	"fmt"
	"github.com/joecomscience/library-api/pkg/domain"
)

type PostgresDBRepository struct {
	*sql.DB
}

func (repo PostgresDBRepository) FindAll() ([]domain.Book, error) {
	var result []domain.Book
	rows, err := repo.DB.Query("SELECT * FROM book")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book domain.Book
		if err := rows.Scan(
			&book.Id,
			&book.Name,
			&book.Author,
			&book.Image,
			&book.Description,
			&book.CreatedDate,
			&book.UpdatedDate,
		); err != nil {
			return nil, err
		}

		result = append(result, book)
	}

	return result, nil
}

func (repo PostgresDBRepository) FindById(id string) (domain.Book, error) {
	var book domain.Book
	row := repo.DB.QueryRow("SELECT * FROM book where id=$1", id)
	if err := row.Scan(
		&book.Id,
		&book.Name,
		&book.Author,
		&book.Image,
		&book.Description,
		&book.CreatedDate,
		&book.UpdatedDate,
	); err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

func (repo PostgresDBRepository) Create(book domain.Book) error {
	query := `
INSERT INTO book(name, author, image, description) 
VALUES($1, $2, $3, $4)
`
	row, err := repo.DB.Exec(
		query,
		book.Name,
		book.Author,
		book.Image,
		book.Description,
	)
	if err != nil {
		return err
	}
	count, err := row.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println(count)
	return nil
}

func (repo PostgresDBRepository) Update(book domain.Book) error {
	query := `
UPDATE SET name = $1, author = $2, image = $3, description = $4 
WHERE id = $5
`
	row, err := repo.DB.Exec(
		query,
		book.Name,
		book.Author,
		book.Image,
		book.Description,
		book.Id,
	)
	if err != nil {
		return err
	}
	count, err := row.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println(count)
	return nil
}

func (repo PostgresDBRepository) Delete(id string) error {
	query := `DELETE FROM book WHERE id = $1`
	row, err := repo.DB.Exec(query, id)
	if err != nil {
		return err
	}
	count, err := row.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println(count)
	return nil
}
