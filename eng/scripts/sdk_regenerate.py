#!/usr/bin/env python

# --------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# --------------------------------------------------------------------------------------------
from typing import Dict, List, Optional
from pathlib import Path
import subprocess
from datetime import datetime
from subprocess import check_call, check_output, call
import argparse
import logging
import json
import re
import glob


def update_emitter_package(sdk_root: str, typespec_go_root: str, use_dev_package: bool):
    if use_dev_package:
        logging.info("Using dev package mode")
        
        # Find the package.json in typespec-go root
        package_json_path = Path(typespec_go_root) / "package.json"
        if not package_json_path.exists():
            logging.error(f"package.json not found at {package_json_path}")
            raise FileNotFoundError(f"package.json not found at {package_json_path}")
        
        # Use PowerShell script to generate emitter-package.json
        logging.info("Update emitter-package.json using New-EmitterPackageJson.ps1")
        try:
            check_call([
                "pwsh",
                "./eng/common/scripts/typespec/New-EmitterPackageJson.ps1",
                "-PackageJsonPath",
                str(package_json_path.absolute()),
                "-OutputDirectory",
                "eng"
            ], cwd=sdk_root, shell=True)
        except Exception as e:
            logging.error("Failed to run New-EmitterPackageJson.ps1")
            logging.error(e)
            raise
        
        # Find the typespec-go.tgz file
        typespec_go_tgz = None
        for item in Path(typespec_go_root).iterdir():
            if "typespec-go" in item.name and item.name.endswith(".tgz"):
                typespec_go_tgz = item
                break
        
        if not typespec_go_tgz:
            logging.error("Cannot find .tgz for typespec-go")
            raise FileNotFoundError("Cannot find .tgz for typespec-go")
        
        # Update emitter-package.json to use the dev package path
        emitter_package_path = Path(sdk_root) / "eng/emitter-package.json"
        with open(emitter_package_path, "r") as f:
            emitter_package = json.load(f)
        
        emitter_package["dependencies"]["@azure-tools/typespec-go"] = typespec_go_tgz.absolute().as_posix()
        
        with open(emitter_package_path, "w") as f:
            json.dump(emitter_package, f, indent=2)
        
        logging.info(f"Updated emitter-package.json to use typespec-go from \"{typespec_go_tgz.absolute()}\"")
        
        # Update emitter-package-lock.json
        logging.info("Update emitter-package-lock.json")
        try:
            check_call("tsp-client generate-lock-file", shell=True, cwd=sdk_root)
        except Exception as e:
            logging.error("Failed to update emitter-package-lock.json")
            logging.error(e)
            raise
    else:
        logging.info("Using released package mode")
        
        # Find the package.json in typespec-go root
        package_json_path = Path(typespec_go_root) / "package.json"
        if not package_json_path.exists():
            logging.error(f"package.json not found at {package_json_path}")
            raise FileNotFoundError(f"package.json not found at {package_json_path}")
        
        # Use tsp-client to generate config files with released package
        logging.info("Update emitter-package.json and emitter-package-lock.json using released package")
        try:
            check_call([
                "tsp-client", 
                "generate-config-files", 
                "--package-json", 
                str(package_json_path.absolute())
            ], cwd=sdk_root, shell=True)
        except Exception as e:
            logging.error("Failed to generate config files with tsp-client")
            logging.error(e)
            raise

def get_latest_commit_id() -> str:
    return (
        check_output(
            "git ls-remote https://github.com/Azure/azure-rest-api-specs.git HEAD | awk '{ print $1}'", shell=True
        )
        .decode("utf-8")
        .split("\n")[0]
        .strip()
    )


def get_typespec_go_commit_hash(typespec_go_root: str) -> str:
    """Get the current commit hash of the typespec-go repository."""
    try:
        return (
            check_output(
                "git rev-parse HEAD", shell=True, cwd=typespec_go_root
            )
            .decode("utf-8")
            .strip()
        )
    except Exception as e:
        logging.warning(f"Failed to get typespec-go commit hash: {e}")
        return "unknown"


