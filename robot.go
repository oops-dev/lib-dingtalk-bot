package libDingtalkBot

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var (
	invokeURL = "https://oapi.dingtalk.com/robot/send"
)

func SetInvokeURL(u string) {
	invokeURL = u
}

type Roboter interface {
	SetSecret(secret string)
	SendMessage(msg *Message) error
	SendText(content string, at *At) error
	SendLink(link *LinkMessage, at *At) error
	SendMarkdown(markdown *MarkdownMessage, at *At) error
	SendActionCard(actionCard *ActionCardMessage, at *At) error
	SendFeedCard(feedCard *FeedCardMessage, at *At) error
}

type robot struct {
	webhook url.URL
	secret  []byte
}

func (r *robot) getURL() string {
	if r.secret == nil {
		return r.webhook.String()
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	query := r.webhook.Query()
	query.Set("timestamp", timestamp)
	query.Set("sign", timestampSign([]byte(timestamp), r.secret))

	u := r.webhook
	u.RawQuery = query.Encode()

	return u.String()
}

func (r *robot) request(msg *Message) error {
	buf := &bytes.Buffer{}
	if err := jsoniter.ConfigFastest.NewEncoder(buf).Encode(msg); err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, r.getURL(), buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var dr DingResponse
	if err = jsoniter.ConfigFastest.NewDecoder(resp.Body).Decode(&dr); err != nil {
		return err
	}

	if dr.ErrCode != 0 {
		return fmt.Errorf("lib-dingtalk-bot: request failed: %s", dr.ErrMessage)
	}

	return nil
}

func (r *robot) SetSecret(secret string) {
	r.secret = []byte(secret)
}

func (r *robot) SendMessage(msg *Message) error {
	return r.request(msg)
}

func (r *robot) SendText(content string, at *At) error {
	return r.request(&Message{
		MessageType: MessageTypeText,
		At:          at,
		Text: &TextMessage{
			Content: content,
		},
	})
}

func (r *robot) SendLink(link *LinkMessage, at *At) error {
	return r.request(&Message{
		MessageType: MessageTypeLink,
		At:          at,
		Link:        link,
	})
}

func (r *robot) SendMarkdown(markdown *MarkdownMessage, at *At) error {
	return r.request(&Message{
		MessageType: MessageTypeMarkdown,
		At:          at,
		Markdown:    markdown,
	})
}

func (r *robot) SendActionCard(actionCard *ActionCardMessage, at *At) error {
	return r.request(&Message{
		MessageType: MessageTypeActionCard,
		At:          at,
		ActionCard:  actionCard,
	})
}

func (r *robot) SendFeedCard(feedCard *FeedCardMessage, at *At) error {
	return r.request(&Message{
		MessageType: MessageTypeFeedCard,
		At:          at,
		FeedCard:    feedCard,
	})
}

func New(accessToken string) (Roboter, error) {
	u, err := url.Parse(invokeURL)
	if err != nil {
		return nil, err
	}
	q := url.Values{}
	q.Set("access_token", accessToken)
	u.RawQuery = q.Encode()

	return &robot{
		webhook: *u,
		secret:  nil,
	}, nil
}
