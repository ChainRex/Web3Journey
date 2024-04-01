const { task } = require("hardhat/config")

task("block-number", "打印区块号").setAction(async (taskArgs, hre) => {
    const blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log(`现在的区块号是：${blockNumber}`)
})

module.exports = {}
