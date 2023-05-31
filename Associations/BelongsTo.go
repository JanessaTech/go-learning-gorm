package associations

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Empoyee struct {
	ID           uint           `gorm:"column:id"`
	Name         string         `gorm:"column:name;default:JanessaTech"`
	Email        string         `gorm:"column:email"`
	Age          uint8          `gorm:"column:age;default:10"` // provided default value for age
	CompanyID    int            `gorm:"column:company_id"`
	Company      Company        `gorm:"foreignKey:CompanyID"`
	Birthday     time.Time      `gorm:"column:birthday"`
	MemberNumber sql.NullString `gorm:"column:member_number"`
	ActivatedAt  sql.NullTime   `gorm:"column:activated_at"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
}

type Company struct {
	gorm.Model
	ID   int    `gorm:"column:id"`
	Name string `gorm:"column:name;default:demoCompany"`
}

func BelongToDemo() {

}
