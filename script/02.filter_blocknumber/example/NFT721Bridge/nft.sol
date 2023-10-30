// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

contract NFT721Bridge {


    event TransferNFT(uint64 indexed nonce, address token, uint256 tokenID, uint16 dstChainId, address sender, address recipient);

    event ReceiveNFT(uint64 indexed nonce, address sourceToken, address token, uint256 tokenID, uint16 sourceChain, uint16 sendChain, address recipient);

}