---
subcategory: "Identity and Access Management (IAM)"
page_title: "Yandex: yandex_iam_service_account"
description: |-
  Get information about a Yandex IAM service account.
---


# yandex_iam_service_account




Get information about a Yandex IAM service account. For more information about accounts, see [Yandex.Cloud IAM accounts](https://cloud.yandex.com/docs/iam/concepts/#accounts).

## Example usage

```terraform
data "yandex_iam_service_account" "builder" {
  service_account_id = "sa_id"
}

data "yandex_iam_service_account" "deployer" {
  name = "sa_name"
}
```

## Argument reference

* `service_account_id` - (Optional) ID of a specific service account.

* `name` - (Optional) Name of a specific service account.

~> **NOTE:** One of `service_account_id` or `name` should be specified.

* `folder_id` - (Optional) Folder that the resource belongs to. If value is omitted, the default provider folder is used.

## Attributes Reference

* `description` - Description of the service account.
