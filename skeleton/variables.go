package skeleton

var baseTemplates = []Template{
	{"/resource/tmpl/terraform/modules/vpc/main.tf.tmpl", "modules/vpc/main.tf"},
	{"/resource/tmpl/terraform/modules/vpc/outputs.tf.tmpl", "modules/vpc/outputs.tf"},
	{"/resource/tmpl/terraform/modules/vpc/variables.tf.tmpl", "modules/vpc/variables.tf"},
	{"/resource/tmpl/terraform/modules/sg/main.tf.tmpl", "modules/sg/main.tf"},
	{"/resource/tmpl/terraform/modules/sg/outputs.tf.tmpl", "modules/sg/outputs.tf"},
	{"/resource/tmpl/terraform/modules/sg/variables.tf.tmpl", "modules/sg/variables.tf"},
	{"/resource/tmpl/terraform/modules/ecs/main.tf.tmpl", "modules/ecs/main.tf"},
	{"/resource/tmpl/terraform/modules/ecs/outputs.tf.tmpl", "modules/ecs/outputs.tf"},
}

var autoScaleAPITemplates = []Template{
	{"/resource/tmpl/terraform/modules/alb/main.tf.tmpl", "modules/alb/main.tf"},
	{"/resource/tmpl/terraform/modules/alb/api.tf.tmpl", "modules/alb/{{ .APIName }}.tf"},
	{"/resource/tmpl/terraform/modules/alb/variables.tf.tmpl", "modules/alb/variables.tf"},
	{"/resource/tmpl/terraform/modules/ecr/api.tf.tmpl", "modules/ecr/{{ .APIName }}.tf"},
	{"/resource/tmpl/terraform/modules/ecr/policy.json.tmpl", "modules/ecr/policy.json"},
}
