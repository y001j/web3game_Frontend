// SPDX-License-Identifier: MIT
pragma solidity >=0.4.21 <0.9.0;

import "./ConvertLib.sol";

// This is a simple example of a coin-like contract.
// It is not standards compatible and cannot be expected to talk to other
// coin/token contracts. If you want to create a standards-compliant
// token, see: https://github.com/ConsenSys/Tokens. Cheers!

contract MetaCoin {
    address public owner;
    mapping (address => uint) balances;
    mapping (address => uint) deposit;

    event TransEvent(address,uint);


    event Transfer(address indexed _from, address indexed _to, uint256 _value);

    constructor()  {
        balances[msg.sender] = 10000;
        owner = msg.sender;
    }

    function sendCoin(address receiver, uint amount) public returns(bool sufficient) {
        if (balances[msg.sender] < amount) return false;
        balances[msg.sender] -= amount;
        balances[receiver] += amount;
        emit Transfer(msg.sender, receiver, amount);
        return true;
    }

    function getBalanceInEth(address addr) public view returns(uint) {
        return ConvertLib.convert(getBalance(addr),2);
    }

    function getBalance(address addr) public view returns(uint) {
        return balances[addr];
    }


    fallback() external payable {
        //emit TransEvent(address(this),1);
    }
    receive() external payable {

        //emit TransEvent(msg.sender,2);
    }

    function getEthBalanceInContract() public view returns(uint){
        return address(this).balance;
    }

    function getEthBalance(address addr) public view returns(uint){
        return deposit[addr];
    }

    function sendEth() public payable{
        //if(msg.value <= msg.sender.balance){
        deposit[msg.sender] += msg.value;
        emit TransEvent(msg.sender,2);
        //}
    }

    function transferEth(address payable recipient, uint amount) public returns(bool){
        require(owner == msg.sender || recipient == msg.sender, "Transfer failed! You are not eligible to claim reward!");
        if(amount <= getEthBalance(recipient)){
            recipient.transfer(amount);
            deposit[recipient] -= amount;
            return true;
        }else{
            return false;
        }

    }
}