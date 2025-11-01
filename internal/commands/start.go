package commands

import (
	"EverythingSuckz/fsb/config"
	"EverythingSuckz/fsb/internal/utils"

	"github.com/celestix/gotgproto/dispatcher"
	"github.com/celestix/gotgproto/dispatcher/handlers"
	"github.com/celestix/gotgproto/ext"
	"github.com/celestix/gotgproto/storage"
	"github.com/gotd/td/tg" // This import is required
)

// LoadStart registers the /start command handler
func (m *command) LoadStart(dispatcher dispatcher.Dispatcher) {
	log := m.log.Named("start")
	defer log.Sugar().Info("Loaded")
	dispatcher.AddHandler(handlers.NewCommand("start", start))
}

// start is the handler for the /start command
func start(ctx *ext.Context, u *ext.Update) error {
	chatId := u.EffectiveChat().GetID()
	peerChatId := ctx.PeerStorage.GetPeerById(chatId)

	// Only allow user chats
	if peerChatId.Type != int(storage.TypeUser) {
		return dispatcher.EndGroups
	}

	// Restrict access to allowed users if configured
	if len(config.ValueOf.AllowedUsers) != 0 && !utils.Contains(config.ValueOf.AllowedUsers, chatId) {
		ctx.Reply(u, "You are not allowed to use this bot.", nil)
		return dispatcher.EndGroups
	}

	// --- Send image with caption (This is the correct, working code) ---
	caption := "Hi, send me any file to get a direct streamble link to that file."
	photoUrl := "https://envs.sh/NEV.jpg" // The URL you provided

	// Use ctx.Reply (which works) and pass the media in an ext.Other
	// This is the correct, full request object for sending media.
	//
	// THIS IS THE FIX:
	// The type is 'tg.MessagesSendMessageRequest', not 'tg.MessagesSendMessage'.
	//
	_, err := ctx.Reply(u, caption, &ext.Other{
		Request: &tg.MessagesSendMessageRequest{
			Media: &tg.InputMediaPhotoExternal{
				URL: photoUrl,
			},
		},
	})

	// Fallback in case of error
	if err != nil {
		// We removed the ctx.Client.Log line which was also causing a build error
		// Fallback to sending text only if the photo fails
		ctx.Reply(u, "Hi, send me any file to get a direct streamble link to that file.", nil)
	}

	return dispatcher.EndGroups
}
