CREATE TABLE author (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  bio  text
);

CREATE TABLE book (
  id          BIGSERIAL PRIMARY KEY,
  name        text      NOT NULL,
  description text
);