# Test TLS-Server

Use to test ssl client with different version
by adjusting min max version

```go
TLSConfig: &tls.Config{
	MinVersion: tls.VersionTLS12,
	MaxVersion: tls.VersionTLS12,
}
```

```
curl --insecure -v https://localhost:8888 --tlsv1
curl --insecure -v https://localhost:8888 --tlsv1.2
```