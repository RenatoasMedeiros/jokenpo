module jokenpo/backend

go 1.24.1

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/uuid v1.6.0
	github.com/gorilla/mux v1.8.1
	golang.org/x/crypto v0.36.0
	jokenpo/database v0.0.0-00010101000000-000000000000
)

require (
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/lib/pq v1.10.9 // indirect
)

replace jokenpo/database => ../database
