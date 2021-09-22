package constants

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/vigneshuvi/GoDateFormat"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4"
)

const (
	Key         = "C4018FB83C31DFCAB605D5B0292C5DA9FAB0167100E1660E0E9F105DE0CF8861"
	IdentifyKey = "id"
	Limit       = 30
	Aes         = "*$#1_k3y_c3t_v1rul_t3ch"
	Sqlite3     = "./berkas-tracker.db"
	DbType      = "postgres"
)

type JwtTandaTanganOnlineClaims struct {
	Username    string `json:"name"`
	Role        string `json:"role"`
	ID          int    `json:"id"`
	UnitKerjaID int    `json:"unit_kerja_id"`
	jwt.StandardClaims
}

type H map[string]interface{}

//HashPassword for hashpasword
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

//CheckPasswordHash for hashpasword
func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func GetFilename() string {
	return time.Now().Format(GoDateFormat.ConvertFormat("yyyymmddHHMMSS"))
}

func GetDatetimeNow() null.Time {

	return null.NewTime(time.Now(), true)

}

func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
