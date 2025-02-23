# ProjectPulse

ProjectPulse is a modern project management system built with Next.js 15 (App Router) for the frontend and Go for the backend. It provides a robust platform for managing projects, tasks, and team collaboration.

## 🚀 Tech Stack

### Frontend
- Next.js 15 with App Router
- TailwindCSS for styling
- shadcn/ui for UI components
- TypeScript for type safety

### Backend
- Go (Golang)
- Gin web framework
- GORM for database operations
- PostgreSQL for data storage

### Infrastructure
- Docker & Docker Compose
- Multi-container architecture
- Environment-based configuration

## 📁 Project Structure

```
project-pulse/
├── frontend/                # Next.js frontend application
├── backend/                 # Go backend application
│   ├── cmd/                # Main applications
│   ├── internal/           # Internal packages
│   │   ├── models/        # Database models
│   │   ├── handlers/      # HTTP handlers
│   │   └── middleware/    # Middleware functions
│   └── pkg/               # Shared packages
│       ├── config/        # Configuration
│       └── database/      # Database utilities
├── db/                     # Database migrations and seeds
└── docker-compose.yml      # Docker compose configuration
```

## 🔧 Prerequisites

- Docker and Docker Compose
- Node.js 18+ (for local frontend development)
- Go 1.24+ (for local backend development)
- PostgreSQL 15 (handled by Docker)

## 🚀 Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/ProjectPulse.git
   cd ProjectPulse
   ```

2. Create and configure the environment variables:
   ```bash
   cp .env.example .env
   # Edit .env with your configurations
   ```

3. Start the application using Docker Compose:
   ```bash
   docker-compose up -d
   ```

4. Access the application:
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080
   - PostgreSQL: localhost:5432

## 💻 Development

### Running Locally

1. Frontend Development:
   ```bash
   cd frontend
   npm install
   npm run dev
   ```

2. Backend Development:
   ```bash
   cd backend
   go mod download
   go run cmd/main.go
   ```

### Database Migrations

The database schema is automatically migrated when the application starts. Initial migrations can be found in `db/init.sql`.

## 🔒 Environment Variables

Create a `.env` file in the root directory with the following variables:

```env
# Database Configuration
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=

# API Configuration
API_PORT=
GIN_MODE=  # Change to 'release' in production
```

## 🛠️ Features

- User Authentication and Authorization
- Project Management
- Task Tracking
- Team Collaboration
- File Attachments
- Comments and Discussions
- Real-time Updates

## 📝 API Documentation

API documentation is available at `http://localhost:8080/swagger/index.html` when running in development mode.

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Next.js](https://nextjs.org/)
- [Gin Web Framework](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [shadcn/ui](https://ui.shadcn.com/)
- [TailwindCSS](https://tailwindcss.com/)
