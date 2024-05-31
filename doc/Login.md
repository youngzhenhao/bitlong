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

POST 开发票： /custodyAccount/invoice/apply
开发票请求参数：
请求头：
Authorization ： "Bearer" token

请求体：
格式 Json
"amount": int 请求发票金额
"memo": string 发票备注信息

返回值：
"invoice": string 发票号码
"error": string 错误信息

### 2.查询发票

POST 查询发票： /custodyAccount/invoice/querybalance
查询发票请求参数：
请求头：
Authorization ： "Bearer" token

请求体：
{
"asset_id":"00"
}

asset_id：string 资产ID,00表示比特币

返回值：
```json
{
    "invoices": [
        {
            "invoice": "lnbcrt1222220n1pn9jlmxpp5f0yjutf4t2z7hsrp826s4q27f964npgjhjfrtva2pzk240weqt2sdq6dysxcmmkv5sx7umnv3shxerpyqcqzzsxqyz5vqsp5hrtme76j03gaaxle3a3tvd83u86va6q6pltcy2fzta7082ju698q9qyyssqmcey7racq2gu03v54j7jujv2fq7ypkqgj74pcvjpv6p9h5r5lfqplqu9c28lv35x4wrvxvw6hdhjjpppnreqk36he2wyfkxyyknmhqgqxxhtrf",
            "asset_id": "00",
            "amount": 122222,
            "status": 0
        }
    ]
}
```

invoice： string 发票号码
asset_id：string 资产ID,预留，00表示比特币
amount：int 发票金额
status： int 发票状态，0表示未支付，1表示已支付，2表示已失效

"error": string 错误信息
### 3.查询余额


POST 查询余额： /custodyAccount/invoice/querybalance
查询余额请求参数：
请求头：
Authorization ： "Bearer" token

请求体：
无

返回值：
"balance": int 账户余额
"error": string 错误信息












