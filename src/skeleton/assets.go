package skeleton

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets2a428c505c470f618c496372adfe27b70b919fb5 = "variable \"availability-zone\" {}"
var _Assets72af51fef0d5e6fe9c11b8bd19b758272de17fbb = "resource \"aws_alb_target_group\" \"{{ .APIName }}\" {\n  name        = \"{{ .Project }}-{{ .APIName }}-tg\"\n  port        = 8000\n  protocol    = \"HTTP\"\n  vpc_id      = var.vpc-vpc-{{ .Project }}-id\n  target_type = \"ip\"\n\n  health_check {\n    interval            = lookup(var.load-balancer-rule, \"{{ .APIName }}.health-check-interval\")\n    path                = \"/health\"\n    port                = 8000\n    protocol            = \"HTTP\"\n    timeout             = lookup(var.load-balancer-rule, \"{{ .APIName }}.health-check-timeout\")\n    healthy_threshold   = lookup(var.load-balancer-rule, \"{{ .APIName }}.healthy-threshold\")\n    unhealthy_threshold = lookup(var.load-balancer-rule, \"{{ .APIName }}.unhealthy-threshold\")\n    matcher             = 200\n  }\n}\n\nresource \"aws_alb_listener_rule\" \"{{ .APIName }}\" {\n  listener_arn = aws_alb_listener.{{ .Project }}-http.arn\n\n  action {\n    type             = \"forward\"\n    target_group_arn = aws_alb_target_group.{{ .APIName }}.arn\n  }\n\n  condition {\n    path_pattern {\n      values = [\"/{{ .APIName }}\"]\n    }\n  }\n}\n\n# ## Auto Scaling\n# resource \"aws_appautoscaling_target\" \"autoscale-{{ .APIName }}\" {\n#   service_namespace  = \"ecs\"\n#   resource_id        = \"service/${var.ecs-{{ .Project }}-cluster.name}/{{ .APIName }}\"\n#   scalable_dimension = \"ecs:service:DesiredCount\"\n#   role_arn           = data.aws_iam_role.ecs-service-autoscaling.arn\n#   min_capacity       = lookup(var.load-balancer-rule, \"{{ .APIName }}.min-capacity\")\n#   max_capacity       = lookup(var.load-balancer-rule, \"{{ .APIName }}.max-capacity\")\n# }\n# \n# resource \"aws_appautoscaling_policy\" \"{{ .APIName }}-scale-up\" {\n#   name               = \"{{ .APIName }}-scale-up\"\n#   service_namespace  = \"ecs\"\n#   resource_id        = \"service/${var.ecs-{{ .Project }}-cluster.name}/{{ .APIName }}\"\n#   scalable_dimension = \"ecs:service:DesiredCount\"\n# \n#   step_scaling_policy_configuration {\n#     adjustment_type         = \"ChangeInCapacity\"\n#     cooldown                = lookup(var.load-balancer-rule, \"{{ .APIName }}.scale-up-cooldown\")\n#     metric_aggregation_type = \"Average\"\n# \n#     step_adjustment {\n#       metric_interval_lower_bound = 0\n#       scaling_adjustment          = 1\n#     }\n#   }\n# \n#   depends_on = [\"aws_appautoscaling_target.autoscale-{{ .APIName }}\"]\n# }\n# \n# resource \"aws_appautoscaling_policy\" \"{{ .APIName }}-scale-down\" {\n#   name               = \"{{ .APIName }}-scale-down\"\n#   service_namespace  = \"ecs\"\n#   resource_id        = \"service/${var.ecs-{{ .Project }}-cluster.name}/{{ .APIName }}\"\n#   scalable_dimension = \"ecs:service:DesiredCount\"\n# \n#   step_scaling_policy_configuration {\n#     adjustment_type         = \"ChangeInCapacity\"\n#     cooldown                = lookup(var.load-balancer-rule, \"{{ .APIName }}.scale-down-cooldown\")\n#     metric_aggregation_type = \"Average\"\n# \n#     step_adjustment {\n#       metric_interval_upper_bound = 0\n#       scaling_adjustment          = -1\n#     }\n#   }\n# \n#   depends_on = [\"aws_appautoscaling_target.autoscale-{{ .APIName }}\"]\n# }\n# \n# resource \"aws_cloudwatch_metric_alarm\" \"{{ .APIName }}-cpu-high\" {\n#   alarm_name          = \"{{ .APIName }}-cpu-high\"\n#   comparison_operator = \"GreaterThanOrEqualToThreshold\"\n#   evaluation_periods  = \"2\"\n#   metric_name         = \"CPUUtilization\"\n#   namespace           = \"AWS/ECS\"\n#   period              = \"60\"\n#   statistic           = lookup(var.load-balancer-rule, \"{{ .APIName }}.cpu-high-statistic\")\n#   threshold           = lookup(var.load-balancer-rule, \"{{ .APIName }}.cpu-high-threshold\")\n# \n#   dimensions = {\n#     ClusterName = var.ecs-{{ .Project }}-cluster.name\n#     ServiceName = \"{{ .APIName }}\"\n#   }\n# \n#   alarm_actions = [\"${aws_appautoscaling_policy.{{ .APIName }}-scale-up.arn}\"]\n# }\n# \n# resource \"aws_cloudwatch_metric_alarm\" \"{{ .APIName }}-cpu-low\" {\n#   alarm_name          = \"{{ .APIName }}-cpu-low\"\n#   comparison_operator = \"LessThanOrEqualToThreshold\"\n#   evaluation_periods  = \"2\"\n#   metric_name         = \"CPUUtilization\"\n#   namespace           = \"AWS/ECS\"\n#   period              = \"60\"\n#   statistic           = lookup(var.load-balancer-rule, \"{{ .APIName }}.cpu-low-statistic\")\n#   threshold           = lookup(var.load-balancer-rule, \"{{ .APIName }}.cpu-low-threshold\")\n# \n#   dimensions = {\n#     ClusterName = var.ecs-{{ .Project }}-cluster.name\n#     ServiceName = \"{{ .APIName }}\"\n#   }\n# \n#   alarm_actions = [\"${aws_appautoscaling_policy.{{ .APIName }}-scale-down.arn}\"]\n# } "
var _Assets02ee67c7712efb78754afe499bd4084638a306e3 = "resource \"aws_ecs_cluster\" \"{{ .Project }}-cluster\" {\n  name = \"{{ .Project }}\"\n\n  tags = {\n    Name        = \"{{ .Project }}\"\n    Environment = terraform.workspace\n  }\n}\n"
var _Assetsffc03a188b966e5e41bdff75413f26d12cea7441 = "resource \"aws_iam_role\" \"{{ .Project }}-task-exec-role\" {\n  name = \"{{ .Project }}-${terraform.workspace}-task-exec-role\"\n  path = \"/system/\"\n\n  assume_role_policy = <<EOF\n{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Sid\": \"1\",\n      \"Effect\": \"Allow\",\n      \"Action\": \"sts:AssumeRole\",\n      \"Principal\": {\n        \"Service\": \"ecs-tasks.amazonaws.com\"\n      }\n    }\n  ]\n}\nEOF\n}\n\nresource \"aws_iam_role_policy\" \"{{ .Project }}-task-exec-policy\" {\n  name = \"{{ .Project }}-${terraform.workspace}-task-exec-policy\"\n  role = aws_iam_role.{{ .Project }}-task-exec-role.id\n\n  policy = <<EOF\n{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Effect\": \"Allow\",\n      \"Action\": [\n        \"ecr:GetAuthorizationToken\",\n        \"ecr:BatchCheckLayerAvailability\",\n        \"ecr:GetDownloadUrlForLayer\",\n        \"ecr:BatchGetImage\",\n        \"logs:CreateLogStream\",\n        \"logs:PutLogEvents\",\n\n        \"ssm:GetParameters\",\n        \"secretsmanager:GetSecretValue\",\n        \"kms:Decrypt\",\n\n        \"s3:*\"\n      ],\n      \"Resource\": \"*\"\n    }\n  ]\n}\nEOF\n}\n"
var _Assetse65d182e7adf57e01416dca18760fe8cb0a96886 = "variable \"vpc-subnet-{{ .Project }}-private-a-cidr-block\" {}\nvariable \"vpc-subnet-{{ .Project }}-private-c-cidr-block\" {}\nvariable \"vpc-subnet-{{ .Project }}-public-a-cidr-block\" {}\nvariable \"vpc-subnet-{{ .Project }}-public-c-cidr-block\" {}\nvariable \"vpc-vpc-{{ .Project }}-id\" {}\n"
var _Assets582407d3914c3d8ea320e5cc5f92672f35fadbe1 = ".terraform"
var _Assetsce9456e7bd1fa1456c982cf65f299df2fd2e1767 = "provider \"aws\" {\n  version = \"~> 2.5\"\n  region  = \"{{ .Region }}\"\n}"
var _Assets949f1ec398d11b3847420f5ac22babc0ac75614b = "# update belong to your configuration\nvariable \"availability-zone\" {\n  type = map(string)\n  default = {\n    \"development.a\" = \"\"\n    \"development.c\" = \"\"\n\n    \"production.a\" = \"\"\n    \"production.c\" = \"\"\n  }\n}\n\nvariable \"load-balancer-rule\" {\n  type = map(string)\n  default = {\n    {{ .LBRule }}\n  }\n}"
var _Assets013f2e25a5383bd1cb863dfcd6190df942f8c395 = "terraform {\n  required_version = \">= 0.12.0\"\n\n  # backend \"s3\" {\n  #   region = \"{{ .Region }}\"\n  #   bucket = \"\"\n  #   key    = \"\"\n  # }\n}"
var _Assetsbacbea7a9bc4dcf5062d63ac8c05ff548dde1924 = "{\n    \"rules\": [\n        {\n            \"rulePriority\": 1,\n            \"description\": \"Expire images older than 30 days\",\n            \"selection\": {\n                \"tagStatus\": \"untagged\",\n                \"countType\": \"sinceImagePushed\",\n                \"countUnit\": \"days\",\n                \"countNumber\": 30\n            },\n            \"action\": {\n                \"type\": \"expire\"\n            }\n        }\n    ]\n}"
var _Assets335e2f5adb3f568fa4e3da8ecef12ec0779a5c06 = "output \"{{ .Project }}-cluster\" {\n  value = aws_ecs_cluster.{{ .Project }}-cluster\n}\n"
var _Assets0884b28cfe9ca3f1bfaa1adac7911cbd4734ac30 = "module \"vpc\" {\n  source = \"../../modules/aws/vpc\"\n\n  availability-zone = var.availability-zone\n}\n\nmodule \"sg\" {\n  source = \"../../modules/aws/sg\"\n\n  vpc-subnet-{{ .Project }}-private-a-cidr-block = module.vpc.subnet-{{ .Project }}-private-a-cidr-block\n  vpc-subnet-{{ .Project }}-private-c-cidr-block = module.vpc.subnet-{{ .Project }}-private-c-cidr-block\n  vpc-subnet-{{ .Project }}-public-a-cidr-block  = module.vpc.subnet-{{ .Project }}-public-a-cidr-block\n  vpc-subnet-{{ .Project }}-public-c-cidr-block  = module.vpc.subnet-{{ .Project }}-public-c-cidr-block\n  vpc-vpc-{{ .Project }}-id                      = module.vpc.vpc-{{ .Project }}-id\n}\n\nmodule \"alb\" {\n  source = \"../../modules/aws/alb\"\n\n  vpc-vpc-{{ .Project }}-id             = module.vpc.vpc-{{ .Project }}-id\n  sg-{{ .Project }}-alb-id              = module.sg.{{ .Project }}-alb-id\n  vpc-subnet-{{ .Project }}-public-a-id = module.vpc.subnet-{{ .Project }}-public-a-id\n  vpc-subnet-{{ .Project }}-public-c-id = module.vpc.subnet-{{ .Project }}-public-c-id\n  ecs-{{ .Project }}-cluster            = module.ecs.{{ .Project }}-cluster\n  load-balancer-rule                    = var.load-balancer-rule\n}\n\nmodule \"ecs\" {\n  source = \"../../modules/aws/ecs\"\n}\n\nmodule \"ecr\" {\n  source = \"../../modules/aws/ecr\"\n}\n\nmodule \"iam\" {\n  source = \"../../modules/aws/iam\"\n}\n\n"
var _Assetsb5f9fb9ed3d9e71486de02e118fabe6d36fbf665 = "resource \"aws_alb\" \"{{ .Project }}\" {\n  name            = \"{{ .Project }}-fargate-alb\"\n  security_groups = [var.sg-{{ .Project }}-alb-id]\n\n  subnets = [\n    var.vpc-subnet-{{ .Project }}-public-a-id,\n    var.vpc-subnet-{{ .Project }}-public-c-id,\n  ]\n\n  internal                   = false\n  enable_deletion_protection = false\n\n  tags = {\n    Environment = terraform.workspace\n  }\n}\n\nresource \"aws_alb_listener\" \"{{ .Project }}-http\" {\n  load_balancer_arn = aws_alb.{{ .Project }}.arn\n  port              = \"8000\"\n  protocol          = \"HTTP\"\n\n  default_action {\n    type = \"fixed-response\"\n\n    fixed_response {\n      content_type = \"text/plain\"\n      message_body = \"No API launched or Misspelling.\"\n      status_code  = \"404\"\n    }\n  }\n}\n\ndata \"aws_iam_role\" \"ecs-service-autoscaling\" {\n  name = \"AWSServiceRoleForApplicationAutoScaling_ECSService\"\n}\n"
var _Assets15a142d2a032e7816f336c2df6d25fce4ad65b2f = "resource \"aws_ecr_repository\" \"{{ .APIName }}\" {\n  name = \"{{ .Project }}/{{ .APIName }}\"\n}\n\nresource \"aws_ecr_lifecycle_policy\" \"{{ .APIName }}\" {\n  repository = aws_ecr_repository.{{ .APIName }}.name\n\n  policy = file(\"${path.module}/policy.json\")\n}\n"
var _Assetsc87eece7149baef6e249d528de3f9980d6e9aacf = "output \"vpc-{{ .Project }}-id\" {\n  value = \"${aws_vpc.{{ .Project }}.id}\"\n}\n\noutput \"subnet-{{ .Project }}-public-a-id\" {\n  value = aws_subnet.{{ .Project }}-public-a.id\n}\n\noutput \"subnet-{{ .Project }}-public-c-id\" {\n  value = aws_subnet.{{ .Project }}-public-c.id\n}\n\noutput \"subnet-{{ .Project }}-public-a-cidr-block\" {\n  value = aws_subnet.{{ .Project }}-public-a.cidr_block\n}\n\noutput \"subnet-{{ .Project }}-public-c-cidr-block\" {\n  value = aws_subnet.{{ .Project }}-public-c.cidr_block\n}\n\noutput \"subnet-{{ .Project }}-private-a-cidr-block\" {\n  value = aws_subnet.{{ .Project }}-private-a.cidr_block\n}\n\noutput \"subnet-{{ .Project }}-private-c-cidr-block\" {\n  value = aws_subnet.{{ .Project }}-private-c.cidr_block\n}\n"
var _Assetsec62645ba4c90c2d247daee1c3c816eb0fcb9761 = "# VPC Settings\nresource \"aws_vpc\" \"{{ .Project }}\" {\n  cidr_block = \"10.1.0.0/16\"\n\n  enable_dns_hostnames = true\n  enable_dns_support   = true\n  enable_classiclink   = false\n\n  instance_tenancy = \"default\"\n\n  tags = {\n    Name        = \"{{ .Project }}\"\n    Environment = terraform.workspace\n    Workspace   = terraform.workspace\n  }\n}\n\n# Public Subnets Settings\nresource \"aws_subnet\" \"{{ .Project }}-public-a\" {\n  vpc_id            = aws_vpc.{{ .Project }}.id\n  cidr_block        = \"10.1.1.0/24\"\n  availability_zone = lookup(var.availability-zone, \"${terraform.workspace}.a\")\n\n  tags = {\n    Name        = \"{{ .Project }}-public-a\"\n    Environment = terraform.workspace\n    Workspace   = terraform.workspace\n  }\n}\n\nresource \"aws_subnet\" \"{{ .Project }}-public-c\" {\n  vpc_id            = aws_vpc.{{ .Project }}.id\n  cidr_block        = \"10.1.3.0/24\"\n  availability_zone = lookup(var.availability-zone, \"${terraform.workspace}.c\")\n\n  tags = {\n    Name        = \"{{ .Project }}-public-c\"\n    Environment = terraform.workspace\n    Workspace   = terraform.workspace\n  }\n}\n\n# Private Subnets Settings\nresource \"aws_subnet\" \"{{ .Project }}-private-a\" {\n  vpc_id            = aws_vpc.{{ .Project }}.id\n  cidr_block        = \"10.1.100.0/24\"\n  availability_zone = lookup(var.availability-zone, \"${terraform.workspace}.a\")\n\n  tags = {\n    Name        = \"{{ .Project }}-private-a\"\n    Environment = terraform.workspace\n    Workspace   = terraform.workspace\n  }\n}\n\nresource \"aws_subnet\" \"{{ .Project }}-private-c\" {\n  vpc_id            = aws_vpc.{{ .Project }}.id\n  cidr_block        = \"10.1.101.0/24\"\n  availability_zone = lookup(var.availability-zone, \"${terraform.workspace}.c\")\n\n  tags = {\n    Name        = \"{{ .Project }}-private-c\"\n    Environment = terraform.workspace\n    Workspace   = terraform.workspace\n  }\n}\n\n# Routes Table Settings\nresource \"aws_route_table\" \"{{ .Project }}-public-rt\" {\n  vpc_id = aws_vpc.{{ .Project }}.id\n\n  route {\n    cidr_block = \"0.0.0.0/0\"\n    gateway_id = aws_internet_gateway.{{ .Project }}-igw.id\n  }\n\n  tags = {\n    Name        = \"{{ .Project }}-public-rt\"\n    Environment = terraform.workspace\n    Workspace   = terraform.workspace\n  }\n}\n\nresource \"aws_route_table_association\" \"{{ .Project }}-rta1\" {\n  subnet_id      = aws_subnet.{{ .Project }}-public-a.id\n  route_table_id = aws_route_table.{{ .Project }}-public-rt.id\n}\n\nresource \"aws_route_table_association\" \"{{ .Project }}-rta2\" {\n  subnet_id      = aws_subnet.{{ .Project }}-public-c.id\n  route_table_id = aws_route_table.{{ .Project }}-public-rt.id\n}\n\n# DHCP option sets\nresource \"aws_vpc_dhcp_options\" \"{{ .Project }}-dhcp\" {\n  domain_name_servers = [\"AmazonProvidedDNS\"]\n\n  tags = {\n    Name        = \"{{ .Project }}-dhcp\"\n    Environment = terraform.workspace\n    Workspace   = terraform.workspace\n  }\n}\n\nresource \"aws_vpc_dhcp_options_association\" \"{{ .Project }}-dhcp-association\" {\n  vpc_id          = aws_vpc.{{ .Project }}.id\n  dhcp_options_id = aws_vpc_dhcp_options.{{ .Project }}-dhcp.id\n}\n\n# Internet Gateway Settings\nresource \"aws_internet_gateway\" \"{{ .Project }}-igw\" {\n  vpc_id = aws_vpc.{{ .Project }}.id\n\n  tags = {\n    Name        = \"{{ .Project }}-igw\"\n    Environment = terraform.workspace\n    Workspace   = terraform.workspace\n  }\n}\n"
var _Assets788a43e2798f9ce7137842580efbe41a7bca66c2 = "variable \"vpc-vpc-{{ .Project }}-id\" {}\nvariable \"vpc-subnet-{{ .Project }}-public-a-id\" {}\nvariable \"vpc-subnet-{{ .Project }}-public-c-id\" {}\n\nvariable \"sg-{{ .Project }}-alb-id\" {}\n\nvariable \"ecs-{{ .Project }}-cluster\" {}\n\nvariable \"load-balancer-rule\" {\n    type = map(string)\n}"
var _Assetse3d0b8bc4917cfbd1a87d600a89e4f7894dc6e33 = "output \"{{.Project }}-private-subnet-all-tcp-id\" {\n  value = aws_security_group.{{.Project }}-private-subnet-all-tcp.id\n}\n\noutput \"{{.Project }}-public-subnet-all-tcp-id\" {\n  value = aws_security_group.{{.Project }}-public-subnet-all-tcp.id\n}\n\noutput \"{{.Project }}-alb-id\" {\n  value = aws_security_group.{{.Project }}-alb.id\n}\n\noutput \"{{.Project }}-ecs-id\" {\n  value = aws_security_group.{{.Project }}-ecs.id\n}\n"
var _Assets3ff5abfa5c08898738456203761e92760f43576f = "resource \"aws_security_group\" \"{{ .Project }}-public-subnet-all-tcp\" {\n  vpc_id = var.vpc-vpc-{{ .Project }}-id\n  name   = \"{{ .Project }}-public-subnet-all-tcp\"\n\n  ingress {\n    from_port = 0\n    to_port   = 0\n    protocol  = \"-1\"\n\n    cidr_blocks = [\n      var.vpc-subnet-{{ .Project }}-public-a-cidr-block,\n      var.vpc-subnet-{{ .Project }}-public-c-cidr-block,\n    ]\n  }\n\n  egress {\n    from_port   = 0\n    to_port     = 0\n    protocol    = \"-1\"\n    cidr_blocks = [\"0.0.0.0/0\"]\n  }\n\n  tags = {\n    Name        = \"{{ .Project }}publicsubnetalltcp\"\n    Environment = terraform.workspace\n  }\n}\n\nresource \"aws_security_group\" \"{{ .Project }}-private-subnet-all-tcp\" {\n  vpc_id = var.vpc-vpc-{{ .Project }}-id\n  name   = \"{{ .Project }}-private-subnet-all-tcp\"\n\n  ingress {\n    from_port = 0\n    to_port   = 0\n    protocol  = \"-1\"\n\n    cidr_blocks = [\n      var.vpc-subnet-{{ .Project }}-private-a-cidr-block,\n      var.vpc-subnet-{{ .Project }}-private-c-cidr-block,\n    ]\n  }\n\n  egress {\n    from_port   = 0\n    to_port     = 0\n    protocol    = \"-1\"\n    cidr_blocks = [\"0.0.0.0/0\"]\n  }\n\n  tags = {\n    Name        = \"{{ .Project }}privatesubnetalltcp\"\n    Environment = terraform.workspace\n  }\n}\n\nresource \"aws_security_group\" \"{{ .Project }}-alb\" {\n  vpc_id = var.vpc-vpc-{{ .Project }}-id\n  name   = \"{{ .Project }}-alb\"\n\n  ingress {\n    from_port   = 80\n    to_port     = 80\n    protocol    = \"tcp\"\n    cidr_blocks = [\"0.0.0.0/0\"]\n  }\n\n  ingress {\n    from_port   = 443\n    to_port     = 443\n    protocol    = \"tcp\"\n    cidr_blocks = [\"0.0.0.0/0\"]\n  }\n\n  ingress {\n    from_port   = 8000\n    to_port     = 8000\n    protocol    = \"tcp\"\n    cidr_blocks = [\"0.0.0.0/0\"]\n  }\n\n  egress {\n    from_port   = 0\n    to_port     = 0\n    protocol    = \"-1\"\n    cidr_blocks = [\"0.0.0.0/0\"]\n  }\n\n  tags = {\n    Name        = \"{{ .Project }}alb\"\n    Environment = terraform.workspace\n  }\n}\n\nresource \"aws_security_group\" \"{{ .Project }}-ecs\" {\n  vpc_id = var.vpc-vpc-{{ .Project }}-id\n  name   = \"{{ .Project }}-ecs\"\n\n  ingress {\n    from_port   = 80\n    to_port     = 80\n    protocol    = \"tcp\"\n    cidr_blocks = [\"0.0.0.0/0\"]\n\n    security_groups = [\n      aws_security_group.{{ .Project }}-alb.id,\n    ]\n  }\n\n  ingress {\n    from_port   = 443\n    to_port     = 443\n    protocol    = \"tcp\"\n    cidr_blocks = [\"0.0.0.0/0\"]\n\n    security_groups = [\n      aws_security_group.{{ .Project }}-alb.id,\n    ]\n  }\n\n  ingress {\n    from_port   = 8000\n    to_port     = 8000\n    protocol    = \"tcp\"\n    cidr_blocks = [\"0.0.0.0/0\"]\n\n    security_groups = [\n      aws_security_group.{{ .Project }}-alb.id,\n    ]\n  }\n\n  egress {\n    from_port   = 0\n    to_port     = 0\n    protocol    = \"-1\"\n    cidr_blocks = [\"0.0.0.0/0\"]\n  }\n\n  tags = {\n    Name        = \"{{ .Project }}ecs\"\n    Environment = terraform.workspace\n  }\n}\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"resource"}, "/resource": []string{}, "/resource/tmpl": []string{}, "/resource/tmpl/terraform": []string{}, "/resource/tmpl/terraform/modules": []string{}, "/resource/tmpl/terraform/modules/alb": []string{"variables.tf.tmpl", "api.tf.tmpl", "main.tf.tmpl"}, "/resource/tmpl/terraform/modules/ecr": []string{"policy.json.tmpl", "api.tf.tmpl"}, "/resource/tmpl/terraform/modules/ecs": []string{"outputs.tf.tmpl", "main.tf.tmpl"}, "/resource/tmpl/terraform/modules/iam": []string{"ecs.tf.tmpl"}, "/resource/tmpl/terraform/modules/sg": []string{"outputs.tf.tmpl", "variables.tf.tmpl", "main.tf.tmpl"}, "/resource/tmpl/terraform/modules/vpc": []string{"outputs.tf.tmpl", "variables.tf.tmpl", "main.tf.tmpl"}, "/resource/tmpl/terraform/regions": []string{".gitignore.tmpl", "aws.tf.tmpl", "backend.tf.tmpl", "modules.tf.tmpl", "variables.tf.tmpl"}}, map[string]*assets.File{
	"/resource/tmpl/terraform/modules/iam/ecs.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/iam/ecs.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1580919725, 1580919725399301242),
		Data:     []byte(_Assetsffc03a188b966e5e41bdff75413f26d12cea7441),
	}, "/resource/tmpl/terraform/modules/sg/variables.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/sg/variables.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1580919309, 1580919309159125199),
		Data:     []byte(_Assetse65d182e7adf57e01416dca18760fe8cb0a96886),
	}, "/resource/tmpl/terraform/regions/.gitignore.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/regions/.gitignore.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1580060273, 1580060273079777488),
		Data:     []byte(_Assets582407d3914c3d8ea320e5cc5f92672f35fadbe1),
	}, "/resource/tmpl/terraform/regions/aws.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/regions/aws.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1580060273, 1580060273080132867),
		Data:     []byte(_Assetsce9456e7bd1fa1456c982cf65f299df2fd2e1767),
	}, "/resource/tmpl/terraform/regions/variables.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/regions/variables.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1580921916, 1580921916951613738),
		Data:     []byte(_Assets949f1ec398d11b3847420f5ac22babc0ac75614b),
	}, "/resource/tmpl/terraform/modules/alb": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/alb",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1580060273, 1580060273078431279),
		Data:     nil,
	}, "/resource/tmpl/terraform/modules/iam": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/iam",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1579935670, 1579935670220051984),
		Data:     nil,
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1580063095, 1580063095186659180),
		Data:     nil,
	}, "/resource/tmpl/terraform/regions/backend.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/regions/backend.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1580060273, 1580060273081471401),
		Data:     []byte(_Assets013f2e25a5383bd1cb863dfcd6190df942f8c395),
	}, "/resource/tmpl/terraform/modules/ecr/policy.json.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/ecr/policy.json.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1579935670, 1579935670219133763),
		Data:     []byte(_Assetsbacbea7a9bc4dcf5062d63ac8c05ff548dde1924),
	}, "/resource/tmpl/terraform/modules/ecs/outputs.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/ecs/outputs.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1580918987, 1580918987444343095),
		Data:     []byte(_Assets335e2f5adb3f568fa4e3da8ecef12ec0779a5c06),
	}, "/resource/tmpl/terraform": &assets.File{
		Path:     "/resource/tmpl/terraform",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1580060273, 1580060273079590806),
		Data:     nil,
	}, "/resource/tmpl/terraform/modules/ecr": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/ecr",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1579935670, 1579935670219025106),
		Data:     nil,
	}, "/resource/tmpl/terraform/regions/modules.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/regions/modules.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1580922167, 1580922167865120313),
		Data:     []byte(_Assets0884b28cfe9ca3f1bfaa1adac7911cbd4734ac30),
	}, "/resource/tmpl/terraform/modules/alb/main.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/alb/main.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1580919671, 1580919671746658011),
		Data:     []byte(_Assetsb5f9fb9ed3d9e71486de02e118fabe6d36fbf665),
	}, "/resource/tmpl/terraform/modules/ecr/api.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/ecr/api.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1580926006, 1580926006807926873),
		Data:     []byte(_Assets15a142d2a032e7816f336c2df6d25fce4ad65b2f),
	}, "/resource/tmpl/terraform/modules/ecs": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/ecs",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1579935670, 1579935670219692585),
		Data:     nil,
	}, "/resource/tmpl/terraform/modules/vpc/outputs.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/vpc/outputs.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1580921532, 1580921532196621659),
		Data:     []byte(_Assetsc87eece7149baef6e249d528de3f9980d6e9aacf),
	}, "/resource/tmpl/terraform/regions": &assets.File{
		Path:     "/resource/tmpl/terraform/regions",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1580060273, 1580060273081960191),
		Data:     nil,
	}, "/resource/tmpl/terraform/modules/sg": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/sg",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1579935670, 1579935670221216750),
		Data:     nil,
	}, "/resource/tmpl/terraform/modules/vpc/main.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/vpc/main.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1580922270, 1580922270687270919),
		Data:     []byte(_Assetsec62645ba4c90c2d247daee1c3c816eb0fcb9761),
	}, "/resource/tmpl/terraform/modules": &assets.File{
		Path:     "/resource/tmpl/terraform/modules",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1579935670, 1579935670221422715),
		Data:     nil,
	}, "/resource/tmpl/terraform/modules/alb/variables.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/alb/variables.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1580921621, 1580921621658936695),
		Data:     []byte(_Assets788a43e2798f9ce7137842580efbe41a7bca66c2),
	}, "/resource/tmpl/terraform/modules/sg/outputs.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/sg/outputs.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1580919864, 1580919864991500307),
		Data:     []byte(_Assetse3d0b8bc4917cfbd1a87d600a89e4f7894dc6e33),
	}, "/resource/tmpl/terraform/modules/sg/main.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/sg/main.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1580919807, 1580919807613343518),
		Data:     []byte(_Assets3ff5abfa5c08898738456203761e92760f43576f),
	}, "/resource/tmpl/terraform/modules/vpc/variables.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/vpc/variables.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1580921909, 1580921909566649740),
		Data:     []byte(_Assets2a428c505c470f618c496372adfe27b70b919fb5),
	}, "/resource": &assets.File{
		Path:     "/resource",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1579935670, 1579935670217341679),
		Data:     nil,
	}, "/resource/tmpl": &assets.File{
		Path:     "/resource/tmpl",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1579935670, 1579935670217473640),
		Data:     nil,
	}, "/resource/tmpl/terraform/modules/alb/api.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/alb/api.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1581177510, 1581177510939791519),
		Data:     []byte(_Assets72af51fef0d5e6fe9c11b8bd19b758272de17fbb),
	}, "/resource/tmpl/terraform/modules/ecs/main.tf.tmpl": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/ecs/main.tf.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1580918989, 1580918989523529018),
		Data:     []byte(_Assets02ee67c7712efb78754afe499bd4084638a306e3),
	}, "/resource/tmpl/terraform/modules/vpc": &assets.File{
		Path:     "/resource/tmpl/terraform/modules/vpc",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1580921888, 1580921888708054217),
		Data:     nil,
	}}, "")
