package config

import (
	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/global"
	"github.com/openPanel/core/app/manager/router"
)

func SaveLocalNodeInfo(i global.NodeInfo) error {
	return Save(constant.ConfigKeyNodeInfo, i, constant.LocalStore)
}

func LoadLocalNodeInfo() (global.NodeInfo, error) {
	i := new(global.NodeInfo)
	err := Load(constant.ConfigKeyNodeInfo, i, constant.LocalStore)
	if err != nil {
		return global.NodeInfo{}, err
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

func LoadNodesCache() ([]router.Node, error) {
	s := new([]router.Node)
	err := Load(constant.ConfigKeyNodesCache, s, constant.LocalStore)
	if err != nil {
		return nil, err
	}
	return *s, nil
}

func UpdateNodesCache(newEntries []router.Node) error {
	return Save(constant.ConfigKeyNodesCache, newEntries, constant.LocalStore)
}
