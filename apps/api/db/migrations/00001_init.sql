-- +goose Up

CREATE TYPE plan_type AS ENUM ('freemium', 'premium');

CREATE TABLE Users(
  _id SERIAL PRIMARY KEY,
  fullname TEXT NOT NULL,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  chessComName TEXT,
  plan plan_type NOT NULL DEFAULT 'freemium',
  reviewCount integer DEFAULT 0,
  lastReviwed DATE DEFAULT CURRENT_DATE
);

CREATE TABLE Games(
  _id UUID PRIMARY KEY,
  gameurl TEXT NOT NULL,
  whiteusername TEXT NOT NULL,
  blackusername TEXT NOT NULL,
  whiterating integer NOT NULL,
  blackrating integer NOT NULL,
  playercolor TEXT NOT NULL,
  timeclass TEXT NOT NULL,
  result TEXT NOT NULL,
  user_id INTEGER NOT NULL REFERENCES Users(_id) ON DELETE CASCADE,
  createdate DATE DEFAULT CURRENT_DATE
);

CREATE TABLE Issues(
  _id UUID PRIMARY KEY,
  game_id UUID NOT NULL REFERENCES Games(_id) ON DELETE CASCADE,
	MoveIndex      int NOT NULL,
	IsMate boolean default false,
	Fen TEXT NOT NULL,
	Move TEXT NOT NULL,
	Color TEXT NOT NULL,
	user_id INTEGER NOT NULL REFERENCES Users(_id) ON DELETE CASCADE,
  createdate DATE DEFAULT CURRENT_DATE
);


-- +goose Down

DROP TABLE IF EXISTS Issues;
DROP TABLE IF EXISTS Games;
DROP TABLE IF EXISTS Users;
