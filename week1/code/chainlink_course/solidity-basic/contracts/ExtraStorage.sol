// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "contracts/SimpleStorage.sol";

contract ExtraStorage is SimpleStorage{

    function store(uint256 _favoriteNumber) public override {
        favoriteNunmber = _favoriteNumber + 5;
    }

}