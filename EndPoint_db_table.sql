CREATE TABLE generated_texts (
 id SERIAL PRIMARY KEY,
 category_id VARCHAR(250),
 text TEXT,
 status VARCHAR(250),
 createData TIMESTAMP
);