package commands

import (
	"GoSortingBot/embed"
	"fmt"
	"strings"

	dg "github.com/bwmarrin/discordgo"
)

var commandsList = []string{"!houses", "!gryffindor", "!ravenclaw", "!hufflepuff", "!slytherin"}

// HelpCommand is the command to display help message
func HelpCommand(sess *dg.Session, msg *dg.Message) {
	messageStr := mentionUser(msg) + ", hi I'm the Sorting Hat!"
	sendMessage(sess, msg.ChannelID, messageStr)
	messageStr = "My commands are: " + strings.Join(commandsList, ", ")
	sendMessage(sess, msg.ChannelID, messageStr)
}

// HouseCommand Displays info about the houses
func HouseCommand(sess *dg.Session, msg *dg.Message) {
	houseStr := "There are 4 different houses: Gryffindor, Hufflepuff, Ravenclaw and Slytherin"
	sendMessage(sess, msg.ChannelID, houseStr)
	houseStr = "To learn more about a particular house, use !slytherin, !ravenclaw, !hufflepuff or !gryffindor"
	sendMessage(sess, msg.ChannelID, houseStr)
}

// GryffindorCommand displays info important to the Gryffindor house
func GryffindorCommand(sess *dg.Session, msg *dg.Message) {

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
func RavenclawCommand(sess *dg.Session, msg *dg.Message) {

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

// HufflepuffCommand displays info important to the Hufflepuff house
func HufflepuffCommand(sess *dg.Session, msg *dg.Message) {

	embed := embed.NewEmbed().SetTitle("Hufflepuff House").
		SetColor(15513913).
		SetURL("https://www.pottermore.com/collection/all-about-hufflepuff").
		SetThumbnail("https://www.hp-lexicon.org/wp-content/uploads/2015/08/hufflepuff-shield-200x0-c-default.jpg").
		AddField("About", "Hufflepuff is one of the four Houses of Hogwarts School of Witchcraft and Wizardry. Its founder was the medieval witch Helga Hufflepuff. Hufflepuff is the most inclusive among the four houses; valuing hard work, dedication, patience, loyalty, and fair play rather than a particular aptitude in its members. The emblematic animal is a badger, and yellow and black are its colours. The Head of Hufflepuff is Pomona Sprout and the Fat Friar is the House's patron ghost.", false).
		AddField("Founder", "[Helga Hufflepuff](https://harrypotter.fandom.com/wiki/Helga_Hufflepuff)", false).
		AddField("Traits", "Dedication\nHardworking\nFairness\nPatience\nKindness\nTolerance\nLoyalty", false).
		AddField("Famous Wizards", "[Cedric Diggory](https://harrypotter.fandom.com/wiki/Cedric_Diggory)\n[Pomona Sprout](https://harrypotter.fandom.com/wiki/Pomona_Sprout)", false).
		AddField("More Info", "https://harrypotter.fandom.com/wiki/Hufflepuff", false).
		MessageEmbed

	_, err := sess.ChannelMessageSendEmbed(msg.ChannelID, embed)
	if err != nil {
		fmt.Print(err)
	}
}

// SlytherinCommand displays info important to the Hufflepuff house
func SlytherinCommand(sess *dg.Session, msg *dg.Message) {

	embed := embed.NewEmbed().SetTitle("Slytherin House").
		SetColor(1722154).
		SetURL("https://www.pottermore.com/collection/all-about-slytherin").
		SetThumbnail("https://cdn.shopify.com/s/files/1/0006/8213/1492/products/Slytherin_Hogwarts_House_Crest_-_Harry_Potter_1024x1024.jpg?v=1547596686").
		AddField("About", "Slytherin is one of the four Houses at Hogwarts School of Witchcraft and Wizardry, founded by Salazar Slytherin. In establishing the house, Salazar instructed the Sorting Hat to pick students who had a few particular characteristics he most valued. Those characteristics include: cunning, resourcefulness, and ambition. Many Slytherin students tend to clique together, often acquiring leaders, which further exemplifies Slytherin's ambitious qualities. Examples of these include Draco Malfoy’s Gang and the Death Eaters.", false).
		AddField("Founder", "[Salazar Slytherin](https://harrypotter.fandom.com/wiki/Salazar_Slytherin)", false).
		AddField("Traits", "Resourcefulness\nCunning\nAmbition\nDetermination\nSelf-Preservation\nFraternity\nCleverness", false).
		AddField("Famous Wizards", "[Serverus Snape](https://harrypotter.fandom.com/wiki/Severus_Snape)\n[Merlin](https://harrypotter.fandom.com/wiki/Merlin)\n[Tom Marvolo Riddle](https://harrypotter.fandom.com/wiki/Tom_Riddle)", false).
		AddField("More Info", "https://harrypotter.fandom.com/wiki/Slytherin", false).
		MessageEmbed

	_, err := sess.ChannelMessageSendEmbed(msg.ChannelID, embed)
	if err != nil {
		fmt.Print(err)
	}
}

func sendMessage(sess *dg.Session, channelid string, message string) error {
	_, err := sess.ChannelMessageSend(channelid, message)
	return err
}

func mentionUser(msg *dg.Message) string {
	return "<@" + msg.Author.ID + ">"
}
