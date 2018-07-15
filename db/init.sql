CREATE TYPE CEFR as ENUM ('A1','A2','B1','B2','C1','C2');

CREATE TABLE reading(
    id INT PRIMARY KEY     NOT NULL,
    title CHAR(50),
    cefr CEFR,
    text TEXT
);

INSERT INTO reading VALUES (1,'Mary Had a Little Lamb','A1','Mary had a little lamb whose fleece was white as snow.');