import { execSync } from "node:child_process";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";

const root = join(dirname(fileURLToPath(import.meta.url)), "..");
const backend = join(root, "backend");
const files = process.argv.slice(2);

for (const file of files) {
  execSync(`gofmt -w "${file}"`, { cwd: root, stdio: "inherit" });
}

if (files.length > 0) {
  execSync("go vet ./...", { cwd: backend, stdio: "inherit" });
}
