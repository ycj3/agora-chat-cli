## agchat fcm test

Test whether the push notification credentials and notification services work properly in FCM

```
agchat fcm test [flags]
```

### Examples

```
# Send a test message to a specific device
$ agchat fcm test --key <service-account-key> --token <device-token>

```

### Options

```
  -h, --help             help for test
  -k, --key string       The service account JSON file
  -m, --message string   JSON string for the push message (default "{\"title\": \"FCM Message\", \"body\": \"This is an FCM notification message!\"}")
  -t, --token string     The device's registration token
```

### Options inherited from parent commands

```
  -v, --verbose   enable verbose output
```

### SEE ALSO

* [agchat fcm](agchat_fcm.md)	 - Test FCM push notifications

