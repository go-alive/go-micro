package service

import (
	"github.com/go-alive/go-micro/config/source"
	proto "github.com/go-alive/go-micro/config/source/service/proto"
)

type watcher struct {
	stream proto.Config_WatchService
}

func newWatcher(stream proto.Config_WatchService) (source.Watcher, error) {
	return &watcher{stream: stream}, nil
}

func (w *watcher) Next() (*source.ChangeSet, error) {
	var rsp proto.WatchResponse
	err := w.stream.RecvMsg(&rsp)
	if err != nil {
		return nil, err
	}
	return toChangeSet(rsp.ChangeSet), nil
}

func (w *watcher) Stop() error {
	return w.stream.Close()
}
