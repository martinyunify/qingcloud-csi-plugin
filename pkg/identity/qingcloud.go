package identity

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/yunify/qingcloud-csi-plugin/pkg/version"
)

//QingCloudCSINodeServer qingcloud node csi server
type QingCloudCSIIndentityServer struct {
}

//GetSupportedVersions get supported csi version
func (server *QingCloudCSIIndentityServer) GetSupportedVersions(context.Context, *csi.GetSupportedVersionsRequest) (*csi.GetSupportedVersionsResponse, error) {
	return &csi.GetSupportedVersionsResponse{
		SupportedVersions: []*csi.Version{
			version.SupportedVersion,
		},
	}, nil
}

//GetPluginInfo get plugin meta info
func (server *QingCloudCSIIndentityServer) GetPluginInfo(context.Context, *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	return &csi.GetPluginInfoResponse{
		Name:          version.PluginName,
		VendorVersion: version.PluginVersion,
		Manifest:      version.PluginInfo,
	}, nil
}
