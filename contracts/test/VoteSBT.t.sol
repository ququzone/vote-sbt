// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import {VoteSBT} from "../src/VoteSBT.sol";
import {ERC721TokenReceiver} from "solmate/tokens/ERC721.sol";

contract VoteSBTTest is Test, ERC721TokenReceiver {
    VoteSBT vote;
    uint256 internal tokenId;
    address internal minter;
    address internal alice;
    address internal bob;

    function setUp() public {
        vote = new VoteSBT("https://nftassets.com/");
        tokenId = vote.newToken();
        alice = vm.addr(0x1);
        bob = vm.addr(0x2);
        minter = vm.addr(0x3);

        vote.addMinter(minter);
    }

    function onERC721Received(
        address, address, uint256, bytes calldata
    ) public pure override returns (bytes4) {
        return ERC721TokenReceiver.onERC721Received.selector;
    }

    function testMintTokens() public {
        // Check old balance
        uint256 existingBalance = vote.balanceOf(alice, tokenId);
        vm.expectRevert("ONLY_MINTER");
        vote.mint{value:0}(alice, tokenId);
        vm.prank(minter);
        vote.mint{value:0}(alice, tokenId);
        uint256 newBalance = vote.balanceOf(alice, tokenId);

        assertEq(vote.uri(tokenId), "https://nftassets.com/1.json");

        // Check that our token balance increased by 1
        if(newBalance - existingBalance != 1){
            revert("No tokens were sent to the test contract.");
        }
    }

    function testCannotTransferTokens() public {
        vm.prank(minter);
        vote.mint{value:0}(alice, tokenId);

        vm.startPrank(alice);
        vm.expectRevert(abi.encodeWithSignature("TokenIsSoulbound()"));
        vote.safeTransferFrom(alice, bob, tokenId, 1, "");

        uint256[] memory tokenIds = new uint256[](1);
        tokenIds[0] = tokenId;
        uint256[] memory amounts = new uint256[](1);
        amounts[0] = 1;
        
        vm.expectRevert(abi.encodeWithSignature("TokenIsSoulbound()"));
        vote.safeBatchTransferFrom(alice, bob, tokenIds, amounts, "");

        vm.stopPrank();
    }

    function testInterface() public {
        assertEq(vote.supportsInterface(0x01ffc9a7), true); // ERC165
        assertEq(vote.supportsInterface(0xd9b67a26), true); // ERC1155
        assertEq(vote.supportsInterface(0x0e89341c), true); // ERC1155MetadataURI
        assertEq(vote.supportsInterface(0xb45a3c0e), true); // ERC5192
    }
}
