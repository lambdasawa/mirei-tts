.PHONY: dev deploy

dev:
	go run application.go

deploy:
	git add -f voice/* text-seed.json
	eb setenv GOPATH='/tmp/go'
	eb deploy --staged
	git restore --staged voice/* text-seed.json
