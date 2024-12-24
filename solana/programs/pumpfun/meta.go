package pumpfun

import "github.com/0xjeffro/tx-parser/utils"

const Program = "6EF8rrecthR5Dkzon8Nwu78hRvfCKubJ14M5uBEwF6P"
const ProgramName = "PumpFun"

var BuyDiscriminator = utils.CalculateDiscriminator("global:buy")
var SellDiscriminator = utils.CalculateDiscriminator("global:sell")
var CreateDiscriminator = utils.CalculateDiscriminator("global:create")

var AnchorSelfCPILogDiscriminator = [8]uint8{228, 69, 165, 46, 81, 203, 154, 29}
var AnchorSelfCPILogSwapDiscriminator = [8]uint8{189, 219, 127, 211, 78, 230, 97, 238}
