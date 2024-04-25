### 4/24

1.列出当前钱包中的所有资产   
-Listassets() (同id资产不合并) or  Listbalances()（合并）
```go
func Listbalances()
```

>{"success":true,"error":"","data":{"asset_balances":{"bc5c740a97f06fef3caddb4991b03a93d14c977cd9c6468cf68526a8da6bc8c2":{"asset_genesis":{"genesis_point":"ade04fbfb66a29aa7939e21baf6cb32a7df26baf92bd1599386cb28853f95d63:1","name":"cat","meta_hash":"bR1LNuWyBHqV7btqKwlL1GsLjW8+Rk8IxqVLxlq6r44=","asset_id":"vFx0Cpfwb+88rdtJkbA6k9FMl3zZxkaM9oUmqNpryMI="},"balance":1000},"e62cc7e6eedfb4babd0c434bd45cb87101e7b67fa5506a8da474b2b6cd709d76":{"asset_genesis":{"genesis_point":"d2f7c0ba01a8b35b5f30cc746452b020f59c16e5635929a7440ef37b1628ab0e:1","name":"lp222","meta_hash":"iN/NS/BbYpF8DZHfsCNupzD9sK9VE9Vr2OTxUeOpwxw=","asset_id":"5izH5u7ftLq9DENL1Fy4cQHntn+lUGqNpHSyts1wnXY="},"balance":1234},"f2eebbfcaec64a69496877f7e5d9262eddca68cf4001fbf3c06f63c213345f25":{"asset_genesis":{"genesis_point":"ba00154187213a88c4e97024cae697550a3e53dff28faed110ea63e4e4af3fdf:1","name":"lp111","meta_hash":"iN/NS/BbYpF8DZHfsCNupzD9sK9VE9Vr2OTxUeOpwxw=","asset_id":"8u67/K7GSmlJaHf35dkmLt3KaM9AAfvzwG9jwhM0XyU="},"balance":1234}}}}


2.从宇宙中查询资产的发行信息
```go
func GetAssetInfo(asset_id string)
```
> {"success":true,"error":"","data":{"asset":{"asset_genesis":{"genesis_point":"d2f7c0ba01a8b35b5f30cc746452b020f59c16e5635929a7440ef37b1628ab0e:1","name":"lp222","meta_hash":"iN/NS/BbYpF8DZHfsCNupzD9sK9VE9Vr2OTxUeOpwxw=","asset_id":"5izH5u7ftLq9DENL1Fy4cQHntn+lUGqNpHSyts1wnXY="},"amount":1234,"script_key":"AkBant7idrJmkZu78r83yMqVj0ONtIdL8BRRgKMJrEUq","chain_anchor":{"anchor_tx":"AgAAAAABAQ6rKBZ78w5EpylZY+UWnPUgsFJkdMwwX1uzqAG6wPfSAQAAAAD/////AugDAAAAAAAAIlEg24bgfoRLKCOpLWCpZj8KvMw0KgJC0KILDSmD1qOyZO1XGKw6AAAAACJRIMARa8VODyBztR7lQoD665av9Vm5xYmr2yKcr7rdSyIaAUC0EiwutLaOh10cPYdHRtesU9ZFaSYVOsjx4q5cen9fRKTcgDALONH6ldKMLX23Q5h49Go6s4+WFLA/3HkUw7IaAAAAAA==","anchor_block_hash":"4758b4e5d15a5ba86ca69c3bd52abf6bcb85ed5b3497ee2ca320394efd419caf","anchor_outpoint":"4e375fa7c727700ddb5634d2b8a30b7d4374e12c305b1bd53a519fe4d3fecc94:0","internal_key":"AwynRKC7UtXshvR3R8I+NKjbtwEwk1+iHSNGMveO+ZkU","merkle_root":"ckYLJDv2gYHYGYs6TWPOWm1A+AGhhIJM5gfqKtulY6o=","block_height":3115}},"meta":"313233313233313233","createTime":1713769570}}

