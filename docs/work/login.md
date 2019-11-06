session:
    client 10.64.40.110:60478 => 10.68.31.11:29005



1. Challenge -> Response

    RobotManager.messageReceived(Challenge ch) 发起 msgQueue 的put
    调用 on(Challenge p) 方法，发送 Response

2. KeyExchange -> KeyExchange

    KeyExchange 加入到msgQueue
    调用 on(KeyExchange p)方法 发送 KeyExchange




3. OnlineAnnounce -> RoleList

4. RoleList_Re -> CreateRole / SelectRole

5. CreateRole_Re -> SelectRole
6. SelectRole_Re -> EnterWorld
7. PlayerChangeGS_Re -> EnterWorld