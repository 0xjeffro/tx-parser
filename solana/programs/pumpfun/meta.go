package pumpfun

import "github.com/0xjeffro/tx-parser/utils"

const Program = "6EF8rrecthR5Dkzon8Nwu78hRvfCKubJ14M5uBEwF6P"

var BuyDiscriminator = utils.CalculateDiscriminator("global:buy")
var SellDiscriminator = utils.CalculateDiscriminator("global:sell")
