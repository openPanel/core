package config

import (
	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/global"
)

func SaveLocalNodeInfo(i global.LocalNodeInfo) error {
	return Save(constant.ConfigKeyNodeInfo, i, constant.LocalStore)
}

func LoadLocalNodeInfo() (global.LocalNodeInfo, error) {
	i := new(global.LocalNodeInfo)
	err := Load(constant.ConfigKeyNodeInfo, i, constant.LocalStore)
	if err != nil {
		return global.LocalNodeInfo{}, err
	}
	return *i, nil
}

func SaveClusterInfo(i global.ClusterInfo) error {
	return Save(constant.ConfigKeyClusterInfo, i, constant.SharedStore)
}

func LoadClusterInfo() (global.ClusterInfo, error) {
	i := new(global.ClusterInfo)
	err := Load(constant.ConfigKeyClusterInfo, i, constant.SharedStore)
	if err != nil {
		return global.ClusterInfo{}, err
	}
	return *i, nil
}
