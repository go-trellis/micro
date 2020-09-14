package etcd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/go-trellis/trellis/configure"
	"github.com/go-trellis/trellis/internal"
	"github.com/go-trellis/trellis/registry"

	"go.etcd.io/etcd/clientv3"
)

type watcher struct {
	conf *configure.Watcher

	client *clientv3.Client

	fullpath string

	ctx    context.Context
	cancel context.CancelFunc

	watchChan clientv3.WatchChan
}

func newWatcher(r *Register, conf *configure.Watcher) (registry.Watcher, error) {
	w := &watcher{
		client: r.client,
		conf:   conf,
	}

	w.ctx, w.cancel = context.WithCancel(context.Background())

	w.fullpath = internal.WorkerETCDPath(w.conf.GetName(), w.conf.GetVersion())

	w.watchChan = w.client.Watch(w.ctx, w.fullpath, clientv3.WithPrefix(), clientv3.WithPrevKV())

	return w, w.ctx.Err()
}

func (p *watcher) Stop() {
	p.cancel()
}

func (p *watcher) Next(ch chan *registry.Result) {

	for {
		select {
		case wresp := <-p.watchChan:
			resp := &registry.Result{
				NodeType: p.conf.LoadBalancing,
			}
			if wresp.Err() != nil {
				resp.Err = wresp.Err()
				ch <- resp
				continue
			}

			if wresp.Canceled {
				resp.Err = errors.New("watcher was canceled")
				ch <- resp
				continue
			}

			for _, ev := range wresp.Events {

				switch ev.Type {
				case clientv3.EventTypePut:
					services, err := p.decode(ev.Kv.Value)
					if err != nil {
						resp.Err = err
						ch <- resp
						continue
					}
					if ev.IsCreate() {
						resp.Action = registry.ActionCreate
					} else if ev.IsModify() {
						resp.Action = registry.ActionUpdate
					}

					resp.Service = services
				case clientv3.EventTypeDelete:
					resp.Action = registry.ActionDelete
					service, err := p.pathToService(ev.Kv.Key)
					if err != nil {
						resp.Err = err
						ch <- resp
						continue
					}
					resp.Service = service
				default:

				}

				ch <- resp
			}
		}
	}
}

func (p *watcher) pathToService(bs []byte) (*configure.RegistService, error) {
	paths := strings.Split(string(bs), "/")
	if len(paths) < 6 {
		return nil, fmt.Errorf("service path is invalid")
	}
	s := &configure.RegistService{
		Name:    paths[3],
		Version: paths[4],
		Domain:  paths[5],
	}
	return s, nil
}

func (p *watcher) decode(bs []byte) (*configure.RegistService, error) {
	s := &configure.RegistService{}
	if err := json.Unmarshal(bs, s); err != nil {
		return nil, err
	}
	return s, nil
}

func (p *watcher) Fullpath() string {
	return p.fullpath
}
