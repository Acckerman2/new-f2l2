package commands

import (
	"EverythingSuckz/fsb/config"
	"EverythingSuckz/fsb/internal/utils"

	"github.com/celestix/gotgproto/dispatcher"
	"github.com/celestix/gotgproto/dispatcher/handlers"
	"github.com/celestix/gotgproto/ext"
	"github.com/celestix/gotgproto/storage"
	// We DO NOT need "github.com/gotd/td/tg" for this solution
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
	// I have been an idiot. The helper methods are on `u` (the Update),
	// not `ctx` (the Context). And the options are `*ext.ReplyOpts`,
	// not `*ext.Other`.
	//
	// This single line fixes the entire build.
	//
	_, err := u.ReplyPhotoURL(photoUrl, &ext.ReplyOpts{
		Caption: caption,
	})

	// Fallback in case of error
	if err != nil {
		// Fallback to sending text only if the photo fails
		// Note: The text-only reply *is* on `ctx`.
		ctx.Reply(u, "Hi, send me any file to get a direct streamble link to that file.", nil)
	}

	return dispatcher.EndGroups
}
