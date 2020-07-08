
email-linebuster is a command line tool for ensuring emails are sent at an exact time

## Problem

Some Rental applications may have language such as
```
open application period starts at 3PM Thursday.
Applications should be sent to $EMAIL_ADDR
Applications will be considered first come first served.
Any application submitted prior to 3PM will be assesed an 8 hour penalty
```

Gmail scheduled send generally incurs a lag of 30-45s off the scheduled time, so you can use this instead to send an email from your Gmail closer to the desired time

## Usage

```
Usage of ./bin/email-linebuster:
  -attachement string
        path to attachment
  -email string
        path to email file (txt only)
  -from string
        from email
  -pass string
        gmail app password, see https://support.google.com/accounts/answer/185833?hl=en
  -subj string
        email subj
  -time string
        time to send the email (default "2020-07-07T14:48:22-07:00")
  -to string
        to email
```
