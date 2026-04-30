# MVP Plan - Digital Proof Approval Workflow

## 1. Project Title

**Digital Proof Approval Workflow MVP**

## 2. Short Summary

This MVP is a web-based workflow for digital proof approval.

It helps replace the current manual email-based process with a clearer system where:

1. A file checker or production location creates/uploads a digital proof.
2. The customer receives a review link.
3. The customer opens the proof in the web app.
4. The customer approves the proof or requests changes.
5. The file checker can see the updated proof status.

The MVP is not a final production system. It is a working prototype to validate the main workflow and prepare the internship assignment scope.

---

## 3. Problem

Digital proofs are currently handled manually by emailing back and forth between:

- the customer
- Print.com
- the production location

This can create problems such as:

- unclear proof status
- difficult tracking of customer approval
- possible confusion about the latest proof version
- manual communication overhead
- slower production workflow

The goal of this MVP is to make the proof approval process more structured, visible, and easier to manage.

---

## 4. Main Goal

The main goal is to create a simple proof approval workflow where customers can check and confirm digital proofs inside a web application.

The MVP should make the workflow easier for both the customer and the internal file checker.

---

## 5. Users / Roles

### Customer

The customer can:

- view proofs waiting for review
- open a proof review page
- approve a digital proof
- request changes with a comment
- see the current status of a proof

### File Checker / Internal Employee

The file checker can:

- view all proof requests
- create a new proof request
- link a proof to an order
- generate a customer review link
- see whether the proof is approved or needs changes

### Production Location

In the real assignment, production locations may also upload digital proofs.

For this MVP, this role is represented by the file checker/internal upload flow.

---

## 6. MVP Scope

### Included in this MVP

- Customer dashboard
- Digital proof review page
- File checker dashboard
- Create/upload proof form
- Fake email preview with customer review link
- Approve proof action
- Request changes action with comment
- Status tracking
- Print.com-inspired UI style
- Vue frontend
- Go backend API
- In-memory storage

### Not included yet

- Real authentication
- Real user roles and permissions
- Real email sending
- Real PDF/image upload
- Real PDF/image preview
- Database storage
- Proof version history
- Production-ready security

---

## 7. Main Workflow

### Step 1 — Create Proof

A file checker creates a new proof request by entering:

- order ID
- customer name
- customer email
- product name
- proof file name

The backend creates a new proof with the status:

```text
WAITING_FOR_APPROVAL
```

### Step 2 — Fake Email Preview

After the proof is created, the system shows a fake email preview.

The email contains:

- customer name
- order ID
- product name
- proof file name
- customer review link

This demonstrates how the customer would receive the proof in a real workflow.

### Step 3 — Customer Review

The customer opens the review link and sees:

- proof information
- order information
- customer information
- proof preview placeholder
- approve button
- request changes form

### Step 4 — Customer Decision

The customer can choose one of two actions:

1. **Approve proof**
2. **Request changes**

If approved, the status becomes:

```text
APPROVED
```

If changes are requested, the status becomes:

```text
CHANGES_REQUESTED
```

A comment is saved together with the proof.

### Step 5 — File Checker Status Overview

The file checker can open the dashboard and see all proofs with their latest statuses.

---

## 8. Proof Statuses

The MVP uses three proof statuses:

```text
WAITING_FOR_APPROVAL
APPROVED
CHANGES_REQUESTED
```

### Status Meaning

| Status | Meaning |
|---|---|
| WAITING_FOR_APPROVAL | The proof has been created and is waiting for the customer |
| APPROVED | The customer approved the proof |
| CHANGES_REQUESTED | The customer requested changes and added a comment |

---

## 9. Pages

### Customer Dashboard

URL:

```text
/customer/dashboard
```

Purpose:

- show all customer proofs
- show proof status
- open proof review page

### Proof Review Page

URL:

```text
/customer/proofs/:id
```

Purpose:

- show proof details
- show preview placeholder
- approve proof
- request changes with comment

### File Checker Dashboard

URL:

```text
/file-checker/dashboard
```

Purpose:

- show all digital proofs
- show proof statuses
- open customer review page
- navigate to upload proof page

### Upload Proof Page

URL:

```text
/file-checker/upload
```

Purpose:

- create a new proof request
- generate customer review link
- show fake email preview

