package handler

// import (
// 	book "book-review/gunk/v1/book"

// 	"database/sql"
// 	"log"
// 	"net/http"
// 	"time"

// 	validation "github.com/go-ozzo/ozzo-validation/v4"
// )

// // "database/sql"
// // "log"
// // "net/http"
// // "time"

// // validation "github.com/go-ozzo/ozzo-validation/v4"
// type Book struct {
// 	ID           int          `json:"id" form:"-" db:"id"`
// 	BookTypeId   int          `json:"booktype_id" db:"booktype_id"`
// 	BookTypeName string       `json:"type_name" db:"type_name"`
// 	Title        string       `json:"book_title" db:"book_title"`
// 	Description  string       `json:"book_description" db:"book_description"`
// 	Author       string       `json:"author" db:"author"`
// 	Published    time.Time    `json:"published" db:"published"`
// 	CreatedAt    time.Time    `json:"created_at" db:"created_at"`
// 	UpdatedAt    time.Time    `json:"updated_at" db:"updated_at"`
// 	DeletedAt    sql.NullTime `json:"deleted_at" db:"deleted_at"`
// }

// // book all validation

// func (b Book) Validate() error {
// 	return validation.ValidateStruct(&b,
// 		validation.Field(&b.Title,
// 			validation.Required.Error("The Book Title field is required."),
// 		),
// 		validation.Field(&b.Description,
// 			validation.Required.Error("The Description field is required."),
// 		),
// 		validation.Field(&b.Author,
// 			validation.Required.Error("The username field is required."),
// 		),
// 		validation.Field(&b.Published,
// 			validation.Required.Error("The Published field is required."),
// 		),
// 	)
// }

// type BKFilter struct {
// 	Books      []Book
// 	SearchTerm string
// }

// func (h Handler) BookList(w http.ResponseWriter, r *http.Request) {
// 	if err := r.ParseForm(); err != nil {
// 		log.Fatalf("%#v", err)
// 	}
// 	st := r.FormValue("SearchTerm")
// 	listBK, err := h.usermgmSvc.BookList(r.Context(), &book.BookListRequest{
// 		SearchTerm: st,
// 	})
// 	if err != nil {
// 		http.Error(w, "Internal Server error", http.StatusInternalServerError)
// 	}

// 	data := []Book{}
// 	if listBK != nil {
// 		for _, v := range listBK.GetBookFilterList().TotalBook {

// 			data = append(data, Book{
// 				ID:           int(v.ID),
// 				BookTypeId:   int(v.BookTypeId),
// 				BookTypeName: v.BookTypeName,
// 				Title:        v.Title,
// 				Description:  v.Description,
// 				Author:       v.Author,
// 				Published:    v.Published.AsTime(),
// 			})
// 		}
// 	}
// 	Data := BKFilter{
// 		Books:      data,
// 		SearchTerm: st,
// 	}

// 	h.ParseBKListTemplate(w, Data)
// }

// func (h Handler) ParseBKListTemplate(w http.ResponseWriter, data any) {
// 	t := h.Templates.Lookup("booklist.html")
// 	if t == nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 	}
// 	if err := t.Execute(w, data); err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 	}
// }