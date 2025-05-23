
# crypTracker

A real time cryptocurrency tracker to stay on top of your precious coins. Get lightning fast updates about the 100 most popular cryptocurrencies, and add them to your watchlist to keep a close eye on it. With a sleek UI, you get exactly what you need without having to go through a bajillion numbers and scary graphs.



## Tech Stack

**Client:** React (Typescript), Zustand, Websockets, TailwindCSS.

**Server:** Go, Fiber, GORM, PostgreSQL.

The backend is built with Go for performance and scalability, paired with PostgreSQL to handle large datasets and frequent queries. Fiber is used over Gin due to its native WebSocket support. On the frontend, React with TypeScript ensures type safety, while Zustand manages global state efficiently. TailwindCSS enables rapid and customizable UI development. WebSockets are preferred over Socket.io for better Go compatibility, and JWT handles authentication seamlessly.

## Installation

#### Prerequisites
```bash
- Go (1.23.4 or higher)
- npm and Node.js
- PostgreSQL 17 (or higher)
```

1. Build the backend
```bash
  go build cmd/crypTracker/main.go
```

2. Install node dependencies in `web` and build the frontend
```bash
  cd web
  npm install
  npm run build
```

3. Create a .env file in the `root` directory with the following values. DB_URL must be for PostgreSQL with the following keys.
```bash
COINGECKO_API_KEY="your_API_key"
DB_URL="host=localhost user={username} password={password} dbname={database_name} port=5432 sslmode=disable"
RUNTIME_ENV="production"
SECRET="JWT Secret"
VITE_SECRET="Same secret as SECRET"
PORT=8080
VITE_PORT=8080
```

4. Go back to the `root` directory and run the build file

```bash
  cd ..
  ./main
```

The app must now be running on `http://localhost:PORT`
## API Reference

### Coin API

#### Test API
```http
GET /test
```
**Response:**
```json
"Testing mount"
```
---

#### Get All Coins
```http
GET /crypto/coins
```
**Response:**
```json
[
  {
    "coin_gecko_id": "string",
    "symbol": "string",
    "name": "string",
    "current_price": "number",
    "market_cap": "number",
    "updated_at": "timestamp",
    "image": "string"
  }
]
```
---

#### Get Coin Details
```http
GET /crypto/coins/:coinId
```
**Response:**
```json
{
  "coin_gecko_id": "string",
  "symbol": "string",
  "name": "string",
  "current_price": "number",
  "market_cap": "number",
  "updated_at": "timestamp",
  "image": "string",
  "description": "string",
  "graph_data": "array"
}
```
**Error Response:**
```json
{
  "message": "'coinId' doesn't exist",
  "status": 400
}
```

### User API

#### Create User
```http
POST /users/
```

**Request Body:**
```json
{
  "username": "string",
  "password": "string"
}
```

**Response:**
```json
{
  "message": "201 : Successfully registered user",
  "user_data": {
    "user_id": "string",
    "username": "string"
  }
}
```

---

#### Login
```http
POST /users/login
```

**Request Body:**
```json
{
  "username": "string",
  "password": "string"
}
```

**Response:**
```json
{
  "message": "Successfully logged in user : username!",
  "user_id": "string",
  "user_data": {
    "user_id": "string",
    "username": "string"
  }
}
```

---

#### Sign Out
```http
GET /users/:userId/signout
```

**Response:**
```json
{
  "message": "Successfully signed out"
}
```

---

#### Get User Details
```http
GET /users/:userId
```

**Response:**
```json
{
  "user_id": "string",
  "username": "string",
  "profile_image": "string",
  "created_at": "timestamp"
}
```

---

#### Delete User
```http
DELETE /users/:userId
```

**Response:**
```json
{
  "message": "Successfully deleted user : userId"
}
```

---

#### Update User
```http
PUT /users/:userId
```

**Request Body:**
```json
{
  "username": "string",
  "password": "string"
}
```

**Response:**
```json
{
  "user_id": "string",
  "username": "string"
}
```

---

### Watchlist API

#### Get Watchlist
```http
GET /users/:userId/watchlist
```

**Response:**
```json
[
  {
    "coin_gecko_id": "string",
    "symbol": "string",
    "name": "string",
    "current_price": "number",
    "market_cap": "number",
    "updated_at": "timestamp",
    "image": "string"
  }
]
```

---

#### Add Coin to Watchlist
```http
POST /users/:userId/watchlist
```

**Request Body:**
```json
{
  "coin_id": "string"
}
```

**Response:**
```json
{
  "user_id": "string",
  "coin_gecko_id": "string"
}
```

---

#### Remove Coin from Watchlist
```http
DELETE /users/:userId/watchlist
```

**Request Body:**
```json
{
  "coin_id": "string"
}
```

**Response:**
```json
{
  "message": "Successfully deleted coin_id from watchlist"
}
```

### WebSocket API

#### WebSocket Connection
```http
GET /ws
```

**Description:**
Establishes a WebSocket connection for real-time communication.

**Headers:**
- `Upgrade: websocket`
- `Connection: Upgrade`

**Response:**
- On success, the client is upgraded to a WebSocket connection.
- On failure, returns:

```json
{
  "message": "426 Upgrade Required"
}
```

---
## Upcoming Features

I've planned on including the following features in the near future
- Coin price graphs (InfluxDB + Grafana)
- Dockerfile (to make setup easier)
- PFP storage using Amazon S3 buckets