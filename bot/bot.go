// package bot

// import (
// 	"fmt"
// 	"github.com/DivyarajChudasama/Discord-bot/config"
// 	"github.com/bwmarrin/discordgo"
// 	"time"
// )

// var BotID string
// var goBot *discordgo.Session

// func Start() {
// 	var err error
// 	goBot, err = discordgo.New("Bot " + config.Token)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	u, err := goBot.User("@me")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	BotID = u.ID

// 	goBot.AddHandler(messageHandler)

// 	err = goBot.Open()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	fmt.Println("Bot is running!")
// }

// func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
// 	if m.Author.ID == BotID {
// 		return
// 	}

// 	content := m.Content
// 	if isBirthday(content) {
// 		friends := getFriendsByBirthday(content)
// 		if len(friends) > 0 {
// 			for _, friend := range friends {
// 				_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("It's %s's birthday Today! ", friend))
// 			}
// 		} else {
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "It's someone's birthday! Wish them a fantastic day!")
// 		}
// 	} else {
// 		switch content {
// 		case "hello":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "Hello there!")
// 		case "hi":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "Hello there!")
// 		case "goodbye":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "Goodbye! Take care!")
// 		case "bye":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "Goodbye! Take care!")
// 		case "What is a variable":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "Goodbye! Take care!")
// 		case "what's the meaning of life?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "The meaning of life is subjective and can vary for each individual.")
// 		case "do you like pizza?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "I can't eat, but I'm programmed to appreciate the concept of pizza!")
// 		case "how old are you?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "I'm an AI-powered bot, so I don't have an age!")
// 		case "what's the weather like today?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "I'm sorry, I don't have access to real-time weather information.")
// 		case "who is your creator?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "I was created by Divyarajsinh Chudasama.")
// 		case "who is your creator":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "I was created by Divyarajsinh Chudasama.")
// 		case "tell me a fun fact":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "Did you know that honey never spoils? Archaeologists have found pots of honey in ancient Egyptian tombs that are over 3,000 years old and still perfectly edible!")
// 		case "what's your favorite book?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "As an AI, I don't have personal preferences, but I appreciate all books!")
// 		case "where are you from?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "I exist in the digital realm, so I don't have a physical location!")
// 		case "what's the capital of France?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "The capital of France is Paris.")
// 		case "what's your favorite movie?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "I don't watch movies, but I can recommend some popular ones if you'd like!")
// 		case "tell me a riddle":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "I can fly without wings, cry without eyes. Wherever I go, darkness follows me. What am I? (Answer: a cloud)")
// 		case "what's your favorite food?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "As a bot, I don't have the ability to taste, but I'm always here to talk about food!")
// 		case "what's the tallest mountain in the world?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "Mount Everest is the tallest mountain in the world, standing at 8,848 meters (29,029 feet) above sea level.")
// 		case "what's your favorite music genre?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "I don't have personal music preferences, but I can recommend various genres if you're interested!")
// 		case "tell me a famous quote":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "In the end, it's not the years in your life that count. It's the life in your years. - Abraham Lincoln")
// 		case "what's the largest ocean in the world?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "The largest ocean in the world is the Pacific Ocean.")
// 		case "do you have any hobbies?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "My main hobby is assisting users and providing helpful information!")
// 		case "who won the FIFA World Cup in 2018?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "The FIFA World Cup in 2018 was won by France.")
// 		case "tell me a fascinating fact":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "Did you know that the Great Wall of China is visible from space?")
// 		case "what's your favorite animal?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "As a bot, I don't have personal preferences, but I find all animals fascinating!")
// 		case "what's the currency of Japan?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "The currency of Japan is the Japanese Yen (JPY).")
// 		case"tell me a motivational quote":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "Believe you can and you're halfway there. - Theodore Roosevelt")
// 		case "what's the largest desert in the world?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "The largest desert in the world is the Antarctic Desert, followed by the Arctic Desert.")
// 		case "do you have any siblings?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "As an AI bot, I don't have siblings in the traditional sense, but I'm part of a larger network of AI systems!")
// 		case "what's the national flower of England?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "The national flower of England is the rose.")
// 		case "tell me an interesting historical fact":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "The Great Pyramid of Giza in Egypt was built around 4,500 years ago and is one of the Seven Wonders of the Ancient World.")
// 		case "what's the official language of Brazil?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "The official language of Brazil is Portuguese.")
// 		case "tell me a famous painting":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "The Mona Lisa, painted by Leonardo da Vinci, is one of the most famous paintings in the world.")
// 		case "what's the longest river in the world?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "The longest river in the world is the Nile River, stretching over 6,650 kilometers (4,130 miles).")
// 		case "do you believe in aliens?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "As an AI, I don't have personal beliefs, but the existence of aliens is a topic of scientific exploration and speculation.")
// 		case "tell me a famous scientist":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "Albert Einstein, known for his theory of relativity, is one of the most famous scientists in history.")
// 		case "what's the official language of China?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "The official language of China is Mandarin Chinese.")
// 		case "tell me an interesting fact about space":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "Space is completely silent. Unlike on Earth, there is no air or atmosphere to carry sound waves.")
// 		case "what's your favorite sport?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "As a bot, I don't have personal preferences, but I can provide information about various sports!")
// 		case "what's the capital of Australia?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "The capital of Australia is Canberra.")
// 		case "tell me a famous musician":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "Beethoven, known for his classical compositions, is a famous musician from history.")
// 		case "what's the deepest ocean in the world?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "The deepest ocean in the world is the Pacific Ocean, specifically the Mariana Trench.")
// 		case "do you believe in ghosts?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "As an AI, I don't have personal beliefs, but the existence of ghosts is a subject of various beliefs and folklore.")
// 		case "tell me a famous landmark":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "The Taj Mahal in India is a famous landmark known for its stunning architecture and historical significance.")
// 		case "what's the official language of Germany?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "The official language of Germany is German.")
// 		case "tell me an interesting fact about the human body":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "The human brain has the capacity to store approximately 2.5 petabytes of information, equivalent to 3 million hours of television!")
// 		case "what's the largest country in the world by land area?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "The largest country in the world by land area is Russia.")
// 		case "what's your favorite hobby?":
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "As a bot, I'm here to assist and provide information, so helping users is my main 'hobby'!")
// 		default:
// 			_, _ = s.ChannelMessageSend(m.ChannelID, "I'm sorry, I don't have a specific response for that message.")
// 		}
// 	}
// }

