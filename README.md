# Game Log Parser

This is a program written in Go to analyze game logs and generate a report with statistics for each game.

## System requirements

- Go (version 1.20.6 or higher)

## How to use

1. Make sure you have Go installed on your system.
2. Clone the repository to your local environment:

```bash
git clone https://github.com/marcelotoledo5000/game-log.git
cd game-log
```

3. Run the program, providing the path of the log file as an argument:

```bash
go run main.go path/to/file.log
```

For example:

```bash
go run main.go log/temp.log
```

## Sample log

```log
 20:37 ------------------------------------------------------------
 20:37 InitGame: \sv_floodProtect\1\sv_maxPing\0\sv_minPing\0\sv_maxRate\10000\sv_minRate\0\sv_hostname\Code Miner Server\
 20:38 ClientConnect: 2
 20:38 ClientUserinfoChanged: 2 n\Isgalamido\t\0\model\uriel/zael\hmodel\uriel/zael\g_redteam\\g_blueteam\
 20:38 ClientBegin: 2
 20:42 Item: 2 item_armor_body
 20:54 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT
 20:59 Item: 2 weapon_rocketlauncher
 21:07 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT
 21:10 ClientDisconnect: 2
 21:15 ClientConnect: 2
 21:17 ClientUserinfoChanged: 2 n\Isgalamido\t\0\model\uriel/zael\hmodel\uriel/zael\g_redteam\\g_blueteam\
 21:17 ClientBegin: 2
 21:34 Item: 2 ammo_rockets
 21:42 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT
 21:49 Item: 2 weapon_rocketlauncher
 21:51 ClientConnect: 3
 21:51 ClientUserinfoChanged: 3 n\Dono da Bola\t\0\model\sarge/krusade\hmodel\sarge/krusade\g_redteam
 21:53 ClientUserinfoChanged: 3 n\Mocinha\t\0\model\sarge\hmodel\sarge\g_redteam\\g_blueteam\\c1\4\c2\5\hc\95\w\0\l\0\tt\0\tl\0
 21:53 ClientBegin: 3
 22:04 Item: 2 ammo_rockets
 22:06 Kill: 2 3 7: Isgalamido killed Mocinha by MOD_ROCKET_SPLASH
 22:11 Item: 2 item_quad
 22:11 ClientDisconnect: 3
 22:18 Kill: 2 2 7: Isgalamido killed Isgalamido by MOD_ROCKET_SPLASH
 22:27 Item: 2 ammo_rockets
 22:40 Kill: 2 2 7: Isgalamido killed Isgalamido by MOD_ROCKET_SPLASH
 22:45 Item: 2 item_armor_body
 25:41 Kill: 1022 2 19: <world> killed Isgalamido by MOD_FALLING
 25:50 Item: 2 item_armor_combat
 25:52 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT
 26:09 Item: 2 weapon_rocketlauncher
 26:10 ShutdownGame:
 20:37 ------------------------------------------------------------
```

## Sample report

```json
"game_1": {
  "total_kills": 8,
  "players": ["Isgalamido", "Mocinha"],
  "kills": {
    "Isgalamido": -4,
  },
  "kills_by_means": {
    "MOD_TRIGGER_HURT": 4,
    "MOD_ROCKET_SPLASH": 3,
    "MOD_FALLING": 1,
  }
}
Player Ranking:
1. Isgalamido: -4 kills
```
