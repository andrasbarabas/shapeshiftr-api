# ShapeShiftr - API

## Debugging in VSCode

### Requirements:

- Visual Studio Code 1.63+
- Go and Delve extensions installed in your VS Code editor

If you haven't installed Delve on your local machine, use the following command to add it:

```
go install github.com/go-delve/delve/cmd/dlv@latest
```

In order to be able to use the debugger properly, you will need a working PostgreSQL/MySQL/etc. database connection, either by spawning a Docker container or starting an instance on your localhost.
If you are working with Docker, it's enough to run the existing `docker-compose.yml` file with the database service only:

```
docker compose up --build db
```

Once ready, you can start the debugger process by pressing `F5` in VSCode.
