# 业务流程

## 钱包解锁 `unlock`

### GenSeed

助记词生成

1.  接口
```go
func GenSeed() [24]string
```

2.  参数列表

|返回类型|用途|
|----|----|
|string[24]|助记词|

3.  详细解释

是用于实例化新 lnd 实例的第一个方法。该方法允许调用者根据可选的口令生成新的加密种子。如果提供了口令，则需要口令来解密密码种子，以显示内部钱包种子。

用户获得并验证密码种子后，应使用 InitWallet 方法提交新生成的种子，并创建钱包。

### InitWallet

初始化钱包

1. 接口

```go
func InitWallet(seed [24]string, password string) bool
```

2. 参数列表

| 输入参数 | 类型 | 用途|
|--------|----|----|
| seed | string[24] | 助记词 |
| password | string | 钱包密码 |

| 返回类型 | 用途 |
|--------|---------|
|boolean|是否初始化成功|

3. 详细解释

首次启动时使用，用于完全初始化守护进程及其内部钱包。至少必须提供一个钱包密码。这将用于加密磁盘上的敏感资料。

在恢复情况下，用户还可以指定自己的密码和口令。如果设置了该密码，守护进程就会使用之前的状态来初始化其内部钱包。

或者，也可以使用 GenSeed RPC 来获取种子，然后将其提交给用户。经用户验证后，可将种子输入此 RPC，以提交新钱包。

### UnlockWallet

解锁钱包

1. 接口

```go
func UnlockWallet(password string) bool
```

2. 参数列表：

| 输入参数   | 类型     | 用途       |
|--------|--------|----------|
|  password  | string | 解锁钱包密码   |

| 返回类型  | 用途       |
|--------|---------|
|  bool  | 是否解锁钱包   |
 
3. 详细解释：

在lnd启动时，UnlockWallet使用提供的解锁密码解锁钱包数据库。


### ChangePassword

更改密码

1. 接口
```go
func ChangePassword(currentPassword, newPassword string) bool
```
2. 参数列表

|输入参数|类型|用途|
|-------|----|----|
|currentPassword|string|当前密码|
|newPassword|string|新密码|

|返回类型|用途|
|----|----|
|bool|是否成功修改密码|
3. 详细解释

ChangePassword用于更改钱包密码，如果成功，会自动解锁钱包数据库。
### GetNewAddress

生成新地址

1. 接口

```go
func GetNewAddress() string
```
2. 参数列表

|返回类型|用途|
|----|----|
|bool|是否生成新地址|

3. 详细解释

NewAddress在本地钱包的控制下创建一个新地址。

### GetWalletBalance

查询钱包余额

1. 接口
```go
GetWalletBalance() string
```

2. 参数列表

|返回类型|用途|
|----|----|
|string|返回余额|

3. 详细解释

返回总未支出输出(已确认和未确认)，所有已确认的未支出输出和钱包控制下的所有未确认的未支出输出。

---



## 通道 `channel`

### ConnectPeer

连接到另一个lnd节点

```go
func ConnectPeer(pubkey, host string) bool
```

|      函数参数      |  类型  |       用途       |
|:------------------:|:------:|:----------------:|
| pubkey         | string | 节点公钥         |
| host  | string   | 目的主机加端口即套接字 |


| 返回参数 |  类型  |
|:--------:|:------:|
| 是否成功   | boolean |


### OpenChannel

开通道

```go
func OpenChannel(nodePubkey string, localFundingAmount int64) string
```

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
|  是否成功  | boolean |

### OpenChannelSync

开通道

> 是 OpenChannel RPC 调用的同步版本。该调用供 REST 代理客户端使用。与所有其他同步调用一样，所有字节片都将以十六进制编码字符串的形式填充

|      函数参数      |  类型  |       用途       |
|:------------------:|:------:|:----------------:|
| nodePubkey         | string | 节点公钥         |
| localFundingAmount | long   | 本地注资数量(聪) |

| 返回参数 |  类型  |
|:--------:|:------:|
|  资金交易的输出点（txid:index）  | string |

### CloseChannel

关通道

```go
func CloseChannel(fundingTxidStr string, outputIndex int) bool
```

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

### PendingChannels

查询pending等待中的通道

