package smtp

import (
	"fmt"
	"net/smtp"
)

type SMTPService struct {
	host string
	port string
	from string
}

func New(
	host string,
	port string,
	from string,
) *SMTPService {
	return &SMTPService{
		host: host,
		port: port,
		from: from,
	}
}

func (s *SMTPService) SendActivationEmail(
	to string,
	token string,
) error {

	link := fmt.Sprintf(
		"http://localhost:8080/auth/activate?token=%s",
		token,
	)

	message := []byte(
		"Subject: Активация аккаунта\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
			"\r\n" +

			"Здравствуйте!\n\n" +

			"Для завершения регистрации и активации аккаунта " +
			"необходимо подтвердить адрес электронной почты.\n\n" +

			"Перейдите по ссылке ниже:\n\n" +
			link + "\n\n" +

			"Ссылка действительна в течение 5 минут.\n\n" +

			"Если вы не выполняли вход в систему, " +
			"проигнорируйте это письмо или обратитесь " +
			"в техническую поддержку.\n\n" +

			"С уважением,\n" +
			"Inno Accounting",
	)

	return smtp.SendMail(
		s.host+":"+s.port,
		nil,
		s.from,
		[]string{to},
		message,
	)
}