package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2", "https://ddm.cn-southwest-2.myhuaweicloud.com")
var CN_SOUTH_2 = region.NewRegion("cn-south-2", "https://ddm.cn-south-2.myhuaweicloud.com")
var CN_NORTH_4 = region.NewRegion("cn-north-4", "https://ddm.cn-north-4.myhuaweicloud.com")
var CN_EAST_3 = region.NewRegion("cn-east-3", "https://ddm.cn-east-3.myhuaweicloud.com")
var CN_EAST_2 = region.NewRegion("cn-east-2", "https://ddm.cn-east-2.myhuaweicloud.com")
var CN_NORTH_1 = region.NewRegion("cn-north-1", "https://ddm.cn-north-1.myhuaweicloud.com")
var CN_SOUTH_1 = region.NewRegion("cn-south-1", "https://ddm.cn-south-1.myhuaweicloud.com")
var AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1", "https://ddm.ap-southeast-1.myhuaweicloud.com")
var LA_NORTH_2 = region.NewRegion("la-north-2", "https://ddm.la-north-2.myhuaweicloud.com")
var SA_BRAZIL_1 = region.NewRegion("sa-brazil-1", "https://ddm.sa-brazil-1.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"cn-southwest-2": CN_SOUTHWEST_2,
	"cn-south-2":     CN_SOUTH_2,
	"cn-north-4":     CN_NORTH_4,
	"cn-east-3":      CN_EAST_3,
	"cn-east-2":      CN_EAST_2,
	"cn-north-1":     CN_NORTH_1,
	"cn-south-1":     CN_SOUTH_1,
	"ap-southeast-1": AP_SOUTHEAST_1,
	"la-north-2":     LA_NORTH_2,
	"sa-brazil-1":    SA_BRAZIL_1,
}

var provider = region.DefaultProviderChain("DDM")

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
