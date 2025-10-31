package commands

import (
	"EverythingSuckz/fsb/config"
	"EverythingSuckz/fsb/internal/utils"

	"github.com/celestix/gotgproto/dispatcher"
	"github.com/celestix/gotgproto/dispatcher/handlers"
	"github.com/celestix/gotgproto/ext"
	"github.com/celestix/gotgproto/storage"
)

func (m *command) LoadStart(dispatcher dispatcher.Dispatcher) {
	log := m.log.Named("start")
	defer log.Sugar().Info("Loaded")
	dispatcher.AddHandler(handlers.NewCommand("start", start))
}

func start(ctx *ext.Context, u *ext.Update) error {
	chatId := u.EffectiveChat().GetID()
	peerChatId := ctx.PeerStorage.GetPeerById(chatId)

	if peerChatId.Type != int(storage.TypeUser) {
		return dispatcher.EndGroups
	}

	if len(config.ValueOf.AllowedUsers) != 0 && !utils.Contains(config.ValueOf.AllowedUsers, chatId) {
		ctx.Reply(u, "You are not allowed to use this bot.", nil)
		return dispatcher.EndGroups
	}

	// --- Send image with caption and button ---
	photoUrl := "https://envs.sh/NEV.jpg" // 🔹 Replace with your image URL

	// Inline keyboard button markup
	keyboard := &tg.ReplyInlineMarkup{
		Rows: []tg.KeyboardButtonRow{
			{
				Buttons: []tg.KeyboardButtonClass{
					&tg.KeyboardButtonURL{
						Text: "👑 Owner",
						URL:  "https://t.me/Acckerman_r2", // 🔹 Replace with your link
					},
				},
			},
		},
	}

	// Caption text (with emojis and formatting)
	caption := "💾 *Smart File Stream Bot*\n" +
		"Upload once — get your file’s instant stream link and direct download URL.\n" +
		"Optimized for speed ⚡️ and reliability 🔒."

	// Send the photo with caption and inline button
	_, err := ctx.Client.SendMedia(ctx, &tg.MessagesSendMediaRequest{
		Peer:    ctx.InputPeer(chatId),
		Media:   &tg.InputMediaPhotoExternal{URL: photoUrl},
		Message: caption,
		ReplyMarkup: keyboard,
		ParseMode:   tg.ParseModeMarkdown,
	})
	if err != nil {
		ctx.Reply(u, "Error sending image.", nil)
		return dispatcher.EndGroups
	}

	return dispatcher.EndGroups
}

