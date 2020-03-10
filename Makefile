.PHONY: dev deploy

dev:
	go run application.go

deploy:
	git add -f voice/* text-seed.json ipa.dic
	eb setenv GOPATH='/tmp/go'
	eb deploy --staged
	git restore --staged voice/* text-seed.json ipa.dic

generate-dictionary:
	./bin/generate-dictionary.sh
