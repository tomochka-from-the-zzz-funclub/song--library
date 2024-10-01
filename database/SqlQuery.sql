CREATE TABLE Users (
    userId SERIAL PRIMARY KEY,
    userName VARCHAR(100) NOT NULL UNIQUE,
    userPassword VARCHAR(100) NOT NULL,
    firstName VARCHAR(100) NOT NULL,
    lastName VARCHAR(100) NOT NULL,
    userMail VARCHAR(100) NOT NULL,
    userPhone VARCHAR(100) NOT NULL
);

SELECT *FROM Users;

CREATE TABLE Admins (
    adminId SERIAL PRIMARY KEY,
    adminName VARCHAR(100) NOT NULL UNIQUE,
    adminPassword VARCHAR(100) NOT NULL
);

SELECT *FROM Admins;

CREATE TABLE Authors (
    authorId SERIAL PRIMARY KEY,
    authorName VARCHAR(100) NOT NULL UNIQUE
);

SELECT *FROM Authors;

CREATE TABLE Categories (
    categoryId SERIAL PRIMARY KEY,
    categoryName VARCHAR(100) NOT NULL UNIQUE
);

SELECT *FROM Categories;

CREATE TABLE Publishers (
    publisherId SERIAL PRIMARY KEY,
    publisherName VARCHAR(100) NOT NULL UNIQUE
);

SELECT *FROM Publishers;

CREATE TABLE Books (
    bookId SERIAL PRIMARY KEY,
    bookName VARCHAR(100),
    bookAmount INT,
    bookAuthotId INT,
    FOREIGN KEY (bookAuthotId) REFERENCES Authors(authorId) ON DELETE CASCADE,
    bookCategoryId INT,
    FOREIGN KEY (bookCategoryId) REFERENCES Categories(categoryId) ON DELETE CASCADE,
    bookPublisherId INT,
    FOREIGN KEY (bookPublisherId) REFERENCES Publishers(publisherId) ON DELETE CASCADE
);

SELECT *FROM Books;

CREATE TABLE booksandusers (
    id SERIAL PRIMARY KEY,
    bookid INT NOT NULL,
    FOREIGN KEY (bookid) REFERENCES books(bookid),
    userid INT NOT NULL,
    FOREIGN KEY (userId) REFERENCES Users(userid),
    timestamp timestamp default current_timestamp
);

SELECT *FROM booksandusers;

CREATE TABLE tags(
    id SERIAL PRIMARY KEY,
    bookid INT NOT NULL,
    FOREIGN KEY (bookid) REFERENCES books(bookid),
    tags VARCHAR(100) NOT NULL  
);

SELECT *FROM tags;