// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract ProviderRegistry {

    struct Provider {
        address providerAddress;
        uint cpu;
        uint memorySize;
        uint reputation;
        bool active;
    }

    mapping(address => Provider) public providers;

    event ProviderRegistered(address provider, uint cpu, uint memorySize);

    function registerProvider(uint cpu, uint memorySize) public {
        providers[msg.sender] = Provider({
            providerAddress: msg.sender,
            cpu: cpu,
            memorySize: memorySize,
            reputation: 100,
            active: true
        });

        emit ProviderRegistered(msg.sender, cpu, memorySize);
    }

    function getProvider(address providerAddr) public view returns (Provider memory) {
        return providers[providerAddr];
    }

    function updateReputation(address providerAddr, uint newScore) public {
        providers[providerAddr].reputation = newScore;
    }
}