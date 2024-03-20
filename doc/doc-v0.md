## 业务流程

### TODO

- 为每个函数添加接口及详细描述

### 钱包解锁 `unlock`

#### GenSeed

助记词生成


#### InitWallet

初始化钱包

#### UnlockWallet

解锁钱包

1. 接口

```go
func UnlockWallet(password string) bool
```

2. 参数列表：
    - `string` password

3. 返回值：
    - `bool` 是否成功解锁
    
4. 详细解释：



#### ChangePassword

更改密码

#### GetNewAddress

生成新地址

---

#### GetWalletBalance

查询钱包余额

### 通道 `channel`

#### ConnectPeer

连接到节点

#### OpenChannel

开通道

#### CloseChannel

关通道

---

#### PendingChannels

等待中的通道

#### ListChannels

列出通道

#### GetChanInfo

查询通道信息

#### ChannelBalance

通道余额

#### ClosedChannels

已关闭通道

### 发票 `invoice` 收款

#### AddInvoice
```go    
> func AddInvoice(value long) string
```      
生成发票     
根据传入的金额生成闪电发票

| 返回参数    | 类型  | 用途      |
|---------|-----|---------|
| value | int | 需要收款的金额 |

| 返回类型   | 用途    |
|--------|-------|
| string | 生成的发票 |

---

#### ListInvoices       
```go
func ListInvoices() string
```
发票列表     
列出当前账户下所持有的发票     
| 返回类型   | 用途    |
|--------|-------|
| string | 返回发票列表 |


#### lookupInvoice
```go
>func LookupInvoice(rhash string) string
```
查询发票信息

| 返回参数    | 类型     | 用途              |
|---------|--------|-----------------|
| string | string | 需要查询的闪电发票支付hash |

| 返回类型   | 用途          |
|--------|-------------|
| string | 指定闪电发票的详细信息 |


### 付款 `pay`

#### DecodePayReq
>func DecodePayReq(pay_req string) int64   

解码发票  
解码支付请求字符串,返回发票金额

| 返回参数    | 类型     | 用途          |
|---------|--------|-------------|
| pay_req | string | 要解码的支付请求字符串 |

| 返回类型  | 用途                         |
|-------|----------------------------|
| int64 | 解码发票的金额. 0：可支付任意金额。-1：解码错误 |


#### EstimateRouteFee
>EstimateRouteFee(dest string, amtsat int64) string  

计算费用  
允许测试发送到目标节点指定金额是否成功

| 返回参数   | 类型     | 用途       |
|--------|--------|----------|
| dest   | string | 目标节点地址   |
| amtsat | int64  | 要测试发送的金额 |

| 返回类型   | 用途     |
|--------|--------|
| string | 返回测试信息 |

#### SendPaymentSync

>func SendPaymentSync(invoice string) string  

支付发票  
支付闪电发票请求  

| 返回参数    | 类型     | 用途         |
|---------|--------|------------|
| invoice | string | 要支付的闪电发票请求 |

| 返回类型   | 用途                 |
|--------|--------------------|
| string | 返回支付hash,可用于跟踪付款进度 |

#### TrackPaymentV2
>func TrackPaymentV2(payhash string) string

交易信息  
返回由付款哈希值标识的付款的更新流。


| 返回参数    | 类型     | 用途            |
|---------|--------|---------------|
| payhash | string | 要查询的支付的支付hash |

| 返回类型   | 用途     |
|--------|--------|
| string | 返回支付状态 |

#### SendCoins
>func SendCoins(addr string, amount int64) string

发送至链上  
向指定比特币地址发送金额。

| 返回参数   | 类型     | 用途   |
|--------|--------|------|
| addr   | string | 目的地址 |
| amount | string | 发送金额 |

| 返回类型   | 用途     |
|--------|--------|
| string | 返回交易ID |


