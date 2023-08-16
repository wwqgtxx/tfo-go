package tfo

import (
	"syscall"

	"golang.org/x/sys/unix"
)

func SetTFOListener(fd uintptr) error {
	return setTFO(fd)
}

func SetTFODialer(fd uintptr) error {
	return setTFO(fd)
}

func setTFO(fd uintptr) error {
	return unix.SetsockoptInt(int(fd), unix.IPPROTO_TCP, unix.TCP_FASTOPEN, 1)
}

func (*Dialer) socket(domain int) (int, error) {
	return unix.Socket(domain, unix.SOCK_STREAM|unix.SOCK_NONBLOCK|unix.SOCK_CLOEXEC, unix.IPPROTO_TCP)
}

func (*Dialer) setIPv6Only(fd int, family int, ipv6only bool) error {
	return setIPv6Only(fd, family, ipv6only)
}

const connectSyscallName = "sendmsg"

func doConnect(fd uintptr, rsa syscall.Sockaddr, b []byte) (int, error) {
	return syscall.SendmsgN(int(fd), b, nil, rsa, 0)
}
