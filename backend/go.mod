module jokenpo/backend

go 1.24.1

require (
	github.com/google/uuid v1.6.0
	github.com/gorilla/mux v1.8.1
	jokenpo/database v0.0.0
)

require github.com/lib/pq v1.10.9 // indirect

replace jokenpo/database => ../database
