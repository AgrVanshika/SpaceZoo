# Ensure that you have a creds.py file
from creds import ACCOUNT, TOKEN

from twilio.rest import Client

client = Client(ACCOUNT, TOKEN)

# func send_sms(to, from, message)

message = client.messages \
                .create(
                     body="SpaceZoo.tech",
                     from_='+15155178952',
                     to='+1'
                 )

print(message.sid)

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

