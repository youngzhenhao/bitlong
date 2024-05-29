DecodeAddr:
```json
{"success":true,"error":"","data":{"encoded":"taprt1qqqsqqspqqzzqgrg25gkuuprc9up7yr0673k9zaxupg72tzyudqufe4krxczrl0kqcssyeuwjychfyr5l39eudwwczrrf2mjshzx7tlwmvx8vac0hgn8ypqypqssx7gh5kkdl7vp84zyj0ql4p7qc7a8xuyu5wnwckh705aw6ks69c38pgqaurp0dpshx6rdv95kcw309akkz6tvvfhhstn5v4ex66twv9kzumrfva58gmnfdenjuar0v3shjw35xsesa0054s","asset_id":"206855116e7023c1781f106fd7a3628ba6e051e52c44e341c4e6b619b021fdf6","asset_type":0,"amount":222,"group_key":"","script_key":"02678e9131749074fc4b9e35cec08634ab7285c46f2feedb0c76770fba26720404","internal_key":"037917a5acdff9813d44493c1fa87c0c7ba73709ca3a6ec5afe7d3aed5a1a2e227","tapscript_sibling":"","taproot_output_key":"b41ca9423a4693cb561de024dc8f5f1e8074a2dc3f86ee452debe44514dbcd2c","proof_courier_addr":"hashmail://mailbox.terminal.lightning.today:443","asset_version":0}}
```

QueryAddrs:
```json
{
    "success": true,
    "error": "",
    "data": [
        {
            "encoded": "taprt1qqqsqqspqqzzq6sgulg2xl4wjtve8mmp66mc3w4yqsx4latm63av93qkttst96hmqcssyyvjkavwu6m5zmusj7sy2qfp0p8wxjnavdwnffrssh78zpftl5n3pqss8h5ddpz2yylzm7us2gvpqsvlxstztklavz0rs5z9kpt3ksmwk0swpgq3grp0dpshx6rdv95kcw309akkz6tvvfhhstn5v4ex66twv9kzumrfva58gmnfdenjuar0v3shjw35xses8n729c",
            "asset_id": "6a08e7d0a37eae92d993ef61d6b788baa4040d5ff57bd47ac2c4165ae0b2eafb",
            "asset_type": 0,
            "amount": 20,
            "group_key": "",
            "script_key": "021192b758ee6b7416f9097a0450121784ee34a7d635d34a47085fc71052bfd271",
            "internal_key": "03de8d6844a213e2dfb90521810419f341625dbfd609e385045b0571b436eb3e0e",
            "tapscript_sibling": "",
            "taproot_output_key": "bde2103b31c48ca94437d10ac490ad5c0db6f3f5a9f33c0d062f3b35d6fb7e7b",
            "proof_courier_addr": "hashmail://mailbox.terminal.lightning.today:443",
            "asset_version": 0
        }
    ]
}
```


