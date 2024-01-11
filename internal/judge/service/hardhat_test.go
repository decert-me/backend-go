package service

import (
	"fmt"
	"github.com/tidwall/gjson"
	"regexp"
	"testing"
)

func TestGetContractName(t *testing.T) {
	code := `
		// SPDX-License-Identifier: MIT
		pragma solidity ^0.8.20;
		
		import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
		
		contract GLDToken {
			constructor(uint256 initialSupply) ERC20("Gold", "GLD") {
				_mint(msg.sender, initialSupply);
			}
		}
	`

	re := regexp.MustCompile(`contract\s+(\w+)(\s+is){`)
	match := re.FindStringSubmatch(code)
	if len(match) > 1 {
		fmt.Println(match[1]) // 输出：test
	}
}

func TestCheckResult(t *testing.T) {
	result := `
		[{"testName()":{"status":"Success","reason":null,"counterexample":null,"logs":[{"address":"0x5615deb798bb3e4dfa0139dfa1b3d433cc23b72f","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x0000000000000000000000000000000000000000000000000000000000000000","0x0000000000000000000000007fa9385be102ac3eac297483dd6233d62b3e1496"],"data":"0x00000000000000000000000000000000000000000000000000000000000f4240"}],"decoded_logs":[],"kind":{"Standard":9584},"traces":[],"labeled_addresses":{},"breakpoints":{}},"testSymbol()":{"status":"Success","reason":null,"counterexample":null,"logs":[{"address":"0x5615deb798bb3e4dfa0139dfa1b3d433cc23b72f","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x0000000000000000000000000000000000000000000000000000000000000000","0x0000000000000000000000007fa9385be102ac3eac297483dd6233d62b3e1496"],"data":"0x00000000000000000000000000000000000000000000000000000000000f4240"}],"decoded_logs":[],"kind":{"Standard":9605},"traces":[],"labeled_addresses":{},"breakpoints":{}},"testTotalSupply()":{"status":"Failure","reason":null,"counterexample":null,"logs":[{"address":"0x5615deb798bb3e4dfa0139dfa1b3d433cc23b72f","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x0000000000000000000000000000000000000000000000000000000000000000","0x0000000000000000000000007fa9385be102ac3eac297483dd6233d62b3e1496"],"data":"0x00000000000000000000000000000000000000000000000000000000000f4240"},{"address":"0x7fa9385be102ac3eac297483dd6233d62b3e1496","topics":["0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50"],"data":"0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000224572726f723a2061203d3d2062206e6f7420736174697366696564205b75696e745d000000000000000000000000000000000000000000000000000000000000"},{"address":"0x7fa9385be102ac3eac297483dd6233d62b3e1496","topics":["0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8"],"data":"0x000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000f4240000000000000000000000000000000000000000000000000000000000000000a2020202020204c65667400000000000000000000000000000000000000000000"},{"address":"0x7fa9385be102ac3eac297483dd6233d62b3e1496","topics":["0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8"],"data":"0x00000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000989681000000000000000000000000000000000000000000000000000000000000000a2020202020526967687400000000000000000000000000000000000000000000"}],"decoded_logs":["Error: a == b not satisfied [uint]","      Left: 1000000","     Right: 10000001"],"kind":{"Standard":22189},"traces":[],"labeled_addresses":{},"breakpoints":{}}}]
	`
	//arr := gjson.Get(result, "@this").Array()
	//fmt.Println(len(arr))
	//fmt.Print(gjson.Get(result, "*.status").String())
	gjson.Parse(result).ForEach(func(key, value gjson.Result) bool {
		value.ForEach(func(key, value gjson.Result) bool {
			fmt.Println(key, value)
			status := gjson.Get(value.String(), "status").String()
			fmt.Println("status", status)
			return true // keep iterating
		})
		return true // keep iterating
	})
}

//result := `{"success":true,"message":"Success","data":{"output":"{\"status\":true,\"gasUsed\":\"0x5208\",\"gasCost\":\"0x5208\",\"output\":\"0x\",\"address\":\"0x
