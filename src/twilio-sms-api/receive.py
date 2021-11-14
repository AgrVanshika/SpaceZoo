import os
import psycopg2
from flask import Flask, request, redirect
from twilio.twiml.messaging_response import MessagingResponse
import sys
sys.path.append("../../")
from getrandomfromcategory import getRandomFromCategory

app = Flask(__name__)

@app.route('/sms', methods = ['GET', 'POST'])

def sms_reply():
    body = request.values.get('Body', None)
    resp = MessagingResponse()
    if body.lower()[0] == 'b' and 'ing' in body.lower():
        technique = getRandomFromCategory('breathing')
        resp.message(f"{technique[0]}: {technique[1]}")
    elif "expression" in body.lower():
        technique = getRandomFromCategory('expression')
        resp.message(f"{technique[0]}: {technique[1]}")
    elif "meditation" in body.lower():
        technique = getRandomFromCategory('meditation')
        resp.message(f"{technique[0]}: {technique[1]}")
    elif "reflection" in body.lower():
        technique = getRandomFromCategory('relaxation')
        resp.message(f"{technique[0]}: {technique[1]}")
    elif "relaxation" in body.lower():
        technique = getRandomFromCategory('relaxation')
        resp.message(f"{technique[0]}: {technique[1]}")
    elif "visualization" in body.lower():
        technique = getRandomFromCategory('visualization')
        resp.message(f"{technique[0]}: {technique[1]}")
    else:
        resp.message("Specify which type of relaxation technique you would like. The available categories are breathing, expression, meditation, reflection, relaxation, and visualization.")
    
    return str(resp)

if __name__ == "__main__":
    app.run(host='0.0.0.0',debug=False)