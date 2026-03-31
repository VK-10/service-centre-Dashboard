# 🔧 Service Center Tracker

A real-time vehicle service management system built with **Go**, **Gin**, **GORM**, and **SQLite**. Customers can submit their vehicle for service and track its status live. Admins manage all services from a protected dashboard.

---

## Features

- **Customer Portal** — Submit a vehicle service request and track status in real time via SSE
- **Admin Dashboard** — View all active services, update status, delete records
- **Live Updates** — Status changes reflect instantly on the customer page without a refresh
- **Session Auth** — Cookie-based admin login with bcrypt password hashing
- **Custom Validators** — Form validation for vehicle types and issue categories

---

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go 1.21+ |
| Web Framework | Gin |
| ORM | GORM |
| Database | SQLite |
| Sessions | gin-contrib/sessions (GORM store) |
| Real-time | Server-Sent Events (SSE) |
| Frontend | Go HTML Templates + Tailwind CSS |
| ID Generation | teris-io/shortid |
| Password Hashing | bcrypt |

---

## Project Structure

```
service-center/
├── cmd/
│   ├── main.go           # Entry point
│   ├── handlers.go       # Handler struct and constructor
│   ├── admin.go          # Login, logout, dashboard, status update, delete
│   ├── customers.go      # New vehicle form, vehicle submission, customer view
│   ├── events.go         # SSE handlers and stream logic
│   ├── middleware.go     # Auth middleware
│   ├── notifications.go  # NotificationManager (pub/sub over channels)
│   ├── routes.go         # Route registration
│   ├── utils.go          # Config, template loading, session helpers
│   └── validators.go     # Custom Gin validators
├── internal/
│   └── models/
│       ├── models.go     # DB init and AutoMigrate
│       ├── vehicles.go   # Vehicle/VehicleItem models and queries
│       └── user.go       # User model, auth, lookup
├── templates/
│   ├── base.tmpl         # Shared top/bottom layout
│   ├── order.tmpl        # New vehicle service form
│   ├── customer.tmpl     # Customer order tracking page
│   ├── admin.tmpl        # Admin dashboard
│   ├── login.tmpl        # Admin login
│   └── static/           # Static assets (logo, etc.)
├── data/
│   └── vehicles.db       # SQLite database (auto-created)
└── go.mod
```

---

## Getting Started

### Prerequisites

- Go 1.21 or higher
- GCC (required by SQLite driver — use [TDM-GCC](https://jmeubank.github.io/tdm-gcc/) on Windows)

### Installation

```bash
# Clone the repo
git clone https://github.com/your-username/service-center.git
cd service-center

# Install dependencies
go mod tidy

# Run the server
go run ./cmd
```

The server starts on `http://localhost:8080` by default.

---

## Environment Variables

| Variable | Default | Description |
|---|---|---|
| `PORT` | `8080` | Server port |
| `DATABASE_URL` | `./data/vehicles.db` | SQLite DB path |
| `SESSION_SECRET_KEY` | `service-centre-secret-key` | Session encryption key |

Set them in your shell or a `.env` file before running.

---

## Creating the First Admin User

There is no signup page — admin users must be inserted directly into the database.

**1. Generate a bcrypt password hash** (in Go):
```go
hash, _ := bcrypt.GenerateFromPassword([]byte("yourpassword"), 12)
fmt.Println(string(hash))
```

**2. Create an `insert.sql` file:**
```sql
INSERT INTO users (id, username, password) VALUES ('admin01234567', 'admin', '$2a$12$yourHashHere');
```

**3. Run it:**
```bash
cd data
sqlite3 vehicles.db ".read insert.sql"
```

---

## Routes

### Public

| Method | Path | Description |
|---|---|---|
| GET | `/` | New vehicle service form |
| POST | `/new-vehicle` | Submit service request |
| GET | `/customers/:id` | Customer tracking page |
| GET | `/notifications?vehicleId=:id` | SSE stream for customer |
| GET | `/login` | Admin login page |
| POST | `/login` | Admin login submit |
| GET | `/logout` | Admin logout |

### Admin (requires auth)

| Method | Path | Description |
|---|---|---|
| GET | `/admin` | Admin dashboard |
| POST | `/admin/vehicles/:id/status` | Update vehicle status |
| POST | `/admin/vehicles/:id` | Delete vehicle |
| GET | `/admin/notifications` | SSE stream for admin |

---

## Vehicle Statuses

Vehicles move through these stages:

```
Checked In → In Service → On Hold → Quality Check → Ready
```

---

## How Real-Time Updates Work

```
Admin updates status
       ↓
HandleOrderPut saves to DB
       ↓
NotificationManager.Notify("vehicle:<id>", JSON payload)
       ↓
SSE pushes event to connected customer browser
       ↓
Customer page updates status bar without reload
```

The `NotificationManager` uses a `map[string]map[chan string]bool` protected by a `sync.RWMutex` to fan out messages to all connected clients for a given vehicle key.

---

## Known Limitations

- No signup flow — admin accounts must be seeded manually via SQL
- SQLite is single-writer; swap to PostgreSQL for production concurrency
- No HTTPS config — add a reverse proxy (Nginx/Caddy) before deploying
- Session `Secure: true` requires HTTPS — set to `false` for local development

---

## License

MIT
