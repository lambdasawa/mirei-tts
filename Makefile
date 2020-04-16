.PHONY: \
	dev \
	run-local \
	run-container \
	deploy \
	generate-data \
	save-data

# Run app with live restart.
# Use https://github.com/watchexec/watchexec .
dev:
	watchexec -r -e go -- make run-local

# Run app in local host machine.
run-local:
	go run application.go

# Run app in container.
run-container:
	docker build -t mirei-tts .
	docker run --rm --interactive --tty --publish 5000 mirei-tts

# Deploy to AWS.
deploy:
	cd infra && yarn && yarn cdk deploy

# Generate static files.
generate-data:
	./generator/clean.sh
	./generator/generate-voice.sh
	./generator/generate-dictionary.sh
	./generator/generate-text-seed.sh

# Save static files used by app to Amazon S3.
save-data:
	aws s3 cp --recursive ./${MTTS_DATA_LOCAL_PREFIX}/ s3://${MTTS_DATA_BUCKET}/

# Fetch static files used by app to Amazon S3.
fetch-data:
	aws s3 sync s3://${MTTS_DATA_BUCKET}/ ./${MTTS_DATA_LOCAL_PREFIX}/
