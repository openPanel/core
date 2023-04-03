package config

import (
	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/global"
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

func LoadNodesCache() ([]NodeCacheEntry, error) {
	s := new([]NodeCacheEntry)
	err := Load(constant.ConfigKeyNodesCache, s, constant.LocalStore)
	if err != nil {
		return nil, err
	}
	return *s, nil
}

func AppendNodesCache(newEntry NodeCacheEntry) error {
	s := new([]NodeCacheEntry)
	err := Load(constant.ConfigKeyNodesCache, s, constant.LocalStore)
	if err != nil {
		return err
	}

	news := make([]NodeCacheEntry, 0, len(*s)+1)
	for _, e := range *s {
		if e.Id != newEntry.Id {
			news = append(news, e)
		} else {
			news = append(news, newEntry)
		}
	}
	return Save(constant.ConfigKeyNodesCache, news, constant.LocalStore)
}

func UpdateNodesCache(newEntries []NodeCacheEntry) error {
	return Save(constant.ConfigKeyNodesCache, newEntries, constant.LocalStore)
}
