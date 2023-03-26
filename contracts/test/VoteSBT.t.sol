// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import {VoteSBT} from "../src/VoteSBT.sol";
import {ERC721TokenReceiver} from "solmate/tokens/ERC721.sol";

contract VoteSBTTest is Test, ERC721TokenReceiver {
    VoteSBT vote;
    address internal alice;
    address internal bob;

    function setUp() public {
        vote = new VoteSBT("I Vote", "IVOTE", "");
        alice = vm.addr(0x2);
        bob = vm.addr(0x3);
    }

    function onERC721Received(address, address, uint256, bytes calldata) public pure override returns (bytes4){
        return ERC721TokenReceiver.onERC721Received.selector;
    }

    function testMintTokens() public{
        // Check old balance
        uint256 existingBalance = vote.balanceOf(alice);
        vote.mint{value:0}(alice);
        uint256 newBalance = vote.balanceOf(alice);

        // Check that our token balance increased by 1
        if(newBalance - existingBalance != 1){
            revert("No tokens were sent to the test contract.");
        }
    }

    function testCannotTransferTokens() public{
        vote.mint{value:0}(alice);

        vm.startPrank(alice);
        uint256 tokenId = vote.currentTokenId() - 1;

        vm.expectRevert(abi.encodeWithSignature("TokenIsSoulbound()"));
        vote.safeTransferFrom(alice, bob, tokenId);

        vm.stopPrank();
    }
}
