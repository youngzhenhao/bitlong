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

连接到另一个lnd节点

|      函数参数      |  类型  |       用途       |
|:------------------:|:------:|:----------------:|
| pubkey         | string | 节点公钥         |
| host  | string   | 目的主机加端口即套接字 |


| 返回参数 |  类型  |
|:--------:|:------:|
| 是否成功   | boolean |


#### OpenChannel

开通道

> 会尝试向远程对等方打开请求中指定的单一注资通道。
> 用户可以指定确认资金交易的目标区块数，或为资金交易手动设定费率。
> 如果两者都未指定，则使用宽松的区块确认目标。每次 OpenStatusUpdate 都会返回正在进行的通道的待定通道 ID。
> 根据 OpenChannelRequest 中指定的参数，该待定通道 ID 可用于手动推进通道资金流。

|      函数参数      |  类型  |       用途       |
|:------------------:|:------:|:----------------:|
| nodePubkey         | string | 节点公钥         |
| localFundingAmount | long   | 本地注资数量(聪) |

| 返回参数 |  类型  |
|:--------:|:------:|
| 通道id   | string |

#### CloseChannel

关通道

> 试图关闭一个由其通道输出点（ChannelPoint）标识的活动通道。
> 该方法的操作还可以在超时后尝试强制关闭不活动的对等设备。如果请求非强制关闭（合作关闭），
> 用户可以指定关闭交易确认前的目标区块数或手动费率。
> 如果两者都未指定，则使用默认的宽松区块确认目标。

资金交易的输出点（txid:index）。有了这个值，Bob 就能为 Alice 版本的承诺交易生成签名

|      函数参数      |  类型  |       用途       |
|:------------------:|:------:|:----------------:|
| fundingTxidStr          | string | 通道的 txid         |
| outputIndex  | int   | 交易输出点索引index |

| 返回参数 |  类型  |
|:--------:|:------:|
| 是否成功   | boolean |

---

#### PendingChannels

等待中的通道

#### ListChannels

列出通道

#### GetChanInfo

查询通道信息

> 返回指定通道的最新认证网络公告，该通道由通道 ID 标识：一个 8 字节整数，用于唯一标识区块链中交易资金输出的位置

|      函数参数      |  类型  |       用途       |
|:------------------:|:------:|:----------------:|
| chanId          | int | 通道的 id         |


| 返回参数 |  类型  |
|:--------:|:------:|
| 通道详情数据   | string |

#### ChannelBalance

通道余额

> 返回所有开放渠道的资金总额报告，按本地/远程、待结算的本地/远程和未结算的本地/远程余额分类

| 返回参数 |  类型  |
|:--------:|:------:|
| 所有通道余额详情数据   | string |

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


