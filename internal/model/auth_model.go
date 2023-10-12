package model

import (
	"github.com/golang-jwt/jwt"
)

// // AuthService represents the AuthService service.
// type AuthService interface {
//	Register(*user.CreateUserRequest) (*RegisterResponse, error)
//	Login(*LoginRequest) (*LoginResponse, error)
//	Logout(*LogoutRequest) (*LogoutResponse, error)
//	GetProfile(*GetUserProfileRequest) (*GetUserProfileResponse, error)
//	GetProtectedData(*GetProtectedDataRequest) (*SecuredResponse, error)
// }

type Claims struct {
	ID	string `json:"ID"`
	Sub string `json:"UUID"`
	jwt.StandardClaims
}

// // User represents the User message.
// type User struct {
//	ID		 string		   `json:"ID"`
//	UUID	 string		   `json:"UUID"`
//	Metadata *UserMetadata `json:"metadata"`
//	Content	 *UserContent  `json:"content"`
// }

// // UserMetadata represents the UserMetadata message.
// type UserMetadata struct {
//	Name  string			 `json:"name"`
//	Dates *CommonDate `json:"dates"`
// }

// // UserContent represents the UserContent message.
// type UserContent struct {
//	Email		 string	  `json:"email"`
//	Phone		 string	  `json:"phone"`
//	LastName	 string	  `json:"lastName"`
//	FirstName	 string	  `json:"firstName"`
//	ProjectRoles []string `json:"projectRoles"`
//	ScrumRoles	 []string `json:"scrumRoles"`
//	Password	 string	  `json:"password"`
// }

// RegisterResponse represents the RegisterResponse message.
type RegisterResponse struct {
	ID	  string `json:"ID"`
	UUID  string `json:"UUID"`
	Email string `json:"email"`
}

// LoginRequest represents the LoginRequest message.
type LoginRequest struct {
	ID		 string `json:"ID" binding:"required"`
	Email	 string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents the LoginResponse message.
type LoginResponse struct {
	ID	  string `json:"ID"`
	UUID  string `json:"UUID"`
	Email string `json:"email"`
	Token string `json:"token" binding:"required"`
}

// LogoutRequest represents the LogoutRequest message.
type LogoutRequest struct {
	ID	  string `json:"ID"`
	UUID  string `json:"UUID"`
	Email string `json:"email"`
	Token string `json:"token"`
}

// LogoutResponse represents the LogoutResponse message.
type LogoutResponse struct {
	ID	  string `json:"ID"`
	UUID  string `json:"UUID"`
	Email string `json:"email"`
}

// SecuredResponse represents the SecuredResponse message.
type SecuredResponse struct {
	ID			string `json:"ID"`
	UUID		string `json:"UUID"`
	Email		string `json:"email"`
	Token		string `json:"token"`
	SecuredData string `json:"securedData"`
}

// GetUserProfileRequest represents the GetUserProfileRequest message.
type GetUserProfileRequest struct {
	ID	  string `json:"ID"`
	UUID  string `json:"UUID"`
	Email string `json:"email"`
	Token string `json:"token"`
}

// GetUserProfileResponse represents the GetUserProfileResponse message.
type GetUserProfileResponse struct {
	ID			string `json:"ID"`
	UUID		string `json:"UUID"`
	Email		string `json:"email"`
	Token		string `json:"token"`
	SecuredData string `json:"securedData"`
}

// GetProtectedDataRequest represents the GetProtectedDataRequest message.
type GetProtectedDataRequest struct {
	ID	  string `json:"ID"`
	UUID  string `json:"UUID"`
	Email string `json:"email"`
	Token string `json:"token"`
}

// GetProtectedDataResponse represents the GetProtectedDataResponse message.
type GetProtectedDataResponse struct {
	ID			string `json:"ID"`
	UUID		string `json:"UUID"`
	Email		string `json:"email"`
	Token		string `json:"token"`
	SecuredData string `json:"securedData"`
}

// package model

// import (
//	"time"

//	"github.com/golang-jwt/jwt"
// )

// type Claims struct {
//	Username string `json:"username"`
//	Sub		 string `json:"UUID"`
//	jwt.StandardClaims
// }

// // LoginRequest represents the user login credentials.
// type LoginRequest struct {
//	Username string `json:"username" binding:"required"`
//	Password string `json:"password" binding:"required"`
// }

// // TokenResponse represents the token response after successful login.
// // type TokenResponse struct {
// type LoginResponse struct {
//	Token string `json:"token" binding:"required"`
// }

// type SecuredRequest struct {
//	Token string `json:"token"`
// }

// type SecuredResponse struct {
//	Message string `json:"message"`
// }

// type GetUserProfileResponse struct {
//	UUID		string	  `json:"UUID"`
//	LastName	string	  `json:"lastName"`
//	FirstName	string	  `json:"firstName"`
//	Role		string	  `json:"role"`
//	DateOfBirth time.Time `json:"dateOfBirth"`
//	Email		string	  `json:"email"`
//	PhoneNumber string	  `json:"phoneNumber"`
//	Username	string	  `json:"username"`
// }
