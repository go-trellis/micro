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

package registry

import (
	"fmt"
	"sync"

	"github.com/go-trellis/trellis/configure"
	"github.com/go-trellis/trellis/internal"
	"github.com/go-trellis/trellis/message/proto"

	"github.com/go-trellis/node"
)

var defaultRegister = &Register{
	newRegistryFuncs: make(map[proto.RegisterType]NewRegistryFunc),
	mapRegistries:    make(map[string]Registry),
	nodeManagers:     make(map[string]node.Manager),
}

// Register 注册机
type Register struct {
	locker sync.RWMutex

	newRegistryFuncs map[proto.RegisterType]NewRegistryFunc
	mapRegistries    map[string]Registry
	nodeManagers     map[string]node.Manager
}

// Regist 注册注册机
func Regist(typ proto.RegisterType, fn NewRegistryFunc) error {
	_, ok := defaultRegister.newRegistryFuncs[typ]
	if ok {
		return fmt.Errorf("registry'type (%d) is already exist", typ)
	}
	defaultRegister.newRegistryFuncs[typ] = fn
	return nil
}

// GetNewRegistryFunc 获取注册机生成函数
func GetNewRegistryFunc(typ proto.RegisterType) (NewRegistryFunc, error) {
	r, ok := defaultRegister.newRegistryFuncs[typ]
	if !ok {
		return nil, fmt.Errorf("registry'name (%d) isnot exist", typ)
	}
	return r, nil
}

// RegistService 注册Service
func RegistService(name string, service *configure.RegistService) error {
	r, ok := defaultRegister.mapRegistries[name]
	if !ok {
		return fmt.Errorf("not found service(%s)'s registry (%s)", service.Name, name)
	}

	return r.Regist(service)
}

// NewRegistry 生成注册机
func NewRegistry(name string, opts *RegistOption) error {
	_, ok := defaultRegister.mapRegistries[name]
	if ok {
		return fmt.Errorf("registry: %s is already exist", name)
	}
	fn, err := GetNewRegistryFunc(opts.RegisterType)
	if err != nil {
		return err
	}
	r := fn()

	if err := r.Init(opts); err != nil {
		return err
	}

	defaultRegister.mapRegistries[name] = r
	return nil
}

// NewRegistryWatcher 生成监控者
func NewRegistryWatcher(name string, wConfig *configure.Watcher) error {
	r, ok := defaultRegister.mapRegistries[name]
	if !ok {
		return fmt.Errorf("not found registry: %s", name)
	}

	watcher, err := r.Watcher(wConfig)
	if err != nil {
		return err
	}

	go runWatcher(watcher)
	return nil
}

// Stop 结束
func Stop() {
	for _, r := range defaultRegister.mapRegistries {
		r.Stop()
	}
}

func runWatcher(w Watcher) {
	ch := make(chan *Result)
	go w.Next(ch)

	for {
		result := <-ch

		if result.Err != nil {
			continue
		}

		path := internal.WorkerTrellisPath(result.Service.Name, result.Service.Version)

		nm, ok := defaultRegister.getNodeManager(path)

		if !ok {
			nm = node.New(result.NodeType, path)
		}

		nd := result.ToNode()
		switch result.Action {
		case ActionCreate, ActionUpdate:
			nm.Add(nd)
		case ActionDelete:
			nm.RemoveByID(nd.ID)
		default:
		}

		defaultRegister.setNodeManager(path, nm)
	}
}

// SetNodeManager 设置节点
func SetNodeManager(key string, nm node.Manager) {
	defaultRegister.setNodeManager(key, nm)
}

// GetNodeManager 获取节点
func GetNodeManager(key string) (node.Manager, bool) {
	return defaultRegister.getNodeManager(key)
}

func (p *Register) setNodeManager(key string, nm node.Manager) {
	p.locker.Lock()
	p.nodeManagers[key] = nm
	p.locker.Unlock()
}

func (p *Register) delNodeManager(key string) {
	p.locker.Lock()
	delete(p.nodeManagers, key)
	p.locker.Unlock()
}

func (p *Register) getNodeManager(key string) (node.Manager, bool) {
	p.locker.RLock()
	nm, ok := p.nodeManagers[key]
	p.locker.RUnlock()
	return nm, ok
}
