package iaas

import (
	"github.com/yunify/qingcloud-sdk-go/service"
	"github.com/container-storage-interface/spec/lib/go/csi"
)

type QingCloudVolumeFactory struct {
	volumeService service.VolumeService
}

func(factory *QingCloudVolumeFactory) CreateVolume(name string, size int) (csi.VolumeInfo,error) {
	
}