<#
  .SYNOPSIS
  Detects changed files in a PR and sets Azure DevOps output variables
  to control which CI jobs should run.

  .DESCRIPTION
  Uses Get-ChangedFiles from eng/common to determine which files changed,
  then evaluates exclusion rules mirroring the original per-pipeline path
  triggers. Sets output variables for each CI job.
#>

[CmdletBinding()]
param (
  [string]$SourceCommittish = "${env:SYSTEM_PULLREQUEST_SOURCECOMMITID}",
  [string]$TargetCommittish = ("origin/${env:SYSTEM_PULLREQUEST_TARGETBRANCH}" -replace "refs/heads/")
)

Set-StrictMode -Version 3

. (Join-Path $PSScriptRoot .. .. common scripts common.ps1)

$changedFiles = Get-ChangedFiles `
  -SourceCommittish $SourceCommittish `
  -TargetCommittish $TargetCommittish `
  -DiffFilterType ''

# On push/CI triggers there is no PR context, so Get-ChangedFiles returns
# empty.  Fall back to running every job in that case.
if (!$changedFiles) {
  Write-Host "No changed files detected (possibly a CI push trigger) — enabling all jobs."
  Write-Host "##vso[task.setvariable variable=run_autorest_go;isOutput=true]true"
  Write-Host "##vso[task.setvariable variable=run_typespec_go;isOutput=true]true"
  Write-Host "##vso[task.setvariable variable=run_gotest;isOutput=true]true"
  exit 0
}

# Exclusion sets mirroring the original pipeline path triggers.
# A job runs if ANY changed file is NOT covered by its exclude prefixes.
$excludeSets = @{
  run_autorest_go = @(
    "packages/autorest.gotest"
    "packages/typespec-go"
  )
  run_typespec_go = @(
    "packages/autorest.go"
    "packages/autorest.gotest"
    "swagger"
  )
  run_gotest = @(
    "packages/autorest.go"
    "packages/codegen.go"
    "packages/codemodel.go"
    "packages/naming.go"
    "packages/typespec-go"
  )
}

foreach ($varName in $excludeSets.Keys) {
  $excludes = $excludeSets[$varName]
  $shouldRun = $false

  foreach ($file in $changedFiles) {
    $excluded = $false
    foreach ($prefix in $excludes) {
      if ($file -like "$prefix/*" -or $file -eq $prefix) {
        $excluded = $true
        break
      }
    }
    if (-not $excluded) {
      $shouldRun = $true
      break
    }
  }

  $value = if ($shouldRun) { "true" } else { "false" }
  Write-Host "Setting $varName = $value"
  Write-Host "##vso[task.setvariable variable=$varName;isOutput=true]$value"
}
