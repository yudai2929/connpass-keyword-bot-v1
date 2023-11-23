GO_FILE_PATH := cmd/batch/send_notification/main.go
FUNCTION_NAME := send-notification
ROLE_NAME=aws-lambda-role
REGION=ap-northeast-1
ACCOUNT_ID=$(shell aws sts get-caller-identity --query Account --output text --region ap-northeast-1)
exec:
	go run cmd/app/main.go

compile:
	GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap $(GO_FILE_PATH)

zip:
	zip function.zip bootstrap


create-role:
	aws iam create-role --role-name $(ROLE_NAME) --assume-role-policy-document file://trust-policy.json


create-function:
	aws lambda create-function --function-name $(FUNCTION_NAME) \
	--runtime provided.al2 --handler bootstrap \
	--role arn:aws:iam::$(ACCOUNT_ID):role/$(ROLE_NAME) \
	--zip-file fileb://function.zip \
	--region $(REGION)

update-function:
	aws lambda update-function-code --function-name $(FUNCTION_NAME) --zip-file fileb://function.zip --region $(REGION)



PHONY: exec compile zip create-role create-function