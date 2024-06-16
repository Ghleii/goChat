module github.com/Ghleii/goChat

go 1.18

require (
	github.com/Ghleii/goChat/handlers 07d3eaf2c488e2f5e46ef1786352a2091a68ef2f
	github.com/go-sql-driver/mysql v1.6.0
)

require (
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e // indirect
)

replace github.com/Ghleii/goChat/handlers => ./handlers
