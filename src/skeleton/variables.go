package skeleton

var baseTemplates = []Template{
	{"/resource/tmpl/terraform/modules/vpc/main.tf.tmpl", "modules/vpc/main.tf"},
	{"/resource/tmpl/terraform/modules/vpc/outputs.tf.tmpl", "modules/vpc/outputs.tf"},
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

var deploymentTemplates = []Template{
	{"/resource/tmpl/terraform/regions/.gitignore.tmpl", "{{ .Region }}/.gitignore"},
	{"/resource/tmpl/terraform/regions/aws.tf.tmpl", "{{ .Region }}/aws.tf"},
	{"/resource/tmpl/terraform/regions/backend.tf.tmpl", "{{ .Region }}/backend.tf"},
	{"/resource/tmpl/terraform/regions/modules.tf.tmpl", "{{ .Region }}/modules.tf"},
	{"/resource/tmpl/terraform/regions/variables.tf.tmpl", "{{ .Region }}/variables.tf"},
}

var regions = []string{
	"us-east-1",
	"us-east-2",
	"us-west-1",
	"us-west-2",
	"ca-central-1",
	"eu-central-1",
	"eu-west-1",
	"eu-west-2",
	"eu-west-3",
	"ap-northeast-1",
	"ap-northeast-2",
	"ap-northeast-3",
	"ap-southeast-1",
	"ap-southeast-2",
	"ap-south-1",
	"sa-east-1",
}