```go
func PendingChannels() string
```

| 返回参数 |  类型  |
|:--------:|:------:|
| 查询到的数据   | string |


### ListChannels

列出通道

> 返回该节点参与的所有开放通道的描述。

```go
func ListChannels() string
```

| 返回参数 |  类型  |
|:--------:|:------:|
| 查询到的数据   | string |


### GetChanInfo


查询通道信息

```go
func GetChanInfo(chanId int) string
```

> 返回指定通道的最新认证网络公告，该通道由通道 ID 标识：一个 8 字节整数，用于唯一标识区块链中交易资金输出的位置

|      函数参数      |  类型  |       用途       |
|:------------------:|:------:|:----------------:|
| chanId          | int | 通道的 id         |


| 返回参数 |  类型  |
|:--------:|:------:|
| 通道详情数据   | string |

### ChannelBalance

```go
func ChannelBalance() string
```

通道余额

> 返回所有开放渠道的资金总额报告，按本地/远程、待结算的本地/远程和未结算的本地/远程余额分类

| 返回参数 |  类型  |
|:--------:|:------:|
| 所有通道余额详情数据   | string |

### ClosedChannels

已关闭通道

> 返回该节点参与的所有封闭通道的描述

```go
func ClosedChannels() string
```

| 返回参数 |  类型  |
|:--------:|:------:|
| 查询到的数据   | string |


## 发票 `invoice` 收款

### AddInvoice
```go    
> func AddInvoice(value long,memo string) string
```      
生成发票     
根据传入的金额生成闪电发票

| 参数    | 类型     | 用途      |
|-------|--------|---------|
| value | long   | 需要收款的金额 |
| memo  | string | 发票附带的信息 |

| 返回类型   | 用途    |
|--------|-------|
| string | 生成的发票 |

---

### ListInvoices       
```go
func ListInvoices() string
```
发票列表     
列出当前账户下所持有的发票     

| 返回类型   | 用途    |
|--------|-------|
| string | 返回发票列表 |


### lookupInvoice
```go
>func LookupInvoice(rhash string) string
```
查询发票信息

| 参数    | 类型     | 用途              |
|---------|--------|-----------------|
| string | string | 需要查询的闪电发票支付hash |

| 返回类型   | 用途          |
|--------|-------------|
| string | 指定闪电发票的详细信息 |


## 付款 `pay`

### DecodePayReq
>func DecodePayReq(pay_req string) int64   

解码发票  
解码支付请求字符串,返回发票金额

| 参数    | 类型     | 用途          |
|---------|--------|-------------|
| pay_req | string | 要解码的支付请求字符串 |

| 返回类型  | 用途                         |
|-------|----------------------------|
| int64 | 解码发票的金额. 0：可支付任意金额。-1：解码错误 |


### EstimateRouteFee
>EstimateRouteFee(dest string, amtsat int64) string  

计算费用  
允许测试发送到目标节点指定金额是否成功

| 参数   | 类型     | 用途       |
|--------|--------|----------|
| dest   | string | 目标节点地址   |
| amtsat | int64  | 要测试发送的金额 |

| 返回类型   | 用途     |
|--------|--------|
| string | 返回测试信息 |

### SendPaymentSync

>func SendPaymentSync(invoice string) string  

支付发票  
支付闪电发票请求  

| 参数    | 类型     | 用途         |
|---------|--------|------------|
| invoice | string | 要支付的闪电发票请求 |

| 返回类型   | 用途                 |
|--------|--------------------|
| string | 返回支付hash,可用于跟踪付款进度 |

### TrackPaymentV2
>func TrackPaymentV2(payhash string) string

交易信息  
返回由付款哈希值标识的付款的更新流。


| 参数    | 类型     | 用途            |
|---------|--------|---------------|
| payhash | string | 要查询的支付的支付hash |

| 返回类型   | 用途     |
|--------|--------|
| string | 返回支付状态 |

### SendCoins
>func SendCoins(addr string, amount int64) string

发送至链上  
向指定比特币地址发送金额。

| 参数   | 类型     | 用途   |
|--------|--------|------|
| addr   | string | 目的地址 |
| amount | string | 发送金额 |

| 返回类型   | 用途     |
|--------|--------|
| string | 返回交易ID |


