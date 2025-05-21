-- I'm a comment
CREATE TABLE author (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  bio  text
);

-- I also deserve a comment
CREATE TABLE book (
  id          BIGSERIAL PRIMARY KEY,
  name        text      NOT NULL,
  description text
);