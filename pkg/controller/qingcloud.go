package controller

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/yunify/qingcloud-sdk-go/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hash/adler32"
)



//QingCloudCSIControllerServer qingcloud node csi server
type QingCloudCSIControllerServer struct {
	volumeService service.VolumeService
	uniqueID      string
}

func (server *QingCloudCSIControllerServer) CreateVolume(ctx context.Context, request *csi.CreateVolumeRequest) (response *csi.CreateVolumeResponse, err error) {
	capacityRange := request.GetCapacityRange()
	limit:=capacityRange.GetLimitBytes()
	if limit < 10 * gib || limit > 500 * gib{
		err = status.Errorf(codes.OutOfRange, "limit %d is less than 10GB or bigger than 500GB", limit)
	}

	if limit %(10*gib) != 0 {
		err = status.Errorf(codes.OutOfRange, "%d can't be devided by 10GB", limit)
	}

	required:= capacityRange.RequiredBytes
	if required > 500 *gib {
		err = status.Errorf(codes.OutOfRange, "required storage %d is bigger than the upper bound of size limit 500GB",required)
	}

	CreateVolumeRequest:= service.CreateVolumesInput{
		Count: service.Int(1),
		VolumeType: service.Int(0),
		Size: service.Int(int(limit/gib)),
		VolumeName: service.String(request.Name),
	}
	apiresponse,err:=server.volumeService.CreateVolumes(&CreateVolumeRequest)
	if err != nil {

	}
	for _,volume := range apiresponse.Volumes {
		response.VolumeInfo = &csi.VolumeInfo{
			Id: *volume,
		}
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
