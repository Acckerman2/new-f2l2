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

	// --- MODIFICATION START (Corrected) ---

	// 1. Define your caption
	caption := "Hi, send me any file to get a direct streamble link to that file."

	// 2. Using your specified URL
	
	/*
	// **Option A: Send a local file**
	imagePath := "path/to/your/start-image.jpg" 
	_, err := ctx.ReplyPhotoPath(u, imagePath, &ext.Other{
		Caption: caption,
	})
	*/

	
	// **Option B: Send from a URL**
	imageURL := "https://envs.sh/NEV.jpg" // <-- Your link is here
	_, err := ctx.ReplyPhotoURL(u, imageURL, &ext.Other{
		Caption: caption,
	})
	

	/*
	// **Option C: Send using a Telegram File ID (Most efficient)**
	fileID := "AgACAgUAAxI...<your_file_id>" 
	_, err := ctx.ReplyPhotoID(u, fileID, &ext.Other{
		Caption: caption,
	})
	*/

	// 3. Add error handling (Corrected)
	if err != nil {
		// **Note: Logging is on `ctx.Client.Log`**
		ctx.Client.Log.Errorf("Failed to send start photo: %v", err)
		// Fallback to sending text only if the photo fails
		ctx.Reply(u, caption, nil)
	}

	// --- MODIFICATION END ---

	return dispatcher.EndGroups
}
