1.批量发送比特币
SendMany(jsonAddr string, feerate int64) string

feeRate单位为sat/vbyte,使用0表示使用默认的费率。

jsonAddr格式：
```json
{"address1":100,"address2":200}
```

example:
```json
{"bcrt1pq83tk5uu0lpwk2gd7f736ttrmexed8xazfz3jmwj0ml26cwyurast4xk3w":1111,"bcrt1pra9w5dphnx75n0pjzcxlc5e8k9vg9sdupttyr36prn2t6ullr9eq0utvac":2222}
```

2.资产支付/批量支付
SendAssets(jsonAddrs string, feeRate int64) string

feeRate单位为sat/vbyte,使用0表示使用默认的费率。

jsonAddrs格式：
```json
["addrs1","addrs2"]
```
example:
```json
["bcrt1pq83tk5uu0lpwk2gd7f736ttrmexed8xazfz3jmwj0ml26cwyurast4xk3w","bcrt1pra9w5dphnx75n0pjzcxlc5e8k9vg9sdupttyr36prn2t6ullr9eq0utvac"]
```

3.生成资产接收地址
NewAddr(assetId string, amt int)
assetId:资产ID
amt:接收资产的数量

4.查询资产接收列表
AddrReceives() string

5.查询资产是否为本地发行
CheckAssetIssuanceIsLocal(assetId string)

返回值：
```json
{"success":true,"error":"","code":200,"data":{"is_local":true,"asset_id":"8513a233c4d011fccb2037aff509070e874293e8cd5f2e42310db5d947e3eb32","batch_txid":"a145851eb09040a1c9822891f9b64c35a54621583a6cb63a47664f055e71fb91","amount":22222}}
```