#!/usr/bin/env python

# --------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# --------------------------------------------------------------------------------------------
from typing import Any, Dict, List, Optional
from pathlib import Path
import subprocess
from datetime import datetime
from subprocess import check_call, check_output, call
import argparse
import logging
import json
import os
import re
import glob
import urllib.request


def get_latest_typespec_go_package_info():
    """Get the latest version and dependencies of @azure-tools/typespec-go from npm registry."""
    try:
        logging.info("Fetching latest @azure-tools/typespec-go info from npm registry")
        
        # Get package info from npm registry
        url = "https://registry.npmjs.org/@azure-tools/typespec-go/latest"
        with urllib.request.urlopen(url) as response:
            package_info = json.loads(response.read().decode())
        
        version = package_info.get("version")
        dev_dependencies = package_info.get("devDependencies", {})
        
        logging.info(f"Latest @azure-tools/typespec-go version: {version}")
        
        return {
            "version": version,
            "devDependencies": dev_dependencies
        }
    except Exception as e:
        logging.error(f"Failed to fetch latest typespec-go info: {e}")
        raise


def update_dev_dependencies(emitter_package: dict, source_deps: dict):
    """Update devDependencies in emitter_package with versions from source_deps."""
    if "devDependencies" not in emitter_package:
        return
    for package_name in emitter_package["devDependencies"].keys():
        if package_name in source_deps:
            emitter_package["devDependencies"][package_name] = source_deps[package_name]
            logging.info(f"Updated {package_name} to version {source_deps[package_name]}")
        else:
            logging.info(f"Package {package_name} not found in dependencies, keeping existing version")


def update_emitter_package(sdk_root: str, typespec_go_root: str, use_dev_package: bool):
    # Load existing emitter-package.json
    emitter_package_path = Path(sdk_root) / "eng/emitter-package.json"
    if not emitter_package_path.exists():
        logging.error(f"emitter-package.json not found at {emitter_package_path}")
        raise FileNotFoundError(f"emitter-package.json not found at {emitter_package_path}")
    logging.info("Loading existing emitter-package.json")
    with open(emitter_package_path, "r", encoding="utf-8") as f:
        emitter_package = json.load(f)
    
    if use_dev_package:
        logging.info("Using dev package mode")
        
        # Find the package.json in typespec-go root
        package_json_path = Path(typespec_go_root) / "package.json"
        if not package_json_path.exists():
            logging.error(f"package.json not found at {package_json_path}")
            raise FileNotFoundError(f"package.json not found at {package_json_path}")
        
        # Load package.json to get dependency versions
        logging.info("Reading package.json to get dependency versions")
        with open(package_json_path, "r", encoding="utf-8") as f:
            package_json = json.load(f)
        dev_deps = package_json.get("devDependencies", {})

        # Update devDependencies in emitter_package
        update_dev_dependencies(emitter_package, dev_deps)

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
        emitter_package["dependencies"]["@azure-tools/typespec-go"] = typespec_go_tgz.absolute().as_posix()
    else:
        logging.info("Using released package mode")

        # Find the package.json in recent released typespec-go
        package_info = get_latest_typespec_go_package_info()

        # Update emitter-package.json to use the released package version
        if "dependencies" not in emitter_package:
            emitter_package["dependencies"] = {}
        emitter_package["dependencies"]["@azure-tools/typespec-go"] = package_info["version"]
        logging.info(f"Updated @azure-tools/typespec-go to version {package_info['version']}")

        # Update devDependencies in emitter_package
        dev_deps = package_info["devDependencies"]
        update_dev_dependencies(emitter_package, dev_deps)
    
    # Print the complete emitter_package before writing
    logging.info("Complete emitter-package.json content:")
    logging.info(json.dumps(emitter_package, indent=2))
    
    # Write the updated emitter-package.json
    with open(emitter_package_path, "w", encoding="utf-8") as f:
        json.dump(emitter_package, f, indent=2)
    
    # Update emitter-package-lock.json
    logging.info("Update emitter-package-lock.json")
    try:
        check_call(["tsp-client", "generate-lock-file"], cwd=sdk_root)
    except Exception as e:
        logging.error("Failed to update emitter-package-lock.json")
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


def get_module_name(package_folder: Path) -> Optional[str]:
    """Read the module name from go.mod in the package folder."""
    go_mod_path = package_folder / "go.mod"
    if not go_mod_path.exists():
        return None
    try:
        with open(go_mod_path, "r", encoding="utf-8") as f:
            for line in f:
                stripped = line.strip()
                if stripped.startswith("module "):
                    return stripped[len("module "):].strip()
    except FileNotFoundError as e:
        logging.warning(f"Failed to read go.mod for {package_folder.name}: {e}")
    return None


def restore_module_name(package_folder: Path, original_module: str) -> Optional[str]:
    """Regen drops the module suffix; restore the original suffixed path. Return it if restored."""
    if not original_module:
        return None
    base_module = re.sub(r"/v\d+$", "", original_module)
    suffix = re.compile(r"(" + re.escape(base_module) + r")/v\d+")
    bumped_module = None
    changed_files = check_output(
        ["git", "diff", "--name-only", "--", str(package_folder)], text=True
    ).splitlines()
    for rel in changed_files:
        file_path = Path(rel)
        if not file_path.is_file():
            continue
        try:
            old = check_output(["git", "show", f"HEAD:{rel}"], text=True)
            new = file_path.read_text(encoding="utf-8")
        except Exception:
            continue
        old_lines = old.splitlines(keepends=True)
        new_lines = new.splitlines(keepends=True)
        if len(old_lines) != len(new_lines):
            continue
        for idx, new_line in enumerate(new_lines):
            if new_line == old_lines[idx]:
                continue
            if suffix.sub(r"\1", old_lines[idx]) == new_line:
                match = suffix.search(old_lines[idx])
                if match:
                    bumped_module = match.group(0)
                new_lines[idx] = old_lines[idx]
        file_path.write_text("".join(new_lines), encoding="utf-8")
    return bumped_module


