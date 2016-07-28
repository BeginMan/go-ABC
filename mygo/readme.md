We need:

- login
- list all rooms
- create rooms
- add rooms
- play
- chat
- exit
- logout

原则上一台服务器可创建上百万个 goroutine, 也就是可以支持上百万room

Person {
    唯一ID
    userName,
    Level,
    Score
}

