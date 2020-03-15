.PHONY: dev deploy

dev:
	go run application.go

deploy:
	git add -f voice/*
	eb setenv GOPATH='/tmp/go'
	eb deploy --staged
	git restore --staged voice/*

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
