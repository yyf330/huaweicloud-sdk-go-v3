package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var AF_SOUTH_1 = region.NewRegion("af-south-1", "https://dcs.af-south-1.myhuaweicloud.com")
var CN_NORTH_2 = region.NewRegion("cn-north-2", "https://dcs.cn-north-2.myhuaweicloud.com")
var CN_NORTH_4 = region.NewRegion("cn-north-4", "https://dcs.cn-north-4.myhuaweicloud.com")
var CN_NORTH_1 = region.NewRegion("cn-north-1", "https://dcs.cn-north-1.myhuaweicloud.com")
var CN_EAST_2 = region.NewRegion("cn-east-2", "https://dcs.cn-east-2.myhuaweicloud.com")
var CN_EAST_3 = region.NewRegion("cn-east-3", "https://dcs.cn-east-3.myhuaweicloud.com")
var CN_SOUTH_1 = region.NewRegion("cn-south-1", "https://dcs.cn-south-1.myhuaweicloud.com")
var CN_SOUTH_2 = region.NewRegion("cn-south-2", "https://dcs.cn-south-2.myhuaweicloud.com")
var CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2", "https://dcs.cn-southwest-2.myhuaweicloud.com")
var AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2", "https://dcs.ap-southeast-2.myhuaweicloud.com")
var AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1", "https://dcs.ap-southeast-1.myhuaweicloud.com")
var AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3", "https://dcs.ap-southeast-3.myhuaweicloud.com")
var RU_NORTHWEST_2 = region.NewRegion("ru-northwest-2", "https://dcs.ru-northwest-2.myhuaweicloud.com")
var SA_BRAZIL_1 = region.NewRegion("sa-brazil-1", "https://dcs.sa-brazil-1.myhuaweicloud.com")
var LA_NORTH_2 = region.NewRegion("la-north-2", "https://dcs.la-north-2.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"af-south-1":     AF_SOUTH_1,
	"cn-north-2":     CN_NORTH_2,
	"cn-north-4":     CN_NORTH_4,
	"cn-north-1":     CN_NORTH_1,
	"cn-east-2":      CN_EAST_2,
	"cn-east-3":      CN_EAST_3,
	"cn-south-1":     CN_SOUTH_1,
	"cn-south-2":     CN_SOUTH_2,
	"cn-southwest-2": CN_SOUTHWEST_2,
	"ap-southeast-2": AP_SOUTHEAST_2,
	"ap-southeast-1": AP_SOUTHEAST_1,
	"ap-southeast-3": AP_SOUTHEAST_3,
	"ru-northwest-2": RU_NORTHWEST_2,
	"sa-brazil-1":    SA_BRAZIL_1,
	"la-north-2":     LA_NORTH_2,
}

var provider = region.DefaultProviderChain("DCS")

func ValueOf(regionId string) *region.Region {
	if regionId == "" {
		panic("unexpected empty parameter: regionId")
	}

	reg := provider.GetRegion(regionId)
	if reg != nil {
		return reg
	}

	if _, ok := staticFields[regionId]; ok {
		return staticFields[regionId]
	}
	panic(fmt.Sprintf("unexpected regionId: %s", regionId))
}
