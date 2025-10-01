# Hotel Bookings & Reservations System

A comprehensive hotel booking and reservation management system built with Go backend and HTML/CSS/JavaScript frontend, featuring room availability checking, reservation management, and administrative dashboard.

## 🚀 Features

- **Room Management**: Multiple room types (General's Quarters, Major's Suite)
- **Reservation System**: Complete booking workflow with availability checking
- **User Authentication**: Secure login system for administrators
- **Admin Dashboard**: Comprehensive reservation management interface
- **Email Integration**: Automated email notifications for bookings
- **Session Management**: Secure session handling with SCS
- **Database Integration**: PostgreSQL with migration support
- **Form Validation**: Client and server-side validation
- **Responsive Design**: Mobile-friendly Bootstrap interface
- **Calendar Integration**: Visual availability calendar
- **Search Functionality**: Advanced availability search

## 🛠 Tech Stack

### Backend
- **Language**: Go 1.23.4
- **Router**: Chi Router (github.com/go-chi/chi/v5)
- **Database**: PostgreSQL
- **Session Management**: SCS (github.com/alexedwards/scs/v2)
- **CSRF Protection**: NoSurf (github.com/justinas/nosurf)
- **Email**: Simple Mail (github.com/xhit/go-simple-mail/v2)
- **Database Migrations**: Buffalo Pop/Fizz

### Frontend
- **Framework**: Server-side rendered HTML templates
- **Styling**: Bootstrap 5
- **JavaScript**: Vanilla JS with SweetAlert2
- **Icons**: Font Awesome
- **Date Picker**: Vanillajs-datepicker
- **Notifications**: Notie.js

### Database & Tools
- **Database**: PostgreSQL
- **Migration Tool**: Buffalo Pop
- **Template Engine**: Go HTML templates
- **Build Tool**: Go build system

## 📁 Project Structure

```
bookings/
├── cmd/                          # Application entry points
│   └── web/                      # Web application
│       ├── main.go              # Main application entry
│       ├── routes.go            # HTTP routes configuration
│       ├── middlewares.go       # HTTP middleware
│       ├── send-mail.go         # Email service
│       └── *_test.go            # Web layer tests
├── internal/                    # Private application code
│   ├── config/                  # Application configuration
│   │   └── config.go           # Config structures
│   ├── drivers/                # Database drivers
│   │   └── drivers.go          # Database connection
│   ├── forms/                  # Form handling and validation
│   │   ├── forms.go            # Form validation logic
│   │   ├── errors.go           # Form error handling
│   │   └── forms_test.go       # Form tests
│   ├── handlers/               # HTTP request handlers
│   │   ├── handlers.go         # Route handlers
│   │   ├── handlers_test.go    # Handler tests
│   │   └── setup_test.go       # Test setup utilities
│   ├── helpers/                # Utility functions
│   │   └── helpers.go          # Helper functions
│   ├── models/                 # Data models
│   │   ├── models.go           # Database models
│   │   └── templatedata.go     # Template data structures
│   ├── render/                 # Template rendering
│   │   ├── render.go           # Template renderer
│   │   ├── render_test.go      # Render tests
│   │   └── setup_test.go       # Render test setup
│   └── repository/             # Data access layer
│       ├── repository.go       # Repository interfaces
│       └── dbrepo/             # Database repository implementations
├── templates/                   # HTML templates
│   ├── base.layout.tmpl        # Base layout template
│   ├── admin.layout.tmpl       # Admin layout template
│   ├── home.page.tmpl          # Homepage template
│   ├── about.page.tmpl         # About page template
│   ├── contact.page.tmpl       # Contact page template
│   ├── generals.page.tmpl      # General's Quarters room page
│   ├── majors.page.tmpl        # Major's Suite room page
│   ├── search-availability.page.tmpl    # Search form
│   ├── make-reservations.page.tmpl      # Reservation form
│   ├── reservation-summary.page.tmpl    # Booking summary
│   ├── choose-room.page.tmpl            # Room selection
│   ├── login.page.tmpl                  # Admin login
│   ├── admin-dashboard.page.tmpl        # Admin dashboard
│   ├── admin-reservations-calender.page.tmpl  # Calendar view
│   ├── admin-all-reservations.page.tmpl       # All reservations
│   ├── admin-new-reservations.page.tmpl       # New reservations
│   └── admin-reservation-show.page.tmpl       # Reservation details
├── static/                      # Static assets
│   ├── css/                    # Stylesheets
│   │   └── styles.css          # Custom styles
│   ├── js/                     # JavaScript files
│   │   └── app.js              # Application JavaScript
│   ├── images/                 # Image assets
│   │   ├── outside.png         # Hotel exterior
│   │   ├── tray.png           # Service tray
│   │   ├── woman-laptop.png    # Hero image
│   │   ├── generals-quarters.png  # Room image
│   │   └── marjors-suite.png      # Room image
│   └── admin/                  # Admin panel assets
│       ├── css/               # Admin stylesheets
│       ├── js/                # Admin JavaScript
│       ├── images/            # Admin images
│       └── vendors/           # Third-party libraries
├── migrations/                  # Database migrations
│   ├── schema.sql              # Complete database schema
│   ├── *_create_*.up.fizz      # Forward migrations
│   ├── *_create_*.down.fizz    # Rollback migrations
│   └── *_seed_*.sql            # Database seeding
├── email-templates/            # Email templates
│   └── basic.html             # Basic email template
├── html-source/               # Static HTML source files
│   ├── index.html             # Homepage source
│   ├── about.html             # About page source
│   ├── contact.html           # Contact page source
│   ├── generals.html          # General's room source
│   ├── majors.html            # Major's room source
│   ├── make-reservation.html  # Reservation form source
│   └── reservation.html       # Reservation page source
├── logs/                      # Application logs
├── go.mod                     # Go module definition
├── go.sum                     # Go module checksums
├── database.yml.example       # Database configuration template
├── run.sh                     # Application startup script
├── update.sh                  # Update script
├── testdbconnection.go        # Database connection test
├── .gitignore                 # Git ignore rules
└── README.md                  # This file
```

