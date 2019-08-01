pragma solidity ^0.4.24;

/**
 * Utility library of inline functions on addresses
 */
library AddressUtils {

  /**
   * Returns whether the target address is a contract
   * @dev This function will return false if invoked during the constructor of a contract,
   * as the code is not actually created until after the constructor finishes.
   * @param _account address of the account to check
   * @return whether the target address is a contract
   */
  function isContract(address _account) internal view returns (bool) {
    uint256 size;
    // XXX Currently there is no better way to check if there is a contract in an address
    // than to check the size of the code at that address.
    // See https://ethereum.stackexchange.com/a/14016/36603
    // for more details about how this works.
    // TODO Check this again before the Serenity release, because all addresses will be
    // contracts then.
    // solium-disable-next-line security/no-inline-assembly
    assembly { size := extcodesize(_account) }
    return size > 0;
  }

}

/**
 * @title SafeMath
 * @dev Math operations with safety checks that revert on error
 */
library SafeMath {

  /**
  * @dev Multiplies two numbers, reverts on overflow.
  */
  function mul(uint256 _a, uint256 _b) internal pure returns (uint256) {
    // Gas optimization: this is cheaper than requiring 'a' not being zero, but the
    // benefit is lost if 'b' is also tested.
    // See: https://github.com/OpenZeppelin/openzeppelin-solidity/pull/522
    if (_a == 0) {
      return 0;
    }

    uint256 c = _a * _b;
    require(c / _a == _b);

    return c;
  }

  /**
  * @dev Integer division of two numbers truncating the quotient, reverts on division by zero.
  */
  function div(uint256 _a, uint256 _b) internal pure returns (uint256) {
    require(_b > 0); // Solidity only automatically asserts when dividing by 0
    uint256 c = _a / _b;
    // assert(_a == _b * c + _a % _b); // There is no case in which this doesn't hold

    return c;
  }

  /**
  * @dev Subtracts two numbers, reverts on overflow (i.e. if subtrahend is greater than minuend).
  */
  function sub(uint256 _a, uint256 _b) internal pure returns (uint256) {
    require(_b <= _a);
    uint256 c = _a - _b;

    return c;
  }

  /**
  * @dev Adds two numbers, reverts on overflow.
  */
  function add(uint256 _a, uint256 _b) internal pure returns (uint256) {
    uint256 c = _a + _b;
    require(c >= _a);

    return c;
  }

  /**
  * @dev Divides two numbers and returns the remainder (unsigned integer modulo),
  * reverts when dividing by zero.
  */
  function mod(uint256 a, uint256 b) internal pure returns (uint256) {
    require(b != 0);
    return a % b;
  }
}


/**
 * @title ERC165
 * @dev https://github.com/ethereum/EIPs/blob/master/EIPS/eip-165.md
 */
interface ERC165 {

  /**
   * @notice Query if a contract implements an interface
   * @param _interfaceId The interface identifier, as specified in ERC-165
   * @dev Interface identification is specified in ERC-165. This function
   * uses less than 30,000 gas.
   */
  function supportsInterface(bytes4 _interfaceId)
    external
    view
    returns (bool);
}

/**
 * @title SupportsInterfaceWithLookup
 * @author Matt Condon (@shrugs)
 * @dev Implements ERC165 using a lookup table.
 */
contract SupportsInterfaceWithLookup is ERC165 {

  bytes4 public constant InterfaceId_ERC165 = 0x01ffc9a7;
  /**
   * 0x01ffc9a7 ===
   *   bytes4(keccak256('supportsInterface(bytes4)'))
   */

  /**
   * @dev a mapping of interface id to whether or not it's supported
   */
  mapping(bytes4 => bool) internal supportedInterfaces;

  /**
   * @dev A contract implementing SupportsInterfaceWithLookup
   * implement ERC165 itself
   */
  constructor()
    public
  {
    _registerInterface(InterfaceId_ERC165);
  }

  /**
   * @dev implement supportsInterface(bytes4) using a lookup table
   */
  function supportsInterface(bytes4 _interfaceId)
    external
    view
    returns (bool)
  {
    return supportedInterfaces[_interfaceId];
  }

  /**
   * @dev private method for registering an interface
   */
  function _registerInterface(bytes4 _interfaceId)
    internal
  {
    require(_interfaceId != 0xffffffff);
    supportedInterfaces[_interfaceId] = true;
  }
}



/**
 * @title ERC721 token receiver interface
 * @dev Interface for any contract that wants to support safeTransfers
 * from ERC721 asset contracts.
 */
contract ERC721Receiver {
  /**
   * @dev Magic value to be returned upon successful reception of an NFT
   *  Equals to `bytes4(keccak256("onERC721Received(address,address,uint256,bytes)"))`,
   *  which can be also obtained as `ERC721Receiver(0).onERC721Received.selector`
   */
  bytes4 internal constant ERC721_RECEIVED = 0x150b7a02;

  /**
   * @notice Handle the receipt of an NFT
   * @dev The ERC721 smart contract calls this function on the recipient
   * after a `safetransfer`. This function MAY throw to revert and reject the
   * transfer. Return of other than the magic value MUST result in the
   * transaction being reverted.
   * Note: the contract address is always the message sender.
   * @param _operator The address which called `safeTransferFrom` function
   * @param _from The address which previously owned the token
   * @param _tokenId The NFT identifier which is being transferred
   * @param _data Additional data with no specified format
   * @return `bytes4(keccak256("onERC721Received(address,address,uint256,bytes)"))`
   */
  function onERC721Received(
    address _operator,
    address _from,
    uint256 _tokenId,
    bytes _data
  )
    public
    returns(bytes4);
}