def update_commit_id(file: Path, commit_id: str):
    with open(file, "r") as f:
        content = f.readlines()
    for idx in range(len(content)):
        if "commit:" in content[idx]:
            content[idx] = f"commit: {commit_id}\n"
            break
    with open(file, "w") as f:
        f.writelines(content)


def get_api_version_from_metadata(package_folder: Path) -> Optional[str]:
    """Extract API version from metadata.json file if it exists."""
    # Construct the metadata.json path based on the package folder structure
    # {package_folder}/testdata/_metadata.json
    metadata_path = package_folder / "testdata" / "_metadata.json"
    
    if metadata_path.exists():
        try:
            with open(metadata_path, "r") as f:
                metadata = json.load(f)
                api_version = metadata.get("apiVersion")
                if api_version:
                    logging.info(f"Found API version {api_version} in metadata.json for {package_folder.name}")
                    return api_version
        except (json.JSONDecodeError, FileNotFoundError) as e:
            logging.warning(f"Failed to read metadata.json for {package_folder.name}: {e}")
    
    return None


def get_api_version_from_client_files(package_folder: Path) -> Optional[str]:
    """Extract API version from client Go files by searching for 'Generated from API version' comment."""
    # Look for *_client.go files in the package folder
    client_files_pattern = str(package_folder / "*_client.go")
    client_files = glob.glob(client_files_pattern)
    
    api_version_pattern = re.compile(r"Generated from API version\s+([^\s,]+)")
    
    for client_file in client_files:
        try:
            with open(client_file, "r", encoding="utf-8") as f:
                content = f.read()
                match = api_version_pattern.search(content)
                if match:
                    api_version = match.group(1)
                    logging.info(f"Found API version {api_version} in {Path(client_file).name} for {package_folder.name}")
                    return api_version
        except (FileNotFoundError, UnicodeDecodeError) as e:
            logging.warning(f"Failed to read client file {client_file}: {e}")
    
    return None


def get_api_version(package_folder: Path) -> Optional[str]:
    """Get API version for a package, first trying metadata.json, then client files."""
    # First, try to get from metadata.json
    api_version = get_api_version_from_metadata(package_folder)
    
    if api_version:
        return api_version
    
    # If not found in metadata, try client files
    api_version = get_api_version_from_client_files(package_folder)
    
    if not api_version:
        logging.warning(f"Could not find API version for {package_folder.name}")
    
    return api_version

def regenerate_sdk(use_latest_spec: bool, service_filter: str, sdk_root: str, typespec_go_root: str) -> Dict[str, List[str]]:
    result = {
        "succeed_to_regenerate": [], 
        "fail_to_regenerate": [], 
        "not_found_api_version": [], 
        "time_to_regenerate": str(datetime.now()),
        "typespec_go_commit_hash": get_typespec_go_commit_hash(typespec_go_root)
    }
    # get all tsp-location.yaml
    commit_id = get_latest_commit_id()
    sdk_resourcemanager_path = Path(sdk_root) / "sdk" / "resourcemanager"
    for item in sdk_resourcemanager_path.rglob("tsp-location.yaml"):
        package_folder = item.parent
        if len(service_filter) > 0 and re.match(service_filter, package_folder.name) is None:
            continue
        logging.info(f"Regenerating {package_folder.name}...")
        if use_latest_spec:
            logging.info("Using latest spec")
            update_commit_id(item, commit_id)
        try:
            # Get API version for this package
            api_version = get_api_version(package_folder)
            
            # Build the tsp-client command with optional API version
            tsp_command = "tsp-client update"
            if api_version:
                tsp_command += f" --emitter-options api-version={api_version}"
                logging.info(f"Using API version {api_version} for {package_folder.name}")
            else:
                logging.info(f"No API version specified for {package_folder.name}, using default behavior")
                result["not_found_api_version"].append(package_folder.name)

            # Use subprocess.run for better control over output
            proc_result = subprocess.run(
                tsp_command, 
                shell=True, 
                cwd=str(package_folder),
                capture_output=True,
                text=True,
                check=True
            )
            
            # Log the output for progress tracking
            if proc_result.stdout:
                logging.info(f"Output for {package_folder.name}:")
                for line in proc_result.stdout.split('\n'):
                    if line.strip():
                        logging.info(f"  {line}")
            
            if proc_result.stderr:
                logging.warning(f"Stderr for {package_folder.name}:")
                for line in proc_result.stderr.split('\n'):
                    if line.strip():
                        logging.warning(f"  {line}")
                        
            # Check for errors in output
            output_lines = proc_result.stdout.split('\n') if proc_result.stdout else []
            errors = [line for line in output_lines if "- error " in line.lower()]
            if errors:
                raise Exception("Errors found in output:\n" + "\n".join(errors))
                
        except subprocess.CalledProcessError as e:
            logging.error(f"Failed to regenerate {package_folder.name}")
            logging.error(f"Command failed with exit code {e.returncode}")
            if e.stdout:
                logging.error(f"Stdout:\n{e.stdout}")
            if e.stderr:
                logging.error(f"Stderr:\n{e.stderr}")
            result["fail_to_regenerate"].append(package_folder.name)
        except Exception as e:
            logging.error(f"Failed to regenerate {package_folder.name}")
            logging.error(f"Error: {str(e)}")
            result["fail_to_regenerate"].append(package_folder.name)
        else:
            logging.info(f"Successfully regenerated {package_folder.name}")
            result["succeed_to_regenerate"].append(package_folder.name)
            
    result["succeed_to_regenerate"].sort()
    result["fail_to_regenerate"].sort()
    result["not_found_api_version"].sort()
    return result


