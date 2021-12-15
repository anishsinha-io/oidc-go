CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    user_id FOREIGN KEY REFERENCES users(id) ON DELETE CASCADE,
    access_token VARCHAR(255),
    refresh_token VARCHAR(255),
    session_hash VARCHAR(500),
    authorization_code VARCHAR(255),
    state VARCHAR(255),
);