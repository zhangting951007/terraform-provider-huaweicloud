---
subcategory: "IoT Device Access (IoTDA)"
---

# huaweicloud_iotda_device

Manages an IoTDA device within HuaweiCloud.

## Example Usage

### Create a directly connected device and an indirectly connected device

```hcl
variable "spaceId" {}
variable "productId" {}
variable "secret" {}

resource "huaweicloud_iotda_device" "device" {
  node_id    = "device_SN_1"
  name       = "device_name"
  space_id   = var.spaceId
  product_id = var.productId
  secret     = var.secret

  tags = {
    foo = "bar"
    key = "value"
  }
}

resource "huaweicloud_iotda_device" "sub_device" {
  node_id    = "device_SN_2"
  name       = "device_name_2"
  space_id   = var.spaceId
  product_id = var.productId
  gateway_id = huaweicloud_iotda_device.device.id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the IoTDA device resource.
If omitted, the provider-level region will be used. Changing this parameter will create a new resource.

* `name` - (Required, String) Specifies the device name, which contains 4 to 256 characters. Only letters,
Chinese characters, digits, hyphens (-), underscore (_) and the following special characters are allowed: `?'#().,&%@!`.

* `node_id` - (Required, String, ForceNew) Specifies the node ID, which contains 4 to 256 characters.
The node ID can be IMEI, MAC address, or serial number. Changing this parameter will create a new resource.

* `space_id` - (Required, String, ForceNew) Specifies the resource space ID which the device belongs to.
Changing this parameter will create a new resource.

* `product_id` - (Required, String, ForceNew) Specifies the product ID which the device belongs to.
Changing this parameter will create a new resource.

* `device_id` - (Optional, String, ForceNew) Specifies the device ID, which contains 4 to 256 characters.
Only letters, digits, hyphens (-) and underscore (_) are allowed. If omitted, the platform will automatically allocate
a device ID. Changing this parameter will create a new resource.

* `secret` - (Optional, String) Specifies a secret for identity authentication, which contains 8 to 32 characters.
Only letters, digits, hyphens (-) and underscore (_) are allowed.

* `fingerprint` - (Optional, String) Specifies a fingerprint of X.509 certificate for identity authentication,
which is a 40-digit or 64-digit hexadecimal string. For more detail, please see
[Registering a Device Authenticated by an X.509 Certificate](https://support.huaweicloud.com/en-us/usermanual-iothub/iot_01_0055.html).

* `gateway_id` - (Optional, String) Specifies the gateway ID which is the device ID of the parent device.
The child device is not directly connected to the platform. If omitted, it means to create a device directly connected
to the platform, the `device_id` of the device is the same as the `gateway_id`.

* `description` - (Optional, String) Specifies the description of device. The description contains a maximum of 2048
characters. Only letters, Chinese characters, digits, hyphens (-), underscore (_) and the following special characters
are allowed: `?'#().,&%@!`.

* `tags` - (Optional, Map) Specifies the key/value pairs to associate with the device.

* `frozen` - (Optional, Bool) Specifies whether to freeze the device. Defaults to `false`.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The device ID in UUID format.

* `status` - The status of device. The valid values are **INACTIVE**, **ONLINE**, **OFFLINE**, **FROZEN**, **ABNORMAL**.

* `auth_type` - The authentication type of device. The options are as follows:
  + **SECRET**: Use a secret for identity authentication.
  + **CERTIFICATES**: Use an x.509 certificate for identity authentication.

* `node_type` - The node type of device. The options are as follows:
  + **GATEWAY**: Directly connected device.
  + **ENDPOINT**: Indirectly connected device.
  + **UNKNOWN**: Unknown type.

## Import

Devices can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_iotda_device.test 10022532f4f94f26b01daa1e424853e1
```
