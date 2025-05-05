# Jokenpo - Real-time Rock Paper Scissors

Jokenpo Realm is a web-based, real-time multiplayer implementation of the classic Rock Paper Scissors game (Jokenpo).

**This project was developed as part of the Cloud Computing course for a Master's Degree program.**

## ✨ Live Demo ✨

A live version of this project is deployed on Microsoft Azure, utilizing Azure Container Apps and Azure Container Registry.

**Access it here:** 👉 **[jokenpo.pt](https://jokenpo.pt)** 👈

*(Note: This domain and the associated Azure hosting are scheduled to be active until approximately July 2025).*

## Architecture Overview 🏗️

The project follows a microservices-like architecture:

1.  **Frontend:** A Vue 3 SPA built with Vite and TypeScript. Handles user interface, routing, and communication with backend services. Uses Pinia for state management and Axios for HTTP requests.
2.  **Backend API (Go):** Handles user authentication (register/login), room creation (HTTP endpoint), and upgrades HTTP connections to WebSocket connections for joining rooms. Uses Gorilla Mux for routing and Gorilla WebSocket for WebSocket handling.
3.  **Statistics API (Go):** A separate service responsible for querying the database and providing game statistics like the leaderboard and move usage percentages. Uses Gorilla Mux for routing.
4.  **Database (PostgreSQL):** Stores user information, game results, and room data.
5.  **WebSocket Communication:** Enables real-time bidirectional communication between players in a game room via the Backend API service.

## Technology Stack 🛠️

* **Frontend:** Vue 3, Vite, TypeScript, Pinia, Axios, Vue Router, CSS
* **Backend:** Go, Gorilla Mux, Gorilla WebSocket, JWT (golang-jwt)
* **Database:** PostgreSQL (Using a SAAS Aiven (Free Tier))
* **Cloud Hosting:** Microsoft Azure (Azure Container Apps)
* **Local Containerization:** Docker, Docker Compose

## Prerequisites (for Local Setup) 📋

* [Git](https://git-scm.com/)
* [Docker](https://www.docker.com/get-started)
* [Docker Compose](https://docs.docker.com/compose/install/) (Usually included with Docker Desktop)
* A `.env` file (see Configuration section)

## Getting Started (Local Setup) 🚀

These instructions are for running the project locally using Docker.

1.  **Clone the Repository:**
    ```bash
    git clone https://github.com/RenatoasMedeiros/jokenpo
    cd jokenpo-main
    (or only cd jokenpo)
    ```

2.  **Configuration (.env file):**
    Create a `.env` file in the root `jokenpo-main` directory by copying the example below, or following the envExample presented on the root of the project. Fill in the required values for *local development*.

    ```dotenv
    # .env - Create this file in the project root for LOCAL setup

    # Backend API Configuration
    BACKEND_ADDR=0.0.0.0
    BACKEND_PORT=8080 # Or your preferred port for the main backend

    # Statistics API Configuration
    STATISTICS_ADDR=0.0.0.0
    STATISTICS_PORT=8081 # Or your preferred port for the statistics service

    # Frontend Configuration
    FRONTEND_PORT=5173 # Default Vite port, change if needed

    # Database Configuration (for connection from backend/statistics services inside Docker)
    DB_HOST=db        # Service name in docker-compose.yml
    DB_PORT=5432      # Default PostgreSQL port
    DB_USER=your_local_db_user     # User for the local Docker DB
    DB_PASSWORD=your_local_db_password # Password for the local Docker DB
    DB_NAME=jokenpo_db
    DB_SSL=disable    # Usually 'disable' for local Docker setup

    # JWT Secret Key (Choose a strong, random secret for local dev)
    JWT_KEY=local_dev_jwt_secret_key

    # Frontend URL (How the backend identifies allowed origins for local dev)
    FRONTEND_URL=http://localhost:5173 # Adjust port if you changed FRONTEND_PORT

    # Vite Environment Variables (Passed to the frontend container for local dev)
    # These URLs are how the *browser* connects to the backend/stats services when running locally
    VITE_BACKEND_URL=http://localhost:8080      # Adjust port if BACKEND_PORT changed
    VITE_STATISTICS_URL=http://localhost:8081  # Adjust port if STATISTICS_PORT changed

3.  **Build and Run with Docker Compose:**
    This command will build the images for the frontend, backend, and statistics services and start them, along with a PostgreSQL database container for local use.
    ```bash
    docker-compose up --build
    ```
    * `--build`: Forces Docker Compose to build the images based on the Dockerfiles.

4.  **Access the Local Application:**
    Open your web browser and navigate to the frontend URL specified in your `.env` file (e.g., `http://localhost:5173`).

5.  **Stop the Local Application:**
    ```bash
    docker-compose down
    ```

## API Endpoints 🌐

* **Auth:**
    * `POST /auth/register`: Create a new player account.
    * `POST /auth/login`: Log in and receive a JWT token.
* **Game Rooms:**
    * `POST /room`: (Authenticated) Creates a new game room, returns the room ID.
    * `GET /join/{room_id}`: Upgrades the HTTP connection to WebSocket to join the specified room.
* **Statistics:**
    * `GET /statistics/ranking`: Returns top players (wins, games played) and move usage percentages.
* **Health Check:**
    * `GET /health`: Simple health check endpoint for the backend service.

## Database Schema 🗄️

The database schema defining the `players`, `games`, and `rooms` tables can be found in `database/schema.sql`.