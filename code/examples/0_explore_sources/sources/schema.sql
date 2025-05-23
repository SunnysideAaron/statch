-- I'm a comment
CREATE TABLE author (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL, -- In line comment
  bio  text
);

/* I also deserve a comment
   on a second line.
*/
CREATE TABLE book (
  id          BIGSERIAL PRIMARY KEY,
  name        text      NOT NULL,
  description text
);