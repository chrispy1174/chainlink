// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../interfaces/LinkTokenInterface.sol";
import "../dev/VRFConsumerBase.sol";

contract VRFConsumer is VRFConsumerBase {

  uint256 public randomnessOutput;
  bytes32 public requestId;

  constructor(address _vrfCoordinator, address _link) public
    // solhint-disable-next-line no-empty-blocks
    VRFConsumerBase(_vrfCoordinator, _link) { /* empty */ }

  function fulfillRandomness(bytes32 _requestId, uint256 _randomness)
    internal override
  {
    randomnessOutput = _randomness;
    requestId = _requestId;
  }

  function testRequestRandomness(bytes32 _keyHash, uint256 _fee, uint256 _seed)
    external returns (bytes32 requestId)
  {
    return requestRandomness(_keyHash, _fee, _seed);
  }
}
