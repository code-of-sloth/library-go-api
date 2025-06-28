package book

import (
	"LibraryGo/src/config"
	"LibraryGo/src/utils"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func AddNewBook(name, author, genre, desc, sku string) (data Book, err error) {
	data.BookID = utils.GenerateRandomChar("book", 10)
	data.Author = author
	data.Desc = desc
	data.Genre = genre
	data.Name = name
	data.Sku = sku
	tx, err := config.DbConn.Begin(context.Background())
	if err != nil {
		err = fmt.Errorf("addnewbook:error init conn ")
		return
	}

	defer tx.Commit(context.Background())
	batch := &pgx.Batch{}
	batch.Queue("INSERT INTO library.bookgroup (name,bookdesc,author,genre,sku,createdat,updatedat) VALUES ($1,$2,$3,$4,$5,now(),now()) ON CONFLICT (sku) DO NOTHING;", data.Name, data.Desc, data.Author, data.Genre, data.Sku)
	batch.Queue("INSERT INTO library.books (bookid,sku,createdat,updatedat,isactive) VALUES ($1,$2,now(),now(),true)", data.BookID, data.Sku)
	br := tx.SendBatch(context.Background(), batch)
	defer br.Close()
	_, err = br.Exec()
	if err != nil {
		tx.Rollback(context.Background())
		err = fmt.Errorf("addnewbook:error adding book")
		return
	}

	return
}

func RemoveBookByID(bookID string) (err error) {
	_, err = config.DbConn.Exec(context.Background(), "UPDATE library.books SET isactive=false,deletedat=now(),updatedat=now() WHERE isactive=true AND bookid=$1", bookID)
	if err != nil {
		err = fmt.Errorf("removebookbyid:error deleting book %w", err)
		return
	}
	return
}

func GetAvailableBooks(pageNum, pageSize int) (data []Book, err error) {
	rows, err := config.DbConn.Query(
		context.Background(),
		"SELECT DISTINCT i.sku,i.name,i.bookdesc,i.author,i.genre,COUNT(b.bookid) OVER (PARTITION BY i.sku),i.createdat FROM library.bookgroup i JOIN library.books b ON i.sku=b.sku WHERE b.isactive=true ORDER BY i.createdat LIMIT $1 OFFSET $2",
		pageSize, (pageNum-1)*pageSize,
	)
	if err != nil {
		err = fmt.Errorf("getavailablebooks:error querying for books %w", err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var bk Book
		var createdat any
		err = rows.Scan(&bk.Sku, &bk.Name, &bk.Desc, &bk.Author, &bk.Genre, &bk.AvailableCount, &createdat)
		if err != nil {
			err = fmt.Errorf("getavailablebooks:error reading for books %w", err)
			return
		}

		data = append(data, bk)
	}

	if len(data) == 0 {
		data = nil
	}
	return
}

// func GetBookByID(bookID string) (data *Book, err error) {
// 	row, err := config.DbConn.Query(context.Background(), "SELECT name,bookdesc,author,genre,isrented FROM library.books WHERE bookid=$1 LIMIT 1", bookID)
// 	if err != nil {
// 		err = fmt.Errorf("getbookbyid:error fetching book with id:%s %w", bookID, err)
// 		return
// 	}

// 	defer row.Close()
// 	if row.Next() {
// 		var bk Book
// 		err = row.Scan(&bk.Name, &bk.Desc, &bk.Author, &bk.Genre, &bk.IsRented)
// 		if err != nil {
// 			err = fmt.Errorf("getbookbyid:error reading book with id:%s %w", bookID, err)
// 			return
// 		}

// 		bk.BookID = bookID
// 		data = &bk
// 	}

// 	return
// }
