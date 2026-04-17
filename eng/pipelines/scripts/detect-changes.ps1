<#
  .SYNOPSIS
  Detects changed files in a PR and sets Azure DevOps output variables
  to control which CI jobs should run.

  This script stands in the place of what was originally 3 separate path triggers across 3 pipelines.
  By utilizing this script to set output variables, we can consolidate into a single pipeline with conditional jobs,
  while maintaining the same behavior of only running jobs when relevant files have changed.

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
# empty.  Fall back to HEAD~1..HEAD so job flags still reflect the actual
# changed files
if (!$changedFiles) {
  Write-Host "No PR context — falling back to HEAD~1..HEAD diff for push trigger."
  $changedFiles = git -c core.quotepath=off diff HEAD~1..HEAD --name-only
  if (!$changedFiles) {
    Write-Host "Still no changed files detected — enabling all jobs as safety fallback."
    Write-Host "##vso[task.setvariable variable=run_autorest_go;isOutput=true]true"
    Write-Host "##vso[task.setvariable variable=run_typespec_go;isOutput=true]true"
    Write-Host "##vso[task.setvariable variable=run_gotest;isOutput=true]true"
    exit 0
  }
  Write-Host "Push trigger diff files:"
  foreach ($file in $changedFiles) {
    Write-Host "    $file"
  }
}

# Exclusion sets governing which jobs run based on which files are changed.
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
