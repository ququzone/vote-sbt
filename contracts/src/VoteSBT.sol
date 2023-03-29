// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "solmate/tokens/ERC1155.sol";
import "solmate/auth/Owned.sol";

contract VoteSBT is Owned, ERC1155 {
    event Locked(uint256 tokenId);
    event Unlocked(uint256 tokenId);

    error TokenNotExist(uint256 tokenId);
    error TokenIsSoulbound();

    mapping (address => bool) public minters;
    mapping (uint256 => bool) private tokens;
    string private baseUri;
    uint256 private currentTokenId;

    modifier onlyMinter() virtual {
        require(minters[msg.sender], "ONLY_MINTER");
        _;
    }

    constructor(string memory _baseUri) Owned(msg.sender) {
        baseUri = _baseUri;
    }

    function newToken() external onlyOwner returns (uint256) {
        uint256 newItemId = ++currentTokenId;
        tokens[newItemId] = true;
        emit Locked(newItemId);
        return newItemId;
    }

    function addMinter(address minter) external onlyOwner {
        minters[minter] = true;
    }

    function removeMinter(address minter) external onlyOwner {
        minters[minter] = false;
    }

    function mint(address recipient, uint256 tokenId) public payable onlyMinter {
        if(tokens[tokenId] != true) {
            revert TokenNotExist(tokenId);
        }
        _mint(recipient, tokenId, 1, "");
    }

    function onlySoulbound(address from, address to) internal pure {
        if (from != address(0) && to != address(0)) {
            revert TokenIsSoulbound();
        }
    }

    function safeTransferFrom(
        address from,
        address to,
        uint256 id,
        uint256 amount,
        bytes calldata data
    ) public virtual override {
        onlySoulbound(from, to);
        super.safeTransferFrom(from, to, id, amount, data);
    }

    function safeBatchTransferFrom(
        address from,
        address to,
        uint256[] calldata ids,
        uint256[] calldata amounts,
        bytes calldata data
    ) public virtual override {
        onlySoulbound(from, to);
        super.safeBatchTransferFrom(from, to, ids, amounts, data);
    }

    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
        return
            interfaceId == 0xb45a3c0e || // ERC165 Interface ID for ERC5192
            super.supportsInterface(interfaceId);
    }

    function uri(uint256 tokenId) public view virtual override returns (string memory) {
        return string(abi.encodePacked(baseUri, toString(tokenId), ".json"));
    }

    function locked(uint256 /*tokenId*/) external pure returns (bool) {
        return true;
    }

    function toString(uint256 value) internal pure returns (string memory) {
        // Inspired by OraclizeAPI's implementation - MIT licence
        // https://github.com/oraclize/ethereum-api/blob/b42146b063c7d6ee1358846c198246239e9360e8/oraclizeAPI_0.4.25.sol

        if (value == 0) {
            return "0";
        }
        uint256 temp = value;
        uint256 digits;
        while (temp != 0) {
            digits++;
            temp /= 10;
        }
        bytes memory buffer = new bytes(digits);
        while (value != 0) {
            digits -= 1;
            buffer[digits] = bytes1(uint8(48 + uint256(value % 10)));
            value /= 10;
        }
        return string(buffer);
    }
}
