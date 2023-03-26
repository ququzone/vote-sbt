// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "solmate/tokens/ERC721.sol";
import "solmate/auth/Owned.sol";

contract VoteSBT is Owned, ERC721 {
    string private baseURI;
    uint256 public currentTokenId;

    constructor(
        string memory _name,
        string memory _symbol,
        string memory _baseURI
    ) Owned(msg.sender) ERC721(_name, _symbol) {
        baseURI = _baseURI;
    }

    function mint(address recipient) public payable onlyOwner returns (uint256) {
        uint256 newItemId = ++currentTokenId;
        _safeMint(recipient, newItemId);
        return newItemId;
    }

    function tokenURI(uint256 /*id*/) public view virtual override returns (string memory) {
        return baseURI;
    }
}
