package convert

import (
  "inc/lib"
"fmt"
  "os"
  
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "(sm|smeme)",
    As:       []string{"smeme"},
    Tags:     "convert",
    IsPrefix: true,
    IsMedia:  true,
     IsWaitt:  true,
     IsQuerry: true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

      byte, _ := client.WA.Download(m.Media)

    randomJpgImg := "./" + lib.GetRandomString(5) + ".jpg"
    if err := os.WriteFile(randomJpgImg, byte, 0600); err != nil {
        fmt.Printf("Failed to save image: %v", err)
        return
    }
    //log.Printf("Saved image in %s", randomJpgImg)
    url, err := lib.Upload(randomJpgImg)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
       res := "https://api.memegen.link/images/custom/-/"+m.Querry+".jpg?background="+url
      
      bytes, err := client.GetBytes(res)
      if err != nil {
         fmt.Println("Error:", err)
        return
      }
      client.SendImage(m.From, bytes, ".s", m.ID)
      os.Remove(randomJpgImg)

    },
  })
}