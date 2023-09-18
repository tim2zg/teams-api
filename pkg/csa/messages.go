package csa

import (
	"encoding/json"
	"fmt"
	"github.com/fossteams/teams-api/pkg/errors"
	models2 "github.com/fossteams/teams-api/pkg/models"
	"github.com/fossteams/teams-api/pkg/util"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func (c *CSASvc) GetMessagesByChannel(channel *models2.Channel) ([]models2.ChatMessage, error) {
	endpointUrl := c.getEndpoint(EndpointMessages,
		fmt.Sprintf("/users/ME/conversations/%s/messages",
			url.PathEscape(channel.Id),
		),
	)
	values := endpointUrl.Query()
	values.Add("view", "msnp24Equivalent|supportsMessageProperties")
	values.Add("pageSize", "200")
	values.Add("startTime", "1")
	endpointUrl.RawQuery = values.Encode()

	req, err := c.AuthenticatedRequest("GET", endpointUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	expectedStatusCode := http.StatusOK
	if resp.StatusCode != expectedStatusCode {
		return nil, errors.NewHTTPError(expectedStatusCode, resp.StatusCode, nil)
	}

	jsonBuffer, err := util.GetJSON(resp, c.debugSave)
	if err != nil {
		return nil, err
	}

	var msgResponse models2.MessagesResponse
	dec := json.NewDecoder(jsonBuffer)
	if c.debugDisallowUnknownFields {
		dec.DisallowUnknownFields()
	}
	err = dec.Decode(&msgResponse)
	if err != nil {
		return nil, fmt.Errorf("unable to decode json: %v", err)
	}

	return msgResponse.Messages, err
}

func (c *CSASvc) SendMessage(channel string, message string) error {
	endpointUrl := c.getEndpoint(EndpointMessages,
		fmt.Sprintf("/users/ME/conversations/%s/messages",
			url.PathEscape(channel),
		),
	)

	body := "{\"content\":\"<p>" + message + "</p>\",\"messagetype\":\"RichText/Html\",\"contenttype\":\"text\",\"amsreferences\":[],\"clientmessageid\":\" " + strconv.Itoa(1000000000000000000+rand.Intn(999999999999999999)) + "\",\"imdisplayname\":\"Tom\",\"properties\":{\"importance\":\"\",\"subject\":\"\"}}"

	req, err := c.AuthenticatedRequest("POST", endpointUrl.String(), strings.NewReader(body))
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	expectedStatusCode := http.StatusCreated
	if resp.StatusCode != expectedStatusCode {
		return errors.NewHTTPError(expectedStatusCode, resp.StatusCode, nil)
	}

	return err
}

func (c *CSASvc) ReactToMessage(channel string, message string, emote string) error {
	endpointUrl := c.getEndpoint(EndpointMessages,
		fmt.Sprintf("/users/ME/conversations/%s/messages/%s/properties?name=emotions&replace=true",
			url.PathEscape(channel),
			url.PathEscape(message),
		),
	)

	timeNow := time.Now().UnixMilli()
	body := "{\"emotions\":{\"key\":\"" + emote + "\",\"value\":" + strconv.FormatInt(timeNow, 10) + "}}"

	req, err := c.AuthenticatedRequest("PUT", endpointUrl.String(), strings.NewReader(body))
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	expectedStatusCode := http.StatusOK
	if resp.StatusCode != expectedStatusCode {
		return errors.NewHTTPError(expectedStatusCode, resp.StatusCode, nil)
	}

	return err
}

func (c *CSASvc) RemoveReactionToMessage(channel string, message string, emote string) error {
	endpointUrl := c.getEndpoint(EndpointMessages,
		fmt.Sprintf("/users/ME/conversations/%s/messages/%s/properties?name=emotions",
			url.PathEscape(channel),
			url.PathEscape(message),
		),
	)

	body := "{\"emotions\":{\"key\":\"" + emote + "\"}}"

	req, err := c.AuthenticatedRequest("DELETE", endpointUrl.String(), strings.NewReader(body))
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	expectedStatusCode := http.StatusOK
	if resp.StatusCode != expectedStatusCode {
		return errors.NewHTTPError(expectedStatusCode, resp.StatusCode, nil)
	}

	return err
}

func (c *CSASvc) DeleteMessage(channel string, message string) error {
	endpointUrl := c.getEndpoint(EndpointMessages,
		fmt.Sprintf("/users/ME/conversations/%s/messages/%s?behavior=softDelete",
			url.PathEscape(channel),
			url.PathEscape(message),
		),
	)

	req, err := c.AuthenticatedRequest("DELETE", endpointUrl.String(), nil)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	expectedStatusCode := http.StatusOK
	if resp.StatusCode != expectedStatusCode {
		return errors.NewHTTPError(expectedStatusCode, resp.StatusCode, nil)
	}

	return err
}
