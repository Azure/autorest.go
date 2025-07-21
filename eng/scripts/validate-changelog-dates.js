#!/usr/bin/env node

/**
 * Script to validate release dates in CHANGELOG.md files
 * Ensures that:
 * 1. Release dates follow the format YYYY-MM-DD
 * 2. Release dates are not in the future (allowing some buffer for timezone differences)
 * 3. "unreleased" entries are allowed
 */

import fs from 'fs';
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

function validateChangelogDates(changelogPath) {
  console.log(`Validating changelog at: ${changelogPath}`);
  
  if (!fs.existsSync(changelogPath)) {
    console.error(`Error: CHANGELOG.md not found at ${changelogPath}`);
    process.exit(1);
  }

  const content = fs.readFileSync(changelogPath, 'utf8');
  const lines = content.split('\n');
  
  // Regex to match version headers with dates: ## x.y.z (YYYY-MM-DD) or ## x.y.z (unreleased)
  const versionHeaderRegex = /^##\s+(\d+\.\d+\.\d+(?:\.\d+)?)\s+\((.+)\)\s*$/;
  
  let hasErrors = false;
  const today = new Date();
  // Add 1 day buffer for timezone differences
  const maxAllowedDate = new Date(today.getTime() + 24 * 60 * 60 * 1000);
  
  for (let i = 0; i < lines.length; i++) {
    const line = lines[i].trim();
    const match = line.match(versionHeaderRegex);
    
    if (match) {
      const version = match[1];
      const dateStr = match[2];
      const lineNumber = i + 1;
      
      console.log(`Found release entry: ${version} (${dateStr}) at line ${lineNumber}`);
      
      // Skip "unreleased" entries
      if (dateStr.toLowerCase() === 'unreleased') {
        console.log(`  ✓ Version ${version} is marked as unreleased - OK`);
        continue;
      }
      
      // Validate date format (YYYY-MM-DD)
      const dateFormatRegex = /^\d{4}-\d{2}-\d{2}$/;
      if (!dateFormatRegex.test(dateStr)) {
        console.error(`  ✗ Error: Invalid date format "${dateStr}" for version ${version} at line ${lineNumber}`);
        console.error(`    Expected format: YYYY-MM-DD`);
        hasErrors = true;
        continue;
      }
      
      // Parse and validate the date
      const releaseDate = new Date(dateStr + 'T00:00:00.000Z'); // Parse as UTC to avoid timezone issues
      
      if (isNaN(releaseDate.getTime())) {
        console.error(`  ✗ Error: Invalid date "${dateStr}" for version ${version} at line ${lineNumber}`);
        hasErrors = true;
        continue;
      }
      
      // Additional validation for reasonable date values
      const currentYear = new Date().getFullYear();
      const releaseYear = releaseDate.getFullYear();
      
      if (releaseYear < 2020 || releaseYear > currentYear + 2) {
        console.error(`  ✗ Error: Unreasonable release date "${dateStr}" for version ${version} at line ${lineNumber}`);
        console.error(`    Release year ${releaseYear} seems outside reasonable range (2020-${currentYear + 2})`);
        hasErrors = true;
        continue;
      }
      
      // Check if date is in the future
      if (releaseDate > maxAllowedDate) {
        console.error(`  ✗ Error: Future release date "${dateStr}" for version ${version} at line ${lineNumber}`);
        console.error(`    Release dates cannot be in the future. Current date: ${today.toISOString().split('T')[0]}`);
        hasErrors = true;
        continue;
      }
      
      console.log(`  ✓ Version ${version} release date ${dateStr} is valid`);
    }
  }
  
  if (hasErrors) {
    console.error('\n❌ CHANGELOG validation failed with errors');
    process.exit(1);
  } else {
    console.log('\n✅ CHANGELOG validation passed');
  }
}

// Get the changelog path from command line argument or default to typespec-go
const changelogPath = process.argv[2] || path.join(__dirname, '../../packages/typespec-go/CHANGELOG.md');
validateChangelogDates(changelogPath);