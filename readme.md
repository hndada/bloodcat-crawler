Bloodcat Crawler
===================================
New bloodcat crawler that can download Loved maps!

# How to use
1. Set osu! Songs folder, search keywords and options at config.txt.
2. Put SetIDs that you don't want to download at ban.txt.
3. Double click the .exe file.

# Note
The program's usefulness comes from simplicity and stability of Bloodcat service itself.
Beware that the program burdens great traffic to Bloodcat server, so don't abuse this. 
It would make Bloodcat put bot-detecting features like reCAPTCHA. 

There're dropped features:
- Update all outdated beatmaps (local ranking goes wiped if update happens)
    * Mapset at Bloodcat might not be latest.

- Update omitted beatmaps from a set
    * It harms simplicity of the code.

- Skip downloading maps that mapped by mappers that you banned
    * It is ambiguous and tricky to handle at the crawler.
