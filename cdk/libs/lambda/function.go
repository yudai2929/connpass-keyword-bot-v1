package lambda

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/jsii-runtime-go"
	"github.com/yudai2929/connpass-keyword-bot-v1/cdk/utils/build"
)

type FunctionProps struct {
	BuildDir        string
	BuildOutputPath string
	GolangPath      string
	Handler         string
	FunctionName    string
}

func NewFunctionProps(buildDir string, golangPath string, functionName string) FunctionProps {
	return FunctionProps{
		BuildDir:        buildDir,
		BuildOutputPath: buildDir + "/" + "bootstrap",
		GolangPath:      golangPath,
		Handler:         "main",
		FunctionName:    functionName,
	}
}

func BuildAndCreateFunction(stack awscdk.Stack, props FunctionProps, env *map[string]*string) (awslambda.Function, error) {
	if err := build.Golang(props.BuildOutputPath, props.GolangPath); err != nil {
		return nil, err
	}

	fn := awslambda.NewFunction(stack, jsii.String(props.FunctionName), &awslambda.FunctionProps{
		FunctionName: jsii.String(props.FunctionName),
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Code:         awslambda.AssetCode_FromAsset(jsii.String(props.BuildDir), nil),
		Handler:      jsii.String(props.Handler),
		Environment:  env,
	})

	return fn, nil
}
