package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 备份实例请求体
type BackupInstanceBody struct {

	// 备份缓存实例的备注信息。
	Remark *string `json:"remark,omitempty"`

	// 备份缓存实例的格式。
	BackupFormat *BackupInstanceBodyBackupFormat `json:"backup_format,omitempty"`
}

func (o BackupInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupInstanceBody struct{}"
	}

	return strings.Join([]string{"BackupInstanceBody", string(data)}, " ")
}

type BackupInstanceBodyBackupFormat struct {
	value string
}

type BackupInstanceBodyBackupFormatEnum struct {
	AOF BackupInstanceBodyBackupFormat
	RDB BackupInstanceBodyBackupFormat
}

func GetBackupInstanceBodyBackupFormatEnum() BackupInstanceBodyBackupFormatEnum {
	return BackupInstanceBodyBackupFormatEnum{
		AOF: BackupInstanceBodyBackupFormat{
			value: "aof",
		},
		RDB: BackupInstanceBodyBackupFormat{
			value: "rdb",
		},
	}
}

func (c BackupInstanceBodyBackupFormat) Value() string {
	return c.value
}

func (c BackupInstanceBodyBackupFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupInstanceBodyBackupFormat) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
