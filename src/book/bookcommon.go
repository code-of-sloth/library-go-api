package book

import (
	"LibraryGo/src/config"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgtype"
)

type Book struct {
	BookID         string `json:"bookID,omitempty"`
	Name           string `json:"name"`
	Author         string `json:"author"`
	Genre          string `json:"genre"`
	Desc           string `json:"desc"`
	AvailableCount int    `json:"availableCount"`
	Sku            string `json:"sku"`
}

type bookLendingRec struct {
	BookID     string    `json:"bookID"`
	UserID     string    `json:"userID"`
	ReturnDate time.Time `json:"returnDate"`
}

func isBookAvailable(sku string) (bookID string, err error) {
	row, err := config.DbConn.Query(context.Background(), "SELECT bookid FROM library.books WHERE sku=$1 AND isactive=true AND isrented=false LIMIT 1", sku)
	if err != nil {
		err = fmt.Errorf("isbookavailable:error querying books for sku %s %w", sku, err)
		return
	}

	defer row.Close()
	if row.Next() {
		err = row.Scan(&bookID)
		if err != nil {
			err = fmt.Errorf("isbookavailable:error reading for sku:%s %w", sku, err)
			return
		}
	} else {
		err = fmt.Errorf("isbookavailable:book with sku %s is not available", bookID)
		return
	}

	return
}

func fetchLendingRecord(bookID string) (lendRec bookLendingRec, err error) {
	row, err := config.DbConn.Query(context.Background(), "SELECT userid,returndate FROM library.lending WHERE bookid=$1 AND isreturned=false LIMIT 1;", bookID)
	if err != nil {
		err = fmt.Errorf("fetchlendingrecord:error querying lending records for id:%s %w", bookID, err)
		return
	}

	defer row.Close()
	if row.Next() {
		var returnDate pgtype.Timestamp
		err = row.Scan(&lendRec.UserID, &returnDate)
		if err != nil {
			err = fmt.Errorf("fetchlendingrecord:error reading lending records for id:%s %w", bookID, err)
			return
		}

		lendRec.BookID = bookID
		lendRec.ReturnDate = returnDate.Time
	} else {
		err = fmt.Errorf("fetchlendingrecord:unable to find lending records for id:%s %w", bookID, err)
		return
	}

	return
}
