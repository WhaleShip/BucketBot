CREATE TABLE Notes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    text VARCHAR(255) NOT NULL
);

CREATE TABLE Users (
    user_id INTEGER PRIMARY KEY,
);

CREATE TABLE users_notes (
    user_id INTEGER,
    note_id INTEGER,
    PRIMARY KEY (user_id, note_id),
    FOREIGN KEY (user_id) REFERENCES Users (id),
    FOREIGN KEY (note_id) REFERENCES Notes (id)
);