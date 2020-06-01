package domain

type Book struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Image       string `json:"image"`
	Description string `json:"description"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
}

type BookRepository interface {
	FindAll() ([]Book, error)
	FindById(string) (Book, error)
	Create(Book) error
	Update(Book) error
	Delete(string) error
}
