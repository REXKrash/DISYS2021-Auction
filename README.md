# DISYS2021-Auction

This is a primitive implementation of a distributed auction system using passive replication (primary-backup replication).
The primary replica manager is selected through a token-ring protocol.

You can start the 3 replica manager nodes by using these commands in 3 separate command line windows:
```
go run .\replica-manager\server.go 0
go run .\replica-manager\server.go 1
go run .\replica-manager\server.go 2
```

A client node can be started using this command:
```
go run .\client\client.go <name>
```

As a client node you can then participate in the auction using the following commands:
```
/result
/bid <amount>
```

You can exit any of the nodes by pressing CTRL + C
