package nebulasio

import (
	"fmt"
	"github.com/blocktree/OpenWallet/crypto"
	"github.com/blocktree/OpenWallet/openwallet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
	"strconv"
)






//BlockchainInfo 本地节点区块链信息
type BlockchainInfo struct {
	Chain                string `json:"chain"`
	Blocks               uint64 `json:"blocks"`
	Headers              uint64 `json:"headers"`
	Bestblockhash        string `json:"bestblockhash"`
	Difficulty           string `json:"difficulty"`
	Mediantime           uint64 `json:"mediantime"`
	Verificationprogress string `json:"verificationprogress"`
	Chainwork            string `json:"chainwork"`
	Pruned               bool   `json:"pruned"`
}

func NewBlockchainInfo(json *gjson.Result) *BlockchainInfo {
	b := &BlockchainInfo{}
	//解析json
	b.Chain = gjson.Get(json.Raw, "chain").String()
	b.Blocks = gjson.Get(json.Raw, "blocks").Uint()
	b.Headers = gjson.Get(json.Raw, "headers").Uint()
	b.Bestblockhash = gjson.Get(json.Raw, "bestblockhash").String()
	b.Difficulty = gjson.Get(json.Raw, "difficulty").String()
	b.Mediantime = gjson.Get(json.Raw, "mediantime").Uint()
	b.Verificationprogress = gjson.Get(json.Raw, "verificationprogress").String()
	b.Chainwork = gjson.Get(json.Raw, "chainwork").String()
	b.Pruned = gjson.Get(json.Raw, "pruned").Bool()
	return b
}

//Unspent 未花记录
type Unspent struct {

	/*
			{
		        "txid" : "d54994ece1d11b19785c7248868696250ab195605b469632b7bd68130e880c9a",
		        "vout" : 1,
		        "address" : "mgnucj8nYqdrPFh2JfZSB1NmUThUGnmsqe",
		        "account" : "test label",
		        "scriptPubKey" : "76a9140dfc8bafc8419853b34d5e072ad37d1a5159f58488ac",
		        "amount" : 0.00010000,
		        "confirmations" : 6210,
		        "spendable" : true,
		        "solvable" : true
		    }
	*/
	Key           string `storm:"id"`
	TxID          string `json:"txid"`
	Vout          uint64 `json:"vout"`
	Address       string `json:"address"`
	AccountID     string `json:"account" storm:"index"`
	ScriptPubKey  string `json:"scriptPubKey"`
	Amount        string `json:"amount"`
	Confirmations uint64 `json:"confirmations"`
	Spendable     bool   `json:"spendable"`
	Solvable      bool   `json:"solvable"`
	HDAddress     openwallet.Address
}

func NewUnspent(json *gjson.Result) *Unspent {
	obj := &Unspent{}
	//解析json
	obj.TxID = gjson.Get(json.Raw, "txid").String()
	obj.Vout = gjson.Get(json.Raw, "vout").Uint()
	obj.Address = gjson.Get(json.Raw, "address").String()
	obj.AccountID = gjson.Get(json.Raw, "account").String()
	obj.ScriptPubKey = gjson.Get(json.Raw, "scriptPubKey").String()
	obj.Amount = gjson.Get(json.Raw, "amount").String()
	obj.Confirmations = gjson.Get(json.Raw, "confirmations").Uint()
	//obj.Spendable = gjson.Get(json.Raw, "spendable").Bool()
	obj.Spendable = true
	obj.Solvable = gjson.Get(json.Raw, "solvable").Bool()

	return obj
}

type UnspentSort struct {
	values     []*Unspent
	comparator func(a, b *Unspent) int
}

func (s UnspentSort) Len() int {
	return len(s.values)
}
func (s UnspentSort) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}
func (s UnspentSort) Less(i, j int) bool {
	return s.comparator(s.values[i], s.values[j]) < 0
}

//type Address struct {
//	Address   string `json:"address" storm:"id"`
//	Account   string `json:"account" storm:"index"`
//	HDPath    string `json:"hdpath"`
//	CreatedAt time.Time
//}

type User struct {
	UserKey string `storm:"id"`     // primary key
	Group   string `storm:"index"`  // this field will be indexed
	Email   string `storm:"unique"` // this field will be indexed with a unique constraint
	Name    string // this field will not be indexed
	Age     int    `storm:"index"`
}