## 🚦 Prerequisites

- **Go**: 1.23.4 or later
- **PostgreSQL**: 12.0 or later
- **Buffalo CLI**: For database migrations (optional)

## ⚡ Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/karthikbhandary2/bookings.git
cd bookings
```

### 2. Database Setup

#### Create PostgreSQL Database

```bash
# Create database
createdb bookings

# Create user (optional)
psql -c "CREATE USER bookings_user WITH PASSWORD 'your_password';"
psql -c "GRANT ALL PRIVILEGES ON DATABASE bookings TO bookings_user;"
```

#### Configure Database Connection

Update the database connection string in `cmd/web/main.go`:

```go
db, err := drivers.ConnectSQL("host=localhost port=5432 dbname=bookings user=your_user password=your_password")
```

Or create a `database.yml` file from the example:

```bash
cp database.yml.example database.yml
# Edit database.yml with your database credentials
```

#### Run Database Migrations

```bash
# If you have Buffalo CLI installed
buffalo pop migrate

# Or execute the schema directly
psql -d bookings -f migrations/schema.sql
```

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Configure Application

Update configuration in `cmd/web/main.go`:

```go
// Set to true for production
app.InProduction = false

// Update port if needed
const portNumber = ":8080"
```

### 5. Start the Application

#### Using the run script:
```bash
chmod +x run.sh
./run.sh
```

#### Or manually:
```bash
go build -o bookings cmd/web/*.go
./bookings
```

### 6. Access the Application

- **Frontend**: http://localhost:8080
- **Admin Panel**: http://localhost:8080/admin/dashboard (after login)

## 🔧 Development

### Available Scripts

```bash
# Build and run application
./run.sh

# Update dependencies
./update.sh

# Test database connection
go run testdbconnection.go

# Run tests
go test ./...

# Build for production
go build -o bookings cmd/web/*.go
```

### Database Operations

#### Create New Migration

```bash
# Using Buffalo CLI
buffalo pop generate migration create_new_table

# Manual migration files in migrations/ directory
```

#### Seed Database

```bash
# Run seed files
psql -d bookings -f migrations/seed_rooms_table.postgres.up.sql
psql -d bookings -f migrations/seed_restrictions_table.postgres.up.sql
```

### Application Routes

#### Public Routes
- `GET /` - Homepage with hero carousel
- `GET /about` - About page
- `GET /contact` - Contact page
- `GET /generals-quarters` - General's Quarters room details
- `GET /majors-suite` - Major's Suite room details
- `GET /search-availability` - Room availability search
- `POST /search-availability` - Process availability search
- `POST /search-availability-json` - AJAX availability search
- `GET /choose-room/:id` - Room selection
- `GET /make-reservation` - Reservation form
- `POST /make-reservation` - Process reservation
- `GET /reservation-summary` - Booking confirmation

#### Admin Routes (Protected)
- `GET /admin/dashboard` - Admin dashboard
- `GET /admin/reservations-new` - New reservations
- `GET /admin/reservations-all` - All reservations
- `GET /admin/reservations-calendar` - Calendar view
- `GET /admin/reservations/:src/:id/show` - Reservation details
- `POST /admin/reservations/:src/:id` - Update reservation
- `POST /admin/process-reservation/:src/:id/:action` - Process reservation actions

#### Authentication Routes
- `GET /user/login` - Login page
- `POST /user/login` - Process login
- `POST /user/logout` - Logout

### Data Models

#### Core Models

```go
// User represents system users (admins)
type User struct {
    ID          int
    FirstName   string
    LastName    string
    Email       string
    Password    string
    AccessLevel int
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

// Room represents hotel rooms
type Room struct {
    ID        int
    RoomName  string
    CreatedAt time.Time
    UpdatedAt time.Time
}

// Reservation represents guest bookings
type Reservation struct {
    ID        int
    FirstName string
    LastName  string
    Email     string
    Phone     string
    StartDate time.Time
    EndDate   time.Time
    RoomID    int
    CreatedAt time.Time
    UpdatedAt time.Time
    Room      Room
    Processed int
}

// RoomRestriction represents room availability restrictions
type RoomRestriction struct {
    ID            int
    StartDate     time.Time
    EndDate       time.Time
    RoomID        int
    ReservationID int
    RestrictionID int
    CreatedAt     time.Time
    UpdatedAt     time.Time
    Room          Room
    Reservation   Reservation
    Restriction   Restriction
}
```

### Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/handlers -v
go test ./internal/forms -v
go test ./internal/render -v

# Run tests with race detection
go test -race ./...
```

### Email Configuration

Configure email settings in `cmd/web/send-mail.go`:

```go
// Update SMTP settings
m := mail.NewMSG()
m.SetServer("your-smtp-server.com")
m.SetAuthentication("username", "password")
m.SetPort(587)
m.SetEncryption(mail.EncryptionSTARTTLS)
```

## 🎨 Frontend Features

### User Interface
- **Responsive Design**: Bootstrap 5 responsive grid system
- **Interactive Calendar**: Date picker for reservation dates
- **Form Validation**: Real-time client-side validation
- **Notifications**: Toast notifications for user feedback
- **Modal Dialogs**: SweetAlert2 for confirmations
- **Image Carousel**: Hero image slideshow on homepage

### JavaScript Libraries
- **SweetAlert2**: Beautiful alert dialogs
- **Notie**: Lightweight notifications
- **Vanillajs-datepicker**: Date selection widget
- **Bootstrap 5**: UI framework and components

### Admin Panel Features
- **Dashboard**: Overview of reservations and statistics
- **Calendar View**: Visual representation of bookings
- **Reservation Management**: CRUD operations for bookings
- **User Authentication**: Secure admin login system
- **Responsive Tables**: Mobile-friendly data tables

## 🔒 Security Features

- **CSRF Protection**: NoSurf middleware for CSRF token validation
- **Session Security**: Secure session configuration with SCS
- **Input Validation**: Server-side form validation
- **SQL Injection Protection**: Parameterized database queries
- **Authentication**: Admin-only access to management features
- **Secure Cookies**: HTTPOnly and Secure cookie flags in production

## 🚀 Deployment

### Environment Configuration

For production deployment:

```go
// In cmd/web/main.go
app.InProduction = true

// Update database connection for production
db, err := drivers.ConnectSQL("host=prod-db port=5432 dbname=bookings user=prod_user password=prod_pass sslmode=require")
```

### Production Checklist

1. **Database**: Set up production PostgreSQL instance
2. **Environment**: Set `app.InProduction = true`
3. **SSL**: Configure HTTPS and secure cookies
4. **Email**: Configure production SMTP settings
5. **Logging**: Set up proper log rotation
6. **Monitoring**: Implement health checks and monitoring
7. **Backup**: Set up database backup strategy

### Docker Deployment (Optional)

Create a `Dockerfile`:

```dockerfile
FROM golang:1.23.4-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bookings cmd/web/*.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/bookings .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/email-templates ./email-templates
EXPOSE 8080
CMD ["./bookings"]
```

## 📊 Database Schema

### Tables

- **users**: System administrators
- **rooms**: Available room types
- **reservations**: Guest bookings
- **restrictions**: Restriction types (reservation, owner block)
- **room_restrictions**: Room availability restrictions

### Key Relationships

- Reservations belong to Rooms
- Room Restrictions link Rooms, Reservations, and Restrictions
- Users manage the system through admin interface

## 🆘 Troubleshooting

### Common Issues

1. **Database Connection Failed**
   ```bash
   # Check PostgreSQL status
   sudo systemctl status postgresql
   
   # Verify connection parameters in main.go
   # Check database exists: psql -l
   ```

2. **Template Not Found**
   ```bash
   # Ensure templates directory exists and contains .tmpl files
   # Check template cache configuration in main.go
   ```

3. **Static Files Not Loading**
   ```bash
   # Verify static directory structure
   # Check file permissions: chmod -R 644 static/
   ```

4. **Port Already in Use**
   ```bash
   # Change port in main.go
   const portNumber = ":8081"
   
   # Or kill existing process
   sudo lsof -ti:8080 | xargs kill -9
   ```

5. **Migration Errors**
   ```bash
   # Check database connection
   # Verify migration files syntax
   # Run migrations manually: psql -d bookings -f migrations/schema.sql
   ```

## 📞 Support

For support and questions:
- Check the existing documentation
- Review the code comments in handlers and models
- Test database connection with `testdbconnection.go`
- Examine log files in the `logs/` directory

## 📝 License

This project is part of a learning exercise for Go web development with hotel booking system implementation.

---

**Happy Booking! 🏨**
