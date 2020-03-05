.PHONY: dev deploy

dev:
	go run application.go

deploy:
	git add -f voice/*
	eb setenv GOPATH='/tmp/go'
	eb deploy --staged
	git restore --staged voice/*

