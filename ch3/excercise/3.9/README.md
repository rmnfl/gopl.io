# Exercise 3.9

Write a web server that renders fractals and writes image data to the client. Allow the client to specify the x, y, and zoom values as parameters to the HTTP request.

## Usage

### Run the server

```bash
go run main.go &
```

By default, the server listens on:

http://localhost:8000

### Request with custom parameters

You can control the generated SVG via query parameters:

| Parameter | Type    | Default | Description                               |
| --------- | ------- | ------- | ----------------------------------------- |
| `x`       | float64 | 0       |  Center point X coordinate                |
| `y`       | float64 | 0       |  Center point Y coordinate                |
| `zoom`    | float64 | 1       | Zoom ratio                                |
