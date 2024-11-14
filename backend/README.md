https://docs.turso.tech/sdk/go/quickstart

Install libsql for Go
`go get github.com/tursodatabase/libsql-client-go/libsql`

Run Locally without Docker (bad idea, make sure you have tesseract and ocrmypdf installed)

```
go build
./pollis-go
```

Build and Run

```
docker build -t pollis-go
docker run -d -p 8080:8080 pollis-go

# test
curl http://localhost:8080/posts
```

Deploy

```
flyctl launch
```
