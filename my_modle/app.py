from flask import Flask, jsonify, request, render_template
from pathlib import Path
from ai_using import AI_MODLE
import format
import scripts_add
import json
import time
import markdown

app = Flask(__name__)
path = Path('lib/data/history.json')  # 加载历史回答
path_modle_format = Path('lib/data/modle_format.json')  # 范式信息
local_time = time.localtime()
LOG_IN_TIME = time.strftime("%Y-%m-%d %H:%M:%S", local_time)  # 加载本地时间，防止打开时加载以前的信息
format.add_new_model()

def initialize_path(path):  # 加载所需文件
    if not path.exists():
        path.write_text(json.dumps({}))

initialize_path(path)

@app.route('/')
def controller():  # 渲染页面
    return render_template('controller.html')

@app.route('/send_data', methods=["POST"])  # 清楚历史信息
def clear_out():
    path.write_text(json.dumps({}))
    return jsonify({'message': 'History cleared'})

@app.route('/register')
def submit_ai():
    return render_template('market.html')

@app.route('/submit_api', methods=["POST"])
def ADD_NEW_MODLE():
    data = request.get_json()
    # name
    INPUT_BASICDATA_NAME = data.get('input1', '')
    # api
    INPUT_BASICDATA_API = data.get('input2', '')
    # robot id
    INPUT_BASICDATA_ROBOTID = data.get('input3', '')
    # promot
    INPUT_BASICDATA_PROMOT = data.get('input4', '')
    # 大模型信息
    INPUT_BASICDATA_MODLE = data.get('dropdown', '')

    # 新的大模型信息
    # 大模型名称
    INPUT_NEWDATA_NAME = data.get('extraModelName', '')
    INPUT_NEWDATA_FORMAT = data.get('extraModelSpec', '')

    # 读取规范
    content = path_modle_format.read_text(encoding='utf-8')
    all_modle = json.loads(content)

    format_selected = all_modle[INPUT_BASICDATA_MODLE]
    new_modle = scripts_add.ADD_MODLE(INPUT_BASICDATA_MODLE, INPUT_BASICDATA_NAME, INPUT_BASICDATA_API,
                                    format_selected, INPUT_BASICDATA_ROBOTID, INPUT_BASICDATA_PROMOT)
    add_model = scripts_add.CREATE_MODEL(INPUT_NEWDATA_NAME, INPUT_NEWDATA_FORMAT)
    # 重新加载
    new_modle.ADD_In_AI()
    new_modle.ADD_In_Main()
    new_modle.ADD_IN_CONTROLLER()
    add_model.add_in_format()
    print('new')
    return jsonify({'message': 'Model added successfully'})

@app.route('/submit', methods=["POST"])  # 发送新加入的模型信息
def submit():
    data = request.get_json()  # 加载用户输入
    option = data.get('modelSelect', '')
    user_input = data.get('user_input', '')  

    if user_input == '':
        user_input = '你好'

    ai_response = AI_MODLE(option, user_input, path)  # ai大模型调用的初始化  
    response = ai_response.call_agent_app()

    response_history = {}

    if path.read_text().strip():
        content = path.read_text(encoding='utf-8')
        response_history = json.loads(content)  # 加载历史信息
    current_time = time.localtime()

    QUESTION_TIME = time.strftime("%Y-%m-%d %H:%M:%S", current_time)  # 将回答信息的时间写入
    response_history[QUESTION_TIME] = response['text']

    history_content = json.dumps(response_history, ensure_ascii=False, indent=4) 
    path.write_text(history_content)

    final_response = ''  # 最后一次的回答，默认为空
    for date_time, content in response_history.items():
        if date_time > LOG_IN_TIME:
            final_response = f"{final_response}\n\n{content}"

    # 将markdown内容渲染为html页面内容
    html_response = markdown.markdown(final_response)
    
    return jsonify({'message': html_response})  # 

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=7880)
