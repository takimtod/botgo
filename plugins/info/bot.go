package info

import (
  
  "inc/lib"
  
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "bot",
    As:       []string{"ping"},
    Tags:     "info",
   // IsPrefix: true,
    Exec: func(client *lib.Event, m *lib.IMessage) {
     
      m.Reply("Bot aktif")
    },
  })
}
