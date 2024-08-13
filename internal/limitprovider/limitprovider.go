/*
 * Copyright (c) Siemens 2021
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package limitprovider

import (
	"encoding/json"
	"errors"
	"log"

	systemapi "systemservice/api/siemens_iedge_dmapi_v1"
	"systemservice/internal/common/utils"
)

//LimitProvider is a struct that provides NFR Limit İnfo
type LimitProvider struct {
	fs   utils.FileSystem
	util utils.Utils
	cfg  *utils.DefaultConfig
}

func (l LimitProvider) getContent(path string) (string, error) {
	content, err := l.fs.ReadFile(path)
	if err != nil {
		log.Println("limitProvider:getContent():, Cannot Read File!", err)
		return "", errors.New("cannot read file")
	}
	return string(content), nil
}

//GetLimitContent method provides Limit information
func (l LimitProvider) GetLimitContent() (*systemapi.Limits, error) {
	var err error
	limitsFile := utils.DefaultConfigPath

	//Read Limits from relevant json file
	jsonContent, err := l.getContent(limitsFile)
	if err != nil {
		return &systemapi.Limits{}, err
	}

	log.Println("limitProvider:GetLimits(), jsonContent", jsonContent)

	err = json.Unmarshal([]byte(jsonContent), &l.cfg)
	if err != nil {
		log.Println("limitProvider:GetLimits(), Unmarshal() Fail, err:", err)
		return &systemapi.Limits{}, err
	}

	log.Println("limitProvider:GetLimits(), limits:", l.cfg.Limits.String())
	return &l.cfg.Limits, nil
}

//CreateLimitProvider to create a new instance of LimitProvider
func CreateLimitProvider(fsVal utils.FileSystem, utVal utils.Utils) *LimitProvider {
	limitProvider := LimitProvider{fs: fsVal, util: utVal}

	return &limitProvider
}
