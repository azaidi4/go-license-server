
$(info Starting Postgres server...)
$(shell docker run --rm --name license-pg -d -p 12345:5432 postgres)
$(info Installing Go packages...)

all: 
	go get github.com/lib/pq
	go get golang.org/x/tools/cmd/goimports
	go get -u github.com/xo/xo
	go get github.com/gorilla/mux
	go get github.com/gorilla/handlers
