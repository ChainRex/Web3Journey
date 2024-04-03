const { assert, expect } = require("chai")
const { deployments, ethers, getNamedAccounts, network } = require("hardhat")
const { developmentChain } = require("../../helper-hardhat-config")

developmentChain.includes(network.name)
    ? describe.skip
    : describe("FundMe", async function () {
          let fundMe
          let deployer
          const sendValue = ethers.utils.parseEther("0.025")
          console.log("集成测试开始")

          beforeEach(async function () {
              deployer = (await getNamedAccounts()).deployer
              console.log(deployer)
              fundMe = await ethers.getContract("FundMe", deployer)
              console.log(fundMe.address)
          })
          // describe("constructor", async function () {
          //     it("喂价地址设置成功", async function () {
          //         const response = await fundMe.getPriceFeed()
          //         console.log(response)
          //         // assert.equal(response, MockV3Aggregator.address)
          //     })
          // })

          describe("fund", async function () {
              it("如果ETH不足则失败", async function () {
                  console.log("如果ETH不足则失败")
                  await expect(fundMe.fund()).to.be.reverted
              })
              it("更新募捐数据成功", async function () {
                  const transactionResponse = await fundMe.fund({ value: sendValue })
                  await transactionResponse.wait(1)
                  const response = await fundMe.getAddressToAmountFunded(deployer)
                  assert.equal(response.toString(), sendValue.toString())
              })
              it("把funder添加到funders", async function () {
                  const response = await fundMe.getFunder(0)
                  assert.equal(response, deployer)
              })
          })
          describe("withdraw", async function () {
              beforeEach(async function () {
                  const transactionResponse = await fundMe.fund({ value: sendValue })
                  await transactionResponse.wait(1)
              })
              it("当只有一个funder时提取ETH", async function () {
                  // 获取fundMe合约的初始余额
                  const startingFundMeBalance = await fundMe.provider.getBalance(fundMe.address)

                  // 获取deployer的初始余额
                  const startingDeployerBalance = await fundMe.provider.getBalance(deployer)

                  const transactionResponse = await fundMe.withdraw()
                  const transactionReceipt = await transactionResponse.wait(1)
                  const { gasUsed, effectiveGasPrice } = transactionReceipt
                  const gasCost = gasUsed.mul(effectiveGasPrice)

                  const endingFundMeBalance = await fundMe.provider.getBalance(fundMe.address)
                  const endingDeployerBalance = await fundMe.provider.getBalance(deployer)

                  assert.equal(endingFundMeBalance, 0)
                  assert.equal(
                      startingFundMeBalance.add(startingDeployerBalance).toString(),
                      endingDeployerBalance.add(gasCost).toString()
                  )
              })

              it("更便宜的提款_单人", async function () {
                  // 获取fundMe合约的初始余额
                  const startingFundMeBalance = await fundMe.provider.getBalance(fundMe.address)

                  // 获取deployer的初始余额
                  const startingDeployerBalance = await fundMe.provider.getBalance(deployer)

                  const transactionResponse = await fundMe.cheaperWithdraw()
                  const transactionReceipt = await transactionResponse.wait(1)
                  const { gasUsed, effectiveGasPrice } = transactionReceipt
                  const gasCost = gasUsed.mul(effectiveGasPrice)

                  const endingFundMeBalance = await fundMe.provider.getBalance(fundMe.address)
                  const endingDeployerBalance = await fundMe.provider.getBalance(deployer)

                  assert.equal(endingFundMeBalance, 0)
                  assert.equal(
                      startingFundMeBalance.add(startingDeployerBalance).toString(),
                      endingDeployerBalance.add(gasCost).toString()
                  )
              })
          })
      })
