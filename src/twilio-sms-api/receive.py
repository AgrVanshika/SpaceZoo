import os
import psycopg2
from flask import Flask, request, redirect
from twilio.twiml.messaging_response import MessagingResponse
from twilio.twiml.voice_response import VoiceResponse, Gather
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

@app.route("/voice", methods=['GET', 'POST'])
def answer_call():
    resp = VoiceResponse()
    
    #if the request already has digits, save them to choice and do this shit
    if 'Digits' in request.values:
        choice = request.values['Digits']

        if choice == '1':
            technique = getRandomFromCategory('breathing')
            resp.say(f"You have chosen breathing. {technique[0]}. {technique[1]}")
            return(str(resp))
        elif choice == '2':
            technique = getRandomFromCategory('expression')
            resp.say(f"You have chosen expression. {technique[0]}. {technique[1]}")
            return(str(resp))
        elif choice == '3':
            technique = getRandomFromCategory('meditation')
            resp.say(f"You have chosen meditation. {technique[0]}. {technique[1]}")
            return(str(resp))
        elif choice == '4':
            technique = getRandomFromCategory('reflection')
            resp.say(f"You have chosen reflection. {technique[0]}. {technique[1]}")
            return(str(resp))
        elif choice == '5':
            technique = getRandomFromCategory('relaxation')
            resp.say(f"You have chosen relaxation. {technique[0]}. {technique[1]}")
            return(str(resp))
        elif choice == '6':
            technique = getRandomFromCategory('visualization')
            resp.say(f"You have chosen visualization. {technique[0]}. {technique[1]}")
            return(str(resp))
        else:
            resp.say("Sorry, I didn't understand that choice.")
    gather = Gather(num_digits=1)
    gather.say("Thank you for calling space zoo. For a breathing technique, press 1. For an espressive technique, press 2. For a meditation technique, press 3. For a reflection technique, press 4. For a relaxation technique, press 5. For a visualization technique, press 6.")
    resp.append(gather)
    #loop!
    resp.redirect('/voice')
    
    return str(resp)

if __name__ == "__main__":
    app.run(host='0.0.0.0',debug=False)