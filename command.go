package main

import (
  "flag"
  "fmt"
  "time"
  "math/rand"
)

func main(){
  flag.Parse()
  rand.Seed(time.Now().UnixNano())
  r := rand.Intn(3) // 0～2のどれかの値を取得する

  PlayerHand := flag.Args()[0]

  if PlayerHand != "g" && PlayerHand != "p" && PlayerHand != "t" {
    fmt.Println("半角小文字で、g か t か p を入力してください")
    return
  }

  var EnemyHand string
  switch r {
    case 0:
      fmt.Println("相手はグーを出しました")
      EnemyHand = "g"
      break
    case 1:
      fmt.Println("相手はチョキを出しました")
      EnemyHand = "t"
      break
    case 2:
      fmt.Println("相手はパーを出しました")
      EnemyHand = "p"
      break
  }

  
  switch PlayerHand {
    case "p":
      switch EnemyHand {
        case "g":
          fmt.Println("あなたの勝ちです")
          break
        case "t":
          fmt.Println("あなたの負けです")
          break
        case "p":
          fmt.Println("引き分けです")
          break
      }
      break
    case "g":
      switch EnemyHand {
        case "g":
          fmt.Println("引き分けです")
          break
        case "t":
          fmt.Println("あなたの勝ちです")
          break
        case "p":
          fmt.Println("あなたの負けです")
          break
      }
      break
    case "t":
      switch EnemyHand {
        case "g":
          fmt.Println("あなたの負けです")
          break
        case "t":
          fmt.Println("引き分けです")
          break
        case "p":
          fmt.Println("あなたの勝ちです")
          break
      }
      break
  }
}