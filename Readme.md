ChatApp API Documentation

Overview
ChatApp is a backend service written in Go, deployed on Railway.
It provides APIs for authentication, project management, proposals, and real-time chat channels.
This document describes the available endpoints and how to test them using Postman or curl.

Base URL
https://chatapp-production-6ce1.up.railway.app
All endpoints are relative to this base URL.

Environment Variables
- baseUrl: API hostname (default: https://chatapp-production-6ce1.up.railway.app)
- jwt: Bearer token returned after login
- projectId: ID of a project (example: 123)
- channelId: ID of a chat channel (example: proposal-123)

Endpoints

1. Auth
Signup:
POST /signup
Body:
{
  "email": "test@example.com",
  "password": "secret123"
}

Login:
POST /login
Body:
{
  "email": "test@example.com",
  "password": "secret123"
}
Returns JWT token to use in Authorization header.

2. Projects
Create Project:
POST /projects
Headers:
Authorization: Bearer {jwt}
Body:
{
  "title": "Website Build",
  "description": "Need a Go backend"
}

Submit Proposal:
POST /projects/{projectId}/proposals
Headers:
Authorization: Bearer {jwt}
Body:
{
  "freelancer_id": "freelancer123",
  "bid": 5000
}

3. Chat
Create Channel:
POST /chat/{channelId}/create
Headers:
Authorization: Bearer {jwt}
Body:
{
  "creator_id": "client123",
  "members": ["client123", "freelancer456"]
}

Send Message:
POST /chat/{channelId}/message
Headers:
Authorization: Bearer {jwt}
Body:
{
  "user_id": "client123",
  "text": "Hello, welcome!"
}

Add Members:
POST /chat/{channelId}/members/add
Headers:
Authorization: Bearer {jwt}
Body:
{
  "members": ["reviewer789"]
}

Remove Members:
POST /chat/{channelId}/members/remove
Headers:
Authorization: Bearer {jwt}
Body:
{
  "members": ["reviewer789"]
}

Quickstart with Postman
1. Import the provided JSON collection into Postman.
2. Set baseUrl to your Railway deployment URL.
3. Run Signup then Login to get a JWT.
4. Save JWT in the jwt variable.
5. Test project and chat endpoints using the collection.

Summary
- Base URL: https://chatapp-production-6ce1.up.railway.app
- Auth: Signup/Login
- Projects: Create, Submit Proposal
- Chat: Create Channel, Send Message, Add/Remove Members
- Use JWT in Authorization header for protected routes.
