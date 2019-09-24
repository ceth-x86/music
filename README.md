# Music

The project allows track playlists on music services and, based on them, notify of new music releases.

## Quick start

```bash
# add your favourite playlists
music playlist add --service spotify --id 37i9dQZF1DX5wgKYQVRARv

# sync
music playlist sync

# get the result
music releases list
```

```
Id        Artist                 Album            Type                                                                                     Genres                                                                                   
---- ------------------- ----------------------- -------- ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------
 59   Voyager             Colours                 single   australian metal                                                                                                                                                          
 60   Cult Of Luna        Lay Your Head to Rest   single   atmospheric sludge,avantgarde metal,drone metal,mathcore,post-doom metal,post-metal,post-rock,progressive metal,progressive sludge,slayer,sludge metal,swedish doom metal 
 61   Tides From Nebula   Ghost Horses            single   cosmic post-rock,instrumental post-rock,polish post-rock,post-metal,post-rock                                                                                             
 62   Voyager             Water over the Bridge   single   australian metal                                                                                                                                                          
 63   TOOL                Fear Inoculum           album    alternative metal,art rock,nu metal,post-grunge,post-metal,progressive metal,progressive rock,rock                                                                        
 64   Klone               Le Grand Voyage         album                                                                                                                                                                              
 65   Leprous             Below                   single   avantgarde metal,djent,jazz metal,norwegian metal,progressive metal                                                                                                       
 66   Vanden Plas         The Ghost Xperiment     single   german metal,neo classical metal,neo-progressive,progressive metal
```

## Usage

### CLI commands

| Command                     | Description             | Examples                                                   |
| ----------------------------| ----------------------- | ---------------------------------------------------------- |
| playlist add                | Add playlist            | playlist add --service spotify --id 37i9dQZF1DXcF6B6QPhFDv | 
| playlist list               | Show list of playlists  | playlist list                                              |
| playlist sync               | Sync playlists          | playlist sync                                              |
| playlist sync --id 4        | Sync playlist           | playlist sync --id 4                                       |
| release list                | Show list of releases   | release list, release list --type album                    |
| release list --type album   | Show list of new albums | release list --type album                                  |

## Roadmap

- Deezer integration
- Yandex.Music integration
