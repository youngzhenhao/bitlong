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

生成发票

---

#### ListInvoices

发票列表

#### lookupInvoice

查询发票信息

### 付款 `pay`

#### DecodePayReq

解码发票  
解码支付请求字符串,返回发票金额

| 返回参数    | 类型     | 用途          |
|---------|--------|-------------|
| pay_req | string | 要解码的支付请求字符串 |

| 返回类型  | 用途                         |
|-------|----------------------------|
| int64 | 解码发票的金额. 0：可支付任意金额。-1：解码错误 |


#### EstimateRouteFee

计算费用

#### SendPaymentV2

支付发票

#### TrackPaymentV2

交易追踪

---

