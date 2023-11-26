package main

import (
	"os"
	"os/exec"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awseventstargets"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
	"github.com/joho/godotenv"
)

type CdkStackProps struct {
	awscdk.StackProps
}

func NewCdkStack(scope constructs.Construct, id string, props *CdkStackProps) (awscdk.Stack, error) {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	functionName := "send_notification"

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	fn, err := createLambdaFunction(stack, functionName)

	if err != nil {
		return nil, err
	}

	ruleName := functionName + "_rule"

	// EventBridgeスケジュールの定義（東京時間で毎日9時）
	rule := awsevents.NewRule(stack, jsii.String(ruleName), &awsevents.RuleProps{
		Schedule: awsevents.Schedule_Cron(&awsevents.CronOptions{
			Minute: jsii.String("0"),
			Hour:   jsii.String("0"),
			Day:    jsii.String("*"),
			Month:  jsii.String("*"),
			Year:   jsii.String("*"),
		}),
	})

	rule.AddTarget(awseventstargets.NewLambdaFunction(fn, &awseventstargets.LambdaFunctionProps{}))

	return stack, nil
}

func createLambdaFunction(stack awscdk.Stack, functionName string) (awslambda.Function, error) {
	if err := golangBuild(getCmdPaths(functionName)); err != nil {
		return nil, err
	}

	fn := awslambda.NewFunction(stack, jsii.String(functionName), &awslambda.FunctionProps{
		FunctionName: jsii.String(functionName),
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Code:         awslambda.AssetCode_FromAsset(jsii.String("bin/"+functionName), nil),
		Handler:      jsii.String("main"),
		Environment: &map[string]*string{
			"CONNPASS_URL":         jsii.String(os.Getenv("CONNPASS_URL")),
			"USER_ID":              jsii.String(os.Getenv("USER_ID")),
			"CHANNEL_SECRET":       jsii.String(os.Getenv("CHANNEL_SECRET")),
			"CHANNEL_ACCESS_TOKEN": jsii.String(os.Getenv("CHANNEL_ACCESS_TOKEN")),
			"SUPABASE_URL":         jsii.String(os.Getenv("SUPABASE_URL")),
			"SUPABASE_KEY":         jsii.String(os.Getenv("SUPABASE_KEY")),
		},
	})

	return fn, nil
}

func getCmdPaths(functionName string) (string, string) {
	buildPath := "bin/" + functionName + "/bootstrap"
	golangPath := "cmd/batch/" + functionName + "/main.go"

	return buildPath, golangPath
}

func golangBuild(buildPath string, golangPath string) error {
	simpleCmd := exec.Command("go", "build", "-tags", " lambda.norpc", "-o", buildPath, golangPath)
	simpleCmd.Env = append(os.Environ(), "GOOS=linux", "GOARCH=amd64")
	_, err := simpleCmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCdkStack(app, "CdkStack", &CdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	// ---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	// ---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
