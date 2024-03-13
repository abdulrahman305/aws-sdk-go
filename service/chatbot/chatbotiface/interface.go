// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package chatbotiface provides an interface to enable mocking the AWS Chatbot service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package chatbotiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/chatbot"
)

// ChatbotAPI provides an interface to enable mocking the
// chatbot.Chatbot service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Chatbot.
//	func myFunc(svc chatbotiface.ChatbotAPI) bool {
//	    // Make svc.CreateChimeWebhookConfiguration request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := chatbot.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockChatbotClient struct {
//	    chatbotiface.ChatbotAPI
//	}
//	func (m *mockChatbotClient) CreateChimeWebhookConfiguration(input *chatbot.CreateChimeWebhookConfigurationInput) (*chatbot.CreateChimeWebhookConfigurationOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockChatbotClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ChatbotAPI interface {
	CreateChimeWebhookConfiguration(*chatbot.CreateChimeWebhookConfigurationInput) (*chatbot.CreateChimeWebhookConfigurationOutput, error)
	CreateChimeWebhookConfigurationWithContext(aws.Context, *chatbot.CreateChimeWebhookConfigurationInput, ...request.Option) (*chatbot.CreateChimeWebhookConfigurationOutput, error)
	CreateChimeWebhookConfigurationRequest(*chatbot.CreateChimeWebhookConfigurationInput) (*request.Request, *chatbot.CreateChimeWebhookConfigurationOutput)

	CreateMicrosoftTeamsChannelConfiguration(*chatbot.CreateMicrosoftTeamsChannelConfigurationInput) (*chatbot.CreateMicrosoftTeamsChannelConfigurationOutput, error)
	CreateMicrosoftTeamsChannelConfigurationWithContext(aws.Context, *chatbot.CreateMicrosoftTeamsChannelConfigurationInput, ...request.Option) (*chatbot.CreateMicrosoftTeamsChannelConfigurationOutput, error)
	CreateMicrosoftTeamsChannelConfigurationRequest(*chatbot.CreateMicrosoftTeamsChannelConfigurationInput) (*request.Request, *chatbot.CreateMicrosoftTeamsChannelConfigurationOutput)

	CreateSlackChannelConfiguration(*chatbot.CreateSlackChannelConfigurationInput) (*chatbot.CreateSlackChannelConfigurationOutput, error)
	CreateSlackChannelConfigurationWithContext(aws.Context, *chatbot.CreateSlackChannelConfigurationInput, ...request.Option) (*chatbot.CreateSlackChannelConfigurationOutput, error)
	CreateSlackChannelConfigurationRequest(*chatbot.CreateSlackChannelConfigurationInput) (*request.Request, *chatbot.CreateSlackChannelConfigurationOutput)

	DeleteChimeWebhookConfiguration(*chatbot.DeleteChimeWebhookConfigurationInput) (*chatbot.DeleteChimeWebhookConfigurationOutput, error)
	DeleteChimeWebhookConfigurationWithContext(aws.Context, *chatbot.DeleteChimeWebhookConfigurationInput, ...request.Option) (*chatbot.DeleteChimeWebhookConfigurationOutput, error)
	DeleteChimeWebhookConfigurationRequest(*chatbot.DeleteChimeWebhookConfigurationInput) (*request.Request, *chatbot.DeleteChimeWebhookConfigurationOutput)

	DeleteMicrosoftTeamsChannelConfiguration(*chatbot.DeleteMicrosoftTeamsChannelConfigurationInput) (*chatbot.DeleteMicrosoftTeamsChannelConfigurationOutput, error)
	DeleteMicrosoftTeamsChannelConfigurationWithContext(aws.Context, *chatbot.DeleteMicrosoftTeamsChannelConfigurationInput, ...request.Option) (*chatbot.DeleteMicrosoftTeamsChannelConfigurationOutput, error)
	DeleteMicrosoftTeamsChannelConfigurationRequest(*chatbot.DeleteMicrosoftTeamsChannelConfigurationInput) (*request.Request, *chatbot.DeleteMicrosoftTeamsChannelConfigurationOutput)

	DeleteMicrosoftTeamsConfiguredTeam(*chatbot.DeleteMicrosoftTeamsConfiguredTeamInput) (*chatbot.DeleteMicrosoftTeamsConfiguredTeamOutput, error)
	DeleteMicrosoftTeamsConfiguredTeamWithContext(aws.Context, *chatbot.DeleteMicrosoftTeamsConfiguredTeamInput, ...request.Option) (*chatbot.DeleteMicrosoftTeamsConfiguredTeamOutput, error)
	DeleteMicrosoftTeamsConfiguredTeamRequest(*chatbot.DeleteMicrosoftTeamsConfiguredTeamInput) (*request.Request, *chatbot.DeleteMicrosoftTeamsConfiguredTeamOutput)

	DeleteMicrosoftTeamsUserIdentity(*chatbot.DeleteMicrosoftTeamsUserIdentityInput) (*chatbot.DeleteMicrosoftTeamsUserIdentityOutput, error)
	DeleteMicrosoftTeamsUserIdentityWithContext(aws.Context, *chatbot.DeleteMicrosoftTeamsUserIdentityInput, ...request.Option) (*chatbot.DeleteMicrosoftTeamsUserIdentityOutput, error)
	DeleteMicrosoftTeamsUserIdentityRequest(*chatbot.DeleteMicrosoftTeamsUserIdentityInput) (*request.Request, *chatbot.DeleteMicrosoftTeamsUserIdentityOutput)

	DeleteSlackChannelConfiguration(*chatbot.DeleteSlackChannelConfigurationInput) (*chatbot.DeleteSlackChannelConfigurationOutput, error)
	DeleteSlackChannelConfigurationWithContext(aws.Context, *chatbot.DeleteSlackChannelConfigurationInput, ...request.Option) (*chatbot.DeleteSlackChannelConfigurationOutput, error)
	DeleteSlackChannelConfigurationRequest(*chatbot.DeleteSlackChannelConfigurationInput) (*request.Request, *chatbot.DeleteSlackChannelConfigurationOutput)

	DeleteSlackUserIdentity(*chatbot.DeleteSlackUserIdentityInput) (*chatbot.DeleteSlackUserIdentityOutput, error)
	DeleteSlackUserIdentityWithContext(aws.Context, *chatbot.DeleteSlackUserIdentityInput, ...request.Option) (*chatbot.DeleteSlackUserIdentityOutput, error)
	DeleteSlackUserIdentityRequest(*chatbot.DeleteSlackUserIdentityInput) (*request.Request, *chatbot.DeleteSlackUserIdentityOutput)

	DeleteSlackWorkspaceAuthorization(*chatbot.DeleteSlackWorkspaceAuthorizationInput) (*chatbot.DeleteSlackWorkspaceAuthorizationOutput, error)
	DeleteSlackWorkspaceAuthorizationWithContext(aws.Context, *chatbot.DeleteSlackWorkspaceAuthorizationInput, ...request.Option) (*chatbot.DeleteSlackWorkspaceAuthorizationOutput, error)
	DeleteSlackWorkspaceAuthorizationRequest(*chatbot.DeleteSlackWorkspaceAuthorizationInput) (*request.Request, *chatbot.DeleteSlackWorkspaceAuthorizationOutput)

	DescribeChimeWebhookConfigurations(*chatbot.DescribeChimeWebhookConfigurationsInput) (*chatbot.DescribeChimeWebhookConfigurationsOutput, error)
	DescribeChimeWebhookConfigurationsWithContext(aws.Context, *chatbot.DescribeChimeWebhookConfigurationsInput, ...request.Option) (*chatbot.DescribeChimeWebhookConfigurationsOutput, error)
	DescribeChimeWebhookConfigurationsRequest(*chatbot.DescribeChimeWebhookConfigurationsInput) (*request.Request, *chatbot.DescribeChimeWebhookConfigurationsOutput)

	DescribeChimeWebhookConfigurationsPages(*chatbot.DescribeChimeWebhookConfigurationsInput, func(*chatbot.DescribeChimeWebhookConfigurationsOutput, bool) bool) error
	DescribeChimeWebhookConfigurationsPagesWithContext(aws.Context, *chatbot.DescribeChimeWebhookConfigurationsInput, func(*chatbot.DescribeChimeWebhookConfigurationsOutput, bool) bool, ...request.Option) error

	DescribeSlackChannelConfigurations(*chatbot.DescribeSlackChannelConfigurationsInput) (*chatbot.DescribeSlackChannelConfigurationsOutput, error)
	DescribeSlackChannelConfigurationsWithContext(aws.Context, *chatbot.DescribeSlackChannelConfigurationsInput, ...request.Option) (*chatbot.DescribeSlackChannelConfigurationsOutput, error)
	DescribeSlackChannelConfigurationsRequest(*chatbot.DescribeSlackChannelConfigurationsInput) (*request.Request, *chatbot.DescribeSlackChannelConfigurationsOutput)

	DescribeSlackChannelConfigurationsPages(*chatbot.DescribeSlackChannelConfigurationsInput, func(*chatbot.DescribeSlackChannelConfigurationsOutput, bool) bool) error
	DescribeSlackChannelConfigurationsPagesWithContext(aws.Context, *chatbot.DescribeSlackChannelConfigurationsInput, func(*chatbot.DescribeSlackChannelConfigurationsOutput, bool) bool, ...request.Option) error

	DescribeSlackUserIdentities(*chatbot.DescribeSlackUserIdentitiesInput) (*chatbot.DescribeSlackUserIdentitiesOutput, error)
	DescribeSlackUserIdentitiesWithContext(aws.Context, *chatbot.DescribeSlackUserIdentitiesInput, ...request.Option) (*chatbot.DescribeSlackUserIdentitiesOutput, error)
	DescribeSlackUserIdentitiesRequest(*chatbot.DescribeSlackUserIdentitiesInput) (*request.Request, *chatbot.DescribeSlackUserIdentitiesOutput)

	DescribeSlackUserIdentitiesPages(*chatbot.DescribeSlackUserIdentitiesInput, func(*chatbot.DescribeSlackUserIdentitiesOutput, bool) bool) error
	DescribeSlackUserIdentitiesPagesWithContext(aws.Context, *chatbot.DescribeSlackUserIdentitiesInput, func(*chatbot.DescribeSlackUserIdentitiesOutput, bool) bool, ...request.Option) error

	DescribeSlackWorkspaces(*chatbot.DescribeSlackWorkspacesInput) (*chatbot.DescribeSlackWorkspacesOutput, error)
	DescribeSlackWorkspacesWithContext(aws.Context, *chatbot.DescribeSlackWorkspacesInput, ...request.Option) (*chatbot.DescribeSlackWorkspacesOutput, error)
	DescribeSlackWorkspacesRequest(*chatbot.DescribeSlackWorkspacesInput) (*request.Request, *chatbot.DescribeSlackWorkspacesOutput)

	DescribeSlackWorkspacesPages(*chatbot.DescribeSlackWorkspacesInput, func(*chatbot.DescribeSlackWorkspacesOutput, bool) bool) error
	DescribeSlackWorkspacesPagesWithContext(aws.Context, *chatbot.DescribeSlackWorkspacesInput, func(*chatbot.DescribeSlackWorkspacesOutput, bool) bool, ...request.Option) error

	GetAccountPreferences(*chatbot.GetAccountPreferencesInput) (*chatbot.GetAccountPreferencesOutput, error)
	GetAccountPreferencesWithContext(aws.Context, *chatbot.GetAccountPreferencesInput, ...request.Option) (*chatbot.GetAccountPreferencesOutput, error)
	GetAccountPreferencesRequest(*chatbot.GetAccountPreferencesInput) (*request.Request, *chatbot.GetAccountPreferencesOutput)

	GetMicrosoftTeamsChannelConfiguration(*chatbot.GetMicrosoftTeamsChannelConfigurationInput) (*chatbot.GetMicrosoftTeamsChannelConfigurationOutput, error)
	GetMicrosoftTeamsChannelConfigurationWithContext(aws.Context, *chatbot.GetMicrosoftTeamsChannelConfigurationInput, ...request.Option) (*chatbot.GetMicrosoftTeamsChannelConfigurationOutput, error)
	GetMicrosoftTeamsChannelConfigurationRequest(*chatbot.GetMicrosoftTeamsChannelConfigurationInput) (*request.Request, *chatbot.GetMicrosoftTeamsChannelConfigurationOutput)

	ListMicrosoftTeamsChannelConfigurations(*chatbot.ListMicrosoftTeamsChannelConfigurationsInput) (*chatbot.ListMicrosoftTeamsChannelConfigurationsOutput, error)
	ListMicrosoftTeamsChannelConfigurationsWithContext(aws.Context, *chatbot.ListMicrosoftTeamsChannelConfigurationsInput, ...request.Option) (*chatbot.ListMicrosoftTeamsChannelConfigurationsOutput, error)
	ListMicrosoftTeamsChannelConfigurationsRequest(*chatbot.ListMicrosoftTeamsChannelConfigurationsInput) (*request.Request, *chatbot.ListMicrosoftTeamsChannelConfigurationsOutput)

	ListMicrosoftTeamsChannelConfigurationsPages(*chatbot.ListMicrosoftTeamsChannelConfigurationsInput, func(*chatbot.ListMicrosoftTeamsChannelConfigurationsOutput, bool) bool) error
	ListMicrosoftTeamsChannelConfigurationsPagesWithContext(aws.Context, *chatbot.ListMicrosoftTeamsChannelConfigurationsInput, func(*chatbot.ListMicrosoftTeamsChannelConfigurationsOutput, bool) bool, ...request.Option) error

	ListMicrosoftTeamsConfiguredTeams(*chatbot.ListMicrosoftTeamsConfiguredTeamsInput) (*chatbot.ListMicrosoftTeamsConfiguredTeamsOutput, error)
	ListMicrosoftTeamsConfiguredTeamsWithContext(aws.Context, *chatbot.ListMicrosoftTeamsConfiguredTeamsInput, ...request.Option) (*chatbot.ListMicrosoftTeamsConfiguredTeamsOutput, error)
	ListMicrosoftTeamsConfiguredTeamsRequest(*chatbot.ListMicrosoftTeamsConfiguredTeamsInput) (*request.Request, *chatbot.ListMicrosoftTeamsConfiguredTeamsOutput)

	ListMicrosoftTeamsConfiguredTeamsPages(*chatbot.ListMicrosoftTeamsConfiguredTeamsInput, func(*chatbot.ListMicrosoftTeamsConfiguredTeamsOutput, bool) bool) error
	ListMicrosoftTeamsConfiguredTeamsPagesWithContext(aws.Context, *chatbot.ListMicrosoftTeamsConfiguredTeamsInput, func(*chatbot.ListMicrosoftTeamsConfiguredTeamsOutput, bool) bool, ...request.Option) error

	ListMicrosoftTeamsUserIdentities(*chatbot.ListMicrosoftTeamsUserIdentitiesInput) (*chatbot.ListMicrosoftTeamsUserIdentitiesOutput, error)
	ListMicrosoftTeamsUserIdentitiesWithContext(aws.Context, *chatbot.ListMicrosoftTeamsUserIdentitiesInput, ...request.Option) (*chatbot.ListMicrosoftTeamsUserIdentitiesOutput, error)
	ListMicrosoftTeamsUserIdentitiesRequest(*chatbot.ListMicrosoftTeamsUserIdentitiesInput) (*request.Request, *chatbot.ListMicrosoftTeamsUserIdentitiesOutput)

	ListMicrosoftTeamsUserIdentitiesPages(*chatbot.ListMicrosoftTeamsUserIdentitiesInput, func(*chatbot.ListMicrosoftTeamsUserIdentitiesOutput, bool) bool) error
	ListMicrosoftTeamsUserIdentitiesPagesWithContext(aws.Context, *chatbot.ListMicrosoftTeamsUserIdentitiesInput, func(*chatbot.ListMicrosoftTeamsUserIdentitiesOutput, bool) bool, ...request.Option) error

	UpdateAccountPreferences(*chatbot.UpdateAccountPreferencesInput) (*chatbot.UpdateAccountPreferencesOutput, error)
	UpdateAccountPreferencesWithContext(aws.Context, *chatbot.UpdateAccountPreferencesInput, ...request.Option) (*chatbot.UpdateAccountPreferencesOutput, error)
	UpdateAccountPreferencesRequest(*chatbot.UpdateAccountPreferencesInput) (*request.Request, *chatbot.UpdateAccountPreferencesOutput)

	UpdateChimeWebhookConfiguration(*chatbot.UpdateChimeWebhookConfigurationInput) (*chatbot.UpdateChimeWebhookConfigurationOutput, error)
	UpdateChimeWebhookConfigurationWithContext(aws.Context, *chatbot.UpdateChimeWebhookConfigurationInput, ...request.Option) (*chatbot.UpdateChimeWebhookConfigurationOutput, error)
	UpdateChimeWebhookConfigurationRequest(*chatbot.UpdateChimeWebhookConfigurationInput) (*request.Request, *chatbot.UpdateChimeWebhookConfigurationOutput)

	UpdateMicrosoftTeamsChannelConfiguration(*chatbot.UpdateMicrosoftTeamsChannelConfigurationInput) (*chatbot.UpdateMicrosoftTeamsChannelConfigurationOutput, error)
	UpdateMicrosoftTeamsChannelConfigurationWithContext(aws.Context, *chatbot.UpdateMicrosoftTeamsChannelConfigurationInput, ...request.Option) (*chatbot.UpdateMicrosoftTeamsChannelConfigurationOutput, error)
	UpdateMicrosoftTeamsChannelConfigurationRequest(*chatbot.UpdateMicrosoftTeamsChannelConfigurationInput) (*request.Request, *chatbot.UpdateMicrosoftTeamsChannelConfigurationOutput)

	UpdateSlackChannelConfiguration(*chatbot.UpdateSlackChannelConfigurationInput) (*chatbot.UpdateSlackChannelConfigurationOutput, error)
	UpdateSlackChannelConfigurationWithContext(aws.Context, *chatbot.UpdateSlackChannelConfigurationInput, ...request.Option) (*chatbot.UpdateSlackChannelConfigurationOutput, error)
	UpdateSlackChannelConfigurationRequest(*chatbot.UpdateSlackChannelConfigurationInput) (*request.Request, *chatbot.UpdateSlackChannelConfigurationOutput)
}

var _ ChatbotAPI = (*chatbot.Chatbot)(nil)