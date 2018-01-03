package nodeserver

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
)

//QingCloudCSINodeServer qingcloud node csi server
type QingCloudCSINodeServer struct {
}

func (server *QingCloudCSINodeServer) NodePublishVolume(context.Context, *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {

}
func (server *QingCloudCSINodeServer) NodeUnpublishVolume(context.Context, *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {

}
func (server *QingCloudCSINodeServer) GetNodeID(context.Context, *csi.GetNodeIDRequest) (*csi.GetNodeIDResponse, error) {

}
func (server *QingCloudCSINodeServer) NodeProbe(context.Context, *csi.NodeProbeRequest) (*csi.NodeProbeResponse, error) {

}
func (server *QingCloudCSINodeServer) NodeGetCapabilities(context.Context, *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {

}
