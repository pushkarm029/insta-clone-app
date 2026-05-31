# Maintenance Policy

This repository is maintained as a portfolio and learning project.

## Current Focus

- Keep the app buildable locally.
- Keep frontend and backend tests runnable from a clean checkout.
- Document setup, Firebase configuration, and known operational constraints.
- Prefer small, reviewable commits for maintenance work.

## Issue Handling

- Bug reports should include reproduction steps, expected behavior, actual behavior, and screenshots when helpful.
- Feature requests should explain the user problem before proposing an implementation.
- Security reports should avoid public proof-of-concept details until a fix is available.

## Dependency Updates

Dependency updates should be tested locally before merging. Avoid broad major-version upgrades unless the migration path is clear and tests cover the affected behavior.

## Release Notes

User-visible maintenance changes should be summarized in `CHANGELOG.md` before tagging a release.
