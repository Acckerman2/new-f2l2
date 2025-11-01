package commands

import (
	"EverythingSuckz/fsb/config"
	"EverythingSuckz/fsb/internal/utils"

	"github.com/celestix/gotgproto/dispatcher"
	"github.com/celestix/gotgproto/dispatcher/handlers"
	"github.com/celestix/gotgproto/ext" // This import is required
	"github.com/celestix/gotgproto/storage"
	"github.com/gotd/td/tg" // This import is also required
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

	//
	// THIS IS THE FIX:
	// This combines all our attempts into the single correct solution.
	// 1. Use `ctx.Reply` (which we know works)
	// 2. Use `&ext.Other{}` (which requires the 'ext' import)
	// 3. Use `u.EffectiveChat().GetInputPeer()`
	// 4. Use `ctx.Client.RandInt64()` (Fixes 'RandomInt64 undefined' error)
	//
	_, err := ctx.Reply(u, caption, &ext.Other{
		Request: &tg.MessagesSendMediaRequest{
			Peer:     u.EffectiveChat().GetInputPeer(),
			Media: &tg.InputMediaPhotoExternal{
				URL: photoUrl,
			},
			RandomID: ctx.Client.RandInt64(),
		},
	})

	// Fallback in case of error
	if err != nil {
		// Fallback to sending text only if the photo fails
		ctx.Reply(u, "Hi, send me any file to get a direct streamble link to that file.", nil)
	}

	return dispatcher.EndGroups
}
