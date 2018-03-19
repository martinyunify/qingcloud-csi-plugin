package version

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
)

//PluginName plugin name
var PluginName = "Qingcloud_csi"

//PluginVersion version number of plugin
var PluginVersion = "git"

//PluginInfo version info of plugin
var PluginInfo map[string]string

//SupportedVersion supported csi version
var SupportedVersion = &csi.Version{Major: 0, Minor: 1, Patch: 0}

