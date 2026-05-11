$ErrorActionPreference = "Stop"

$repoRoot = (& git rev-parse --show-toplevel).Trim()
$backendDir = Join-Path $repoRoot "7/backend"

if (-not (Test-Path $backendDir)) {
    exit 0
}

$changed = & git diff --cached --name-only --diff-filter=ACMR
$backendChanges = $changed | Where-Object { $_ -like "7/backend/*" }
if (-not $backendChanges) {
    exit 0
}

if (-not (Get-Command gofmt -ErrorAction SilentlyContinue)) {
    Write-Error "gofmt not found in PATH"
    exit 1
}

$unformatted = & gofmt -l $backendDir
if ($unformatted) {
    Write-Error "gofmt found unformatted files:`n$unformatted"
    exit 1
}

if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    Write-Error "go not found in PATH"
    exit 1
}

Push-Location $backendDir
try {
    & go vet ./...
    if ($LASTEXITCODE -ne 0) {
        Write-Error "go vet failed"
        exit 1
    }
}
finally {
    Pop-Location
}

exit 0
