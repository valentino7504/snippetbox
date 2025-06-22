# Snippetbox

**Snippetbox** is a minimalist web application for creating, viewing, and listing plain-text snippets.  
It is designed to be simple, secure, and easy to understand — ideal for learning the fundamentals of web development in Go.

This project is based on the book  
_Let's Go: Learn to build professional web applications with Go_  
by **Alex Edwards**  
[https://lets-go.alexedwards.net](https://lets-go.alexedwards.net)

---

## Purpose

The goal of Snippetbox is to help developers understand how to build a full-featured web application in Go using only the standard library and minimal dependencies. It emphasizes clear architecture, proper separation of concerns, and secure web development practices.

This is an educational project, not meant for production use.

---

## Features

- Create and display short text snippets
- View individual snippets by unique ID
- Display a list of recent snippets
- HTML templating with dynamic data injection
- Environment-based configuration (via flags)
- Clean project structure with testable, modular components

---

## Tech Stack

- **Language**: Go (Golang)
- **Standard Library Packages**:
  - `net/http` – HTTP server and routing
  - `html/template` – Templating engine
  - `database/sql` – Database access
- **Database**: MariaDB
- **Logger**: `log/slog` for structured logging
- **External Libraries**:
  - `github.com/go-sql-driver/mysql` – MySQL driver
  - Other minimal packages as needed

---

## Credits

This project is a direct result of working through  
_Let's Go: Learn to build professional web applications with Go_  
by Alex Edwards  
[https://lets-go.alexedwards.net](https://lets-go.alexedwards.net)

All design and architectural patterns are drawn from that resource.
