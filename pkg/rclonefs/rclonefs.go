package rclonefs

import (
	"path"

	"github.com/fastschema/fastschema/app"
)

func NewFromConfig(diskConfigs []*app.DiskConfig, localRoot string) ([]app.Disk, error) {
	var disks []app.Disk

	for _, diskConfig := range diskConfigs {
		switch diskConfig.Driver {
		case "s3":
			s3Disk, err := NewS3(&RcloneS3Config{
				Name:            diskConfig.Name,
				Root:            diskConfig.Root,
				Provider:        diskConfig.Provider,
				Bucket:          diskConfig.Bucket,
				Region:          diskConfig.Region,
				Endpoint:        diskConfig.Endpoint,
				AccessKeyID:     diskConfig.AccessKeyID,
				SecretAccessKey: diskConfig.SecretAccessKey,
				BaseURL:         diskConfig.BaseURL,
				ACL:             diskConfig.ACL,
			})

			if err != nil {
				return nil, err
			}

			disks = append(disks, s3Disk)
		case "local":
			localDisk, err := NewLocal(&RcloneLocalConfig{
				Name:       diskConfig.Name,
				Root:       path.Join(localRoot, diskConfig.Root),
				BaseURL:    diskConfig.BaseURL,
				GetBaseURL: diskConfig.GetBaseURL,
			})

			if err != nil {
				return nil, err
			}

			disks = append(disks, localDisk)
		}
	}

	return disks, nil
}
