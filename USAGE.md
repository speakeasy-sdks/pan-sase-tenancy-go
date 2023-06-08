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