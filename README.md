# 基于区块链的AI大模型使用平台

## 启动方法：

1. 安装docker：
```bash
#下载docker 
# 官方脚本当前已无法下载，使用gitee备份的脚本:
curl -fsSL https://gitee.com/real__cool/fabric_install/raw/main/docker_install.sh | bash -s docker --mirror Aliyun
#添加当前用户到docker用户组 
sudo usermod -aG docker $USER 
newgrp docker 
sudo mkdir -p /etc/docker
#配置docker镜像加速，使用阿里云加速
sudo tee /etc/docker/daemon.json <<-'EOF'
{
	"registry-mirrors": ["https://4wgtxa6q.mirror.aliyuncs.com"]
}
EOF

#重启docker 
sudo systemctl daemon-reload
sudo systemctl restart docker
```

2. 安装go、node、jq：
```bash
#下载二进制包
wget https://golang.google.cn/dl/go1.19.linux-amd64.tar.gz
#将下载的二进制包解压至 /usr/local目录
sudo tar -C /usr/local -xzf go1.19.linux-amd64.tar.gz
mkdir $HOME/go
#将以下内容添加至环境变量 ~/.bashrc
export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export PATH=$GOROOT/bin:$PATH
export PATH=$GOPATH/bin:$PATH
#更新环境变量
source  ~/.bashrc 
#设置代理
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

#下载nvm安装脚本
wget https://gitee.com/real__cool/fabric_install/raw/main/nvminstall.sh
#安装nvm；屏幕输出内容添加环境变量
chmod +x nvminstall.sh
./nvminstall.sh
# 将环境变量写入.bashrc
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion
# 更新环境变量
source  ~/.bashrc
# 安装node16（其他版本测试也可以）
nvm install 16
#换源
npm config set registry https://registry.npmmirror.com

#安装jq 
sudo apt install jq
```

3. 克隆项目：
```bash
git clone https://github.com/xiyuii/Chaincode
```

4. 启动区块链（在blockchain/network目录下）：
```bash
# 仅在首次使用执行：下载Fabric Docker镜像。如果拉取速度过慢或失败请检查是否完成docker换源，或者更换一个其他的镜像源再试。
./install-fabric.sh -f 2.5.6 d 
```
```bash
# 启动区块链网络
./start.sh
```
**如果在启动区块链网络时报错可以使用以下命令：**
```bash
# 执行清理所有的容器指令：
docker rm -f $(docker ps -aq)
```
**然后再重启区块链网络**

5. 修改后端ip，换成本地ip或者云服务器的ip。
```bash
fabric-trace/application/web/.env.production
fabric-trace/application/web/.env.development
fabric-trace/application/web/src/router/index.js
```

6. 启动调用AI大模型部分（在my_modle目录下）：
```bash
# 启用虚拟环境并安装所必需的依赖
./venv.sh

# 启动程序
./start.sh
```

7. 在application/web目录下启动：
```bash
# 仅在首次运行执行：安装依赖
npm install 
```
返回application目录：
```bash
./start.sh
```