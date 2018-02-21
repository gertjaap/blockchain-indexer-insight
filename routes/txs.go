package routes

import (
	"encoding/json"
    "net/http"
)

type TxsResponseTxVinScriptSig struct {
	Hex 				string					`json:"hex"`
	Asm  				string 					`json:"asm"`
}

type TxsResponseTxVin struct {	
	CoinBase 			string 						`json:"coinbase"`
	Sequence  			int64						`json:"sequence"`
	Number				int 						`json:"n"`
	TxId 				string 						`json:"txid"`
	Vout 				int 						`json:"vout"`
	ScriptSig 			TxsResponseTxVinScriptSig	`json:"scriptSig"`
	Address 			string 						`json:"addr"`
	ValueSatoshi		int64						`json:"valueSat"`
	Value 				float64 					`json:"value"`
	DoubleSpentTxID 	string 						`json:"doubleSpentTxID"`
}

type TxsResponseTxVoutScriptPubKey struct {
	Hex 				string					`json:"hex"`
	Asm  				string 					`json:"asm"`
	Addresses			[]string 				`json:"addresses"`
	Type 				string 					`json:"type"`
}

type TxsResponseTxVout struct {
	Value 				float64							`json:"value"`
	Number 				int 							`json:"n"`
	ScriptPubKey 		TxsResponseTxVoutScriptPubKey	`json:"scriptPubKey"`
	SpentTxId 			string							`json:"spentTxId"`
	SpentIndex 			int 							`json:"spentIndex"`
	SpentHeight 		int  							`json:"spentHeight"`
}

type TxsResponseTx struct {
	TxId 				string					`json:"txid"`
	Version				int 					`json:"version"`
	LockTime 			int 					`json:"locktime"`
	Vin 				[]TxsResponseTxVin 		`json:"vin"`
	Vout 				[]TxsResponseTxVout 	`json:"vout"`
	BlockHash 			string 					`json:"blockhash"`
	BlockHeight 		int 					`json:"blockheight"`
	Confirmations 		int 					`json:"confirmations"`
	Time 				int64 					`json:"time"`
	BlockTime 			int64 					`json:"blocktime"`
	IsCoinBase 			bool 					`json:"isCoinBase"`
	ValueOut 			float64 				`json:"valueOut"`
	Size 				int 					`json:"size"`
}

type TxsResponse struct {
	PagesTotal			int 					`json:"pagesTotal"`
	Txs					[]TxsResponseTx			`json:"txs"`
}

