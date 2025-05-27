## Event Managemet API Project

- this project entails build an API tho manage event creation and bookings by users. It uses Go with gin package and SQLite powered database with GORM for effective database communication.



### API Documentation

This document describes the endpoints exposed by the Go API using the Gin framework.

Base URL; `http://localhost:8080`

Public Routes (No Authentication Required)
List All Events

    Method: GET

    Path: /events

    Description:
    Retrieves a list of all events.

    Controller: controllers.GetEvents

Get Event by ID

    Method: GET

    Path: /events/:id

    Description:
    Retrieves details of a specific event by its ID.

    Controller: controllers.GetEvent

User Signup

    Method: POST

    Path: /signup

    Description:
    Registers a new user.

    Controller: controllers.Signup

User Login

    Method: POST

    Path: /login

    Description:
    Authenticates a user and returns a token.

    Controller: controllers.Login

Authenticated Routes

All routes below require authentication.

Create Event

    Method: POST

    Path: /events

    Description:
    Creates a new event.

    Controller: controllers.CreateEvent

List Registered Events

    Method: GET

    Path: /events/registered

    Description:
    Lists all events the authenticated user has registered for.

    Controller: controllers.GetRegisteredEvents

List Created Events

    Method: GET

    Path: /events/created

    Description:
    Lists all events created by the authenticated user.

    Controller: controllers.GetCreatedEvents

Update Event

    Method: PUT

    Path: /events/:id

    Description:
    Updates an existing event by ID.

    Controller: controllers.UpdateEvent

Delete Event

    Method: DELETE

    Path: /events/:id

    Description:
    Deletes an event by ID.

    Controller: controllers.DeleteEvent

Register for Event

    Method: POST

    Path: /events/:id/register

    Description:
    Registers the authenticated user for the specified event.

    Controller: controllers.RegisterForEvent

Cancel Registration

    Method: DELETE

    Path: /events/:id/register

    Description:
    Cancels the authenticated user’s registration for the specified event.

    Controller: controllers.CancelRegistration
