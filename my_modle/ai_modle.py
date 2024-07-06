from http import HTTPStatus
from dashscope import Application
import json
from pathlib import Path

class QW_MAX:
    def __init__(self, question, path):
        self.question = question
        self.path = path

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
    def __init__(self, question, path):
        self.question = question
        self.path = path

    def history(self):
        content = self.path.read_text()
        history_content = json.loads(content)
        history_response = '\n'.join(history_content)
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
        
    