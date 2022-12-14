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

/*
 * Microsoft Bot Connector API - v3.0
 *
 * The Bot Connector REST API allows your bot to send and receive messages to channels configured in the  [Bot Framework Developer Portal](https://dev.botframework.com). The Connector service uses industry-standard REST  and JSON over HTTPS.    Client libraries for this REST API are available. See below for a list.    Many bots will use both the Bot Connector REST API and the associated [Bot State REST API](/en-us/restapi/state). The  Bot State REST API allows a bot to store and retrieve state associated with users and conversations.    Authentication for both the Bot Connector and Bot State REST APIs is accomplished with JWT Bearer tokens, and is  described in detail in the [Connector Authentication](/en-us/restapi/authentication) document.    # Client Libraries for the Bot Connector REST API    * [Bot Builder for C#](/en-us/csharp/builder/sdkreference/)  * [Bot Builder for Node.js](/en-us/node/builder/overview/)  * Generate your own from the [Connector API Swagger file](https://raw.githubusercontent.com/Microsoft/BotBuilder/master/CSharp/Library/Microsoft.Bot.Connector.Shared/Swagger/ConnectorAPI.json)    ?? 2016 Microsoft
 *
 * API version: v3
 * Contact: botframework@microsoft.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package schema

// AudioCard - Audio card
type AudioCard struct {

	// Title of this card
	Title string `json:"title,omitempty"`

	// Subtitle of this card
	Subtitle string `json:"subtitle,omitempty"`

	// Text of this card
	Text string `json:"text,omitempty"`

	Image ThumbnailURL `json:"image,omitempty"`

	// Media URLs for this card. When this field contains more than one URL, each URL is an alternative format of the same content.
	Media []MediaURL `json:"media,omitempty"`

	// Actions on this card
	Buttons []CardAction `json:"buttons,omitempty"`

	// This content may be shared with others (default:true)
	Shareable bool `json:"shareable,omitempty"`

	// Should the client loop playback at end of content (default:true)
	Autoloop bool `json:"autoloop,omitempty"`

	// Should the client automatically start playback of media in this card (default:true)
	Autostart bool `json:"autostart,omitempty"`

	// Aspect ratio of thumbnail/media placeholder. Allowed values are \"16:9\" and \"4:3\"
	Aspect string `json:"aspect,omitempty"`

	// Describes the length of the media content without requiring a receiver to open the content. Formatted as an ISO 8601 Duration field.
	Duration string `json:"duration,omitempty"`

	// Supplementary parameter for this card
	Value map[string]interface{} `json:"value,omitempty"`
}
