package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Type    string `yaml:"type"`
	Network string `yaml:"network"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Master  struct {
		Snowflake struct {
			NodeID int64 `yaml:"node_id"`
		} `yaml:"snowflake"`
		Registry struct {
			Enabled bool `yaml:"enabled"`
			Expiry  struct {
				ExpireDuration       string `yaml:"expire_duration"`
				ExpiryLimit          int    `yaml:"expiry_limit"`
				ExpiredEvictInterval string `yaml:"expired_evict_interval"`
			} `yaml:"expiry"`
		} `yaml:"registry"`
	} `yaml:"master"`
	Worker struct {
		MasterAddress       string `yaml:"master_address"`
		HeartbeatDuration   string `yaml:"heartbeat_duration"`
		AcquireTaskDuration string `yaml:"acquire_task_duration"`
		Plugin              struct {
			Enabled bool   `yaml:"enabled"`
			Name    string `yaml:"name"`
			Path    string `yaml:"path"`
		} `yaml:"plugin"`
	} `yaml:"worker"`
	S3Endpoint         string `yaml:"s3_endpoint"`
	AccessKey          string `yaml:"access_key"`
	AccessSecret       string `yaml:"access_secret"`
	TaskBucket         string `yaml:"task_bucket"`
	IntermediateBucket string `yaml:"intermediate_bucket"`
	IntermediatePrefix string `yaml:"intermediate_prefix"`
	ResultBucket       string `yaml:"result_bucket"`
	ResultPrefix       string `yaml:"result_prefix"`
}

func Load(filename string, c *Config) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(bytes, c); err != nil {
		return err
	}
	return nil
}
