CREATE TABLE IF NOT EXISTS users (
  "id" serial PRIMARY KEY,
  "chat_id" varchar NOT NULL,
  "login" varchar NOT NULL,
  "role" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
