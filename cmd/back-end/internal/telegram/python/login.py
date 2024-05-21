from telethon import TelegramClient
import os, sys
import asyncio

args = sys.argv

api_id = int(args[args.index("-id_api") + 1])
api_hash = args[args.index("-hash_api") + 1]

phone = args[args.index("-phone") + 1]
code = args[args.index("-code") + 1]
hash = args[args.index("-hash") + 1]
password = args[args.index("-pass") + 1]

async def login():
    sess_name = "phone-" + phone.replace("+", "")
    client = TelegramClient(f'sessions/{sess_name}', api_id, api_hash)
    await client.connect()
    
    if await client.is_user_authorized():
        print("logined")
        return
        
    if hash == "0":
        res = await client.send_code_request(phone)
        print(res.phone_code_hash)
        return
    
    if password == "0":
        try:
            await client.sign_in(phone, code, phone_code_hash=hash)
            print("logined")
            return
        except:
            print("2af")
            return
    
    try:
        await client.sign_in(password=password)
        print("logined")
        return
    except:
        print("fail")
        return

asyncio.run(login())
