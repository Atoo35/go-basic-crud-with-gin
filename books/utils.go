package books

import (
	"github.com/Atoo35/basic-crud/models"
	"github.com/Atoo35/basic-crud/schema"
)

func FormatBookResponse(data models.Book, book *schema.CreateBook) {
	book.ID = data.ID
	book.Name = data.Name
	book.Author = data.Author
	book.Publication = data.Publication
}
