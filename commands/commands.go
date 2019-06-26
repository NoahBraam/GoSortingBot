package commands

import (
	"fmt"
	"sorting_bot/embed"

	"github.com/bwmarrin/discordgo"
)

// HelpCommand is the command to display help message
func HelpCommand(sess *discordgo.Session, msg *discordgo.Message) {
	messageStr := mentionUser(msg) + ", hi I'm the Sorting Hat!"
	sendMessage(sess, msg.ChannelID, messageStr)
}

// HouseCommand Displays info about the houses
func HouseCommand(sess *discordgo.Session, msg *discordgo.Message) {
	houseStr := "There are 4 different houses: Gryffindor, Hufflepuff, Ravenclaw and Slytherin"
	sendMessage(sess, msg.ChannelID, houseStr)
	houseStr = "To learn more about a particular house, use !slytherin, !ravenclaw, !hufflepuff or !gryffindor"
	sendMessage(sess, msg.ChannelID, houseStr)
}

// GryffindorCommand displays info important to the Gryffindor house
func GryffindorCommand(sess *discordgo.Session, msg *discordgo.Message) {

	embed := embed.NewEmbed().SetTitle("Gryffindor House").
		SetColor(7602177).
		SetURL("https://www.pottermore.com/collection/all-about-gryffindor").
		SetThumbnail("https://www.hp-lexicon.org/wp-content/uploads/2015/08/gryffindor-shield-200x0-c-default.jpg").
		AddField("About", "Gryffindor is one of the four Houses of Hogwarts School of Witchcraft and Wizardry and was founded by Godric Gryffindor. Godric instructed the Sorting Hat to choose students possessing characteristics he most valued, such as courage, chivalry, and determination, to be Sorted into his house. The emblematic animal is a lion, and its colours are scarlet and gold. Sir Nicholas de Mimsy-Porpington, also known as \"Nearly Headless Nick\" is the House ghost.", false).
		AddField("Founder", "[Godrick Gryffindor](https://harrypotter.fandom.com/wiki/Godric_Gryffindor)", false).
		AddField("Traits", "Bravery\nNerve\nCourage\nChivalry\nDaring", false).
		AddField("Famous Wizards", "[Harry Potter](https://harrypotter.fandom.com/wiki/Harry_Potter)\n[Albus Dumbledore](https://harrypotter.fandom.com/wiki/Albus_Dumbledore)\n[The Weasleys](https://harrypotter.fandom.com/wiki/Weasley_family)\n[Hermione Granger](https://harrypotter.fandom.com/wiki/Hermione_Granger)", false).
		AddField("More Info", "https://harrypotter.fandom.com/wiki/Gryffindor", false).
		MessageEmbed

	_, err := sess.ChannelMessageSendEmbed(msg.ChannelID, embed)
	if err != nil {
		fmt.Print(err)
	}
}

// RavenclawCommand displays info important to the Ravenclaw house
func RavenclawCommand(sess *discordgo.Session, msg *discordgo.Message) {

	embed := embed.NewEmbed().SetTitle("Ravenclaw House").
		SetColor(924224).
		SetURL("https://www.pottermore.com/collection/all-about-ravenclaw").
		SetThumbnail("https://www.hp-lexicon.org/wp-content/uploads/2015/08/shield_rav-200x0-c-default.jpg").
		AddField("About", "Ravenclaw is one of the four Houses of Hogwarts School of Witchcraft and Wizardry. Its founder was a medieval witch Rowena Ravenclaw. Members of this house are characterised by their wit, learning, and wisdom. The emblematic animal is symbol is an eagle, and blue and bronze are its colours. The Head of Ravenclaw is Filius Flitwick and the house ghost is the Grey Lady, otherwise known as Helena Ravenclaw.", false).
		AddField("Founder", "[Rowena Ravenclaw](https://harrypotter.fandom.com/wiki/Rowena_Ravenclaw)", false).
		AddField("Traits", "Intelligence\nWit\nWisdom\nCreativity\nOriginality\nIndividuality\nAcceptance", false).
		AddField("Famous Wizards", "[Luna Lovegood](https://harrypotter.fandom.com/wiki/Luna_Lovegood)\n[Filius Flitwick](https://harrypotter.fandom.com/wiki/Filius_Flitwick)\n[Gilderoy Lockheart](https://harrypotter.fandom.com/wiki/Gilderoy_Lockhart)", false).
		AddField("More Info", "https://harrypotter.fandom.com/wiki/Ravenclaw", false).
		MessageEmbed

	_, err := sess.ChannelMessageSendEmbed(msg.ChannelID, embed)
	if err != nil {
		fmt.Print(err)
	}
}

func sendMessage(sess *discordgo.Session, channelid string, message string) error {
	_, err := sess.ChannelMessageSend(channelid, message)
	return err
}

func mentionUser(msg *discordgo.Message) string {
	return "<@" + msg.Author.ID + ">"
}
