exec:
	go run cmd/app/main.go

deploy:
	cd cdk && cdk deploy


create-sam-template:
	@echo "Creating SAM template"
	@cdk synth --no-staging > template.yaml

start-lambda:
	@echo "Starting Lambda"
	@make create-sam-template
	@echo "Creating SAM template"
	@cdk synth --no-staging > template.yaml
	@sam local invoke --template template.yaml


