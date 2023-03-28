package config

import (
	"github.com/openPanel/core/app/global"
)

type Key string

const (
	ServerNodeInfo Key = "local.nodeInfo"
)

func SaveLocalNodeInfo(i global.LocalNodeInfo) error {
	return Save(ServerNodeInfo, i)
}

func LoadLocalNodeInfo() (global.LocalNodeInfo, error) {
	i := new(global.LocalNodeInfo)
	err := Load(ServerNodeInfo, &i)
	if err != nil {
		return global.LocalNodeInfo{}, err
	}
	return *i, nil
}
