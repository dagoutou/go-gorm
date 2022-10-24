package main

import (
	common "common_x"
)

//Privilege 权限表
type Privilege struct {
	ID             int            `json:"id" gorm:"primary_key;not null" required:"false" example:"1"`               // 记录ID
	Name           string         `json:"name" gorm:"type:varchar(10);not null" validate:"required" required:"true"` // 知识库名称
	Description    string         `json:"description" gorm:"type:text;"`                                             // 权限描述
	LabelID        string         `json:"label_id"`                                                                  // 所属组ID
	PrivilegeLabel PrivilegeLabel `gorm:"foreignKey:LabelID"`                                                        // 所属权限
	Role           []Role         `gorm:"many2many:privilege_role"`                                                  // 角色
}

type PrivilegeLabel struct {
	ID          int         `json:"id" gorm:"primary_key;not null" required:"false" example:"1"`               // 记录ID
	Name        string      `json:"name" gorm:"type:varchar(10);not null" validate:"required" required:"true"` // 所属组名称
	Description string      `json:"description" gorm:"type:text;"`
	Privilege   []Privilege `grom:"foreignKey:LabelID"`
}

//PrivilegesGroup 权限信息
type PrivilegesGroup struct {
	LabelID    int             `json:"label_id"`   // 权限分组id
	LabelName  string          `json:"label_name"` // 权限分组名称
	Privileges []PrivilegeInfo `json:"privileges"` // 权限列表
}

//PrivilegeInfo 权限信息
type PrivilegeInfo struct {
	ID   int    `json:"id"`   // id
	Name string `json:"name"` // 权限名称
}

//PrivilegeRole 角色与权限的关联表
type PrivilegeRole struct {
	RoleID      int `json:"role_id"`      //角色ID
	PrivilegeID int `json:"privilege_id"` // 权限ID
}

//Role 角色
type Role struct {
	ID          int         `json:"id" gorm:"primary_key;not null" required:"false" example:"1"`                 // ID
	Name        string      `json:"name" gorm:"type:varchar(20);not null;unique"  required:"true" example:"管理员"` //角色名称
	Description string      `json:"description" gorm:"type:text;"`                                               // 权限描述
	Privilege   []Privilege `json:"privilege" gorm:"many2many:privilege_role"`                                   // 关联表
	User        []User      `json:"user" gorm:"many2many:role_relationship;joinForeignKey:RoleID"`               // 关联角色 其实joinForeignKey为关联表的字段
	common.Base
}

//RoleRelationship 角色与用户的关联表
type RoleRelationship struct {
	RoleID     int    `json:"role_id"`                                                            // 角色ID
	UnionID    string `json:"union_id"`                                                           // 用户ID
	Status     string `json:"status" gorm:"type:varchar(50)" enums:"forever,interim,uneffective"` // 有效期类型
	ExpireData int64  `json:"expire_data" gorm:"type:bigint"`                                     // 过期时间
}

//RoleUserInfo 角色用户信息
type RoleUserInfo struct {
	ID         string `json:"id"`          // 角色ID
	Name       string `json:"name"`        // 姓名
	Phone      string `json:"phone"`       // 电话
	Email      string `json:"email"`       // 邮箱
	Position   string `json:"position"`    // 职位
	Status     string `json:"status"`      // 绑定状态
	ExpireData int64  `json:"expire_data"` // 有效期
}

//User 用户
type User struct {
	ID       string `json:"id"  gorm:"<-:false" example:"1"`                                                             // 记录ID
	Username string `json:"username" gorm:"type:varchar(50);not null;unique" validate:"required,max=50" example:"admin"` // 用户名
	Password string `json:"password" gorm:"type:varchar(128);not null" validate:"required,max=128" example:"admin"`      // 密码
	Role     []Role `json:"role" gorm:"many2many:role_relationship;joinForeignKey:UnionID"`                              // 关联角色 其实joinForeignKey为关联表的字段
}

func main() {

}
