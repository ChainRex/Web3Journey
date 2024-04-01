const ethers = require("ethers")
const fs = require("fs-extra")
require("dotenv").config()

async function main() {
    // 127.0.0.1:8545
    const provider = new ethers.providers.JsonRpcProvider(process.env.RPC_URL)
    const wallet = new ethers.Wallet(process.env.PRIVATE_KEY, provider)
    const abi = fs.readFileSync(
        "./SimpleStorage_sol_SimpleStorage.abi",
        "utf-8",
    )
    const binay = fs.readFileSync(
        "./SimpleStorage_sol_SimpleStorage.bin",
        "utf-8",
    )
    const contractFactory = new ethers.ContractFactory(abi, binay, wallet)
    console.log("部署中")
    const contract = await contractFactory.deploy()
    await contract.deployTransaction.wait(1)
    console.log(`合约地址：${contract.address}`)
    console.log("部署完成")

    const currentFavoriteNumber = await contract.retrieve()
    console.log(`FavoriteNumber: ${currentFavoriteNumber.toString()}`)
    const transactionResponse = await contract.store("7")
    const transactionReceipt = await transactionResponse.wait(1)
    const updatedFavoriteNumber = await contract.retrieve()
    console.log(`FavoriteNumber: ${updatedFavoriteNumber.toString()}`)
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error)
        process.exit(1)
    })
