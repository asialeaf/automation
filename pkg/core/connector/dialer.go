package connector

import (
	"sync"
	"time"
)

type Dialer struct {
	lock        sync.Mutex
	connections map[string]Connection
}

func NewDialer() *Dialer {
	return &Dialer{
		connections: make(map[string]Connection),
	}
}

func (d *Dialer) Connect(host Host) (Connection, error) {
	var err error

	d.lock.Lock()
	defer d.lock.Unlock()

	conn, ok := d.connections[host.GetName()]
	if !ok {
		opts := Config{
			Username: host.GetUser(),
			Port:     host.GetPort(),
			Address:  host.GetAddress(),
			Password: host.GetPassword(),
			// PrivateKey: host.GetPrivateKey(),
			// KeyFile:    host.GetPrivateKeyPath(),
			Timeout: time.Duration(host.GetTimeout()) * time.Second,
		}
		conn, err = NewConnection(opts)
		if err != nil {
			return nil, err
		}
		d.connections[host.GetName()] = conn
	}

	return conn, nil
}

func (d *Dialer) Close(host Host) {
	conn, ok := d.connections[host.GetName()]
	if !ok {
		return
	}

	conn.Close()
	// logger.Log.Debugf("close connection %s", host.GetName())

	c := conn.(*connection)
	d.forgetConnection(c)
}

func (d *Dialer) forgetConnection(conn *connection) {
	d.lock.Lock()
	defer d.lock.Unlock()

	for k := range d.connections {
		if d.connections[k] == conn {
			delete(d.connections, k)
		}
	}
}
