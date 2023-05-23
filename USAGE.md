<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"sase-tenancy"
	"sase-tenancy/pkg/models/shared"
	"sase-tenancy/pkg/models/operations"
)

func main() {
    s := sasetenancy.New()

    ctx := context.Background()
    res, err := s.TenancyGroup.Create(ctx, shared.TenantServiceGroupCreate{
        DisplayName: "Example TSG",
        ParentID: sasetenancy.String("1378242802"),
        SupportContact: sasetenancy.String("user@example.com"),
        Vertical: shared.TenantServiceGroupCreateVerticalHighTech.ToPointer(),
    }, operations.PostTenancyV1TenantServiceGroupsSecurity{
        Bearer: "YOUR_BEARER_TOKEN_HERE",
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