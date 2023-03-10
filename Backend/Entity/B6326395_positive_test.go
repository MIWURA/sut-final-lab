package entity_test

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string `valid:"required~Name not null"` // ต้องไม่เป็นค่าว่าง
	Email      string //`valid:"email~Email not True in formatches"`
	EmployeeID string `valid:"matches(^[M|J|S]\\d{8}$)~EM_id not True in formatches"` // รหัสพนักงานขึ7นต้นด้วย J หรือ M หรือ S แล้วตามด้วยตัวเลขจํานวน 8 ตัว
}

func TestEmpositive(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("case 1", func(t *testing.T) {
		em := Employee{
			Name:       "Mew",
			Email:      "Mew@gmail.com",
			EmployeeID: "J12345678",
		}

		ok, err := govalidator.ValidateStruct(em)
		g.Expect(ok).To(gomega.BeTrue())
		g.Expect(err).To(gomega.BeNil())
		//g.Expect(err.Error()).To(gomega.Equal("Employ..."))

	})


}
