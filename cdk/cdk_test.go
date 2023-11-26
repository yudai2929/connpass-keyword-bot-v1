package main

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/assertions"
	"github.com/aws/jsii-runtime-go"
)

func TestNewCdkStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack, err := NewCdkStack(app, "MyStack", nil)

	// THEN
	if err != nil {
		t.Fatalf("Failed to create stack: %v", err)
	}

	template := assertions.Template_FromStack(stack)

	// Check if the stack contains a Lambda function named "send_notification"
	template.HasResourceProperties(jsii.String("AWS::Lambda::Function"), map[string]interface{}{
		"FunctionName": "send_notification",
	})
}

func TestNewCdkStackWithProps(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)
	props := &CdkStackProps{
		StackProps: awscdk.StackProps{
			Env: &awscdk.Environment{
				Account: jsii.String("123456789012"),
				Region:  jsii.String("us-east-1"),
			},
		},
	}

	// WHEN
	stack, err := NewCdkStack(app, "MyStack", props)

	// THEN
	if err != nil {
		t.Fatalf("Failed to create stack: %v", err)
	}

	template := assertions.Template_FromStack(stack)

	// Check if the stack contains a Lambda function named "send_notification"
	template.HasResourceProperties(jsii.String("AWS::Lambda::Function"), map[string]interface{}{
		"FunctionName": "send_notification",
	})
}
