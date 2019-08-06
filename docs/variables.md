# variables.tf

```
variable "load_balancer_rule" {
  type = "map"
  default = {
    "APINAME.health_check_interval" = 10
    "APINAME.health_check_timeout"  = 5
    "APINAME.healthy_threshold"     = 3
    "APINAME.unhealthy_threshold"   = 2
    "APINAME.min_capacity"          = 1
    "APINAME.max_capacity"          = 20
    "APINAME.cpu_high_statistic"    = "Average"
    "APINAME.cpu_low_statistic"     = "Average"
    "APINAME.cpu_high_threshold"    = 60
    "APINAME.cpu_low_threshold"     = 30
    "APINAME.scale_up_cooldown"     = 600
    "APINAME.scale_down_cooldown"   = 600
  }
}
```