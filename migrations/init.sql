CREATE TABLE Notes (
    id SERIAL PRIMARY KEY,
    text VARCHAR(255) NOT NULL
);

CREATE TABLE Users (
    user_id INTEGER PRIMARY KEY,
    note_id INTEGER,
    FOREIGN KEY (note_id) REFERENCES Notes (id)
);