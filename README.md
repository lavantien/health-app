# Simple Health App

## System Overview

### Requirements Breakdown

- Time: 12-16 hours
- No Authentication and Authorization of Users
- No unit tests
- Static Landing Page without login
- Input Pages
- Two Read-only Pages with buttons to redirect to their appropriate Input Pages
- 5 Endpoints: Users, Meals, Records, Exercises, Diaries
- Each endpoint only needs Create and Read All actions

### Business Decisions

- Single Rest API with Mutex and Caching, without external libraries; exposes multiple endpoints
- Simple backend database layer to communicate with external database
- Simple SPA frontend with routings consume the backend's endpoints and render them appropriately according to the Figma designs
- All functionalities and endpoints should be tested

### Technical Decisions

- Dev Env: Go, Node
- Database: In-memory
- Communication: HTTP1.1 (REST Endpoints)
- Backend: Go
- Frontend: React

### Demo Instruction

- Install Go and NodeJS
- Create an optimized build of the React client for Go server to serve:

```bash
cd client && npm run build
```

- Run the Go server, which will serve both backend and frontend:

```bash
cd .. && go run .
```

- Go to `<http://localhost:4200>' to visit the website

### Development Environment

- Install Go and Node
- Live editing React client by first go to the `client` directory and run `npm start` for hot reload, then go to `http://localhost:3000`
- After make changes to the React client, go to the `client` directory and run `npm run build` to rebuild the client
- Run Go server as normal, `go run .`

## Design Documents

### Business Logics

#### Users

- When creating User, will check if the name already existed
- There currently no Auth for User endpoints
- There currently no password encoding

### Database Schemas

### CURL Endpoints Documentation
