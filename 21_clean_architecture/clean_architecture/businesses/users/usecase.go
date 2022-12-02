package users

import(
  "Go_Aldimas/21_clean_architecture/clean_architecture/app/middlewares"
)

type userUsecase struct {
  userRepository Repository
  jwtAuth *middlewares.ConfigJWT
}

func NewUserUsecase(ur Repository, jwtAuth *middlewares.ConfigJWT) Usecase {
	return &userUsecase {
		userRepository: ur,
		jwtAuth: jwtAuth,
	}
}

func (uu *userUsecase) GetAll() []Domain {
	return uu.userRepository.GetAll()
}

func (uu *userUsecase) Register(userDomain *Domain) Domain {
	return uu.userRepository.Register(userDomain)
}

func (uu *userUsecase) Login(userDomain *Domain) string {
	var user Domain = uu.userRepository.GetByEmail(userDomain)
	if user.ID == 0 {
		return ""
	}
	token := uu.jwtAuth.GenerateToken(user.ID)
	return token
}