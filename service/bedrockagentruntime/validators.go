// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagentruntime

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpInvokeAgent struct {
}

func (*validateOpInvokeAgent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpInvokeAgent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*InvokeAgentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpInvokeAgentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRetrieveAndGenerate struct {
}

func (*validateOpRetrieveAndGenerate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRetrieveAndGenerate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RetrieveAndGenerateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRetrieveAndGenerateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRetrieve struct {
}

func (*validateOpRetrieve) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRetrieve) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RetrieveInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRetrieveInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpInvokeAgentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpInvokeAgent{}, middleware.After)
}

func addOpRetrieveAndGenerateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRetrieveAndGenerate{}, middleware.After)
}

func addOpRetrieveValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRetrieve{}, middleware.After)
}

func validateApiResult(v *types.ApiResult) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ApiResult"}
	if v.ActionGroup == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ActionGroup"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateByteContentDoc(v *types.ByteContentDoc) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ByteContentDoc"}
	if v.Identifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Identifier"))
	}
	if v.ContentType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContentType"))
	}
	if v.Data == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Data"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateExternalSource(v *types.ExternalSource) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExternalSource"}
	if len(v.SourceType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("SourceType"))
	}
	if v.S3Location != nil {
		if err := validateS3ObjectDoc(v.S3Location); err != nil {
			invalidParams.AddNested("S3Location", err.(smithy.InvalidParamsError))
		}
	}
	if v.ByteContent != nil {
		if err := validateByteContentDoc(v.ByteContent); err != nil {
			invalidParams.AddNested("ByteContent", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateExternalSources(v []types.ExternalSource) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExternalSources"}
	for i := range v {
		if err := validateExternalSource(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateExternalSourcesGenerationConfiguration(v *types.ExternalSourcesGenerationConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExternalSourcesGenerationConfiguration"}
	if v.GuardrailConfiguration != nil {
		if err := validateGuardrailConfiguration(v.GuardrailConfiguration); err != nil {
			invalidParams.AddNested("GuardrailConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateExternalSourcesRetrieveAndGenerateConfiguration(v *types.ExternalSourcesRetrieveAndGenerateConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExternalSourcesRetrieveAndGenerateConfiguration"}
	if v.ModelArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ModelArn"))
	}
	if v.Sources == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Sources"))
	} else if v.Sources != nil {
		if err := validateExternalSources(v.Sources); err != nil {
			invalidParams.AddNested("Sources", err.(smithy.InvalidParamsError))
		}
	}
	if v.GenerationConfiguration != nil {
		if err := validateExternalSourcesGenerationConfiguration(v.GenerationConfiguration); err != nil {
			invalidParams.AddNested("GenerationConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateFilterAttribute(v *types.FilterAttribute) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "FilterAttribute"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateFunctionResult(v *types.FunctionResult) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "FunctionResult"}
	if v.ActionGroup == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ActionGroup"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateGenerationConfiguration(v *types.GenerationConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GenerationConfiguration"}
	if v.GuardrailConfiguration != nil {
		if err := validateGuardrailConfiguration(v.GuardrailConfiguration); err != nil {
			invalidParams.AddNested("GuardrailConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateGuardrailConfiguration(v *types.GuardrailConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GuardrailConfiguration"}
	if v.GuardrailId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GuardrailId"))
	}
	if v.GuardrailVersion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GuardrailVersion"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateInvocationResultMember(v types.InvocationResultMember) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InvocationResultMember"}
	switch uv := v.(type) {
	case *types.InvocationResultMemberMemberApiResult:
		if err := validateApiResult(&uv.Value); err != nil {
			invalidParams.AddNested("[apiResult]", err.(smithy.InvalidParamsError))
		}

	case *types.InvocationResultMemberMemberFunctionResult:
		if err := validateFunctionResult(&uv.Value); err != nil {
			invalidParams.AddNested("[functionResult]", err.(smithy.InvalidParamsError))
		}

	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateKnowledgeBaseQuery(v *types.KnowledgeBaseQuery) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "KnowledgeBaseQuery"}
	if v.Text == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Text"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateKnowledgeBaseRetrievalConfiguration(v *types.KnowledgeBaseRetrievalConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "KnowledgeBaseRetrievalConfiguration"}
	if v.VectorSearchConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VectorSearchConfiguration"))
	} else if v.VectorSearchConfiguration != nil {
		if err := validateKnowledgeBaseVectorSearchConfiguration(v.VectorSearchConfiguration); err != nil {
			invalidParams.AddNested("VectorSearchConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateKnowledgeBaseRetrieveAndGenerateConfiguration(v *types.KnowledgeBaseRetrieveAndGenerateConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "KnowledgeBaseRetrieveAndGenerateConfiguration"}
	if v.KnowledgeBaseId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KnowledgeBaseId"))
	}
	if v.ModelArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ModelArn"))
	}
	if v.RetrievalConfiguration != nil {
		if err := validateKnowledgeBaseRetrievalConfiguration(v.RetrievalConfiguration); err != nil {
			invalidParams.AddNested("RetrievalConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.GenerationConfiguration != nil {
		if err := validateGenerationConfiguration(v.GenerationConfiguration); err != nil {
			invalidParams.AddNested("GenerationConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateKnowledgeBaseVectorSearchConfiguration(v *types.KnowledgeBaseVectorSearchConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "KnowledgeBaseVectorSearchConfiguration"}
	if v.Filter != nil {
		if err := validateRetrievalFilter(v.Filter); err != nil {
			invalidParams.AddNested("Filter", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRetrievalFilter(v types.RetrievalFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RetrievalFilter"}
	switch uv := v.(type) {
	case *types.RetrievalFilterMemberAndAll:
		if err := validateRetrievalFilterList(uv.Value); err != nil {
			invalidParams.AddNested("[andAll]", err.(smithy.InvalidParamsError))
		}

	case *types.RetrievalFilterMemberEquals:
		if err := validateFilterAttribute(&uv.Value); err != nil {
			invalidParams.AddNested("[equals]", err.(smithy.InvalidParamsError))
		}

	case *types.RetrievalFilterMemberGreaterThan:
		if err := validateFilterAttribute(&uv.Value); err != nil {
			invalidParams.AddNested("[greaterThan]", err.(smithy.InvalidParamsError))
		}

	case *types.RetrievalFilterMemberGreaterThanOrEquals:
		if err := validateFilterAttribute(&uv.Value); err != nil {
			invalidParams.AddNested("[greaterThanOrEquals]", err.(smithy.InvalidParamsError))
		}

	case *types.RetrievalFilterMemberIn:
		if err := validateFilterAttribute(&uv.Value); err != nil {
			invalidParams.AddNested("[in]", err.(smithy.InvalidParamsError))
		}

	case *types.RetrievalFilterMemberLessThan:
		if err := validateFilterAttribute(&uv.Value); err != nil {
			invalidParams.AddNested("[lessThan]", err.(smithy.InvalidParamsError))
		}

	case *types.RetrievalFilterMemberLessThanOrEquals:
		if err := validateFilterAttribute(&uv.Value); err != nil {
			invalidParams.AddNested("[lessThanOrEquals]", err.(smithy.InvalidParamsError))
		}

	case *types.RetrievalFilterMemberListContains:
		if err := validateFilterAttribute(&uv.Value); err != nil {
			invalidParams.AddNested("[listContains]", err.(smithy.InvalidParamsError))
		}

	case *types.RetrievalFilterMemberNotEquals:
		if err := validateFilterAttribute(&uv.Value); err != nil {
			invalidParams.AddNested("[notEquals]", err.(smithy.InvalidParamsError))
		}

	case *types.RetrievalFilterMemberNotIn:
		if err := validateFilterAttribute(&uv.Value); err != nil {
			invalidParams.AddNested("[notIn]", err.(smithy.InvalidParamsError))
		}

	case *types.RetrievalFilterMemberOrAll:
		if err := validateRetrievalFilterList(uv.Value); err != nil {
			invalidParams.AddNested("[orAll]", err.(smithy.InvalidParamsError))
		}

	case *types.RetrievalFilterMemberStartsWith:
		if err := validateFilterAttribute(&uv.Value); err != nil {
			invalidParams.AddNested("[startsWith]", err.(smithy.InvalidParamsError))
		}

	case *types.RetrievalFilterMemberStringContains:
		if err := validateFilterAttribute(&uv.Value); err != nil {
			invalidParams.AddNested("[stringContains]", err.(smithy.InvalidParamsError))
		}

	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRetrievalFilterList(v []types.RetrievalFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RetrievalFilterList"}
	for i := range v {
		if err := validateRetrievalFilter(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRetrieveAndGenerateConfiguration(v *types.RetrieveAndGenerateConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RetrieveAndGenerateConfiguration"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.KnowledgeBaseConfiguration != nil {
		if err := validateKnowledgeBaseRetrieveAndGenerateConfiguration(v.KnowledgeBaseConfiguration); err != nil {
			invalidParams.AddNested("KnowledgeBaseConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.ExternalSourcesConfiguration != nil {
		if err := validateExternalSourcesRetrieveAndGenerateConfiguration(v.ExternalSourcesConfiguration); err != nil {
			invalidParams.AddNested("ExternalSourcesConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRetrieveAndGenerateInput(v *types.RetrieveAndGenerateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RetrieveAndGenerateInput"}
	if v.Text == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Text"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRetrieveAndGenerateSessionConfiguration(v *types.RetrieveAndGenerateSessionConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RetrieveAndGenerateSessionConfiguration"}
	if v.KmsKeyArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KmsKeyArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateReturnControlInvocationResults(v []types.InvocationResultMember) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ReturnControlInvocationResults"}
	for i := range v {
		if err := validateInvocationResultMember(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3ObjectDoc(v *types.S3ObjectDoc) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3ObjectDoc"}
	if v.Uri == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Uri"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSessionState(v *types.SessionState) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SessionState"}
	if v.ReturnControlInvocationResults != nil {
		if err := validateReturnControlInvocationResults(v.ReturnControlInvocationResults); err != nil {
			invalidParams.AddNested("ReturnControlInvocationResults", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpInvokeAgentInput(v *InvokeAgentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InvokeAgentInput"}
	if v.SessionState != nil {
		if err := validateSessionState(v.SessionState); err != nil {
			invalidParams.AddNested("SessionState", err.(smithy.InvalidParamsError))
		}
	}
	if v.AgentId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AgentId"))
	}
	if v.AgentAliasId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AgentAliasId"))
	}
	if v.SessionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRetrieveAndGenerateInput(v *RetrieveAndGenerateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RetrieveAndGenerateInput"}
	if v.Input == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Input"))
	} else if v.Input != nil {
		if err := validateRetrieveAndGenerateInput(v.Input); err != nil {
			invalidParams.AddNested("Input", err.(smithy.InvalidParamsError))
		}
	}
	if v.RetrieveAndGenerateConfiguration != nil {
		if err := validateRetrieveAndGenerateConfiguration(v.RetrieveAndGenerateConfiguration); err != nil {
			invalidParams.AddNested("RetrieveAndGenerateConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.SessionConfiguration != nil {
		if err := validateRetrieveAndGenerateSessionConfiguration(v.SessionConfiguration); err != nil {
			invalidParams.AddNested("SessionConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRetrieveInput(v *RetrieveInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RetrieveInput"}
	if v.KnowledgeBaseId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KnowledgeBaseId"))
	}
	if v.RetrievalQuery == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RetrievalQuery"))
	} else if v.RetrievalQuery != nil {
		if err := validateKnowledgeBaseQuery(v.RetrievalQuery); err != nil {
			invalidParams.AddNested("RetrievalQuery", err.(smithy.InvalidParamsError))
		}
	}
	if v.RetrievalConfiguration != nil {
		if err := validateKnowledgeBaseRetrievalConfiguration(v.RetrievalConfiguration); err != nil {
			invalidParams.AddNested("RetrievalConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
