CREATE TABLE IF NOT EXISTS books(
   id INT AUTO_INCREMENT PRIMARY KEY,
   title VARCHAR(40),
   author VARCHAR(40)
);

INSERT IGNORE INTO books
    (id, title, author)
VALUES
    (1, 'Doraemon', 'Fujiko F. Fujio'),
    (2, 'Harry Potter', 'J.K. Rowling'),
    (3, 'The Lord of The Ring', 'J.R.R. Tolkien');