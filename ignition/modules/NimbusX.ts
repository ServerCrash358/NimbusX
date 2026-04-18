import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

/**
 * NimbusX deployment module.
 *
 * Deployment order matters because of contract dependencies:
 *   1. Escrow           (no deps)
 *   2. ProviderRegistry (no deps)
 *   3. Marketplace      (needs Escrow + ProviderRegistry addresses + scheduler address)
 *   4. Escrow.setMarketplace(marketplace address)
 *
 * The scheduler address below defaults to the first Hardhat test account for local runs.
 * Override it via module parameters for production deployments.
 */
const NimbusXModule = buildModule("NimbusXModule", (m) => {
  // Configurable scheduler address – defaults to account[0] for local testing
  const schedulerAddress = m.getParameter(
    "schedulerAddress",
    "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266" // hardhat account[0]
  );

  // 1. Deploy Escrow first (no constructor args)
  const escrow = m.contract("Escrow");

  // 2. Deploy ProviderRegistry (no constructor args)
  const registry = m.contract("ProviderRegistry");

  // 3. Deploy Marketplace wired to Escrow + Registry + Scheduler
  const marketplace = m.contract("Marketplace", [registry, escrow, schedulerAddress]);

  // 4. Wire Escrow back to Marketplace (avoids circular constructor dependency)
  m.call(escrow, "setMarketplace", [marketplace]);

  return { escrow, registry, marketplace };
});

export default NimbusXModule;
