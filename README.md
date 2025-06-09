# Event Management API Project

This project is a backend API designed to handle event creation and user bookings efficiently. Built using Go with the Gin web framework, it leverages SQLite as the database, with GORM as the ORM for smooth and reliable database interactions. The API provides a solid foundation for managing events, enabling features such as event listing, creation, and user reservations.


📘 API Documentation

Base URL: http://localhost:8080

- Public Routes (No Authentication Required)
Method	Endpoint	Description	Controller
GET	/events	Retrieves a list of all events.	controllers.GetEvents
GET	/events/:id	Retrieves details of a specific event.	controllers.GetEvent
POST	/signup	Registers a new user.	controllers.Signup
POST	/login	Authenticates a user and returns a token.	controllers.Login

- Authenticated Routes (Require JWT Token)
Method	Endpoint	Description	Controller
POST	/events	Creates a new event.	controllers.CreateEvent
GET	/events/registered	Lists all events the authenticated user has registered for.	controllers.GetRegisteredEvents
GET	/events/created	Lists all events created by the authenticated user.	controllers.GetCreatedEvents
PUT	/events/:id	Updates an existing event by ID.	controllers.UpdateEvent
DELETE	/events/:id	Deletes an event by ID.	controllers.DeleteEvent
POST	/events/:id/register	Registers the user for the specified event.	controllers.RegisterForEvent
DELETE	/events/:id/register	Cancels the user's registration for the specified event.	controllers.CancelRegistration
🚀 Running the Project Locally

To run this project on your local machine, follow the steps below:
- Prerequisites

    Go installed (v1.18 or later recommended)

    Git installed

    (Optional) make if using Makefile for automation

🛠️ Setup Steps

# 1. Clone the repository
git clone https://github.com/your-username/event-management-api.git
cd event-management-api

# 2. Download dependencies
go mod tidy

# 3. Run the API
go run main.go

The API will start on http://localhost:8080

    🗒️ Note: The project uses SQLite as the database. On first run, the database file will be created automatically.
