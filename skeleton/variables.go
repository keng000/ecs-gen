package skeleton

var baseTemplates = []Template{
	{"skeleton/resource/tmpl/terraform/modules/vpc/main.tf.tmpl", "modules/vpc/main.tf"},
	{"skeleton/resource/tmpl/terraform/modules/vpc/outputs.tf.tmpl", "modules/vpc/outputs.tf"},
	{"skeleton/resource/tmpl/terraform/modules/vpc/variables.tf.tmpl", "modules/vpc/variables.tf"},
	{"skeleton/resource/tmpl/terraform/modules/sg/main.tf.tmpl", "modules/sg/main.tf"},
	{"skeleton/resource/tmpl/terraform/modules/sg/outputs.tf.tmpl", "modules/sg/outputs.tf"},
	{"skeleton/resource/tmpl/terraform/modules/sg/variables.tf.tmpl", "modules/sg/variables.tf"},
	{"skeleton/resource/tmpl/terraform/modules/ecs/main.tf.tmpl", "modules/ecs/main.tf"},
	{"skeleton/resource/tmpl/terraform/modules/ecs/outputs.tf.tmpl", "modules/ecs/outputs.tf"},
}

var autoScaleAPITemplates = []Template{
	{"skeleton/resource/tmpl/terraform/modules/alb/main.tf.tmpl", "modules/alb/main.tf"},
	{"skeleton/resource/tmpl/terraform/modules/alb/api.tf.tmpl", "modules/alb/{{ .APIName }}.tf"},
	{"skeleton/resource/tmpl/terraform/modules/alb/variables.tf.tmpl", "modules/alb/variables.tf"},
	{"skeleton/resource/tmpl/terraform/modules/ecr/api.tf.tmpl", "modules/ecr/{{ .APIName }}.tf"},
	{"skeleton/resource/tmpl/terraform/modules/ecr/policy.json.tmpl", "modules/ecr/policy.json"},
}
