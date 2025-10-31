package commands

import (
	"EverythingSuckz/fsb/config"
	"EverythingSuckz/fsb/internal/utils"

	"github.com/celestix/gotgproto/dispatcher"
	"github.com/celestix/gotgproto/dispatcher/handlers"
	"github.com/celestix/gotgproto/ext"
	"github.com/celestix/gotgproto/storage"
)

// start is the handler for /start command
func start(ctx *ext.Context, update *ext.Update) error {
	chat := update.EffectiveChat()
	peer := ctx.PeerStorage.GetPeerById(chat.ID)

	// Only allow user chats
	if peer.Type != int(storage.TypeUser) {
		return dispatcher.EndGroups
	}

	// Restrict access to allowed users if configured
	if len(config.ValueOf.AllowedUsers) != 0 && !utils.Contains(config.ValueOf.AllowedUsers, chat.ID) {
		ctx.Reply(update, "🚫 You are not allowed to use this bot.", nil)
		return dispatcher.EndGroups
	}

	// Send welcome message with inline button and image
	imageUrl := "https://envs.sh/NEV.jpg" // 🔹 Replace with your image
	buttons := [][]ext.InlineKeyboardButton{
		{
			{Text: "👑 Owner", Url: "https://t.me/Acckerman_r2"}, // 🔹 Replace link
		},
	}

	ctx.Bot.SendPhoto(chat.ID, &ext.SendPhotoOpts{
		Photo: imageUrl,
		Caption: "💾 *Smart File Stream Bot*\n" +
			"Upload files and get instant stream links ⚡\n" +
			"Fast • Secure • Reliable 💫",
		ParseMode:   "Markdown",
		ReplyMarkup: &ext.InlineKeyboardMarkup{InlineKeyboard: buttons},
	})

	return dispatcher.EndGroups
}

// LoadStart registers the /start command
func (m *command) LoadStart(dispatcher dispatcher.Dispatcher) {
	log := m.log.Named("start")
	defer log.Sugar().Info("Loaded")
	dispatcher.AddHandler(handlers.NewCommand("start", start))
}
