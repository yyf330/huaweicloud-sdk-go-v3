package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var AF_SOUTH_1 = region.NewRegion("af-south-1", "https://swr-api.af-south-1.myhuaweicloud.com")
var AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1", "https://swr-api.ap-southeast-1.myhuaweicloud.com")
var AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2", "https://swr-api.ap-southeast-2.myhuaweicloud.com")
var AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3", "https://swr-api.ap-southeast-3.myhuaweicloud.com")
var CN_EAST_2 = region.NewRegion("cn-east-2", "https://swr-api.cn-east-2.myhuaweicloud.com")
var CN_EAST_3 = region.NewRegion("cn-east-3", "https://swr-api.cn-east-3.myhuaweicloud.com")
var CN_NORTH_1 = region.NewRegion("cn-north-1", "https://swr-api.cn-north-1.myhuaweicloud.com")
var CN_NORTH_2 = region.NewRegion("cn-north-2", "https://swr-api.cn-north-2.myhuaweicloud.com")
var CN_NORTH_4 = region.NewRegion("cn-north-4", "https://swr-api.cn-north-4.myhuaweicloud.com")
var CN_SOUTH_1 = region.NewRegion("cn-south-1", "https://swr-api.cn-south-1.myhuaweicloud.com")
var CN_SOUTH_2 = region.NewRegion("cn-south-2", "https://swr-api.cn-south-2.myhuaweicloud.com")
var CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2", "https://swr-api.cn-southwest-2.myhuaweicloud.com")
var RU_NORTHWEST_2 = region.NewRegion("ru-northwest-2", "https://swr-api.ru-northwest-2.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"af-south-1":     AF_SOUTH_1,
	"ap-southeast-1": AP_SOUTHEAST_1,
	"ap-southeast-2": AP_SOUTHEAST_2,
	"ap-southeast-3": AP_SOUTHEAST_3,
	"cn-east-2":      CN_EAST_2,
	"cn-east-3":      CN_EAST_3,
	"cn-north-1":     CN_NORTH_1,
	"cn-north-2":     CN_NORTH_2,
	"cn-north-4":     CN_NORTH_4,
	"cn-south-1":     CN_SOUTH_1,
	"cn-south-2":     CN_SOUTH_2,
	"cn-southwest-2": CN_SOUTHWEST_2,
	"ru-northwest-2": RU_NORTHWEST_2,
}

var provider = region.DefaultProviderChain("SWR")

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
