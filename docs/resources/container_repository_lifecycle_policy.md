---
subcategory: "Container Registry"
page_title: "Yandex: yandex_container_repository_lifecycle_policy"
description: |-
  
---

# yandex_container_repository_lifecycle_policy

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `repository_id` (String)
- `status` (String)

### Optional

- `description` (String)
- `name` (String)
- `rule` (Block List) (see [below for nested schema](#nestedblock--rule))
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `created_at` (String)
- `id` (String) The ID of this resource.

<a id="nestedblock--rule"></a>
### Nested Schema for `rule`

Optional:

- `description` (String)
- `expire_period` (String)
- `retained_top` (Number)
- `tag_regexp` (String)
- `untagged` (Boolean)


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `default` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).