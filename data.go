package main

// Replace with your data by using -ldflags like this:
//
// tinygo flash -target pybadge -ldflags="-X main.YourName=@myontwitter -X main.YourTitleA1='Amazing human' -X main.YourTitleA2='also kind'"
//
// See Makefile for more info.
var (
	YourName, YourCompany, YourTitleA1, YourTitleA2, YourTitleB1, YourTitleB2 string
	YourMarqueeTop, YourMarqueeMiddle, YourMarqueeBottom, YourQRText          string
)

const (
	DefaultName          = "@TinyGolang"
	DefaultTitleA1       = "Go Compiler"
	DefaultTitleA2       = "Small Places"
	DefaultMarqueeTop    = "This badge"
	DefaultMarqueeMiddle = "runs"
	DefaultMarqueeBottom = "TINYGO"
	DefaultQRText        = "https://gopherbadge.com"
	DefaultTitleB1       = "I enjoy"
	DefaultTitleB2       = "TINYGO"
)

func setCustomData() {
	if YourName == "" {
		YourName = DefaultName
	}

	if YourTitleA1 == "" {
		YourTitleA1 = DefaultTitleA1
	}

	if YourTitleA2 == "" {
		YourTitleA2 = DefaultTitleA2
	}

	if YourTitleB1 == "" {
		YourTitleB1 = DefaultTitleB1
	}

	if YourTitleB2 == "" {
		YourTitleB2 = DefaultTitleB2
	}

	if YourMarqueeTop == "" {
		YourMarqueeTop = DefaultMarqueeTop
	}

	if YourMarqueeMiddle == "" {
		YourMarqueeMiddle = DefaultMarqueeMiddle
	}

	if YourMarqueeBottom == "" {
		YourMarqueeBottom = DefaultMarqueeBottom
	}

	if YourQRText == "" {
		YourQRText = DefaultQRText
	}

	if YourCompany == "" {
		YourCompany = "TinyGo"
	}
}

type Talk struct {
	startHour, endHour, line1, line2, line3 string
}

type Day struct {
	title string
	talks []Talk
}

