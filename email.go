
// Sending Email Using Smtp in Golang
package main
 
import (
    "fmt"
    "net/smtp"
    "os"
    "os/exec"
    "bufio"
    "time"
)

func main () {

	var penis string
	
	from := "golangbot699@gmail.com"
   	password := "kdxy zlvv unbk rssy"


	loop: for {
		Clear()
		fmt.Println("List:")
		fmt.Println("0: quit \n1: send single email \n2: send to multiple emails \n3: smsOverSmtp")

		fmt.Scanln(&penis)

		switch penis {

		case "0":
			break loop
		case "1":
			var vic string
			fmt.Println("Type victim email here address here:")
			fmt.Scanln(&vic)
			email(from, password, "penis tool", vic)
		case "2":
			multiSender(from, password, "fuck")
		case "3":
			textOverSmtp(from, password)
		}
	}
}

func textOverSmtp(from, password string) {
	var number string
	var text string
	var A string
	var gate string

	gateway := []string{"@msg.koodomobile.com", "@txt.att.net", "@cingularme.com", "@txt.bellmobility.com", "@sms.myboostmobile.com", "@text.scsouth1.com", "@vmobl.com", "@vtext.com", "@msg.telus.com", "@tms.suncom.com", "@messaging.sprintpcs.com", "@email.uscc.net", "@sms.sasktel.com", "@text.republicwireless.com", "@qwestmp.com", "@vtext.com", "@page.nextel.com", "@cspire1.com", "@sms.cricketwireless.net", "@mailmymobile.net", "@cwemail.com"}

	fmt.Println("what is the victims phone number?")
	fmt.Scanln(&number)
	
	fmt.Println("what is the text you want to send?")
	fmt.Scanln(&text)
	
	fmt.Println("Do you know the cell provider gateway? Y/N")
	fmt.Scanln(&A)
	if A == "n" || A == "N" {
		for i := 0; i <= 10; i++ {
		//vic := number + "@msg.koodomobile.com"
	
		vic := number + gateway[i]
		fmt.Println(vic)
		email(from, password, text, vic)
		fmt.Println(vic)
	
		}
	
	} else {
	
	fmt.Println("what is the sms gateway?")
        fmt.Scanln(&gate)

	vic := number + gate

	email(from, password, text, vic)
        fmt.Println(vic)
	
	}
	
	time.Sleep(2 * time.Second)

}

func multiSender(from, password, msg string) {
	var file string
	fmt.Println("Paste your list of emails here, it should just be emails in a list with no interuptions.")
	fmt.Scanln(&file)

	emailArray := fileLines(file)
	numOfEmails := len(emailArray)

	for i := 0; i < numOfEmails; i++ {
		email(from, password, msg, emailArray[i])
	}
	time.Sleep(2 * time.Second)
}

// Main function
func email(from, password, msg, victim string) {
 
    // from is senders email address
     
    // we used environment variables to load the
    // email address and the password from the shell
    // you can also directly assign the email address
    // and the password
    //from := "golangbot699@gmail.com"
    //password := "kdxy zlvv unbk rssy"
 
    // toList is list of email address that email is to be sent.
    toList := []string{victim}
 
    // host is address of server that the
    // sender's email address belongs,
    // in this case its gmail.
    // For e.g if your are using yahoo
    // mail change the address as smtp.mail.yahoo.com
    host := "smtp.gmail.com"
 
    // Its the default port of smtp server
    port := "587"
 
    // This is the message to send in the mail
    //msg := "The nukes are cumming!!!"
    //However I commented that out and it's now declared in the func call!
    // We can't send strings directly in mail,
    // strings need to be converted into slice bytes
    body := []byte(msg)
 
    // PlainAuth uses the given username and password to
    // authenticate to host and act as identity.
    // Usually identity should be the empty string, 
    // to act as username.
    auth := smtp.PlainAuth("", from, password, host)
 
    // SendMail uses TLS connection to send the mail
    // The email is sent to all address in the toList,
    // the body should be of type bytes, not strings
    // This returns error if any occurred.
    err := smtp.SendMail(host+":"+port, auth, from, toList, body)
 
    // handling the errors
    check(err)

    fmt.Println("Successfully sent mail to all user in toList")
}

func fileLines(f string) []string {

        filePath := f
        readFile, err := os.Open(filePath)
        check(err)

        fileScanner := bufio.NewScanner(readFile)
        fileScanner.Split(bufio.ScanLines)
        var fileLines []string


                for fileScanner.Scan() {
                fileLines = append(fileLines, fileScanner.Text())
        }

        readFile.Close()

        return fileLines
}


func Clear() {
        
	cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()


    	//check(err)

}

func check (err error) {

	if err != nil {
        	fmt.Println(err)
        	os.Exit(1)
    	}
}

