package model

import (
	"time"

	"github.com/fishjar/gin-rest-boilerplate/db"
)

// User 用户模型
type User struct {
	Base
	Name         *string    `json:"name" gorm:"column:name;type:VARCHAR(32);not null" binding:"min=3,max=20"`                            // 姓名
	Nickname     *string    `json:"nickname" gorm:"column:nickname" binding:"omitempty"`                                                 // 昵称
	Gender       *int       `json:"gender" gorm:"column:gender;type:TINYINT;default:0" binding:"omitempty,eq=0|eq=1|eq=2"`               // 性别
	Avatar       *string    `json:"avatar" gorm:"column:avatar" binding:"omitempty"`                                                     // 昵称
	Mobile       *string    `json:"mobile" gorm:"column:mobile;type:VARCHAR(16)" binding:"omitempty"`                                    // 手机
	Email        *string    `json:"email" gorm:"column:email" binding:"omitempty,email"`                                                 // 邮箱
	Homepage     *string    `json:"homepage" gorm:"column:homepage" binding:"omitempty,url"`                                             // 个人主页
	Birthday     *time.Time `json:"birthday" gorm:"column:birthday;type:DATE" binding:"omitempty"`                                       // 生日
	Height       *float32   `json:"height" gorm:"column:height;type:FLOAT" binding:"omitempty,min=0.01,max=300"`                         // 身高(cm)
	BloodType    *string    `json:"bloodType" gorm:"column:blood_type;type:VARCHAR(8)" binding:"omitempty,eq=A|eq=B|eq=AB|eq=O|eq=NULL"` // 血型(ABO)
	Notice       *string    `json:"notice" gorm:"column:notice;type:TEXT" binding:"omitempty"`                                           // 备注
	Intro        *string    `json:"intro" gorm:"column:intro;type:TEXT" binding:"omitempty"`                                             // 简介
	Address      *string    `json:"address" gorm:"column:address;type:JSON" binding:"omitempty"`                                         // 地址
	Lives        *string    `json:"lives" gorm:"column:lives;type:JSON" binding:"omitempty"`                                             // 生活轨迹
	Tags         *string    `json:"tags" gorm:"column:tags;type:JSON" binding:"omitempty"`                                               // 标签
	LuckyNumbers *string    `json:"luckyNumbers" gorm:"column:lucky_numbers;type:JSON" binding:"omitempty"`                              // 幸运数字
	Score        *int       `json:"score" gorm:"column:score;default:0" binding:"omitempty"`                                             // 积分
	Auths        []*Auth    `json:"auths" gorm:"foreignkey:UserID"`                                                                      // 帐号
	// Roles        []*Role    `json:"roles" gorm:"many2many:userrole;"`                                                                    // 角色
	// Groups       []*Group   `json:"groups" gorm:"many2many:usergroup;"`                                                                  // 组
}

// TableName 自定义用户表名
func (User) TableName() string {
	return "user"
}

func init() {
	db.DB.AutoMigrate(&User{})

	// 插入默认用户
	// userName := "gabe"
	// userType := "admin"
	// passWord := "123456"
	// user := User{
	// 	UserName: userName,
	// 	UserType: userType,
	// 	PassWord: utils.MD5Pwd(passWord),
	// }
	// name := "Gabe Yuan"
	// user.Name = &name
	// // db.DB.Create(&user)
	// if err := db.DB.Where(&user).FirstOrCreate(&user).Error; err != nil {
	// 	fmt.Println("默认用户创建失败：", err)
	// } else {
	// 	fmt.Println("默认用户已创建")
	// 	fmt.Println("用户名：", userName)
	// 	fmt.Println("用户类型：", userType)
	// 	fmt.Println("密码：", passWord)
	// }
}
