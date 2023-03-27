// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "solmate/tokens/ERC721.sol";
import "solmate/auth/Owned.sol";

contract VoteSBT is Owned, ERC721 {
    event Locked(uint256 tokenId);
    event Unlocked(uint256 tokenId);

    error TokenIsSoulbound();

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
        emit Locked(newItemId);
        return newItemId;
    }

    function onlySoulbound(address from, address to) internal pure {
        if (from != address(0) && to != address(0)) {
            revert TokenIsSoulbound();
        }
    }

    function transferFrom(address from, address to, uint256 id) public override {
        onlySoulbound(from, to);
        super.transferFrom(from, to, id);
    }

    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
        return
            interfaceId == 0xb45a3c0e || // ERC165 Interface ID for ERC5192
            super.supportsInterface(interfaceId);
    }

    function tokenURI(uint256 /*tokenId*/) public view virtual override returns (string memory) {
        return baseURI;
    }

    function locked(uint256 /*tokenId*/) external pure returns (bool) {
        return true;
    }
}
