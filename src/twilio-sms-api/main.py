# Ensure that you have a creds.py file
from creds import ACCOUNT, TOKEN

from twilio.rest import Client
from sys import argv

client = Client(ACCOUNT, TOKEN)

# func send_sms(to, from, message) 

def send_sms():
    message = client.messages.create(
                     body= argv[2],
                     from_='+15155178952',
                     to= argv[1]
                 )
    return message

send_sms()