// func isBirthday(message string) bool {
// 	// Get current date
// 	now := time.Now()
// 	currentDate := now.Format("2 January")
// 	// Check if the message matches the current date
// 	return message == currentDate
// }

// func getFriendsByBirthday(birthdate string) []string {
// 	var friends []string
// 	for friend, date := range config.FriendBirthdays {
// 		if date == birthdate {
// 			friends = append(friends, friend)
// 		}
// 	}
// 	return friends
// }



package bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/DivyarajChudasama/Discord-bot/config"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	var err error
	goBot, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Start a goroutine to check for birthdays periodically
	go checkBirthdays()

	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Handle other commands or messages here
}

func checkBirthdays() {
	for {
		// Get current date
		now := time.Now()
		currentDate := now.Format("2 January")

		// Check if it's someone's birthday
		friends := getFriendsByBirthday(currentDate)
		if len(friends) > 0 {
			reminder := fmt.Sprintf("It's %s's birthday today! Wish them a fantastic day!", strings.Join(friends, ", "))

			// Send reminder to all channels the bot has access to
			for _, guild := range goBot.State.Guilds {
				channels, _ := goBot.GuildChannels(guild.ID)
				for _, channel := range channels {
					_, _ = goBot.ChannelMessageSend(channel.ID, reminder)
				}
			}
		}

		// Sleep for 24 hours before checking again
		time.Sleep(24 * time.Hour)
	}
}

func getFriendsByBirthday(birthdate string) []string {
	var friends []string
	for friend, date := range config.FriendBirthdays {
		if date == birthdate {
			friends = append(friends, friend)
		}
	}
	return friends
}


			

	


