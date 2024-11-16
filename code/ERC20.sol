// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

abstract contract ERC20 {
    event Transfer(address indexed _from, address indexed _to, uint256 _value);
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);

    /*
        @dev Stores balances of accounts
    */
    mapping(address owner => uint256) private _balances;

    /*
        @dev Stores the amounts of tokens to be allowed to spend from another account
    */
    mapping(address owner => mapping(address spender => uint256)) private _allowances;

    /*
        @dev The number of tokens on the whole blockchain
    */
    uint256 private _totalSupply;

    string private _name;
    string private _symbol;
    address private _authority;

    constructor(string memory name_, string memory symbol_, address authority_) {
        _name = name_;
        _symbol = symbol_;
        _authority = authority_;
    }

    /*
        @dev Returns the name of the token
    */
    function name() public view virtual returns (string memory) {
        return _name;
    }

    /*
        @dev Returns the ticker of the token
    */
    function symbol() public view virtual returns (string memory){
        return _symbol;
    }

    /*
        @dev Decimal helps us to divide token into "cents"
    */
    function decimals() public view virtual returns (uint8) {
        return 18;
    }

    /*
        @dev Return the total supply of tokens
    */
    function totalSupply() public view virtual returns (uint256) {
        return _totalSupply;
    }

    /*
        @dev Returns the balance of the _owner
    */
    function balanceOf(address _owner) public view virtual returns (uint256 balance) {
        return _balances[_owner];
    }

    /*
        @dev Returns how much tokens _spender is allowed to spend on behalf of _owner
    */
    function allowance(address _owner, address _spender) public view virtual returns(uint256 remaining) {
        return _allowances[_owner][_spender];
    }

    /*
        @dev Mints new tokens on blockchain
    */
    function mint(address _account, uint256 _value) external virtual {
        require(msg.sender == _authority, "not owner");
        require(_account != address(0), "invalid account");

        _update(address(0), _account, _value);
    }

    /*
        @dev Transfers tokens of the message sender
    */
    function transfer(address _to, uint256 _value) public virtual returns (bool success) {
        _transfer(msg.sender, _to, _value);
        return true;
    }

    /*
        @dev Transfers tokens on behalf of _from, if message sender is allowed
    */
    function transferFrom(address _from, address _to, uint256 _value) public virtual returns (bool success) {
        _spendAllowance(_from, msg.sender, _value);
        _transfer(_from, _to, _value);
        return true;
    }

    /*
        @dev Approving spending of tokens to _spender on behalf of message sender
    */
    function approve(address _spender, uint256 _value) public virtual returns (bool success) {
        _approve(msg.sender, _spender, _value);
        return true;
    }

    /*
        @dev Used to transfer money from one account to another, if there is enough of funds
    */
    function _transfer(address _from, address _to, uint256 _value) internal virtual {
        require(_from != address(0), "invalid spender");
        require(_to != address(0), "invalid receiver");
        _update(_from, _to, _value);
    }

    /*
        @dev Updates values of accounts.
        Can be used for minting as well as burning of tokens
    */
    function _update(address _from, address _to, uint256 _value) internal virtual {
        if (_from == address(0)) {
            _totalSupply += _value;
        } else {
            require(_balances[_from] >= _value, "sender has insufficient funds");
            unchecked {
                _balances[_from] -= _value;
            }
        }

        if (_to == address(0)) {
            _totalSupply -= _value;
        } else {
            unchecked {
                _balances[_to] += _value;
            }
        }

        emit Transfer(_from, _to, _value);
    }

    /*
        @dev Used to set allowance
    */
    function _approve(address _owner, address _spender, uint256 _value) internal virtual {
        _approve(_owner, _spender, _value, true);
    }

    /*
        @dev Used to update allowance.
        Does not check for availability of such amount.
        The availability is checked via _transfer()
    */
    function _spendAllowance(address _owner, address _spender, uint256 _value) internal virtual {
        uint256 _currentAllowance = allowance(_owner, _spender);

        require(_currentAllowance >= _value, "not enough of allowance");

        unchecked {
            _approve(_owner, _spender, _value, false);
        }
    }

    /*
        @dev Sets allowance and emits Approve event, if needed
    */
    function _approve(address _owner, address _spender, uint256 _value, bool _emitEvent) internal virtual {
        require(_owner!=address(0), "invalid owner");
        require(_spender!=address(0), "invalid spender");

        _allowances[_owner][_spender] = _value;

        if (_emitEvent) {
            emit Approval(_owner, _spender, _value);
        }
    }

    /*
        @dev Mints new tokens on blockchain
    */
    function _mint(address _account, uint256 _value) internal virtual {
        require(_account != address(0), "invalid account");

        _update(address(0), _account, _value);
    }
}

