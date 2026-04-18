import hre from "hardhat";

/**
 * Deploys all NimbusX contracts to the configured network.
 *
 * Deployment order:
 *   1. Escrow
 *   2. ProviderRegistry
 *   3. Marketplace (wired to Escrow + Registry)
 *   4. Escrow.setMarketplace(marketplace)
 */
async function main() {
  const [deployer] = await hre.ethers.getSigners();

  console.log("Deploying NimbusX contracts with:", deployer.address);
  console.log("Network:", hre.network.name);
  console.log("");

  // 1. Deploy Escrow
  const Escrow = await hre.ethers.getContractFactory("Escrow");
  const escrow = await Escrow.deploy();
  await escrow.waitForDeployment();
  const escrowAddr = await escrow.getAddress();
  console.log("✅ Escrow deployed to:           ", escrowAddr);

  // 2. Deploy ProviderRegistry
  const Registry = await hre.ethers.getContractFactory("ProviderRegistry");
  const registry = await Registry.deploy();
  await registry.waitForDeployment();
  const registryAddr = await registry.getAddress();
  console.log("✅ ProviderRegistry deployed to: ", registryAddr);

  // 3. Deploy Marketplace
  //    Scheduler address defaults to deployer for testing; set a real scheduler for production.
  const schedulerAddr = deployer.address;
  const Marketplace = await hre.ethers.getContractFactory("Marketplace");
  const marketplace = await Marketplace.deploy(registryAddr, escrowAddr, schedulerAddr);
  await marketplace.waitForDeployment();
  const marketplaceAddr = await marketplace.getAddress();
  console.log("✅ Marketplace deployed to:      ", marketplaceAddr);

  // 4. Wire Escrow → Marketplace
  await escrow.setMarketplace(marketplaceAddr);
  console.log("✅ Escrow wired to Marketplace");

  // 5. Wire Registry → Marketplace
  await registry.setMarketplace(marketplaceAddr);
  console.log("✅ ProviderRegistry wired to Marketplace");

  console.log("\n── Deployment Summary ──────────────────────────");
  console.log(`Escrow:           ${escrowAddr}`);
  console.log(`ProviderRegistry: ${registryAddr}`);
  console.log(`Marketplace:      ${marketplaceAddr}`);
  console.log(`Scheduler:        ${schedulerAddr}`);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});