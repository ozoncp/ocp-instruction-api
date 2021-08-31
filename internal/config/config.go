package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Pg_master_conn            string   `yaml:"pg_master_conn"`
	Pg_slave_conn            string   `yaml:"pg_slave_conn"`
	Kafka_addr         []string `yaml:"kafka_addr"`
	Jaeger_addr        string   `yaml:"jaeger_addr"`
	Inserts_chank_size int      `yaml:"inserts_chank_size"`
	Grpc_Listen        string   `yaml:"grpc_listen"`
	Grpc_Endpoint      string   `yaml:"grpc_endpoint"`
	Grpc_Jsongw_Listen string   `yaml:"grpc_jsongw_listen"`
	Metrics_Listen     string   `yaml:"metrics_listen"`
}

var Data Config

func Load() error {
	file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, &Data)
	if err != nil {
		return err
	}

	return nil
}
