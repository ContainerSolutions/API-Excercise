package postgres_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gitlab.com/DGuedes/API-Exercise/titanic"
	"gitlab.com/DGuedes/API-Exercise/util"
)

var _ = Describe("Person DAO", func() {
	BeforeEach(func() {
		p.AutoMigrate()
	})

	AfterEach(func() {
		p.DB.Exec("DROP TABLE if exists people cascade")
	})

	It("Deletes person successfully", func() {
		person := titanic.Person{Name: "abc"}
		err := p.Insert(&person)
		util.Check(err)
		Expect(person.UUID).NotTo(BeNil())

		q1, _ := p.GetAll()
		Expect(q1).To(HaveLen(1))

		p.Destroy(person.UUID)

		q2, _ := p.GetAll()
		Expect(q2).To(HaveLen(0))
	})

	It("Edit person successfully", func() {
		person := titanic.Person{Name: "abc"}
		err := p.Insert(&person)
		util.Check(err)

		person.Name = "newname"
		p.Update(&person)

		q1, _ := p.Find(person.UUID)
		Expect(q1.Name).To(Equal("newname"))
	})

	It("Finds person successfully", func() {
		person := titanic.Person{Name: "abc"}
		p.Insert(&person)

		q1, _ := p.Find(person.UUID)
		Expect(q1.Name).To(Equal("abc"))
	})

	It("Inserts a new person successfully", func() {
		people, _ := p.GetAll()
		Expect(people).To(BeEmpty())

		err := p.Insert(&titanic.Person{})
		Expect(err).To(BeNil())

		people, _ = p.GetAll()
		Expect(people).To(HaveLen(1))
	})

	It("Inserts person in bulk successfully", func() {
		err := p.BulkInsert(titanic.People{
			titanic.Person{Name: "Person1", PassengerClass: 3},
			titanic.Person{Name: "Person2", PassengerClass: 5},
		})
		util.Check(err)

		// BulkInsert isn't atomic and needs sometime to be concluded
		time.Sleep(time.Millisecond * 100)

		people, _ := p.GetAll()
		Expect(people).To(HaveLen(2))
		Expect(people[0].Name).To(Equal("Person1"))
		Expect(people[0].PassengerClass).To(Equal(3))
	})
})
