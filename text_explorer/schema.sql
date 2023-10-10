CREATE DATABASE hfedoriv_texts;

CREATE TABLE articles (
   id SERIAL PRIMARY KEY,
   title VARCHAR(250) UNIQUE NOT NULL,
   text VARCHAR(5000) NOT NULL,
   user_id INT
);

SELECT * FROM articles;

SELECT * FROM articles WHERE id = 1;

INSERT INTO articles (title, text, user_id)
VALUES ('test title', 'test text', 045);

UPDATE articles
SET title='title updated', text='text updated'
WHERE id = 1;

DELETE FROM articles WHERE id = 1;