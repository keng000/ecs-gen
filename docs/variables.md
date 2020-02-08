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
    "APINAME.cpu-high-statistic"    = "Average"
    "APINAME.cpu-low-statistic"     = "Average"
    "APINAME.cpu-high-threshold"    = 60
    "APINAME.cpu-low-threshold"     = 30
    "APINAME.scale-up-cooldown"     = 600
    "APINAME.scale-down-cooldown"   = 600
  }
}
```