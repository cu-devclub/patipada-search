package repositories

type userJordanWrightEmailing struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewUserJordanWrightEmailing(name string, fromEmailAddress string, fromEmailPassword string) UserEmailing {
	return &userJordanWrightEmailing{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}
