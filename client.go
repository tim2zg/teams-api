package teams_api

import (
	"fmt"
	"github.com/fossteams/teams-api/pkg/csa"
	"github.com/fossteams/teams-api/pkg/models"
	"github.com/fossteams/teams-api/pkg/mt"
	"net/http"
)
import "github.com/fossteams/teams-api/pkg"

type TeamsClient struct {
	httpClient *http.Client
	chatSvc    *csa.CSASvc
	mtSvc      *mt.Service
}

func (t *TeamsClient) Debug(debugFlag bool) {
	t.chatSvc.DebugSave(debugFlag)
}

func (t *TeamsClient) ChatSvc() *csa.CSASvc {
	return t.chatSvc
}

func New() (*TeamsClient, error) {
	teamsClient := TeamsClient{}

	// Get Skype Spaces Token
	skypeSpaces, err := api.GetSkypeSpacesToken()
	if err != nil {
		return nil, fmt.Errorf("unable to get SkypeSpaces token: %v", err)
	}

	chatSvcToken, err := api.GetChatSvcAggToken()
	if err != nil {
		return nil, fmt.Errorf("unable to get chat service token: %v", err)
	}

	skypeToken, err := api.GetSkypeToken()
	if err != nil {
		return nil, err
	}

	teamsToken, err := api.GetTeamsToken()
	if err != nil {
		return nil, err
	}

	// Initialize Chat Service
	csaSvc, err := csa.NewCSAService(chatSvcToken, skypeToken)
	if err != nil {
		return nil, fmt.Errorf("unable to init Chat Service")
	}
	teamsClient.chatSvc = csaSvc

	// Initialize MT Service
	mtSvc, err := mt.NewMiddleTierService(api.Emea, skypeSpaces, teamsToken)
	if err != nil {
		return nil, fmt.Errorf("unable to init MT Service: %v", err)
	}
	teamsClient.mtSvc = mtSvc

	return &teamsClient, err
}

func (t *TeamsClient) GetConversations() (*models.ConversationResponse, error) {
	return t.chatSvc.GetConversations()
}

func (t *TeamsClient) GetMessages(channel *models.Channel) ([]models.ChatMessage, error) {
	return t.chatSvc.GetMessagesByChannel(channel)
}

func (t *TeamsClient) GetMe() (*models.User, error) {
	return t.mtSvc.GetMe()
}

func (t *TeamsClient) FetchShortProfile(mris []string) ([]models.User, error) {
	return t.mtSvc.FetchShortProfile(mris...)
}

func (t *TeamsClient) GetProfilePicture(emailOrId string) ([]byte, error) {
	return t.mtSvc.GetProfilePicture(emailOrId)
}

func (t *TeamsClient) GetTeamsProfilePicture(emailOrId string) ([]byte, error) {
	return t.mtSvc.GetTeamsProfilePicture(emailOrId)
}

func (t *TeamsClient) GetTenants() ([]models.Tenant, error) {
	return t.mtSvc.GetTenants()
}

func (t *TeamsClient) GetPinnedChannels() ([]csa.ChannelId, error) {
	return t.chatSvc.GetPinnedChannels()
}

func (t *TeamsClient) SendMessage(channel string, message string) error {
	return t.chatSvc.SendMessage(channel, message)
}

func (t *TeamsClient) ReactToMessage(channel string, messageID string, reaction string) error {
	return t.chatSvc.ReactToMessage(channel, messageID, reaction)
}

func (t *TeamsClient) RemoveReactionToMessage(channel string, messageID string, reaction string) error {
	return t.chatSvc.RemoveReactionToMessage(channel, messageID, reaction)
}

func (t *TeamsClient) DeleteMessage(channel string, message string) (err error) {
	return t.chatSvc.DeleteMessage(channel, message)
}
