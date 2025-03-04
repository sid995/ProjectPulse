# ProjectPulse

<div align="center">
  <img src="docs/assets/logo.png" alt="ProjectPulse Logo" width="300" />
  <p><strong>A modern, efficient project management system built with Go and Next.js</strong></p>
</div>

## 📋 Overview

ProjectPulse is a comprehensive project management platform designed to streamline task tracking, team collaboration, and project oversight. With a powerful Go backend and a responsive Next.js frontend, ProjectPulse offers a robust solution for teams of all sizes.

### ✨ Key Features

- **Project Management:** Create, organize, and track projects with customizable workflows
- **Ticket Tracking:** Manage the full lifecycle of tasks and issues
- **Team Collaboration:** Real-time commenting, notifications, and activity tracking
- **User Management:** Role-based access control and detailed user profiles
- **Analytics & Reporting:** Gain insights into project performance and team productivity
- **Modern UI:** Responsive, accessible interface built with React and Tailwind CSS

## 🚀 Getting Started

### Prerequisites

- Docker and Docker Compose
- Git
- Node.js 18+ (for local frontend development)
- Go 1.24 (for local backend development)

### Quick Start with Docker

1. Clone the repository
   ```bash
   git clone https://github.com/sid995/projectpulse.git
   cd projectpulse
   ```

2. Create your environment configuration
   ```bash
   cp .env.example .env
   # Edit .env file with your configuration
   ```

3. Start the application
   ```bash
   docker-compose up
   ```

4. Access the application
   - Frontend: http://localhost:3000
   - API: http://localhost:4000

### Development Setup

For development with hot reloading:

```bash
# Start the development environment
./dev.sh
```

## 🏗️ Architecture

ProjectPulse follows a clean architecture approach with Domain-Driven Design principles:

### Backend (Go)
- **API Layer:** Gin HTTP framework with middleware for authentication, logging, and error handling
- **Service Layer:** Core business logic implementation
- **Repository Layer:** Data access using GORM with PostgreSQL
- **Domain Layer:** Business entities and domain logic
- **Infrastructure:** Redis caching, email services, and external integrations

### Frontend (Next.js + TypeScript)
- **UI Components:** Reusable components built with Tailwind CSS
- **State Management:** React Query for server state and Context API for local state
- **API Integration:** Axios for API communication
- **Authentication:** JWT-based authentication with secure HTTP-only cookies

## 📁 Project Structure

```
projectpulse/
├── api/               # Go backend
│   ├── cmd/           # Application entrypoints
│   ├── internal/      # Private application code
│   ├── pkg/           # Public library code
│   └── Dockerfile     # Backend Docker configuration
├── web/               # Next.js frontend
│   ├── components/    # React components
│   ├── hooks/         # Custom React hooks
│   ├── pages/         # Next.js pages
│   └── Dockerfile     # Frontend Docker configuration
├── docker-compose.yml # Docker Compose configuration
├── .env               # Environment variables
└── README.md          # Project documentation
```

## 🧪 Testing

```bash
# Run backend tests
cd api
go test ./...

# Run frontend tests
cd web
npm test
```

## 📦 Deployment

ProjectPulse can be deployed using Docker Compose for small-scale deployments or Kubernetes for larger installations.

### Production Deployment

```bash
# Build and start production containers
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d
```

## 🛠️ Technology Stack

### Backend
- **Language:** Go
- **Web Framework:** Gin
- **ORM:** GORM
- **Database:** PostgreSQL
- **Caching:** Redis

### Frontend
- **Framework:** React with Next.js
- **Language:** TypeScript
- **Styling:** Tailwind CSS
- **State Management:** React Query + Context API

### DevOps
- **Containerization:** Docker
- **CI/CD:** GitHub Actions
- **Monitoring:** Prometheus + Grafana (planned)

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgements

- [Gin Web Framework](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [Next.js](https://nextjs.org/)
- [Tailwind CSS](https://tailwindcss.com/)
- [React Query](https://react-query.tanstack.com/)

---

<div align="center">
  <p>Made with ❤️ by the ProjectPulse Team</p>
  <p>
    <a href="https://github.com/yourusername/projectpulse/issues">Report Bug</a> ·
    <a href="https://github.com/yourusername/projectpulse/issues">Request Feature</a>
  </p>
</div>