package csi

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/yunify/qingcloud-csi-plugin/pkg/controller"
	"github.com/yunify/qingcloud-csi-plugin/pkg/identity"
	"github.com/yunify/qingcloud-csi-plugin/pkg/nodeserver"

	"google.golang.org/grpc"
)

type QingCloudPluginServer struct {
	server *grpc.Server
}

func (server *QingCloudPluginServer) start() {
	server.server = grpc.NewServer()
	csi.RegisterIdentityServer(server.server, &identity.QingCloudCSIIndentityServer{})
	csi.RegisterNodeServer(server.server, &nodeserver.QingCloudCSINodeServer{})
	csi.RegisterControllerServer(server.server, &controller.QingCloudCSIControllerServer{})
}
