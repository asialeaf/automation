package connector

import (
	"io"
	"os"
)

type Connection interface {
	Exec(cmd string, host Host) (stdout string, code int, err error)
	PExec(cmd string, stdin io.Reader, stdout io.Writer, stderr io.Writer, host Host) (code int, err error)
	Fetch(local, remote string, host Host) error
	Scp(local, remote string, host Host) error
	RemoteFileExist(remote string, host Host) bool
	RemoteDirExist(remote string, host Host) (bool, error)
	MkDirAll(path string, mode string, host Host) error
	Chmod(path string, mode os.FileMode) error
	Close()
}

type Connector interface {
	Connect(host Host) (Connection, error)
	Close(host Host)
}

type Host interface {
	GetName() string
	SetName(name string)
	GetAddress() string
	SetAddress(str string)
	// GetInternalAddress() string
	// SetInternalAddress(str string)
	GetPort() int
	SetPort(port int)
	GetUser() string
	SetUser(u string)
	GetPassword() string
	SetPassword(password string)
	// GetPrivateKey() string
	// SetPrivateKey(privateKey string)
	// GetPrivateKeyPath() string
	// SetPrivateKeyPath(path string)
	// GetArch() string
	// SetArch(arch string)
	GetTimeout() int64
	SetTimeout(timeout int64)
	// GetRoles() []string
	// SetRoles(roles []string)
	// IsRole(role string) bool
	// GetCache() *cache.Cache
	// SetCache(c *cache.Cache)
}
