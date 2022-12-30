---
subcategory: "Cloud Firewall (CRF)"
---

# huaweicloud_cfw_protection_rule

Manages a CFW protection rule resource within HuaweiCloud.

## Example Usage

```HCL
variable "name" {}
variable "description" {}
variable "object_id" {}

resource "huaweicloud_cfw_protection_rule" "test" {
  name                = var.name
  object_id           = var.object_id
  description         = var.description
  type                = 0
  address_type        = 0
  action_type         = 0
  long_connect_enable = 0
  status              = 1

  source {
    type    = 0
    address = "192.168.0.1"
  }

  destination {
    type    = 0
    address = "192.168.0.2"
  }

  service {
    type        = 0
    protocol    = 6
    source_port = 8001
    dest_port   = 8002
  }

  sequence {
    top = 1
  }
}
```

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the resource.
  If omitted, the provider-level region will be used. Changing this parameter will create a new resource.

* `name` - (Required, String) The rule name.

* `object_id` - (Required, String, ForceNew) The protected object ID

  Changing this parameter will create a new resource.

* `type` - (Required, Int) The rule type.
  The value can be **0** (Internet rule), **1** (VPC rule), or **2** (NAT rule).

* `action_type` - (Required, Int) The action type.
  The value can be **0** (allow) **1** (deny).

* `address_type` - (Required, Int) The address type.
  The value can be **0** (IPv4), **1** (IPv6), or **2** (domain).

* `sequence` - (Required, String) The sequence configuration.
The [Order Rule](#ProtectionRule_OrderRuleAcl) structure is documented below.

* `service` - (Required, String) The service configuration.
The [Rule Service](#ProtectionRule_RuleService) structure is documented below.

* `source` - (Required, String) The source configuration.
The [Rule Source Address](#ProtectionRule_RuleSourceAddress) structure is documented below.

* `destination` - (Required, String) The destination configuration.
The [Rule Destination Address](#ProtectionRule_RuleDestinationAddress) structure is documented below.

* `status` - (Required, Int) The rule status. The options are as follows:
  + **0**: disabled;
  + **1**: enabled;

* `long_connect_enable` - (Required, Int) Whether to support persistent connections.
  The options are as follows:
  + **0**: supported;
  + **1**: not supported;

* `long_connect_time_hour` - (Optional, Int) The persistent connection duration (hour).

* `long_connect_time_minute` - (Optional, Int) The persistent connection duration (minute).

* `long_connect_time_second` - (Optional, Int) The persistent Connection Duration (second).

* `description` - (Optional, String) The description.

* `direction` - (Optional, Int) The direction. The options are as follows:
  + **0**: inbound;
  + **1**: outbound;

<a name="ProtectionRule_OrderRuleAcl"></a>
The `sequence` block supports:

* `dest_rule_id` - (Optional, String) The ID of the rule that the added rule will follow.
  This parameter cannot be left blank if the rule is not pinned on top, and is empty when the added rule is pinned on top.

* `top` - (Optional, Int) Whether to pin on top.
  The options are as follows:
  + **0**: no;
  + **1**: yes;

<a name="ProtectionRule_RuleService"></a>
The `service` block supports:

* `type` - (Required, Int) The service input type.
  The value **0** indicates manual input, and the value **1** indicates automatic input.

* `dest_port` - (Optional, String) The destination port.

* `protocol` - (Optional, Int) The protocol type. The options are as follows:
  + **6**: TCP;
  + **17**: UDP;
  + **1**: ICMP;
  + **58**: ICMPv6;
  + **-1**: any protocol;
  
  Regarding the addition type, a null value indicates it is automatically added.

* `service_set_id` - (Optional, String) The service group ID.
  This parameter is left blank for the manual type and cannot be left blank for the automatic type.

* `service_set_name` - (Optional, String) The service group name.

* `source_port` - (Optional, String) The source port.

<a name="ProtectionRule_RuleSourceAddress"></a>
The `source` block supports:

* `type` - (Required, Int) The Source type. The options are as follows:
  + **0**: manual input;
  + **1**: associated IP address group;
  + **2**: domain name;

* `address` - (Optional, String) The IP address.
  The value cannot be empty for the manual type, and cannot be empty for the automatic or domain type.

* `address_set_id` - (Optional, String) The ID of the associated IP address group.
  The value cannot be empty for the automatic type or for the manual or domain type.

* `address_set_name` - (Optional, String) The IP address group name.

* `address_type` - (Optional, Int) The address type. The options are as follows:
  + **0**: IPv4;
  + **1**: IPv6;

* `domain_address_name` - (Optional, String) The name of the domain name address.
  This parameter cannot be left empty for the domain name type, and is empty for the manual or automatic type.

<a name="ProtectionRule_RuleDestinationAddress"></a>
The `destination` block supports:

* `type` - (Required, Int) The Source type. The options are as follows:
  + **0**: manual input;
  + **1**: associated IP address group;
  + **2**: domain name;

* `address` - (Optional, String) The IP address.
  The value cannot be empty for the manual type, and cannot be empty for the automatic or domain type.

* `address_set_id` - (Optional, String) The ID of the associated IP address group.
  The value cannot be empty for the automatic type or for the manual or domain type.

* `address_set_name` - (Optional, String) The IP address group name.

* `address_type` - (Optional, Int) The address type. The options are as follows:
  + **0**: IPv4;
  + **1**: IPv6;

* `domain_address_name` - (Optional, String) The name of the domain name address.
  This parameter cannot be left empty for the domain name type, and is empty for the manual or automatic type.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID.

## Import

The protectionrule can be imported using `object_id`, `id`, separated by a slash, e.g.

```sh
$ terraform import huaweicloud_cfw_protection_rule.test <object_id>/<id>
```