QueryAssetTransfers
```json
{
    "success": true,
    "error": "",
    "data": [
        {
            "transfer_timestamp": 1716964271,
            "anchor_tx_hash": "004529349dbc7e536e9d21e09c3828857ee49cf6aea14b4c0f5cc563c26436a1",
            "anchor_tx_height_hint": 5162,
            "anchor_tx_chain_fees": 12725,
            "inputs": [
                {
                    "anchor_point": "35e5a9f9ddc20134a11bb5ae1ed2a34ba8439e06c18470ae17326f5c2d8a117a:1",
                    "asset_id": "206855116e7023c1781f106fd7a3628ba6e051e52c44e341c4e6b619b021fdf6",
                    "script_key": "02aaf12d1254d81742bcef3aa05ce27f2db9b5642b9d6ca92e86ab04dd5ec88c1b",
                    "amount": 222
                }
            ],
            "outputs": [
                {
                    "anchor": {
                        "outpoint": "a13664c263c55c0f4c4ba1aef69ce47e8528389ce0219d6e537ebc9d34294500:0",
                        "value": 1000,
                        "internal_key": "02bc4eaa5471db67d286e3f1c4af7cca967b5f3e439f219a2fb53cd52f4fa9bc42",
                        "taproot_asset_root": "0095721a214f02624b936f6ddebd8edce412e752157287f8e3efc663d31cfc4e",
                        "merkle_root": "0095721a214f02624b936f6ddebd8edce412e752157287f8e3efc663d31cfc4e",
                        "tapscript_sibling": "",
                        "num_passive_assets": 0
                    },
                    "script_key": "02df80ce1e1241430c82f67403094243e6718479c100438cc32d308fa5eb783251",
                    "script_key_is_local": true,
                    "amount": 123,
                    "new_proof_blob": "",
                    "split_commit_root_hash": "b3010b2c905c976946ca666c2f8a69ff5d6dde34919dec5396b4be96cf084a43",
                    "output_type": "OUTPUT_TYPE_SPLIT_ROOT",
                    "asset_version": "ASSET_VERSION_V0"
                },
                {
                    "anchor": {
                        "outpoint": "a13664c263c55c0f4c4ba1aef69ce47e8528389ce0219d6e537ebc9d34294500:1",
                        "value": 1000,
                        "internal_key": "03839e790a08af2fc3c16375514ed1f832cb4bb2be2a777b8c7d461f14d8af1e59",
                        "taproot_asset_root": "e2ac8146ec9a481c3d2ff9dbd4e462aeefbbf12418e6213a61d022776bea6150",
                        "merkle_root": "e2ac8146ec9a481c3d2ff9dbd4e462aeefbbf12418e6213a61d022776bea6150",
                        "tapscript_sibling": "",
                        "num_passive_assets": 0
                    },
                    "script_key": "021621cc549c356e74f06746b21ec6f2338ea104608c20c09b94678c1dccbaa2f0",
                    "script_key_is_local": false,
                    "amount": 99,
                    "new_proof_blob": "",
                    "split_commit_root_hash": "",
                    "output_type": "OUTPUT_TYPE_SIMPLE",
                    "asset_version": "ASSET_VERSION_V0"
                }
            ]
        }
    ]
}
```

