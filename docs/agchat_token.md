## agchat token

Generate and parse agora tokens

```
agchat token [flags]
```

### Examples

```
# Generate token for a specific user
$ agchat token --user <user-id>

# Parse an agora token
$ agchat token --parse <user-token>

	#Service type
		ServiceTypeRtc       = 1
		ServiceTypeRtm       = 2
		ServiceTypeFpa       = 4
		ServiceTypeChat      = 5
		ServiceTypeEducation = 7

	#Rtc
		PrivilegeJoinChannel        = 1
		PrivilegePublishAudioStream = 2
		PrivilegePublishVideoStream = 3
		PrivilegePublishDataStream  = 4

	#Rtm
	#Fpa
		PrivilegeLogin = 1

	#Chat
		PrivilegeChatUser = 1
		PrivilegeChatApp  = 2

	#Education
		PrivilegeEducationRoomUser = 1
		PrivilegeEducationUser     = 2
		PrivilegeEducationApp      = 3

```

### Options

```
  -h, --help           help for token
  -p, --parse string   parse an agora token
  -u, --user string    generate a new user token for use in SDK APIs
```

### Options inherited from parent commands

```
  -v, --verbose   enable verbose output
```

### SEE ALSO

* [agchat](agchat.md)	 - Agora Chat CLI

