// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

contract Cat{
    uint private x;
    constructor(){}
    function myFunction() external{
        x = 3;
    }
}

contract Kitten is Cat{
    constructor(){}
}