
返回json的接口：  

### 4/17   

打开的通道列表   

func ListChannels()    

挂起的通道列表  

func PendingChannels()   

获取通道状态  
func GetChannelState(chanPoint string)

true:
活动：ACTIVE  
已打开但未活动：INACTIVE   
挂起待打开：PENDING_OPEN   
挂起待关闭：PENDING_CLOSE   
已关闭：CLOSED

false:   
未找到：NO_FIND_CHANNEL


获取打开的通道信息  
func GetChannelInfo(chanPoint string)
 
true:   
返回指定的已打开通道的详细信息

false:   
未找到：NO_FIND_CHANNEL

新增:显示Lit所有服务活动信息   
SubServerStatus()   


### 4/18   v0.0.1

获取钱包余额：  
func GetWalletBalance()    

获取LND节点信息：  
func GetInfoOfLnd()  

获取地址列表：  
func ListAddress() string   

获取账户列表：  
func ListAccounts() string   

查找指定账户信息：  
func FindAccount(name string) string 

发送比特币：
func SendCoins(addr string, amount int64) string   

发送所有比特币：  
func SendAllCoins(addr string) string   





