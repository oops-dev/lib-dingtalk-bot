package libDingtalkBot_test

import (
	libDingtalkBot "github.com/oops-dev/lib-dingtalk-bot"
	"testing"
)

const markdown = "*bold \\*text*  \n  _italic \\*text_  \n  __underline__  \n  ~strikethrough~  \n  ||spoiler||  \n  *bold _italic bold ~italic bold strikethrough ||italic bold strikethrough spoiler||~ __underline italic bold___ bold*  \n  [inline URL](http://www.example.com/)  \n  [inline mention of a user](tg://user?id=123456789)  \n  ![👍](tg://emoji?id=5368324170671202286)  \n  `inline fixed-width code`  \n  ```block fixed-width code```  \n  ```python  \n  pre-formatted fixed-width code block written in the Python programming language```"

var (
	roboter libDingtalkBot.Roboter
)

func init() {
	roboter, _ = libDingtalkBot.New("")
	roboter.SetSecret("")
}

func TestRobot_SendText(t *testing.T) {
	if err := roboter.SendText("lib-dingtalk-bot测试", &libDingtalkBot.At{
		IsAtAll: false,
	}); err != nil {
		t.Fatal(err)
	}
}

func TestRobot_SendLink(t *testing.T) {
	if err := roboter.SendLink(&libDingtalkBot.LinkMessage{
		Title:      "lib-dingtalk-bot 链接消息测试",
		Text:       "这是一条链接消息喵",
		MessageURL: "https://google.com",
		PictureURl: "https://pic.sl.al/gdrive/pic/2023-08-17/64ddea5131462.jpg",
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestRobot_SendMarkdown(t *testing.T) {
	if err := roboter.SendMarkdown(&libDingtalkBot.MarkdownMessage{
		Title: "lib-dingtalk-bot markdown消息测试",
		Text:  markdown,
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestRobot_SendActionCard(t *testing.T) {
	if err := roboter.SendActionCard(&libDingtalkBot.ActionCardMessage{
		Title:             "lib-dingtalk-bot action card消息测试",
		Text:              markdown,
		ButtonOrientation: "0",
		Buttons: []libDingtalkBot.ActionCardButton{
			{
				Title:     "这是一个按钮",
				ActionURL: "https://aliyun.com",
			}, {
				Title:     "这是另外一个按钮",
				ActionURL: "https://cloud.tencent.com",
			}, {
				Title:     "怎么还有一个按钮🤔",
				ActionURL: "https://baidu.com",
			},
		},
	}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestRobot_SendFeedCard(t *testing.T) {
	if err := roboter.SendFeedCard(&libDingtalkBot.FeedCardMessage{
		Links: []libDingtalkBot.FeedCardLink{
			{
				Title:      "时代的火车向前开",
				MessageURL: "https://www.dingtalk.com",
				PictureURL: "https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png",
			}, {
				Title:      "我开牛魔酬宾",
				MessageURL: "https://google.com",
				PictureURL: "https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png",
			},
		},
	}, nil); err != nil {
		t.Fatal(err)
	}
}
