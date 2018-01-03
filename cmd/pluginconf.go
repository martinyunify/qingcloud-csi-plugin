package cmd

import (
	"github.com/yunify/qingcloud-csi-plugin/pkg/csi"
)

//QingCloudCSIPluginConf configuration of qingcloud iaas storage resource plugin
type QingCloudCSIPluginConf struct {
	server   string
	uniqueID string
}

func (config *QingCloudCSIPluginConf) NewPluginInstance() *csi.QingCloudPluginServer {

}
