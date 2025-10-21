# Exercise 1.10

Find a web site that produces a large amount of data. Investigate caching by running `fetchall` twice in succession to see whether the reported time changes much. Do you get the same content each time? Modify `fetchall` to print its output to a file so it can be examined.

## Test run


```bash
➜  1.10 git:(master) ✗ go run main.go https://alibaba.com                  
2.68s    78647  https://alibaba.com
2.68s elapsed
➜  1.10 git:(master) ✗ go run main.go https://alibaba.com
1.53s    78647  https://alibaba.com
1.53s elapsed
➜  1.10 git:(master) ✗ go run main.go https://alibaba.com
1.43s    78647  https://alibaba.com
1.43s elapsed
```

Two consecutive requests to https://alibaba.com produced nearly identical HTML files (output1.html, output2.html).
A diff comparison shows only minor differences in dynamic tracking parameters and timestamps:
```bash
➜  1.10 git:(master) ✗ diff output1.html output2.html 
137c137
< window.server_aplus=true;with(document)with(body)with(insertBefore(createElement("script"),firstChild))setAttribute("exparams","userid=&aplus&ug_tag=&eagleeye_traceid=21613e1617609180705088033e0dda&ip=95%2e174%2e91%2e224&dmtrack_c={}&pageid=5fae5be00b884eb41760918071&hn=haumea011136078180%2erg%2dde%2d3%2ede198&asid=AQAAAAA3evVoaOfaUwAAAABCgY/kEbpl6g==&at_bu=icbu",id="beacon-aplus",src="//s.alicdn.com/@g/alilog/??aplus_plugin_icbufront/index.js,mlog/aplus_v2.js")
---
> window.server_aplus=true;with(document)with(body)with(insertBefore(createElement("script"),firstChild))setAttribute("exparams","userid=&aplus&ug_tag=&eagleeye_traceid=2161386017609180824973516e0e5f&ip=95%2e174%2e91%2e224&dmtrack_c={}&pageid=5fae5be00b884a1c1760918083&hn=haumea011136074028%2erg%2dde%2d3%2ede198&asid=AQAAAABDevVotgZQNAAAAABpmdf3S55NJw==&at_bu=icbu",id="beacon-aplus",src="//s.alicdn.com/@g/alilog/??aplus_plugin_icbufront/index.js,mlog/aplus_v2.js")
149c149
< <div class="home-container-newuser"  id="root"><!-- Silkworm Render: 0b884eb417609178485321992d0c99 --><style>
---
> <div class="home-container-newuser"  id="root"><!-- Silkworm Render: 0b884a1c17609179918198718d0bc2 --><style>
846c846
< <!--1760917848552 Sun Oct 19 16:50:48 PDT 2025-->
---
> <!--1760917992022 Sun Oct 19 16:53:12 PDT 2025-->
```
These differences correspond to session identifiers and rendering timestamps, which are generated dynamically.
The main page structure and content are identical.

## Results

- The first request took 2.68 seconds, while subsequent runs were significantly faster (≈1.5 seconds).
- The response size (78 647 bytes) was identical across all runs.
- This indicates that HTTP caching (either browser-level or CDN-level) likely reduced network latency on subsequent requests.

Repeated requests to the same URL can be served faster due to caching mechanisms, even though the content remains almost the same.