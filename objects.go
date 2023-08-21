package libDingtalkBot

type MessageType int32

const (
	MessageTypeText MessageType = iota
	MessageTypeLink
	MessageTypeMarkdown
	MessageTypeActionCard
	MessageTypeFeedCard
)

var (
	messageTypeValue = map[string]MessageType{
		"text":       MessageTypeText,
		"link":       MessageTypeLink,
		"markdown":   MessageTypeMarkdown,
		"actionCard": MessageTypeActionCard,
		"feedCard":   MessageTypeFeedCard,
	}
	messageTypeName = map[MessageType]string{
		MessageTypeText:       "text",
		MessageTypeLink:       "link",
		MessageTypeMarkdown:   "markdown",
		MessageTypeActionCard: "actionCard",
		MessageTypeFeedCard:   "feedCard",
	}
)

func (t MessageType) String() string {
	s, ok := messageTypeName[t]
	if !ok {
		return "<none>"
	}
	return s
}

func (t MessageType) MarshalText() ([]byte, error) {
	return string2Bytes(t.String()), nil
}

type DingResponse struct {
	ErrCode    int    `json:"errcode"`
	ErrMessage string `json:"errmsg"`
}

type At struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	AtUserIds []string `json:"atUserIds,omitempty"`
	IsAtAll   bool     `json:"isAtAll,omitempty"`
}

type TextMessage struct {
	Content string `json:"content"`
}

type LinkMessage struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	MessageURL string `json:"messageUrl"`
	PictureURl string `json:"picUrl,omitempty"`
}

type MarkdownMessage struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type ActionCardButton struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

type ActionCardMessage struct {
	Title             string             `json:"title"`
	Text              string             `json:"text"`
	SingleTitle       string             `json:"singleTitle"`
	SingleURL         string             `json:"singleURL"`
	ButtonOrientation string             `json:"btnOrientation,omitempty"`
	Buttons           []ActionCardButton `json:"btns,omitempty"`
}

type FeedCardLink struct {
	Title      string `json:"title"`
	MessageURL string `json:"messageURL"`
	PictureURL string `json:"picURL"`
}

type FeedCardMessage struct {
	Links []FeedCardLink `json:"links"`
}

type Message struct {
	MessageType MessageType `json:"msgtype"`
	At          *At         `json:"at,omitempty"`

	Text       *TextMessage       `json:"text,omitempty"`
	Link       *LinkMessage       `json:"link,omitempty"`
	Markdown   *MarkdownMessage   `json:"markdown,omitempty"`
	ActionCard *ActionCardMessage `json:"actionCard,omitempty"`
	FeedCard   *FeedCardMessage   `json:"feedCard,omitempty"`
}
