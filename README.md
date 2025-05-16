# Project-Base-Training

A minimal Gin + GORM + PostgreSQL app, containerized with Docker Compose.

## 🛠️ Prerequisites

- **Docker** & **Docker Compose** installed on your machine.
- (Optional) A local install of Go 1.24+ if you want to run outside Docker.

## 🚀 Getting Started

1. **Clone the repo**  
   ```bash
   git clone git@github.com:MaTb3aa/Project-Base-Training.git
   cd Project-Base-Training
   ```
2. **Build the Docker images**  
   ```bash
    docker-compose up --build
    ```
   
3. **Try End point**
   ```bash
   curl http://localhost:8888/ping 
   ```
