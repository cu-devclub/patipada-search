package usecases

import "log"

func (s *UsersUsecaseImpl) VerifyUsername(username string) (bool, error) {
	log.Println("Verifying username: ", username, " ....")
	_, err := s.usersRepository.GetUserByUsername(username)
	if err != nil {
		return false, err
	}

	log.Println("Username: ", username, " is verified")
	return true, nil
}
