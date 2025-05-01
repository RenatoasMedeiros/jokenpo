module statistics

go 1.24.1

replace jokenpo/database => ../database

require (
	github.com/gorilla/mux v1.8.1
	jokenpo/database v0.0.0-00010101000000-000000000000
)

require github.com/lib/pq v1.10.9 // indirect
