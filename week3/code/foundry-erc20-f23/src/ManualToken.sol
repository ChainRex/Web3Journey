// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

contract ManualToken {
    mapping(address => uint256) private s_balances;

    string public name = "Manual Token";

    function totalSupply() public pure returns (uint256) {
        return 100 ether;
    }

    function decimals() public pure returns (uint8) {
        return 18;
    }

    function balanceOf(address _owner) public view returns (uint256 balance) {
        return s_balances[_owner];
    }

    function transfer(address _to, uint256 _amount) public {
        uint256 previousBalances = s_balances[msg.sender] + s_balances[_to];
        require(s_balances[msg.sender] >= _amount);
        s_balances[msg.sender] -= _amount;
        s_balances[_to] += _amount;
        assert(s_balances[msg.sender] + s_balances[_to] == previousBalances);
    }
}
