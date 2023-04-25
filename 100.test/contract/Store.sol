// SPDX-License-Identifier: MIT
/*
这里只是一个简单的合约，就是一个键/值存储，只有一个外部方法来设置任何人的键/值对。 我们还在设置值后添加了要发出的事件。
*/
pragma solidity =0.7.6;

contract Store {
    event ItemSet(string key, string value);

    string public version;

    mapping(string => string) public items;

    constructor(string memory _version)  {
        version = _version;
    }

    function setItem(string memory key, string memory value) external {
        items[key] = value;
        emit ItemSet(key, value);
        if (bytes(key).length == 5){
            revert();
        }
    }
}
