# Backup Admin

> see https://aka.ms/autorest

This is the AutoRest configuration file for Backup Admin.

---
## Getting Started
To build the SDK for Backup Admin, simply [Install AutoRest](https://aka.ms/autorest/install) and in this folder, run:

> `autorest`

To see additional help and options, run:

> `autorest --help`
---

## Configuration

### Basic Information
These are the global settings for the Backup Admin API.

``` yaml
title: BackupAdminClient
description: Backup Admin Client
openapi-type: arm
tag: package-2016-05-01
```

### Tag: package-2018-09-01

These settings apply only when `--tag=package-2018-09-01` is specified on the command line.

``` yaml $(tag) == 'package-2018-09-01'
input-file:
    - Microsoft.Backup.Admin/preview/2018-09-01/Backup.json
    - Microsoft.Backup.Admin/preview/2018-09-01/BackupLocations.json
    - Microsoft.Backup.Admin/preview/2018-09-01/Backups.json
```

### Tag: package-2016-05-01

These settings apply only when `--tag=package-2016-05-01` is specified on the command line.

``` yaml $(tag) == 'package-2016-05-01'
input-file:
    - Microsoft.Backup.Admin/stable/2016-05-01/Backup.json
    - Microsoft.Backup.Admin/stable/2016-05-01/BackupLocations.json
    - Microsoft.Backup.Admin/stable/2016-05-01/Backups.json
```

---
# Code Generation

## C#

``` yaml $(csharp)
csharp:
  azure-arm: true
  license-header: MICROSOFT_MIT_NO_VERSION
  namespace: Microsoft.AzureStack.Management.Backup.Admin
  payload-flattening-threshold: 1
  output-folder: $(csharp-sdks-folder)/Generated
  clear-output-folder: true
```

## Multi-API/Profile support for AutoRest v3 generators 

AutoRest V3 generators require the use of `--tag=all-api-versions` to select api files.

This block is updated by an automatic script. Edits may be lost!

``` yaml $(tag) == 'all-api-versions' /* autogenerated */
# include the azure profile definitions from the standard location
require: $(this-folder)/../../../../profiles/readme.md

# all the input files across all versions
input-file:
  - $(this-folder)/Microsoft.Backup.Admin/preview/2018-09-01/Backup.json
  - $(this-folder)/Microsoft.Backup.Admin/preview/2018-09-01/BackupLocations.json
  - $(this-folder)/Microsoft.Backup.Admin/preview/2018-09-01/Backups.json
  - $(this-folder)/Microsoft.Backup.Admin/stable/2016-05-01/Backup.json
  - $(this-folder)/Microsoft.Backup.Admin/stable/2016-05-01/BackupLocations.json
  - $(this-folder)/Microsoft.Backup.Admin/stable/2016-05-01/Backups.json

```

If there are files that should not be in the `all-api-versions` set, 
uncomment the  `exclude-file` section below and add the file paths.

``` yaml $(tag) == 'all-api-versions'
#exclude-file: 
#  - $(this-folder)/Microsoft.Example/stable/2010-01-01/somefile.json
```
