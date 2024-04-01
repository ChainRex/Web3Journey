// SPDX-License-Identifier: MIT
pragma solidity 0.8.8;

contract SimpleStorage {
    uint256 favoriteNunmber;

    mapping(string => uint256) public nameToFavoriteNumber;

    struct People {
        uint256 favoriteNunmber;
        string name;
    }

    // uint256[] public favoriteNunmberList;
    People[] public people;

    function store(uint256 _favoriteNumber) public virtual {
        favoriteNunmber = _favoriteNumber;
    }

    function retrieve() public view returns (uint256) {
        return favoriteNunmber;
    }

    function addPerson(string memory _name, uint256 _favoriteNunmber) public {
        people.push(People(_favoriteNunmber, _name));
        nameToFavoriteNumber[_name] = _favoriteNunmber;
    }
}
