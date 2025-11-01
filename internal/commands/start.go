package commands

import (
	"EverythingSuckz/fsb/config"
	"EverythingSuckz/fsb/internal/utils"

	"github.com/celestix/gotgproto/dispatcher"
	"github.com/celestix/gotgproto/dispatcher/handlers"
	"github.com/celestix/gotgproto/ext"
	"github.com/celestix/gotgproto/storage"
	"github.com/gotd/td/tg" // This import is still required
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

	// --- Send image with caption (This is the correct way for v1.0.0-beta18) ---
	caption := "Hi, send me any file to get a direct streamble link to that file."
	photoUrl := "https://envs.sh/NEV.jpg" // The URL you provided

	// 1. Create the media object for the photo URL
	media := &tg.InputMediaPhotoExternal{
		URL: photoUrl,
	}

	// 2. Create the SendMessage parameters object.
	// The ctx.Reply function expects this as its *third* argument.
	params := &tg.MessagesSendMessage{
		Media: media,
	}

	// 3. Call ctx.Reply.
	// The 2nd argument (caption) will be used as the message.
	// The 3rd argument (params) provides the media.
	_, err := ctx.Reply(u, caption, params)

	// 4. Handle errors
	if err != nil {
		// Just return the error. Do not use ctx.Client.Log
		return err
	}

	return dispatcher.EndGroups
}
