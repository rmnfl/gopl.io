# Exercise 3.4

Following the approach of the Lissajous example in Section 1.7, construct a web server that computes surfaces and writes SVG data to the client. The server must set the Content-Type header like this:

`w.Header().Set("Content-Type", "image/svg+xml")`

(The Lissajous example omitted this because the server used heuristics to recognize common formats such as PNG.)

Allow the client to specify values like `height`, `width`, and `color` as HTTP request parameters.

## Usage

### Run the server

```bash
go run main.go &
```

By default, the server listens on:

http://localhost:8000

### Request with custom parameters

You can control the generated SVG via query parameters:

| Parameter | Type   | Default | Description                              |
| --------- | ------ | ------- | ---------------------------------------- |
| `width`   | int    | 600     | Image canvas width in pixels             |
| `height`  | int    | 320     | Image canvas height in pixels            |
| `cells`   | int    | 100     | Number of grid cells per side            |
| `xyrange` | float  | 30.0    | Axis range for x and y (units)           |
| `color`   | string | `#ffffff` | Fill color for polygons (hex `#RRGGBB`) |

### Example requests

Default (no parameters):

```
http://localhost:8000/
```

Change the canvas size:

```
http://localhost:8000/?width=800&height=600
```

Use a red fill color (note `#` must be URL-encoded as `%23`):

```
http://localhost:8000/?color=%23ff0000
```

Increase detail (more cells):

```
http://localhost:8000/?cells=200
```