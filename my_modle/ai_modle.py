from http import HTTPStatus
from dashscope import Application
import json
import openai


class QW_MAX:
    def __init__(self, question, path, promot='你是一个有帮助的AI助手'):
        self.question = question
        self.path = path
        self.promot = promot

    def history(self):
        content = self.path.read_text()
        history_content = json.loads(content)
        history_response = ''
        for value in history_content.values():
            history_response = f"{value}\n{history_response}"
        return history_response

    def call_agent_app(self):
        history_response = self.history()

        response = Application.call(
            app_id = 'e7b9e9c8748d41f3972271f9524309a5',
            prompt = f"这是以前的回答{history_response}，回答时以此为参考，但不要有任何提及，只回答这个问题：{self.question}",
            api_key = 'sk-d04d8b67485c4168ab25ea0bae100dcf'
        )

        if response.status_code != HTTPStatus.OK:
            return response.message
        
        else:
            return response.output

class CHATGPT_3_5:
    def __init__(self, question, path, promot='你是一个有帮助的AI助手'):
        self.question = question
        self.path = path
        self.promot = promot

    def history(self):
        content = self.path.read_text()
        history_content = json.loads(content)
        history_response = ''
        for value in history_content.values():
            history_response = f'''{value}
{history_response}'''
        return history_response
    
    def call_agent_app(self):
        history_resposne=self.history()

        openai.api_key = self.api
        resposne = openai.ChatCompletion.create(
            model='gpt-3.5-turbo',
            message=[
                {'role': 'system', 'content': self.promot},
                {'role': 'user', 'content': f'这是以前的回答{history_resposne}，回答时以此为参考，但不要有任何提及，只回答这个问题：{self.question}'},
            ]
        )

        return resposne['choices'][0]['message']['content']
        
