FROM golang:1.20 as build
WORKDIR /usr/src/app
# Copy dependencies list
COPY go.mod go.sum ./
# Build with optional lambda.norpc tag
COPY cmd/batch/send_notification/main.go .
RUN GOARCH=amd64 GOOS=linux go build -tags lambda.norpc -o main main.go
# Copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /app/main ./main
ENTRYPOINT [ "./main" ]