var scheduleData = []Day{
	{"Monday August 25 (TZ EDT)",
		[]Talk{
			{"18:00", "20:00", "WOMEN WHO GO MEETUP", "Angelica Hill, Cassie Coyle,", "Kate Pong & Samantha Coyle"},
			{"", "", "", "", ""},
			{"", "", "", "", ""},
			{"", "", "", "", ""},
		},
	},
	{"Tuesday August 26 (TZ EDT)",
		[]Talk{
			{"07:00", "11:00", "Morning Coffee Break", "Foyer, Level 4, North Javits", ""},
			{"07:00", "17:30", "Registration & Gopher Gear Shop", "Foyer, Level 5, North Javits", ""},
			{"08:30", "12:30", "WORKSHOP: Building and deployin AI", "agents with Go", "Johnny Boursiquot"},
			{"08:30", "12:30", "WORKSHOP: Profiling and optimizing", "Go programs", "Cory LaNou"},
			{"08:30", "17:30", "WORKSHOP: AI-powered systems in Go:", "RAG & tool calling", "Florin Patan"},
			{"08:30", "17:30", "WORKSHOP", "Ultimate Go", "Miki Tebeka"},
			{"08:30", "17:30", "WORKSHOP", "Ultimate software design and engineering", "William Kennedy"},
			{"08:30", "17:30", "", "HALLWAY TRACK", "Foyer, Level 4, North Javits"},
			{"09:00", "12:00", "Contributors summit", "Ian Cottrell, Michael Knyszek", "& Michael Pratt"},
			{"09:00", "17:00", "Challenge series: how deep", "does the code go?", "Neil Primmer & Benji Vesterby"},
			{"10:00", "16:00", "TINYGO HARDWARE HACK SESSION", "Ron Evans & Patricio Whittingslow", "Room 409/411, Level 4, North Javits"},
			{"10:00", "16:00", "Community Roundtables", "Room 404/405, Level 4, North Javits", ""},
			{"12:00", "13:30", "Grab & Go lunch", "Foyer, Level 4, North Javits", ""},
			{"13:30", "16:30", "Community day:", "Meet the Go team", "Room 413, Level 4, North Javits"},
			{"13:30", "17:30", "WORKSHOP: Advanced Go for", "experienced programmers", "Johnny Boursiquot"},
			{"13:30", "17:30", "WORKSHOP: Advanced testing", "in Go programs", "Cory LaNou"},
			{"13:30", "17:30", "WORKSHOP: Go concurrency: debugging", "goroutines and channels", "Derek Parker"},
			{"15:00", "17:00", "Afternoon break", "Foyer, Level 4, North Javits", ""},
			{"17:00", "19:00", "Meetup organizers (past, present & future)", "Benji Vesterby, Cassie Coyle", "& Samantha Coyle"},
			{"18:00", "20:00", "Neurospicy meetup", "Andy Haskell, Charles Pustejovsky,", "Kaylyn Gibilterra & Erik St. Martin"},
			{"20:00", "22:00", "RainGo Connection Meetup", "Angelica Hill, Kris Brandow", "& Zaq? Question"},
		},
	},
	{"Wednesday August 27 (TZ EDT)",
		[]Talk{
			{"08:00", "10:00", "Coffee with the exhibitors", "Room 503/504, Level 5, North Javits", ""},
			{"08:00", "17:00", "Challenge series: how deep", "does the code go?", "Neil Primmer & Benji Vesterby"},
			{"08:00", "17:45", "GopherCon exhibition", "Room 503/504, Level 5, North Javits", ""},
			{"08:00", "17:45", "", "HALLWAY TRACK", "Foyer, Level 5, North Javits"},
			{"09:05", "09:30", "Go's next frontier", "Cameron Balahan", ""},
			{"09:35", "10:00", "Demystifying AI agents:", "A new primitive for flow control", "Solomon Hykes"},
			{"10:05", "10:30", "Goroutines and cells:", "Lessons in goal-directed systems", "Carlisia Campos"},
			{"10:30", "11:15", "Morning beverage break", "Room 503/504, Level 5, North Javits", ""},
			{"10:45", "11:00", "Go: a taylor-made language", "for AWS lambda", "Capital One"},
			{"11:15", "11:40", "The code you reviewed is", "not the code you built", "Jess McClintock"},
			{"11:45", "12:10", "My protobuf module is faster", " than yours (I cheated)", "Tom Lyons"},
			{"12:15", "12:40", "Building a decentralize social", "media app with Go and ATProto", "Gautam Dey"},
			{"12:15", "12:40", "Building a decentralize social", "media app with Go and ATProto", "Gautam Dey"},
			{"12:40", "14:25", "Lunch service", "Room 503/504, Level 5, North Javits", "Bayer"},
			{"13:40", "14:25", "Lightning talks", "Angelica Hill, Kaylyn Gibilterra,", "James Heller & Sahid Velji"},
			{"14:25", "14:50", "Analysis and transformation tools", "for Go codebase modernization", "Alan A. Donovan"},
			{"14:55", "15:20", "Plowing through data:", "building flexible pipelines with Go", "Mindy Ratcliff"},
			{"15:20", "16:00", "Afternoon beverage break", "Room 503/504, Level 5, North Javits", ""},
			{"16:00", "16:25", "Understanding escape analysis", "to speed up your code", "PJ Malloy"},
			{"16:30", "16:55", "Scaling LLMs with Go: production patterns", "for handling millions of AI requests", "John Wang"},
			{"17:00", "17:45", "Porting the TypeScript compiler", "to Go for a 10x speedup", "Jake Bailey"},
			{"18:00", "21:00", "Rooftop reception with Skool", "Rooftop, 570 10th Avenue", ""},
			{"19:00", "21:00", "Asian alliance meetup", "Arjun Malhotra, Ellen Gao,", "Hannah Kim & Madhav Jovrajani"},
			{"19:00", "21:00", "United Fo meetup", "Edmondo Porcu, Johnny Boursiquot", "& Kris BRandow"},
		},
	},
	{"Thursday August 28 (TZ EDT)",
		[]Talk{
			{"08:00", "10:00", "Coffee with the exhibitors", "Room 503/504, Level 5, North Javits", ""},
			{"08:00", "16:00", "Challenge series: how deep", "does the code go?", "Neil Primmer & Benji Vesterby"},
			{"08:00", "16:00", "GopherCon exhibition", "Room 503/504, Level 5, North Javits", ""},
			{"08:00", "17:30", "", "HALLWAY TRACK", "Foyer, Level 5, North Javits"},
			{"08:00", "18:00", "Attendee services & Gopher Gear Shop", "Foyer, Level 5, North Javits", ""},
			{"08:40", "08:55", "From chaos to cohesion:", "a community of practive story", "Capital One"},
			{"09:00", "09:05", "", "WELCOME BACK, GOPHERS!", ""},
			{"09:05", "09:30", "An operating system in Go", "Patricio Whittingslow", ""},
			{"09:35", "10:05", "Go's trace tooling", "and concurrency", "William Kennedy"},
			{"10:10", "10:35", "Profiling request latency", "with critical path analysis", "Felix Geisendorfer"},
			{"10:35", "11:05", "Morning beverage break", "Room 503/504, Level 5, North Javits", ""},
			{"11:05", "11:30", "Invisible insight: strategies for", "auto-instrumenting go apps", "Hannah Kim"},
			{"11:35", "12:00", "Next-gen AI tooling", "with MCP servers in Go", "Katie Hockman"},
			{"12:05", "12:30", "Supercharging ML pipelines", "with Go", "Vaidehi Thete"},
			{"12:30", "14:15", "Lunch service", "Room 503/504, Level 5, North Javits", ""},
			{"13:30", "14:15", "Lightning talks", "Angelica Hill, Kaylyn Gibilterra,", "James Heller & Sahid Velji"},
			{"14:15", "14:45", "Go faster: integrating CUDA", "in Go for GPU acceleration", "Sam Burns"},
			{"14:45", "15:15", "Go plays nice with your computer", "race detection and freedom", "Reghav Roy"},
			{"15:15", "15:45", "Afternoon beverage break", "Room 503/504, Level 5, North Javits", ""},
			{"15:45", "16:10", "Advancing Go garbage collection", "with Green Tea", "Michael Knyszek"},
			{"16:15", "16:40", "The Go cryptography state of the union", "Filippo Valsorda", ""},
			{"16:45", "17:30", "AI & Go: opportunities & challenges", "Johnny Boursiquote, David Soria, Gari Singh,", "Ian Cottrell, Jaana Dogan & Solomon Hykes"},
		},
	},
}

