all: render

build:
	go build .

render: build
	./rda render
