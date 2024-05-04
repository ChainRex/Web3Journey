// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test} from "forge-std/Test.sol";
import {MyGovernor, GovernorCountingSimple} from "../src/MyGovernor.sol";
import {Box} from "../src/Box.sol";
import {GovToken} from "../src/GovToken.sol";
import {TImeLock} from "../src/TImeLock.sol";
import {IGovernor} from "@openzeppelin/contracts/governance/IGovernor.sol";

contract MyGovernorTest is Test {
    MyGovernor governor;
    Box box;
    TImeLock timelock;
    GovToken govToken;

    address public USER = makeAddr("user");
    uint256 public constant INITIAL_SUPPLY = 100 ether;
    uint256 public constant MIN_DELAY = 3600;
    uint256 public constant VOTING_DELAY = 1;
    uint256 public constant VOTING_PERIOD = 50400;

    function setUp() public {
        govToken = new GovToken();
        govToken.mint(USER, INITIAL_SUPPLY);
        vm.startPrank(USER);
        govToken.delegate(USER);
        timelock = new TImeLock(MIN_DELAY, new address[](0), new address[](0)); // 所有人都可以提案和执行
        governor = new MyGovernor(govToken, timelock);

        bytes32 propoerserRole = timelock.PROPOSER_ROLE();
        bytes32 executorRole = timelock.EXECUTOR_ROLE();
        bytes32 adminRole = timelock.DEFAULT_ADMIN_ROLE();

        timelock.grantRole(propoerserRole, address(governor));
        timelock.grantRole(executorRole, address(0));
        timelock.revokeRole(adminRole, USER);
        vm.stopPrank();

        box = new Box();
        box.transferOwnership(address(timelock)); // 注意是timelock，不是governor
    }

    function testCantUpdateBoxWithoutGovernance() public {
        vm.expectRevert();
        box.store(1);
    }

    function testGovernanceUpdatesBox() public {
        vm.startPrank(USER); // 手上token数量越多，权重越大
        uint256 valueToStore = 888;
        string memory description = "store 888";
        bytes[] memory datas = new bytes[](1);
        bytes memory data = abi.encodeWithSignature("store(uint256)", valueToStore);
        datas[0] = data;
        uint256[] memory value = new uint256[](1);
        value[0] = 0;
        address[] memory targets = new address[](1);
        targets[0] = address(box);

        // 向Dao提出提案
        uint256 id = governor.propose(targets, value, datas, description);
        assertEq(uint256(governor.state(id)), uint256(IGovernor.ProposalState.Pending));

        vm.warp(block.timestamp + VOTING_DELAY + 1);
        vm.roll(block.number + VOTING_DELAY + 1);
        assertEq(uint256(governor.state(id)), uint256(IGovernor.ProposalState.Active));

        // 投票
        string memory reason = "vote yes";
        uint8 voteType = uint8(GovernorCountingSimple.VoteType.For);
        governor.castVoteWithReason(id, voteType, reason);

        vm.warp(block.timestamp + VOTING_PERIOD + 1);
        vm.roll(block.number + VOTING_PERIOD + 1);
        assertEq(uint256(governor.state(id)), uint256(IGovernor.ProposalState.Succeeded));

        // 交易排队
        bytes32 descriptionHash = keccak256(abi.encodePacked(description));
        governor.queue(targets, value, datas, descriptionHash);
        assertEq(uint256(governor.state(id)), uint256(IGovernor.ProposalState.Queued));
        vm.warp(block.timestamp + MIN_DELAY + 1);
        vm.roll(block.number + MIN_DELAY + 1);

        // 执行
        governor.execute(targets, value, datas, descriptionHash);
        assertEq(uint256(governor.state(id)), uint256(IGovernor.ProposalState.Executed));
        assertEq(box.getNumber(), valueToStore);

        vm.stopPrank();
    }
}
