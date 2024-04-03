const { getNamedAccounts, ethers } = require("hardhat")

async function main() {
    const { deployer } = await getNamedAccounts()
    const fundMe = await ethers.getContract("FundMe", deployer)

    console.log("捐赠中")

    const transactionResponse = await fundMe.fund({ value: ethers.utils.parseEther("1") })

    await transactionResponse.wait(1)
    console.log("捐赠成功")
}
main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error)
        process.exit(1)
    })
