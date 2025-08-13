#!/usr/bin/env node

/**
 * Simple test script for the changelog date validator
 */

import { execSync } from 'child_process';
import fs from 'fs';
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const testDir = '/tmp/changelog-validation-tests';
const validatorScript = path.join(__dirname, 'validate-changelog-dates.js');

console.log('Testing CHANGELOG.md validation script...\n');

// Clean up test directory if it exists
if (fs.existsSync(testDir)) {
  fs.rmSync(testDir, { recursive: true, force: true });
}
fs.mkdirSync(testDir, { recursive: true });

// Test 1: Valid changelog should pass
const validChangelog = `# Release History

## 0.2.0 (unreleased)

### Other Changes

* Upcoming release.

## 0.1.0 (2025-07-15)

### Features Added

* Initial release.
`;

fs.writeFileSync(path.join(testDir, 'valid-changelog.md'), validChangelog);

try {
  execSync(`node "${validatorScript}" "${path.join(testDir, 'valid-changelog.md')}"`, { 
    stdio: 'pipe' 
  });
  console.log('✅ Test 1 passed: Valid changelog accepted');
} catch (error) {
  console.error('❌ Test 1 failed: Valid changelog rejected');
  console.error(error.stdout?.toString() || error.message);
  process.exit(1);
}

// Test 2: Invalid changelog should fail
const invalidChangelog = `# Release History

## 0.2.0 (2030-12-31)

### Other Changes

* Future release.

## 0.1.0 (invalid-date)

### Features Added

* Invalid date format.
`;

fs.writeFileSync(path.join(testDir, 'invalid-changelog.md'), invalidChangelog);

try {
  execSync(`node "${validatorScript}" "${path.join(testDir, 'invalid-changelog.md')}"`, { 
    stdio: 'pipe' 
  });
  console.error('❌ Test 2 failed: Invalid changelog accepted');
  process.exit(1);
} catch (error) {
  console.log('✅ Test 2 passed: Invalid changelog rejected');
}

// Cleanup
fs.rmSync(testDir, { recursive: true, force: true });

console.log('\n🎉 All tests passed! Changelog validation script is working correctly.');