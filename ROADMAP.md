# Roadmap

This roadmap keeps future work explicit without implying a delivery guarantee.

## Maintenance

- Keep local setup instructions current.
- Keep tests runnable without external services where possible.
- Remove committed secrets and document credential rotation.
- Review dependency risk before broad upgrades.

## Product

- Improve responsive layouts across feed, profile, and overlays.
- Add infinite scroll for feed and explore pages.
- Add delete support for posts and comments.
- Improve real-time messaging support with Firebase.
- Consider TypeScript migration after the test surface is stronger.

## Backend

- Split route setup, Firebase initialization, and handlers into smaller units.
- Add handler tests around request validation and error responses.
- Normalize API response shapes.
