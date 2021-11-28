DROP DATABASE IF EXISTS test_db;
CREATE DATABASE test_db;
DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS Userdata;

\c test_db;

CREATE TABLE Users (
    ID SERIAL,
    Username VARCHAR(100) PRIMARY KEY
);
CREATE TABLE Userdata (
    UserID Int NOT NULL,
    Name VARCHAR(100),
    Surname VARCHAR(100),
    Description VARCHAR(200)
);