type Block struct {
	/*
{
    "result": {
        "hash": "37c1afd37a7340fce54e4b32401e02b537ee7d7d8ff02796e85811a497b5660b",
        "parent_hash": "52f0f98a6f755fa1088b296cc461a7b266c06c284e2fd8f43003de460cc64398",
        "height": "1098023",
        "nonce": "0",
        "coinbase": "n1EzGmFsVepKduN1U5QFyhLqpzFvM9sRSmG",
        "timestamp": "1539314535",
        "chain_id": 1001,
        "state_root": "b481d6a6bddab87558d7b31f590bdaea1b30aaf9d11dee56d8a5f8a27dc79e6b",
        "txs_root": "d9f90ebe015f42651778764bbcae85558de36b696bc8ed45145735b70f6f7b08",
        "events_root": "540004b851de9a5da07b457076a6b398245b0886a10b5815991b38117ca62056",
        "consensus_root": {
            "timestamp": "1539314535",
            "proposer": "GVfjikwHkVvZrFz0l460IVAxpDi6hi9ILA0=",
            "dynasty_root": "Kst/n5bFEWA/FD8qPKn029Y8cE9dfEgp6VMuJp5AGfU="
        },
        "miner": "n1bFyMzjrjML5ax8ry7NQ4dWHHmnBU9ghXi",
        "randomSeed": "e306e461d9c9e7f343fe732347204e1865da008c9fec5658bbdfbf767ae2dca8",
        "randomProof": "f70bc1978048f54971e9dd6a529a94208248ae2f79590f7a17f73e0b708b9b54a15fe0f7e27ead9f0f8f63208811c2b61d1d1fda1959b825dc0d0739ef8ee98304e541261fcd1142700f715d0cd9c9267f8eab1accca1d9a12b37a63f335d00afa6f0850fd8a1902fa1e35aa1b73d9e9ac31a1aabfcff47671d30734187a8a4ee2",
        "is_finality": true,
        "transactions": [
            {
                "hash": "08057a5635865bb8a90fa06e833b1171fdd2eaedb1dc673e44601122f1ff4519",
                "chainId": 1001,
                "from": "n1JdmmyhrrqBuESseZSbrBucnvugSewSMTE",
                "to": "n1wR7zue5zXjUEmQofcTZV8t1H61MbBdpj5",
                "value": "0",
                "nonce": "59",
                "timestamp": "1539314513",
                "type": "call",
                "data": "eyJGdW5jdGlvbiI6ImFkZCIsIkFyZ3MiOiJbXCJuMUpkbW15aHJycUJ1RVNzZVpTYnJCdWNudnVnU2V3U01URVwiLFwi5rWL6K+V77yabjFKZG1teWhycnFCdUVTc2VaU2JyQnVjbnZ1Z1Nld1NNVEVcIl0ifQ==",
                "gas_price": "1000000",
                "gas_limit": "200000",
                "contract_address": "",
                "status": 1,
                "gas_used": "20322",
                "execute_error": "",
                "execute_result": "\"\"",
                "block_height": "1098023"
            }
        ]
    }
}
*/
	Hash              string
	Confirmations     uint64
	Merkleroot        string
	tx                []string //只记录hash(txid)
	Previousblockhash string
	Height            uint64 `storm:"id"`
	Version           uint64
	Time              uint64
	Fork              bool
}

func NewBlock(json *gjson.Result) *Block {
	obj := &Block{}
	//解析json
	obj.Hash = json.Get( "hash").String()
	obj.Previousblockhash = json.Get( "parent_hash").String()
	obj.Height,_ = strconv.ParseUint(json.Get( "height").String(),10,64)
	obj.Time,_ = strconv.ParseUint(json.Get( "timestamp").String(),10,64)
	txs := make([]string, 0)
	for _, tx := range json.Get( "transactions").Array() {
	//	fmt.Printf("tx.Get().String()=%v\n",tx.Get("hash").String())
		txs = append(txs, tx.Get("hash").String())
	}
	obj.tx = txs

	//	obj.Confirmations = gjson.Get(json.Raw, "confirmations").Uint()
	//	obj.Merkleroot = gjson.Get(json.Raw, "merkleroot").String()
	//	obj.Version = gjson.Get(json.Raw, "version").Uint()

	return obj
}

