# TenancyGroup

### Available Operations

* [Create](#create) - Create a tenant service group
* [Delete](#delete) - Delete a tenant service group
* [Get](#get) - Get a tenant service group
* [List](#list) - List all tenant service groups
* [ListAncestors](#listancestors) - List tenant service group ancestors
* [ListChildren](#listchildren) - List tenant service group children
* [Update](#update) - Update a tenant service group

## Create

Create a tenant service group.
The service account used to authenticate this request
is granted `msp_superuser` access to the new tenant
service group.


### Example Usage

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

## Delete

Delete a tenant service group. If the TSG ID supplied
in this API's path does not match the TSG ID contained in
the access token used to authenticate this request, this
request will fail.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"sase-tenancy"
	"sase-tenancy/pkg/models/operations"
)

func main() {
    s := sasetenancy.New()

    ctx := context.Background()
    res, err := s.TenancyGroup.Delete(ctx, operations.DeleteTenancyV1TenantServiceGroupsTsgIDRequest{
        TsgID: "1378242802",
    }, operations.DeleteTenancyV1TenantServiceGroupsTsgIDSecurity{
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

## Get

Get a tenant service group by TSG ID.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"sase-tenancy"
	"sase-tenancy/pkg/models/operations"
)

func main() {
    s := sasetenancy.New()

    ctx := context.Background()
    res, err := s.TenancyGroup.Get(ctx, operations.GetTenancyV1TenantServiceGroupsTsgIDRequest{
        TsgID: "1378242802",
    }, operations.GetTenancyV1TenantServiceGroupsTsgIDSecurity{
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

## List

Get a list of all the tenant service groups
that are available to the service account used to
authenticate this request.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"sase-tenancy"
	"sase-tenancy/pkg/models/operations"
)

func main() {
    s := sasetenancy.New()

    ctx := context.Background()
    res, err := s.TenancyGroup.List(ctx, operations.GetTenancyV1TenantServiceGroupsRequest{
        Hierarchy: sasetenancy.Bool(false),
    }, operations.GetTenancyV1TenantServiceGroupsSecurity{
        Bearer: "YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetTenancyV1TenantServiceGroups200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## ListAncestors

List the ancestor tenants of the tenant service group
specified in this request. If the TSG ID supplied
in this API's path does not match the TSG ID contained in
the access token used to authenticate this request, this
request will fail.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"sase-tenancy"
	"sase-tenancy/pkg/models/operations"
)

func main() {
    s := sasetenancy.New()

    ctx := context.Background()
    res, err := s.TenancyGroup.ListAncestors(ctx, operations.PostTenancyV1TenantServiceGroupsTsgIDOperationsListAncestorsRequest{
        Fields: sasetenancy.String("corrupti"),
        IncludeSelf: sasetenancy.Bool(false),
        Sort: operations.PostTenancyV1TenantServiceGroupsTsgIDOperationsListAncestorsSortDesc.ToPointer(),
        TsgID: "1378242802",
    }, operations.PostTenancyV1TenantServiceGroupsTsgIDOperationsListAncestorsSecurity{
        Bearer: "YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostTenancyV1TenantServiceGroupsTsgIDOperationsListAncestors200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## ListChildren

List the child tenants of the tenant service group
specified in this request. If the TSG ID supplied
in this API's path does not match the TSG ID contained in
the access token used to authenticate this request, this
request will fail.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"sase-tenancy"
	"sase-tenancy/pkg/models/operations"
)

func main() {
    s := sasetenancy.New()

    ctx := context.Background()
    res, err := s.TenancyGroup.ListChildren(ctx, operations.PostTenancyV1TenantServiceGroupsTsgIDOperationsListChildrenRequest{
        Hierarchy: sasetenancy.Bool(false),
        IncludeSelf: sasetenancy.Bool(false),
        TsgID: "1378242802",
    }, operations.PostTenancyV1TenantServiceGroupsTsgIDOperationsListChildrenSecurity{
        Bearer: "YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostTenancyV1TenantServiceGroupsTsgIDOperationsListChildren200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## Update

Update a tenant service group. If the TSG ID supplied 
in this API's path does not match the TSG ID contained in
the access token used to authenticate this request, this 
request will fail.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"sase-tenancy"
	"sase-tenancy/pkg/models/operations"
	"sase-tenancy/pkg/models/shared"
)

func main() {
    s := sasetenancy.New()

    ctx := context.Background()
    res, err := s.TenancyGroup.Update(ctx, operations.PutTenancyV1TenantServiceGroupsTsgIDRequest{
        TenantServiceGroupUpdate: shared.TenantServiceGroupUpdate{
            DisplayName: sasetenancy.String("Example TSG"),
            SupportContact: sasetenancy.String("user@example.com"),
            Vertical: shared.TenantServiceGroupUpdateVerticalHighTech.ToPointer(),
        },
        TsgID: "1378242802",
    }, operations.PutTenancyV1TenantServiceGroupsTsgIDSecurity{
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
