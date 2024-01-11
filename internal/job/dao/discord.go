package dao

import (
	"backend-go/internal/app/model"
	"backend-go/pkg/log"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"time"
)

func (d *Dao) AirdropSuccessNotice(address string, tokenId string) {
	// 查询用户信息
	var user model.Users
	err := d.db.Model(&model.Users{}).Select("socials").Where("address = ?", address).First(&user).Error
	if err != nil {
		log.Errorv("查询用户失败", zap.String("address", address), zap.Error(err))
		return
	}
	Token := d.c.Discord.Token
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Errorv("error creating Discord session", zap.Error(err))
		return
	}
	// Cleanly close down the Discord session.
	defer dg.Close()
	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages
	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Errorv("error opening connection", zap.Error(err))
		return
	}
	// 拼接信息
	link := fmt.Sprintf("https://polygonscan.com/token/%s?a=%d", d.c.Contract.Badge, tokenID)
	description := fmt.Sprintf("Address: %s \n TokenID: %d \n %s \n <@%s>", address, tokenID, link, gjson.Get(string(user.Socials), "discord.id"))
	footer := discordgo.MessageEmbedFooter{Text: "decertme-bot"}
	embeds := []*discordgo.MessageEmbed{
		{Color: 0x0099FF, Title: "SBT 空投",
			Description: description,
			URL:         fmt.Sprintf("https://decert.me/quests/%d", tokenID),
			Footer:      &footer,
			Timestamp:   time.Now().Format(time.RFC3339),
		},
	}
	msg := discordgo.MessageSend{Embeds: embeds}
	_, err = dg.ChannelMessageSendComplex(d.c.Discord.SuccessChannel, &msg)
	if err != nil {
		log.Errorv("ChannelMessageSendComplex error", zap.Error(err))
	}
}

func (d *Dao) AirdropFailNotice(address string, tokenId string, reason string) {
	Token := d.c.Discord.Token
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Errorv("error creating Discord session", zap.Error(err))
		return
	}
	// Cleanly close down the Discord session.
	defer dg.Close()
	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Errorv("error opening connection", zap.Error(err))
		return
	}
	// 拼接信息
	description := fmt.Sprintf("Address: %s \n TokenID: %d \n Reason: %s", address, tokenID, reason)
	footer := discordgo.MessageEmbedFooter{Text: "decertme-bot"}
	embeds := []*discordgo.MessageEmbed{
		{Color: 0x0099FF, Title: "SBT 空投失败",
			Description: description,
			URL:         fmt.Sprintf("https://decert.me/quests/%d", tokenID),
			Footer:      &footer,
			Timestamp:   time.Now().Format(time.RFC3339),
		},
	}
	msg := discordgo.MessageSend{Embeds: embeds}
	_, err = dg.ChannelMessageSendComplex(d.c.Discord.FailedChannel, &msg)
	if err != nil {
		log.Errorv("ChannelMessageSendComplex error", zap.Error(err))
	}
}
