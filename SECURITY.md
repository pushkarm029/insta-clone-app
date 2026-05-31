# Security Policy

## Reporting

Please report security issues privately to the repository maintainer. Do not open a public issue with exploit details, credentials, tokens, or personal data.

## Supported Scope

The current supported scope is the latest commit on the active maintenance branch and the latest tagged maintenance release.

## Credential Handling

- Do not commit `.env` files, Firebase Admin service account JSON, private keys, or deployment secrets.
- Use `FIREBASE_CREDENTIALS_FILE` or `FIREBASE_CREDENTIALS_JSON` for the backend.
- Rotate any credential that was committed or shared publicly.

## Local Review Checklist

- Run frontend tests.
- Run backend tests.
- Confirm no secrets are staged with `git diff --cached`.
- Check changed dependencies before committing lockfile updates.
