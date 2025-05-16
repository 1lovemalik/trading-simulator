CREATE TABLE users (
    id TEXT PRIMARY KEY, 
    user_name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    passcode TEXT NOT NULL
);

CREATE TABLE portfolios (
    id TEXT PRIMARY KEY,
    user_id TEXT REFERENCES users(id) ON DELETE CASCADE,
    portfolio_name TEXT NOT NULL,
    balance NUMERIC(12, 2) DEFAULT 0.00
);

CREATE TABLE stocks (
    id SERIAL PRIMARY KEY,
    ticker_name TEXT,
    quantity INTEGER,
    portfolio_id TEXT REFERENCES portfolios(id) ON DELETE CASCADE,
    UNIQUE(ticker_name, portfolio_id)  -- Prevents duplicates per portfolio
);

CREATE TABLE trades (
    id SERIAL PRIMARY KEY,
    trade_type TEXT NOT NULL,
    execution_time TIMESTAMP NOT NULL,
    ticker_name TEXT,
    portfolio_id TEXT REFERENCES portfolios(id)
);