.PHONY: deploy

deploy:
	git add -f voice/*
	eb setenv GOPATH='/tmp/go'
	eb deploy --staged
	git restore --staged voice/*

