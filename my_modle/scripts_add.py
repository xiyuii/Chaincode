from pathlib import Path
import re
import csv

path = Path('add_ai_modle.sh')

class ADD_MODLE:
    def __init__(self, base_model, modle_name, modle_api, modle_format, robot_id, modle_promot='你是一个有帮助的AI助手'):
        self.name = modle_name
        self.format = modle_format
        self.api = modle_api
        self.robot_id = robot_id
        self.promot = modle_promot
        self.base_model = base_model
        self.modle = ('QW_MAX', 'ChatGpt-3.5', 'Wenxin')

    def ADD_In_AI(self):
        # 向存放ai模型的文件追加
        added_content = f'''
class {self.name}:
    def __init__(self, question, path, promot):
        self.question = question
        self.path = path
        self.api = '{self.api}'
        self.robot = '{self.robot_id}'
        self.promot = promot
{self.format}

'''
        # 将修改后的文件加入
        # 读取文件内容
        file_path = 'ai_modle.py'  # 替换为你的文件路径

        # 写回文件
        with open(file_path, 'a', encoding='utf-8') as file:
            file.write(added_content)
        

    def ADD_In_Main(self):
        # 向主运行程序内追加
        # 定义要插入的新内容
        new_content = f'''    if option == "{self.name}":
        return ai_modle.{self.name}(user_input, path, promot)
    # Insert into this line    
'''

        # 定义标记
        marker = '# Insert into this line'

        # 读取文件内容
        file_path = 'ai_using.py'  # 替换为你的文件路径
        with open(file_path, 'r', encoding='utf-8') as file:
            content = file.read()

        # 使用正则表达式查找标记并插入新内容
        updated_content = re.sub(f'({marker})', rf'\1\n{new_content}', content)

        # 写回文件
        with open(file_path, 'w', encoding='utf-8') as file:
            file.write(updated_content)
    
    def ADD_IN_CONTROLLER(self):
        # 向controller.html内追加
        # 定义要插入的新内容
        new_content = f'''                <option value='{self.name}'>{self.name}</option>
                    <add_in_here></add_in_here>\n'''

        # 定义标记
        marker = '<add_in_here></add_in_here>'

        # 读取文件内容
        file_path = 'templates/controller.html'  # 替换为你的文件路径
        csv_path = 'lib/data/all_model.csv'
        csv_content = [self.name, self.base_model]
        with open(file_path, 'r') as file:
            content = file.read()

        # 使用正则表达式查找标记并插入新内容
        updated_content = re.sub(f'({marker})', rf'\1\n{new_content}', content)

        # 写回文件
        with open(file_path, 'w', encoding='utf-8') as file:
            file.write(updated_content)

        with open(csv_path, 'a', encoding='utf-8', newline='') as file_new:
            csvwriter = csv.writer(file_new)
            csvwriter.writerow(csv_content)

class CREATE_MODEL:
    def __init__(self, model_name, model_format):
        self.name = model_name
        self.format = model_format

    def add_in_format(self):
        # 向存放ai模型的文件追加
        new_content = f'''{self.name} = """
        {self.format}
        """
    # New Model Here

'''
        # 定义标记
        marker = '# New Model Here'

        # 读取文件内容
        file_path = 'format.py'  # 替换为你的文件路径
        # 读取文件内容
        file_path = 'ai_using.py'  # 替换为你的文件路径
        with open(file_path, 'r', encoding='utf-8') as file:
            content = file.read()

        # 使用正则表达式查找标记并插入新内容
        updated_content = re.sub(f'({marker})', rf'\1\n{new_content}', content)

        # 写回文件
        with open(file_path, 'w', encoding='utf-8') as file:
            file.write(updated_content)


        add_content = f"dict_ai['{self.name}'] = {self.name}"
    # Add new model"
        # 定义标记
        marker = '# Add new model'

        # 读取文件内容
        file_path = 'format.py'  # 替换为你的文件路径
        # 读取文件内容
        file_path = 'ai_using.py'  # 替换为你的文件路径
        with open(file_path, 'r', encoding='utf-8') as file:
            content = file.read()

        # 使用正则表达式查找标记并插入新内容
        updated_content = re.sub(f'({marker})', rf'\1\n{add_content}', content)

        # 写回文件
        with open(file_path, 'w', encoding='utf-8') as file:
            file.write(updated_content)