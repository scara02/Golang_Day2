package domain

type Contact struct {
    ID           int
    FullName     string
    PhoneNumber  string
}

func NewContact(id int, fullName string, phoneNumber string) Contact {
    return Contact{
        ID:           id,
        FullName:     fullName,
        PhoneNumber:  phoneNumber,
    }
}

func (c *Contact) SetFullName(lastName, firstName, middleName string) {
    c.FullName = lastName + " " + firstName + " " + middleName
}

func (c *Contact) GetFullName() string {
    return c.FullName
}

func (c *Contact) SetPhoneNumber(phoneNumber string) {
    var cleanedNumber string
    for _, char := range phoneNumber {
        if char >= '0' && char <= '9' {
            cleanedNumber += string(char)
        }
    }
    c.PhoneNumber = cleanedNumber
}