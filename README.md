[![go]()]()
[![Fiber]()]()
[![Prisma]()]()

# Go Fiber Backend

This is a simple implementation of [GoFiber]() with [Prisma ORM]()

## Install Modules

```
go mod init main
```

```
go mod tidy
```

## Database Configuration

```
go run github.com/steebchen/prisma-client-go db push
```

## Usage

Runing as project:

```
go run main.go
```

```
 ┌───────────────────────────────────────────────────┐
 │                      GoBack                       │
 │                   Fiber v2.52.4                   │
 │                http://127.0.0.1:80                │
 │        (bound on host 0.0.0.0 and port 80)        │
 │                                                   │
 │ Handlers ............. 8  Processes ........... 1 │
 │ Prefork ....... Disabled  PID ............. 15744 │
 └───────────────────────────────────────────────────┘
```

If you considere add values edit the **schema.prisma** and **controllers/register.go** file and push the changes using:

```
go run github.com/steebchen/prisma-client-go db push
```

## Structure of project

```
> controllers
  > home.go |> get all user information using token
  > login.go |> username or email using a identifier
  > logout.go |> logout using a token
  > main.go |>  api test directory
  > register.go |> register directory
```

## API Routing

```
> Main Routes
  > http://localhost/*
```

```
> Account Routes
    > http://localhost/account/home
    > http://localhost/account/login
    > http://localhost/account/logout
    > http://localhost/account/register
```

# Important

> 🛑 This project is for experimental use only is not recomended for a profesional use in a production project !
