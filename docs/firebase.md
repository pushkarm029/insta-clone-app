# Firebase Setup

The app uses Firebase in two places:

- The React app uses Firebase Authentication, Firestore, Storage, and optional Analytics through the Firebase web SDK.
- The Go backend uses Firebase Admin credentials to access Firestore.

## Frontend Configuration

Create `frontend/.env.local` for local frontend values:

```shell
REACT_APP_FIREBASE_API_KEY=your-api-key
REACT_APP_FIREBASE_AUTH_DOMAIN=your-project.firebaseapp.com
REACT_APP_FIREBASE_PROJECT_ID=your-project-id
REACT_APP_FIREBASE_STORAGE_BUCKET=your-project.appspot.com
REACT_APP_FIREBASE_MESSAGING_SENDER_ID=your-sender-id
REACT_APP_FIREBASE_APP_ID=your-app-id
REACT_APP_FIREBASE_MEASUREMENT_ID=your-measurement-id
```

Firebase web config values are not equivalent to Admin SDK secrets, but local environment files should still stay out of Git.

## Backend Configuration

The backend should use a Firebase Admin service account supplied outside source control. Prefer one of these local forms:

```shell
export FIREBASE_CREDENTIALS_FILE=/absolute/path/to/service-account.json
```

or:

```shell
export FIREBASE_CREDENTIALS_JSON="$(cat /absolute/path/to/service-account.json)"
```

Rotate any service account that was ever committed to a public repository.
