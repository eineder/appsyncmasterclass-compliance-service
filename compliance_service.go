package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ComplianceServiceStackProps struct {
	awscdk.StackProps
}

func NewComplianceServiceStack(scope constructs.Construct, id string, props *ComplianceServiceStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	fn := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("handler-on-new-tweet-posted"), &awscdklambdagoalpha.GoFunctionProps{
		Entry:   jsii.String("lambda/on_tweet_posted.go"),
		Runtime: awslambda.Runtime_PROVIDED_AL2(),
	})

	rule := awsevents.NewRule(stack, jsii.String("rule-on-new-tweet-posted"), &awsevents.RuleProps{
		EventPattern: &awsevents.EventPattern{
			DetailType: jsii.Strings("new_tweet_posted"),
		},
	})

	target := awseventstargets.NewLambdaFunction(fn, &awseventstargets.LambdaFunctionProps{})
	rule.AddTarget(target)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewComplianceServiceStack(app, "appsyncmasterclass-ComplianceServiceStack", &ComplianceServiceStackProps{
		awscdk.StackProps{},
	})

	app.Synth(nil)
}
