package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var CN_NORTH_1 = region.NewRegion("cn-north-1", "https://ecs.cn-north-1.myhuaweicloud.com")
var CN_NORTH_4 = region.NewRegion("cn-north-4", "https://ecs.cn-north-4.myhuaweicloud.com")
var CN_SOUTH_1 = region.NewRegion("cn-south-1", "https://ecs.cn-south-1.myhuaweicloud.com")
var CN_EAST_2 = region.NewRegion("cn-east-2", "https://ecs.cn-east-2.myhuaweicloud.com")
var CN_EAST_3 = region.NewRegion("cn-east-3", "https://ecs.cn-east-3.myhuaweicloud.com")
var CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2", "https://ecs.cn-southwest-2.myhuaweicloud.com")
var AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1", "https://ecs.ap-southeast-1.myhuaweicloud.com")
var AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2", "https://ecs.ap-southeast-2.myhuaweicloud.com")
var AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3", "https://ecs.ap-southeast-3.myhuaweicloud.com")
var AF_SOUTH_1 = region.NewRegion("af-south-1", "https://ecs.af-south-1.myhuaweicloud.com")
var SA_BRAZIL_1 = region.NewRegion("sa-brazil-1", "https://ecs.sa-brazil-1.myhuaweicloud.com")
var LA_NORTH_2 = region.NewRegion("la-north-2", "https://ecs.la-north-2.myhuaweicloud.com")
var CN_SOUTH_4 = region.NewRegion("cn-south-4", "https://ecs.cn-south-4.myhuaweicloud.com")
var NA_MEXICO_1 = region.NewRegion("na-mexico-1", "https://ecs.na-mexico-1.myhuaweicloud.com")
var LA_SOUTH_2 = region.NewRegion("la-south-2", "https://ecs.la-south-2.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"cn-north-1":     CN_NORTH_1,
	"cn-north-4":     CN_NORTH_4,
	"cn-south-1":     CN_SOUTH_1,
	"cn-east-2":      CN_EAST_2,
	"cn-east-3":      CN_EAST_3,
	"cn-southwest-2": CN_SOUTHWEST_2,
	"ap-southeast-1": AP_SOUTHEAST_1,
	"ap-southeast-2": AP_SOUTHEAST_2,
	"ap-southeast-3": AP_SOUTHEAST_3,
	"af-south-1":     AF_SOUTH_1,
	"sa-brazil-1":    SA_BRAZIL_1,
	"la-north-2":     LA_NORTH_2,
	"cn-south-4":     CN_SOUTH_4,
	"na-mexico-1":    NA_MEXICO_1,
	"la-south-2":     LA_SOUTH_2,
}

var provider = region.DefaultProviderChain("ECS")

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
