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

	// --- MODIFICATION START ---

	// 1. Define your caption
	caption := "Hi, send me any file to get a direct streamble link to that file."

	// 2. Choose ONE option to send your image
	
	// **Option A: Send a local file**
	// The path must be accessible by your bot's process.
	imagePath := "https://envs.sh/NEV.jpg" // <-- IMPORTANT: Change this path
	_, err := u.ReplyPhotoPath(imagePath, &ext.Other{
		Caption: caption,
	})

	/*
	// **Option B: Send from a URL**
	imageURL := "https://example.com/images/welcome.png" // <-- IMPORTANT: Change this URL
	_, err := u.ReplyPhotoURL(imageURL, &ext.Other{
		Caption: caption,
	})
	*/

	/*
	// **Option C: Send using a Telegram File ID (Most efficient)**
	// Get this by sending your photo to @FileID_Bot or similar
	fileID := "AgACAgUAAxI...<your_file_id>" // <-- IMPORTANT: Change this File ID
	_, err := u.ReplyPhotoID(fileID, &ext.Other{
		Caption: caption,
	})
	*/

	// 3. Add error handling (optional but recommended)
	if err != nil {
		ctx.Log.Errorf("Failed to send start photo: %v", err)
		// Fallback to sending text only if the photo fails
		ctx.Reply(u, caption, nil)
	}

	// --- MODIFICATION END ---

	return dispatcher.EndGroups
}
