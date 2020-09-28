package plugins

const (
	configurationAsCodePlugin           = "configuration-as-code:1.38"
	gitPlugin                           = "git:4.2.2"
	jobDslPlugin                        = "job-dsl:1.77"
	kubernetesCredentialsProviderPlugin = "kubernetes-credentials-provider:0.13"
	kubernetesPlugin                    = "kubernetes:1.25.2"
	workflowAggregatorPlugin            = "workflows-aggregator:2.6"
	workflowJobPlugin                   = "workflows-job:2.39"
)

// basePluginsList contains plugins to install by operator.
var basePluginsList = []Plugin{
	Must(New(kubernetesPlugin)),
	Must(New(workflowJobPlugin)),
	Must(New(workflowAggregatorPlugin)),
	Must(New(gitPlugin)),
	Must(New(jobDslPlugin)),
	Must(New(configurationAsCodePlugin)),
	Must(New(kubernetesCredentialsProviderPlugin)),
}

// BasePlugins returns list of plugins to install by operator.
func BasePlugins() []Plugin {
	return basePluginsList
}
