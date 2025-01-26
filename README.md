# Chat Room Application

This is a **real-time chat room application** built with a Go backend (WebSocket server) and a React frontend. The Go server handles WebSocket connections and facilitates real-time communication between clients, while the React frontend provides the user interface.

## Project Structure

```bash
.
├── backend/    # Go (Golang) WebSocket server (backend)
├── frontend/          # React frontend for the chat room
└── README.md          # Project documentation
```
## Prerequisites
- Go: version 1.16 or higher.
- Node.js: Node v14 or higher recommended


## Setting Up the React Frontend (frontend)
The React frontend serves as the user interface for the chat room.

### Steps to Run the Frontend
1. **Navigate to the frontend/ directory**:
```bash
cd frontend/
```
2. **Install dependencies**:
```bash
npm install
```
3. **Start the development server**:
```bash
npm start
```
The frontend will be available at http://localhost:3000.

## Frontend File Structure
```bash
frontend/
├── public/
│   └── index.html    # Main HTML file for the React app
├── src/
│   ├── components/   # Contains React components (frontend, etc.)
│   ├── App.tsx       # Main app component
│   ├── index.tsx     # Entry point for the React app
│   └── styles.css    # Custom styles for the chat room
├── package.json      # NPM package file
└── tsconfig.json     # TypeScript configuration file
```

## Key Features
- Real-time messaging using WebSockets.
- Emoji support using the emoji-mart library.
- Responsive frontend layout.

## Running Both Frontend and Backend

1. **Start the Go server: Open a terminal and navigate to the backend folder**:
```bash
cd backend/
go run main.go
```

2. **Start the React frontend: In a separate terminal, navigate to the frontend folder**:
```bash
cd frontend/
npm start
```