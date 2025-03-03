package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var CN_NORTH_4 = region.NewRegion("cn-north-4", "https://gaussdb-opengauss.cn-north-4.myhuaweicloud.com")
var CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2", "https://gaussdb-opengauss.cn-southwest-2.myhuaweicloud.com")
var CN_SOUTH_1 = region.NewRegion("cn-south-1", "https://gaussdb-opengauss.cn-south-1.myhuaweicloud.com")
var RU_NORTHWEST_2 = region.NewRegion("ru-northwest-2", "https://gaussdb-opengauss.ru-northwest-2.myhuaweicloud.com")
var AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3", "https://gaussdb-opengauss.ap-southeast-3.myhuaweicloud.com")
var CN_NORTH_2 = region.NewRegion("cn-north-2", "https://gaussdb-opengauss.cn-north-2.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"cn-north-4":     CN_NORTH_4,
	"cn-southwest-2": CN_SOUTHWEST_2,
	"cn-south-1":     CN_SOUTH_1,
	"ru-northwest-2": RU_NORTHWEST_2,
	"ap-southeast-3": AP_SOUTHEAST_3,
	"cn-north-2":     CN_NORTH_2,
}

var provider = region.DefaultProviderChain("GAUSSDBFOROPENGAUSS")

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
