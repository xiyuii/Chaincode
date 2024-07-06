from flask import Flask, jsonify, request, render_template
import ai_modle
import json
from pathlib import Path
import time

app = Flask(__name__)
path = Path('lib/data/history.json')
local_time = time.localtime()
LOG_IN_TIME = time.strftime("%Y-%m-%d %H:%M:%S", local_time)

def AI_MODLE(option = 'QW_Max', user_input = 'Hello', path = ''):
    if option == 'QW-Max':
        return ai_modle.QW_MAX(user_input, path)

def initialize_path(path):
    if not path.exists():
        path.write_text(json.dumps({}))

initialize_path(path)

@app.route('/')
def controller():
    return render_template('controller.html')

@app.route('/submit')
def api_market():
    return render_template('api_market.html')

@app.route('/send_data', methods=['POST'])
def clear_out():
    data = request.get_json()
    if data:
        path.write_text(json.dumps({}))
        

@app.route('/submit', methods=["POST"])
def submit():
    data = request.get_json()
    option = data.get('modelSelect', '')
    user_input = data.get('user_input', '')
    if user_input == '':
        user_input = 'Hello'
    ai_response =  AI_MODLE(option, user_input, path)
    response = ai_response.call_agent_app()

    response_history = {}
    if path.read_text().strip():
        content = path.read_text()
        response_history = json.loads(content)
    current_time = time.localtime()
    QUESTION_TIME = time.strftime("%Y-%m-%d %H:%M:%S", current_time)
    response_history[QUESTION_TIME] = response['text']

    history_content = json.dumps(response_history)
    path.write_text(history_content)

    final_response = ''
    for date_time, content in response_history.items():
        if date_time > LOG_IN_TIME:
            final_response = f"{final_response}<br><br>{content}"

    return jsonify({'message': final_response})

if __name__ == '__main__':
    app.run(debug = True)