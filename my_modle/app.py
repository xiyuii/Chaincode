from flask import Flask, jsonify, request, render_template
from pathlib import Path
import format
import scripts_add
import ai_modle
import json
import time
import markdown
import re

app = Flask(__name__)
path = Path('lib/data/history.json')  # 存放历史问答
path_modle_format = Path('lib/data/modle_format.json')  # 存放范式
local_time = time.localtime()
LOG_IN_TIME = time.strftime("%Y-%m-%d %H:%M:%S", local_time)  # 记录当前打开浏览器的时间，防止载入之前的数据
format.add_new_model()

def AI_MODLE(option='QW_Max', user_input='Hello', path=''):  # 选择模型
    if option == 'QW-Max':
        return ai_modle.QW_MAX(user_input, path)
    # Insert into this line


def initialize_path(path):  # 初始化history文件
    if not path.exists():
        path.write_text(json.dumps({}))

initialize_path(path)

@app.route('/')
def controller():  # 渲染界面
    return render_template('controller.html')

@app.route('/send_data', methods=["POST"])  # 消除历史问答
def clear_out():
    path.write_text(json.dumps({}))
    return jsonify({'message': 'History cleared'})

@app.route('/register')
def submit_ai():
    return render_template('market.html')

@app.route('/submit_api', methods=["POST"])
def ADD_NEW_MODLE():
    data = request.get_json()
    # 名称
    INPUT_BASICDATA_NAME = data.get('input1', '')
    # api
    INPUT_BASICDATA_API = data.get('input2', '')
    # robot id
    INPUT_BASICDATA_ROBOTID = data.get('input3', '')
    # 模型选择
    INPUT_BASICDATA_MODLE = data.get('dropdown', '')

    # 读取范式
    content = path_modle_format.read_text()
    all_modle = json.loads(content)

    format_selected = all_modle[INPUT_BASICDATA_MODLE]
    new_modle = scripts_add.ADD_MODLE(INPUT_BASICDATA_NAME, INPUT_BASICDATA_API,
                                    format_selected, INPUT_BASICDATA_ROBOTID)
    # 添加
    new_modle.ADD_In_AI()
    new_modle.ADD_In_Main()
    print('new')
    return jsonify({'message': 'Model added successfully'})

@app.route('/submit', methods=["POST"])  # 调用api实现本地化
def submit():
    data = request.get_json()  # 选择大模型
    option = data.get('modelSelect', '')
    user_input = data.get('user_input', '')  # 接受用户输入

    if user_input == '':
        user_input = 'Hello'

    ai_response = AI_MODLE(option, user_input, path)  # 实例化AI大模型
    response = ai_response.call_agent_app()  # 输出

    response_history = {}

    if path.read_text().strip():
        content = path.read_text()
        response_history = json.loads(content)  # 读取内容
    current_time = time.localtime()

    QUESTION_TIME = time.strftime("%Y-%m-%d %H:%M:%S", current_time)  # 以时间作为键值
    response_history[QUESTION_TIME] = response['text']

    history_content = json.dumps(response_history)  # 将回答存入history.json中
    path.write_text(history_content)

    final_response = ''  # 仅加载打开浏览器时的所有回答
    for date_time, content in response_history.items():
        if date_time > LOG_IN_TIME:
            final_response = f"{final_response}\n\n{content}"

    # 将Markdown文本转换为HTML
    html_response = markdown.markdown(final_response)
    
    return jsonify({'message': html_response})  # 将结果返回html

if __name__ == '__main__':
    app.run(debug=True)
