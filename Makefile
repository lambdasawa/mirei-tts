.PHONY: dev deploy run-container

dev:
	go run application.go

deploy:
	git add -f voice/* text-seed.json ipa.dic
	eb setenv GOPATH='/tmp/go'
	eb deploy --staged
	git restore --staged voice/* text-seed.json ipa.dic

run-container:
	docker build -t mirei-tts .
	docker run --rm --interactive --tty --publish 5000 mirei-tts

generate-dictionary:
	./bin/generate-dictionary.sh

archive-data:
	rm -rf data
	mkdir data
	cp -r \
		voice/ \
		tweet.json \
		sentence.json \
		text-seed.json \
		ipa.dic \
		cmd/trim/あいうえお郡道.wav \
		mirei-tts \
		data/
	zip -r data.zip data
