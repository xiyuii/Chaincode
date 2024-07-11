# 大纲
  目前是一个连接前端和后端的Python代码，使用flask架构。
  ```
  |
  |——templates                //模板文件
  |     |——controller.html    //主页面，可以就行对话
  |     |——market.html        //到时候api请求就在这罗列出来（现在只有添加）
  |
  |——static                   //静态脚本文件
  |     |——css                //css文件，风格化
  |     |——js                 //javascript文件，用于协调前后端，实现数据传输和动态刷新
  |
  |——ai_modle.py              //大模型的api，目前用来初始化
  |——app.py                   //主函数，调用flask架构
  |——format.py                //添加各种模型的调用规范
  |——scripts_add.py           //更新模型
  |——requirements.txt         //安装必要的python库（python版本3.6及以上）
  |——add_ai_modle.sh          //更新文件脚本
  |
  |——lib                      //到时候放一些链码之类的
       |——data                //目前存储问题记录
  ```

  运行时建议先创建虚拟环境，创建venv虚拟环境如下：
  ```
  cmd
  python3 -m venv <venv name>
  ```
