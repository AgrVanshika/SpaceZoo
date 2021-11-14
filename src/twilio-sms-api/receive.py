from flask import Flask, request, redirect
from twilio.twiml.messaging_response import MessagingResponse

app = Flask(__name__)

@app.route('/sms', methods = ['GET', 'POST'])

def sms_reply():
    body = request.values.get('Body', None)
    resp = MessagingResponse()
    if body.lower()[0] == 'b' and 'ing' in body.lower():
        resp.message("breathing technique incoming!")
    elif "expression" in body.lower():
        resp.message("expression technique incoming!")
    elif "meditation" in body.lower():
        resp.message("meditation technique incoming!")
    elif "reflection" in body.lower():
        resp.message("reflection technique incoming!")
    elif "relaxation" in body.lower():
        resp.message("relaxation technique incoming!")
    elif "visualization" in body.lower():
        resp.message("visualization technique incoming!")
    else:
        resp.message("Specify which type of relaxation technique you would like. The available categories are breathing, expression, meditation, reflection, relaxation, and visualization.")
    
    return str(resp)

if __name__ == "__main__":
    app.run(debug=False)