## 降低gas的方法

### constant

函数外分配一次变量（直接赋值），然后永远不再不改变它，可以改为常量，使其不再占用一个存储空间，而是直接写入字节码。可以节省赋值和读取的gas费用。

全部字母大写命名

```solidity
uint256 public constant MINIMUM_USD = 50 * 1e18;
```

### immutable

不可变变量，只能被赋值一次，可以节省赋值和读取的gas费用。

以i_xxx命名变量

```solidity
address public immutable i_owner;
```

### Custom error

避免报错信息字符串单独存储

```solidity
// require(msg.sender == i_owner,"Sender is not owner");
if(msg.sender != i_owner){revert NotOwner();}
```

