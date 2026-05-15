# SSE Test Project

This project is a minimal demonstration of Server-Sent Events (SSE) using a Go backend and a simple HTML/JS frontend.

## How to Run

1.  Navigate to the `test` directory:
    ```bash
    cd test
    ```

2.  Run the Go server:
    ```bash
    go run main.go
    ```

3.  The server will start at `http://localhost:8080`.

## How to Test

1.  Open your web browser and go to `http://localhost:8080`.
2.  You should see a page titled "Server-Sent Events Test".
3.  The page will display a new line every second with the current server time, received via SSE.

## Project Structure

- `main.go`: The Go HTTP server with an `/events` endpoint for streaming data.
- `index.html`: The frontend client using `EventSource` to consume the SSE stream.
- `go.mod`: Go module definition.
