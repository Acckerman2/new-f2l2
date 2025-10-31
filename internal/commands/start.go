package commands

import (
	"EverythingSuckz/fsb/config"
	"EverythingSuckz/fsb/internal/utils"

	"github.com/celestix/gotgproto/dispatcher"
	"github.com/celestix/gotgproto/dispatcher/handlers"
	"github.com/celestix/gotgproto/ext"
	"github.com/celestix/gotgproto/storage"
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

	// --- Send image with caption ---
	caption := "Hi, send me any file to get a direct streamble link to that file."
	photoUrl := "https://envs.sh/NEV.jpg" // The URL you provided

	// Use the photoUrl variable to send the photo
	_, err := ctx.ReplyPhotoURL(u, photoUrl, &ext.Other{
		Caption: caption,
	})

	// Add error handling
	if err != nil {
		ctx.Client.Log.Errorf("Failed to send start photo: %v", err)
		// Fallback to sending text only if the photo fails
		ctx.Reply(u, caption, nil)
	}

	return dispatcher.EndGroups
}
