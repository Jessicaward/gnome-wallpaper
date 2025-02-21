# gnome-wallpaper
Wallpaper switcher for GNOME linux environments

## How to run
I haven't added this to any package managers or anything (yet).

So to run, download the source code, navigate to it in a terminal. Then run `go build`.

To run the program, you'll just need to pass your image directory (like `/path/to/exec/gnome-wallpaper /home/jess/wallpapers`).

### How to *actually* run it
So, you don't want to run this every 5 mins lol, so register a cronjob like so.

Add the following to your crontab list (`crontab -e`).

```cron
0 * * * * env DISPLAY=:0 DBUS_SESSION_BUS_ADDRESS=unix:path=/run/user/$(id -u)/bus XDG_SESSION_TYPE=wayland /path/to/exec/gnome-wallpaper /path/to/wallpapers >> ~/wallpaper_changer.log 2>&1
```

## Troubleshooting
If it's unable to change your wallpaper, make sure you don't have a trailing slash in the directory (i should prolly fix that).