/**
 * @title ERC721 Non-Fungible Token Standard basic interface
 * @dev see https://github.com/ethereum/EIPs/blob/master/EIPS/eip-721.md
 */
contract ERC721Basic is ERC165 {

  bytes4 internal constant InterfaceId_ERC721 = 0x80ac58cd;
  /*
   * 0x80ac58cd ===
   *   bytes4(keccak256('balanceOf(address)')) ^
   *   bytes4(keccak256('ownerOf(uint256)')) ^
   *   bytes4(keccak256('approve(address,uint256)')) ^
   *   bytes4(keccak256('getApproved(uint256)')) ^
   *   bytes4(keccak256('setApprovalForAll(address,bool)')) ^
   *   bytes4(keccak256('isApprovedForAll(address,address)')) ^
   *   bytes4(keccak256('transferFrom(address,address,uint256)')) ^
   *   bytes4(keccak256('safeTransferFrom(address,address,uint256)')) ^
   *   bytes4(keccak256('safeTransferFrom(address,address,uint256,bytes)'))
   */

  bytes4 internal constant InterfaceId_ERC721Enumerable = 0x780e9d63;
  /**
   * 0x780e9d63 ===
   *   bytes4(keccak256('totalSupply()')) ^
   *   bytes4(keccak256('tokenOfOwnerByIndex(address,uint256)')) ^
   *   bytes4(keccak256('tokenByIndex(uint256)'))
   */

  bytes4 internal constant InterfaceId_ERC721Metadata = 0x5b5e139f;
  /**
   * 0x5b5e139f ===
   *   bytes4(keccak256('name()')) ^
   *   bytes4(keccak256('symbol()')) ^
   *   bytes4(keccak256('tokenURI(uint256)'))
   */

  event Transfer(
    address indexed _from,
    address indexed _to,
    uint256 indexed _tokenId
  );
  event Approval(
    address indexed _owner,
    address indexed _approved,
    uint256 indexed _tokenId
  );
  event ApprovalForAll(
    address indexed _owner,
    address indexed _operator,
    bool _approved
  );

  function balanceOf(address _owner) public view returns (uint256 _balance);
  function ownerOf(uint256 _tokenId) public view returns (address _owner);

  function approve(address _to, uint256 _tokenId) public;
  function getApproved(uint256 _tokenId)
    public view returns (address _operator);

  function setApprovalForAll(address _operator, bool _approved) public;
  function isApprovedForAll(address _owner, address _operator)
    public view returns (bool);

  function transferFrom(address _from, address _to, uint256 _tokenId) public;
  function safeTransferFrom(address _from, address _to, uint256 _tokenId)
    public;

  function safeTransferFrom(
    address _from,
    address _to,
    uint256 _tokenId,
    bytes _data
  )
    public;
}

/**
 * @title ERC-721 Non-Fungible Token Standard, optional enumeration extension
 * @dev See https://github.com/ethereum/EIPs/blob/master/EIPS/eip-721.md
 */
contract ERC721Enumerable is ERC721Basic {
  function totalSupply() public view returns (uint256);
  function tokenOfOwnerByIndex(
    address _owner,
    uint256 _index
  )
    public
    view
    returns (uint256 _tokenId);

  function tokenByIndex(uint256 _index) public view returns (uint256);
}


/**
 * @title ERC-721 Non-Fungible Token Standard, optional metadata extension
 * @dev See https://github.com/ethereum/EIPs/blob/master/EIPS/eip-721.md
 */
contract ERC721Metadata is ERC721Basic {
  function name() external view returns (string _name);
  function symbol() external view returns (string _symbol);
  function tokenURI(uint256 _tokenId) public view returns (string);
}

/**
 * @title ERC-721 Non-Fungible Token Standard, full implementation interface
 * @dev See https://github.com/ethereum/EIPs/blob/master/EIPS/eip-721.md
 */
contract ERC721 is ERC721Basic, ERC721Enumerable, ERC721Metadata {
}

/**
 * @title ERC721 Non-Fungible Token Standard basic implementation
 * @dev see https://github.com/ethereum/EIPs/blob/master/EIPS/eip-721.md
 */
