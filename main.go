package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"

	"example.com/ticker/imports/k8s"
	"example.com/ticker/internal/values"
)

type MyChartProps struct {
	cdk8s.ChartProps
}

func NewMyChart(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)
	val := values.Read()

	matchLabels := &map[string]*string{
		"app": jsii.String(val.Name),
		"env": jsii.String(val.Env),
	}

	metadata := &k8s.ObjectMeta{
		Name:      jsii.String(val.Name),
		Namespace: jsii.String(val.Namespace),
		Labels:    matchLabels,
	}

	k8s.NewKubeNamespace(chart, jsii.String("namespace"), &k8s.KubeNamespaceProps{
		Metadata: metadata,
	})

	envNameToValue := map[string]*string{
		"APIKEY": &val.ApiKey,
		"NDAYS":  &val.NumOfDays,
		"SYMBOL": &val.Symbol,
	}

	envValues := make(map[string]*string)
	for envKey, envVal := range envNameToValue {
		envValues[envKey] = envVal
	}

	k8s.NewKubeSecret(chart, jsii.String("env"), &k8s.KubeSecretProps{
		Metadata:   metadata,
		StringData: &envValues,
	})

	k8s.NewKubeDeployment(chart, jsii.String("deployment"), &k8s.KubeDeploymentProps{
		Metadata: metadata,

		Spec: &k8s.DeploymentSpec{
			Replicas: jsii.Number(val.Replicas),
			Selector: &k8s.LabelSelector{
				MatchLabels: matchLabels,
			},

			Template: &k8s.PodTemplateSpec{
				Metadata: metadata,
				Spec: &k8s.PodSpec{
					Containers: &[]*k8s.Container{
						{
							Name:  jsii.String(val.Name),
							Image: jsii.String(val.Image),
							EnvFrom: &[]*k8s.EnvFromSource{
								{
									SecretRef: &k8s.SecretEnvSource{
										Name: metadata.Name,
									},
								},
							},
							Resources: &k8s.ResourceRequirements{
								Requests: &map[string]k8s.Quantity{
									"cpu":    k8s.Quantity_FromString(&val.Resources.Requests.CPU),
									"memory": k8s.Quantity_FromString(&val.Resources.Requests.Memory),
								},
							},
						},
					},
				},
			},
		},
	})

	k8s.NewKubeService(chart, jsii.String("service"), &k8s.KubeServiceProps{
		Metadata: metadata,
		Spec: &k8s.ServiceSpec{
			Selector: matchLabels,
			Ports: &[]*k8s.ServicePort{
				{
					Name: &val.Service.Name,
					Port: &val.Service.Port,
				},
			},
		},
	})
	return chart
}

func main() {
	app := cdk8s.NewApp(nil)
	NewMyChart(app, "ticker", nil)
	app.Synth()
}
