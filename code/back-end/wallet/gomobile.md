# 钱包结构目录说明（初期）

这是初期钱包项目的目录结构说明。该结构设计旨在提供清晰的模块划分，使项目易于理解和扩展。

## 目录结构

```plaintext
wallet
├── address           # 地址实现
│   └── btc.go        # BTC地址实现
├── api               # API接口实现
│   └── wallteApi.go  # 钱包模块功能对外暴露的接口(btc主链资产相关实现)
│   └── star.go       # lnd节点启动引用
├── config            # 项目配置
│   └── apiConfig.go  # 常量模块配置
├── taproot           # Taproot资产相关实现
├── other             # btc主链资产相关API
│   └── otherApi.go   # Api请求逻辑
│   └── otheReq.go    # http请求函数
├── tests             # 测试文件
│   └── addr_test.go  # 地址API暴露测试文件
├── tx                # 离线签名相关实现
├── README.md  
└── go.mod
```


# Gomobile 安装流程

## 1. 安装 Go 环境
- 访问 [Go 官方网站](https://golang.org/dl/) 下载并安装 Go。
- 安装完成后，验证安装结果：运行 `go version` 命令。

## 2. 安装 Android Studio
- 下载并安装 [Android Studio](https://developer.android.com/studio)。
- 在 Android Studio 中安装 SDK，确保安装了 NDK（版本不要超过22）。

## 3. 安装 Java 环境 (JDK)
- 下载并安装 [Java Development Kit (JDK)](https://www.oracle.com/java/technologies/javase-downloads.html)。

## 4. 安装 Gomobile
- 执行以下命令安装 Gomobile 和 Gobind：
```bash
   go install golang.org/x/mobile/cmd/gomobile@latest
```
```bash
  gomobile init
```
```bash
  go install golang.org/x/mobile/cmd/gobind@latest
```
- 需要下载
```bash
go get golang.org/x/mobile/bind
```

## 5. 打包命令
```bash
cd 到需要打包的文件目录
gomobile bind -target android 
```

## 6. 参考信息
- https://blog.51cto.com/u_15588078/6626686
- https://juejin.cn/post/7303749383165804555

## btc免费节点服务商 blockdaemon 
https://www.blockdaemon.com/sg