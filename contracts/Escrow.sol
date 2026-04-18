// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title Escrow
/// @notice Holds client payments and releases them only on job completion or failure.
///         The Marketplace is the only authorized caller.
contract Escrow {

    address public marketplace;

    struct EscrowEntry {
        address client;
        uint256 amount;
        bool    released;
    }

    mapping(uint256 => EscrowEntry) public entries;

    event Deposited(uint256 indexed jobId, address indexed client, uint256 amount);
    event Released(uint256 indexed jobId, address indexed recipient, uint256 amount);
    event Refunded(uint256 indexed jobId, address indexed client, uint256 amount);

    modifier onlyMarketplace() {
        require(msg.sender == marketplace, "Only marketplace");
        _;
    }

    /// @dev Marketplace address is set after deployment to avoid circular dependency.
    constructor() {}

    function setMarketplace(address marketplaceAddr) external {
        require(marketplace == address(0), "Already set");
        marketplace = marketplaceAddr;
    }

    /// @notice Deposit funds for a specific job
    function deposit(uint256 jobId, address client) external payable onlyMarketplace {
        require(entries[jobId].client == address(0), "Already deposited");
        require(msg.value > 0, "No value sent");

        entries[jobId] = EscrowEntry({
            client:   client,
            amount:   msg.value,
            released: false
        });

        emit Deposited(jobId, client, msg.value);
    }

    /// @notice Release funds to provider after successful job completion
    function release(uint256 jobId, address payable recipient) external onlyMarketplace {
        EscrowEntry storage entry = entries[jobId];
        require(entry.client != address(0), "No escrow entry");
        require(!entry.released, "Already released");

        entry.released = true;
        uint256 amount = entry.amount;

        (bool ok,) = recipient.call{value: amount}("");
        require(ok, "Transfer failed");

        emit Released(jobId, recipient, amount);
    }

    /// @notice Refund client when a job fails or is cancelled
    function refund(uint256 jobId, address payable client) external onlyMarketplace {
        EscrowEntry storage entry = entries[jobId];
        require(entry.client != address(0), "No escrow entry");
        require(!entry.released, "Already released");
        require(entry.client == client, "Wrong client");

        entry.released = true;
        uint256 amount = entry.amount;

        (bool ok,) = client.call{value: amount}("");
        require(ok, "Refund failed");

        emit Refunded(jobId, client, amount);
    }

    function getEntry(uint256 jobId) external view returns (EscrowEntry memory) {
        return entries[jobId];
    }

    receive() external payable {}
}
