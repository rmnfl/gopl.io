# Exercise 1.12

Modify the Lissajous server to read parameter values from the URL. For exam- ple, you might arrange it so that a URL like `http://localhost:8000/?cycles=20` sets the number of cycles to 20 instead of the default 5. Use the `strconv.Atoi` function to convert the string parameter into an integer.You can see its documentation with `go doc strconv.Atoi`

## Usage

### Run the server

```bash
go run main.go &
```

By default, the server listens on:
http://localhost:8000

### Request with custom parameters
You can control the generated figure via query parameters:

| Parameter | Type  | Default | Description                                   |
| --------- | ----- | ------- | --------------------------------------------- |
| `cycles`  | float | 5       | Number of complete x oscillator revolutions   |
| `res`     | float | 0.001   | Angular resolution (smaller → smoother curve) |
| `size`    | int   | 100     | Image canvas covers `[-size..+size]`          |
| `nframes` | int   | 64      | Number of animation frames                    |
| `delay`   | int   | 8       | Delay between frames in 10ms units            |
| `freq`    | float | random  | Relative frequency of y oscillator            |

### Example requests

Default (no parameters):
```
http://localhost:8000/
```

Increase cycles (denser figure):
```
http://localhost:8000/?cycles=20
```

Faster frequency and bigger canvas:
```
http://localhost:8000/?freq=2.5&size=150
```

High resolution and fewer frames:
```
http://localhost:8000/?res=0.0005&nframes=32
```