AddrReceives
```json
{
    "success": true,
    "error": "",
    "data": [
        {
            "creation_time_unix_seconds": 1716791915,
            "addr": {
                "encoded": "taprt1qqqsqqspqqzzqgrg25gkuuprc9up7yr0673k9zaxupg72tzyudqufe4krxczrl0kqcssy0xvas777q4ya0pgn524r2sy7vjhn32ezwhfecux3vtxvtnc0mrlpqssxg6fdcmdtv02wqnu54x054l7luqa5axzcf405q8zg6lvfmgum4p2pgqusrp0dpshx6rdv95kcw309akkz6tvvfhhstn5v4ex66twv9kzumrfva58gmnfdenjuar0v3shjw35xsesq0k6a9",
                "asset_id": "206855116e7023c1781f106fd7a3628ba6e051e52c44e341c4e6b619b021fdf6",
                "asset_type": 0,
                "amount": 200,
                "group_key": "",
                "script_key": "023cccec3def02a4ebc289d1551aa04f32579c55913ae9ce3868b16662e787ec7f",
                "internal_key": "0323496e36d5b1ea7027ca54cfa57feff01da74c2c26afa00e246bec4ed1cdd42a",
                "tapscript_sibling": "",
                "taproot_output_key": "59e07fa8b42a2f693340c73bb4e5c0a120ca4e9390d725864eafbca81ee2e84b",
                "proof_courier_addr": "hashmail://mailbox.terminal.lightning.today:443",
                "asset_version": 0
            },
            "status": "ADDR_EVENT_STATUS_COMPLETED",
            "outpoint": "a2a94ee99d7dc8d684007b1dc6f81e80330904198749337b07f099da9cfcef4e:1",
            "utxo_amt_sat": 1000,
            "taproot_sibling": "",
            "confirmation_height": 4864,
            "has_proof": true
        },
        {
            "creation_time_unix_seconds": 1716792953,
            "addr": {
                "encoded": "taprt1qqqsqqspqqzzqgrg25gkuuprc9up7yr0673k9zaxupg72tzyudqufe4krxczrl0kqcss92h395f9fkqhg27w7w4qtn387tdek4jzh8tv4yhgd2cym40v3rqmpqssy970yymj0e9r3x9tx7pk0m2t5t8p3qxluukuyjlz6nxhr538yjwspgqaurp0dpshx6rdv95kcw309akkz6tvvfhhstn5v4ex66twv9kzumrfva58gmnfdenjuar0v3shjw35xsesnq366m",
                "asset_id": "206855116e7023c1781f106fd7a3628ba6e051e52c44e341c4e6b619b021fdf6",
                "asset_type": 0,
                "amount": 222,
                "group_key": "",
                "script_key": "02aaf12d1254d81742bcef3aa05ce27f2db9b5642b9d6ca92e86ab04dd5ec88c1b",
                "internal_key": "0217cf213727e4a3898ab378367ed4ba2ce1880dfe72dc24be2d4cd71d227249d0",
                "tapscript_sibling": "",
                "taproot_output_key": "017c464429ad7ba731ea05c85a9cacad406f56ea92e1835e2fe2e1cdf5c0d2c5",
                "proof_courier_addr": "hashmail://mailbox.terminal.lightning.today:443",
                "asset_version": 0
            },
            "status": "ADDR_EVENT_STATUS_COMPLETED",
            "outpoint": "168fc91148942ad869b420837628dca276408f1433759dfd7a4a728803f78bf1:1",
            "utxo_amt_sat": 1000,
            "taproot_sibling": "",
            "confirmation_height": 4866,
            "has_proof": true
        },
        {
            "creation_time_unix_seconds": 1716792953,
            "addr": {
                "encoded": "taprt1qqqsqqspqqzzqgrg25gkuuprc9up7yr0673k9zaxupg72tzyudqufe4krxczrl0kqcssyeuwjychfyr5l39eudwwczrrf2mjshzx7tlwmvx8vac0hgn8ypqypqssx7gh5kkdl7vp84zyj0ql4p7qc7a8xuyu5wnwckh705aw6ks69c38pgqaurp0dpshx6rdv95kcw309akkz6tvvfhhstn5v4ex66twv9kzumrfva58gmnfdenjuar0v3shjw35xsesa0054s",
                "asset_id": "206855116e7023c1781f106fd7a3628ba6e051e52c44e341c4e6b619b021fdf6",
                "asset_type": 0,
                "amount": 222,
                "group_key": "",
                "script_key": "02678e9131749074fc4b9e35cec08634ab7285c46f2feedb0c76770fba26720404",
                "internal_key": "037917a5acdff9813d44493c1fa87c0c7ba73709ca3a6ec5afe7d3aed5a1a2e227",
                "tapscript_sibling": "",
                "taproot_output_key": "b41ca9423a4693cb561de024dc8f5f1e8074a2dc3f86ee452debe44514dbcd2c",
                "proof_courier_addr": "hashmail://mailbox.terminal.lightning.today:443",
                "asset_version": 0
            },
            "status": "ADDR_EVENT_STATUS_COMPLETED",
            "outpoint": "168fc91148942ad869b420837628dca276408f1433759dfd7a4a728803f78bf1:2",
            "utxo_amt_sat": 1000,
            "taproot_sibling": "",
            "confirmation_height": 4866,
            "has_proof": true
        },
        {
            "creation_time_unix_seconds": 1716796601,
            "addr": {
                "encoded": "taprt1qqqsqqspqqzzqgrg25gkuuprc9up7yr0673k9zaxupg72tzyudqufe4krxczrl0kqcss92h395f9fkqhg27w7w4qtn387tdek4jzh8tv4yhgd2cym40v3rqmpqssy970yymj0e9r3x9tx7pk0m2t5t8p3qxluukuyjlz6nxhr538yjwspgqaurp0dpshx6rdv95kcw309akkz6tvvfhhstn5v4ex66twv9kzumrfva58gmnfdenjuar0v3shjw35xsesnq366m",
                "asset_id": "206855116e7023c1781f106fd7a3628ba6e051e52c44e341c4e6b619b021fdf6",
                "asset_type": 0,
                "amount": 222,
                "group_key": "",
                "script_key": "02aaf12d1254d81742bcef3aa05ce27f2db9b5642b9d6ca92e86ab04dd5ec88c1b",
                "internal_key": "0217cf213727e4a3898ab378367ed4ba2ce1880dfe72dc24be2d4cd71d227249d0",
                "tapscript_sibling": "",
                "taproot_output_key": "017c464429ad7ba731ea05c85a9cacad406f56ea92e1835e2fe2e1cdf5c0d2c5",
                "proof_courier_addr": "hashmail://mailbox.terminal.lightning.today:443",
                "asset_version": 0
            },
            "status": "ADDR_EVENT_STATUS_TRANSACTION_CONFIRMED",
            "outpoint": "a187ffa4fc76374b739f4d1c22f56049fd6ddd02e4d62c59c6bb5963a3ad2e48:1",
            "utxo_amt_sat": 1000,
            "taproot_sibling": "",
            "confirmation_height": 4873,
            "has_proof": true
        },
        {
            "creation_time_unix_seconds": 1716796601,
            "addr": {
                "encoded": "taprt1qqqsqqspqqzzqgrg25gkuuprc9up7yr0673k9zaxupg72tzyudqufe4krxczrl0kqcssyeuwjychfyr5l39eudwwczrrf2mjshzx7tlwmvx8vac0hgn8ypqypqssx7gh5kkdl7vp84zyj0ql4p7qc7a8xuyu5wnwckh705aw6ks69c38pgqaurp0dpshx6rdv95kcw309akkz6tvvfhhstn5v4ex66twv9kzumrfva58gmnfdenjuar0v3shjw35xsesa0054s",
                "asset_id": "206855116e7023c1781f106fd7a3628ba6e051e52c44e341c4e6b619b021fdf6",
                "asset_type": 0,
                "amount": 222,
                "group_key": "",
                "script_key": "02678e9131749074fc4b9e35cec08634ab7285c46f2feedb0c76770fba26720404",
                "internal_key": "037917a5acdff9813d44493c1fa87c0c7ba73709ca3a6ec5afe7d3aed5a1a2e227",
                "tapscript_sibling": "",
                "taproot_output_key": "b41ca9423a4693cb561de024dc8f5f1e8074a2dc3f86ee452debe44514dbcd2c",
                "proof_courier_addr": "hashmail://mailbox.terminal.lightning.today:443",
                "asset_version": 0
            },
            "status": "ADDR_EVENT_STATUS_TRANSACTION_CONFIRMED",
            "outpoint": "a187ffa4fc76374b739f4d1c22f56049fd6ddd02e4d62c59c6bb5963a3ad2e48:2",
            "utxo_amt_sat": 1000,
            "taproot_sibling": "",
            "confirmation_height": 4873,
            "has_proof": true
        },
        {
            "creation_time_unix_seconds": 1716960245,
            "addr": {
                "encoded": "taprt1qqqsqqspqqzzqgrg25gkuuprc9up7yr0673k9zaxupg72tzyudqufe4krxczrl0kqcss92h395f9fkqhg27w7w4qtn387tdek4jzh8tv4yhgd2cym40v3rqmpqssy970yymj0e9r3x9tx7pk0m2t5t8p3qxluukuyjlz6nxhr538yjwspgqaurp0dpshx6rdv95kcw309akkz6tvvfhhstn5v4ex66twv9kzumrfva58gmnfdenjuar0v3shjw35xsesnq366m",
                "asset_id": "206855116e7023c1781f106fd7a3628ba6e051e52c44e341c4e6b619b021fdf6",
                "asset_type": 0,
                "amount": 222,
                "group_key": "",
                "script_key": "02aaf12d1254d81742bcef3aa05ce27f2db9b5642b9d6ca92e86ab04dd5ec88c1b",
                "internal_key": "0217cf213727e4a3898ab378367ed4ba2ce1880dfe72dc24be2d4cd71d227249d0",
                "tapscript_sibling": "",
                "taproot_output_key": "017c464429ad7ba731ea05c85a9cacad406f56ea92e1835e2fe2e1cdf5c0d2c5",
                "proof_courier_addr": "hashmail://mailbox.terminal.lightning.today:443",
                "asset_version": 0
            },
            "status": "ADDR_EVENT_STATUS_COMPLETED",
            "outpoint": "35e5a9f9ddc20134a11bb5ae1ed2a34ba8439e06c18470ae17326f5c2d8a117a:1",
            "utxo_amt_sat": 1000,
            "taproot_sibling": "",
            "confirmation_height": 5012,
            "has_proof": true
        }
    ]
}
```