// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./ProviderRegistry.sol";
import "./Escrow.sol";

/// @title Marketplace
/// @notice Users post jobs here; the scheduler matches them to providers.
contract Marketplace {

    enum JobStatus { Open, Assigned, Running, Completed, Failed, Cancelled }

    struct Job {
        uint256     jobId;
        address     client;
        address     assignedProvider;
        uint256     cpuRequired;
        uint256     memoryRequired;  // MB
        uint256     maxDuration;     // seconds
        uint256     maxPricePerHour; // wei
        JobStatus   status;
        uint256     createdAt;
        uint256     startedAt;
        uint256     completedAt;
        bytes32     resultHash;      // keccak256 of result data
    }

    uint256 public nextJobId;
    ProviderRegistry public registry;
    Escrow public escrow;
    address public scheduler; // off-chain scheduler address (trusted)

    mapping(uint256 => Job) public jobs;
    uint256[] public openJobs;

    event JobPosted(uint256 indexed jobId, address indexed client, uint256 cpuRequired, uint256 memoryRequired);
    event JobAssigned(uint256 indexed jobId, address indexed provider);
    event JobStarted(uint256 indexed jobId);
    event JobCompleted(uint256 indexed jobId, bytes32 resultHash);
    event JobFailed(uint256 indexed jobId);
    event JobCancelled(uint256 indexed jobId);

    modifier onlyScheduler() {
        require(msg.sender == scheduler, "Only scheduler");
        _;
    }

    modifier onlyProvider(uint256 jobId) {
        require(msg.sender == jobs[jobId].assignedProvider, "Not assigned provider");
        _;
    }

    modifier jobExists(uint256 jobId) {
        require(jobs[jobId].client != address(0), "Job not found");
        _;
    }

    constructor(address registryAddr, address payable escrowAddr, address schedulerAddr) {
        registry = ProviderRegistry(registryAddr);
        escrow = Escrow(escrowAddr);
        scheduler = schedulerAddr;
    }

    /// @notice Post a new compute job with escrowed payment
    function postJob(
        uint256 cpuRequired,
        uint256 memoryRequired,
        uint256 maxDuration,
        uint256 maxPricePerHour
    ) external payable returns (uint256) {
        require(cpuRequired > 0 && memoryRequired > 0, "Invalid job spec");
        require(msg.value > 0, "Must escrow payment");

        uint256 jobId = nextJobId++;
        jobs[jobId] = Job({
            jobId:            jobId,
            client:           msg.sender,
            assignedProvider: address(0),
            cpuRequired:      cpuRequired,
            memoryRequired:   memoryRequired,
            maxDuration:      maxDuration,
            maxPricePerHour:  maxPricePerHour,
            status:           JobStatus.Open,
            createdAt:        block.timestamp,
            startedAt:        0,
            completedAt:      0,
            resultHash:       bytes32(0)
        });

        openJobs.push(jobId);
        escrow.deposit{value: msg.value}(jobId, msg.sender);

        emit JobPosted(jobId, msg.sender, cpuRequired, memoryRequired);
        return jobId;
    }

    /// @notice Scheduler assigns a provider to an open job
    function assignJob(uint256 jobId, address provider) external onlyScheduler jobExists(jobId) {
        require(jobs[jobId].status == JobStatus.Open, "Job not open");
        require(registry.getProvider(provider).active, "Provider not active");

        jobs[jobId].assignedProvider = provider;
        jobs[jobId].status = JobStatus.Assigned;

        emit JobAssigned(jobId, provider);
    }

    /// @notice Provider confirms job has started
    function startJob(uint256 jobId) external onlyProvider(jobId) jobExists(jobId) {
        require(jobs[jobId].status == JobStatus.Assigned, "Job not assigned to you");

        jobs[jobId].status = JobStatus.Running;
        jobs[jobId].startedAt = block.timestamp;

        emit JobStarted(jobId);
    }

    /// @notice Provider marks job complete and submits result hash
    function completeJob(uint256 jobId, bytes32 resultHash) external onlyProvider(jobId) jobExists(jobId) {
        require(jobs[jobId].status == JobStatus.Running, "Job not running");

        jobs[jobId].status = JobStatus.Completed;
        jobs[jobId].completedAt = block.timestamp;
        jobs[jobId].resultHash = resultHash;

        // Release escrowed funds to provider
        escrow.release(jobId, payable(jobs[jobId].assignedProvider));

        emit JobCompleted(jobId, resultHash);
    }

    /// @notice Scheduler marks a job as failed and refunds client
    function failJob(uint256 jobId) external onlyScheduler jobExists(jobId) {
        require(
            jobs[jobId].status == JobStatus.Running ||
            jobs[jobId].status == JobStatus.Assigned,
            "Job not active"
        );

        jobs[jobId].status = JobStatus.Failed;

        // Refund escrowed funds to client
        escrow.refund(jobId, payable(jobs[jobId].client));

        // Slash provider reputation
        address provider = jobs[jobId].assignedProvider;
        ProviderRegistry.Provider memory p = registry.getProvider(provider);
        uint256 newRep = p.reputation > 10 ? p.reputation - 10 : 0;
        registry.updateReputation(provider, newRep);

        emit JobFailed(jobId);
    }

    /// @notice Client cancels an open (unassigned) job and gets refund
    function cancelJob(uint256 jobId) external jobExists(jobId) {
        require(jobs[jobId].client == msg.sender, "Not your job");
        require(jobs[jobId].status == JobStatus.Open, "Cannot cancel");

        jobs[jobId].status = JobStatus.Cancelled;
        escrow.refund(jobId, payable(msg.sender));

        emit JobCancelled(jobId);
    }

    function getJob(uint256 jobId) external view returns (Job memory) {
        return jobs[jobId];
    }

    function getOpenJobs() external view returns (uint256[] memory) {
        return openJobs;
    }
}
