import { expect } from "chai";
import { network } from "hardhat";

const { ethers } = await network.connect();

/**
 * NimbusX Smart Contract Integration Tests
 *
 * Tests the full job lifecycle:
 *   Provider registers → Client posts job → Scheduler assigns → Provider completes → Funds released
 */
describe("NimbusX Contracts", function () {
  // ── Helpers ────────────────────────────────────────────────────────────────

  async function deployAll() {
    const [owner, scheduler, provider, client] = await ethers.getSigners();

    // Deploy in dependency order
    const Escrow = await ethers.getContractFactory("Escrow");
    const escrow = await Escrow.deploy();

    const Registry = await ethers.getContractFactory("ProviderRegistry");
    const registry = await Registry.deploy();

    const Marketplace = await ethers.getContractFactory("Marketplace");
    const marketplace = await Marketplace.deploy(
      await registry.getAddress(),
      await escrow.getAddress(),
      scheduler.address
    );

    // Wire escrow to marketplace
    await escrow.setMarketplace(await marketplace.getAddress());

    return { escrow, registry, marketplace, owner, scheduler, provider, client };
  }

  // ── ProviderRegistry ───────────────────────────────────────────────────────

  describe("ProviderRegistry", function () {
    it("Should allow a provider to register", async function () {
      const { registry, provider } = await deployAll();

      await registry.connect(provider).registerProvider(8, 16384, ethers.parseEther("0.01"));
      const p = await registry.getProvider(provider.address);

      expect(p.active).to.be.true;
      expect(p.cpu).to.equal(8n);
      expect(p.memoryMB).to.equal(16384n);
      expect(p.reputation).to.equal(100n);
    });

    it("Should prevent double registration", async function () {
      const { registry, provider } = await deployAll();

      await registry.connect(provider).registerProvider(4, 8192, ethers.parseEther("0.005"));
      await expect(
        registry.connect(provider).registerProvider(4, 8192, ethers.parseEther("0.005"))
      ).to.be.revertedWith("Already registered");
    });

    it("Should allow a provider to deregister", async function () {
      const { registry, provider } = await deployAll();

      await registry.connect(provider).registerProvider(4, 8192, ethers.parseEther("0.005"));
      await registry.connect(provider).deregisterProvider();

      const p = await registry.getProvider(provider.address);
      expect(p.active).to.be.false;
    });

    it("Should allow owner to update reputation", async function () {
      const { registry, owner, provider } = await deployAll();

      await registry.connect(provider).registerProvider(4, 8192, ethers.parseEther("0.005"));
      await registry.connect(owner).updateReputation(provider.address, 75n);

      const p = await registry.getProvider(provider.address);
      expect(p.reputation).to.equal(75n);
    });
  });

  // ── Escrow ─────────────────────────────────────────────────────────────────

  describe("Escrow", function () {
    it("Should reject setMarketplace if already set", async function () {
      const { escrow, marketplace } = await deployAll();
      await expect(
        escrow.setMarketplace(await marketplace.getAddress())
      ).to.be.revertedWith("Already set");
    });
  });

  // ── Full Job Lifecycle ─────────────────────────────────────────────────────

  describe("Job Lifecycle", function () {
    it("Should complete the full happy path: post → assign → start → complete", async function () {
      const { registry, marketplace, escrow, scheduler, provider, client } = await deployAll();

      // 1. Provider registers
      await registry.connect(provider).registerProvider(8, 16384, ethers.parseEther("0.01"));

      // 2. Client posts job with 1 ETH escrow
      const payment = ethers.parseEther("1");
      const tx = await marketplace.connect(client).postJob(4, 8192, 3600, ethers.parseEther("0.02"), {
        value: payment,
      });
      const receipt = await tx.wait();
      const jobId = 0n;

      // Verify escrow holds the funds
      const entry = await escrow.getEntry(jobId);
      expect(entry.amount).to.equal(payment);

      // 3. Scheduler assigns provider
      await marketplace.connect(scheduler).assignJob(jobId, provider.address);

      // 4. Provider starts
      await marketplace.connect(provider).startJob(jobId);

      // 5. Provider completes with result hash
      const resultHash = ethers.keccak256(ethers.toUtf8Bytes("output-data"));
      const providerBefore = await ethers.provider.getBalance(provider.address);

      await marketplace.connect(provider).completeJob(jobId, resultHash);

      // Funds should be released to provider
      const providerAfter = await ethers.provider.getBalance(provider.address);
      expect(providerAfter).to.be.greaterThan(providerBefore);

      // Job status = Completed (3)
      const job = await marketplace.getJob(jobId);
      expect(job.status).to.equal(3n);
      expect(job.resultHash).to.equal(resultHash);
    });

    it("Should refund client and slash reputation on job failure", async function () {
      const { registry, marketplace, scheduler, provider, client } = await deployAll();

      await registry.connect(provider).registerProvider(8, 16384, ethers.parseEther("0.01"));

      const payment = ethers.parseEther("1");
      await marketplace.connect(client).postJob(4, 8192, 3600, ethers.parseEther("0.02"), {
        value: payment,
      });
      const jobId = 0n;

      await marketplace.connect(scheduler).assignJob(jobId, provider.address);
      await marketplace.connect(provider).startJob(jobId);

      const clientBefore = await ethers.provider.getBalance(client.address);
      await marketplace.connect(scheduler).failJob(jobId);
      const clientAfter = await ethers.provider.getBalance(client.address);

      // Client gets refunded
      expect(clientAfter).to.be.greaterThan(clientBefore);

      // Provider reputation is slashed
      const p = await registry.getProvider(provider.address);
      expect(p.reputation).to.equal(90n); // 100 - 10
    });

    it("Should allow client to cancel an open job", async function () {
      const { marketplace, client } = await deployAll();

      const payment = ethers.parseEther("0.5");
      await marketplace.connect(client).postJob(2, 4096, 1800, ethers.parseEther("0.01"), {
        value: payment,
      });
      const jobId = 0n;

      const clientBefore = await ethers.provider.getBalance(client.address);
      await marketplace.connect(client).cancelJob(jobId);
      const clientAfter = await ethers.provider.getBalance(client.address);

      expect(clientAfter).to.be.greaterThan(clientBefore);

      const job = await marketplace.getJob(jobId);
      expect(job.status).to.equal(5n); // Cancelled
    });
  });
});
