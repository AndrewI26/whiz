package logtar

import (
	"encoding/json"
	"os"
	"fmt"
)

const (
	OneKB = 1024;
	FiveKB = 5 * 1024; 
	TenKB = 10 * 1024;  
	TwentyKB = 20 * 1024; 
	FiftyKB = 50 * 1024; 
	HundredKB = 100 * 1024;

	HalfMB = 512 * 1024;
	OneMB = 1024 * 1024;
	FiveMB = 5 * 1024 * 1024;  
	TenMB = 10 * 1024 * 1024; 
	TwentyMB = 20 * 1024 * 1024;
	FiftyMB = 50 * 1024 * 1024; 
	HundredMB = 100 * 1024 * 1024;

	Minutely = 60;
	Hourly = 60 * Minutely; 
	Daily = 24 * Hourly;
	Weekly = 7 * Daily;
    Monthly = 30 * Daily;
    Yearly = 12 * Monthly;
)

type RollingConfig struct {
	TimeThreshold int64 `json:"timeThreshold"`
	SizeThreshold int64 `json:"sizeThreshold"`
}

// NewRollingConfig constructs a new rolling config struct from a given json config file.
func NewRollingConfig(path string) (*RollingConfig, error) {
	rollingConfig := RollingConfig{}
	file, err := os.Open(path)
	if err != nil {
		return &rollingConfig, fmt.Errorf("failed to open config file %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&rollingConfig)
	if err != nil {
		return &rollingConfig, fmt.Errorf("failed to parse config file %w", err)
	}

	return &rollingConfig, nil
}
