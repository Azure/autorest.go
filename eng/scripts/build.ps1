param(
    [string] $BuildNumber,
    [string] $Output,
    [string] $BuildAlphaVersion
)

[bool]$BuildAlphaVersion = $BuildAlphaVersion -in 'true', '1', 'yes', 'y'

$ErrorActionPreference = 'Stop'

$root = (Resolve-Path "$PSScriptRoot/../..").Path.Replace('\', '/')

if (!$Output) {
  $Output = "$root/artifacts"
}

$packagesFolder = New-Item "$Output/packages" -ItemType Directory -Force

function invoke($command) {
  Write-Host "> $command"
  Invoke-Expression $command
  if ($LASTEXITCODE -ne 0) {
    Write-Host "Command failed: $command"
    exit $LASTEXITCODE
  }
}

function updatePackage($name, $path) {
  Push-Location $path
  try {
    $version = (Get-Content "./package.json" -Raw | ConvertFrom-Json).version

    if ($BuildAlphaVersion) {
      $version = "$($version.Split('-')[0])-alpha.$BuildNumber"
      invoke "npm version $version --allow-same-version --no-workspaces-update"
      if ($LASTEXITCODE) { exit $LASTEXITCODE }
    }

    invoke "npm pack --pack-destination=$packagesFolder"
    if ($LASTEXITCODE) { exit $LASTEXITCODE }
    
    $packageMatrix[$name] = $version
  }
  finally {
      Pop-Location
  }
}

$packageMatrix = @{}

updatePackage "generator" "$root/packages/autorest.go"

$packageMatrix | ConvertTo-Json | Set-Content $output/package-versions.json

exit 0
