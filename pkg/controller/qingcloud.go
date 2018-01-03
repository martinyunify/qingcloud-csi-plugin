package controller

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/yunify/qingcloud-sdk-go/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	kib    uint64 = 1024
	mib    uint64 = kib * 1024
	gib    uint64 = mib * 1024
	gib100 uint64 = gib * 100
	tib    uint64 = gib * 1024
	tib100 uint64 = tib * 100
)

//QingCloudCSIControllerServer qingcloud node csi server
type QingCloudCSIControllerServer struct {
	volumeService service.VolumeService
	uniqueID      string
}

func (server *QingCloudCSIControllerServer) CreateVolume(ctx context.Context, request *csi.CreateVolumeRequest) (response *csi.CreateVolumeResponse, err error) {
	capacityRange := request.GetCapacityRange()
	if capacityRange.GetLimitBytes()%(10*gib) != 0 {
		err = status.Errorf(codes.OutOfRange, "%d can't be devided by 10GB", capacityRange.GetLimitBytes())
	}

}
func (server *QingCloudCSIControllerServer) DeleteVolume(context.Context, *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {

}
func (server *QingCloudCSIControllerServer) ControllerPublishVolume(context.Context, *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {

}
func (server *QingCloudCSIControllerServer) ControllerUnpublishVolume(context.Context, *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {

}
func (server *QingCloudCSIControllerServer) ValidateVolumeCapabilities(ctx context.Context, request *csi.ValidateVolumeCapabilitiesRequest) (response *csi.ValidateVolumeCapabilitiesResponse, err error) {
	describeVolumeRequest := service.DescribeVolumesInput{}
	volumeid := request.GetVolumeId()
	describeVolumeRequest.Volumes = []*string{&volumeid}
	capabilities := request.GetVolumeCapabilities()
	voluemresponse, err := server.volumeService.DescribeVolumes(&describeVolumeRequest)
	if err != nil {
		response.Supported = false
		return
	} else if *voluemresponse.TotalCount > 0 {
		for _, target := range capabilities {
			if target.AccessMode.GetMode() == csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER {
				response.Supported = true
			} else {
				response.Supported = false
			}
		}
	}
	return

}
func (server *QingCloudCSIControllerServer) ListVolumes(context.Context, *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {

}

func (server *QingCloudCSIControllerServer) GetCapacity(context.Context, *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	return &csi.GetCapacityResponse{
		AvailableCapacity: tib * 4,
	}, nil
}

func (server *QingCloudCSIControllerServer) ControllerProbe(context.Context, *csi.ControllerProbeRequest) (*csi.ControllerProbeResponse, error) {
	return &csi.ControllerProbeResponse{}, nil
}

func (server *QingCloudCSIControllerServer) ControllerGetCapabilities(context.Context, *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: []*csi.ControllerServiceCapability{
			&csi.ControllerServiceCapability{
				Type: &csi.ControllerServiceCapability_Rpc{
					Rpc: &csi.ControllerServiceCapability_RPC{
						Type: csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
					},
				},
			},
			&csi.ControllerServiceCapability{
				Type: &csi.ControllerServiceCapability_Rpc{
					Rpc: &csi.ControllerServiceCapability_RPC{
						Type: csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
					},
				},
			},
			&csi.ControllerServiceCapability{
				Type: &csi.ControllerServiceCapability_Rpc{
					Rpc: &csi.ControllerServiceCapability_RPC{
						Type: csi.ControllerServiceCapability_RPC_LIST_VOLUMES,
					},
				},
			},
			&csi.ControllerServiceCapability{
				Type: &csi.ControllerServiceCapability_Rpc{
					Rpc: &csi.ControllerServiceCapability_RPC{
						Type: csi.ControllerServiceCapability_RPC_GET_CAPACITY,
					},
				},
			},
		},
	}, nil
}
