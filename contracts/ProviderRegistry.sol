// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title ProviderRegistry
/// @notice Tracks compute providers that offer resources on the NimbusX network.
contract ProviderRegistry {

    struct Provider {
        address providerAddress;
        uint256 cpu;          // cores
        uint256 memoryMB;     // megabytes
        uint256 pricePerHour; // wei per hour
        uint256 reputation;   // 0-100 score
        bool    active;
    }

    address public owner;
    address public marketplace;
    mapping(address => Provider) public providers;
    address[] public providerList;

    event ProviderRegistered(address indexed provider, uint256 cpu, uint256 memoryMB, uint256 pricePerHour);
    event ProviderDeregistered(address indexed provider);
    event ReputationUpdated(address indexed provider, uint256 newScore);

    modifier onlyOwner() {
        require(msg.sender == owner, "Not owner");
        _;
    }

    modifier onlyRegistered() {
        require(providers[msg.sender].active, "Provider not registered");
        _;
    }

    constructor() {
        owner = msg.sender;
    }

    /// @notice Set the authorized marketplace address
    function setMarketplace(address _marketplace) external onlyOwner {
        marketplace = _marketplace;
    }

    /// @notice Register as a compute provider
    function registerProvider(uint256 cpu, uint256 memoryMB, uint256 pricePerHour) external {
        require(cpu > 0 && memoryMB > 0, "Invalid resources");
        require(!providers[msg.sender].active, "Already registered");

        providers[msg.sender] = Provider({
            providerAddress: msg.sender,
            cpu: cpu,
            memoryMB: memoryMB,
            pricePerHour: pricePerHour,
            reputation: 100,
            active: true
        });
        providerList.push(msg.sender);

        emit ProviderRegistered(msg.sender, cpu, memoryMB, pricePerHour);
    }

    /// @notice Deregister and stop receiving jobs
    function deregisterProvider() external onlyRegistered {
        providers[msg.sender].active = false;
        emit ProviderDeregistered(msg.sender);
    }

    /// @notice Update the hourly price
    function updatePrice(uint256 newPrice) external onlyRegistered {
        providers[msg.sender].pricePerHour = newPrice;
    }

    /// @notice Called by the Marketplace to adjust reputation
    function updateReputation(address providerAddr, uint256 newScore) external {
        require(msg.sender == owner || msg.sender == marketplace, "Not authorized");
        require(newScore <= 100, "Score out of range");
        providers[providerAddr].reputation = newScore;
        emit ReputationUpdated(providerAddr, newScore);
    }

    function getProvider(address providerAddr) external view returns (Provider memory) {
        return providers[providerAddr];
    }

    function getProviderCount() external view returns (uint256) {
        return providerList.length;
    }

    function getAllProviders() external view returns (address[] memory) {
        return providerList;
    }
}