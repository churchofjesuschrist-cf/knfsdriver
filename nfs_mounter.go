package nfsdriver

import (
	"code.cloudfoundry.org/goshims/execshim"
)

type nfsMounter struct {
	exec execshim.Exec
}

func NewNfsMounter(exec execshim.Exec) Mounter {
	return &nfsMounter{exec}
}

func (m *nfsMounter) Mount(source string, target string, fstype string, flags uintptr, data string) ([]byte, error) {
	cmd := m.exec.Command("mount", "-t", fstype, "-o", data, source, target)
	return cmd.CombinedOutput()
}

func (m *nfsMounter) Unmount(target string, flags int) (err error) {
	cmd := m.exec.Command("umount", target)
	return cmd.Run()
}
