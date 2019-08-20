package skeleton

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets72af51fef0d5e6fe9c11b8bd19b758272de17fbb = "resource \"aws_alb_target_group\" \"{{ .APIName }}\" {\n  name        = \"{{ .Project }}-{{ .APIName }}-tg\"\n  port        = 8000\n  protocol    = \"HTTP\"\n  vpc_id      = var.vpc_vpc_{{ .Project }}_id\n  target_type = \"ip\"\n\n  health_check {\n    interval            = lookup(var.load_balancer_rule, \"{{ .APIName }}.health_check_interval\")\n    path                = \"/health\"\n    port                = 8000\n    protocol            = \"HTTP\"\n    timeout             = lookup(var.load_balancer_rule, \"{{ .APIName }}.health_check_timeout\")\n    unhealthy_threshold = lookup(var.load_balancer_rule, \"{{ .APIName }}.unhealthy_threshold\")\n    matcher             = 200\n  }\n}\n\nresource \"aws_alb_listener_rule\" \"{{ .APIName }}\" {\n  listener_arn = aws_alb_listener.{{ .Project }}_http.arn\n\n  action {\n    type             = \"forward\"\n    target_group_arn = aws_alb_target_group.{{ .APIName }}.arn\n  }\n\n  condition {\n    field  = \"path-pattern\"\n    values = [\"/{{ .APIName }}\"]\n  }\n}\n\n# ## Auto Scaling\n# resource \"aws_appautoscaling_target\" \"ac_{{ .APIName }}\" {\n#   service_namespace  = \"ecs\"\n#   resource_id        = \"service/${var.ecs_{{ .Project }}_cluster.name}/{{ .APIName }}\"\n#   scalable_dimension = \"ecs:service:DesiredCount\"\n#   role_arn           = data.aws_iam_role.ecs_service_autoscaling.arn\n#   min_capacity       = lookup(var.load_balancer_rule, \"{{ .APIName }}.min_capacity\")\n#   max_capacity       = lookup(var.load_balancer_rule, \"{{ .APIName }}.max_capacity\")\n# }\n# \n# resource \"aws_appautoscaling_policy\" \"{{ .APIName }}_scale_up\" {\n#   name               = \"{{ .APIName }}_scale_up\"\n#   service_namespace  = \"ecs\"\n#   resource_id        = \"service/${var.ecs_{{ .Project }}_cluster.name}/{{ .APIName }}\"\n#   scalable_dimension = \"ecs:service:DesiredCount\"\n# \n#   step_scaling_policy_configuration {\n#     adjustment_type         = \"ChangeInCapacity\"\n#     cooldown                = lookup(var.load_balancer_rule, \"{{ .APIName }}.scale_up_cooldown\")\n#     metric_aggregation_type = \"Average\"\n# \n#     step_adjustment {\n#       metric_interval_lower_bound = 0\n#       scaling_adjustment          = 1\n#     }\n#   }\n# \n#   depends_on = [\"aws_appautoscaling_target.ac_{{ .APIName }}\"]\n# }\n# \n# resource \"aws_appautoscaling_policy\" \"{{ .APIName }}_scale_down\" {\n#   name               = \"{{ .APIName }}_scale_down\"\n#   service_namespace  = \"ecs\"\n#   resource_id        = \"service/${var.ecs_{{ .Project }}_cluster.name}/{{ .APIName }}\"\n#   scalable_dimension = \"ecs:service:DesiredCount\"\n# \n#   step_scaling_policy_configuration {\n#     adjustment_type         = \"ChangeInCapacity\"\n#     cooldown                = lookup(var.load_balancer_rule, \"{{ .APIName }}.scale_down_cooldown\")\n#     metric_aggregation_type = \"Average\"\n# \n#     step_adjustment {\n#       metric_interval_upper_bound = 0\n#       scaling_adjustment          = -1\n#     }\n#   }\n# \n#   depends_on = [\"aws_appautoscaling_target.ac_{{ .APIName }}\"]\n# }\n# \n# resource \"aws_cloudwatch_metric_alarm\" \"{{ .APIName }}_cpu_high\" {\n#   alarm_name          = \"{{ .APIName }}_cpu_high\"\n#   comparison_operator = \"GreaterThanOrEqualToThreshold\"\n#   evaluation_periods  = \"2\"\n#   metric_name         = \"CPUUtilization\"\n#   namespace           = \"AWS/ECS\"\n#   period              = \"60\"\n#   statistic           = lookup(var.load_balancer_rule, \"{{ .APIName }}.cpu_high_statistic\")\n#   threshold           = lookup(var.load_balancer_rule, \"{{ .APIName }}.cpu_high_threshold\")\n# \n#   dimensions = {\n#     ClusterName = var.ecs_{{ .Project }}_cluster.name\n#     ServiceName = \"{{ .APIName }}\"\n#   }\n# \n#   alarm_actions = [\"${aws_appautoscaling_policy.{{ .APIName }}_scale_up.arn}\"]\n# }\n# \n# resource \"aws_cloudwatch_metric_alarm\" \"{{ .APIName }}_cpu_low\" {\n#   alarm_name          = \"{{ .APIName }}_cpu_low\"\n#   comparison_operator = \"LessThanOrEqualToThreshold\"\n#   evaluation_periods  = \"2\"\n#   metric_name         = \"CPUUtilization\"\n#   namespace           = \"AWS/ECS\"\n#   period              = \"60\"\n#   statistic           = lookup(var.load_balancer_rule, \"{{ .APIName }}.cpu_low_statistic\")\n#   threshold           = lookup(var.load_balancer_rule, \"{{ .APIName }}.cpu_low_threshold\")\n# \n#   dimensions = {\n#     ClusterName = var.ecs_{{ .Project }}_cluster.name\n#     ServiceName = \"{{ .APIName }}\"\n#   }\n# \n#   alarm_actions = [\"${aws_appautoscaling_policy.{{ .APIName }}_scale_down.arn}\"]\n# }"
var _Assetsffc03a188b966e5e41bdff75413f26d12cea7441 = "resource \"aws_iam_role\" \"{{ .Project }}_task_exec_role\" {\n  name = \"{{ .Project }}_${terraform.workspace}_task_exec_role\"\n  path = \"/system/\"\n\n  assume_role_policy = <<EOF\n{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Sid\": \"1\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"sts:AssumeRole\",\n      \"Principal\": {\n        \"Service\": \"ecs-tasks.amazonaws.com\"\n      }\n    }\n  ]\n}\nEOF\n}\n\nresource \"aws_iam_role_policy\" \"{{ .Project }}_task_exec_policy\" {\n  name = \"{{ .Project }}_${terraform.workspace}_task_exec_policy\"\n  role = aws_iam_role.{{ .Project }}_task_exec_role.id\n\n  policy = <<EOF\n{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Effect\": \"Allow\",\n      \"Action\": [\n        \"ecr:GetAuthorizationToken\",\n        \"ecr:BatchCheckLayerAvailability\",\n        \"ecr:GetDownloadUrlForLayer\",\n        \"ecr:BatchGetImage\",\n        \"logs:CreateLogStream\",\n        \"logs:PutLogEvents\",\n\n        \"ssm:GetParameters\",\n        \"secretsmanager:GetSecretValue\",\n        \"kms:Decrypt\",\n\n        \"s3:*\"\n      ],\n      \"Resource\": \"*\"\n    }\n  ]\n}\nEOF\n}\n"
var _Assets3ff5abfa5c08898738456203761e92760f43576f = "resource \"aws_security_group\" \"{{ .Project }}_public_subnet_all_tcp\" {\n  vpc_id = var.vpc_vpc_{{ .Project }}_id\n  name   = \"{{ .Project }}_public_subnet_all_tcp\"\n\n  ingress {\n    from_port = 0\n    to_port   = 0\n    protocol  = \"-1\"\n\n    cidr_blocks = [\n      var.vpc_subnet_{{ .Project }}_public_a_cidr_block,\n      var.vpc_subnet_{{ .Project }}_public_c_cidr_block,\n    ]\n  }\n\n  egress {\n    from_port   = 0\n    to_port     = 0\n    protocol    = \"-1\"\n    cidr_blocks = [\"0.0.0.0/0\"]\n  }\n\n  tags = {\n    Name        = \"{{ .Project }}publicsubnetalltcp\"\n    Environment = terraform.workspace\n  }\n}\n\nresource \"aws_security_group\" \"{{ .Project }}_private_subnet_all_tcp\" {\n  vpc_id = var.vpc_vpc_{{ .Project }}_id\n  name   = \"{{ .Project }}_private_subnet_all_tcp\"\n\n  ingress {\n    from_port = 0\n    to_port   = 0\n    protocol  = \"-1\"\n\n    cidr_blocks = [\n      var.vpc_subnet_{{ .Project }}_private_a_cidr_block,\n      var.vpc_subnet_{{ .Project }}_private_c_cidr_block,\n    ]\n  }\n\n  egress {\n    from_port   = 0\n    to_port     = 0\n    protocol    = \"-1\"\n    cidr_blocks = [\"0.0.0.0/0\"]\n  }\n\n  tags = {\n    Name        = \"{{ .Project }}privatesubnetalltcp\"\n    Environment = terraform.workspace\n  }\n}\n\nresource \"aws_security_group\" \"{{ .Project }}_alb\" {\n  vpc_id = var.vpc_vpc_{{ .Project }}_id\n  name   = \"{{ .Project }}_alb\"\n\n  ingress {\n    from_port   = 80\n    to_port     = 80\n    protocol    = \"tcp\"\n    cidr_blocks = [\"0.0.0.0/0\"]\n  }\n\n  ingress {\n    from_port   = 443\n    to_port     = 443\n    protocol    = \"tcp\"\n    cidr_blocks = [\"0.0.0.0/0\"]\n  }\n\n  ingress {\n    from_port   = 8000\n    to_port     = 8000\n    protocol    = \"tcp\"\n    cidr_blocks = [\"0.0.0.0/0\"]\n  }\n\n  egress {\n    from_port   = 0\n    to_port     = 0\n    protocol    = \"-1\"\n    cidr_blocks = [\"0.0.0.0/0\"]\n  }\n\n  tags = {\n    Name        = \"{{ .Project }}alb\"\n    Environment = terraform.workspace\n  }\n}\n\nresource \"aws_security_group\" \"{{ .Project }}_ecs\" {\n  vpc_id = var.vpc_vpc_{{ .Project }}_id\n  name   = \"{{ .Project }}_ecs\"\n\n  ingress {\n    from_port   = 80\n    to_port     = 80\n    protocol    = \"tcp\"\n    cidr_blocks = [\"0.0.0.0/0\"]\n\n    security_groups = [\n      aws_security_group.{{ .Project }}_alb.id,\n    ]\n  }\n\n  ingress {\n    from_port   = 443\n    to_port     = 443\n    protocol    = \"tcp\"\n    cidr_blocks = [\"0.0.0.0/0\"]\n\n    security_groups = [\n      aws_security_group.{{ .Project }}_alb.id,\n    ]\n  }\n\n  ingress {\n    from_port   = 8000\n    to_port     = 8000\n    protocol    = \"tcp\"\n    cidr_blocks = [\"0.0.0.0/0\"]\n\n    security_groups = [\n      aws_security_group.{{ .Project }}_alb.id,\n    ]\n  }\n\n  egress {\n    from_port   = 0\n    to_port     = 0\n    protocol    = \"-1\"\n    cidr_blocks = [\"0.0.0.0/0\"]\n  }\n\n  tags = {\n    Name        = \"{{ .Project }}ecs\"\n    Environment = terraform.workspace\n  }\n}\n"
var _Assetse3d0b8bc4917cfbd1a87d600a89e4f7894dc6e33 = "output \"{{.Project }}_private_subnet_all_tcp_id\" {\n  value = aws_security_group.{{.Project }}_private_subnet_all_tcp.id\n}\n\noutput \"{{.Project }}_public_subnet_all_tcp_id\" {\n  value = aws_security_group.{{.Project }}_public_subnet_all_tcp.id\n}\n\noutput \"{{.Project }}_alb_id\" {\n  value = aws_security_group.{{.Project }}_alb.id\n}\n\noutput \"{{.Project }}_ecs_id\" {\n  value = aws_security_group.{{.Project }}_ecs.id\n}\n"
var _Assetsbacbea7a9bc4dcf5062d63ac8c05ff548dde1924 = "{\n    \"rules\": [\n        {\n            \"rulePriority\": 1,\n            \"description\": \"Expire images older than 30 days\",\n            \"selection\": {\n                \"tagStatus\": \"untagged\",\n                \"countType\": \"sinceImagePushed\",\n                \"countUnit\": \"days\",\n                \"countNumber\": 30\n            },\n            \"action\": {\n                \"type\": \"expire\"\n            }\n        }\n    ]\n}"
var _Assets15a142d2a032e7816f336c2df6d25fce4ad65b2f = "resource \"aws_ecr_repository\" \"{{ .APIName }}\" {\n  name = \"{{ .Project }}/{{ .APIName }}\"\n}\n\nresource \"aws_ecr_lifecycle_policy\" \"{{ .APIName }}\" {\n  repository = \"${aws_ecr_repository.{{ .APIName }}.name}\"\n\n  policy = \"${file(\"${path.module}/policy.json\")}\"\n}\n"
var _Assets335e2f5adb3f568fa4e3da8ecef12ec0779a5c06 = "output \"{{ .Project }}_cluster\" {\n  value = aws_ecs_cluster.{{ .Project }}_cluster\n}\n"
var _Assetse65d182e7adf57e01416dca18760fe8cb0a96886 = "variable \"vpc_subnet_{{ .Project }}_private_a_cidr_block\" {}\nvariable \"vpc_subnet_{{ .Project }}_private_c_cidr_block\" {}\nvariable \"vpc_subnet_{{ .Project }}_public_a_cidr_block\" {}\nvariable \"vpc_subnet_{{ .Project }}_public_c_cidr_block\" {}\nvariable \"vpc_vpc_{{ .Project }}_id\" {}\n"
var _Assetsc87eece7149baef6e249d528de3f9980d6e9aacf = "output \"vpc_{{ .Project }}_id\" {\n  value = \"${aws_vpc.{{ .Project }}.id}\"\n}\n\noutput \"subnet_{{ .Project }}_public_a_id\" {\n  value = aws_subnet.{{ .Project }}_public_a.id\n}\n\noutput \"subnet_{{ .Project }}_public_c_id\" {\n  value = aws_subnet.{{ .Project }}_public_c.id\n}\n\noutput \"subnet_{{ .Project }}_public_a_cidr_block\" {\n  value = aws_subnet.{{ .Project }}_public_a.cidr_block\n}\n\noutput \"subnet_{{ .Project }}_public_c_cidr_block\" {\n  value = aws_subnet.{{ .Project }}_public_c.cidr_block\n}\n\noutput \"subnet_{{ .Project }}_private_a_cidr_block\" {\n  value = aws_subnet.{{ .Project }}_private_a.cidr_block\n}\n\noutput \"subnet_{{ .Project }}_private_c_cidr_block\" {\n  value = aws_subnet.{{ .Project }}_private_c.cidr_block\n}\n"
var _Assetsec62645ba4c90c2d247daee1c3c816eb0fcb9761 = "# VPC Settings\nresource \"aws_vpc\" \"{{ .Project }}\" {\n  cidr_block = \"10.1.0.0/16\"\n\n  enable_dns_hostnames = true\n  enable_dns_support   = true\n  enable_classiclink   = false\n\n  instance_tenancy = \"default\"\n\n  tags = {\n    Name        = \"{{ .Project }}\"\n    Environment = terraform.workspace\n    Workspace   = terraform.workspace\n  }\n}\n\n# Public Subnets Settings\nresource \"aws_subnet\" \"{{ .Project }}_public_a\" {\n  vpc_id            = aws_vpc.{{ .Project }}.id\n  cidr_block        = \"10.1.1.0/24\"\n  availability_zone = lookup(var.availability_zone, \"${terraform.workspace}.a\")\n\n  tags = {\n    Name        = \"{{ .Project }}_public_a\"\n    Environment = terraform.workspace\n    Workspace   = terraform.workspace\n  }\n}\n\nresource \"aws_subnet\" \"{{ .Project }}_public_c\" {\n  vpc_id            = aws_vpc.{{ .Project }}.id\n  cidr_block        = \"10.1.3.0/24\"\n  availability_zone = lookup(var.availability_zone, \"${terraform.workspace}.c\")\n\n  tags = {\n    Name        = \"{{ .Project }}_public_c\"\n    Environment = terraform.workspace\n    Workspace   = terraform.workspace\n  }\n}\n\n# Private Subnets Settings\nresource \"aws_subnet\" \"{{ .Project }}_private_a\" {\n  vpc_id            = aws_vpc.{{ .Project }}.id\n  cidr_block        = \"10.1.100.0/24\"\n  availability_zone = lookup(var.availability_zone, \"${terraform.workspace}.a\")\n\n  tags = {\n    Name        = \"{{ .Project }}_private_a\"\n    Environment = terraform.workspace\n    Workspace   = terraform.workspace\n  }\n}\n\nresource \"aws_subnet\" \"{{ .Project }}_private_c\" {\n  vpc_id            = aws_vpc.{{ .Project }}.id\n  cidr_block        = \"10.1.101.0/24\"\n  availability_zone = lookup(var.availability_zone, \"${terraform.workspace}.c\")\n\n  tags = {\n    Name        = \"{{ .Project }}_private_c\"\n    Environment = terraform.workspace\n    Workspace   = terraform.workspace\n  }\n}\n\n# Routes Table Settings\nresource \"aws_route_table\" \"{{ .Project }}-public-rt\" {\n  vpc_id = aws_vpc.{{ .Project }}.id\n\n  route {\n    cidr_block = \"0.0.0.0/0\"\n    gateway_id = aws_internet_gateway.{{ .Project }}-igw.id\n  }\n\n  tags = {\n    Name        = \"{{ .Project }}_public_rt\"\n    Environment = terraform.workspace\n    Workspace   = terraform.workspace\n  }\n}\n\nresource \"aws_route_table_association\" \"{{ .Project }}-rta1\" {\n  subnet_id      = aws_subnet.{{ .Project }}_public_a.id\n  route_table_id = aws_route_table.{{ .Project }}-public-rt.id\n}\n\nresource \"aws_route_table_association\" \"{{ .Project }}-rta2\" {\n  subnet_id      = aws_subnet.{{ .Project }}_public_c.id\n  route_table_id = aws_route_table.{{ .Project }}-public-rt.id\n}\n\n# DHCP option sets\nresource \"aws_vpc_dhcp_options\" \"{{ .Project }}-dhcp\" {\n  domain_name_servers = [\"AmazonProvidedDNS\"]\n\n  tags = {\n    Name        = \"{{ .Project }}_dhcp\"\n    Environment = terraform.workspace\n    Workspace   = terraform.workspace\n  }\n}\n\nresource \"aws_vpc_dhcp_options_association\" \"{{ .Project }}-dhcp-association\" {\n  vpc_id          = aws_vpc.{{ .Project }}.id\n  dhcp_options_id = aws_vpc_dhcp_options.{{ .Project }}-dhcp.id\n}\n\n# Internet Gateway Settings\nresource \"aws_internet_gateway\" \"{{ .Project }}-igw\" {\n  vpc_id = aws_vpc.{{ .Project }}.id\n\n  tags = {\n    Name        = \"{{ .Project }}_igw\"\n    Environment = terraform.workspace\n    Workspace   = terraform.workspace\n  }\n}\n"
var _Assetsb5f9fb9ed3d9e71486de02e118fabe6d36fbf665 = "resource \"aws_alb\" \"{{ .Project }}\" {\n  name            = \"{{ .Project }}-fargate-alb\"\n  security_groups = [var.sg_{{ .Project }}_alb_id]\n\n  subnets = [\n    var.vpc_subnet_{{ .Project }}_public_a_id,\n    var.vpc_subnet_{{ .Project }}_public_c_id,\n  ]\n\n  internal                   = false\n  enable_deletion_protection = false\n\n  tags = {\n    Environment = terraform.workspace\n  }\n}\n\nresource \"aws_alb_listener\" \"{{ .Project }}_http\" {\n  load_balancer_arn = aws_alb.{{ .Project }}.arn\n  port              = \"8000\"\n  protocol          = \"HTTP\"\n\n  default_action {\n    type = \"fixed-response\"\n\n    fixed_response {\n      content_type = \"text/plain\"\n      message_body = \"No API launched or Misspelling.\"\n      status_code  = \"404\"\n    }\n  }\n}\n\ndata \"aws_iam_role\" \"ecs_service_autoscaling\" {\n  name = \"AWSServiceRoleForApplicationAutoScaling_ECSService\"\n}\n"
var _Assets02ee67c7712efb78754afe499bd4084638a306e3 = "resource \"aws_ecs_cluster\" \"{{ .Project }}_cluster\" {\n  name = \"{{ .Project }}\"\n\n  tags = {\n    Name        = \"{{ .Project }}\"\n    Environment = terraform.workspace\n  }\n}\n"
var _Assets788a43e2798f9ce7137842580efbe41a7bca66c2 = "variable \"vpc_vpc_{{ .Project }}_id\" {}\nvariable \"vpc_subnet_{{ .Project }}_public_a_id\" {}\nvariable \"vpc_subnet_{{ .Project }}_public_c_id\" {}\n\nvariable \"sg_{{ .Project }}_alb_id\" {}\n\nvariable \"ecs_{{ .Project }}_cluster\" {}\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"resource"}, "/resource": []string{}, "/resource/tmpl": []string{}, "/resource/tmpl/terraform": []string{}, "/resource/tmpl/terraform/modules": []string{}, "/resource/tmpl/terraform/modules/alb": []string{"variables.tf.tmpl", "api.tf.tmpl", "main.tf.tmpl"}, "/resource/tmpl/terraform/modules/ecr": []string{"policy.json.tmpl", "api.tf.tmpl"}, "/resource/tmpl/terraform/modules/ecs": []string{"outputs.tf.tmpl", "main.tf.tmpl"}, "/resource/tmpl/terraform/modules/iam": []string{"ecs.tf.tmpl"}, "/resource/tmpl/terraform/modules/sg": []string{"outputs.tf.tmpl", "variables.tf.tmpl", "main.tf.tmpl"}, "/resource/tmpl/terraform/modules/vpc": []string{"outputs.tf.tmpl", "main.tf.tmpl"}}, map[string]*assets.File{
	"/resource/tmpl/terraform/modules/sg": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/sg",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1564179044, 1564179044218696158),
		Data:     nil,
	}, "/resource/tmpl/terraform/modules/sg/variables.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/sg/variables.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1564179044, 1564179044220000000),
		Data:     []byte(_Assetse65d182e7adf57e01416dca18760fe8cb0a96886),
	}, "/resource/tmpl/terraform/modules/vpc/outputs.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/vpc/outputs.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1564146199, 1564146199421000000),
		Data:     []byte(_Assetsc87eece7149baef6e249d528de3f9980d6e9aacf),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1564333093, 1564333093434021210),
		Data:     nil,
	}, "/resource/tmpl/terraform/modules/ecr": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/ecr",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1565845014, 1565845014061470207),
		Data:     nil,
	}, "/resource/tmpl/terraform/modules/ecr/policy.json.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/ecr/policy.json.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1564328677, 1564328677575000000),
		Data:     []byte(_Assetsbacbea7a9bc4dcf5062d63ac8c05ff548dde1924),
	}, "/resource/tmpl/terraform/modules/ecr/api.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/ecr/api.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1565845013, 1565845013947329636),
		Data:     []byte(_Assets15a142d2a032e7816f336c2df6d25fce4ad65b2f),
	}, "/resource/tmpl/terraform/modules/ecs/outputs.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/ecs/outputs.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1564179132, 1564179132255000000),
		Data:     []byte(_Assets335e2f5adb3f568fa4e3da8ecef12ec0779a5c06),
	}, "/resource/tmpl/terraform/modules/vpc/main.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/vpc/main.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1564146163, 1564146163800177847),
		Data:     []byte(_Assetsec62645ba4c90c2d247daee1c3c816eb0fcb9761),
	}, "/resource": &assets.File{
		Path:     "/resource",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1564145870, 1564145870383245725),
		Data:     nil,
	}, "/resource/tmpl/terraform/modules/alb/variables.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/alb/variables.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1564321318, 1564321318092000000),
		Data:     []byte(_Assets788a43e2798f9ce7137842580efbe41a7bca66c2),
	}, "/resource/tmpl/terraform/modules/alb/main.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/alb/main.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1564179271, 1564179271064000000),
		Data:     []byte(_Assetsb5f9fb9ed3d9e71486de02e118fabe6d36fbf665),
	}, "/resource/tmpl/terraform/modules/ecs/main.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/ecs/main.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1564179111, 1564179111535000000),
		Data:     []byte(_Assets02ee67c7712efb78754afe499bd4084638a306e3),
	}, "/resource/tmpl/terraform/modules/ecs": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/ecs",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1564179132, 1564179132254876058),
		Data:     nil,
	}, "/resource/tmpl/terraform/modules/iam": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/iam",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1564179069, 1564179069392485687),
		Data:     nil,
	}, "/resource/tmpl/terraform/modules/iam/ecs.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/iam/ecs.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1564179069, 1564179069393000000),
		Data:     []byte(_Assetsffc03a188b966e5e41bdff75413f26d12cea7441),
	}, "/resource/tmpl": &assets.File{
		Path:     "/resource/tmpl",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1564145911, 1564145911378901332),
		Data:     nil,
	}, "/resource/tmpl/terraform": &assets.File{
		Path:     "/resource/tmpl/terraform",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1564146764, 1564146764279206481),
		Data:     nil,
	}, "/resource/tmpl/terraform/modules": &assets.File{
		Path:     "/resource/tmpl/terraform/modules",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1564329060, 1564329060440874074),
		Data:     nil,
	}, "/resource/tmpl/terraform/modules/alb": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/alb",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1564321318, 1564321318092000289),
		Data:     nil,
	}, "/resource/tmpl/terraform/modules/alb/api.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/alb/api.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1566317872, 1566317872103734408),
		Data:     []byte(_Assets72af51fef0d5e6fe9c11b8bd19b758272de17fbb),
	}, "/resource/tmpl/terraform/modules/sg/main.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/sg/main.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1564178995, 1564178995612000000),
		Data:     []byte(_Assets3ff5abfa5c08898738456203761e92760f43576f),
	}, "/resource/tmpl/terraform/modules/vpc": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/vpc",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1566319378, 1566319378396927258),
		Data:     nil,
	}, "/resource/tmpl/terraform/modules/sg/outputs.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/sg/outputs.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1564179020, 1564179020035000000),
		Data:     []byte(_Assetse3d0b8bc4917cfbd1a87d600a89e4f7894dc6e33),
	}}, "")
