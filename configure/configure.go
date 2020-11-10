/*
Copyright © 2020 Henry Huang <hhh@rutcode.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package configure

import (
	"encoding/json"

	"github.com/go-trellis/common/logger"
	"github.com/go-trellis/trellis/message/proto"

	"github.com/go-trellis/config"
	"github.com/go-trellis/node"

	// cobra in main
	_ "github.com/spf13/cobra"
)

// Config project configure
type Config struct {
	Project *Project `yaml:"project" json:"project"`
}

// Project project config info
type Project struct {
	Logger LoggerConfig `yaml:"logger" json:"logger"`

	Services map[string]*Service `yaml:"services" json:"services"`

	Registries map[string]*Registry `yaml:"registries" json:"registries"`
}

// LoggerConfig logger's configure
type LoggerConfig struct {
	Level      logger.Level `yaml:"level" json:"level"`
	ChanBuffer int          `yaml:"chan_buffer" json:"chan_buffer"`
	Separator  string       `yaml:"separator" json:"separator"`
}

// Service service info
type Service struct {
	Service proto.Service `yaml:",inline" json:",inline"`

	Options config.Options `yaml:"options" json:"options"`

	Registry *ServiceRegistryOptions `yaml:"registry" json:"registry"`
}

type ServiceRegistryOptions struct {
	Name     string `yaml:"name" json:"name"`
	Domain   string `yaml:"domain" json:"domain"`
	Protocol string `yaml:"protocol" json:"protocol"`
	Weight   uint32 `yaml:"weight" json:"weight"`
}

// Registry run configure
type Registry struct {
	Type string `yaml:"type" json:"type"`

	Options config.Options `yaml:"options" json:"options"`

	Watchers []*Watcher `yaml:"watchers" json:"watchers"`
}

// RegistService service which should regist into registry
type RegistService struct {
	Service proto.Service `yaml:",inline" json:",inline"`

	Domain   string `yaml:"domain" json:"domain"`
	Protocol string `yaml:"protocal" json:"protocal"`
	Weight   uint32 `yaml:"weight" json:"weight"`
}

func (p *RegistService) String() string {
	bs, _ := json.Marshal(p)
	return string(bs)
}

func ToRegistService(str string) (*RegistService, error) {
	s := &RegistService{}
	err := json.Unmarshal([]byte(str), s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

type RegistServices []*RegistService

type Watcher struct {
	Service proto.Service `yaml:",inline" json:",inline"`

	LoadBalancing node.Type `yaml:"load_balancing" json:"load_balancing"`
}

// Fullpath fullname
func (p *Watcher) Fullpath() string {
	return p.Service.String()
}
