# Digital Proof Approval Workflow MVP

This project is a small MVP for an automated digital proof approval workflow, inspired by a professional print platform workflow.

The goal of the prototype is to explore how a manual email-based proof approval process can be replaced by a clearer web-based workflow.

## Goal

Digital proofs are usually images or PDF files that need explicit approval from a customer before production can continue.

In the current assignment context, this process is manually handled by emailing back and forth between the customer, Print.com, and the production location.

This MVP demonstrates a simple workflow where:

1. A file checker creates a proof request  
2. A customer receives a review link  
3. The customer checks the proof in the web app  
4. The customer approves the proof or requests changes  
5. The file checker can see the updated proof status  

## Users

### Customer

The customer can:

- View proofs in a customer dashboard  
- Open a proof review page  
- Approve a proof  
- Request changes with a comment  

### File Checker / Internal Employee

The file checker can:

- View all proof requests  
- Create a new proof request  
- Generate a customer review link  
- See the current status of each proof  

## MVP Features

- Customer dashboard  
- File checker dashboard  
- Upload/create proof form  
- Digital proof review page  
- Fake email preview with customer review link  
- Approve proof action  
- Request changes with a comment  
- Proof status tracking  
- Print.com-inspired user interface  

## Tech Stack

- Vue.js frontend  
- Vue Router  
- Axios  
- Go backend  
- Gin API framework  
- In-memory storage (MVP only)  
- GitHub Pages (frontend deployment)  

## Demo Flow

1. Open the file checker dashboard  
2. Create a new proof request  
3. View the generated fake email preview  
4. Open the customer review link  
5. Approve the proof or request changes  
6. Return to the file checker dashboard to see the updated status  

## How to Run Locally

### Backend

```bash
cd backend
go run main.go
```

Backend runs on:
`   http://localhost:8080   `

### Frontend

`   cd frontend npm install npm run dev   `

Frontend runs on:
`   http://localhost:5173   `

Main Pages
----------

Customer dashboard:
`   /customer/dashboard   `

File checker dashboard:
`   /file-checker/dashboard   `

Upload proof page:
`   /file-checker/upload   `

Proof review page:
`   /customer/proofs/:id   `

Current Limitations

This is only an MVP prototype. Some parts are simplified:

*   Backend uses in-memory storage
    
*   Data resets when backend restarts
    
*   No real file upload yet
    
*   Proof preview is a placeholder
    
*   No real email sending
    
*   No authentication or roles
    
*   GitHub Pages hosts only the frontend
    

