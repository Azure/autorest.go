param(
    [string] $UseTypeSpecNext
)

[bool]$UseTypeSpecNext = $useTypeSpecNext -in 'true', '1', 'yes', 'y'

$ErrorActionPreference = 'Stop'

$root = (Resolve-Path "$PSScriptRoot/../..").Path.Replace('\', '/')

function invoke($command) {
    Write-Host "> $command"
    Invoke-Expression $command
    if ($LASTEXITCODE -ne 0) {
        Write-Host "Command failed: $command"
        exit $LASTEXITCODE
    }
}

function invoke-npx($package, $npxArgs) {
  $packageJson = Get-Content "$root/package.json" -Raw | ConvertFrom-Json
  $version = $packageJson.devDependencies.$package
  invoke "npx --yes $package@$version $npxArgs"
}

Push-Location $root
try {
  if ($UseTypeSpecNext) {
    invoke-npx '@azure-tools/typespec-bump-deps' "--use-peer-ranges packages/autorest.go/package.json"
  }

  invoke-npx "@microsoft/rush" 'update'
}
finally {
    Pop-Location
}
