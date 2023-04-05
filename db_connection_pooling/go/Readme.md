
1. Tutorial

    ```
    go-wrk -c <số goroutine> -d <số giây> <url>
    ```    

    ```
    Usage: go-wrk <options> <url>
    Options:
    -H       Header to add to each request (you can define multiple -H flags) (Default )
    -M       HTTP method (Default GET)
    -T       Socket/request timeout in ms (Default 1000)
    -body    request body string or @filename (Default )
    -c       Number of goroutines to use (concurrent connections) (Default 10)
    -ca      CA file to verify peer against (SSL/TLS) (Default )
    -cert    CA certificate file to verify peer against (SSL/TLS) (Default )
    -d       Duration of test in seconds (Default 10)
    -f       Playback file name (Default <empty>)
    -help    Print help (Default false)
    -host    Host Header (Default )
    -http    Use HTTP/2 (Default true)
    -key     Private key file name (SSL/TLS (Default )
    -no-c    Disable Compression - Prevents sending the "Accept-Encoding: gzip" header (Default false)
    -no-ka   Disable KeepAlive - prevents re-use of TCP connections between different HTTP requests (Default false)
    -no-vr   Skip verifying SSL certificate of the server (Default false)
    -redir   Allow Redirects (Default false)
    -v       Print version details (Default false)
    ```
2. Ref

   https://viblo.asia/p/database-connection-pooling-tong-quan-va-implement-benchmark-38X4EN0oJN2?fbclid=IwAR1ho9LkvFqdnMgFRQIl3WK4I88OtnEBnwDSpx_usMPp2NbtsC3qzNohuk8&mibextid=S66gvF

   https://viblo.asia/p/su-dung-sqldb-dung-cach-de-nang-cao-hieu-nang-cho-ung-dung-1VgZvrx2ZAw