// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "IERC.sol";

contract WTC is IERC20{
    // /**
    //  * @dev 释放条件：当 `value` 单位的货币从账户 (`from`) 转账到另一账户 (`to`)时.
    //  */
    // event Transfer(address indexed from, address indexed to, uint256 value);

    // /**
    //  * @dev 释放条件：当 `value` 单位的货币从账户 (`owner`) 授权给另一账户 (`spender`)时.
    //  */
    // event Approval(address indexed owner, address indexed spender, uint256 value);

    error NotOwner();
    error InsufficientBalance();
    error InsufficientAllowance();

    uint public override totalSupply;

    mapping(address => uint) public override balanceOf;

    mapping(address => mapping(address => uint)) public override allowance;

    string name;

    string symbol;

    address owner;


    constructor(string memory _name,string memory _symbol){
        name = _name;
        symbol = _symbol;
        owner = msg.sender;
    }

    modifier onlyOwner{
        if(msg.sender != owner) revert NotOwner();
        _;
    }

    /**
     * @dev 转账 `amount` 单位代币，从调用者账户到另一账户 `to`.
     *
     * 如果成功，返回 `true`.
     *
     * 释放 {Transfer} 事件.
     */
    function transfer(address _to, uint256 _amount) external override returns (bool){
        if(balanceOf[msg.sender] < _amount) revert InsufficientBalance();
        balanceOf[msg.sender] -= _amount;
        balanceOf[_to] += _amount;
        emit Transfer(msg.sender, _to, _amount);
        return true;
    }


    /**
     * @dev 调用者账户给`spender`账户授权 `amount`数量代币。
     *
     * 如果成功，返回 `true`.
     *
     * 释放 {Approval} 事件.
     */
    function approve(address _spender, uint256 _amount) external override returns (bool){
        allowance[msg.sender][_spender] = _amount;
        emit Approval(msg.sender, _spender, _amount);
        return true;
    }

    /**
     * @dev 通过授权机制，从`from`账户向`to`账户转账`amount`数量代币。转账的部分会从调用者的`allowance`中扣除。
     *
     * 如果成功，返回 `true`.
     *
     * 释放 {Transfer} 事件.
     */
    function transferFrom(
        address _from,
        address _to,
        uint256 _amount
    ) external override returns (bool){
        if(allowance[_from][msg.sender] < _amount) revert InsufficientAllowance();
        allowance[_from][msg.sender] -= _amount;
        balanceOf[_from] -= _amount;
        balanceOf[_to] += _amount;
        emit Transfer(_from, _to, _amount);
        return true; 
    }

    function mint(uint _amount) external onlyOwner{
        balanceOf[msg.sender] += _amount;
        totalSupply += _amount;
        emit Transfer(address(0), msg.sender, _amount);
    }
}