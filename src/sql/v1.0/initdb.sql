CREATE SCHEMA library;

CREATE TABLE IF NOT EXISTS library.users(  
    userid VARCHAR(15) PRIMARY KEY,
    name VARCHAR(15) NOT NULL,
    mobile VARCHAR(10) NOT NULL,
    createdat timestamp NOT NULL,
    updatedat timestamp NOT NULL,
    deletedat timestamp,
    isactive BOOLEAN NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS unique_active_mobile ON library.users(mobile) WHERE isactive = true;

CREATE TABLE IF NOT EXISTS library.bookgroup(
    sku VARCHAR(15) PRIMARY KEY,
    name VARCHAR(25) NOT NULL,
    bookdesc VARCHAR(250) NOT NULL,
    author VARCHAR(25) NOT NULL,
    genre VARCHAR(25) NOT NULL,
    createdat timestamp NOT NULL,
    updatedat timestamp NOT NULL
)

CREATE TABLE IF NOT EXISTS library.books(  
    bookid VARCHAR(15) PRIMARY KEY,
    sku VARCHAR(15) REFERENCES library.bookgroup (sku) NOT NULL,
    isrented BOOLEAN NOT NULL DEFAULT false,
    createdat timestamp NOT NULL,
    updatedat timestamp NOT NULL,
    deletedat timestamp,
    isactive BOOLEAN NOT NULL 
);

CREATE TABLE IF NOT EXISTS library.lending(
    lendingid VARCHAR(25) PRIMARY KEY,
    userid VARCHAR(15) REFERENCES library.users (userid) NOT NULL,
    bookid VARCHAR(15) REFERENCES library.books (bookid) NOT NULL,
    returndate timestamp NOT NULL,
    returnedat timestamp,
    isreturned BOOLEAN NOT NULL DEFAULT false,
    createdat timestamp NOT NULL,
    updatedat timestamp NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS unique_book_rent ON library.lending(bookid) WHERE isreturned = false;

