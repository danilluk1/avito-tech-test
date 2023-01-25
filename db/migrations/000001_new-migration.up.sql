CREATE TABLE announcements (
                               id SERIAL PRIMARY KEY,
                               name VARCHAR(200) NOT NULL,
                               description VARCHAR(1000) NOT NULL,
                               price decimal(12, 2) NOT NULL,
                               created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE photos (
                        id SERIAL PRIMARY KEY,
                        link TEXT,
                        announcement_id INTEGER,
                        CONSTRAINT announcement_id
                            FOREIGN KEY (announcement_id)
                                REFERENCES announcements(id) ON DELETE CASCADE
);