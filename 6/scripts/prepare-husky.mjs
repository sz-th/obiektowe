import { execSync } from "node:child_process";
import { existsSync } from "node:fs";
import { dirname, join, resolve } from "node:path";
import { fileURLToPath } from "node:url";

const pkgRoot = resolve(dirname(fileURLToPath(import.meta.url)), "..");
let dir = pkgRoot;

while (dir !== dirname(dir)) {
  if (existsSync(join(dir, ".git"))) {
    execSync(`npx husky "${join(pkgRoot, ".husky")}"`, {
      cwd: dir,
      stdio: "inherit",
    });
    break;
  }
  dir = dirname(dir);
}
