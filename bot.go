package main

import (
	"log"
	// "os"
	"strings"

	"github.com/src-d/flamingo"
	"github.com/src-d/flamingo/slack"
	"github.com/jpcano/thiago/thiago"
)

type helloController struct{}

func (c *helloController) HandleIntro(bot flamingo.Bot, channel flamingo.Channel) error {
	_, err := bot.Say(flamingo.NewOutgoingMessage("Hey! I am Thiago, a bot for resourcing.\nHi"))
	return err
}

func (c *helloController) CanHandle(msg flamingo.Message) bool {
	return strings.ToLower(strings.TrimSpace(msg.Text)) == "hello"
}

func (c *helloController) Handle(bot flamingo.Bot, msg flamingo.Message) error {
	if _, err := bot.Say(flamingo.NewOutgoingMessage("hello!")); err != nil {
		return err
	}

	_, resp, err := bot.Ask(flamingo.NewOutgoingMessage("how are you?"))
	if err != nil {
		return err
	}

	text := strings.ToLower(strings.TrimSpace(resp.Text))
	if text == "good" || text == "fine" {
		if _, err := bot.Say(flamingo.NewOutgoingMessage("i'm glad!")); err != nil {
			return err
		}
	} else {
		if _, err := bot.Say(flamingo.NewOutgoingMessage(":(")); err != nil {
			return err
		}
	}

	return nil
}


type subscribeController struct{}


func (c *subscribeController) CanHandle(msg flamingo.Message) bool {
	return strings.ToLower(strings.TrimSpace(msg.Text)) == "subscribe"
}

func (c *subscribeController) Handle(bot flamingo.Bot, msg flamingo.Message) error {
	_, resp, err := bot.Ask(flamingo.NewOutgoingMessage("which languages do you want to subscribe to?"))
	if err != nil {
		return err
	}

	langs := strings.Split(resp.Text, " ")

	if err := thiago.Subscribe(msg.User.Username, langs); err != nil {
		return err
	}
	return nil
}

type findController struct{}


func (c *findController) CanHandle(msg flamingo.Message) bool {
	return strings.ToLower(strings.TrimSpace(msg.Text)) == "find"
}

func (c *findController) Handle(bot flamingo.Bot, msg flamingo.Message) error {
	_, resp, err := bot.Ask(flamingo.NewOutgoingMessage("which languages?"))
	if err != nil {
		return err
	}

	langs := strings.Split(resp.Text, " ")

	names, err := thiago.FindSubscriberByTags(langs)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := bot.Say(flamingo.NewOutgoingMessage("The subscribers to `" + strings.Join(langs[:],", ") + "` are: `" + strings.Join(names[:],", ") + "`")); err != nil {
		return err
	}
	return nil
}

type publishController struct{}


func (c *publishController) CanHandle(msg flamingo.Message) bool {
	return strings.ToLower(strings.TrimSpace(msg.Text)) == "publish"
}

func (c *publishController) Handle(bot flamingo.Bot, msg flamingo.Message) error {
	_, resp, err := bot.Ask(flamingo.NewOutgoingMessage("which "))
	if err != nil {
		return err
	}

	langs := strings.Split(resp.Text, " ")

	names, err := thiago.FindSubscriberByTags(langs)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := bot.Say(flamingo.NewOutgoingMessage("The subscribers to `" + strings.Join(langs[:],", ") + "` are: `" + strings.Join(names[:],", ") + "`")); err != nil {
		return err
	}
	return nil
}
func main() {
	// token := os.Getenv("SLACK_TOKEN")
	// id := os.Getenv("BOT_ID")
	token := "xoxb-71041051398-hMnp7KB41pDJ5DQdjwYWG1Nx"
	id := "B2314H38U"
	client := slack.NewClient(token, slack.ClientOptions{
		Debug: true,
	})

	hello := &helloController{}
	subscribe := &subscribeController{}
	find := &findController{}
	publish := &publishController{}
	client.AddController(hello)
	client.AddController(subscribe)
	client.AddController(find)
	client.AddController(publish)
	client.AddBot(id, token)

	log.Fatal(client.Run())
}
