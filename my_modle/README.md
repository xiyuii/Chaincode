# 大纲
  目前是一个连接前端和后端的Python代码，使用flask架构。
  ```
  |
  |——templates              //模板文件
  |     |——controller.html  //主页面，可以就行对话
  |     |——api_market.html  //到时候api请求就在这罗列出来
  |
  |——static                 //静态脚本文件
  |     |——style.css        //css文件，风格化
  |     |——scripts.js       //javascript文件，用于协调前后端，实现数据传输和动态刷新
  |
  |——ai_modle.py            //大模型的api，目前用来初始化
  |
  |——app.py                 //主函数，调用flask架构
  ```