// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

import {ERC20} from "./ERC20.sol";

contract G42 is ERC20 {
    constructor(address _authority) ERC20("Grigorian42", "G42", _authority) {
        _mint(msg.sender, 1000 * 10 ** decimals());
    }
}
