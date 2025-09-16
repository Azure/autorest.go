# SDK Regeneration Pipeline Guideline

This document provides a guideline about the SDK regeneration pipeline for TypeSpec Go emitter developers.

## Pipeline Overview

> **Pipeline Link**: [SDK Regeneration Pipeline](https://dev.azure.com/azure-sdk/internal/_build?definitionId=7860)

The SDK regeneration pipeline regenerates Azure SDK for Go packages using any branch of TypeSpec Go emitter with the same API version as the released SDK (which to make sure SDK changes are due to TypeSpec or Go emitter release, not API version). This pipeline validates emitter changes and ensures that SDK packages stay up-to-date with TypeSpec changes and emitter changes.

### Limitations

⚠️ **Management Plane Only**: This pipeline is currently designed exclusively for **management plane (resource manager) SDKs**. It does **not** support:
- Data plane SDKs
- Client libraries outside the `sdk/resourcemanager` directory
- Custom SDK implementations

### Pipeline Process
The SDK regeneration pipeline automates the process of updating Azure SDK packages with the latest TypeSpec Go emitter changes. Here are the key stages:
> 💡For detailed implementation, see the complete pipeline definition in [`/eng/pipelines/sdk-regenerate.yml`](https://github.com/Azure/autorest.go/blob/main/eng/pipelines/sdk-regenerate.yml) and the regeneration script in [`/eng/scripts/sdk_regenerate.py`](https://github.com/Azure/autorest.go/blob/main/eng/scripts/sdk_regenerate.py)

#### 1. Setup & Build
- Sets up environment (Node.js, Go, pnpm, tsp-client)
- Builds the current TypeSpec Go emitter from current branch
- Update emitter version of Azure Go SDK `emitter-package.json`

#### 2. Package Discovery
- Scans `sdk/resourcemanager` for packages with `tsp-location.yaml`
- Applies service filtering (if specified)
- Extracts **original API versions** from existing packages (`_metadata.json` or client files)

#### 3. SDK Generation
- Runs `tsp-client update` for each package with original API version
- Tracks success/failure for each package

#### 4. Results & PR Creation
- Generates results report (`regenerate-sdk-result.json`)
- Creates draft PR in azure-sdk-for-go with all changes, PR title: `[Automation] Regenerate SDK based on typespec-go branch {branch-name}`

### Pipeline Parameters

#### Required Parameters

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `UseLatestSpec` | boolean | `false` | Whether to use the latest API specifications from [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs) or the original commit of `tsp-location.yml` |
| `ServiceFilter` | string | `.*` | Regex pattern to filter which services to regenerate. Matches against the service package name (e.g., `armcompute`, `armstorage`) |

#### Usage Examples

##### Generate All SDKs with Latest Specs (Default)
```yaml
UseLatestSpec: true
ServiceFilter: '.*'
```

##### Generate Specific Services
```yaml
UseLatestSpec: false
ServiceFilter: 'armcompute|armstorage|armchaos'
```

##### Generate One Service
```yaml
UseLatestSpec: true
ServiceFilter: 'armcompute'
```

## Pipeline Usage
Currently, this pipeline only supports manually triggered when you need to verify emitter changes.
1. Navigate to the [pipeline](https://dev.azure.com/azure-sdk/internal/_build?definitionId=7860)
2. Select which branch of emitter you want to validate
3. Configure parameters:
   - Set `UseLatestSpec` to `true` for latest API specs, `false` for the original commit of `tsp-location.yml`
   - Set `ServiceFilter` for specific services (use `.*` for all)
4. Click "Run"

## Pipeline Result
You could find the generated SDK pull request link from pipeline logs
<img width="923" height="610" alt="image" src="https://github.com/user-attachments/assets/a06cfe51-1cff-45b2-945f-a4398069cee0" />

### PR Structure
- **Title**: `[Automation] Regenerate SDK based on typespec-go branch {branch-name}`
- **Status**: Opens as draft PR for review
- **Base Branch**: `main` branch of azure-sdk-for-go
- **Branch Name**: `typespec-go-regenerate-{branch-name}`

### PR Contents
- Updated SDK packages with latest TypeSpec Go emitter and original spec api version
- Regeneration result: regenerate-sdk-result.json
```json
{
  "succeed_to_regenerate": ["package1", "package2"],
  "fail_to_regenerate": ["package3"]
}
```

## Validate and Refresh SDK

### Development Workflow

```mermaid
graph TB
    A[TypeSpec Go Emitter Changes] --> B[Trigger Regeneration Pipeline]
    B --> C[Review Generated SDK PR]
    C --> D{Purpose?}
    D -->|Dev Validation| E[Merge Emitter PR]
    D -->|Production Release| F[Merge Emitter PR]
    E --> G[Close Regeneration PR]
    F --> H[Release Emitter Version]
    H --> I[Merge Regeneration PR]
    I --> J[SDK Baseline Updated]
    
    K[TypeSpec/Emitter Version Released] --> L[Trigger Refresh Regeneration]
    L --> M[Review Refresh PR]
    M --> N[Merge Refresh PR]
    N --> O[SDK Baseline Refreshed]
```

#### For Development Validation
1. **Develop & Test**: Complete TypeSpec Go emitter development and testing on feature branch
2. **Trigger Pipeline**: Trigger regeneration pipeline for your feature branch
3. **Review Changes**: Review the generated PR to ensure all changes are due to TypeSpec or Go emitter release
4. **Merge Emitter**: Merge TypeSpec Go emitter PR into main branch
5. **Close Regeneration PR**: Close/abandon the regeneration PR (it was only for validation)

#### For Regular SDK Refresh
After each TypeSpec or Go emitter version release:
1. **Trigger Pipeline**: Trigger regeneration pipeline for `main` branch
2. **Review Changes**: Review the generated PR to ensure all changes are due to TypeSpec or Go emitter release
3. **Merge Regeneration PR**: Merge the refresh PR to keep SDK up-to-date

### Quality Gates
- All pipelines must be pass
- All SDK code changes are made by either TypeSpec changes or Go emitter changes
- Module versions must not be changed
- API versions must not be changed

## Related Resources

- [TypeSpec Go Emitter Repository](https://github.com/Azure/autorest.go)
- [Azure SDK for Go Repository](https://github.com/Azure/azure-sdk-for-go)
- [Azure Rest API Spec Repository](https://github.com/Azure/azure-rest-api-specs)