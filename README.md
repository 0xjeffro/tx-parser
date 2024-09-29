<p align="center">
<img src="https://www.jeffro.io/assets/wiki/tx-parser/whale.png" alt="tx-parser" width="100">
</p>
<h1 align="center">On-chain Transaction Parser</h1>

> A powerful library for parsing on-chain transactions into clear, human-readable actions, streamlining blockchain data analysis and interpretation.

[![Build Status](https://github.com/0xjeffro/tx-parser/workflows/tests/badge.svg)](https://github.com/0xjeffro/tx-parser/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/0xjeffro/tx-parser)](https://goreportcard.com/report/github.com/0xjeffro/tx-parser)
[![codecov](https://codecov.io/github/0xjeffro/tx-parser/graph/badge.svg?token=1VPAKE8N6P)](https://codecov.io/github/0xjeffro/tx-parser)
[![commit](https://img.shields.io/github/last-commit/0xjeffro/tx-parser)](https://github.com/0xjeffro/tx-parser/commits/master)
[![GoDoc](https://pkg.go.dev/badge/github.com/0xjeffro/tx-parser?status.svg)](https://pkg.go.dev/github.com/0xjeffro/tx-parser@v1.0.0?tab=doc)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/0xjeffro/tx-parser)
![GitHub License](https://img.shields.io/github/license/0xjeffro/tx-parser)

## Usage

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/0xjeffro/tx-parser/solana"
	"github.com/0xjeffro/tx-parser/solana/types"
	"io/ioutil"
	"net/http"
	"strings"
)

type RpcRsp struct {
	JsonRpc string      `json:"jsonrpc"`
	Result  types.RawTx `json:"result"`
	Id      int         `json:"id"`
}

func main() { 
	// Fetch a raw transaction from the Solana RPC 
	url := "https://api.mainnet-beta.solana.com/"
	method := "POST"
	tx := "5tuQKcRyFFNQw8XmD2Rg7ZuvXHxJynf9Z3oWFGewU9i5MREpQFhko1d5e6i3z15DqngcRGsXNBtpDvqc5EToAaRd"

	payload := strings.NewReader(fmt.Sprintf(`{
		"jsonrpc": "2.0",
		"id": 1,
		"method": "getTransaction",
		"params": [
		  "%s",
		  {
			"encoding": "json",
			"maxSupportedTransactionVersion": 0
		  }
		]
	}`, tx))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Parse the raw transaction
	var result RpcRsp
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
		return
	}

	var txs []types.RawTx
	txs = append(txs, result.Result)
	b, err := json.Marshal(txs)
	if err != nil {
		return
	}
	parsed, err := solana.Parser(b)
	if err != nil {
		return
	}
	parsedJson, err := json.Marshal(parsed)
	if err != nil {
		return
	}
	fmt.Println(string(parsedJson))
}
```

**Output:**

```json
[
    {
        "actions": [
            {
                "programId": "DCA265Vj8a9CEuX1eb1LWRnDT7uK6q1xMipnNyatn23M",
                "programName": "Jupiter DCA Program",
                "instructionName": "OpenDcaV2",
                "inAmount": 495872257157,
                "inAmountPerCycle": 123968064289,
                "cycleFrequency": 60,
                "minOutAmount": 0,
                "maxOutAmount": 0,
                "startAt": 0,
                "dca": "Dy4fBZPYvBwLoUD3EbhHGMz7DyasB74mzUxGEhH4BCkB",
                "user": "8pKmTuPtNDtCacv2JqyLBfPeZGoQYgqCMJ6Jnubfyuwy",
                "payer": "8pKmTuPtNDtCacv2JqyLBfPeZGoQYgqCMJ6Jnubfyuwy",
                "inputMint": "A8C3xuqscfmyLrte3VmTqrAq8kgMASius9AFNANwpump",
                "outputMint": "3S8qX1MsMqRbiwKg2cQyx7nis1oHMgaCuc9c4VfvVdPN",
                "userAta": "6kzoVLfX3rkUeaYNGt1HB3yd19HhhLBCaWkxn9v6y1xb",
                "inAta": "56n2qpWS2YtTeYTcqjwMWmAst6nqwmWk7yzuLW8Q4Gjk",
                "payerAta": "EXFSAYSHEBN8FRgu48vmyrqYTVpZA5aWLVxSj4i3U7oD"
            },
            {
                "programId": "ComputeBudget111111111111111111111111111111",
                "programName": "ComputeBudget",
                "instructionName": "SetComputeUnitLimit",
                "computeUnitLimit": 300000
            },
            {
                "programId": "ComputeBudget111111111111111111111111111111",
                "programName": "ComputeBudget",
                "instructionName": "SetComputeUnitPrice",
                "microLamports": 25761
            }
        ]
    }
]
```

