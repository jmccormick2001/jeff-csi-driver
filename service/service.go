package service

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"k8s.io/klog"
)

const (
	// Name is the name of this CSI SP.
	Name = "jeff-csi-driver"

	// VendorVersion is the version of this CSP SP.
	VendorVersion = "0.0.0"
)

// Service is a CSI SP and idempotency.Provider.
type Service interface {
	csi.ControllerServer
	csi.IdentityServer
	csi.NodeServer
}

type service struct {
	nodeID   string
	nodeName string
}

// New returns a new Service.
func New() Service {
	return &service{}
}

func (s *service) getNodeFQDN() string {
	var err error
	defer func() {
		if res := recover(); res != nil && err == nil {
			err = errors.New("Recovered from getNodeFQDN  " + fmt.Sprint(res))
		}
	}()
	cmd := "hostname -f"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		klog.Warningf("could not get fqdn with cmd : 'hostname -f', get hostname with 'echo $HOSTNAME'")
		cmd = "echo $HOSTNAME"
		out, err = exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			klog.Errorf("Failed to execute command: %s", cmd)
			return s.nodeName
		}
	}
	nodeFQDN := string(out)
	if nodeFQDN == "" {
		klog.Warningf("node fqnd not found, setting node name as node fqdn instead")
		nodeFQDN = s.nodeName
	}
	nodeFQDN = strings.TrimSuffix(nodeFQDN, "\n")
	return nodeFQDN
}
