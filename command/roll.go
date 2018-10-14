package command

func Roll() string {
	return intToString(random(1,100))
}
