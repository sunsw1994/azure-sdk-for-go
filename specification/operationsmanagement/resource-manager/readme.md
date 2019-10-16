# OperationsManagement

> see https://aka.ms/autorest

This is the AutoRest configuration file for OperationsManagement



---
## Getting Started
To build the SDK for OperationsManagement, simply [Install AutoRest](https://aka.ms/autorest/install) and in this folder, run:

> `autorest`

To see additional help and options, run:

> `autorest --help`
---

## Configuration



### Basic Information
These are the global settings for the OperationsManagement API.

``` yaml
title: OperationsManagementClient
description: Operations Management Client
openapi-type: arm
tag: package-2015-11-preview
```


### Tag: package-2015-11-preview

These settings apply only when `--tag=package-2015-11-preview` is specified on the command line.

``` yaml $(tag) == 'package-2015-11-preview'
input-file:
- Microsoft.OperationsManagement/preview/2015-11-01-preview/OperationsManagement.json
```


---
# Code Generation


## Swagger to SDK

This section describes what SDK should be generated by the automatic system.
This is not used by Autorest itself.

``` yaml $(swagger-to-sdk)
swagger-to-sdk:
  - repo: azure-sdk-for-net
  - repo: azure-sdk-for-go
  - repo: azure-sdk-for-js
  - repo: azure-sdk-for-node
```


## C#

These settings apply only when `--csharp` is specified on the command line.
Please also specify `--csharp-sdks-folder=<path to "SDKs" directory of your azure-sdk-for-net clone>`.

```yaml $(csharp)
csharp:
  # last generated using AutoRest.1.0.0-Nightly20170126
  azure-arm: true
  namespace: Microsoft.Azure.Management.OperationsManagement
  payload-flattening-threshold: 1
  license-header: MICROSOFT_MIT_NO_VERSION
  output-folder: $(csharp-sdks-folder)/operationsmanagement/Microsoft.Azure.Management.OperationsManagement/src/Generated
  clear-output-folder: true
```

## Go

See configuration in [readme.go.md](./readme.go.md)

## Java

These settings apply only when `--java` is specified on the command line.
Please also specify `--azure-libraries-for-java-folder=<path to the root directory of your azure-libraries-for-java clone>`.

``` yaml $(java)
azure-arm: true
fluent: true
namespace: com.microsoft.azure.management.operationsmanagement
license-header: MICROSOFT_MIT_NO_CODEGEN
payload-flattening-threshold: 1
output-folder: $(azure-libraries-for-java-folder)/azure-mgmt-operationsmanagement
```

### Java multi-api

``` yaml $(java) && $(multiapi)
batch:
  - tag: package-2015-11-preview
```

### Tag: package-2015-11-preview and java

These settings apply only when `--tag=package-2015-11-preview --java` is specified on the command line.
Please also specify `--azure-libraries-for-java=<path to the root directory of your azure-sdk-for-java clone>`.

``` yaml $(tag) == 'package-2015-11-preview' && $(java) && $(multiapi)
java:
  namespace: com.microsoft.azure.management.operationsmanagement.v2015_11_01_preview
  output-folder: $(azure-libraries-for-java-folder)/operationsmanagement/resource-manager/v2015_11_01_preview
regenerate-manager: true
generate-interface: true
```



## Multi-API/Profile support for AutoRest v3 generators 

AutoRest V3 generators require the use of `--tag=all-api-versions` to select api files.

This block is updated by an automatic script. Edits may be lost!

``` yaml $(tag) == 'all-api-versions' /* autogenerated */
# include the azure profile definitions from the standard location
require: $(this-folder)/../../../profiles/readme.md

# all the input files across all versions
input-file:
  - $(this-folder)/Microsoft.OperationsManagement/preview/2015-11-01-preview/OperationsManagement.json

```

If there are files that should not be in the `all-api-versions` set, 
uncomment the  `exclude-file` section below and add the file paths.

``` yaml $(tag) == 'all-api-versions'
#exclude-file: 
#  - $(this-folder)/Microsoft.Example/stable/2010-01-01/somefile.json
```