---

## 10. Backend API Endpoints

### Health Check

```http
GET /api/health
```

Purpose:

Check if the backend is running.

### Get All Proofs

```http
GET /api/proofs
```

Purpose:

Get all digital proofs.

### Get One Proof

```http
GET /api/proofs/:id
```

Purpose:

Get one proof by ID.

### Create Proof

```http
POST /api/proofs
```

Example body:

```json
{
  "orderId": "ORD-1003",
  "customerName": "Sarah Miller",
  "customerEmail": "sarah@example.com",
  "productName": "Poster A3",
  "fileName": "poster-a3-proof.pdf"
}
```

Purpose:

Create a new proof request.

### Approve Proof

```http
POST /api/proofs/:id/approve
```

Purpose:

Change proof status to approved.

### Request Changes

```http
POST /api/proofs/:id/request-changes
```

Example body:

```json
{
  "comment": "Please make the logo bigger."
}
```

Purpose:

Save customer feedback and change proof status to changes requested.

---

## 11. Data Model

### Proof

```text
id
orderId
customerName
customerEmail
productName
fileName
fileUrl
status
comment
createdAt
updatedAt
```

### Example Proof

```json
{
  "id": 1,
  "orderId": "ORD-1001",
  "customerName": "Kiarash Delavar",
  "customerEmail": "customer@example.com",
  "productName": "Business Cards",
  "fileName": "business-card-proof.pdf",
  "fileUrl": "/uploads/business-card-proof.pdf",
  "status": "WAITING_FOR_APPROVAL",
  "comment": "",
  "createdAt": "2026-04-29T20:28:10Z",
  "updatedAt": "2026-04-29T20:28:10Z"
}
```

---

## 12. Tech Stack

### Frontend

- Vue.js
- Vue Router
- Axios
- CSS
- Vite

### Backend

- Go
- Gin framework
- CORS middleware
- In-memory storage

### Deployment

- GitHub repository
- GitHub Pages for frontend deployment
- GitHub Actions workflow for frontend build/deploy

---

## 13. Current Limitation

The deployed GitHub Pages version only hosts the frontend.

The Go backend does not run on GitHub Pages.

For full functionality, the backend must run locally:

```bash
cd backend
go run main.go
```

Then the frontend can call:

```text
http://localhost:8080/api
```

In a future version, the backend could be deployed separately using Render, Railway, Fly.io, or another hosting provider.

---

## 14. Future Improvements

### Real File Upload

Replace the current file-name-only input with real PDF/image upload.

### PDF/Image Preview

Show the uploaded proof directly inside the proof review page.

### Real Email Sending

Send an actual email to the customer when a proof is ready.

### Authentication and Roles

Add login and role-based access for:

- customer
- file checker
- production location
- admin

### Database Storage

Replace in-memory storage with a real database such as:

- SQLite
- PostgreSQL
- MySQL

### Proof Version History

Allow multiple proof versions after customer feedback.

Example:

```text
Version 1 → customer requested changes
Version 2 → customer approved
```

### Audit Log

Track actions such as:

- proof created
- email sent
- proof viewed
- proof approved
- changes requested

### Reminder System

Send reminders when a proof is waiting too long for customer approval.

---

## 15. Success Criteria for the MVP

The MVP is successful if:

- a file checker can create a proof request
- a customer can open the proof review page
- a customer can approve the proof
- a customer can request changes with a comment
- the file checker can see the updated proof status
- the workflow is clear and easy to demonstrate

---

## 16. Demo Flow

1. Start the backend:

```bash
cd backend
go run main.go
```

2. Start the frontend:

```bash
cd frontend
npm run dev
```

3. Open the file checker dashboard:

```text
http://localhost:5173/file-checker/dashboard
```

4. Open the upload proof page:

```text
http://localhost:5173/file-checker/upload
```

5. Create a new proof request.

6. Show the fake email preview.

7. Open the customer review page.

8. Approve the proof or request changes.

9. Return to the file checker dashboard.

10. Check that the status has changed.

---

## 17. Summary

This MVP demonstrates a realistic full-stack workflow for digital proof approval. It shows how a manual email-based process can be turned into a structured web application with:

- proof creation
- customer review
- approval status tracking
- fake email preview
- customer and internal user flows

