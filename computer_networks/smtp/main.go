package main

import (
	"bytes"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/scorredoira/email"
	"html/template"
	"log"
	"net/mail"
	"net/smtp"
	"strconv"
)

type EmailConfig struct {
	Username string
	Password string
	Host     string
	Port     int
}

const INDEX_HTML = `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN">
	<html>
		<head>
		</head>
		<body>
			<table cellpadding="0" cellspacing="0" border="0" width="100%"
			style="background: whitesmoke; min-width: 320px; font-size: 1px;
			line-height: normal;">
				<tr>
					<td align="center" valign="top">
						<table cellpadding="0" cellspacing="0" border="0" width="700"
						style="background: steelblue; color: whitesmoke;
						font-family: Arial, Helvetica, sans-serif;">
							<tr>
								<td align="center" valign="top">
									<span style="font-size: 20px;
									font-weight: bold;
									line-height: 40px;
									-webkit-text-size-adjust:none; display: block;">
										Hello, {{.Name}}!
									</span>
									<hr width="600" size="1"
									color="whitesmoke" noshade>
									<span style="font-size: 16px;
									font-style: italic;
									line-height: 40px;
									-webkit-text-size-adjust:none; display: block;">
										{{.Message}}
									</span>
								</td>
							</tr>
						</table>
					</td>
				</tr>
			</table>
		</body>
	</html>`

var indexHtml = template.Must(template.New("index").Parse(INDEX_HTML))

type Content struct {
	Name    string
	Message string
}

func main() {
	db, err := sql.Open("mysql", db_login+":"+db_password+"@tcp("+db_host+")/"+db_name)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

// 	smtpHost
// 	smtpPort
// 	smtpPass
// 	smtpUser

	emailConf := &EmailConfig{smtpUser, smtpPass, smtpHost, smtpPort}

	emailauth := smtp.PlainAuth("", emailConf.Username, emailConf.Password, emailConf.Host)

	sender := mail.Address{
		Name:    "",
		Address: "",
	}

	var to string
	fmt.Print("Enter email: ")
	fmt.Scanln(&to)

	receivers := []string{
		to,
	}

	var subject string
	fmt.Print("Enter subject: ")
	fmt.Scanln(&subject)
	fmt.Print("Enter body message: ")
	var message string
	fmt.Scanln(&message)
	var name string
	fmt.Print("Enter name: ")
	fmt.Scanln(&name)

	var buf bytes.Buffer
	tmp := Content{
		Name:    name,
		Message: message,
	}

	err = indexHtml.Execute(&buf, tmp)

	emailContent := email.NewHTMLMessage(subject, buf.String())

	emailContent.From = sender
	emailContent.To = receivers

	files := []string{
		"main.go",
	}

	for _, filename := range files {
		err := emailContent.Attach(filename)

		if err != nil {
			fmt.Println(err)
		}
	}

	err = email.Send(smtpHost+":"+strconv.Itoa(emailConf.Port),
		emailauth,
		emailContent)

	var code int
	var decrypt string

	if err != nil {
		fmt.Println(err)
		code, _ = strconv.Atoi(err.Error()[:3])
		decrypt = err.Error()[3:]
	} else {
		code = 250
		decrypt = "Ok"
	}
	
	_, err = db.Exec("insert into iu9networkslabs.iu9alexeev (mail, subject, content, name, smtp_code, smtp_code_decryption)"+
		"values (?, ?, ?, ?, ?, ?)", to, subject, message, name, code, decrypt)
	if err != nil {
		log.Fatal(err)
	}
}
