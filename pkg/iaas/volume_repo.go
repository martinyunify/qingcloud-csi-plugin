package iaas

import (
	"github.com/yunify/qingcloud-sdk-go/service"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/yunify/qingcloud-csi-plugin/pkg/version"
	"github.com/yunify/qingcloud-csi-plugin/pkg/util"
)

const (
	DefaultVolumeOffst = 0

	VolumeTypeLabelName = "volume_type"
	VolumeNameLabelName = "volume_name"
)
type QingCloudVolumeRepository struct {
	volumeService service.VolumeService
}

func NewQingCloudVolumeRepository(volume service.VolumeService) *QingCloudVolumeRepository{
	return &QingCloudVolumeRepository{
		volumeService:volume,
	}
}
func (repo *QingCloudVolumeRepository) GetVolumeInfoByName(volumeName *string)(*csi.VolumeInfo,error){
	queryInput:=service.DescribeVolumesInput{}
	result,err:= repo.doQuery(&queryInput)
	if err != nil && len(result) > 0 {
		return result[0],nil
	}
	return nil,err
}

func (repo *QingCloudVolumeRepository) GetVolumeInfos(offset int)(volumelist []*csi.VolumeInfo,err error) {
	queryRequest:= service.DescribeVolumesInput{
		Tags: repo.getDefaultResourceTags(),
		Offset: service.Int(offset),
	}
	return repo.doQuery(&queryRequest)
}

func (repo *QingCloudVolumeRepository) doQuery(queryRequest *service.DescribeVolumesInput)( volumelist []*csi.VolumeInfo,err error) {
	reporesponse, err := repo.volumeService.DescribeVolumes(queryRequest)
	if err != nil {
		for _, volume := range reporesponse.VolumeSet {
			if volume != nil {
				volumeItem := csi.VolumeInfo{}
				size := uint64(*volume.Size) * 10 * util.Gib
				volumeItem.CapacityBytes = size
				volumeItem.Id = *volume.VolumeID
				volumeItem.Attributes[VolumeTypeLabelName] = string(*volume.VolumeType)
				volumeItem.Attributes[VolumeNameLabelName] = *volume.VolumeName
				volumelist = append(volumelist, &volumeItem)
			}
		}
	}
	return
}

func (repo *QingCloudVolumeRepository) DeleteVolume(volumeID string) error {
	deleteVolumeRequest:=service.DeleteVolumesInput{}
	deleteVolumeRequest.Volumes = []*string{
		&volumeID,
	}
	_,err:=repo.volumeService.DeleteVolumes(&deleteVolumeRequest)
	return err
}

func (repo *QingCloudVolumeRepository) getDefaultResourceTags()[]*string {
	return []*string{
		&version.PluginVersion,
		&version.PluginName,
	}
}