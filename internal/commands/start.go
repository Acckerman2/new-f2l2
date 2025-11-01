package commands

import (
	"EverythingSuckz/fsb/config"
	"EverythingSuckz/fsb/internal/utils"

	"github.com/celestix/gotgproto/dispatcher"
	"github.com/celestix/gotgproto/dispatcher/handlers"
	"github.com/celestix/gotgproto/ext"
	"github.com/celestix/gotgproto/storage"
	"github.com/gotd/td/tg" // <-- This import is required
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

	// --- Send image with caption (This is the correct way) ---
	
	// This is your caption
	caption := "Hi, send me any file to get a direct streamble link to that file."
	
	// This is your image URL
	photoUrl := "https://envs.sh/NEV.jpg" 

	// This function sends the photo AND the caption.
	// Telegram automatically puts the photo above the caption.
	_, err := ctx.Reply(u, caption, &ext.Other{
		Media: &tg.InputMediaPhotoExternal{
			URL: photoUrl,
		},
	})

	// Just return the error. Do not use ctx.Client.Log
	if err != nil {
		return err
	}

	return dispatcher.EndGroups
}
