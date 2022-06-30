package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type GoServerlessCDKStackProps struct {
	awscdk.StackProps
}

func NewGoServerlessStack(scope constructs.Construct, id string, props *GoServerlessCDKStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	dynamoDBTable := awsdynamodb.NewTable(stack, jsii.String("dynamodb-table"), &awsdynamodb.TableProps{PartitionKey: &awsdynamodb.Attribute{Name: jsii.String("email"), Type: awsdynamodb.AttributeType_STRING}})

	function := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("demo-func"), &awscdklambdagoalpha.GoFunctionProps{Runtime: awslambda.Runtime_GO_1_X(), Environment: &map[string]*string{"TABLE_NAME": dynamoDBTable.TableName()}, Entry: jsii.String("../part2")})

	dynamoDBTable.GrantWriteData(function)
	dynamoDBTable.ApplyRemovalPolicy(awscdk.RemovalPolicy_DESTROY)

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewGoServerlessStack(app, "GoServerlessStack", &GoServerlessCDKStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return nil
}
