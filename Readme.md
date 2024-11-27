# Digital Wallet API

A backend system for managing digital wallets, transactions, group wallets, and contributions. Built with **Golang** and **Gin Framework**, this project offers robust features for secure and scalable wallet management.

---

## Features

1. **User Management**  
   - Register and login functionality with secure password storage.
   - Update, retrieve, and delete user profiles.

2. **Wallet Management**  
   - Create wallets for users with initial balance setup.
   - Perform wallet operations such as top-up, balance inquiry, and wallet freezing/unfreezing.

3. **Transaction Handling**  
   - Transfer funds between wallets with detailed transaction records.
   - Retrieve transaction histories, categorized by income, expense, or full details.

4. **Group Wallet Management**  
   - Create shared wallets with a savings goal and manage group contributions.
   - Add, remove, and update group members and their contributions.

5. **Authentication and Authorization**  
   - Secure endpoints using JWT-based middleware to protect sensitive data.

---

## Project Structure

```plaintext
.
├── config/               # Configuration files (e.g., environment variables)
├── database/             # Database schema and initialization scripts
├── internal/
│   ├── controllers/      # HTTP controllers for handling requests
│   ├── middleware/       # JWT authentication middleware
│   ├── models/           # Database models and payload definitions
│   ├── repository/       # Data access layer for database queries
│   ├── service/          # Business logic for different modules
├── routes/               # API route definitions
├── utils/                # Utility functions (e.g., hashing)
├── main.go               # Application entry point
└── go.mod                # Go module definition
```

---

## API Endpoints

### Authentication
- `POST /api/users/login` - Login and retrieve JWT token.
- `POST /api/users/register` - Register a new user.

### User
- `GET /api/users/` - Get all users.
- `GET /api/users/:id` - Get a user by ID.
- `PUT /api/users/:id` - Update user details.
- `DELETE /api/users/:id` - Delete a user.

### Wallet
- `POST /api/wallet/:id` - Create a wallet.
- `GET /api/wallet/:wallet_id/balance` - Check wallet balance.
- `GET /api/wallet/:wallet_id/info` - Retrieve wallet details.
- `PUT /api/wallet/:wallet_id/topup` - Add funds to a wallet.
- `PUT /api/wallet/:wallet_id/freeze` - Freeze a wallet.
- `PUT /api/wallet/:wallet_id/unfreeze` - Unfreeze a wallet.

### Transaction
- `POST /api/transactions/transfer` - Transfer funds between wallets.
- `GET /api/transactions/:wallet_id/income` - View income transaction history.
- `GET /api/transactions/:wallet_id/expense` - View expense transaction history.
- `GET /api/transactions/:wallet_id` - View all transactions for a wallet.

### Group Wallet
- `POST /api/gw/:wallet_id` - Create a group wallet.
- `PUT /api/gw/:wallet_id` - Update group wallet goal.
- `GET /api/gw/:wallet_id` - Retrieve group wallet details.
- `DELETE /api/gw/:wallet_id` - Delete a group wallet.

### Group Wallet Member
- `POST /api/gwm/:wallet_id` - Add a member to a group wallet.
- `PUT /api/gwm/:wallet_id/:member_id` - Update a member's contribution.
- `GET /api/gwm/:wallet_id` - Get all group wallet members.
- `DELETE /api/gwm/:wallet_id/:member_id` - Remove a group wallet member.

---

## Installation and Setup

### Prerequisites
- **Golang**: Ensure Go is installed (v1.17+ recommended).
- **PostgreSQL**: Install and set up PostgreSQL as the database.

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/made-brio/digital-wallet-api.git
   cd digital-wallet-api
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Set up the database:
   - Create a PostgreSQL database.
   - Run migration scripts from the `database/migrations/` folder.

4. Configure `.env` file:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   JWT_SECRET=your_jwt_secret
   ```

5. Start the application:
   ```bash
   go run main.go
   ```

6. Access API at:  
   `http://localhost:8085`

---

## Tech Stack

- **Language**: Go (Golang)
- **Framework**: Gin Web Framework
- **Database**: PostgreSQL
- **Authentication**: JSON Web Token (JWT)

---

## Future Improvements

- Group Wallet Transaction Logic 

---

## License

This project is licensed under the MIT License. See the LICENSE file for details.
