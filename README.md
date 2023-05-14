# GOFS: Go-Fullstack

CLI to generate fullstack Golang projects all-in-one.

Currently supports a plain Go backend w/ vite (react-ts) client.

## Usage

```
go build

./gofs create -name PROJ_NAME -dest /home/user/projects/ -type react -username GITHUB_USER
```

### Todo

- [ ] Add more frontend frameworks (svelte, vue, angular)
- [ ] Initialize Gin framework?
- [ ] Boilerplate code connecting frontend + backend
- [ ] Better path handling for -dest flag
- [ ] Upload JSON config? 
- [ ] Better project structure (CLI)
