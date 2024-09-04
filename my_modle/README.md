# 大纲
  使用Python完成此部分，使用flask架构。
  ```
  |
  |——templates                         //模板文件
  |     |——controller.html             //主页面，可以就行对话
  |     |——market.html                 //到时候api请求就在这罗列出来（现在只有添加）
  |     |——api_market.html             //所有大模型罗列出来，理论上不在我的代码里面（放里面为了测试）
  |
  |——static                            //静态脚本文件
  |     |——css                         //css文件，风格化
  |     |——js                          //javascript文件，用于协调前后端，实现数据传输和动态刷新
  |
  |——ai_modle.py                       //大模型的api，目前用来初始化
  |——app.py                            //主函数，调用flask架构
  |——format.py                         //添加各种模型的调用规范
  |——scripts_add.py                    //更新模型
  |——requirements.txt                  //安装必要的python库（python版本3.6及以上，最好在python3.12以下，可能版本太高会有一些问题）
  |——start.sh                          //启动脚本
  |
  |——lib                               
  |    |——data                         //存放一些信息
  |        |——history.json             //存放历史记录
  |        |——modle_format.json        //存放一些调用规范
  ```

  start.sh中的运行内容如下：<br>
  + 运行时先创建虚拟环境，创建venv虚拟环境如下：
  ```
  cmd
  python3 -m venv <venv name>
  ```
  + 进入虚拟环境：
  ```
  cmd
  source .venv/bin/activate
  ``` 
  + 下载所需要的包：
  ```
  cmd
  pip install -r requirements.txt
  ```
  + 运行python脚本


  使用个平台的api-key来进行对大模型的调用，目前只能进行生成文本的大模型，有一些局限性。

  页面支持使用选择的模型进行对话，可以对历史的对话记录进行检索，也可以删去历史信息。

  “我的AI”可以加入自己创建的一些已经创建好的大模型，以便后续在选择窗口中进行选择。

  目前的局限是这一部分并未使用vue、数据库等，所以生成的页面对于每个用户都是一样的。
  
