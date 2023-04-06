# Simple Health App

## System Overview

### Requirements Breakdown

- Time: 12-16 hours
- Authentication and Authorization of Users
- Static Landing Page without login
- Input Pages
- Two Read-only Pages with buttons to redirect to their appropriate Input Pages
- 4 Endpoints: Meals, Records, Exercises, Diaries
- Each endpoint only needs Create and Read All actions

### Business Decisions

- Single Rest API with Mutex and Caching, without external libraries; exposes multiple endpoints
- Simple backend database layer to communicate with external database
- Simple SPA frontend with routings consume the backend's endpoints and render them appropriately according to the Figma designs
- All functionalities and endpoints should be tested

### Technical Decisions

- CI/CD: Git, GitHub Actions
- Dev Env & Deployment: Docker, Podman, Docker Compose, AWS EC2
- Database: MongoDB
- Communication: HTTP1.1 (REST Endpoints)
- Backend: Go
- Frontend: Remix React
- Migration Service: Go

### Production Deployment

- Link: [REDACTED]

### Local Development Environment Setup

- Install Docker and Docker Compose
- Run docker-compose up

## Design Documents

### Business Logics

### Database Schemas

### CURL Endpoints Documentation
