<div align="center">
    <img src="https://github.com/speakeasy-sdks/pan-sase-tenancy-ts/assets/6267663/e0205c2a-fa61-4b1f-aaa0-599896e022da" width="300">
    <h1>SASE Tenancy Go SDK</h1>
   <p>Containers used to build your tenant hierachy.</p>
   <a href="https://pan.dev/category/sase/api/tenancy/tenant-service-group/"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/pan-sase-tenancy-go/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/pan-sase-tenancy-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/pan-sase-tenancy-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"sase-tenancy"
	"sase-tenancy/pkg/models/shared"
)

func main() {
    s := sasetenancy.New(
        sasetenancy.WithSecurity(shared.Security{
            Bearer: "",
        }),
    )

    ctx := context.Background()
    res, err := s.TenancyGroup.Create(ctx, shared.TenantServiceGroupCreate{
        DisplayName: "Example TSG",
        ParentID: sasetenancy.String("1378242802"),
        SupportContact: sasetenancy.String("user@example.com"),
        Vertical: shared.TenantServiceGroupCreateVerticalHighTech.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TenantServiceGroup != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [TenancyGroup](docs/sdks/tenancygroup/README.md)

* [Create](docs/sdks/tenancygroup/README.md#create) - Create a tenant service group
* [Delete](docs/sdks/tenancygroup/README.md#delete) - Delete a tenant service group
* [Get](docs/sdks/tenancygroup/README.md#get) - Get a tenant service group
* [List](docs/sdks/tenancygroup/README.md#list) - List all tenant service groups
* [ListAncestors](docs/sdks/tenancygroup/README.md#listancestors) - List tenant service group ancestors
* [ListChildren](docs/sdks/tenancygroup/README.md#listchildren) - List tenant service group children
* [Update](docs/sdks/tenancygroup/README.md#update) - Update a tenant service group
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