type Scene struct {
	description, optionA, optionB, optionC string
	sceneA, sceneB, sceneC                 int
}

var sceneData = []Scene{
	{
		"As usually, you are reading Golang Weekly, among very useful information you notice the CFP for GopherCon is open! It will be in New York, 25-28th August. What do you want to do?",
		"Oh no! you don't have any idea for a talk.",
		"Let's talk about AI, it's the new pink!",
		"You already have the slides for a TinyGo talk.",
		1, 2, 3,
	},
	{
		"It's ok, you still take advantage of the date and got an early-bird ticket at a discounted price:",
		"Self-payin gopher (385$)",
		"Corporate gopher (500$)",
		"Premium gopher + workshops (850$)",
		4, 4, 4,
	},
	{
		"You worked hard on your slides and send your abstract in time. A few days passed and received the bad news that your talk wasn't accepted, " +
			"nobody is interested anymore in AI. But they offer you a discount price for the tickets:",
		"Self-payin gopher (335$)",
		"Corporate gopher (450$)",
		"Premium gopher + workshops (770$)",
		4, 4, 4,
	},
	{
		"Flying drones, hens, LEDs, lasers, music,... who will not like such a talk about TinyGo? Your talk of course was accepted and your trip is fully paid.",
		"YAY!",
		"Super yay!",
		"OMG!!1! a dream come true",
		4, 4, 4,
	},
	{
		"BEEP BEEP BEEP. It's your alarm clock. Today is the day, your need to prepare and go to the airport, New York and a bunch of gophers are waiting for you.",
		"Take a taxi to the airport. No time to waste!",
		"Have Liam's signature pancakes.",
		"Sleep a bit more",
		5, 6, 4,
	},
	{ // 5
		"You took a taxi, arrived at the airport with enough time, there was no need to rush that much. When you go to the check-in baggage desktop you realize you forgot your suitcase with everything. You are only wearing your unicorn-pijama.",
		"No time to go back, you board the plane.",
		"You shop something at the airport's shop.",
		"You don't care. YOLO! Go to your gate.",
		7, 8, 7,
	},
	{
		"You have the best breakfast ever and feel full of energy, so much you pick up your suitcase and RUN to the airport. You pass the security check without issues.",
		"Better not waste time, board the plane.",
		"You shop something at the airport's shop.",
		"A quick visit to the toilet before boarding.",
		7, 8, 7,
	},
	{
		"You are finally at the plane, go up to your seat: 27B. Your seat neighborg looks like a gopher too, who is holding a shark plushie and a network switch.",
		"Ask about the shark.",
		"Ask about the network switch.",
		"Goeiemorgen.",
		9, 10, 11,
	},
	{
		"You've bought the most expensive suit and bowtie of your life, but holy guacomole, what a nice and superb suit. You look awesome.",
		"Go to the gate and board the plane",
		"Go to the gate and board the plane",
		"Go to the gate and board the plane",
		7, 7, 7,
	},
	{
		"-\"This shark? It's name is Blahaj and is the administrator of our mastodon instance. This is probably the weirdest train I've ever been in, I've never seen a train with wings before!\"",
		"Yeah... sure... ",
		"Oh cool",
		"Me neither",
		12, 12, 12,
	},
	{ // 10
		"-\"This switch? I need it for the uplink of our mastodon instance. This is probably the weirdest train I've ever been in, I've never seen a train with wings before!\"",
		"Yeah... sure... ",
		"Oh cool",
		"Me neither",
		12, 12, 12,
	},
	{
		"-\"Dit is waarschijnlijk de vreemdste trein waar ik ooit in heb gezeten, ik heb nog nooit een trein met vleugels gezien!\"",
		"Yeah... sure... ",
		"Oh cool",
		"Ik ook niet",
		12, 12, 12,
	},
	{
		"The plane finally landed. You check in the hotel and have a few hours left to visit the city. New York is famous for:",
		"Hotdog is a matter of pride for New Yorkers",
		"You know the lyrics of all JAY-Z's songs",
		"The Statue of Liberty is a must",
		13, 14, 15,
	},
	{
		"You ask the concierge at the hotel about the best Hotdog in the city and go there. You make a mess of yourself, but it was delicious. When you start leaving the place, a weird looking person, too clean to be a bum, but with crazy person vibes and a tinfoil hat approaches. He tries to stop you.",
		"You run as fast as you can",
		"Try to ignore him",
		"Give him 5USD & proceed to listen to him",
		16, 16, 17,
	},
	{
		"Unfortunately there is no JAY-Z concert today so you decide to go for a walk around the city to make time. When walking through a dark alley a weird looking person, too clean to be a bum, but with crazy person vibes and a tinfoil hat approaches. He tries to stop you.",
		"You run as fast as you can",
		"Try to ignore him",
		"Give him 5USD & proceed to listen to him",

		16, 16, 17,
	},
	{ // 15
		"You visit the Statue of Liberty, it's amazong. While walking back to the hotel, you go through a dark alley and a weird looking person, too clean to be a bum, but with crazy person vibes and a tinfoil hat approaches. He tries to stop you.",
		"You run as fast as you can",
		"Try to ignore him",
		"Give him 5USD & proceed to listen to him",
		16, 16, 17,
	},
	{
		"-\"Hey you... yeah " + YourName + "! " + YourName + " who works at " + YourCompany + " listen to me, I'm Ron Evans!, but the Ron Evans from the future!!, the year 2053, the last gopher. You need to listen carefully I have very little time, remember the code : <<IDKFA>> You are the only hope to save the planet and the humanity! (truust me)\" ... zooosh... and he disappeared",
		"Wait! What just happened?",
		"IDKFA... IDKFA... IDKFA... you try to remember",
		"....",
		18, 18, 18,
	},
	{
		"-\"Oh hello " + YourName + "! Thank you for listening to me, I'm Ron Evans, but the Ron Evans from the future!!, the year 2053, the last gopher. You need to listen carefully I have very little time, remember the code : <<IDKFA>> You are the only hope to save the planet and the humanity! (truust me)\" ... zooosh... and he disappeared",
		"Wait! What just happened?",
		"IDKFA... IDKFA... IDKFA... you try to remember",
		"....",
		18, 18, 18,
	},
	{
		"After this misterious encounter, you want to go back to the hotel as fast as you can. In front of you there is a group of friendly gophers. Thankfully they are going to the same hotel and can carry you. But each one has it's own method of transportation.",
		"You truust Francesc with his bicycle",
		"Takasago can carry you in his motocross bike",
		"Share an e-scooter with Manolo",
		19, 19, 19,
	},
	{
		"You all arrive at the hotel and decided to take a drink at the bar. There, you encounter Natalie and Mat recording a GoTime podcast, they ask you about your unpopular opinion:",
		"It's pronounced data, not data",
		"Rabbits are not rodent, they are lagomorphs",
		"Cheese cake is the best dessert",
		20, 21, 22,
	},
	{ // 20
		"-\"Hey wonderful listeners, we welcome " + YourName + " to our program, the gopher who states that data should be pronounced data instead of data. What do you think Natalie? I'm a bit socked myself. And with this we end our episode for today\"",
		"Time to go to bed",
		"Time to go to bed",
		"Time to go to bed",
		23, 23, 23,
	},
	{
		"-\"Hey wonderful listeners, we welcome " + YourName + " to our program, the gopher who states that rabbits are ... lagomorphs? Isn't that the Alien movie? That doesn't sound like an opinion but more like a fact. What do you think Mat? I'm a bit socked myself. And with this we end our episode for today\"",
		"Time to go to bed",
		"Time to go to bed",
		"Time to go to bed",
		23, 23, 23,
	},
	{
		"-\"Hey wonderful listeners, we welcome " + YourName + " to our program, the gopher who states that the best dessert is the cheesecake. I've to say I totally agree with that. What do you think Natalie? And with this we end our episode for today\"",
		"Time to go to bed",
		"Time to go to bed",
		"Time to go to bed",
		23, 23, 23,
	},
	{
		"After a long day, you arrive to your room, your bed for the next days is waiting for you. You close your eyes. Looks like it's just a second, but wake up fully rested and ready for the first day of GopherCon. You skip breakfast because I'm too tired of adding options.",
		"-----",
		"Go to North Javits",
		"(event location)",
		24, 24, 24,
	},
	{
		"You arrive at the door. A familiar face greets you, it's the wonderful Angelica. You pick up your new TinyGo powered e-ink badge and look at it.",
		"-----",
		"Look at your badge",
		"-----",
		25, 25, 25,
	},
	{ // 25
		"",
		"-----",
		"-----",
		"-----",
		26, 26, 26,
	},
	{
		"You notice there's a SCHEDULE function in your badge, you could navigate today's schedule and pick your favourite talks you want to attend.",
		"An Operating System in Go - Patricio Whittingslow",
		"Goroutines and cells - Carlisia Campos",
		"Next-gen AI tooling - Katie Hockman",
		27, 27, 27,
	},
	{ // left empty, redirect to proper scene according the clothes.
		"",
		"",
		"",
		"",
		28, 29, 30,
	},
	{
		"The talk was super fun and you learned a lot. You are preparing to go to the next talk but some people stops you. They hand you a book and ask you to sign it, they have confused you with someone else. Since you are sporting your unicorn pijama, they think you are Aurelie, the famous Go book author.",
		"Sign it!",
		"Try to explain you are not her",
		"Je ne parle pas français",
		31, 31, 31,
	},
	{
		"The talk was super fun and you learned a lot. You are preparing to go to the next talk but some people stops you. They hand you a martini and ask you for crypto investment advice, they have confused you with someone else. Since you are sporting the faboulous suit and bow tie, they think you are Tanguy Herman, the best dressed gopher in the world.",
		"Shaken, not stirred",
		"Buy high, sell low... or the other way around",
		"Try to explain you are not her",
		31, 31, 31,
	},
	{ // 30
		"The talk was super fun and you learned a lot. You are preparing to go to the next talk but you find some old friends and catch up with your lives. You talked so much you missed the next talk and it's lunch time already.",
		"Yay! Veggie sandwiches.",
		"This quinoa salad is top.",
		"Not sure what I'm eating but it's superb.",
		32, 32, 32,
	},
	{
		"You try to explain they are mistaken to no avail. Anyway you talked so much you missed the next talk and it's lunch time already.",
		"Yay! Veggie sandwiches.",
		"This quinoa salad is top.",
		"Not sure what I'm eating but it's superb.",
		32, 32, 32,
	},
	{ // REDIRECT IN CASE  YOU GAVE A TALK OR NOT
		"",
		"",
		"",
		"",
		33, 34, 34,
	},
	{
		"After lunch it's time for your talk: flying drones, hens, LEDs, lasers, music,... you left the auditorium in awe, asking for more. Time for the grand finale, you only have one shot, one opportunity:",
		"FIREWORKS never fail.",
		"Launch a weather ballon.",
		"One word: Robot-laser-tag.",
		35, 36, 37,
	},
	{
		"After lunch it's time for another talk. The Ron Evans from your timeline is flying some drones inside the auditorium, one is going crazy and is going to attack some gophers. You could help him by rebooting the drone remotely, but need to introduce the reboot CODE:",
		"IDGAF",
		"IDKFA",
		"IDDQD",
		38, 39, 38,
	},
	{ // 35
		"Fireforks inside an auditorium? Yeah why not? Not sure who approved this but you need to introduce the launch CODE:",
		"IDGAF",
		"IDKFA",
		"IDDQD",
		38, 39, 38,
	},
	{
		"You all go outside, a big weather balloon is waiting on the ground. Introduce the launch CODE to initiate the countdown:",
		"IDGAF",
		"IDKFA",
		"IDDQD",
		38, 39, 38,
	},
	{
		"Autonomous robots armed with lasers? Yeah why not? Not sure who approved this but you need to introduce the login CODE for them to boot up:",
		"IDGAF",
		"IDKFA",
		"IDDQD",
		38, 39, 38,
	},
	{
		"Wrong code, rememmber what that crazy person yesterday was telling you. You have another opportunity:",
		"IDDQD",
		"IDKFA",
		"IDGAF",
		38, 39, 38,
	},
	{
		"Correct code, sequence initiated. YAY! The talk is a big success, #TinyGo is trending topic. GopherCon is coming to an end and only the social mixer is left to attend. You talked with many gopher there, laugh at bad tech jokes and make some friends.",
		"Click A to continue.",
		"Click B to continue.",
		"Click C to continue.",
		40, 40, 40,
	},
	{ // 40
		"This is the end of this adventure. Continue to the next screen to see some stats of your adventure.",
		"Click A to continue.",
		"Click B to continue.",
		"Click C to continue.",
		41, 41, 41,
	},
	{ // STATS PAGE
		"",
		"Click A to continue.",
		"Click B to continue.",
		"Click C to continue.",
		42, 42, 42,
	},
	{ // QR PAGE
		"", "", "", "",
		43, 43, 43,
	},
	{ // THE END
		"Thank you for playing this adventure. I hope you enjoyed playing it as much as I (CONEJO) did making it. I hid some easter eggs (not very well hidden) and references. Feel free to share your opinion about it in person or online at @conejo@social.tinygo.org",
		"Click A to continue.",
		"Click B to continue.",
		"Click C to continue.",
		44, 44, 44,
	},
	{ // THE REAL END
		"", "", "", "",
		45, 45, 45,
	},
}