contract ERC721BasicToken is SupportsInterfaceWithLookup, ERC721Basic {

  using SafeMath for uint256;
  using AddressUtils for address;

  // Equals to `bytes4(keccak256("onERC721Received(address,address,uint256,bytes)"))`
  // which can be also obtained as `ERC721Receiver(0).onERC721Received.selector`
  bytes4 private constant ERC721_RECEIVED = 0x150b7a02;

  // Mapping from token ID to owner
  mapping (uint256 => address) internal tokenOwner;

  // Mapping from token ID to approved address
  mapping (uint256 => address) internal tokenApprovals;

  // Mapping from owner to number of owned token
  mapping (address => uint256) internal ownedTokensCount;

  // Mapping from owner to operator approvals
  mapping (address => mapping (address => bool)) internal operatorApprovals;

  constructor()
    public
  {
    // register the supported interfaces to conform to ERC721 via ERC165
    _registerInterface(InterfaceId_ERC721);
  }

  /**
   * @dev Gets the balance of the specified address
   * @param _owner address to query the balance of
   * @return uint256 representing the amount owned by the passed address
   */
  function balanceOf(address _owner) public view returns (uint256) {
    require(_owner != address(0));
    return ownedTokensCount[_owner];
  }

  /**
   * @dev Gets the owner of the specified token ID
   * @param _tokenId uint256 ID of the token to query the owner of
   * @return owner address currently marked as the owner of the given token ID
   */
  function ownerOf(uint256 _tokenId) public view returns (address) {
    address owner = tokenOwner[_tokenId];
    require(owner != address(0));
    return owner;
  }

  /**
   * @dev Approves another address to transfer the given token ID
   * The zero address indicates there is no approved address.
   * There can only be one approved address per token at a given time.
   * Can only be called by the token owner or an approved operator.
   * @param _to address to be approved for the given token ID
   * @param _tokenId uint256 ID of the token to be approved
   */
  function approve(address _to, uint256 _tokenId) public {
    address owner = ownerOf(_tokenId);
    require(_to != owner);
    require(msg.sender == owner || isApprovedForAll(owner, msg.sender));

    tokenApprovals[_tokenId] = _to;
    emit Approval(owner, _to, _tokenId);
  }

  /**
   * @dev Gets the approved address for a token ID, or zero if no address set
   * @param _tokenId uint256 ID of the token to query the approval of
   * @return address currently approved for the given token ID
   */
  function getApproved(uint256 _tokenId) public view returns (address) {
    return tokenApprovals[_tokenId];
  }

  /**
   * @dev Sets or unsets the approval of a given operator
   * An operator is allowed to transfer all tokens of the sender on their behalf
   * @param _to operator address to set the approval
   * @param _approved representing the status of the approval to be set
   */
  function setApprovalForAll(address _to, bool _approved) public {
    require(_to != msg.sender);
    operatorApprovals[msg.sender][_to] = _approved;
    emit ApprovalForAll(msg.sender, _to, _approved);
  }

  /**
   * @dev Tells whether an operator is approved by a given owner
   * @param _owner owner address which you want to query the approval of
   * @param _operator operator address which you want to query the approval of
   * @return bool whether the given operator is approved by the given owner
   */
  function isApprovedForAll(
    address _owner,
    address _operator
  )
    public
    view
    returns (bool)
  {
    return operatorApprovals[_owner][_operator];
  }

  /**
   * @dev Transfers the ownership of a given token ID to another address
   * Usage of this method is discouraged, use `safeTransferFrom` whenever possible
   * Requires the msg sender to be the owner, approved, or operator
   * @param _from current owner of the token
   * @param _to address to receive the ownership of the given token ID
   * @param _tokenId uint256 ID of the token to be transferred
  */
  function transferFrom(
    address _from,
    address _to,
    uint256 _tokenId
  )
    public
  {
    require(isApprovedOrOwner(msg.sender, _tokenId));
    require(_to != address(0));

    clearApproval(_from, _tokenId);
    removeTokenFrom(_from, _tokenId);
    addTokenTo(_to, _tokenId);

    emit Transfer(_from, _to, _tokenId);
  }

  /**
   * @dev Safely transfers the ownership of a given token ID to another address
   * If the target address is a contract, it must implement `onERC721Received`,
   * which is called upon a safe transfer, and return the magic value
   * `bytes4(keccak256("onERC721Received(address,address,uint256,bytes)"))`; otherwise,
   * the transfer is reverted.
   *
   * Requires the msg sender to be the owner, approved, or operator
   * @param _from current owner of the token
   * @param _to address to receive the ownership of the given token ID
   * @param _tokenId uint256 ID of the token to be transferred
  */
  function safeTransferFrom(
    address _from,
    address _to,
    uint256 _tokenId
  )
    public
  {
    // solium-disable-next-line arg-overflow
    safeTransferFrom(_from, _to, _tokenId, "");
  }

  /**
   * @dev Safely transfers the ownership of a given token ID to another address
   * If the target address is a contract, it must implement `onERC721Received`,
   * which is called upon a safe transfer, and return the magic value
   * `bytes4(keccak256("onERC721Received(address,address,uint256,bytes)"))`; otherwise,
   * the transfer is reverted.
   * Requires the msg sender to be the owner, approved, or operator
   * @param _from current owner of the token
   * @param _to address to receive the ownership of the given token ID
   * @param _tokenId uint256 ID of the token to be transferred
   * @param _data bytes data to send along with a safe transfer check
   */
  function safeTransferFrom(
    address _from,
    address _to,
    uint256 _tokenId,
    bytes _data
  )
    public
  {
    transferFrom(_from, _to, _tokenId);
    // solium-disable-next-line arg-overflow
    require(checkAndCallSafeTransfer(_from, _to, _tokenId, _data));
  }

  /**
   * @dev Returns whether the specified token exists
   * @param _tokenId uint256 ID of the token to query the existence of
   * @return whether the token exists
   */
  function _exists(uint256 _tokenId) internal view returns (bool) {
    address owner = tokenOwner[_tokenId];
    return owner != address(0);
  }

  /**
   * @dev Returns whether the given spender can transfer a given token ID
   * @param _spender address of the spender to query
   * @param _tokenId uint256 ID of the token to be transferred
   * @return bool whether the msg.sender is approved for the given token ID,
   *  is an operator of the owner, or is the owner of the token
   */
  function isApprovedOrOwner(
    address _spender,
    uint256 _tokenId
  )
    internal
    view
    returns (bool)
  {
    address owner = ownerOf(_tokenId);
    // Disable solium check because of
    // https://github.com/duaraghav8/Solium/issues/175
    // solium-disable-next-line operator-whitespace
    return (
      _spender == owner ||
      getApproved(_tokenId) == _spender ||
      isApprovedForAll(owner, _spender)
    );
  }

  /**
   * @dev Internal function to mint a new token
   * Reverts if the given token ID already exists
   * @param _to The address that will own the minted token
   * @param _tokenId uint256 ID of the token to be minted by the msg.sender
   */
  function _mint(address _to, uint256 _tokenId) internal {
    require(_to != address(0));
    addTokenTo(_to, _tokenId);
    emit Transfer(address(0), _to, _tokenId);
  }

  /**
   * @dev Internal function to burn a specific token
   * Reverts if the token does not exist
   * @param _tokenId uint256 ID of the token being burned by the msg.sender
   */
  function _burn(address _owner, uint256 _tokenId) internal {
    clearApproval(_owner, _tokenId);
    removeTokenFrom(_owner, _tokenId);
    emit Transfer(_owner, address(0), _tokenId);
  }

  /**
   * @dev Internal function to clear current approval of a given token ID
   * Reverts if the given address is not indeed the owner of the token
   * @param _owner owner of the token
   * @param _tokenId uint256 ID of the token to be transferred
   */
  function clearApproval(address _owner, uint256 _tokenId) internal {
    require(ownerOf(_tokenId) == _owner);
    if (tokenApprovals[_tokenId] != address(0)) {
      tokenApprovals[_tokenId] = address(0);
    }
  }

  /**
   * @dev Internal function to add a token ID to the list of a given address
   * @param _to address representing the new owner of the given token ID
   * @param _tokenId uint256 ID of the token to be added to the tokens list of the given address
   */
  function addTokenTo(address _to, uint256 _tokenId) internal {
    require(tokenOwner[_tokenId] == address(0));
    tokenOwner[_tokenId] = _to;
    ownedTokensCount[_to] = ownedTokensCount[_to].add(1);
  }

  /**
   * @dev Internal function to remove a token ID from the list of a given address
   * @param _from address representing the previous owner of the given token ID
   * @param _tokenId uint256 ID of the token to be removed from the tokens list of the given address
   */
  function removeTokenFrom(address _from, uint256 _tokenId) internal {
    require(ownerOf(_tokenId) == _from);
    ownedTokensCount[_from] = ownedTokensCount[_from].sub(1);
    tokenOwner[_tokenId] = address(0);
  }

  /**
   * @dev Internal function to invoke `onERC721Received` on a target address
   * The call is not executed if the target address is not a contract
   * @param _from address representing the previous owner of the given token ID
   * @param _to target address that will receive the tokens
   * @param _tokenId uint256 ID of the token to be transferred
   * @param _data bytes optional data to send along with the call
   * @return whether the call correctly returned the expected magic value
   */
  function checkAndCallSafeTransfer(
    address _from,
    address _to,
    uint256 _tokenId,
    bytes _data
  )
    internal
    returns (bool)
  {
    if (!_to.isContract()) {
      return true;
    }
    bytes4 retval = ERC721Receiver(_to).onERC721Received(
      msg.sender, _from, _tokenId, _data);
    return (retval == ERC721_RECEIVED);
  }
}

