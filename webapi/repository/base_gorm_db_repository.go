/*
 * @FilePath: \webapi\repository\base_gorm_db_repository.go
 * @Descripttion:
 *
 * @Date: 2020-07-23 17:37:54
 * @Author: yuanhao
 *
 */

package repository

import "github.com/jinzhu/gorm"

// 基础repository
type BaseGormDBRepository struct {
	DB *gorm.DB
}
