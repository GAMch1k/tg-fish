from telethon import TelegramClient
from time import sleep
import os, sys
import asyncio

args = sys.argv

api_id = int(args[args.index("-id_api") + 1])
api_hash = args[args.index("-hash_api") + 1]

phone = args[args.index("-phone") + 1]



async def dump():
    sess_name = "phone-" + phone.replace("+", "")
    client = TelegramClient(f'sessions/{sess_name}', api_id, api_hash)
    await client.connect()
    
    if not await client.is_user_authorized():
        print("dump failed")
        return
    
    dialogs = await client.get_dialogs()
    
    os.mkdir(os.path.join("./dumps", sess_name))
    
    for dialog in dialogs:
        dialog_id = dialog.entity.id
        print(f"Processing {dialog_id} {dialog.name}...")
        
        os.mkdir(os.path.join("./dumps", sess_name, f"{dialog_id} {dialog.name}"))
        
        with open(os.path.join("./dumps", sess_name, f"{dialog_id} {dialog.name}", "chat.txt"), 
                  "w", encoding="utf-8") as f:
            f.write("")

        media_id = 0
        async for msg in client.iter_messages(dialog_id):
            try:
                media = ""
                if msg.media:
                    await client.download_media(msg.media, os.path.join("./dumps", sess_name, f"{dialog_id} {dialog.name}", str(media_id)))
                    media = f"PHOTO/VIDEO id - {media_id}"
                    media_id += 1
                
                date = "none"
                if msg.date:
                    date = str(msg.date)
                
                name = dialog.name
                
                with open(
                    os.path.join("./dumps", sess_name, f"{dialog_id} {dialog.name}", "chat.txt"),
                    "r+", encoding="utf-8"
                ) as f:
                    content = f.read()
                    f.seek(0, 0)
                    f.write(f"{date} {name}: {msg.message}\n{media}".rstrip("\r\n") + "\n\n" + content)
                sleep(0.02)
            except Exception as e:
                print(e)
                sleep(2)
    print("Finished")
    await client.disconnect()
                    
        


asyncio.run(dump())
