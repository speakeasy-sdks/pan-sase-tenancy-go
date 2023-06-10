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

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [shared.TenantServiceGroupCreate](../../models/shared/tenantservicegroupcreate.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.PostTenancyV1TenantServiceGroupsResponse](../../models/operations/posttenancyv1tenantservicegroupsresponse.md), error**


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
    s := sasetenancy.New(
        sasetenancy.WithSecurity(shared.Security{
            Bearer: "",
        }),
    )

    ctx := context.Background()
    res, err := s.TenancyGroup.Delete(ctx, "1378242802")
    if err != nil {
        log.Fatal(err)
    }

    if res.TenantServiceGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `tsgID`                                               | *string*                                              | :heavy_check_mark:                                    | A unique identifier for the tenant service group.<br/> | 1378242802                                            |


### Response

**[*operations.DeleteTenancyV1TenantServiceGroupsTsgIDResponse](../../models/operations/deletetenancyv1tenantservicegroupstsgidresponse.md), error**


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
    s := sasetenancy.New(
        sasetenancy.WithSecurity(shared.Security{
            Bearer: "",
        }),
    )

    ctx := context.Background()
    res, err := s.TenancyGroup.Get(ctx, "1378242802")
    if err != nil {
        log.Fatal(err)
    }

    if res.TenantServiceGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `tsgID`                                               | *string*                                              | :heavy_check_mark:                                    | A unique identifier for the tenant service group.<br/> | 1378242802                                            |


### Response

**[*operations.GetTenancyV1TenantServiceGroupsTsgIDResponse](../../models/operations/gettenancyv1tenantservicegroupstsgidresponse.md), error**


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
    s := sasetenancy.New(
        sasetenancy.WithSecurity(shared.Security{
            Bearer: "",
        }),
    )

    ctx := context.Background()
    res, err := s.TenancyGroup.List(ctx, false)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetTenancyV1TenantServiceGroups200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                       | Type                                                                                                                                                                                                                                            | Required                                                                                                                                                                                                                                        | Description                                                                                                                                                                                                                                     |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                                              | The context to use for the request.                                                                                                                                                                                                             |
| `hierarchy`                                                                                                                                                                                                                                     | **bool*                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                              | Indicates whether the response structure lists groups in<br/>their hierarchy, or as an array of TSGs without regard to<br/>hierarchy. Default is false (don't show hierarchy).<br/><br/>If false, the order of the TSGs in the result array is not<br/>guaranteed.<br/> |


### Response

**[*operations.GetTenancyV1TenantServiceGroupsResponse](../../models/operations/gettenancyv1tenantservicegroupsresponse.md), error**


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
    s := sasetenancy.New(
        sasetenancy.WithSecurity(shared.Security{
            Bearer: "",
        }),
    )

    ctx := context.Background()
    res, err := s.TenancyGroup.ListAncestors(ctx, "1378242802", "corrupti", false, operations.PostTenancyV1TenantServiceGroupsTsgIDOperationsListAncestorsSortDesc)
    if err != nil {
        log.Fatal(err)
    }

    if res.PostTenancyV1TenantServiceGroupsTsgIDOperationsListAncestors200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                   | Type                                                                                                                                                                        | Required                                                                                                                                                                    | Description                                                                                                                                                                 | Example                                                                                                                                                                     |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                       | :heavy_check_mark:                                                                                                                                                          | The context to use for the request.                                                                                                                                         |                                                                                                                                                                             |
| `tsgID`                                                                                                                                                                     | *string*                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                          | A unique identifier for the tenant service group.<br/>                                                                                                                      | 1378242802                                                                                                                                                                  |
| `fields`                                                                                                                                                                    | **string*                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                          | Provide a comma-separated list of fields you want returned.<br/>                                                                                                            |                                                                                                                                                                             |
| `includeSelf`                                                                                                                                                               | **bool*                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                          | Indicates if the TSG used to generate this hierarchy is<br/>included in the resulting TSG list. `true` to include<br/>self. Default is `false`.<br/>                        |                                                                                                                                                                             |
| `sort`                                                                                                                                                                      | [*operations.PostTenancyV1TenantServiceGroupsTsgIDOperationsListAncestorsSort](../../models/operations/posttenancyv1tenantservicegroupstsgidoperationslistancestorssort.md) | :heavy_minus_sign:                                                                                                                                                          | Identifies the response structure's sort order:<br/><br/>* `asc` : From root to leaf.<br/>* `desc` : From leaf to root.<br/>                                                |                                                                                                                                                                             |


### Response

**[*operations.PostTenancyV1TenantServiceGroupsTsgIDOperationsListAncestorsResponse](../../models/operations/posttenancyv1tenantservicegroupstsgidoperationslistancestorsresponse.md), error**


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
    s := sasetenancy.New(
        sasetenancy.WithSecurity(shared.Security{
            Bearer: "",
        }),
    )

    ctx := context.Background()
    res, err := s.TenancyGroup.ListChildren(ctx, "1378242802", false, false)
    if err != nil {
        log.Fatal(err)
    }

    if res.PostTenancyV1TenantServiceGroupsTsgIDOperationsListChildren200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                         | Type                                                                                                                                                              | Required                                                                                                                                                          | Description                                                                                                                                                       | Example                                                                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                                                             | :heavy_check_mark:                                                                                                                                                | The context to use for the request.                                                                                                                               |                                                                                                                                                                   |
| `tsgID`                                                                                                                                                           | *string*                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                | A unique identifier for the tenant service group.<br/>                                                                                                            | 1378242802                                                                                                                                                        |
| `hierarchy`                                                                                                                                                       | **bool*                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                | If `true`, return the entire descendent hierarchy.<br/>If `false`, return only the immediate children of the<br/>TSG identified in this call's path. Default is<br/>`false`.<br/> |                                                                                                                                                                   |
| `includeSelf`                                                                                                                                                     | **bool*                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                | Indicates if the TSG used to generate this hierarchy is<br/>included in the resulting TSG list. `true` to include<br/>self. Default is `false`.<br/>              |                                                                                                                                                                   |


### Response

**[*operations.PostTenancyV1TenantServiceGroupsTsgIDOperationsListChildrenResponse](../../models/operations/posttenancyv1tenantservicegroupstsgidoperationslistchildrenresponse.md), error**


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
    s := sasetenancy.New(
        sasetenancy.WithSecurity(shared.Security{
            Bearer: "",
        }),
    )

    ctx := context.Background()
    res, err := s.TenancyGroup.Update(ctx, shared.TenantServiceGroupUpdate{
        DisplayName: sasetenancy.String("Example TSG"),
        SupportContact: sasetenancy.String("user@example.com"),
        Vertical: shared.TenantServiceGroupUpdateVerticalHighTech.ToPointer(),
    }, "1378242802")
    if err != nil {
        log.Fatal(err)
    }

    if res.TenantServiceGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `tenantServiceGroupUpdate`                                                         | [shared.TenantServiceGroupUpdate](../../models/shared/tenantservicegroupupdate.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `tsgID`                                                                            | *string*                                                                           | :heavy_check_mark:                                                                 | A unique identifier for the tenant service group.<br/>                             | 1378242802                                                                         |


### Response

**[*operations.PutTenancyV1TenantServiceGroupsTsgIDResponse](../../models/operations/puttenancyv1tenantservicegroupstsgidresponse.md), error**

