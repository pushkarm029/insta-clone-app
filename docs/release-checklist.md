# Release Checklist

Use this checklist before tagging a maintenance release.

- [ ] Confirm `git status --short` is clean.
- [ ] Run frontend tests.
- [ ] Run frontend build.
- [ ] Run backend tests.
- [ ] Run `gofmt -w .` in `backend/`.
- [ ] Review `git diff origin/master...HEAD`.
- [ ] Update `CHANGELOG.md`.
- [ ] Confirm no secrets are staged.
- [ ] Create an annotated tag.

Example:

```shell
git tag -a v0.2.0-maintenance-refresh -m "2026 maintenance refresh"
```
