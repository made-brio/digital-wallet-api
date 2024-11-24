
-- +migrate Up
-- +migrate StatementBegin

-- Create user_account Table
CREATE TABLE user_account (
    id SERIAL PRIMARY KEY,         -- Auto-incrementing ID for the user
    username VARCHAR(255) NOT NULL,    -- User's name (non-nullable)
    wallet_id INT DEFAULT 0,       -- Linked wallet ID
    password VARCHAR(255) NOT NULL
);

-- Create Wallet Table
CREATE TABLE wallet (
    id SERIAL PRIMARY KEY,         -- Auto-incrementing wallet ID
    user_id INT NOT NULL,          -- Foreign key to the users table
    balance DECIMAL(15, 2) DEFAULT 0.00, -- Current wallet balance (with 2 decimal places)
    status VARCHAR(50) DEFAULT 'active', -- Status of the wallet (e.g., active, frozen)
    FOREIGN KEY (user_id) REFERENCES user_account(id) ON DELETE CASCADE ON UPDATE CASCADE -- Foreign key constraint
);

-- Create Transactions Table
CREATE TABLE transaction (
    id SERIAL PRIMARY KEY,         -- Auto-incrementing transaction ID
    wallet_id INT NOT NULL,        -- Foreign key to the wallets table
    type VARCHAR(50) NOT NULL,     -- Transaction type (topup, transfer)
    amount DECIMAL(15, 2) NOT NULL, -- Transaction amount (with 2 decimal places)
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Timestamp of the transaction (defaults to current time)
    other_party VARCHAR(255),      -- ID of the other wallet involved (if applicable)
    description TEXT,              -- Description of the transaction
    FOREIGN KEY (wallet_id) REFERENCES wallet(id) ON DELETE CASCADE -- Foreign key constraint
);

-- Create GroupWallets Table
CREATE TABLE group_wallet (
    id SERIAL PRIMARY KEY,                -- Auto-incrementing group wallet ID
    name VARCHAR(255) NOT NULL,           -- Name of the group wallet
    goal DECIMAL(15, 2) NOT NULL,         -- Goal amount for the group wallet
    created_by INT NOT NULL,              -- User ID who created the group wallet
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Creation timestamp
    FOREIGN KEY (created_by) REFERENCES user_account(id) ON DELETE CASCADE ON UPDATE CASCADE -- Foreign key constraint
);

-- Create GroupWalletMembers Table
CREATE TABLE group_wallet_member (
    id SERIAL PRIMARY KEY,                -- Auto-incrementing group wallet member ID
    group_wallet_id INT NOT NULL,         -- Foreign key to the group_wallet table
    user_id INT NOT NULL,                 -- Foreign key to the users table
    name VARCHAR(255) NOT NULL,           -- Name of the member
    contribution DECIMAL(15, 2) NOT NULL DEFAULT 0.00, -- Contribution amount by the member
    FOREIGN KEY (group_wallet_id) REFERENCES group_wallet(id) ON DELETE CASCADE ON UPDATE CASCADE, -- Foreign key constraint
    FOREIGN KEY (user_id) REFERENCES user_account(id) ON DELETE CASCADE ON UPDATE CASCADE -- Foreign key constraint
);

-- +migrate StatementEnd