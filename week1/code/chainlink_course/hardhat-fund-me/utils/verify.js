const { run } = require("hardhat")
async function verify(contractAddress, args) {
    console.log("验证合约中")
    try {
        await run("verify:verify", {
            address: contractAddress,
            constructorArguments: args,
        })
    } catch (e) {
        if (e.message.toLowerCase().includes("already verified")) {
            console.log("合约已经被验证过了")
        } else {
            console.log(e)
        }
    }
}

module.exports = { verify }