def checkout_branch(branch: str, sync_main: bool = False):
    try:
        check_call(f"git fetch azure-sdk {branch}", shell=True)
        check_call(f"git checkout {branch}", shell=True)
        if sync_main:
            logging.info(f"sync {branch} with main branch")
            call(f"git pull azure-sdk main", shell=True)
    except Exception:
        check_call(f"git checkout -b {branch}", shell=True)


def prepare_branch(typespec_go_branch: str):
    check_call("git remote add azure-sdk https://github.com/azure-sdk/azure-sdk-for-go.git", shell=True)

    if typespec_go_branch == "main":
        checkout_branch("typespec-go-main", typespec_go_branch == "main")
    else:
        checkout_branch(f"typespec-go-{typespec_go_branch}")

def git_add():
    check_call("git add .", shell=True)


def main(sdk_root: str, typespec_go_root: str, typespec_go_branch: str, use_latest_spec: bool, service_filter: str, use_dev_package: bool):
    # Configure logging for better pipeline visibility
    logging.basicConfig(
        level=logging.INFO,
        format='%(asctime)s - %(levelname)s - %(message)s',
        handlers=[
            logging.StreamHandler(),  # Console output for pipeline
        ]
    )
    
    prepare_branch(typespec_go_branch)
    update_emitter_package(sdk_root, typespec_go_root, use_dev_package)
    result = regenerate_sdk(use_latest_spec, service_filter, sdk_root, typespec_go_root)
    with open("regenerate-sdk-result.json", "w") as f:
        json.dump(result, f, indent=2)
    git_add()


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="SDK regeneration")

    parser.add_argument(
        "--sdk-root",
        help="SDK repo root folder",
        type=str,
    )

    parser.add_argument(
        "--typespec-go-root",
        help="typespec-go repo root folder",
        type=str,
    )

    parser.add_argument(
        "--typespec-go-branch",
        help="branch of typespec-go repo",
        type=str,
    )

    parser.add_argument(
        "--use-latest-spec",
        help="Whether to use the latest spec",
        type=lambda x: x.lower() == 'true',
        default=False,
    )

    parser.add_argument(
        "--service-filter",
        help="An regex filter to specify which service to regenerate. If not specified, all services will be regenerated.",
        type=str,
    )

    parser.add_argument(
        "--use-dev-package",
        help="Whether to use dev package or released package",
        type=lambda x: x.lower() == 'true',
        default=False,
    )

    args = parser.parse_args()

    main(args.sdk_root, args.typespec_go_root, args.typespec_go_branch, args.use_latest_spec, args.service_filter, args.use_dev_package)