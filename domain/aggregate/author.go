package aggregate

import "github.com/huimingz/ddd-go/domain/entity"

type Author struct {
	user  *entity.User
	books []*entity.Book
}

func (c *Author) AddBook(book *entity.Book) {
	c.books = append(c.books, book)
}

func (c *Author) AllBooks() []*entity.Book {
	return c.books
}
