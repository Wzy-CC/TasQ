package config

import "crypto/tls"

// Config holds all configuration for our program
type Config struct {
	Broker                  string           `yaml:"broker" envconfig:"BROKER"`
	Lock                    string           `yaml:"lock" envconfig:"LOCK"`
	MultipleBrokerSeparator string           `yaml:"multiple_broker_separator" envconfig:"MULTIPLE_BROKEN_SEPARATOR"`
	DefaultQueue            string           `yaml:"default_queue" envconfig:"DEFAULT_QUEUE"`
	ResultBackend           string           `yaml:"result_backend" envconfig:"RESULT_BACKEND"`
	ResultsExpireIn         int              `yaml:"results_expire_in" envconfig:"RESULTS_EXPIRE_IN"`
	AMQP                    *AMQPConfig      `yaml:"amqp"`
	SQS                     *SQSConfig       `yaml:"sqs"`
	Redis                   *RedisConfig     `yaml:"redis"`
	GCPPubSub               *GCPPubSubConfig `yaml:"-" ignored:"true"`
	MongoDB                 *MongoDBConfig   `yaml:"-" ignored:"true"`
	TLSConfig               *tls.Config
	// NoUnixSignals - when set disables signal handling in machinery
	NoUnixSignals bool            `yaml:"no_unix_signals" envconfig:"NO_UNIX_SIGNALS"`
	DynamoDB      *DynamoDBConfig `yaml:"dynamodb"`
}
