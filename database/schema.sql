-- Enable UUID generation
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Players Table
CREATE TABLE players (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

-- Games Table
CREATE TABLE games (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    player_1 UUID REFERENCES players(id) NOT NULL,
    player_2 UUID REFERENCES players(id) NOT NULL,
    player_1_choice TEXT CHECK (player_1_choice IN ('rock', 'paper', 'scissors')),
    player_2_choice TEXT CHECK (player_2_choice IN ('rock', 'paper', 'scissors')),
    winner UUID REFERENCES players(id)
);


