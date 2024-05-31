# Login

## 获取用户名和密码

### 用户名

func GetPublicKey() string 

###  密码

用手机的唯一标识码作为密码。

## 登录流程

### 1. Login 接口

请求参数：

- username：用户名
- password：密码

func Login(username, password string) (string, error)

返回一个token ，需要保存到客户端，每次请求都需要带上这个token。
错误返回空值，错误信息捕捉异常

### 2.refresh
暂时用Login接口代替


# 托管账户：
### 1.开户（开发票）
第一次开发票的时候会在服务器给用户生成一个托管账户

POST 开发票： 
开发票请求参数：
func ApplyInvoiceRequest(amount int64, memo string, token string) ([]byte, error)

amount：开发票金额
memo：开发票备注
token：登录token

返回值：json字节流







