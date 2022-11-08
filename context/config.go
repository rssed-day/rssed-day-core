package context

import (
	"encoding/json"
	pb "github.com/rssed-day/rssed-day-core/grpc/protos/pipeline"
	"github.com/rssed-day/rssed-day-core/plugins/aggregators"
	"github.com/rssed-day/rssed-day-core/plugins/inputs"
	"github.com/rssed-day/rssed-day-core/plugins/outputs"
	"github.com/rssed-day/rssed-day-core/plugins/processors"
	"gopkg.in/yaml.v3"
	"os"
)

// Preference -
type Preference struct {
	Debug bool
	// TODO: logger
}

// Config -
type Config struct {
	Preference           *Preference                    `json:"preference" yaml:"preference"`
	InputConfigs         []inputs.InputConfig           `json:"inputs" yaml:"inputs"`
	PreProcessorConfigs  []processors.ProcessorConfig   `json:"pre-processors" yaml:"pre-processors"`
	AggregatorConfigs    []aggregators.AggregatorConfig `json:"aggregators" yaml:"aggregators"`
	PostProcessorConfigs []processors.ProcessorConfig   `json:"post-processors" yaml:"post-processors"`
	OutputConfigs        []outputs.OutputConfig         `json:"outputs" yaml:"outputs"`
}

// ConfigFactory -
type ConfigFactory interface {
	Config() (*Config, error)
}

// BaseConfigFactory -
type BaseConfigFactory struct {
	config *Config
}

// NewBaseConfigFactory -
func NewBaseConfigFactory(config *Config) *BaseConfigFactory {
	return &BaseConfigFactory{config: config}
}

// Config -
func (b *BaseConfigFactory) Config() (*Config, error) {
	return b.config, nil
}

// FileConfigFactory -
type FileConfigFactory struct {
	path string
}

// NewFileConfigFactory -
func NewFileConfigFactory(path string) *FileConfigFactory {
	return &FileConfigFactory{
		path: path,
	}
}

// Config -
func (f *FileConfigFactory) Config() (*Config, error) {
	bts, err := os.ReadFile(f.path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err = yaml.Unmarshal(bts, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// ProtoConfigFactory -
type ProtoConfigFactory struct {
	config *pb.Config
}

// NewProtoConfigFactory -
func NewProtoConfigFactory(config *pb.Config) *ProtoConfigFactory {
	return &ProtoConfigFactory{
		config: config,
	}
}

// Config -
func (f *ProtoConfigFactory) Config() (*Config, error) {
	var cfg Config

	bts, err := json.Marshal(f.config)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bts, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
