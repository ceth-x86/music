# Music

The project allows track playlists on music services and, based on them, notify of new music releases.

## Quick start

```bash
# add your favourite playlists
music playlist add --service spotify --id 37i9dQZF1DX5wgKYQVRARv
music playlist add --service yandex --id 1004

# sync
music playlist sync

# get the result
music releases list
```

```
  Id                  Artist                                                 Album                                       Date       Type                                                                                                                                          Genres                                                                                                                                                                  Playlist                         
------ ------------------------------------ ------------------------------------------------------------------------ ------------ -------- ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- ----------------------------------------------------------
 2207   Cult Of Luna                         Lay Your Head to Rest                                                    2019-09-05   single   atmospheric sludge,avantgarde metal,drone metal,mathcore,post-doom metal,post-metal,post-rock,progressive metal,progressive sludge,slayer,sludge metal,swedish doom metal                                                                                                             Progressive Metal                                        
 2324   Otherwise                            Lifted                                                                   2019-09-05   single   alternative metal,christian rock,gymcore,hard alternative,nu metal,post-grunge                                                                                                                                                                                                        Rock Hard                                                
 2369   Grimes                               Violence                                                                 2019-09-05   single   alternative dance,art pop,canadian electropop,chillwave,dance pop,dream pop,electropop,escape room,grave wave,indie pop,indietronica,metropopolis,new rave,pop                                                                                                                        Ultimate Indie                                           
 2534   Solence                              Spit It Out                                                              2019-09-05   single   gymcore,post-screamo                                                                                                                                                                                                                                                                  New Metal Tracks                                         
 2552   Overcoats                            The Fool                                                                 2019-09-05   single   electropop,nyc pop                                                                                                                                                                                                                                                                    All New Indie                                            
 2151   Post Malone                          Hollywood's Bleeding                                                     2019-09-06   album    dfw rap,pop,rap                                                                                                                                                                                                                                                                       Rock This                                                
```

## Usage

### CLI commands

| Command                     | Description             | Examples                                                   |
| ----------------------------| ----------------------- | ---------------------------------------------------------- |
| playlist add                | Add playlist            | playlist add --service spotify --id 37i9dQZF1DXcF6B6QPhFDv | 
|                             |                         | playlist add --service yandex --id 1001                    |
| playlist list               | Show list of playlists  | playlist list                                              |
| playlist sync               | Sync playlists          | playlist sync                                              |
| playlist sync --id 4        | Sync playlist           | playlist sync --id 4                                       |
| release list                | Show list of releases   | release list, release list --type album                    |
| release list --type album   | Show list of new albums | release list --type album                                  |

## Roadmap

- Deezer integration
- Send releases to email
