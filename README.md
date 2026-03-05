# U-Short

U-Short is a simple and fast **URL Shortener and QR Code Generator** built with **Golang**.  
The application uses **server-side rendering with Go templates** and **HTMX** for dynamic interactions without full page reloads.

![Demo](./demo/u-short.gif)

## Live App
Try it here → [https://u-short.fly.dev](https://u-short.fly.dev)

## Features

- Shorten long URLs into simple links
- Generate QR Codes instantly
- Fast URL redirection
- Server-side rendered pages
- Dynamic UI updates using HTMX


## Tech Stack

- **Go** — backend server  
- **SQLite** — database  
- **HTMX** — dynamic UI updates  
- **TailwindCSS** — styling  


## How It Works

1. User opens the main page.
2. User submits a long URL.
3. The server generates a short code.
4. The URL is stored in the database.
5. The server renders the result using Go templates.
6. A QR Code is generated for easy mobile access.
7. When the short link is visited, the server redirects to the original URL.


## Routes

### Home

```
GET /
```

Render the landing page with the URL shortener and QR code generator.


### Shorten URL

```
POST /shorten
```

Create a shortened URL and render the result.


### Redirect

```
GET /{shortCode}
```

Redirect to the original URL.

Example:

```
http://localhost:3000/abc123
```


## Template Rendering

The application uses Go's built-in **html/template** package.

Templates are located in:

```
web/templates
```

The `layout.html` file acts as the base template.


## Running Locally

Clone the repository:

```
git clone https://github.com/tosrv/u-short.git
cd u-short
```

Install dependencies:

```
go mod tidy
```

Run the server:

```
go run cmd/server/main.go
```

The server will start at:

```
http://localhost:3000
```


## License

MIT License