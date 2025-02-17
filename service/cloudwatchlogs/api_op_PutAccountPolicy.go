// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an account-level data protection policy or subscription filter policy
// that applies to all log groups or a subset of log groups in the account.
//
// # Data protection policy
//
// A data protection policy can help safeguard sensitive data that's ingested by
// your log groups by auditing and masking the sensitive log data. Each account can
// have only one account-level data protection policy.
//
// Sensitive data is detected and masked when it is ingested into a log group.
// When you set a data protection policy, log events ingested into the log groups
// before that time are not masked.
//
// If you use PutAccountPolicy to create a data protection policy for your whole
// account, it applies to both existing log groups and all log groups that are
// created later in this account. The account-level policy is applied to existing
// log groups with eventual consistency. It might take up to 5 minutes before
// sensitive data in existing log groups begins to be masked.
//
// By default, when a user views a log event that includes masked data, the
// sensitive data is replaced by asterisks. A user who has the logs:Unmask
// permission can use a [GetLogEvents]or [FilterLogEvents] operation with the unmask parameter set to true to
// view the unmasked log events. Users with the logs:Unmask can also view unmasked
// data in the CloudWatch Logs console by running a CloudWatch Logs Insights query
// with the unmask query command.
//
// For more information, including a list of types of data that can be audited and
// masked, see [Protect sensitive log data with masking].
//
// To use the PutAccountPolicy operation for a data protection policy, you must be
// signed on with the logs:PutDataProtectionPolicy and logs:PutAccountPolicy
// permissions.
//
// The PutAccountPolicy operation applies to all log groups in the account. You
// can use [PutDataProtectionPolicy]to create a data protection policy that applies to just one log group.
// If a log group has its own data protection policy and the account also has an
// account-level data protection policy, then the two policies are cumulative. Any
// sensitive term specified in either policy is masked.
//
// # Subscription filter policy
//
// A subscription filter policy sets up a real-time feed of log events from
// CloudWatch Logs to other Amazon Web Services services. Account-level
// subscription filter policies apply to both existing log groups and log groups
// that are created later in this account. Supported destinations are Kinesis Data
// Streams, Firehose, and Lambda. When log events are sent to the receiving
// service, they are Base64 encoded and compressed with the GZIP format.
//
// The following destinations are supported for subscription filters:
//
//   - An Kinesis Data Streams data stream in the same account as the subscription
//     policy, for same-account delivery.
//
//   - An Firehose data stream in the same account as the subscription policy, for
//     same-account delivery.
//
//   - A Lambda function in the same account as the subscription policy, for
//     same-account delivery.
//
//   - A logical destination in a different account created with [PutDestination], for
//     cross-account delivery. Kinesis Data Streams and Firehose are supported as
//     logical destinations.
//
// Each account can have one account-level subscription filter policy. If you are
// updating an existing filter, you must specify the correct name in PolicyName .
// To perform a PutAccountPolicy subscription filter operation for any destination
// except a Lambda function, you must also have the iam:PassRole permission.
//
// [PutDestination]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDestination.html
// [PutDataProtectionPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDataProtectionPolicy.html
// [Protect sensitive log data with masking]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data.html
// [FilterLogEvents]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_FilterLogEvents.html
// [GetLogEvents]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetLogEvents.html
func (c *Client) PutAccountPolicy(ctx context.Context, params *PutAccountPolicyInput, optFns ...func(*Options)) (*PutAccountPolicyOutput, error) {
	if params == nil {
		params = &PutAccountPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAccountPolicy", params, optFns, c.addOperationPutAccountPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAccountPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAccountPolicyInput struct {

	// Specify the policy, in JSON.
	//
	// Data protection policy
	//
	// A data protection policy must include two JSON blocks:
	//
	//   - The first block must include both a DataIdentifer array and an Operation
	//   property with an Audit action. The DataIdentifer array lists the types of
	//   sensitive data that you want to mask. For more information about the available
	//   options, see [Types of data that you can mask].
	//
	// The Operation property with an Audit action is required to find the sensitive
	//   data terms. This Audit action must contain a FindingsDestination object. You
	//   can optionally use that FindingsDestination object to list one or more
	//   destinations to send audit findings to. If you specify destinations such as log
	//   groups, Firehose streams, and S3 buckets, they must already exist.
	//
	//   - The second block must include both a DataIdentifer array and an Operation
	//   property with an Deidentify action. The DataIdentifer array must exactly match
	//   the DataIdentifer array in the first block of the policy.
	//
	// The Operation property with the Deidentify action is what actually masks the
	//   data, and it must contain the "MaskConfig": {} object. The "MaskConfig": {}
	//   object must be empty.
	//
	// For an example data protection policy, see the Examples section on this page.
	//
	// The contents of the two DataIdentifer arrays must match exactly.
	//
	// In addition to the two JSON blocks, the policyDocument can also include Name ,
	// Description , and Version fields. The Name is different than the operation's
	// policyName parameter, and is used as a dimension when CloudWatch Logs reports
	// audit findings metrics to CloudWatch.
	//
	// The JSON specified in policyDocument can be up to 30,720 characters long.
	//
	// Subscription filter policy
	//
	// A subscription filter policy can include the following attributes in a JSON
	// block:
	//
	//   - DestinationArn The ARN of the destination to deliver log events to.
	//   Supported destinations are:
	//
	//   - An Kinesis Data Streams data stream in the same account as the subscription
	//   policy, for same-account delivery.
	//
	//   - An Firehose data stream in the same account as the subscription policy, for
	//   same-account delivery.
	//
	//   - A Lambda function in the same account as the subscription policy, for
	//   same-account delivery.
	//
	//   - A logical destination in a different account created with [PutDestination], for
	//   cross-account delivery. Kinesis Data Streams and Firehose are supported as
	//   logical destinations.
	//
	//   - RoleArn The ARN of an IAM role that grants CloudWatch Logs permissions to
	//   deliver ingested log events to the destination stream. You don't need to provide
	//   the ARN when you are working with a logical destination for cross-account
	//   delivery.
	//
	//   - FilterPattern A filter pattern for subscribing to a filtered stream of log
	//   events.
	//
	//   - DistributionThe method used to distribute log data to the destination. By
	//   default, log data is grouped by log stream, but the grouping can be set to
	//   Random for a more even distribution. This property is only applicable when the
	//   destination is an Kinesis Data Streams data stream.
	//
	// [PutDestination]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDestination.html
	// [Types of data that you can mask]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data-types.html
	//
	// This member is required.
	PolicyDocument *string

	// A name for the policy. This must be unique within the account.
	//
	// This member is required.
	PolicyName *string

	// The type of policy that you're creating or updating.
	//
	// This member is required.
	PolicyType types.PolicyType

	// Currently the only valid value for this parameter is ALL , which specifies that
	// the data protection policy applies to all log groups in the account. If you omit
	// this parameter, the default of ALL is used.
	Scope types.Scope

	// Use this parameter to apply the subscription filter policy to a subset of log
	// groups in the account. Currently, the only supported filter is LogGroupName NOT
	// IN [] . The selectionCriteria string can be up to 25KB in length. The length is
	// determined by using its UTF-8 bytes.
	//
	// Using the selectionCriteria parameter is useful to help prevent infinite loops.
	// For more information, see [Log recursion prevention].
	//
	// Specifing selectionCriteria is valid only when you specify
	// SUBSCRIPTION_FILTER_POLICY for policyType .
	//
	// [Log recursion prevention]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Subscriptions-recursion-prevention.html
	SelectionCriteria *string

	noSmithyDocumentSerde
}

type PutAccountPolicyOutput struct {

	// The account policy that you created.
	AccountPolicy *types.AccountPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAccountPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutAccountPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutAccountPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutAccountPolicy"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutAccountPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAccountPolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutAccountPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutAccountPolicy",
	}
}
