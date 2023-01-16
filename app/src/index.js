import Web3 from "web3";
import metaCoinArtifact from "../../build/contracts/MetaCoin.json";

const App = {
  web3: null,
  account: null,
  meta: null,

  start: async function() {
    const { web3 } = this;

    try {
      // get contract instance
      const networkId = await web3.eth.net.getId();
      const deployedNetwork = metaCoinArtifact.networks[networkId];
      this.meta = new web3.eth.Contract(
        metaCoinArtifact.abi,
          "0xF2edF3b2c4cA91df1dAda13610e47b0a42165b9a",
          //"0x0E3322426c0dc6aeba9FEB6e6F102438b37723a8",
       //deployedNetwork.address,
      );

      // get accounts
      const accounts = await web3.eth.getAccounts();
      this.account = accounts[0];

      //this.refreshBalance();
      //this.refreshEthBalance();
      //this.refreshConBalance();
    } catch (error) {
      console.error("Could not connect to contract or chain.");
    }
  },

  refreshBalance: async function() {
    const { getBalance } = this.meta.methods;
    const balance = await getBalance(this.account).call();

    const balanceElement = document.getElementsByClassName("balance")[0];
    balanceElement.innerHTML = balance;
  },

  refreshConBalance: async function() {
    const { getEthBalanceInContract } = this.meta.methods;
    const balance = await getEthBalanceInContract().call();

    const balanceElement = document.getElementsByClassName("balance3")[0];
    balanceElement.innerHTML = balance;
  },

  refreshEthBalance: async function() {
    const { getEthBalance } = this.meta.methods;
    const balance = await getEthBalance(this.account).call();

    const balanceElement = document.getElementsByClassName("balance2")[0];
    balanceElement.innerHTML = balance;
  },

  sendCoin: async function() {
    const amount = parseInt(document.getElementById("amount").value);
    const receiver = document.getElementById("receiver").value;

    this.setStatus("Initiating transaction... (please wait)");

    const { sendCoin } = this.meta.methods;
    await sendCoin(receiver, amount).send({ from: this.account });

    this.setStatus("Transaction complete!");
    this.refreshBalance();
  },

  sendEth: async function() {
    const amount = parseInt(document.getElementById("amount").value);
    //const receiver = document.getElementById("receiver").value;

    this.setStatus("Initiating deposit... (please wait)");

    const { sendEth } = this.meta.methods;
    await sendEth().send({from: this.account, value:amount});

    this.setStatus("Deposit complete!");
    await this.refreshEthBalance();
    await this.refreshConBalance();
  },

  transferEth: async function() {
    const amount = document.getElementById("amount").value.toString();//parseInt(document.getElementById("amount").value);
    const receiver = document.getElementById("receiver").value;

    this.setStatus("Initiating deposit... (please wait)");

    const { transferEth } = this.meta.methods;
    await transferEth(receiver, amount).send({from: this.account});

    this.setStatus("Deposit complete!");
    await this.refreshEthBalance();
    await this.refreshConBalance();
  },
  setStatus: function(message) {
    const status = document.getElementById("status");
    status.innerHTML = message;
  },
};

window.App = App;
window.playerWalletAddress = null;
//网页加载的时候能够创建登录以太坊钱包地址
window.onload = async (event) => {
  if (window.ethereum) {
    // use MetaMask's provider
    App.web3 = new Web3(window.ethereum);
    //window.ethereum.enable(); // get permission to access accounts
  } else {
    alert("Please install MetaMask or any Ethereum Extension Wallet.");
  }
  localStorage.getItem("userWalletAddress") ? window.playerWalletAddress = localStorage.getItem("userWalletAddress") : window.playerWalletAddress = null;
  if (window.playerWalletAddress!=null) {
    //alert(window.playerWalletAddress);
    document.querySelector(".login-btn").innerHTML =  "🌐 Connected " + window.playerWalletAddress.substring(0,6)+ "…" + window.playerWalletAddress.substring(38);
  }
  //App.start();
};

// Web3 login function
const loginWithEth = async () => {
  // check if there is global window.web3 instance
  if (window.App.web3) {
    try {
      // get the user's ethereum account - prompts metamask to login
      const selectedAccount = await window.ethereum
          .request({
            method: "eth_requestAccounts",
          })
          .then((accounts) => accounts[0])
          .catch(() => {
            // if the user cancels the login prompt
            throw Error("Please select an account");
          });
      // set the global userWalletAddress variable to selected account
      window.playerWalletAddress = Web3.utils.toChecksumAddress(selectedAccount);
      //localStorage.setItem("playerWalletAddress", window.playerWalletAddress);
      console.log(selectedAccount);
      // store the user's wallet address in local storage
      window.localStorage.setItem("userWalletAddress",  window.playerWalletAddress );
      // show the user dashboard
      //showUserDashboard();
      //alert(window.playerWalletAddress);
      //document.querySelector(".login-section").style.display = "none";
      document.querySelector(".login-btn").innerHTML = "🌐 Connected " + window.playerWalletAddress.substring(0,6)+ "…" + window.playerWalletAddress.substring(38);
    } catch (error) {
      alert(error);
    }
  } else {
    alert("wallet not found");
  }
};

// when the user clicks the login button run the loginWithEth function
document.querySelector(".login-btn").addEventListener("click", loginWithEth);
