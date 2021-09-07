CREATE TABLE IF NOT EXISTS persons (
                         id INT PRIMARY KEY,
                         userId INT NOT NULL,
                         firstName VARCHAR ( 100 ),
                         lastName VARCHAR ( 100 ),
                         middleName VARCHAR ( 100 ),
                         createdAt TIMESTAMP NOT NULL DEFAULT (now()),
                         updatedAt TIMESTAMP,
                         deletedAt TIMESTAMP
);