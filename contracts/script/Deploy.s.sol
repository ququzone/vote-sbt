// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import "../src/VoteSBT.sol";

contract Deploy is Script {
    function run() external {
        vm.startBroadcast();
        new VoteSBT("https://nftassets.com/");
        vm.stopBroadcast();
    }
}
