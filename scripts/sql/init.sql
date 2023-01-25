CREATE TABLE announcement (
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    description VARCHAR(1000) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE photos (
    link TEXT PRIMARY KEY,
    announcement_id INTEGER,
        CONSTRAINT announcement_id
            FOREIGN KEY (announcement_id)
                REFERENCES announcement(id) ON DELETE CASCADE
);