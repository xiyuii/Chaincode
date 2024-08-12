import subprocess
from pathlib import Path
import re

path = Path('add_ai_modle.sh')

class ADD_MODLE:
    def __init__(self, modle_name, modle_api, modle_format, robot_id):
        self.name = modle_name
        self.format = modle_format
        self.api = modle_api
        self.robot_id = robot_id
        self.modle = ('QW_MAX', 'ChatGpt-3.5', 'Wenxin')

    def ADD_In_AI(self):
        # 向存放ai模型的文件追加
        # 修改add_ai_modle.sh内的内容
        added_content = f'''cat <<EOT >>ai_modle.py
class {self.name}:
    def __init__(self, question, path):
        self.question = question
        self.path = path
        self.api = '{self.api}'
        self.robot = '{self.robot_id}'
{self.format}
'''
        # 将修改后的文件加入并执行
        path = Path('add_ai_modle.sh')
        path.write_text(added_content)
        subprocess.run(['sudo', 'chmod', '+x', 'add_ai_modle.sh'])
        subprocess.run(['bash', 'add_ai_modle.sh'],
            capture_output = True, text = True)


    def ADD_In_Main(self):
        # 向主运行程序内追加
        # 定义要插入的新内容
        new_content = f'''    if option == ai_modle.{self.name}(user_input, path):
        return ai_modle.{self.name}()
        # Insert into this line\n   '''

        # 定义标记
        marker = '# Insert into this line'

        # 读取文件内容
        file_path = 'app.py'  # 替换为你的文件路径
        with open(file_path, 'r') as file:
            content = file.read()

        # 使用正则表达式查找标记并插入新内容
        updated_content = re.sub(f'({marker})', rf'\1\n{new_content}', content)

        # 写回文件
        with open(file_path, 'w') as file:
            file.write(updated_content)
    
    def ADD_IN_CONTROLLER(self):
        # 向controller.html内追加
        # 定义要插入的新内容
        new_content = f'''            <option value='{self.name}'>{self.name}</option>
                <add_in_here></add_in_here>\n'''

        # 定义标记
        marker = '<add_in_here></add_in_here>'

        # 读取文件内容
        file_path = 'templates/controller.html'  # 替换为你的文件路径
        with open(file_path, 'r') as file:
            content = file.read()

        # 使用正则表达式查找标记并插入新内容
        updated_content = re.sub(f'({marker})', rf'\1\n{new_content}', content)

        # 写回文件
        with open(file_path, 'w') as file:
            file.write(updated_content)

    def ADD_IN_MARKET(self):
        # 向market.html中追加
        # 定义要插入的新内容
        new_content = f'''            <option value='{self.name}'>{self.name}</option>
                <add_in_here></add_in_here>\n'''

        # 定义标记
        marker = '<add_in_here></add_in_here>'

        # 读取文件内容
        file_path = 'templates/market.html'  # 替换为你的文件路径
        with open(file_path, 'r') as file:
            content = file.read()

        # 使用正则表达式查找标记并插入新内容
        updated_content = re.sub(f'({marker})', rf'\1\n{new_content}', content)

        # 写回文件
        with open(file_path, 'w') as file:
            file.write(updated_content)
