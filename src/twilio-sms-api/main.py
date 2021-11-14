# Ensure that you have a creds.py file
from creds import ACCOUNT, TOKEN

from twilio.rest import Client
from sys import argv

client = Client(ACCOUNT, TOKEN)

# func send_sms(to, from, message) 

def send_sms(to, from_, message):
    message = client.messages.create(
                     body= argv[2],
                     from_='+15155178952',
                     to= argv[1]
                 )
    return message

send_sms('+19739807798','+15155178952',"SpaceZoo.tech")
'''{ 
    "account_sid": ACCOUNT,
    "api_version": "2010-04-01",
    "body": message,
    "date_created": "Thu, 13  2015 20:12:31 +0000",
    "date_sent": "Thu, 30 Jul 2015 20:12:33 +0000",
    "date_updated": "Thu, 30 Jul 2015 20:12:33 +0000",
    "direction": "outbound-api",
    "error_code": null,
    "error_message": null,
    "from": from_,
    "messaging_service_sid": "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
    "num_media": "0",
    "num_segments": "1",
    "price": null,
    "price_unit": null,
    "sid": "SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
    "status": "sent",
    "subresource_uris": {
        "media": "/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media.json"
        },
    "to": to,
    "uri": "/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json"
        }
print(message.sid)
send_sms('+19739807798','+15155178952',"SpaceZoo.tech")'''

'''
# https://www.twilio.com/code-exchange/send-sms-message-twilio-phone-number
# https://www.twilio.com/docs/sms/quickstart/python


# try:
#     client = twilio.rest.TwilioRestClient(account_pyp, suth_token)
#     message = client.messages.create(
#         body = "Hello world",
#         from_ = '+3456722222',
#         to = '7832104955'
#         )
# except twilio.TwilioRestException as e:
#     print(e)    
'''
