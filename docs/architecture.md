# Architecture Overview

Insta-Clone-App is split into a React frontend and a Go backend API.

## Frontend

The frontend lives in `frontend/` and is built with Create React App, React Router, Redux Toolkit, Firebase client SDKs, and Jest tests. It owns the user-facing flows: sign in, account creation, feed, explore, search, reels, messages, post creation, profile, and chill-zone.

## Backend

The backend lives in `backend/` and exposes REST-style endpoints with Gin. Handlers read and write Firebase Firestore data for users, posts, likes, comments, followers, and feeds.

## Data Services

Firebase Authentication handles frontend user auth. Firestore stores user and post documents. Firebase Storage hosts uploaded media.

## Local Runtime

For local development, run the frontend on port `3000` and the backend on port `8080`. The frontend development server proxies `/api` requests to the backend.
