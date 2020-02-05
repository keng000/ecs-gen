package skeleton

//go:generate go-assets-builder -p skeleton resource/ -o assets.go

var baseTemplates = []Template{
	{"/resource/tmpl/terraform/modules/vpc/main.tf.tmpl", "modules/aws/vpc/main.tf"},
	{"/resource/tmpl/terraform/modules/vpc/variables.tf.tmpl", "modules/aws/vpc/variables.tf"},
	{"/resource/tmpl/terraform/modules/vpc/outputs.tf.tmpl", "modules/aws/vpc/outputs.tf"},
	{"/resource/tmpl/terraform/modules/sg/main.tf.tmpl", "modules/aws/sg/main.tf"},
	{"/resource/tmpl/terraform/modules/sg/outputs.tf.tmpl", "modules/aws/sg/outputs.tf"},
	{"/resource/tmpl/terraform/modules/sg/variables.tf.tmpl", "modules/aws/sg/variables.tf"},
	{"/resource/tmpl/terraform/modules/ecs/main.tf.tmpl", "modules/aws/ecs/main.tf"},
	{"/resource/tmpl/terraform/modules/ecs/outputs.tf.tmpl", "modules/aws/ecs/outputs.tf"},
	{"/resource/tmpl/terraform/modules/iam/ecs.tf.tmpl", "modules/aws/iam/ecs.tf"},
}

var autoScaleAPITemplates = []Template{
	{"/resource/tmpl/terraform/modules/alb/main.tf.tmpl", "modules/aws/alb/main.tf"},
	{"/resource/tmpl/terraform/modules/alb/api.tf.tmpl", "modules/aws/alb/{{ .APIName }}.tf"},
	{"/resource/tmpl/terraform/modules/alb/variables.tf.tmpl", "modules/aws/alb/variables.tf"},
	{"/resource/tmpl/terraform/modules/ecr/api.tf.tmpl", "modules/aws/ecr/{{ .APIName }}.tf"},
	{"/resource/tmpl/terraform/modules/ecr/policy.json.tmpl", "modules/aws/ecr/policy.json"},
}

var deploymentTemplates = []Template{
	{"/resource/tmpl/terraform/regions/.gitignore.tmpl", "aws/{{ .Region }}/.gitignore"},
	{"/resource/tmpl/terraform/regions/aws.tf.tmpl", "aws/{{ .Region }}/aws.tf"},
	{"/resource/tmpl/terraform/regions/backend.tf.tmpl", "aws/{{ .Region }}/backend.tf"},
	{"/resource/tmpl/terraform/regions/modules.tf.tmpl", "aws/{{ .Region }}/modules.tf"},
	{"/resource/tmpl/terraform/regions/variables.tf.tmpl", "aws/{{ .Region }}/variables.tf"},
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
