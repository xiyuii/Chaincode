# 生成大模型调用范式，若要添加直接在add_new_model填写即可
from pathlib import Path
import json

path = Path('lib/data/modle_format.json')
path.write_text(json.dumps({}))

def add_new_model():
    Qwen_max = """

    def history(self):
        content = self.path.read_text()
        history_content = json.loads(content)
        history_response = ''
        for value in history_content.values():
            history_response = f'''{value}
{history_response}'''
        return history_response

    def call_agent_app(self):
        history_response = self.history()

        response = Application.call(
            app_id = self.robot,
            prompt = f"这是以前的回答{history_response}，回答时以此为参考，但不要有任何提及，只回答这个问题：{self.question}",
            api_key = self.api
        )

        if response.status_code != HTTPStatus.OK:
            return response.message
        
        else:
            return response.output

    """

    Chat_GPT_3_5 = """

    def history(self):
        content = self.path.read_text()
        history_content = json.loads(content)
        history_response = ''
        for value in history_content.values():
            history_response = f'''{value}
{history_response}'''
        return history_response
    
    def call_agent_app(self):
        history_response=self.history

        openai.api_key = self.api
        response = openai.ChatCompletion.create(
            model='gpt-3.5-turbo',
            message=[
                {'role': 'system', 'content': self.promot},
                {'role': 'user', 'content': f'这是以前的回答{history_response}，回答时以此为参考，但不要有任何提及，只回答这个问题：{self.question}'},
            ]
        )

        return response['choices'][0]['message']['content']
"""
    # New Model Here

    dict_ai = {}
    dict_ai['Qwen-max'] = Qwen_max
    dict_ai['ChatGpt-3.5'] = Chat_GPT_3_5
    #Add new model
    content = json.dumps(dict_ai, ensure_ascii=False, indent=4)
    path.write_text(content, encoding='utf-8')
