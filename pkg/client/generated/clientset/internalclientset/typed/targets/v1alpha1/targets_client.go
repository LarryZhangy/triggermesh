/*
Copyright 2022 TriggerMesh Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"net/http"

	v1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/targets/v1alpha1"
	"github.com/triggermesh/triggermesh/pkg/client/generated/clientset/internalclientset/scheme"
	rest "k8s.io/client-go/rest"
)

type TargetsV1alpha1Interface interface {
	RESTClient() rest.Interface
	AWSComprehendTargetsGetter
	AWSDynamoDBTargetsGetter
	AWSEventBridgeTargetsGetter
	AWSKinesisTargetsGetter
	AWSLambdaTargetsGetter
	AWSS3TargetsGetter
	AWSSNSTargetsGetter
	AWSSQSTargetsGetter
	AzureEventHubsTargetsGetter
	AzureSentinelTargetsGetter
	CloudEventsTargetsGetter
	DatadogTargetsGetter
	ElasticsearchTargetsGetter
	GoogleCloudFirestoreTargetsGetter
	GoogleCloudPubSubTargetsGetter
	GoogleCloudStorageTargetsGetter
	GoogleCloudWorkflowsTargetsGetter
	GoogleSheetTargetsGetter
	HTTPTargetsGetter
	IBMMQTargetsGetter
	JiraTargetsGetter
	KafkaTargetsGetter
	LogzMetricsTargetsGetter
	LogzTargetsGetter
	MongoDBTargetsGetter
	OracleTargetsGetter
	SalesforceTargetsGetter
	SendGridTargetsGetter
	SlackTargetsGetter
	SolaceTargetsGetter
	SplunkTargetsGetter
	TwilioTargetsGetter
	ZendeskTargetsGetter
}

// TargetsV1alpha1Client is used to interact with features provided by the targets.triggermesh.io group.
type TargetsV1alpha1Client struct {
	restClient rest.Interface
}

func (c *TargetsV1alpha1Client) AWSComprehendTargets(namespace string) AWSComprehendTargetInterface {
	return newAWSComprehendTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) AWSDynamoDBTargets(namespace string) AWSDynamoDBTargetInterface {
	return newAWSDynamoDBTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) AWSEventBridgeTargets(namespace string) AWSEventBridgeTargetInterface {
	return newAWSEventBridgeTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) AWSKinesisTargets(namespace string) AWSKinesisTargetInterface {
	return newAWSKinesisTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) AWSLambdaTargets(namespace string) AWSLambdaTargetInterface {
	return newAWSLambdaTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) AWSS3Targets(namespace string) AWSS3TargetInterface {
	return newAWSS3Targets(c, namespace)
}

func (c *TargetsV1alpha1Client) AWSSNSTargets(namespace string) AWSSNSTargetInterface {
	return newAWSSNSTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) AWSSQSTargets(namespace string) AWSSQSTargetInterface {
	return newAWSSQSTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) AzureEventHubsTargets(namespace string) AzureEventHubsTargetInterface {
	return newAzureEventHubsTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) AzureSentinelTargets(namespace string) AzureSentinelTargetInterface {
	return newAzureSentinelTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) CloudEventsTargets(namespace string) CloudEventsTargetInterface {
	return newCloudEventsTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) DatadogTargets(namespace string) DatadogTargetInterface {
	return newDatadogTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) ElasticsearchTargets(namespace string) ElasticsearchTargetInterface {
	return newElasticsearchTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) GoogleCloudFirestoreTargets(namespace string) GoogleCloudFirestoreTargetInterface {
	return newGoogleCloudFirestoreTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) GoogleCloudPubSubTargets(namespace string) GoogleCloudPubSubTargetInterface {
	return newGoogleCloudPubSubTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) GoogleCloudStorageTargets(namespace string) GoogleCloudStorageTargetInterface {
	return newGoogleCloudStorageTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) GoogleCloudWorkflowsTargets(namespace string) GoogleCloudWorkflowsTargetInterface {
	return newGoogleCloudWorkflowsTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) GoogleSheetTargets(namespace string) GoogleSheetTargetInterface {
	return newGoogleSheetTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) HTTPTargets(namespace string) HTTPTargetInterface {
	return newHTTPTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) IBMMQTargets(namespace string) IBMMQTargetInterface {
	return newIBMMQTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) JiraTargets(namespace string) JiraTargetInterface {
	return newJiraTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) KafkaTargets(namespace string) KafkaTargetInterface {
	return newKafkaTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) LogzMetricsTargets(namespace string) LogzMetricsTargetInterface {
	return newLogzMetricsTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) LogzTargets(namespace string) LogzTargetInterface {
	return newLogzTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) MongoDBTargets(namespace string) MongoDBTargetInterface {
	return newMongoDBTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) OracleTargets(namespace string) OracleTargetInterface {
	return newOracleTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) SalesforceTargets(namespace string) SalesforceTargetInterface {
	return newSalesforceTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) SendGridTargets(namespace string) SendGridTargetInterface {
	return newSendGridTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) SlackTargets(namespace string) SlackTargetInterface {
	return newSlackTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) SolaceTargets(namespace string) SolaceTargetInterface {
	return newSolaceTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) SplunkTargets(namespace string) SplunkTargetInterface {
	return newSplunkTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) TwilioTargets(namespace string) TwilioTargetInterface {
	return newTwilioTargets(c, namespace)
}

func (c *TargetsV1alpha1Client) ZendeskTargets(namespace string) ZendeskTargetInterface {
	return newZendeskTargets(c, namespace)
}

// NewForConfig creates a new TargetsV1alpha1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*TargetsV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new TargetsV1alpha1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*TargetsV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &TargetsV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new TargetsV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *TargetsV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new TargetsV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *TargetsV1alpha1Client {
	return &TargetsV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *TargetsV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
