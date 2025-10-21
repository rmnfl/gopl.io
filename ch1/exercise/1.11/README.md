# Exercise 1.11

Try fetchall with longer argument lists, such as samples from the top million web sites available at alexa.com. How does the program behave if a web site just doesn’t respond? (Section 8.9 describes mechanisms for coping in such cases.)

# Test run

```bash
➜  1.11 git:(master) ✗ go run main.go http://google.com https://yahoo.com https://alibaba.com https://instagram.com      
0.16s    20169  http://google.com
1.06s       23  https://yahoo.com
1.18s    78648  https://alibaba.com
Get "https://instagram.com": net/http: TLS handshake timeout
10.13s elapsed
```

## Observations

- google.com: quick response, small page (20 169 bytes)
- yahoo.com: small redirect (23 bytes)
- alibaba.com: full HTML content (78 648 bytes)
- instagram.com: failed with TLS handshake timeout

The total elapsed time is dominated by the slowest request — here, Instagram’s unresponsive server.

## Analysis

`fetchall` launches parallel goroutines for each URL.
Each goroutine either completes successfully, or returns an error (e.g., timeout, DNS failure).
The main function waits for all goroutines to finish, so unresponsive servers can significantly delay total execution.