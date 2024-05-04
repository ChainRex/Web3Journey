// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import {TimelockController} from "@openzeppelin/contracts/governance/extensions/GovernorTimelockControl.sol";

contract TImeLock is TimelockController {
    constructor(uint256 minDelay, address[] memory proposers, address[] memory executors)
        TimelockController(minDelay, proposers, executors, msg.sender)
    {}
}
