# Ζήτημα 1: Ανάπτυξη RESTful Υπηρεσίας

## Εισαγωγή

Σε αυτό το ζήτημα, υλοποιήσαμε μια RESTful υπηρεσία διαχείρισης χρηστών με τις ακόλουθες προδιαγραφές:

- Γλώσσα προγραμματισμού: Go
- Web Framework: Gin
- ORM: GORM
- Βάση Δεδομένων: PostgreSQL

## Προαπαιτούμενα

- Go 1.22.2 ή νεότερη έκδοση
- PostgreSQL 13 ή νεότερη έκδοση
- Git

## Δομή Έργου

```
app/
├── cmd/
│   └── main.go
├── internal/
│   ├── handlers/
│   │   └── user_handlers.go
│   ├── models/
│   │   └── user.go
│   ├── routes/
│   │   └── routes.go
│   └── validators/
│       └── user_validators.go
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
└── README.md
```

## Εγκατάσταση και Ρύθμιση

1. Κλωνοποίηση του repository:
   ```
   git clone https://github.com/meter-peter/icsd-cloud-2024.git
   cd icsd-cloud-2024/app
   ```

2. Εγκατάσταση εξαρτήσεων:
   ```
   go mod tidy
   ```

3. Δημιουργία και ρύθμιση του αρχείου .env:
   ```
   cp .env.example .env
   ```
   Επεξεργαστείτε το .env με τις κατάλληλες τιμές για τη βάση δεδομένων.

4. Αρχικοποίηση της βάσης δεδομένων:
   ```
   psql -U your_username -d your_database_name -f ../migrations/init.sql
   ```

## Εκτέλεση της Εφαρμογής

```
go run cmd/main.go
```

Η εφαρμογή θα ξεκινήσει στη διεύθυνση `http://localhost:8080`.

## API Endpoints

- `GET /users`: Ανάκτηση όλων των χρηστών
- `GET /users/:id`: Ανάκτηση συγκεκριμένου χρήστη
- `POST /users`: Δημιουργία νέου χρήστη
- `PUT /users/:id`: Ενημέρωση υπάρχοντος χρήστη
- `DELETE /users/:id`: Διαγραφή χρήστη
- `GET /users/search`: Αναζήτηση χρηστών με βάση κριτήρια

## Swagger UI

Το Swagger UI είναι διαθέσιμο στη διεύθυνση: `http://localhost:8080/swagger/index.html`


[Επόμενο: Ζήτημα 2 - Δοχειοποίηση με Docker](../zitima2-scripts/README.md)