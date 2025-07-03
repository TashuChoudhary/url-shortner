# 🔗 URL Shortener (Go + Docker + Nginx)

A simple and lightweight URL shortener service built using Go. This backend-only project allows you to shorten long URLs and redirect users to the original destination.

---

## 🚀 Features

- Shortens any valid URL
- Redirects using custom short codes
- Built with Go (Golang)
- Containerized with Docker
- Served via Nginx reverse proxy
- CI/CD pipeline using GitHub Actions
- Deployed on Render

---

## 🌐 Live Demo

[Visit the deployed app](https://url-shortner-3ivd.onrender.com/)

## 🛠️ Tech Stack

| Layer            | Technology       |
|------------------|------------------|
| Backend          | Go (Golang)      |
| Containerization | Docker           |
| Web Server       | Nginx            |
| CI/CD            | GitHub Actions   |
| Deployment       | Render           |

---

## 📁 Project Structure

.
├── main.go # Entry point for the Go application
├── Dockerfile # Docker config to containerize the app
├── nginx.conf # Nginx reverse proxy config
├── .github/workflows/ # CI/CD workflow files
└── README.md # Project documentation


---

## 🧪 How to Run Locally (with Docker)

1. Clone the repo:
   ```bash
   git clone https://github.com/yourusername/url-shortener.git
   cd url-shortener

2. Build and run with Docker compose
   docker compose up --build
   or
   docker-compose up

3. Open in browser
   http://localhost
   
---

## 🔄 CI/CD Pipeline (GitHub Actions)

This project uses GitHub Actions to automate:

- Docker image build
- Linting/Testing (if added)
- Deployment to Render on every `main` branch push

CI/CD file location:  
`.github/workflows/docker-build.yml`

---

## 💡 Future Improvements

- Add a database (PostgreSQL or Redis) for persistent storage

- Write Terraform scripts to deploy infrastructure on AWS/GCP

- Add authentication for managing links

- Add analytics (clicks, location, etc.)

- Create a frontend interface

---

## Author

Tashu Choudhary
Cloud & DevOps Enthusiast
Learning AWS, Docker, Terraform, and Go

- "I believe in building, breaking, and learning fast."