/**
 * @title Full ERC721 Token
 * This implementation includes all the required and some optional functionality of the ERC721 standard
 * Moreover, it includes approve all functionality using operator terminology
 * @dev see https://github.com/ethereum/EIPs/blob/master/EIPS/eip-721.md
 */
contract ERC721Token is SupportsInterfaceWithLookup, ERC721BasicToken, ERC721 {

  // Token name
  string internal name_;

  // Token symbol
  string internal symbol_;

  // Mapping from owner to list of owned token IDs
  mapping(address => uint256[]) internal ownedTokens;

  // Mapping from token ID to index of the owner tokens list
  mapping(uint256 => uint256) internal ownedTokensIndex;

  // Array with all token ids, used for enumeration
  uint256[] internal allTokens;

  // Mapping from token id to position in the allTokens array
  mapping(uint256 => uint256) internal allTokensIndex;

  // Optional mapping for token URIs
  mapping(uint256 => string) internal tokenURIs;

  /**
   * @dev Constructor function
   */
  constructor(string _name, string _symbol) public {
    name_ = _name;
    symbol_ = _symbol;

    // register the supported interfaces to conform to ERC721 via ERC165
    _registerInterface(InterfaceId_ERC721Enumerable);
    _registerInterface(InterfaceId_ERC721Metadata);
  }

  /**
   * @dev Gets the token name
   * @return string representing the token name
   */
  function name() external view returns (string) {
    return name_;
  }

  /**
   * @dev Gets the token symbol
   * @return string representing the token symbol
   */
  function symbol() external view returns (string) {
    return symbol_;
  }

  /**
   * @dev Returns an URI for a given token ID
   * Throws if the token ID does not exist. May return an empty string.
   * @param _tokenId uint256 ID of the token to query
   */
  function tokenURI(uint256 _tokenId) public view returns (string) {
    require(_exists(_tokenId));
    return tokenURIs[_tokenId];
  }

  /**
   * @dev Gets the token ID at a given index of the tokens list of the requested owner
   * @param _owner address owning the tokens list to be accessed
   * @param _index uint256 representing the index to be accessed of the requested tokens list
   * @return uint256 token ID at the given index of the tokens list owned by the requested address
   */
  function tokenOfOwnerByIndex(
    address _owner,
    uint256 _index
  )
    public
    view
    returns (uint256)
  {
    require(_index < balanceOf(_owner));
    return ownedTokens[_owner][_index];
  }

  /**
   * @dev Gets the total amount of tokens stored by the contract
   * @return uint256 representing the total amount of tokens
   */
  function totalSupply() public view returns (uint256) {
    return allTokens.length;
  }

  /**
   * @dev Gets the token ID at a given index of all the tokens in this contract
   * Reverts if the index is greater or equal to the total number of tokens
   * @param _index uint256 representing the index to be accessed of the tokens list
   * @return uint256 token ID at the given index of the tokens list
   */
  function tokenByIndex(uint256 _index) public view returns (uint256) {
    require(_index < totalSupply());
    return allTokens[_index];
  }

  /**
   * @dev Internal function to set the token URI for a given token
   * Reverts if the token ID does not exist
   * @param _tokenId uint256 ID of the token to set its URI
   * @param _uri string URI to assign
   */
  function _setTokenURI(uint256 _tokenId, string _uri) internal {
    require(_exists(_tokenId));
    tokenURIs[_tokenId] = _uri;
  }

  /**
   * @dev Internal function to add a token ID to the list of a given address
   * @param _to address representing the new owner of the given token ID
   * @param _tokenId uint256 ID of the token to be added to the tokens list of the given address
   */
  function addTokenTo(address _to, uint256 _tokenId) internal {
    super.addTokenTo(_to, _tokenId);
    uint256 length = ownedTokens[_to].length;
    ownedTokens[_to].push(_tokenId);
    ownedTokensIndex[_tokenId] = length;
  }

  /**
   * @dev Internal function to remove a token ID from the list of a given address
   * @param _from address representing the previous owner of the given token ID
   * @param _tokenId uint256 ID of the token to be removed from the tokens list of the given address
   */
  function removeTokenFrom(address _from, uint256 _tokenId) internal {
    super.removeTokenFrom(_from, _tokenId);

    // To prevent a gap in the array, we store the last token in the index of the token to delete, and
    // then delete the last slot.
    uint256 tokenIndex = ownedTokensIndex[_tokenId];
    uint256 lastTokenIndex = ownedTokens[_from].length.sub(1);
    uint256 lastToken = ownedTokens[_from][lastTokenIndex];

    ownedTokens[_from][tokenIndex] = lastToken;
    // This also deletes the contents at the last position of the array
    ownedTokens[_from].length--;

    // Note that this will handle single-element arrays. In that case, both tokenIndex and lastTokenIndex are going to
    // be zero. Then we can make sure that we will remove _tokenId from the ownedTokens list since we are first swapping
    // the lastToken to the first position, and then dropping the element placed in the last position of the list

    ownedTokensIndex[_tokenId] = 0;
    ownedTokensIndex[lastToken] = tokenIndex;
  }

  /**
   * @dev Internal function to mint a new token
   * Reverts if the given token ID already exists
   * @param _to address the beneficiary that will own the minted token
   * @param _tokenId uint256 ID of the token to be minted by the msg.sender
   */
  function _mint(address _to, uint256 _tokenId) internal {
    super._mint(_to, _tokenId);

    allTokensIndex[_tokenId] = allTokens.length;
    allTokens.push(_tokenId);
  }

  /**
   * @dev Internal function to burn a specific token
   * Reverts if the token does not exist
   * @param _owner owner of the token to burn
   * @param _tokenId uint256 ID of the token being burned by the msg.sender
   */
  function _burn(address _owner, uint256 _tokenId) internal {
    super._burn(_owner, _tokenId);

    // Clear metadata (if any)
    if (bytes(tokenURIs[_tokenId]).length != 0) {
      delete tokenURIs[_tokenId];
    }

    // Reorg all tokens array
    uint256 tokenIndex = allTokensIndex[_tokenId];
    uint256 lastTokenIndex = allTokens.length.sub(1);
    uint256 lastToken = allTokens[lastTokenIndex];

    allTokens[tokenIndex] = lastToken;
    allTokens[lastTokenIndex] = 0;

    allTokens.length--;
    allTokensIndex[_tokenId] = 0;
    allTokensIndex[lastToken] = tokenIndex;
  }

}

