/*
Copyright 2018 Blindside Networks

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

package main

import (
	"fmt"
)

type Configuration struct {
	BASE_URL     string
	SALT         string
}

func (p *Plugin) OnConfigurationChange() error {
	var configuration Configuration
	// loads configuration from our config ui page
	err := p.API.LoadPluginConfiguration(&configuration)
	// stores the config in an Atomic.Value place
	p.configuration.Store(&configuration)
	return err
}
func (p *Plugin) config() *Configuration {
	// returns the config file we had stored in Atomic.Value
	return p.configuration.Load().(*Configuration)
}

func (c *Configuration) IsValid() error {
	if len(c.BASE_URL) == 0 {
		return fmt.Errorf("BASE URL is not configured.")
	} else if len(c.SALT) == 0 {
		return fmt.Errorf("SALT is not configured.")
	}

	return nil
}
