package config

import (
	"encoding/json"
	"time"

	"github.com/spf13/viper"
)

// cheers up the unparam linter
var _ defaulter = (*CacheConfig)(nil)

// CacheConfig contains fields, which enable and configure
// Flipt's various caching mechanisms.
//
// Currently, flipt support in-memory and redis backed caching.
type CacheConfig struct {
	Enabled bool              `json:"enabled" mapstructure:"enabled"`
	TTL     time.Duration     `json:"ttl,omitempty" mapstructure:"ttl"`
	Backend CacheBackend      `json:"backend,omitempty" mapstructure:"backend"`
	Memory  MemoryCacheConfig `json:"memory,omitempty" mapstructure:"memory"`
	Redis   RedisCacheConfig  `json:"redis,omitempty" mapstructure:"redis"`
}

func (c *CacheConfig) setDefaults(v *viper.Viper) error {
	v.SetDefault("cache", map[string]any{
		"enabled": false,
		"backend": CacheMemory,
		"ttl":     1 * time.Minute,
		"redis": map[string]any{
			"host":     "localhost",
			"port":     6379,
			"password": "",
			"db":       0,
		},
		"memory": map[string]any{
			"enabled":           false, // deprecated (see below)
			"eviction_interval": 5 * time.Minute,
		},
	})

	return nil
}

// CacheBackend is either memory or redis
type CacheBackend uint8

func (c CacheBackend) String() string {
	return cacheBackendToString[c]
}

func (c CacheBackend) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

const (
	_ CacheBackend = iota
	// CacheMemory ...
	CacheMemory
	// CacheRedis ...
	CacheRedis
)

var (
	cacheBackendToString = map[CacheBackend]string{
		CacheMemory: "memory",
		CacheRedis:  "redis",
	}

	stringToCacheBackend = map[string]CacheBackend{
		"memory": CacheMemory,
		"redis":  CacheRedis,
	}
)

// MemoryCacheConfig contains fields, which configure in-memory caching.
type MemoryCacheConfig struct {
	EvictionInterval time.Duration `json:"evictionInterval,omitempty" mapstructure:"eviction_interval"`
}

// RedisCacheConfig contains fields, which configure the connection
// credentials for redis backed caching.
type RedisCacheConfig struct {
	Host            string        `json:"host,omitempty" mapstructure:"host"`
	Port            int           `json:"port,omitempty" mapstructure:"port"`
	RequireTLS      bool          `json:"requireTLS" mapstructure:"require_tls"`
	Password        string        `json:"password,omitempty" mapstructure:"password"`
	DB              int           `json:"db,omitempty" mapstructure:"db"`
	PoolSize        int           `json:"poolSize" mapstructure:"pool_size"`
	MinIdleConn     int           `json:"minIdleConn" mapstructure:"min_idle_conn"`
	ConnMaxIdleTime time.Duration `json:"connMaxIdleTime" mapstructure:"conn_max_idle_time"`
	NetTimeout      time.Duration `json:"netTimeout" mapstructure:"net_timeout"`
}
