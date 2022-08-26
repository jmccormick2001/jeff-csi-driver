package service

import (
	"golang.org/x/net/context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/protobuf/ptypes/wrappers"
)

func (s *service) GetPluginInfo(
	ctx context.Context,
	req *csi.GetPluginInfoRequest) (
	*csi.GetPluginInfoResponse, error) {

	foo := &csi.GetPluginInfoResponse{
		Name:          "jeff-csi-driver",
		VendorVersion: "2.1.2",
	}

	return foo, nil
}

func (s *service) GetPluginCapabilities(
	ctx context.Context,
	req *csi.GetPluginCapabilitiesRequest) (
	*csi.GetPluginCapabilitiesResponse, error) {
	return &csi.GetPluginCapabilitiesResponse{
		Capabilities: []*csi.PluginCapability{
			{
				Type: &csi.PluginCapability_Service_{
					Service: &csi.PluginCapability_Service{
						Type: csi.PluginCapability_Service_CONTROLLER_SERVICE,
					},
				},
			},
			// To be reported when topology contstraints are supported
			// {
			//  Type: &csi.PluginCapability_Service_{
			//      Service: &csi.PluginCapability_Service{
			//          Type: csi.PluginCapability_Service_VOLUME_ACCESSIBILITY_CONSTRAINTS,
			//      },
			//  },
			// },
		},
	}, nil

}

func (s *service) Probe(
	ctx context.Context,
	req *csi.ProbeRequest) (
	*csi.ProbeResponse, error) {

	ready := new(wrappers.BoolValue)
	ready.Value = true
	resp := new(csi.ProbeResponse)
	resp.Ready = ready
	return resp, nil
}
