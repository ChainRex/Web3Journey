const { ethers, run, network } = require("hardhat")

async function main() {
    const simpleStorageFactory =
        await ethers.getContractFactory("SimpleStorage")
    console.log("部署中")
    const simpleStorage = await simpleStorageFactory.deploy()
    await simpleStorage.deployed()
    console.log(`部署到${simpleStorage.address}`)
    if (network.config.chainId === 11155111 && process.env.ETHERSCAN_API_KEY) {
        await simpleStorage.deployTransaction.wait(3)
        await verify(simpleStorage.address, [])
    }

    const currentValue = await simpleStorage.retrieve()

    console.log(`现在的值是${currentValue}`)

    const transactionResponse = await simpleStorage.store(7)
    await transactionResponse.wait(1)
    const updatedValue = await simpleStorage.retrieve()
    console.log(`现在的值是${updatedValue}`)
}

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

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error)
        process.exit(1)
    })
