package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awseventstargets"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
	"github.com/yudai2929/connpass-keyword-bot-v1/cdk/config"
	"github.com/yudai2929/connpass-keyword-bot-v1/cdk/libs/lambda"
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

	lp := lambda.NewFunctionProps(
		"bin/send_notification",
		"cmd/batch/send_notification/main.go",
		"send_notification",
	)

	fn, err := lambda.BuildAndCreateFunction(stack, lp, &map[string]*string{
		"CONNPASS_URL":         jsii.String(os.Getenv("CONNPASS_URL")),
		"USER_ID":              jsii.String(os.Getenv("USER_ID")),
		"CHANNEL_SECRET":       jsii.String(os.Getenv("CHANNEL_SECRET")),
		"CHANNEL_ACCESS_TOKEN": jsii.String(os.Getenv("CHANNEL_ACCESS_TOKEN")),
		"SUPABASE_URL":         jsii.String(os.Getenv("SUPABASE_URL")),
		"SUPABASE_KEY":         jsii.String(os.Getenv("SUPABASE_KEY")),
		"YAHOO_CLIENT_ID":      jsii.String(os.Getenv("YAHOO_CLIENT_ID")),
	})

	if err != nil {
		return nil, err
	}

	ruleName := lp.FunctionName + "_rule"

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

func main() {
	defer jsii.Close()

	if err := config.SetupEnv(); err != nil {
		panic(err)
	}

	app := awscdk.NewApp(nil)

	NewCdkStack(app, "CdkStack", &CdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {

	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("AWS_ACCOUNT_ID")),
		Region:  jsii.String(os.Getenv("AWS_DEFAULT_REGION")),
	}
}
