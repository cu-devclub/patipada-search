package usecases

func (s *UsersUsecaseImpl) VerifyUsername(username string) (bool, error) {
	_, err := s.usersRepository.GetUserByUsername(username)
	if err != nil {
		return false,err
	}

	return true,nil 
}
