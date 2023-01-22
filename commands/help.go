package commands

import (
	"ds-bot/tmp"
	"github.com/bwmarrin/discordgo"
)

const (
	greetMsg = `
		Here is some rules, you have to follow, to become a High Overload, Leader, or War chief:
			Do NOT:
				Harass, bully, or threat marginalized or vulnerable groups of people,
				Engage in content manipulation (spamming, subscriber fraud, vote manipulation, or ban evasion),
				Post or threaten to post intimate or sexually explicit photos or videos of another person without their consent,
				Impersonate someone in a misleading way,
		 		Post illegal content,
		 		Do anything that stops the normal use of this server,
			For The Horde!	
		`
	botPrefix       = "."
	ownerChatId     = "1058658554325770320"
	forTheHordeLink = "https://media.giphy.com/media/xThtatRttFzLD9oEtG/giphy.gif"
	message         = `
	Hey!
		My name is Saurfang, and I am extremely useful bot!
		I have such commands: 
			$search {keyword} 	  	- Send gif, which represents keyword
			$help 			      	- List of commands, which this bot wields
			$ban	{user} {reason} - Kick out user from server 
			$listrules 		  	    - Prints all rules, in this server
			$unban 					- Unbans user(need to reply to message)
	`
	rules = `
		Harass, bully, or threat marginalized or vulnerable groups of people,
		Engage in content manipulation (spamming, subscriber fraud, vote manipulation, or ban evasion),
		Post or threaten to post intimate or sexually explicit photos or videos of another person without their consent,
		Impersonate someone in a misleading way,
		Label the content and communities improperly (especially graphic content),
		Post suggestive or sexual content that involves minors,
		Post illegal content,
		Do anything that stops the normal use of this server,
		Transmit, distribute, or upload any viruses, worms, or other malware intended to interfere with this servers service,
		Use the platform to violate the law or infringe on intellectual and other property right,
		Engage in actions that could disrupt, disable, overburden, or impair this servers service,
		Attempt to gain access to another user's account,
		Access, search, or collect data from this server,
		Use the platform in any way that may be abusive or fraudulent`
)

func Help(s *discordgo.Session, ms *discordgo.MessageCreate) {
	if ms.Author.Bot {
		return
	}
	if ms.Content == botPrefix+"help" {
		s.ChannelMessageSendEmbed(ms.ChannelID, tmp.CreateEmbedInfoMessage("Help: ", message).Return())
	}
}
