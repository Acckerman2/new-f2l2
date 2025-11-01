from pyrogram import filters
from pyrogram.types import Message, InlineKeyboardMarkup, InlineKeyboardButton  # <-- ADDED IMPORTS

from WebStreamer.vars import Var 
from WebStreamer.bot import StreamBot

@StreamBot.on_message(filters.command(["start", "help"]) & filters.private)
async def start(_, m: Message):
    if Var.ALLOWED_USERS and not ((str(m.from_user.id) in Var.ALLOWED_USERS) or (m.from_user.username in Var.ALLOWED_USERS)):
        return await m.reply(
            "You are not in the allowed list of users who can use me. \
            Check <a href='https://github.com/EverythingSuckz/TG-FileStreamBot#optional-vars'>this link</a> for more info.",
            disable_web_page_preview=True, quote=True
        )
    
    await m.reply_photo(
        photo="https://envs.sh/NEV.jpg",
        caption="✨ ʜɪ ɪ'ᴍ Sydney Sweeney! 📁🔗\n\n"
                "🚀 ᴜᴘʟᴏᴀᴅ ᴀɴʏ ꜰɪʟᴇ ᴀɴᴅ ɢᴇᴛ ɪɴꜱᴛᴀɴᴛ ᴅɪʀᴇᴄᴛ ʟɪɴᴋꜱ 🌐\n\n"
                "💎 ꜰᴀꜱᴛ ⚡ | ꜱᴇᴄᴜʀᴇ 🔒 | ᴇᴀꜱʏ ᴛᴏ ᴜꜱᴇ 💫\n\n"
                "💬 ᴊᴜꜱᴛ ꜱᴇɴᴅ ᴀ ᴘʜᴏᴛᴏ, ᴠɪᴅᴇᴏ, ᴏʀ ᴅᴏᴄ — ᴀɴᴅ ɪ'ʟʟ ʜᴀɴᴅʟᴇ ᴛʜᴇ ʀᴇꜱᴛ 😎",
        
        # --- THIS IS THE NEW PART ---
        reply_markup=InlineKeyboardMarkup(
            [
                [InlineKeyboardButton("👑 Owner", url="https://t.me/Acckerman_r2")]
                # You can change the URL to your own Telegram link
            ]
        )
        # ----------------------------
    )
