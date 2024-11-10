# Watch_Dogs_tools
I played too much watch dogs, so like the virgin I am, I won't let fiction stay fiction, I am now trying to create a bunch of watch dogs themed hacking tools. All I got right now is a shitty phishing tool though.

You'll want to instal figlet! It's a tool for creating ascii logos. I've added to my hash cracker

-sudo apt install figlet

# Phishing_tool
goPerch is the name of our smtp phishing tool.

How to run:

cd goPerch

go run email.go

You got 3 options:

1: send a single phishing email to one person

2: send a single phishing email to a list of people

3: send a sms message using smtp (meaning you send a text message from your email)

# CatPhish
Have you ever noticed the sheer amount of creeps, weirdos, and losers on discord? well these chumps will click on any almost any link. 
This webserver serves a fake discord image which you can trick people into clicking. You'll want to disguise the link of course.
This is a phishing server which pretends to be an image link, and gets peoples location using this (https://www.w3schools.com/html/html5_geolocation.asp).

which probably isn't very effecient but, it works, on firefox running on a laptop. I don't think it will work on mobile. There's also another version of the phishing server in there for login credentials.

How to use CatPhish:

cd catfish

go run catPhishServer.go



# GoPhish
Really doing alot of phishing tools here. This is another webserver meant to emulate login pages. It's very simple to use and modify.
Just run the server and try localhost:8080/instagram

cd Phishing

go run PhishingServer.go

# Network Scanner
I needed to do something that wasn't a phishing tool. This works well for discovering devices, such as cameras, on network. Even a network as far as serbia.

go run networkScanner.go

# Wifi network guesser
On one hand, I want to write cool names for all these tools, but on the other... It's much easier to creates these self explanatory file names.
Seriously, this one just guesses the password for wifi networks. Great for when you forget your own.

sudo bash wifiBruteForce.sh