func Txs(w http.ResponseWriter, req *http.Request) {
	jsonString := `{"pagesTotal":58,"txs":[{"txid":"bc4d6ea94c702fd48d3ef543d7e1349402e3a3fe2186b9a9ea46894091a8e7cb","version":1,"locktime":0,"vin":[{"coinbase":"0354c907272f706f6f6c2e626974636f696e2e636f6d2f4542312f41443939392f464732403439343738342f104fc2cb00352f9104f4b0dc5230322600","sequence":4294967295,"n":0}],"vout":[{"value":"12.80352408","n":0,"scriptPubKey":{"hex":"76a914a7ea99e1245c4d54d6efce374434b2e091fc481788ac","asm":"OP_DUP OP_HASH160 a7ea99e1245c4d54d6efce374434b2e091fc4817 OP_EQUALVERIFY OP_CHECKSIG","addresses":["1GJrtQWQNxNxoSLqMG2LXfYU7EH2gvQJek"],"type":"pubkeyhash"},"spentTxId":null,"spentIndex":null,"spentHeight":null},{"value":"0.00000000","n":1,"scriptPubKey":{"hex":"6a24aa21a9ed60089d23761323a01b264b94bf63bcbfcf4cce2ffe82dd64e070e6c9bc46c7d5","asm":"OP_RETURN aa21a9ed60089d23761323a01b264b94bf63bcbfcf4cce2ffe82dd64e070e6c9bc46c7d5"},"spentTxId":null,"spentIndex":null,"spentHeight":null}],"blockhash":"00000000000000000044a01dc989bf8c097c3270a42eaf77200649ed9830ed8b","blockheight":510292,"confirmations":1,"time":1519243687,"blocktime":1519243687,"isCoinBase":true,"valueOut":12.80352408,"size":202},{"txid":"e7f9f457403857c2d39a383b22b3914c78cb2ba7b60f792c659efde5b9d9d3b1","version":1,"locktime":0,"vin":[{"txid":"fe4a9652f28af3552f5c9cbb775c50aa1334af241c7f38c1f59e9406ee4045b7","vout":1,"sequence":4294967295,"n":0,"scriptSig":{"hex":"483045022100901bb0cc208aa7f9debf394918bad67b0dd502b326f7d6f3f62f458099012c0e02206aa3bbdabee941de91bd8d116cdfde72ec2812fdb40c82fa7758324585c49f3e012102f16be836851b266ec1fc0f8c44ca8ba64e6706ec5eff5c8f33f03457d796275d","asm":"3045022100901bb0cc208aa7f9debf394918bad67b0dd502b326f7d6f3f62f458099012c0e02206aa3bbdabee941de91bd8d116cdfde72ec2812fdb40c82fa7758324585c49f3e[ALL] 02f16be836851b266ec1fc0f8c44ca8ba64e6706ec5eff5c8f33f03457d796275d"},"addr":"1JLKV32uQ5NV8EizbFVdd1im7TF31JZPFq","valueSat":79268839,"value":0.79268839,"doubleSpentTxID":null}],"vout":[{"value":"0.70236739","n":0,"scriptPubKey":{"hex":"76a914bf1a1f2fb5c14b990d20d4263a40326d894dfce888ac","asm":"OP_DUP OP_HASH160 bf1a1f2fb5c14b990d20d4263a40326d894dfce8 OP_EQUALVERIFY OP_CHECKSIG","addresses":["1JRTMchtgZDdyMwwMGdUtnvbw6CS4zFqo2"],"type":"pubkeyhash"},"spentTxId":"fa0ae2b5fb33eac005346b67803080b315277e37404431f0eb1775f97575e3a3","spentIndex":0,"spentHeight":-1},{"value":"0.08612100","n":1,"scriptPubKey":{"hex":"76a914eee1c4192e8d2bdbfc4803127af04d68aa39ffac88ac","asm":"OP_DUP OP_HASH160 eee1c4192e8d2bdbfc4803127af04d68aa39ffac OP_EQUALVERIFY OP_CHECKSIG","addresses":["1Nn6H9saJCYen2QcyzMnVQ1wJmi6uwLhd6"],"type":"pubkeyhash"},"spentTxId":null,"spentIndex":null,"spentHeight":null}],"blockhash":"00000000000000000044a01dc989bf8c097c3270a42eaf77200649ed9830ed8b","blockheight":510292,"confirmations":1,"time":1519243687,"blocktime":1519243687,"valueOut":0.78848839,"size":226,"valueIn":0.79268839,"fees":0.0042},{"txid":"ca1afd6fccbacf0e058b4ff881988b05cdf49f92ca3b9d4e0d3988af825d41d1","version":1,"locktime":0,"vin":[{"txid":"0ad3e73171ed2308991b72db6ffbab4884e803878ab664dd8f70f0cb4bf48d65","vout":3,"sequence":4294967295,"n":0,"scriptSig":{"hex":"","asm":null},"addr":null,"valueSat":74264913,"value":0.74264913,"doubleSpentTxID":null}],"vout":[{"value":"0.56930000","n":0,"scriptPubKey":{"hex":"76a9140b23428dafac37222f22fb7b9256f4452106f5cb88ac","asm":"OP_DUP OP_HASH160 0b23428dafac37222f22fb7b9256f4452106f5cb OP_EQUALVERIFY OP_CHECKSIG","addresses":["121tg1wCDzzA97bAMkFeqRHm3Gcn8uZm73"],"type":"pubkeyhash"},"spentTxId":"a7e8cc906e02746d7d2846a241f0a148dea72f4e96af26b12237303d77c793da","spentIndex":0,"spentHeight":-1},{"value":"0.17154913","n":1,"scriptPubKey":{"hex":"0020701a8d401c84fb13e6baf169d59684e17abd9fa216c8cc5b9fc63d622ff8c58d","asm":"0 701a8d401c84fb13e6baf169d59684e17abd9fa216c8cc5b9fc63d622ff8c58d"},"spentTxId":null,"spentIndex":null,"spentHeight":null}],"blockhash":"00000000000000000044a01dc989bf8c097c3270a42eaf77200649ed9830ed8b","blockheight":510292,"confirmations":1,"time":1519243687,"blocktime":1519243687,"valueOut":0.74084913,"size":192,"valueIn":0.74264913,"fees":0.0018},{"txid":"538f24f7e9141cde42cd9a6d71e8076144391918dbc785e51a96d5d5c10fdf8c","version":2,"locktime":510290,"vin":[{"txid":"7e8e2c0fe0658b45092b89acdc85981e5b9f53aa164c538cb704350129a97e34","vout":0,"sequence":4294967294,"n":0,"scriptSig":{"hex":"473044022071e86fb14e29f3bf153c169f36fe2de5de025c15b1a7b84d6e6a08127360ea1c022025ef4207e32895c0b477f30b778f85ca52f56ee2dd7b6d9c309588feaf3ccba601210357c80a979d767ee1dc9e631037e4d1d4046875f61cda9aa2d40c5189af989deb","asm":"3044022071e86fb14e29f3bf153c169f36fe2de5de025c15b1a7b84d6e6a08127360ea1c022025ef4207e32895c0b477f30b778f85ca52f56ee2dd7b6d9c309588feaf3ccba6[ALL] 0357c80a979d767ee1dc9e631037e4d1d4046875f61cda9aa2d40c5189af989deb"},"addr":"17fKzf61BDJj3AL4Bbn4P32fueogduPshR","valueSat":1493324570,"value":14.9332457,"doubleSpentTxID":null}],"vout":[{"value":"8.82763770","n":0,"scriptPubKey":{"hex":"76a914f52d0f3a578d8ca22c1dd981c4834fedce471bb688ac","asm":"OP_DUP OP_HASH160 f52d0f3a578d8ca22c1dd981c4834fedce471bb6 OP_EQUALVERIFY OP_CHECKSIG","addresses":["1PMNXnc4qdUjgyRQaYmcmZ5MMUs51S6jgU"],"type":"pubkeyhash"},"spentTxId":null,"spentIndex":null,"spentHeight":null},{"value":"6.10380000","n":1,"scriptPubKey":{"hex":"76a9148a2202c9b0947f3761be360b942839383319669688ac","asm":"OP_DUP OP_HASH160 8a2202c9b0947f3761be360b9428393833196696 OP_EQUALVERIFY OP_CHECKSIG","addresses":["1DbNzbG15Akt3LhezyqRTGBogYeQxTqGfd"],"type":"pubkeyhash"},"spentTxId":null,"spentIndex":null,"spentHeight":null}],"blockhash":"00000000000000000044a01dc989bf8c097c3270a42eaf77200649ed9830ed8b","blockheight":510292,"confirmations":1,"time":1519243687,"blocktime":1519243687,"valueOut":14.9314377,"size":225,"valueIn":14.9332457,"fees":0.001808},{"txid":"491c653c35263bccd716bb586843f38bf6b9d7535a5b4dfb3cf319ddced668a4","version":2,"locktime":510195,"vin":[{"txid":"26bb721d61b18e4aac1aa151aa86f87e877c88f01464fcfa38c5ba5950c9cf1f","vout":1,"sequence":4294967294,"n":0,"scriptSig":{"hex":"47304402200ac047839ae66617005e3bb0ffcbbb626c401b431559f3033d2ceb78e88ffee8022052c2e9d3f5bc306ec3133ed8780d09208937b5a063e2936a98b9367a44f2794c0121031f15f8bc5a96b60d8d97e7714113232641db7ffaf386ecaa97213126cc23c656","asm":"304402200ac047839ae66617005e3bb0ffcbbb626c401b431559f3033d2ceb78e88ffee8022052c2e9d3f5bc306ec3133ed8780d09208937b5a063e2936a98b9367a44f2794c[ALL] 031f15f8bc5a96b60d8d97e7714113232641db7ffaf386ecaa97213126cc23c656"},"addr":"1CRkbGiPVepjJGGf4W1tmJnLy61hYH1kVM","valueSat":33900000,"value":0.339,"doubleSpentTxID":null},{"txid":"ce7b3c4e196946bd49d0e0ec69a3d939e169a7fa799829db54eb4bf0b53c18bb","vout":0,"sequence":4294967294,"n":1,"scriptSig":{"hex":"47304402205b9631f5802274430f9604de670e49a8a9ead20bb4e1608d999fa3c398bdef7d0220739aeae15cfda536ef227670e4278abe619f67fe8635f7dab0b5884d26124b28012102f9aa542f85fa256c417ea5e586ebd0aca31bbf1c52e502f2d2a44831519ca4e1","asm":"304402205b9631f5802274430f9604de670e49a8a9ead20bb4e1608d999fa3c398bdef7d0220739aeae15cfda536ef227670e4278abe619f67fe8635f7dab0b5884d26124b28[ALL] 02f9aa542f85fa256c417ea5e586ebd0aca31bbf1c52e502f2d2a44831519ca4e1"},"addr":"1MzyVR2b9KJeuznNdEsmpMnjySWGBAGmFu","valueSat":43329122,"value":0.43329122,"doubleSpentTxID":null}],"vout":[{"value":"0.09479622","n":0,"scriptPubKey":{"hex":"76a914d89a3068850eb59c0d23414156cc848ef5bd46dd88ac","asm":"OP_DUP OP_HASH160 d89a3068850eb59c0d23414156cc848ef5bd46dd OP_EQUALVERIFY OP_CHECKSIG","addresses":["1LkHfgXZ6zphaj66GKG9EaHNBjXCQA4YP4"],"type":"pubkeyhash"},"spentTxId":"715b17f949ed6ee3000f90c8a085f0b4af6bbe9923e0827f4773140800232013","spentIndex":0,"spentHeight":510292},{"value":"0.67469000","n":1,"scriptPubKey":{"hex":"76a91499d34bb2dde0a79f9a0f073a5606fe01f8381a6788ac","asm":"OP_DUP OP_HASH160 99d34bb2dde0a79f9a0f073a5606fe01f8381a67 OP_EQUALVERIFY OP_CHECKSIG","addresses":["1F2MWS62vKcSut2kUUkdoXgU1o4BZvMVFq"],"type":"pubkeyhash"},"spentTxId":null,"spentIndex":null,"spentHeight":null}],"blockhash":"00000000000000000044a01dc989bf8c097c3270a42eaf77200649ed9830ed8b","blockheight":510292,"confirmations":1,"time":1519243687,"blocktime":1519243687,"valueOut":0.76948622,"size":372,"valueIn":0.77229122,"fees":0.002805},{"txid":"715b17f949ed6ee3000f90c8a085f0b4af6bbe9923e0827f4773140800232013","version":2,"locktime":510291,"vin":[{"txid":"491c653c35263bccd716bb586843f38bf6b9d7535a5b4dfb3cf319ddced668a4","vout":0,"sequence":4294967294,"n":0,"scriptSig":{"hex":"47304402205eee7a98eb6ca81848ed51d1c11350b4123fe66460ad2e23b50a003e6cf9db4702205b29eb33f65d797459350e044f4d6fdaa4ab60f87b696b421e7a5d1bd9523b6e0121021729e45429fc55c9bcd145f570d53b9e466692cb6848c4d6f3ea4b111a6ee48c","asm":"304402205eee7a98eb6ca81848ed51d1c11350b4123fe66460ad2e23b50a003e6cf9db4702205b29eb33f65d797459350e044f4d6fdaa4ab60f87b696b421e7a5d1bd9523b6e[ALL] 021729e45429fc55c9bcd145f570d53b9e466692cb6848c4d6f3ea4b111a6ee48c"},"addr":"1LkHfgXZ6zphaj66GKG9EaHNBjXCQA4YP4","valueSat":9479622,"value":0.09479622,"doubleSpentTxID":null}],"vout":[{"value":"0.05933564","n":0,"scriptPubKey":{"hex":"76a914ddb6ad4b18b2498f6bd70d072cac8da2e5f2c0e688ac","asm":"OP_DUP OP_HASH160 ddb6ad4b18b2498f6bd70d072cac8da2e5f2c0e6 OP_EQUALVERIFY OP_CHECKSIG","addresses":["1MDKB1aRhDuNkEdAjwU1wUs6nBBMfTczR5"],"type":"pubkeyhash"},"spentTxId":null,"spentIndex":null,"spentHeight":null},{"value":"0.03376558","n":1,"scriptPubKey":{"hex":"76a914ff3612040c625d6a39662a79066d13ec34d8f0a788ac","asm":"OP_DUP OP_HASH160 ff3612040c625d6a39662a79066d13ec34d8f0a7 OP_EQUALVERIFY OP_CHECKSIG","addresses":["1QGS5nx3QNGrs25BarGwaxs7xthkHPXMXg"],"type":"pubkeyhash"},"spentTxId":null,"spentIndex":null,"spentHeight":null}],"blockhash":"00000000000000000044a01dc989bf8c097c3270a42eaf77200649ed9830ed8b","blockheight":510292,"confirmations":1,"time":1519243687,"blocktime":1519243687,"valueOut":0.09310122,"size":225,"valueIn":0.09479622,"fees":0.001695},{"txid":"750696b8096681029b175cdf4c7e4c756200a48bd7ca552b5de86694974a15a5","version":2,"locktime":510291,"vin":[{"txid":"cdec78eef9881344d44221c9bef3ae176bdea578b50cda38771054e7a847520a","vout":1,"sequence":4294967294,"n":0,"scriptSig":{"hex":"48304502210085e7dab7385f971b38f64a13be03d41ac563ebc4af8149e447a7ecdb71179e4a0220578d51c2e19c3954265b97a4452e2d532af873cc742cbf149cf992069095002f012103af6a2b019430e5f0e8d8ccffa45d6512a867b9b1ed8b5058a40bec8bd78dbb5a","asm":"304502210085e7dab7385f971b38f64a13be03d41ac563ebc4af8149e447a7ecdb71179e4a0220578d51c2e19c3954265b97a4452e2d532af873cc742cbf149cf992069095002f[ALL] 03af6a2b019430e5f0e8d8ccffa45d6512a867b9b1ed8b5058a40bec8bd78dbb5a"},"addr":"16Y1JS8dg3p9Yh6bbz6ogRMp9qjUdu5xZB","valueSat":7818703,"value":0.07818703,"doubleSpentTxID":null},{"txid":"94fdc4fa7756e3ac41fb6709d3ce4cd28fb8b2bcdeb7004b9dc6154584186c13","vout":0,"sequence":4294967294,"n":1,"scriptSig":{"hex":"47304402206463027fb55a2027eac0c4b1649b4311d92896812ab3dddadf5d2b046bb13b1e02204885d96d609ec938f91f8bcdd12fdf6dc346eec48fe9d921fbdd150b5201f6a90121028b40bd6109e4695112356a4bef90f1a1a14840b03e1dcc140db59cf869dfc616","asm":"304402206463027fb55a2027eac0c4b1649b4311d92896812ab3dddadf5d2b046bb13b1e02204885d96d609ec938f91f8bcdd12fdf6dc346eec48fe9d921fbdd150b5201f6a9[ALL] 028b40bd6109e4695112356a4bef90f1a1a14840b03e1dcc140db59cf869dfc616"},"addr":"112veNaQPnTSwN4ikKAfg4Sksio2UQLgwp","valueSat":129200,"value":0.001292,"doubleSpentTxID":null},{"txid":"cbc4073450b245f995da1c75c43d6b1db85decbbc237a7ec8fd4d342a794d025","vout":1,"sequence":4294967294,"n":2,"scriptSig":{"hex":"4730440220467e273ac136d5bab7c9aa23e8c32b0bb78457381b504128808ebed851c9f7c302202ae37e3653697d893d56efd124dd710a820ccc8aa6b458bda885a1aa4b757eac0121033cf859e2d40d94a317fb20dd0bc390969fb7afd8c6b66b441b57519675acf31c","asm":"30440220467e273ac136d5bab7c9aa23e8c32b0bb78457381b504128808ebed851c9f7c302202ae37e3653697d893d56efd124dd710a820ccc8aa6b458bda885a1aa4b757eac[ALL] 033cf859e2d40d94a317fb20dd0bc390969fb7afd8c6b66b441b57519675acf31c"},"addr":"1NRXRkY4Ho1VZrqhLDo36YeNFkjvHiVNyx","valueSat":1165464,"value":0.01165464,"doubleSpentTxID":null},{"txid":"10ea2a882cfb9635719813b102805f9839403cd78f2c6a7594a09ae3f0f69a28","vout":1,"sequence":4294967294,"n":3,"scriptSig":{"hex":"483045022100f29fb26cbec40dceeae4b8cf6147c8ff9867657f9dbb6017771665dcbce8ebd002201fa8e03c2c74eb3d6e0012ad9ebe1fc484cfe81386354bdb0e4024d295f7b10c012102e2acca348566081a848b70411231dca93555c1a71c5f77e64fc63e994296d12d","asm":"3045022100f29fb26cbec40dceeae4b8cf6147c8ff9867657f9dbb6017771665dcbce8ebd002201fa8e03c2c74eb3d6e0012ad9ebe1fc484cfe81386354bdb0e4024d295f7b10c[ALL] 02e2acca348566081a848b70411231dca93555c1a71c5f77e64fc63e994296d12d"},"addr":"1Pw4xaz1QMTyz6DLpS3N2SrF26GaiqJfiB","valueSat":759544,"value":0.00759544,"doubleSpentTxID":null},{"txid":"ca394a1aa9285bc048afa15b7c9caea28c91e311ea3f853ed5446d0919b83859","vout":1,"sequence":4294967294,"n":4,"scriptSig":{"hex":"473044022003b1921e8780beb194890c920a0a769c1cd5e5de1d83a7c13f9585d20909778f022038803d4d68d2b8452e927fb8e8e900209fd2769b0aacae5caa6a9a56db305adb012103d038a62334671b8d86a807f6448dc7d278ed0e3094b95ca078341883b8735a4d","asm":"3044022003b1921e8780beb194890c920a0a769c1cd5e5de1d83a7c13f9585d20909778f022038803d4d68d2b8452e927fb8e8e900209fd2769b0aacae5caa6a9a56db305adb[ALL] 03d038a62334671b8d86a807f6448dc7d278ed0e3094b95ca078341883b8735a4d"},"addr":"1D7sB8CYXB7W7aTHV4Dh8sinub2jwayanP","valueSat":480000,"value":0.0048,"doubleSpentTxID":null},{"txid":"1aace6b5bb36623a7849200ca72dc13fd88093d48cda20a87bc5c480188ceb5c","vout":370,"sequence":4294967294,"n":5,"scriptSig":{"hex":"483045022100f75648621496731fcec758e4566d9ecc9002fcfb612a57038f85b00b6ce1abee0220435b79a1340defecac975545681800acff3a5e6900d3d1a9a4b19a01bf0d66cb012102acb8c2615b4c4bf139599fa6d73e0373f1aa8a95f490d48bbdcdb371806d65da","asm":"3045022100f75648621496731fcec758e4566d9ecc9002fcfb612a57038f85b00b6ce1abee0220435b79a1340defecac975545681800acff3a5e6900d3d1a9a4b19a01bf0d66cb[ALL] 02acb8c2615b4c4bf139599fa6d73e0373f1aa8a95f490d48bbdcdb371806d65da"},"addr":"17KkUoATWjU9QAm9qLJJ4Y7BkT3R2eBHLR","valueSat":956110,"value":0.0095611,"doubleSpentTxID":null},{"txid":"74b305ab9bfb9220524ff3d5244ae87b016ef915a2ac2b4789c51f404fad1184","vout":0,"sequence":4294967294,"n":6,"scriptSig":{"hex":"47304402207b861c3c41c938e3e025be6a777e544c5370b7664ce53f65194cddda9ff8b60f02203ac00bbde42d1df32b216331ec3ea03d06c63c27008cb984b738e50a1670e87b012103d55c12f1171177e24168b92663fb889cae96b2b96a3e5f8e6608db50d7dd2bfc","asm":"304402207b861c3c41c938e3e025be6a777e544c5370b7664ce53f65194cddda9ff8b60f02203ac00bbde42d1df32b216331ec3ea03d06c63c27008cb984b738e50a1670e87b[ALL] 03d55c12f1171177e24168b92663fb889cae96b2b96a3e5f8e6608db50d7dd2bfc"},"addr":"1AXCH7ozuNUPcxVNUpvPuVYCYErvfmZu6W","valueSat":1333404,"value":0.01333404,"doubleSpentTxID":null},{"txid":"ca0180bb704aa818223a16cb4d451436bf80001f793f35a0f0731c73a56d548d","vout":0,"sequence":4294967294,"n":7,"scriptSig":{"hex":"473044022029fb14f0d2bca7da5facdd13902eee468048a4c47a3b1b5ebad985d83574028802207fa24ff594416d3daf363317412c1292bf4e03789f0ac35e2a69d1e5b16dc39b0121024c4ab1a8c2c8cea8deba44395a14c5b9247bef8817b129d439b0e6d41cb02847","asm":"3044022029fb14f0d2bca7da5facdd13902eee468048a4c47a3b1b5ebad985d83574028802207fa24ff594416d3daf363317412c1292bf4e03789f0ac35e2a69d1e5b16dc39b[ALL] 024c4ab1a8c2c8cea8deba44395a14c5b9247bef8817b129d439b0e6d41cb02847"},"addr":"1GByFyr8hSD9SbVe23MRECHVY8733G71RL","valueSat":623007,"value":0.00623007,"doubleSpentTxID":null},{"txid":"25425969579628d4d9ce45836e5724c151f0d088eb3f2cb92435f5b938ad16c4","vout":2,"sequence":4294967294,"n":8,"scriptSig":{"hex":"483045022100ff3b0c4b097b470989399e7032b9a24bf2a178ad76f9ad01862725c9fd3e961902203b84371fc7672b604ed42d074b0702b48a1f2845d5a6b7f7ded4c43c2cf9af380121022ada8cb6d9c369c52f4f7fce6e6bc2c7165b1e4e7b44b88b6ebf29297ab15c72","asm":"3045022100ff3b0c4b097b470989399e7032b9a24bf2a178ad76f9ad01862725c9fd3e961902203b84371fc7672b604ed42d074b0702b48a1f2845d5a6b7f7ded4c43c2cf9af38[ALL] 022ada8cb6d9c369c52f4f7fce6e6bc2c7165b1e4e7b44b88b6ebf29297ab15c72"},"addr":"1Bg52ooUy2dsdii5dzX1Lv9PgfRWF176Nn","valueSat":3607357,"value":0.03607357,"doubleSpentTxID":null},{"txid":"46f220e8928c42607dd51f0861991050c34143d7e16ddbefe40350f02d5a2ede","vout":79,"sequence":4294967294,"n":9,"scriptSig":{"hex":"473044022050f228da6d137ab83a4d86c7dad99f9e41e729eacc61e7e1380bfbf29633523f0220107942cc67397be4f2f920ddcb4601343fe098b4d1e1d5202fd115176cf3167601210206434f32bd3b4856712acd8d7292e3ada65c19608b7c0ee73e117c508a1b7c75","asm":"3044022050f228da6d137ab83a4d86c7dad99f9e41e729eacc61e7e1380bfbf29633523f0220107942cc67397be4f2f920ddcb4601343fe098b4d1e1d5202fd115176cf31676[ALL] 0206434f32bd3b4856712acd8d7292e3ada65c19608b7c0ee73e117c508a1b7c75"},"addr":"1AXfxkgeuNz7Pq1P1UM3znUbFbW7hvxUPk","valueSat":1572155,"value":0.01572155,"doubleSpentTxID":null},{"txid":"149a849cdf86e99f12d01d8a8ae7bbc97b90d5b34312a8015a7a9a9019f810df","vout":20,"sequence":4294967294,"n":10,"scriptSig":{"hex":"47304402207ec52b7020dd5fcd94602a0ebb4ffbd874db429273ce6100836773c4c0570291022011d20356893deecde077897fc062ae06faa629aac905356cd7318606df3c3f2d0121038d588053a29779e3c2b10d9be59b9695c091d62268ed05cf5329a58f89373c2b","asm":"304402207ec52b7020dd5fcd94602a0ebb4ffbd874db429273ce6100836773c4c0570291022011d20356893deecde077897fc062ae06faa629aac905356cd7318606df3c3f2d[ALL] 038d588053a29779e3c2b10d9be59b9695c091d62268ed05cf5329a58f89373c2b"},"addr":"1EruNcryxv71TTMvBJC2w7kYm8NPZ8Sapv","valueSat":4039794,"value":0.04039794,"doubleSpentTxID":null},{"txid":"40e80072657f499707378b4d8d41a4a5313cb3ae011a4556aff6f8a57a3ad9df","vout":0,"sequence":4294967294,"n":11,"scriptSig":{"hex":"483045022100d52c9d62af91d1cba542f76da4bb7dff91c0a85bcda106cb22179ef042e87be102205b61261dd95d7baf156c195c33d3b28f320699a41684febec13ea062e646be63012103f73f2eb0d2d8afb7ad8ff56a86ab6629afdf2eebf5e89ccfa76094c3bf793fb7","asm":"3045022100d52c9d62af91d1cba542f76da4bb7dff91c0a85bcda106cb22179ef042e87be102205b61261dd95d7baf156c195c33d3b28f320699a41684febec13ea062e646be63[ALL] 03f73f2eb0d2d8afb7ad8ff56a86ab6629afdf2eebf5e89ccfa76094c3bf793fb7"},"addr":"19GgQMspuAMRdHXAnN8NWCFwu5arjCDjhv","valueSat":4776693,"value":0.04776693,"doubleSpentTxID":null},{"txid":"fc3d4d407bd24a009b96d7b52c6cc999e9f32ca6db09bea9f766e39a66ce03fa","vout":0,"sequence":4294967294,"n":12,"scriptSig":{"hex":"483045022100d219cfa68214c4012340ea7fc232059b7b0e4be60b8566aafea9ebcef7ce53300220268f85721c88e38e5549c22a534557e44d425d544108d330e7b2ec6f57323dd1012103bad81316707857d7ef77232de95d45a5bc4c13cd95538e80e958d396b163fd10","asm":"3045022100d219cfa68214c4012340ea7fc232059b7b0e4be60b8566aafea9ebcef7ce53300220268f85721c88e38e5549c22a534557e44d425d544108d330e7b2ec6f57323dd1[ALL] 03bad81316707857d7ef77232de95d45a5bc4c13cd95538e80e958d396b163fd10"},"addr":"19dcxxWxp4DjkHwkaChgvqHYZVmR6mYjgk","valueSat":71060,"value":0.0007106,"doubleSpentTxID":null},{"txid":"7e71b9f603499769fb798ef8b1e830ac325992ed6346eb56bf9b153a6218f1fd","vout":0,"sequence":4294967294,"n":13,"scriptSig":{"hex":"483045022100d5078535b006baf14f4318e0400cee04609b08e303d14d1e67c585bb8fbb1479022066a1de6d6db5d9cb50177b405cfc48f409b303c9ae3a4fd9e0a136c0ab778777012103eb9d7fa585ed54de0985774e1a40124589c94ef5cae0b73d6adaf0db1bca3a5c","asm":"3045022100d5078535b006baf14f4318e0400cee04609b08e303d14d1e67c585bb8fbb1479022066a1de6d6db5d9cb50177b405cfc48f409b303c9ae3a4fd9e0a136c0ab778777[ALL] 03eb9d7fa585ed54de0985774e1a40124589c94ef5cae0b73d6adaf0db1bca3a5c"},"addr":"1Hgc55z8UzhT8mTTZH9xwcADaZZdSN2SZz","valueSat":57976,"value":0.00057976,"doubleSpentTxID":null}],"vout":[{"value":"0.25654670","n":0,"scriptPubKey":{"hex":"76a91413bed82761836a74906e39bdac5b692fa90d12d688ac","asm":"OP_DUP OP_HASH160 13bed82761836a74906e39bdac5b692fa90d12d6 OP_EQUALVERIFY OP_CHECKSIG","addresses":["12oQTk5bGMhGRCX7AkwXD9jBTLFZr1YZMH"],"type":"pubkeyhash"},"spentTxId":null,"spentIndex":null,"spentHeight":null},{"value":"0.00123297","n":1,"scriptPubKey":{"hex":"76a914c12c12c6b3239d00341d2d4a3eaea868d2108f4188ac","asm":"OP_DUP OP_HASH160 c12c12c6b3239d00341d2d4a3eaea868d2108f41 OP_EQUALVERIFY OP_CHECKSIG","addresses":["1JcQDHzQSdRvz8YE7mKmuJScJVgvrDknpP"],"type":"pubkeyhash"},"spentTxId":null,"spentIndex":null,"spentHeight":null}],"blockhash":"00000000000000000044a01dc989bf8c097c3270a42eaf77200649ed9830ed8b","blockheight":510292,"confirmations":1,"time":1519243687,"blocktime":1519243687,"valueOut":0.25777967,"size":2143,"valueIn":0.27390467,"fees":0.016125},{"txid":"0934461ba0b92a30b47a877d0390e81dd2ecbc7b4c3c5410547e4ce6f8142b39","version":2,"locktime":510291,"vin":[{"txid":"ad376165e33f8d482ed7add195e1af1013b5620e64c06a3d731711e3004fd114","vout":20,"sequence":4294967294,"n":0,"scriptSig":{"hex":"4830450221008c47064283cf0f428c88d5b312f1d5c54cea326e4829cf9f2c66c5b6e816370402200a0b90ae4056e0998f9ffa35b8825375b391910044aca144b39da0871cc2d6850121021f710666649165283df3464a1165b08fd14e35d0d7a107162a4d3d8ede5f60ec","asm":"30450221008c47064283cf0f428c88d5b312f1d5c54cea326e4829cf9f2c66c5b6e816370402200a0b90ae4056e0998f9ffa35b8825375b391910044aca144b39da0871cc2d685[ALL] 021f710666649165283df3464a1165b08fd14e35d0d7a107162a4d3d8ede5f60ec"},"addr":"1Kay1mJhvnE4LZi4z9X67g1us1WXazssMp","valueSat":1500000000,"value":15,"doubleSpentTxID":null}],"vout":[{"value":"1.72433960","n":0,"scriptPubKey":{"hex":"a9143c82bf8a0b34131d48c5e5864d20d20acd50c87787","asm":"OP_HASH160 3c82bf8a0b34131d48c5e5864d20d20acd50c877 OP_EQUAL","addresses":["37Cy6bdD3JUAxHqE2v3tGXvCr42iC9qk8h"],"type":"scripthash"},"spentTxId":null,"spentIndex":null,"spentHeight":null},{"value":"13.27398040","n":1,"scriptPubKey":{"hex":"76a914feb63e990198dbe91c5c20b594ee5a2e9322905688ac","asm":"OP_DUP OP_HASH160 feb63e990198dbe91c5c20b594ee5a2e93229056 OP_EQUALVERIFY OP_CHECKSIG","addresses":["1QDnxJBpDxAo2eCRMy7gShEzPDk2YhyyKi"],"type":"pubkeyhash"},"spentTxId":"031f49889f31c8bc3b9e369e81c1526f5afc496dfc67428d10bc3d09cd75810e","spentIndex":0,"spentHeight":510292}],"blockhash":"00000000000000000044a01dc989bf8c097c3270a42eaf77200649ed9830ed8b","blockheight":510292,"confirmations":1,"time":1519243687,"blocktime":1519243687,"valueOut":14.99832,"size":224,"valueIn":15,"fees":0.00168},{"txid":"031f49889f31c8bc3b9e369e81c1526f5afc496dfc67428d10bc3d09cd75810e","version":2,"locktime":510291,"vin":[{"txid":"0934461ba0b92a30b47a877d0390e81dd2ecbc7b4c3c5410547e4ce6f8142b39","vout":1,"sequence":4294967294,"n":0,"scriptSig":{"hex":"483045022100a59b101b8d7c98a5cc8ef02f07b2634375c59d97f33284c2b876d1e963bb296802207d528b1afb1528d249c56753eee7993ffb19d438314709cdbedfb4028fe458fc012103a2a830928db9fd0ba899a353eaefe955edad066bd412a59c288de724a26b5242","asm":"3045022100a59b101b8d7c98a5cc8ef02f07b2634375c59d97f33284c2b876d1e963bb296802207d528b1afb1528d249c56753eee7993ffb19d438314709cdbedfb4028fe458fc[ALL] 03a2a830928db9fd0ba899a353eaefe955edad066bd412a59c288de724a26b5242"},"addr":"1QDnxJBpDxAo2eCRMy7gShEzPDk2YhyyKi","valueSat":1327398040,"value":13.2739804,"doubleSpentTxID":null}],"vout":[{"value":"0.35252419","n":0,"scriptPubKey":{"hex":"76a914e507bbbd92726c9248a705113d950aeecd3d54ff88ac","asm":"OP_DUP OP_HASH160 e507bbbd92726c9248a705113d950aeecd3d54ff OP_EQUALVERIFY OP_CHECKSIG","addresses":["1Mt11FeXgidoxtBKEqEKdHZvaCXnm3HUE1"],"type":"pubkeyhash"},"spentTxId":null,"spentIndex":null,"spentHeight":null},{"value":"12.91976121","n":1,"scriptPubKey":{"hex":"76a9145d912972362982e40ac17cbcd85323ed8667d9f388ac","asm":"OP_DUP OP_HASH160 5d912972362982e40ac17cbcd85323ed8667d9f3 OP_EQUALVERIFY OP_CHECKSIG","addresses":["19XjjsDX2zKE21gxUZcut9npJ89TWVsn5K"],"type":"pubkeyhash"},"spentTxId":null,"spentIndex":null,"spentHeight":null}],"blockhash":"00000000000000000044a01dc989bf8c097c3270a42eaf77200649ed9830ed8b","blockheight":510292,"confirmations":1,"time":1519243687,"blocktime":1519243687,"valueOut":13.2722854,"size":226,"valueIn":13.2739804,"fees":0.001695},{"txid":"aa80905f4e439a635337688563ee1bc3b21fa3aa5a3faa559ea975c43b3e89c3","version":2,"locktime":510291,"vin":[{"txid":"f76ad06472fc36acb36e933ead4444af576a7a89f95b04b3b812af629e56764b","vout":0,"sequence":4294967294,"n":0,"scriptSig":{"hex":"483045022100d9f8abb810b90f9360d20c01cef3fcfab4a7513488e82ab800b7a04d76a95735022041430b4c894358a39d1882c0be8193384724a9e08d265ccb94c12b707c4bec50012103eb0cdb2a6884dcf158bb8824eb2ce8c726c77a8983c91a1dd3202a5ccb90517b","asm":"3045022100d9f8abb810b90f9360d20c01cef3fcfab4a7513488e82ab800b7a04d76a95735022041430b4c894358a39d1882c0be8193384724a9e08d265ccb94c12b707c4bec50[ALL] 03eb0cdb2a6884dcf158bb8824eb2ce8c726c77a8983c91a1dd3202a5ccb90517b"},"addr":"1HgLUHnaU4QHyrfzLMXJUc37yYU7tgbdPR","valueSat":1345018802,"value":13.45018802,"doubleSpentTxID":null}],"vout":[{"value":"11.35353822","n":0,"scriptPubKey":{"hex":"76a9142e1ca6360733300a5bc690e7150a07d9531c462a88ac","asm":"OP_DUP OP_HASH160 2e1ca6360733300a5bc690e7150a07d9531c462a OP_EQUALVERIFY OP_CHECKSIG","addresses":["15CpQQmcZwkLYLqxjzezxqQXR9Hg9gehPB"],"type":"pubkeyhash"},"spentTxId":"614b4426e085c8f115ec70d48d4363b4d08fd9c656891e028e983375511d9b0e","spentIndex":0,"spentHeight":510292},{"value":"2.09496980","n":1,"scriptPubKey":{"hex":"a9144cb8b8281e23f428299d3e0d5aec6d9342e8678e87","asm":"OP_HASH160 4cb8b8281e23f428299d3e0d5aec6d9342e8678e OP_EQUAL","addresses":["38ggZec12MGsjJkhjU5dZ9wfs93Kc2RbwT"],"type":"scripthash"},"spentTxId":null,"spentIndex":null,"spentHeight":null}],"blockhash":"00000000000000000044a01dc989bf8c097c3270a42eaf77200649ed9830ed8b","blockheight":510292,"confirmations":1,"time":1519243687,"blocktime":1519243687,"valueOut":13.44850802,"size":224,"valueIn":13.45018802,"fees":0.00168}]}`
	response := TxsResponse{}
	json.Unmarshal([]byte(jsonString), &response);
	
	js, err := json.Marshal(response)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
