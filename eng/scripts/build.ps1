param(
    [string] $BuildNumber,
    [string] $Output,
    [string] $BuildAlphaVersion
)

[bool]$BuildAlphaVersion = $BuildAlphaVersion -in 'true', '1', 'yes', 'y'

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

function updateVersion($suffix) {
  $version = (Get-Content "package.json" -Raw | ConvertFrom-Json).version.Split('-')[0]
  invoke "npm version $version$suffix --allow-same-version --no-workspaces-update"
}

if (!$Output) {
    $Output = "$root/artifacts"
}

$packagesFolder = New-Item "$Output/packages" -ItemType Directory -Force

Push-Location "$root/packages/autorest.go"
try {
    if ($BuildAlphaVersion) {
      updateVersion  "-alpha.$BuildNumber"
    }

    invoke "npm pack --pack-destination=$packagesFolder"

    exit $LASTEXITCODE
}
finally {
    Pop-Location
}
