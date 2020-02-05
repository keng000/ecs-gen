# variables.tf

```
variable "load-balancer-rule" {
  type = "map"
  default = {
    "APINAME.health-check-interval" = 10
    "APINAME.health-check-timeout"  = 5
    "APINAME.healthy-threshold"     = 3
    "APINAME.unhealthy-threshold"   = 2
    "APINAME.min-capacity"          = 1
    "APINAME.max-capacity"          = 20
    "APINAME.cpu_high_statistic"    = "Average"
    "APINAME.cpu_low_statistic"     = "Average"
    "APINAME.cpu_high_threshold"    = 60
    "APINAME.cpu_low_threshold"     = 30
    "APINAME.scale_up_cooldown"     = 600
    "APINAME.scale_down_cooldown"   = 600
  }
}
```