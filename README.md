# Digital Proof Approval Workflow MVP

This project is an MVP for an automated digital proof approval workflow inspired by a professional print platform workflow.

## Goal

The goal is to replace a manual email-based digital proof approval process with a clear web-based workflow.

## Users

- Customer
- File checker / internal employee

## MVP Features

- Upload a digital proof
- Link proof to an order
- Show customer dashboard
- Preview proof file
- Approve proof
- Request changes with a comment
- Track proof status

## Tech Stack

- Vue.js frontend
- Go backend
- Gin API framework
- In-memory storage for the first MVP version

## Planned Workflow

1. File checker uploads a proof.
2. The proof is linked to an order.
3. Customer sees the proof in their dashboard.
4. Customer opens the proof review page.
5. Customer approves or requests changes.
6. File checker sees the updated status.