contract TicketSale is ERC721Token {
  using SafeMath for uint8;
  using SafeMath for uint16;
  using SafeMath for uint256;

  Hub hub;
  mapping(address => bool) public hosts;
  // 目前正在二級市場交易的票數，一開始為 0 張。
  uint256 public tradings = 0;
  // 儲存目前所有正在二級市場交易的活動票券資訊，Trade struct 包含擁有者（購買者）、票券 Id，販售價格資訊。
  mapping(uint256 => Trade) public tradingList;
  // 紀錄此活動票券擁有者（購買者）正在二級市場販售的活動票券數。
  mapping(address => uint256) public tradingTicketsOfOwner;
  mapping(uint256 => bool) public usedTickets;

  // 目前已販售的票數，一開始為販售 0 張。
  uint256 public tickets = 0;
  // 此活動票券價格，單位是 wei。
  uint256 public price = 20000000000000000;
  // 此活動下一張的票券價格，單位是 wei。
  uint256 public nextPrice = 400000000000000;
  // 限定同一個買家最多可購買此活動票券的數量。
  uint8 public limit = 10;
  // 可賣出的活動票券數量上限值。
  uint256 public maxAttendees = 10;
  // tradeFee 為二級市場的交易服務費（團隊營運費），單位是 wei。
  uint16 public tradeFee = 0;
  // 在使用 setUsedTickets(uint256[]) 函數時，限制輸入的矩陣長度，以防止交易手續費過高。
  uint8 public maxMarkedTickets = 10;

  uint256 public startTime;
  uint256 public dueTime;

  // serviceFeeRatio / 10000 = ratio
  // 收取的服務費（團隊營運費）serviceFee 為：票價 * serviceFeeRatio / 10000，單位是 wei。
  uint16 public serviceFeeRatio = 100;
  uint256 public serviceFee = 0;
  // 收取的回饋費為 rewardFee 為：票價 * r回ewardFeeRatio / 10000，單位是 wei。
  uint16 public rewardFeeRatio = 240;
  uint256 public rewardFee = 0;

  // 儲存目前正在二級市場交易的活動票券資訊，包含擁有者（購買者）、票券 Id，販售價格資訊。
  struct Trade {
    address owner;
    uint256 ticketId;
    uint256 value;
  }

  constructor (string name, address hubAddr, uint16 ratio, uint256 _startTime, uint256 _dueTime, uint256 _price) public
      ERC721Token(name, name)
  {
    hub = Hub(hubAddr);
    hosts[msg.sender] = true;
    serviceFeeRatio = ratio;
    startTime = _startTime;
    dueTime = _dueTime;
    price = _price;
  }

  // 條飾詞：只能是這個活動票券的 Host。
  // 在這邊使用 Hub 合約來建立 TicketSale 子合約，其 Host 為 Hub 合約地址、以及建立 Hub 合約的人。細節請查看 Hub 合約內的 createEvent(string,uint256,uint256,uint256) 函數。
  modifier onlyHost() {
    require(hosts[msg.sender] == true, "only host can access it");
    _;
  }

  // ticketAmount 是購買者 msg 想要一次購買的票券數量。
  //function register(uint8 ticketAmount) external payable {
  function register() external payable {
    // 購買者發起的交易，傳入的 ETH 要大於「票價 * 購買數量」。
    //require(msg.value >= price.mul(ticketAmount), "not enough value");
    require(msg.value >= price, "not enough value");
    // 每一個購買者所持有的活動票數，必須大於 limit 變數。
    //require(balanceOf(msg.sender).add(ticketAmount) <= limit, "too many ticket amount");
    require(balanceOf(msg.sender) <= limit, "too many ticket amount");
    // 賣出的活動票券總數量一定要小於 maxAttendees 變數。
    //require(tickets < maxAttendees, "register is ended");
    //require(tickets.add(ticketAmount) <= maxAttendees, "too many ticket register");
    require(tickets <= maxAttendees, "too many ticket register");

    // 是否在可以買票的開放時間內。
    if (startTime != 0) {      
      require(now >= startTime, "ticket sale is not started yet");
    }
    if (dueTime != 0) {
      require(now <= dueTime, "ticket sale is ended");
    }

    // 收取的服務費（團隊營運費）serviceFee 為：票價 * serviceFeeRatio / 10000。
    //serviceFee = serviceFee.add(msg.value.mul(serviceFeeRatio).div(10000));
    serviceFee = serviceFee.add(price.mul(serviceFeeRatio).div(10000));
    rewardFee = rewardFee.add(price.mul(rewardFeeRatio).div(10000));

    // 這邊的 index 只是迴圈用的暫時變數，與 ticket id 無關。
    // for (uint256 index = 0; index < ticketAmount; index++) {
    //   // 目前的總票券數量（ticketId）加 1。
    //   tickets++;
    //   // 函數：_mint(address _to, uint256 _tokenId)：買家從 0x0 地址取得 Id 為 tickets 的票券
    //   _mint(msg.sender, tickets);
    // }
    tickets = tickets.add(1);
    _mint(msg.sender, tickets);

    // 如果買家傳入合約的 ETH 超過「票價 * 購買數量」，就將多出來的費用退還給買家。
    //uint256 rest = msg.value.sub(ticketAmount.mul(price));
    uint256 rest = msg.value.sub(price);
    if (rest > 0) {
      msg.sender.transfer(rest);
    }
    price = price.add(tickets.mul(nextPrice));
  }

  // 設置可賣出的活動票券數量上限值。
  //function setMaxAttendees(uint256 _maxAttendees) external {
  function setMaxAttendees(uint256 _maxAttendees) external onlyHost {
    require(_maxAttendees > maxAttendees, "maxAttendees must be greater than original");
    maxAttendees = _maxAttendees;
  }

  // 新增這個活動票券的 Host。
  function setHost(address host) external onlyHost {
    hosts[host] = true;
  }

  // 將指定的 Address 從這個這個活動票券的 Host 中移除。
  function removeHost(address addr) external onlyHost {
    require(hosts[addr] == true, "only host can be remove");
    hosts[addr] = false;
  }

  // 設置同一個買家最多可購買此活動票券的數量。
  function setLimit(uint8 _limit) external onlyHost {
    require(_limit > limit, "limit must be greater than original");
    limit = _limit;
  }

  // 設置此活動票券的價格，單位是 wei。
  function setPrice(uint256 _price) external onlyHost {
    price = _price;
  }

  function setNextPrice(uint256 _nextPrice) external onlyHost {
    nextPrice = _nextPrice;
  }

  // 設置二級市場的交易手續費（團隊營運費）serviceFee 為：票價 * tradeFeeRatio / 10000。
  function setTradeFee(uint16 _fee) external onlyHost {
    tradeFee = _fee;
  }

  // 活動票券的擁有者（購買者）在二級市場發出想將此活動票券賣出的資訊。
  function requestTrading(uint256 _ticketId, uint256 _value) external {
    // 必須是票券擁有者。
    address owner = ownerOf(_ticketId);
    require(msg.sender == owner || isApprovedForAll(owner, msg.sender), "only owner can sell");
    // 目前正在二級市場交易的票數加 1。
    tradings++;
    // 儲存目前正在二級市場交易的活動票券資訊，包含擁有者（購買者）、票券 Id，販售價格。
    tradingList[tradings] = Trade({ticketId: _ticketId, value: _value, owner: msg.sender});
    // 此活動票券擁有者（購買者）正在二級市場販售的活動票券數加 1。
    tradingTicketsOfOwner[msg.sender]++;
  }

  // 在二級市場完成交易，或活動票券擁有者（購買者）欲中途取消交易時，執行此程式。
  function _cancelTrade(uint256 _tradeId, address _owner) internal {
    // 此張活動票券擁有者（購買者）正在二級市場販售的活動票券數減 1。
    tradingTicketsOfOwner[_owner]--;
    // 刪除此張目前正在二級市場交易的活動票券資訊。
    delete tradingList[_tradeId];
    // 目前正在二級市場交易的票數減 1。
    tradings--;
  }

  // 取消目前正在二級市場交易的活動票券。
  function cancelTrade(uint256 _tradeId) external {
    // 取出目前正在二級市場交易的活動票券資訊。
    Trade memory t = tradingList[_tradeId];
    // 發起活動票券取消在二級市場交易的地址必須為票券擁有者（購買者）。
    require(t.owner == msg.sender, "only owner can cancel trading");
    // 執行取消交易。
    _cancelTrade(_tradeId, msg.sender);
  }

  // 有買家相中目前正在二級市場交易的活動票券，執行此購買票券函數。
  function trade(uint256 _tradeId) public payable {
    // 此票券必須正在 tradingList 的列表中，表示確實正在二級市場販售。
    require(tradingList[_tradeId].ticketId != 0, "");
    // 取出目前正在二級市場交易的活動票券資訊。
    Trade memory t = tradingList[_tradeId];
    // 買家（msg.sender）所傳入的 ETH 必須要大於此活動票券的購買價格。
    require(t.value <= msg.value, "ticket price is not enough");

    // 當在二級市場活動票券成交後，買家需向賣方支付購票費用。
    t.owner.transfer(t.value - tradeFee);
    // 收取的交易服務費（團隊營運費）serviceFee 為：tradeFee。
    serviceFee = serviceFee.add(tradeFee);
    // 如果買家傳入合約的 ETH 超過「二級市場交易的票價」，就將多出來的費用退還給買家。
    uint256 rest = msg.value.sub(t.value);
    if (rest > 0) {
      msg.sender.transfer(rest);
    }

    // 取消此活動票券可在二級市場販售的狀態。
    _cancelTrade(t.ticketId, t.owner);
    // 移除此活動票券的舊擁有者（擁有者將設置為 0x0 地址）。
    removeTokenFrom(t.owner, t.ticketId);
    // 此活動票券的擁有者設置為新買家（msg.sender）。
    addTokenTo(msg.sender, t.ticketId);
  }

  // 設定哪一些 ticket id 要被設置成已使用（買家已掃描 QR Code 並參與活動）。
  //function setUsedTickets(uint256[] _ticketIds) public onlyHost {
    function setUsedTickets(uint256 _ticketId) internal {
    // 限制批次設置已使用票券票的數量，不能大於 maxMarkedTickets，以避免手續費過高。
    //require(_ticketIds.length <= maxMarkedTickets, "set too many tickets in the same time");
    // 將 _ticketIds 矩陣內的 ticket id 設置為已使用。
    //for (uint8 index = 0; index < _ticketIds.length; index++) {
      //usedTickets[_ticketIds[index]] = true;
      usedTickets[_ticketId] = true;
    //}
  }

  // 查看此票券是否已使用。
  function isUsedTicket(uint256 _ticketId) public view returns (bool) {
    return usedTickets[_ticketId];
  }

  // 此 TicketSale 合約的 Host 可取出扣除 serviceFee 手續費（= 票價 * serviceFeeRatio / 10000）的活動票價金額。
  function withdraw() external onlyHost {
    // BUGGY HERE
    msg.sender.transfer(address(this).balance.sub(serviceFee));
  }

  // 從 TicketSale 合約發送 serviceFee 手續費（= 票價 * serviceFeeRatio / 10000）至 Hub 合約。
  function withdrawFee() external onlyHost {
    // Hub 合約地址必須不為 0x0 地址。
    require(address(hub) != address(0), "hub contract does not exist");
    // 儲存的 serviceFee 手續費的量必須要大於 0。
    require(serviceFee > 0, "service fee is not required");
    // 此 TicketSale 合約內的 ETH 必須要大於 serviceFee 手續費的量。
    require(serviceFee <= address(this).balance, "balance is not enough");

    // 從 TicketSale 合約發送 serviceFee 手續費（= 票價 * serviceFeeRatio / 10000）至 Hub 合約。
    hub.deposit.value(serviceFee)();
    // serviceFee 手續費的量清為 0。
    serviceFee = 0;
  }

bool public isWin = false;
bool public isSettle = false;
address private lastLose;
uint256 public totalBalance = 0;

  function getIsSettle() public view returns (bool){
      return isSettle;
  }
  
    function setIsSettle(bool _s) public onlyHost {
      isSettle = _s;
  }
  
  function setLastLose(address _lastLose) public onlyHost {
    require(_lastLose != address(0), "last lose contract does not exist");
    lastLose = _lastLose;
  }
  
  function addServiceFee(uint256 _serviceFee) public onlyHost {
      serviceFee = serviceFee.add(_serviceFee);
  }
  
  function addRewardFee(uint _rewardFee) public onlyHost {
      rewardFee = rewardFee.add(_rewardFee);
  }
  
  function addTotalBalance(uint256 _totalBalance) public onlyHost {
      totalBalance = totalBalance.add(_totalBalance);
  }

// 由 hub 或 host 可以執行此合約。
  function Win(address _winTicket) public onlyHost {
    // Hub 合約地址必須不為 0x0 地址。
    require(address(hub) != address(0), "hub contract does not exist");
    // 必須在購票結束時間。
    require(dueTime != 0 && now >= dueTime, "ticket is on sale");
    // 必須未清算
    require(isSettle == false, "already settle");
    // 若贏的地址不是此票券，就把所有錢傳過去 winTicket，再由那邊進行結算。
    if(address(_winTicket) != address(this)){
      // 設置贏的票券中 lastLose 的地址。
      //***
      TicketSale ts = TicketSale(_winTicket);
      ts.setLastLose(ownerOf(tickets));
      ts.addServiceFee(serviceFee);
      ts.addRewardFee(rewardFee);
      ts.addTotalBalance(address(this).balance);
      // 將輸的票券所有錢轉到贏的票券合約內。
      _winTicket.transfer(address(this).balance);
      isWin = false;
      isSettle = true;
    } else {
      addTotalBalance(address(this).balance);
      isWin = true;
      isSettle = true;
    }
  }

// 同一場比賽的不同隊票券要全部 settle 好，否則回傳 false。
// 應該是要指定哪此票是同一場比賽，但暫想不到紀錄票的解法。
//function allSettle(uint256[] sameEvents) view public returns(bool){
  function allSettle() view public returns(bool){
  //require(sameEvents.length <= hub.events(), "input.length must be less than or equal events");
  // public關鍵字自動生成的getter函數來獲取值
  // https://ethereum.stackexchange.com/questions/38317/access-to-public-variable-from-other-contract
  //for (uint8 index = 0; index < sameEvents.length; index++) {
  for (uint256 index = 1; index <= hub.getEvents(); index++) {
    // 不允許在其他合約中訪問 mapping 類型。您必須在數據協定中創建 getter / setter。
    // https://ethereum.stackexchange.com/questions/28199/accessing-mapping-from-another-contract
    TicketSale ts = TicketSale(hub.eventList(index));
    //TicketSale ts = TicketSale(hub.getEventList(sameEvents[index]));
    // public關鍵字自動生成的getter函數來獲取值
    // https://ethereum.stackexchange.com/questions/38317/access-to-public-variable-from-other-contract
    if(ts.isSettle() == false){
    //if(ts.getIsSettle() == false){
      return false;
    }
  }
  return true;
}

  function Reward(uint256 _ticketId, address _lastLose) external {
    // 必須是票券的擁有者、或買到輸的那一隊的最後一人。
    address owner = ownerOf(_ticketId);
    require(msg.sender == owner || isApprovedForAll(owner, msg.sender) || msg.sender == _lastLose, "only owner can sell");
    // 必須所有與這場比賽有關的票券都要清算。
    require(allSettle() == true, "must be all settle");
    // Hub 合約地址必須不為 0x0 地址。
    require(address(hub) != address(0), "hub contract does not exist");
    // 必須在購票結束時間。
    require(dueTime != 0 && now >= dueTime, "ticket is on sale");
    // 此票券必須是贏的那一隊才可以執行。
    require(isWin == true, "this ticket is lose team");

    if(msg.sender == _lastLose){
      // 分配獎項。
      _lastLose.transfer(rewardFee);
      // 設置獎勵為 0，避免被再次領獎。
      rewardFee = 0;
    } else{
        // 此票券必須沒有使用過。
        require(usedTickets[_ticketId] == false);
        // 設置此票券已使用過。
        setUsedTickets(_ticketId);
        // 分配獎項。
        owner.transfer((totalBalance.sub(serviceFee).sub(rewardFee)).div(tickets));
    }
  }

}

contract Hub {
  mapping (address => bool) admins;
  mapping (uint256 => address) public eventList;
  uint256 public events;
  uint16 public serviceFeeRate = 500;

  constructor() public {
    admins[msg.sender] = true;
  }

  modifier onlyAdmin() {
    require(admins[msg.sender] == true, "only admin can access it");
    _;
  }

  function isAdmin() external view returns (bool) {
    return admins[msg.sender];
  }

  function setFeeRate(uint16 rate) external onlyAdmin {
    serviceFeeRate = rate;
  }

  function createEvent(string name, uint256 startTime, uint256 dueTime, uint256 price) external {
    events++;
    TicketSale ts = new TicketSale(name, address(this), serviceFeeRate, startTime, dueTime, price);
    ts.setHost(msg.sender);
    eventList[events] = ts;
  }

  function withdraw() external onlyAdmin {
    msg.sender.transfer(address(this).balance);
  }

  function getBalance() external view returns (uint256) {
    return address(this).balance;
  }

  function getEventList(uint256 id) public view returns (address){
      require(id <= events, "input must be less than or equal events");
      return eventList[id];
  }
  
  function getEvents() public view returns (uint256){
      return events;
  }

  function() external payable {}

  function deposit() external payable {}
}
