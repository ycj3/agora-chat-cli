## agchat message send

Send a message via Agora Chat Server

```
agchat message send [flags]
```

### Examples

```
$ agchat message send --users --text --content "Hello, world" --receivers demo_user_1,demo_user_2
$ agchat message send --users --cmd --action "custom_action" --receivers demo_user_1,demo_user_2
$ agchat message send --users --custom --custom-event "custom_event" --custom-exts '{"key1": "value1", "key2": "value2"}' --receivers demo_user_1,demo_user_2
$ agchat message send --users --loc --lat "39.916668" --lon "116.383331" --addr "beijing" --receivers demo_user_1,demo_user_2

```

### Options

```
  -a, --action string         The action of the command message
  -A, --addr string           The address of the location message
      --cmd                   Command type message
  -c, --content string        The content of the text message
      --custom                Custom type message
  -e, --custom-event string   The event of the custom message
  -E, --custom-exts string    The exts of the custom messagen JSON format (e.g., '{"key1": "value1", "key2": "value2"}')
  -g, --groups                Send a message to groups
  -h, --help                  help for send
  -l, --lat string            The latitude of the location message
      --loc                   Location type message
  -L, --lon string            The longitude of the location message
  -r, --receivers strings     Receivers of the message
  -R, --rooms                 Send a message to rooms
  -s, --sender string         Sender of the message (default 'admin')
      --text                  Text type message
  -u, --users                 Send a message to users
```

### Options inherited from parent commands

```
  -v, --verbose   enable verbose output
```

### SEE ALSO

* [agchat message](agchat_message.md)	 - Manage messages

