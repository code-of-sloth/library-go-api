package book

import (
	"LibraryGo/src/config"
	"LibraryGo/src/user"
	"LibraryGo/src/utils"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func LendBookToUser(userID, sku string) (err error) {
	_, err = user.FetchUserByID(userID)
	if err != nil {
		err = fmt.Errorf("lendbooktouser:error fetching active user for id:%s %w", userID, err)
		return
	}

	bookID, err := isBookAvailable(sku)
	if err != nil {
		err = fmt.Errorf("lendbooktouser:error checking availability of sku:%s ", sku)
		return
	} else if bookID != "" {
		err = fmt.Errorf("lendbooktouser:book is not available for lending")
		return
	}

	tx, err := config.DbConn.Begin(context.Background())
	if err != nil {
		err = fmt.Errorf("lendbooktouser:error init conn ")
		return
	}

	defer tx.Commit(context.Background())
	batch := &pgx.Batch{}
	batch.Queue("UPDATE library.books SET isrented=true,updatedat=now() WHERE bookid=$1", bookID)
	lendingID := utils.GenerateRandomChar("TXN", 21)
	returnDate := utils.GetPgTime(time.Now().AddDate(0, 0, 7))
	batch.Queue("INSERT INTO library.lending (lendingid,userid,bookid,returndate,createdat,updatedat) VALUES ($1,$2,$3,$4,now(),now())", lendingID, userID, bookID, returnDate)
	br := tx.SendBatch(context.Background(), batch)
	defer br.Close()
	_, err = br.Exec()
	if err != nil {
		tx.Rollback(context.Background())
		err = fmt.Errorf("lendbooktouser:error updating lending details")
		return
	}

	return
}

func RetunBookByUser(userID, bookID string) (err error) {
	_, err = user.FetchUserByID(userID)
	if err != nil {
		err = fmt.Errorf("retunbookbyuser:error fetching active user for id:%s %w", userID, err)
		return
	}

	lendRec, err := fetchLendingRecord(bookID)
	if err != nil {
		err = fmt.Errorf("retunbookbyuser:error fetching lending rec for bookid:%s %w", bookID, err)
		return
	}

	if userID != lendRec.UserID {
		err = fmt.Errorf("retunbookbyuser:only the user who borreowed the book can return it")
		return
	}

	if time.Now().After(lendRec.ReturnDate) {
		err = fmt.Errorf("retunbookbyuser:user has crossed the return date")
		return
	}

	tx, err := config.DbConn.Begin(context.Background())
	if err != nil {
		err = fmt.Errorf("retunbookbyuser:error init conn ")
		return
	}

	defer tx.Commit(context.Background())
	batch := &pgx.Batch{}
	batch.Queue("UPDATE library.books SET isrented=false,updatedat=now() WHERE bookid=$1", bookID)
	batch.Queue("UPDATE library.lending SET returnedat=now(),isreturned=true,updatedat=now() WHERE bookid=$1 AND isreturned=false", bookID)
	br := tx.SendBatch(context.Background(), batch)
	defer br.Close()
	_, err = br.Exec()
	if err != nil {
		tx.Rollback(context.Background())
		err = fmt.Errorf("retunbookbyuser:error updating lending details")
		return
	}

	return
}
