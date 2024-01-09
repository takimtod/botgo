package owner

import (
  "inc/lib"
  "fmt"
  "io/ioutil"
  "os"
  "os/exec"
  "path/filepath"

)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "backup",
    As:       []string{"backup"},
    Tags:     "owner",
    IsPrefix: true,
    IsOwner: true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

      dirs, err := ioutil.ReadDir(".")
      if err != nil {
        fmt.Println("Unable to scan directory:", err)
        return
      }

      var filePaths []string
      for _, dir := range dirs {
        if dir.Name() != "main" && dir.Name() != ".cache" && dir.Name() != ".git" {
          filePaths = append(filePaths, filepath.Join(".", dir.Name()))
        }
      }

      // Zip the selected directories and files
      zipCommand := exec.Command("zip", append([]string{"-r", "backup.zip"}, filePaths...)...)
      err = zipCommand.Run()
      if err != nil {
        fmt.Println("Error creating zip file:", err)
        return
      }

      // Send the zip file to the desired location
      err = ioutil.WriteFile("backup.zip", nil, 0644)
      if err != nil {
        fmt.Println("Error writing zip file:", err)
        return
      }

      
      bytes, err := ioutil.ReadFile("backup.zip")
      if err != nil {
        fmt.Println("Error reading file:", err)
        return
      }

      fmt.Println("Byte data:")
      client.SendDocument(m.From, bytes, fmt.Sprintf("%s.zip", "backup"), "in", m.ID)
     os.Remove("backup.zip")
      
    },
  })
}
