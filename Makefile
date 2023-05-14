
example-text:	
	cd example/text && go run main.go

example-template:	
	cd example/template && go run main.go


format:
	goimports-reviser -rm-unused -use-cache -set-alias -format ./...

test:
	go test -v ./...