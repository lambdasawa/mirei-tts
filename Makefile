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

# Save data used by application to Amazon S3.
# Usage: make save-data BUCKET=mireittsstack-databucketxxxxx
save-data:
	aws s3 cp --recursive ./voice s3://${BUCKET}/voice/
	aws s3 cp ./text-seed.json s3://${BUCKET}/text-seed.json
	aws s3 cp ./ipa.dic s3://${BUCKET}/ipa.dic
