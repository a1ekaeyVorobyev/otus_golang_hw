package main

import (
	"fmt"
	"reflect"

	fs "github.com/bxcodec/faker/v3"
	"syreclabs.com/go/faker"
)

type T struct {
	_ int
	_ bool
}

type SomeStructWithTags struct {
	Latitude         float32 `faker:"lat"`
	Longitude        float32 `faker:"long"`
	CreditCardNumber string  `faker:"cc_number"`
	CreditCardType   string  `faker:"cc_type"`
	Email            string  `faker:"email"`
	DomainName       string  `faker:"domain_name"`
	IPV4             string  `faker:"ipv4"`
	IPV6             string  `faker:"ipv6"`
	Password         string  `faker:"password"`
	//	Jwt                string  `faker:"jwt"`
	PhoneNumber        string  `faker:"phone_number"`
	MacAddress         string  `faker:"mac_address"`
	URL                string  `faker:"url"`
	UserName           string  `faker:"username"`
	TollFreeNumber     string  `faker:"toll_free_number"`
	E164PhoneNumber    string  `faker:"e_164_phone_number"`
	TitleMale          string  `faker:"title_male"`
	TitleFemale        string  `faker:"title_female"`
	FirstName          string  `faker:"first_name"`
	FirstNameMale      string  `faker:"first_name_male"`
	FirstNameFemale    string  `faker:"first_name_female"`
	LastName           string  `faker:"last_name"`
	Name               string  `faker:"name"`
	UnixTime           int64   `faker:"unix_time"`
	Date               string  `faker:"date"`
	Time               string  `faker:"time"`
	MonthName          string  `faker:"month_name"`
	Year               string  `faker:"year"`
	DayOfWeek          string  `faker:"day_of_week"`
	DayOfMonth         string  `faker:"day_of_month"`
	Timestamp          string  `faker:"timestamp"`
	Century            string  `faker:"century"`
	TimeZone           string  `faker:"timezone"`
	TimePeriod         string  `faker:"time_period"`
	Word               string  `faker:"word"`
	Sentence           string  `faker:"sentence"`
	Paragraph          string  `faker:"paragraph"`
	Currency           string  `faker:"currency"`
	Amount             float64 `faker:"amount"`
	AmountWithCurrency string  `faker:"amount_with_currency"`
	UUIDHypenated      string  `faker:"uuid_hyphenated"`
	UUID               string  `faker:"uuid_digit"`
	Skip               string  `faker:"-"`
	PaymentMethod      string  `faker:"oneof: cc, paypal, check, money order"` // oneof will randomly pick one of the comma-separated values supplied in the tag
	AccountID          int     `faker:"oneof: 15, 27, 61"`                     // use commas to separate the values for now. Future support for other separator characters may be added
	Price32            float32 `faker:"oneof: 4.95, 9.99, 31997.97"`
	Price64            float64 `faker:"oneof: 47463.9463525, 993747.95662529, 11131997.978767990"`
	NumS64             int64   `faker:"oneof: 1, 2"`
	NumS32             int32   `faker:"oneof: -3, 4"`
	NumS16             int16   `faker:"oneof: -5, 6"`
	NumS8              int8    `faker:"oneof: 7, -8"`
	NumU64             uint64  `faker:"oneof: 9, 10"`
	NumU32             uint32  `faker:"oneof: 11, 12"`
	NumU16             uint16  `faker:"oneof: 13, 14"`
	NumU8              uint8   `faker:"oneof: 15, 16"`
	NumU               uint    `faker:"oneof: 17, 18"`
}

func main() {
	k := faker.PhoneNumber().PhoneNumber()
	i := faker.RandomInt(0, 1000)
	d := faker.Number().Number(100)
	fmt.Println("l=", reflect.TypeOf(d))
	l := faker.Number().NumberInt32(10)
	fmt.Println("l=", l, " ", reflect.TypeOf(l))
	fmt.Println("i=", i)
	//fmt.Println("l=",l)
	fmt.Println("Hello, playground i=", k)
	/*
		digits := []int{1, 2, 4, 9, 10}
		for _, d := range digits {
			rx := fmt.Sprintf(`\d{%d}`, d)
			for i := 0; i < 10; i++ {
				res := strconv.FormatInt(int64(faker.Number().NumberInt32(d)), 10)
				fmt.Println("res=",res," rx =",rx)
			}
		}
	*/
	var t1 = T{123, true}
	var t2 = T{123, true}
	fmt.Printf("%#v \n", t1)
	fmt.Printf("%#v", t2)

	a1 := SomeStructWithTags{}
	err := fs.FakeData(&a1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a1)
	fmt.Printf("\n %#v", a1.Century)

	type SomeStruct struct {
		StringENG string `faker:"lang=eng"`
		StringCHI string `faker:"lang=chi"`
		StringRUS string `faker:"lang=rus"`
	}
	fmt.Println(a1.DayOfWeek)
	fmt.Println("---")
	aa := SomeStruct{}
	_ = fs.SetRandomStringLength(5)
	_ = fs.FakeData(&aa)
	fmt.Printf("%+v", aa)
}
