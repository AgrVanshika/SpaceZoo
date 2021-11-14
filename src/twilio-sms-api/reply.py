from flask import Flask, request, redirect
from twilio.twiml.messaging_response import MessagingResponse

app = Flask(__name__)

@app.route('/sms', methods = ['GET', 'POST'])
def sms_reply():
    body = request.values.get('Body',None)
    resp = MessagingResponse()
    if body == "Hello":
        resp.message("Hi!")
    elif body == "bye":
        resp.message("seeya!")
    return str(resp)
if __name__ == "__main__":
    app.run(debug=True)