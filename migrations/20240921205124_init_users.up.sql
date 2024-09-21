CREATE TABLE IF NOT EXISTS users (
  id varchar not null,
  chat_id varchar not null,
  login varchar not null,
  role varchar not null,
  password varchar not null,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(id)
);
