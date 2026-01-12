package value

import (
	"fmt"
	"regexp"
)

type Email string

// NewEmail はemailのバリデーションを行い、Email型を返します
func NewEmail(email string) (Email, error) {
	if !isValidEmail(email) {
		return "", fmt.Errorf("invalid email: %s", email)
	}
	return Email(email), nil
}

func isValidEmail(email string) bool {
	// 一般的なメールアドレスの形式をチェックする正規表現
	// 1. ローカル部: 英数字、ドット、各種記号
	// 2. @ 記号
	// 3. ドメイン名: 英数字、ハイフン、ドット
	// 4. トップレベルドメイン: 2文字以上の英字
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	reg := regexp.MustCompile(emailRegexPattern)
	return reg.MatchString(email)
}

func (e Email) String() string {
	return string(e)
}