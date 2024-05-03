// SPDX-License-Identifier: mit
pragma solidity ^0.8.18;

import {Script} from "forge-std/Script.sol";
import {DevOpsTools} from "lib/foundry-devops/src/DevOpsTools.sol";
import {BoxV2} from "../src/BoxV2.sol";
import {BoxV1} from "../src/BoxV1.sol";

contract UpgradeBox is Script {
    function run() external returns (address) {
        address mostRecentlyDeployed = DevOpsTools.get_most_recent_deployment("ERC1967Proxy", block.chainid);
        vm.startBroadcast();
        BoxV2 box = new BoxV2();
        vm.stopBroadcast();
        address proxy = upgradeBox(mostRecentlyDeployed, address(box));
        return proxy;
    }

    function upgradeBox(address proxy, address newImplementation) public returns (address) {
        vm.startBroadcast();
        BoxV1 box = BoxV1(proxy);
        box.upgradeToAndCall(address(newImplementation), "");
        vm.stopBroadcast();
        return address(box);
    }
}
