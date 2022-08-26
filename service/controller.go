package service

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/container-storage-interface/spec/lib/go/csi"
)

func (s *service) CreateVolume(
	ctx context.Context,
	req *csi.CreateVolumeRequest) (
	*csi.CreateVolumeResponse, error) {

	return nil, nil
}

func (s *service) DeleteVolume(
	ctx context.Context,
	req *csi.DeleteVolumeRequest) (
	*csi.DeleteVolumeResponse, error) {

	return nil, nil
}

func (s *service) ControllerPublishVolume(
	ctx context.Context,
	req *csi.ControllerPublishVolumeRequest) (
	*csi.ControllerPublishVolumeResponse, error) {

	return nil, nil
}

func (s *service) ControllerUnpublishVolume(
	ctx context.Context,
	req *csi.ControllerUnpublishVolumeRequest) (
	*csi.ControllerUnpublishVolumeResponse, error) {

	return nil, nil
}

func (s *service) ValidateVolumeCapabilities(
	ctx context.Context,
	req *csi.ValidateVolumeCapabilitiesRequest) (
	*csi.ValidateVolumeCapabilitiesResponse, error) {

	return nil, nil
}

func (s *service) ListVolumes(
	ctx context.Context,
	req *csi.ListVolumesRequest) (
	*csi.ListVolumesResponse, error) {

	return nil, nil
}

func (s *service) GetCapacity(
	ctx context.Context,
	req *csi.GetCapacityRequest) (
	*csi.GetCapacityResponse, error) {

	return nil, nil
}

func (s *service) ControllerGetVolume(
	_ context.Context, _ *csi.ControllerGetVolumeRequest,
) (*csi.ControllerGetVolumeResponse, error) {

	// Infinidat does not support ControllerGetVolume
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *service) ControllerGetCapabilities(
	ctx context.Context,
	req *csi.ControllerGetCapabilitiesRequest) (
	*csi.ControllerGetCapabilitiesResponse, error) {
	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: []*csi.ControllerServiceCapability{
			{
				Type: &csi.ControllerServiceCapability_Rpc{
					Rpc: &csi.ControllerServiceCapability_RPC{
						Type: csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
					},
				},
			},
			// &csi.ControllerServiceCapability{
			//  Type: &csi.ControllerServiceCapability_Rpc{
			//      Rpc: &csi.ControllerServiceCapability_RPC{
			//          Type: csi.ControllerServiceCapability_RPC_LIST_VOLUMES,
			//      },
			//  },
			// },

			// {
			//  Type: &csi.ControllerServiceCapability_Rpc{
			//      Rpc: &csi.ControllerServiceCapability_RPC{
			//          Type: csi.ControllerServiceCapability_RPC_GET_CAPACITY,
			//      },
			//  },
			// },

			{
				Type: &csi.ControllerServiceCapability_Rpc{
					Rpc: &csi.ControllerServiceCapability_RPC{
						Type: csi.ControllerServiceCapability_RPC_CREATE_DELETE_SNAPSHOT,
					},
				},
			},
			{
				Type: &csi.ControllerServiceCapability_Rpc{
					Rpc: &csi.ControllerServiceCapability_RPC{
						Type: csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
					},
				},
			},
			{
				Type: &csi.ControllerServiceCapability_Rpc{
					Rpc: &csi.ControllerServiceCapability_RPC{
						Type: csi.ControllerServiceCapability_RPC_CLONE_VOLUME,
					},
				},
			},
			// &csi.ControllerServiceCapability{
			//  Type: &csi.ControllerServiceCapability_Rpc{
			//      Rpc: &csi.ControllerServiceCapability_RPC{
			//          Type: csi.ControllerServiceCapability_RPC_LIST_SNAPSHOTS,
			//      },
			//  },
			// },
			{
				Type: &csi.ControllerServiceCapability_Rpc{
					Rpc: &csi.ControllerServiceCapability_RPC{
						Type: csi.ControllerServiceCapability_RPC_EXPAND_VOLUME,
					},
				},
			},
		},
	}, nil

}

func (s *service) CreateSnapshot(
	ctx context.Context,
	req *csi.CreateSnapshotRequest) (
	*csi.CreateSnapshotResponse, error) {

	return nil, nil
}

func (s *service) DeleteSnapshot(
	ctx context.Context,
	req *csi.DeleteSnapshotRequest) (
	*csi.DeleteSnapshotResponse, error) {

	return nil, nil
}

func (s *service) ListSnapshots(
	ctx context.Context,
	req *csi.ListSnapshotsRequest) (
	*csi.ListSnapshotsResponse, error) {

	return nil, nil
}

func (s *service) ControllerExpandVolume(
	ctx context.Context,
	req *csi.ControllerExpandVolumeRequest) (
	*csi.ControllerExpandVolumeResponse, error) {

	return nil, nil
}
