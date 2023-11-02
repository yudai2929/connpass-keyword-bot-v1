FNCTION_NAME=notify-events-by-keyword
ENTRY_POINT=KeyWordNotification

TOPIC_NAME=connpass-keyword-bot

REGION=asia-northeast1
PROJECT_ID=connpass-keyword-bot

create-topic:
	gcloud pubsub topics create $(TOPIC_NAME)

create-subscription:
	gcloud pubsub subscriptions create $(TOPIC_NAME)-sub --topic $(TOPIC_NAME)

create-job:
	gcloud scheduler jobs create pubsub $(TOPIC_NAME)-job \
	--schedule="* * * * *" \
	--topic=$(TOPIC_NAME) \
	--location=$(REGION) \
	--time-zone="Asia/Tokyo" \
	--message-body="job-test"


list-job:
	gcloud scheduler jobs list --location=$(REGION)


delete-job:
	gcloud scheduler jobs delete $(TOPIC_NAME)-job --location=$(REGION)

pull-job:
	gcloud pubsub subscriptions pull $(TOPIC_NAME)-sub --limit=10

run-job:
	gcloud scheduler jobs run $(TOPIC_NAME)-job --location=$(REGION)

deploy-function:
	gcloud functions deploy $(FNCTION_NAME) \
	--gen2 \
	--runtime=go121 \
	--region=$(REGION) \
	--source=. \
	--entry-point=HelloPubSub \
	--trigger-topic=$(TOPIC_NAME)

delete-function:
	gcloud functions delete $(FNCTION_NAME) --region=$(REGION)


# 動かない
curl-test:
	curl -X POST -H "Content-Type: application/json" \
	https://$(REGION)-$(PROJECT_ID).cloudfunctions.net/$(FNCTION_NAME) \
	-d '{"message":"Hello World!"}'


publish-topic:
	gcloud pubsub topics publish $(TOPIC_NAME) --message "publish topic test"

log-function:
	gcloud functions logs read --region=$(REGION) --limit=50 $(FNCTION_NAME)

enable-api:
	gcloud services enable cloudfunctions.googleapis.com && \
	gcloud services enable cloudscheduler.googleapis.com && \
	gcloud services enable pubsub.googleapis.com && \
	gcloud services enable cloudbuild.googleapis.com && \
	gcloud services enable run.googleapis.com && \
	gcloud services enable eventarc.googleapis.com

login:
	gcloud auth login

list-project:
	gcloud projects list

select-project:
	gcloud config set project $(PROJECT_ID)

zip-file:
	zip -r function-source.zip .