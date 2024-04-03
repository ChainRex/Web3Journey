import { ethers } from "./ethers-5.2.esm.min.js"
import { abi, contractAddress } from "./constants.js"

const connectButton = document.getElementById("connectButton")
const fundButton = document.getElementById("fundButton")
const balanceButton = document.getElementById("balaceButton")
const withdrawButton = document.getElementById("withdrawButton")
const msg = document.getElementById("msg")

connectButton.onclick = connect
fundButton.onclick = fund
balanceButton.onclick = getBalance
withdrawButton.onclick = withdraw

async function connect() {
    if (typeof window.ethereum != undefined) {
        console.log("已安装metamask")
        await window.ethereum.request({ method: "eth_requestAccounts" })
        console.log("连接成功")
        connectButton.innerHTML = "已连接"
    } else {
        connectButton.innerHTML = "请安装MetaMask"
    }
}

async function fund() {
    if (typeof window.ethereum != undefined) {
        const ethAmount = document.getElementById("ethAmount").value
        const provider = new ethers.providers.Web3Provider(window.ethereum)
        const signer = provider.getSigner()
        const contract = new ethers.Contract(contractAddress, abi, signer)
        try {
            const transactionResponse = await contract.fund({
                value: ethers.utils.parseEther(ethAmount),
            })
            msg.innerText = "捐赠中，请等待..."
            await listenForTransactionMine(transactionResponse, provider)
            msg.innerText = "捐赠成功！"
        } catch (error) {
            console.log(error)
        }
    }
}

async function getBalance() {
    if (typeof window.ethereum != undefined) {
        const provider = new ethers.providers.Web3Provider(window.ethereum)
        const balance = await provider.getBalance(contractAddress)
        document.getElementById("balance").innerText = ethers.utils.formatEther(balance)
    }
}

async function withdraw() {
    if (typeof window.ethereum != undefined) {
        const provider = new ethers.providers.Web3Provider(window.ethereum)
        const signer = provider.getSigner()
        const contract = new ethers.Contract(contractAddress, abi, signer)
        try {
            const transactionResponse = await contract.withdraw()
            msg.innerText = "提款中，请等待..."
            await listenForTransactionMine(transactionResponse, provider)
            msg.innerText = "提款成功！"
        } catch (error) {
            console.log(error)
        }
    }
}

function listenForTransactionMine(transactionResponse, provider) {
    console.log(`Mining${transactionResponse.hash}...`)
    return new Promise((resolve, reject) => {
        provider.once(transactionResponse.hash, (transactionReceipt) => {
            console.log(`完成${transactionReceipt.confirmations}笔交易`)
            resolve()
        })
    })
}
