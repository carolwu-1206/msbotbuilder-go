// Copyright (c) 2020 InfraCloud Technologies
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package activity

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path"

	"github.com/carolwu-1206/msbotbuilder-go/connector/client"
	"github.com/carolwu-1206/msbotbuilder-go/schema"
	"github.com/pkg/errors"
)

// Response provides functionalities to send activity to the connector service.
type Response interface {
	SendActivity(ctx context.Context, activity schema.Activity) ([]byte, error)
	DeleteActivity(ctx context.Context, activity schema.Activity) ([]byte, error)
	UpdateActivity(ctx context.Context, activity schema.Activity) ([]byte, error)
	GetConversationMember(ctx context.Context, activity schema.Activity) ([]byte, error)
}

const (
	// APIVersion for response URLs
	APIVersion = "v3"

	sendToConversationURL    = "/%s/conversations/%s/activities"
	activityResourceURL      = "/%s/conversations/%s/activities/%s"
	getConversationMemberURL = "/%s/conversations/%s/members/%s"
)

// DefaultResponse is the default implementation of Response.
type DefaultResponse struct {
	Client client.Client
}

// DeleteActivity sends a Delete activity method to the BOT connector service.
func (response *DefaultResponse) DeleteActivity(ctx context.Context, activity schema.Activity) ([]byte, error) {
	u, err := url.Parse(activity.ServiceURL)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to parse ServiceURL %s.", activity.ServiceURL)
	}

	respPath := fmt.Sprintf(activityResourceURL, APIVersion, activity.Conversation.ID, activity.ID)

	// Send activity to client
	u.Path = path.Join(u.Path, respPath)
	resp, err := response.Client.Delete(ctx, *u, activity)
	return resp, errors.Wrap(err, "Failed to delete response.")
}

// SendActivity sends an activity to the BOT connector service.
func (response *DefaultResponse) SendActivity(ctx context.Context, activity schema.Activity) ([]byte, error) {
	u, err := url.Parse(activity.ServiceURL)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to parse ServiceURL %s.", activity.ServiceURL)
	}

	respPath := fmt.Sprintf(sendToConversationURL, APIVersion, activity.Conversation.ID)

	// if ReplyToID is set in the activity, we send reply to that particular activity
	if activity.ReplyToID != "" {
		respPath = fmt.Sprintf(activityResourceURL, APIVersion, activity.Conversation.ID, activity.ReplyToID)
	}

	// Send activity to client
	u.Path = path.Join(u.Path, respPath)
	_, err = response.Client.Post(ctx, *u, activity)

	reqBytes, _ := json.Marshal(activity)

	return reqBytes, errors.Wrap(err, "Failed to send response.")
}

// UpdateActivity sends a Put activity method to the BOT connector service.
func (response *DefaultResponse) UpdateActivity(ctx context.Context, activity schema.Activity) ([]byte, error) {
	u, err := url.Parse(activity.ServiceURL)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to parse ServiceURL %s.", activity.ServiceURL)
	}

	respPath := fmt.Sprintf(activityResourceURL, APIVersion, activity.Conversation.ID, activity.ID)

	// Send activity to client
	u.Path = path.Join(u.Path, respPath)
	resp, err := response.Client.Put(ctx, *u, activity)
	return resp, errors.Wrap(err, "Failed to update response.")
}

// GetConversationMember sends a Get BOT connector service to get conversation member details.
func (response *DefaultResponse) GetConversationMember(ctx context.Context, activity schema.Activity) ([]byte, error) {
	u, err := url.Parse(activity.ServiceURL)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to parse ServiceURL %s.", activity.ServiceURL)
	}

	respPath := fmt.Sprintf(getConversationMemberURL, APIVersion, activity.Conversation.ID, activity.From.ID)

	// Send activity to client
	u.Path = path.Join(u.Path, respPath)
	resp, err := response.Client.Get(ctx, *u)
	return resp, errors.Wrap(err, "Failed to update response.")
}

// NewActivityResponse provides a DefaultResponse implementaton of Response.
func NewActivityResponse(connectorClient client.Client) (Response, error) {
	if connectorClient == nil {
		return nil, errors.New("Invalid connector client for ActivityResponse")
	}

	return &DefaultResponse{connectorClient}, nil
}
