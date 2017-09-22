package nodeStorage

import "github.com/dzen-it/di/system/db"

var Storage *NodeStorage

func init() {
	Storage = new(NodeStorage)
}

const (
	// TIMEALIVE is time in seconds
	TIMEALIVE                = 60
	CLIENT_CACHE TypeStorage = 0
	CLIENT_REDIS TypeStorage = 1
)

type TypeStorage int8

type NodeStorage struct {
}

func (s *NodeStorage) Set(name string, addr string) error {
	keys, err := db.Client.Keys(name)
	if err != nil {
		return err
	}
	if len(keys) > 0 {
		if err := db.Client.Del(keys...); err != nil {
			return err
		}
	}

	return db.Client.Set(createKeyForDB(name, addr), addr)
}

func (s *NodeStorage) Add(name string, addr string) error {
	return db.Client.Set(createKeyForDB(name, addr), addr)
}

func (s *NodeStorage) Get(name string) (result []string, err error) {
	return db.Client.Keys(name)
}

func createKeyForDB(name, addr string) string {
	return name + "|" + addr
}
