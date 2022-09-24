SET time_zone='Asia/Tokyo';
CREATE TABLE users (
  id int AUTO_INCREMENT PRIMARY KEY,
  name varchar(128) NOT NULL,
  email varchar(128) NOT NULL UNIQUE,
  password_hash varchar(128) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;mmmP