package main

import (
    "fmt"
    // "io/ioutil"
    // "log"
    "os"
    // "math"
    // "time"
    "gopkg.in/yaml.v2"
    // // "encoding/json"
    // "github.com/tealeg/xlsx"
)


// type Space struct {
//     Name  string
//     Roles []string
// }
// type Org struct {
//     Name   string
//     Roles  []string
//     Spaces []Space
// }
// type Space struct {
//     Name  string
//     Roles []string
// }
// type Org struct {
//     Name   string
//     Roles  []string
//     Spaces []Space
// }

// type User struct {
//     Uid        string
//     Externalid string
//     Emails     []string
//     Orgs       []Org
// }
// type Config struct {
//     Users []User
// }
// var uaa = `
// users:
// - uid: jcalabrese@pivotal.io
//   externalid: uid=jcalabrese,ou=People,dc=homelab,dc=io
//   emails:
//   - jcalabrese@pivotal.io
//   orgs: 
//       - name: org1 
//         roles:
//           - managers
//           - auditors
// `

// type Config struct {
//     Users []User
// }
// type User struct {
//     Uid        string
//     Externalid string
//     Emails     []string
//     Orgs       []Org
// }
// type Org struct {
//     Name   string
//     Roles  []string
//     Spaces []Space
// }
// type Space struct {
//     Name  string
//     Roles []string
// }

type Config struct {
    User struct {
    Uid        string `yaml:"uid"` 
    Externalid string `yaml:"externalid"`
    Emails     []string  `yaml:"emails"`
    Org struct {
          Name   string  `yaml:"name"`
          Roles  []string  `yaml:"roles"`
          Space struct {
                  Name  string  `yaml:"name"`
                  Roles []string  `yaml:"roles"`
              }
      }
  }
}



func main() {
  // 1.获取xlsx中增加的行- finish
  // 2.将获取数据按照User结构，插入到user.yaml文件中 
   // struct set value fail
    // 3.append file 
// uaa,_ := ioutil.ReadFile("users.yml")
// fl, _:= os.OpenFile("users.yml", os.O_APPEND|os.O_CREATE, 0644)
fl, _:= os.OpenFile("users.yml", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
     data := Config{}
      // datas := Config{}
// emails := make([]Email, 0)

 data.User.Uid  ="test"
 data.User.Emails  =[]string{"lizhanbing@bankcomm.com"}
 data.User.Externalid = "lizhanbing@bankcomm.com"
 data.User.Org.Name = "dabing"
 data.User.Org.Roles = []string{"dabing","manager"}
 data.User.Org.Space.Name = "YDHL-DEV1"
  data.User.Org.Space.Roles = []string{"role1","role2"}


  // data.User.Uid  ="test"
 fmt.Printf("\n%s",&data)

    // yaml.Unmarshal(uaa, &datas)
    // datas = append(datas,data)
     // data.Users.Uid ="test1"
    fmt.Printf("--- t:\n%v\n\n", data)
 d, _ := yaml.Marshal(&data)    
 fmt.Printf("--- t dump:\n%s\n\n", string(d))
 defer fl.Close()
 fl.WriteString(string(d))

 // defer fl.Close()
 // fl.Write(d)
  // fmt.Printf("--- t dump:\n%s\n\n", string(uaa))
   // output, _ := yaml.Marshal(uaadata)

   //  content := []byte(output)
   //  ioutil.WriteFile("/home/Project/golang/meta.yaml", content, os.ModePerm)

}

