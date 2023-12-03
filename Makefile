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


generate-mock:
	@echo "Generating mocks"
	@mockgen -source=pkg/domain/repository/event_repository_interface.go -destination=./mocks/repository/event_repository_mock.go