//BlockHeader 区块链头
func (b *Block) BlockHeader() *openwallet.BlockHeader {

	obj := openwallet.BlockHeader{}
	//解析json
	obj.Hash = b.Hash
	obj.Confirmations = b.Confirmations
	obj.Merkleroot = b.Merkleroot
	obj.Previousblockhash = b.Previousblockhash
	obj.Height = b.Height
	obj.Version = b.Version
	obj.Time = b.Time
	obj.Symbol = Symbol

	return &obj
}

//UnscanRecords 扫描失败的区块及交易
type UnscanRecord struct {
	ID          string `storm:"id"` // primary key
	BlockHeight uint64
	TxID        string
	Reason      string
}

func NewUnscanRecord(height uint64, txID, reason string) *UnscanRecord {
	obj := UnscanRecord{}
	obj.BlockHeight = height
	obj.TxID = txID
	obj.Reason = reason
	obj.ID = common.Bytes2Hex(crypto.SHA256([]byte(fmt.Sprintf("%d_%s", height, txID))))
	return &obj
}

type NasTransaction struct {

/*
	TxID          string
	Size          uint64
	Version       uint64
	LockTime      int64
	Hex           string
	BlockHash     string
	BlockHeight   uint64
	Confirmations uint64
	Blocktime     int64
	IsCoinBase    bool
	Fees          string

	Vins  []*Vin
	Vouts []*Vout
*/
	/*
		//nas
		{
			"hash": "c98c1c117442554a674262b6c566e8ba12b7397fdcbf09abb6f3bd06a1646b83",
			"chainId": 100,
			"from": "n1NM2eETQG5Es7cCc7sh29NJr9cP94QZcXR",
			"to": "n1Qzgp7TE4TNQTC9LQpEsuYZFbGU14cpDZk",
			"value": "17000000000000000000",
			"nonce": "39",
			"timestamp": "1",
			"type": "binary",
			"data": null,
			"gas_price": "1000000",
			"gas_limit": "2000000",
			"contract_address": "",
			"status": 1,
			"gas_used": "20000",
			"execute_error": "",
			"execute_result": ""
		}
	*/
	Hash          		string   //txid
	ChainId       		uint64
	From       			string
	To      			string
	Value           	decimal.Decimal
	Nonce     			string
	Timestamp   		string
	Type				string
	Data 				string
	Gas_price     		string
	Gas_limit    		string
	Contract_address    string
	Status				uint64
	Gas_used			string
	Execute_error		string
	Execute_result		string

	FromAccountId    string  //transaction scanning 的时候对其进行赋值
	ToAccountId      string  //transaction scanning 的时候对其进行赋值

	//交易对应的区块信息
	BlockTime        uint64
	BlockHeight		 uint64
	BlockHash		 string
}

/*
type Vin struct {
	Coinbase string
	TxID     string
	Vout     uint64
	N        uint64
	Addr     string
	Value    string
}

type Vout struct {
	N            uint64
	Addr         string
	Value        string
	ScriptPubKey string
	Type         string
}
*/
/*
func newTxByExplorer(json *gjson.Result) *Transaction {

	/*
			{
			"txid": "9f5eae5b95016825a437ceb9c9224d3e30d3b351f1100e4df5cc0cacac4e668c",
			"version": 1,
			"locktime": 1433760,
			"vin": [],
			"vout": [],
			"blockhash": "0000000000003ac968ee1ae321f35f76d4dcb685045968d60fc39edb20b0eed0",
			"blockheight": 1433761,
			"confirmations": 5,
			"time": 1539050096,
			"blocktime": 1539050096,
			"valueOut": 0.14652549,
			"size": 814,
			"valueIn": 0.14668889,
			"fees": 0.0001634
		}
	*/
