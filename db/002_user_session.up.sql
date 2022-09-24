CREATE TABLE user_sessions (
  id varchar(128) PRIMARY KEY,
  user_id int NOT NULL UNIQUE,
  expire_date Date NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);