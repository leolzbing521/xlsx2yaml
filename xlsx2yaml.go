package main

import (
    "fmt"
    "io/ioutil"
    // "log"
    // "os"
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
    Uid        string
    Externalid string
    Emails     []string
    Org struct {
          Name   string
          Roles  []string
          Space struct {
                  Name  string
                  Roles []string
              }
      }
  }
}



func main() {
  // 1.获取xlsx中增加的行- finish
  // 2.将获取数据按照User结构，插入到user.yaml文件中 
   // struct set value fail
uaa,_ := ioutil.ReadFile("users.yml")
     data := Config{}
// emails := make([]Email, 0)

 data.User.Uid  ="test"
 data.User.Org.Name = "dabing"
 data.User.Org.Space.Name = "YDHL-DEV1"
  data.User.Org.Space.Roles = []string{"role1","role2"}


  // data.User.Uid  ="test"
 fmt.Printf("\n%s",&data)
 //     data.Users.Orgs.Name = "A171616"
    yaml.Unmarshal([]byte(uaa), &data)
    
     // data.Users.Uid ="test1"
    fmt.Printf("--- t:\n%v\n\n", data)
 d, _ := yaml.Marshal(&data)    
 fmt.Printf("--- t dump:\n%s\n\n", string(d))
   // output, _ := yaml.Marshal(uaadata)

   //  content := []byte(output)
   //  ioutil.WriteFile("/home/Project/golang/meta.yaml", content, os.ModePerm)

}