/*
	obj := Transaction{}
	//解析json
	obj.TxID = gjson.Get(json.Raw, "txid").String()
	obj.Version = gjson.Get(json.Raw, "version").Uint()
	obj.LockTime = gjson.Get(json.Raw, "locktime").Int()
	obj.BlockHash = gjson.Get(json.Raw, "blockhash").String()
	obj.BlockHeight = gjson.Get(json.Raw, "blockheight").Uint()
	obj.Confirmations = gjson.Get(json.Raw, "confirmations").Uint()
	obj.Blocktime = gjson.Get(json.Raw, "blocktime").Int()
	obj.Size = gjson.Get(json.Raw, "size").Uint()
	obj.Fees = gjson.Get(json.Raw, "fees").String()

	obj.Vins = make([]*Vin, 0)
	if vins := gjson.Get(json.Raw, "vin"); vins.IsArray() {
		for _, vin := range vins.Array() {
			input := newTxVinByExplorer(&vin)
			obj.Vins = append(obj.Vins, input)
		}
	}

	obj.Vouts = make([]*Vout, 0)
	if vouts := gjson.Get(json.Raw, "vout"); vouts.IsArray() {
		for _, vout := range vouts.Array() {
			output := newTxVoutByExplorer(&vout)
			obj.Vouts = append(obj.Vouts, output)
		}
	}

	return &obj
}
*/

//构建NAS专属交易单,包含交易和区块相关信息
func newNasTransaction(transaction *gjson.Result,block *Block) *NasTransaction {

	/*
		//btc
		{
			"txid": "6595e0d9f21800849360837b85a7933aeec344a89f5c54cf5db97b79c803c462",
			"hash": "f758cb5181d51f8bee1512b4a862faad5b51c7c85a1a11cd6092ffc1c6649bc5",
			"version": 2,
			"size": 249,
			"vsize": 168,
			"locktime": 1414190,
			"vin": [],
			"vout": [],
			"hex": "02000000000101cc8a3077023c08040e677647ad0e528564764f456b01d8519828df165ab3c4550100000017160014aa59f94152351c79b57b14a53e538a923e332468feffffff02a716167c6f00000017a914a0fe07f130a36d9c7581ccd2886895c049b0cc8287ece29c00000000001976a9148c0bceb59d452b3e077f73a420b8bfe09e0550a788ac0247304402205e667171c1798cde426282bb8bff45901866ad6bf0d209e856c1765eda65ba4802203aaa319ea3de00eccef0006e6ee2089aed4b91ada7953f420a47c9c258d424ca0121033cfda2f93d13b01d46ecc406b03ebaba3e1bd526d2148a0a5d579d52f8c7cf022e941500",
			"blockhash": "0000000040730ea7935cce346ce68bf4c07c10b137ba31960bf8a47c4f7da4ec",
			"confirmations": 20076,
			"time": 1537841342,
			"blocktime": 1537841342
		}
		//nas
		{
			"hash": "c98c1c117442554a674262b6c566e8ba12b7397fdcbf09abb6f3bd06a1646b83",
			"chainId": 100,
			"from": "n1NM2eETQG5Es7cCc7sh29NJr9cP94QZcXR",
			"to": "n1Qzgp7TE4TNQTC9LQpEsuYZFbGU14cpDZk",
			"value": "17000000000000000000",
			"nonce": "39",
			"timestamp": "1",
			"type": "binary",
			"data": null,
			"gas_price": "1000000",
			"gas_limit": "2000000",
			"contract_address": "",
			"status": 1,
			"gas_used": "20000",
			"execute_error": "",
			"execute_result": ""
       }
	*/
	obj := NasTransaction{}
	//解析json
	obj.Hash = transaction.Get( "hash").String()
	obj.ChainId = transaction.Get( "chainId").Uint()
	obj.From = transaction.Get( "from").String()
	obj.To = transaction.Get( "to").String()
	obj.Value = decimal.RequireFromString(transaction.Get( "value").String())
	obj.Nonce = transaction.Get( "nonce").String()
	obj.Timestamp = transaction.Get( "timestamp").String()
	obj.Type = transaction.Get( "type").String()
	obj.Data = transaction.Get( "data").String()
	obj.Gas_price = transaction.Get( "gas_price").String()
	obj.Gas_limit = transaction.Get( "gas_limit").String()
	obj.Contract_address = transaction.Get( "contract_address").String()
	obj.Status = transaction.Get( "status").Uint()
	obj.Gas_used = transaction.Get( "gas_used").String()
	obj.Execute_error = transaction.Get( "execute_error").String()
	obj.Execute_result = transaction.Get( "execute_result").String()

	//区块相关信息
	obj.BlockTime = block.Time
	obj.BlockHeight = block.Height
	obj.BlockHash = block.Hash

	return &obj
}