def get_spec_directory(package_folder: Path) -> Optional[str]:
    """Read the spec repo directory (containing tspconfig.yaml) from tsp-location.yaml."""
    tsp_location = package_folder / "tsp-location.yaml"
    if not tsp_location.exists():
        return None
    try:
        with open(tsp_location, "r", encoding="utf-8") as f:
            for line in f:
                stripped = line.strip()
                if stripped.startswith("directory:"):
                    return stripped[len("directory:"):].strip().strip('"')
    except FileNotFoundError as e:
        logging.warning(f"Failed to read tsp-location.yaml for {package_folder.name}: {e}")
    return None



def regenerate_sdk(use_latest_spec: bool, service_filter: str, sdk_root: str, typespec_go_root: str) -> Dict[str, Any]:
    result = {
        "succeed_to_regenerate": [],
        "fail_to_regenerate": [],
        "not_found_api_version": [],
        "module_version_changed": {},
        "time_to_regenerate": str(datetime.now()),
        "typespec_go_commit_hash": get_typespec_go_commit_hash(typespec_go_root),
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
        # Record the original module name so it is not changed by regeneration
        original_module = get_module_name(package_folder)
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
        finally:
            # Keep the original module name; do not bump the module version
            if original_module:
                bumped_module = restore_module_name(package_folder, original_module)
                if bumped_module:
                    spec_directory = get_spec_directory(package_folder)
                    if spec_directory:
                        result["module_version_changed"][spec_directory] = bumped_module
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


def bump_tspconfig_module(spec_root: str, spec_directory: str, bumped_module: str) -> bool:
    """Update only the go module version suffix in tspconfig.yaml.

    Returns True if the file was changed.
    """
    tspconfig_path = Path(spec_root) / spec_directory / "tspconfig.yaml"
    if not tspconfig_path.exists():
        logging.warning(f"tspconfig.yaml not found at {tspconfig_path}")
        return False
    suffix_match = re.search(r"/(v\d+)$", bumped_module)
    new_suffix = f"/{suffix_match.group(1)}" if suffix_match else ""
    with open(tspconfig_path, "r", encoding="utf-8") as f:
        content = f.readlines()
    changed = False
    for idx in range(len(content)):
        match = re.match(r"^(\s*module:\s*\"?)(\S+?)(/v\d+)?(\"?\s*)$", content[idx])
        if match:
            updated = f"{match.group(1)}{match.group(2)}{new_suffix}{match.group(4)}"
            if updated != content[idx]:
                content[idx] = updated
                changed = True
                logging.info(f"Updated module suffix in {tspconfig_path} to '{new_suffix}'")
            break
    if changed:
        with open(tspconfig_path, "w", encoding="utf-8") as f:
            f.writelines(content)
    return changed


def apply_tspconfig_module_bumps(spec_root: str, module_version_changed: dict) -> bool:
    """Bump go module suffixes in tspconfig.yaml for changed packages.

    Returns True if any tspconfig.yaml was modified. The pipeline is responsible for
    committing and opening the PR against the spec repo.
    """
    if not module_version_changed:
        logging.info("No module version changes; nothing to update in spec repo")
        return False
    if not spec_root or not Path(spec_root).exists():
        logging.warning("spec-root not provided or does not exist; skipping tspconfig update")
        return False

    changed_any = False
    for spec_directory, bumped_module in module_version_changed.items():
        if bump_tspconfig_module(spec_root, spec_directory, bumped_module):
            changed_any = True
    return changed_any


def main(sdk_root: str, typespec_go_root: str, typespec_go_branch: str, use_latest_spec: bool, service_filter: str, use_dev_package: bool, create_spec_pr_flag: bool, spec_root: str):
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

    # Print the result instead of committing it to the repo
    result_json = json.dumps(result, indent=2)
    logging.info("Regenerate SDK result:\n%s", result_json)

    # Write the result to the artifact staging directory so it can be published as a pipeline artifact
    staging_dir = os.environ.get("BUILD_ARTIFACTSTAGINGDIRECTORY")
    if staging_dir:
        result_path = Path(staging_dir) / "regenerate-sdk-result.json"
        with open(result_path, "w") as f:
            f.write(result_json)
        logging.info(f"Wrote regenerate-sdk-result.json to {result_path}")

    git_add()

    if create_spec_pr_flag:
        apply_tspconfig_module_bumps(spec_root, result.get("module_version_changed", {}))


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

    parser.add_argument(
        "--create-spec-pr",
        help="Whether to create a PR in the spec repo to bump go module suffixes in tspconfig.yaml",
        type=lambda x: x.lower() == 'true',
        default=False,
    )

    parser.add_argument(
        "--spec-root",
        help="azure-rest-api-specs repo root folder (required when --create-spec-pr is true)",
        type=str,
        default="",
    )

    args = parser.parse_args()

    main(args.sdk_root, args.typespec_go_root, args.typespec_go_branch, args.use_latest_spec, args.service_filter, args.use_dev_package, args.create_spec_pr, args.spec_root)