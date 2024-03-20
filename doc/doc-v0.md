## 业务流程

### TODO

- 为每个函数添加接口及详细描述

### 钱包解锁 `unlock`

#### GenSeed

助记词生成

1.  接口
```go
func GenSeed() string
```

2.  参数列表

|返回类型|用途|
|----|----|
|string|助记词|

3.  详细解释

是用于实例化新 lnd 实例的第一个方法。该方法允许调用者根据可选的口令生成新的加密种子。如果提供了口令，则需要口令来解密密码种子，以显示内部钱包种子。

用户获得并验证密码种子后，应使用 InitWallet 方法提交新生成的种子，并创建钱包。

#### InitWallet

初始化钱包

1. 接口

```go
func InitWallet(seed string, password string) bool
```

2. 参数列表

| 输入参数 | 类型 | 用途|
|--------|----|----|
| seed | string | 助记词 |

| 返回类型 | 用途 |
|--------|---------|
|bool|是否初始化成功|

3. 详细解释

首次启动时使用，用于完全初始化守护进程及其内部钱包。至少必须提供一个钱包密码。这将用于加密磁盘上的敏感资料。

在恢复情况下，用户还可以指定自己的密码和口令。如果设置了该密码，守护进程就会使用之前的状态来初始化其内部钱包。

或者，也可以使用 GenSeed RPC 来获取种子，然后将其提交给用户。经用户验证后，可将种子输入此 RPC，以提交新钱包。

#### UnlockWallet

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


#### ChangePassword

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
#### GetNewAddress

生成新地址

1. 接口

```go
func GetNewAddress() bool
```
2. 参数列表

|返回类型|用途|
|----|----|
|bool|是否生成新地址|

3. 详细解释

NewAddress在本地钱包的控制下创建一个新地址。

#### GetWalletBalance

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

| 参数    | 类型  | 用途      |
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

| 参数    | 类型     | 用途              |
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

| 参数    | 类型     | 用途          |
|---------|--------|-------------|
| pay_req | string | 要解码的支付请求字符串 |

| 返回类型  | 用途                         |
|-------|----------------------------|
| int64 | 解码发票的金额. 0：可支付任意金额。-1：解码错误 |


#### EstimateRouteFee
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

#### SendPaymentSync

>func SendPaymentSync(invoice string) string  

支付发票  
支付闪电发票请求  

| 参数    | 类型     | 用途         |
|---------|--------|------------|
| invoice | string | 要支付的闪电发票请求 |

| 返回类型   | 用途                 |
|--------|--------------------|
| string | 返回支付hash,可用于跟踪付款进度 |

#### TrackPaymentV2
>func TrackPaymentV2(payhash string) string

交易信息  
返回由付款哈希值标识的付款的更新流。


| 参数    | 类型     | 用途            |
|---------|--------|---------------|
| payhash | string | 要查询的支付的支付hash |

| 返回类型   | 用途     |
|--------|--------|
| string | 返回支付状态 |

#### SendCoins
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


