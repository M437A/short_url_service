CREATE TABLE IF NOT EXISTS links (
    id SERIAL   PRIMARY KEY,
    user_id     INT NOT NULL,
    short_url   VARCHAR(255) NOT NULL,
    long_url    TEXT NOT NULL
    );

CREATE TABLE IF NOT EXISTS users (
    id          SERIAL PRIMARY KEY,
    user_name   VARCHAR(255) NOT NULL,
    email       VARCHAR(255) NOT NULL
    );

CREATE TABLE IF NOT EXISTS user_links (
    user_id INT REFERENCES users(id),
    link_id INT REFERENCES links(id),
    PRIMARY KEY (user_id, link_id)
    );

CREATE INDEX IF NOT EXISTS idx_user_links_user_id ON user_links(user_id);
CREATE INDEX IF NOT EXISTS idx_user_links_link_id ON user_links(link_id);