3.在2未成功的情况下，可以从指定宇宙中同步资产的发行信息   
```go
func SyncUniverse(universeHost  string, asset_id    string)
```
> {"success":true,"error":"","data":{"synced_universes":[{"old_asset_root":{},"new_asset_root":{"id":{"Id":{"AssetId":"lDXqdJZFLXh/5ECh1Tfvnf0/FQW1Xy6kaqE4/TY1BcY="},"proof_type":1},"mssmt_root":{"root_hash":"Tl+nRpv+/7V6RwG1SPRTgnazeOPdiz+CC35OrC/F1kY=","root_sum":100}},"new_asset_leaves":[{"asset":{"asset_genesis":{"genesis_point":"3a6c8b6df84e6e242be2fbd4c8bba167fa7fbcf7fc6974790624ca8a32d717a5:1","name":"ghj","meta_hash":"WQWoIog+ROlsWmN+kHvLdGxeIF86HBYyGMIQgQyYYxA=","asset_id":"lDXqdJZFLXh/5ECh1Tfvnf0/FQW1Xy6kaqE4/TY1BcY="},"amount":100,"script_key":"As3HIClk95jnyqVWNwrnIsDKsrMvRGixwMW9CwXy4xdW","prev_witnesses":[{"prev_id":{"anchor_point":"0000000000000000000000000000000000000000000000000000000000000000:0","asset_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=","script_key":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"}}]},"proof":"VEFQUAAEAAAAAAIkpRfXMorKJAZ5dGn897x/+mehu8jU++IrJG5O+G2LbDoAAAABBFAAAAAg+D/dT83rvHg4IIVkx3A91q1UsqFOQdntEwAAAAAAAACClafkrPnzKYWfSGi9IVEii9Kh3jCqjbEDnj/xmIwLy7YIImbVFy4ZSpBcCQbNAgAAAAABAaUX1zKKyiQGeXRp/Pe8f/pnobvI1PviKyRuTvhti2w6AQAAAAD/////AugDAAAAAAAAIlEgpflHp9fsqYmGoHXNvjXtBSwYq3LLFhcpItvrn7G7HD//BQAAAAAAACJRIA+idauUVRRMORb9KKV/CmmNsLuSa7a6zoABw+q2Q84gAUAoW5bKwOKgjW7+3g6QdXnKj0HIdOFU7dA/muk5/zfb3vLCdVngHtTjMRrmcCDwHNGvqcC2th/J3u5BxH6yjMzcAAAAAAj9AWMLdfXkh+lxWLnyG2JRUK1i43C7fQdFZM7D34uCWPn1M2ffLW3L97iZZF/Ep343tPad1VxOx0jOi8y5rfbXHZ93Y9oX82GJrhYcLP8nzc4G2pbb9sOXuEqkHr2IhxHYJMEONhk9uY4XqLLC2p9O8aNpSuFZ6OU/BAgvNXgmdbW2BJGZ5OprsA3Da7VMCGnqxBYcW36uOMcRlgJqRYd7ZcDyxpZQ4oMmATvYOgg9Jg+SzZPXkex9MHkM2cc1t9hKaUONhOg305pFbujwZhXjPL4kz9YNHhLBt/E8RDH0V3LffEYJjKV5VXC1s5FN/1cdkBRd+ICkDQ/rPCrqlZoTB7kbfnLpT7cpSpi1tubEprUsnAa+q9CtS2E6OGo7OP4Mx19pd+4QjqozrfxS30By3M8+NdTSVx3c8LL8eGpRYgV9Ki1PEPHLWA/b8RaUbafQ5uwGMMDlPwd1S8xobm9YexIr/ZkFCuoAAQACTaUX1zKKyiQGeXRp/Pe8f/pnobvI1PviKyRuTvhti2w6AAAAAQNnaGpZBagiiD5E6WxaY36Qe8t0bF4gXzocFjIYwhCBDJhjEAAAAAAABAEABgFkC2kBZwFlAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAOAgAAECECzccgKWT3mOfKpVY3CuciwMqysy9EaLHAxb0LBfLjF1YMnwAEAAAAAAIhA3f3VR/Dt9ysSiQ0y4rGpH29OLQN5n9L1ig7hKexUAKQA3QBSQABAAIglDXqdJZFLXh/5ECh1Tfvnf0/FQW1Xy6kaqE4/TY1BcYEIgAA//////////////////////////////////////////8CJwABAAIiAAD//////////////////////////////////////////w0wAS4ABAAAAAECIQIxGffjHSG/Y+WX75XuXjpOTWbnIWYZV1zaYjsEhxPJvQUDBAEBEQ4AAQACCWNoampra2pqahYEACd8LBdNpRfXMorKJAZ5dGn897x/+mehu8jU++IrJG5O+G2LbDoAAAABA2doalkFqCKIPkTpbFpjfpB7y3RsXiBfOhwWMhjCEIEMmGMQAAAAAAA="}]}]}}
- universeHost: 宇宙地址, 如"http://192.168.1.100:8080",当输入空字符串时，默认使用测试网公共宇宙
- 返回NOT_NEW_DATA错误码时，表示该资产已经同步过了，无需重复同步。

4.查询资产元数据
```go
func api.FetchAssetMeta(isHash bool, data string)
```
>{"success":true,"error":"","data":"63686a6a6b6b6a6a6a"}

-isHash 为true时，data为资产元数据的hash；为false时，data为资产id。

5.json信息解释
```        
   "asset_genesis":{
                "genesis_point": "d2f7c0ba01a8b35b5f30cc746452b020f59c16e5635929a7440ef37b1628ab0e:1", // 资产创世点
                "name": "lp222",        // 资产名称
                "meta_hash": "iN/NS/BbYpF8DZHfsCNupzD9sK9VE9Vr2OTxUeOpwxw=",       // 资产元数据的hash (base64编码)
                "asset_id": "5izH5u7ftLq9DENL1Fy4cQHntn+lUGqNpHSyts1wnXY="     // 资产id(base64编码)
            },
            "balance": 1000, // 资产余额
            "amount": 1234, // 发行数量
            "meta": "313233313233313233" // 资产元数据
            "createTime": 1713769570 // 资产创建时间(Unix时间戳)
```


工具:base64解码
Base64Decode(s string)


# 带图像的元数据
1.格式
暂时采用 Tiramisu 钱包使用的 taproot assets 协议的元数据格式。
参考：https://github.com/snow884/taproot-assets-metadata-format   
```
{
    "acronym": string,(缩写名)
    "description": string, (描述)
    "image_data": [<mediatype>][;base64],<data> (dataurl格式的字符串)
}

```
example:  

```
{
    "acronym": "MTC",
    "description": "This is my test coin", 
    "image_data": "data:image/jpg;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAIAAAD8GO2jAAABc0lEQVR4nGK5skGaAQaUnTbA2ee2JsHZ00zuwtlzLMPhbO3f/Qi9zR/g7Gch7HA2EwONwdC3gOU5UzKcY/7TCs7+eagOzj5jzg9nh3HegLPv7JWFs63vF8HZ/E2IuBn6QUT7OFjXuh7O+TAbkZbPvxeAszO+NcHZSTwX4exCgxlw9s1QCThb0P0fnD30g4jmFjAeU9oI53yMeAFnux5B2N3aJgNnb/lqAmfHB5+Esy83voezWyKM4OyhH0S0zwdnlk6Fc67ORqR3v42OcLbWUic4O79/Jpy958IEODt79xU4m4+tE84e+kFE+3zgLrodzmlb+w7O9gqphbOPCCM0GG/dCmffi1oGZz+WegBn//2PiJuhH0S0j4MVlWlwzm0FGzh7veotONsquw/OnsSHyB/WPxFxMNvOAs62QGSVYRBEtI+D8GZEPvB2QdSr0qdz4Ow8cS84u/+RKZxtLoqoJ1q3IurwbHeE3qEfRDS3ABAAAP//hnBeQ7p8R+8AAAAASUVORK5CYII="
}
```

2.创建一个标准格式元数据

```go
func CreateNewMeta(acronym string, description string, imagefile string) string
```
acronym: 缩写名(可选)
description: 资产描述(必填)
imagefile: 图片文件路径(可选)

返回值: 元数据字符串

可能返回的错误码:
OPEN_IMAGE_FILE_ERROR 打开图片文件失败

3.创建资产

```go
func MintAsset(name string, assetTypeIsCollectible bool, assetMetaData string, amount int, newGroupedAsset bool) string
```
将2中创建的元数据字符串作为参数传入assetMetaData，创建资产。

name: 资产名称(必填)
assetTypeIsCollectible: 是否为收藏品
assetMetaData:元数据类型
amount:数量
newGroupedAsset:是否可增发

4.增发资产
```go
func AddGroupAsset(name string, assetTypeIsCollectible bool, assetMetaData string, amount int,groupKey string) string
```
当2中的资产创建为可增发资产时，调用该方法将资产增发。

groupKey 用于锚定增发资产组的key


5.解析资产元数据:判断资产元数据是否符合格式要求
```go
func CheckMetaStandard(Meta string) bool
```
Meta 资产元数据原始字符串
返回值 bool 是否符合标准
当符合标准使，即可用资产元数据标准解析
不符合标准时，可将整个元数据字符串作为一个无图像的资产的description

6.解析资产元数据:解析资产元数据中的图像
```go
func DecodeBase64ForImage(image string, dir string, name string) string 
```
资产元数据为标准的json格式，可通过json解析获得acronym、description、image_data三个字段

image_data为base64编码的图像数据，需要解码并保存到指定目录下，返回保存路径
image:json解析后的image_data字段
dir:保存图像的目录
name:保存图像文件名

返回值 string 保存图像的路径

可能返回的错误码：
OPEN_IMAGE_FILE_ERROR 打开图像文件失败
WRITE_IMAGE_FILE_ERROR 写入图像文件失败
DATA_NOT_DATAURL 数据不是DATAURL格式编码
DATA_NOT_IMAGE 数据不是图片格式