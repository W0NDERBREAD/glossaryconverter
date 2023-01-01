package converter

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	type args struct {
		glossaryPath string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]Entry
		wantErr bool
	}{
		{
			name: "Converts word with defenition",
			args: args{glossaryPath: filepath.Join("testdata", "standard.txt")},
			want: map[string]Entry{
				"Hundred Companions, the": {
					word:        "Hundred Companions, the",
					defenitions: []Defenition{{defenition: "One hundred male Aes Sedai, among the most powerful of the Age of Legends, who, led by Lews Therin Telamon, launched the final stroke that ended the War of the Shadow by sealing the Dark One back into his prison. The Dark One’s counterstroke tainted saidin; the Hundred Companions went mad and began the Breaking of the World."}},
				},
			},
			wantErr: false,
		},
		{
			name: "Converts word with pronunciation",
			args: args{glossaryPath: filepath.Join("testdata", "pronunciation.txt")},
			want: map[string]Entry{
				"Ajah": {
					word:          "Ajah",
					pronunciation: "AH-jah",
					defenitions:   []Defenition{{defenition: "Societies among the Aes Sedai, to which all Aes Sedai belong. They are designated by colors: Blue Ajah, Red Ajah, White Ajah, Green Ajah, Brown Ajah, Yellow Ajah, and Gray Ajah. Each follows a specific philosophy of the use of the One Power and purposes of the Aes Sedai. For example, the Red Ajah bends all its energies to finding and gentling men who are attempting to wield the Power. The Brown Ajah, on the other hand, forsakes involvement with the world and dedicates itself to seeking knowledge. There are rumors (hotly denied, and never safely mentioned in front of any Aes Sedai) of a Black Ajah, dedicated to serving the Dark One."}},
				},
			},
			wantErr: false,
		},
		{
			name: "Converts defenition spanning multiple lines",
			args: args{glossaryPath: filepath.Join("testdata", "multiline.txt")},
			want: map[string]Entry{
				"currency": {
					word: "currency",
					defenitions: []Defenition{{defenition: `After many centuries of trade, the standard terms for coins are the same in every land: crowns (the largest coin in size), marks and pennies. Crowns and marks can be minted of gold or silver, while pennies can be silver or copper, the last often called simply a copper. In different lands, however, these coins are of different sizes and weights. Even in one nation, coins of different sizes and weights have been minted by different rulers. Because of trade, the coins of many nations can be found almost anywhere, and for that reason, bankers, moneylenders and merchants all use scales to determine the value of any given coin. Even large numbers of coins are weighed.

The heaviest coins come from Andor and Tar Valon, and in those two places the relative values are: 10 copper pennies = 1 silver penny; 100 silver pennies = 1 silver mark; 10 silver marks = 1 silver crown; 10 silver crowns = 1 gold mark; 10 gold marks = 1 gold crown. By contrast, in Altara, where the larger coins contain less gold or silver, the relative values are: 10 copper pennies = 1 silver penny; 21 silver pennies = 1 silver mark; 20 silver marks = 1 silver crown; 20 silver crowns = 1 gold mark; 20 gold marks = 1 gold crown.`}},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Convert(tt.args.glossaryPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getExpectedEntries() map[string]Entry {
	entries := make(map[string]Entry)

	entries["Ajah"] = Entry{
		word:          "Ajah",
		pronunciation: "AH-jah",
		defenitions:   []Defenition{{defenition: "Societies among the Aes Sedai, to which all Aes Sedai belong. They are designated by colors: Blue Ajah, Red Ajah, White Ajah, Green Ajah, Brown Ajah, Yellow Ajah, and Gray Ajah. Each follows a specific philosophy of the use of the One Power and purposes of the Aes Sedai. For example, the Red Ajah bends all its energies to finding and gentling men who are attempting to wield the Power. The Brown Ajah, on the other hand, forsakes involvement with the world and dedicates itself to seeking knowledge. There are rumors (hotly denied, and never safely mentioned in front of any Aes Sedai) of a Black Ajah, dedicated to serving the Dark One."}},
	}

	entries["Amyrlin Seat"] = Entry{
		word:          "Amyrlin Seat",
		pronunciation: "AHM-ehr-lin",
		defenitions: []Defenition{{defenition: `(1) The title of the leader of the Aes Sedai.

Elected for life by the Hall of the Tower, the highest council of the Aes Sedai, which consists of three representatives from each of the seven Ajahs. The Amyrlin Seat has, theoretically at least, almost supreme authority among the Aes Sedai. She ranks as the equal of a king or queen. (2) The throne upon which the leader of the Aes Sedai sits.`}},
	}

	entries["Andor"] = Entry{
		word:          "Andor",
		pronunciation: "AN-door",
		defenitions:   []Defenition{{defenition: "The realm within which the Two Rivers lies. The sign of Andor is a rampant white lion on a field of red."}},
	}

	entries["Breaking of the World, the"] = Entry{
		word:          "Breaking of the World, the",
		pronunciation: "",
		defenitions:   []Defenition{{defenition: "When Lews Therin Telamon and the Hundred Companions resealed the Dark One’s prison, the counterstroke tainted saidin. Eventually every male Aes Sedai went horribly insane. In their madness these men, who could wield the One Power to a degree now unknown, changed the face of the earth. They caused great earthquakes, leveled mountain ranges, raised new mountains, lifted dry land where seas had been, made the ocean rush in where dry land had been. Many parts of the world were completely depopulated, and the survivors were scattered like dust on the wind. This destruction is remembered in stories, legends and history as the Breaking of the World. See also Hundred Companions, the."}},
	}

	entries["channel"] = Entry{
		word:          "channel",
		pronunciation: "",
		defenitions:   []Defenition{{defenition: "(1) (verb) To control the flow of the One Power. (2) (noun) The act of controlling the flow of the One Power."}},
	}

	entries["currency"] = Entry{
		word:          "currency",
		pronunciation: "",
		defenitions: []Defenition{{defenition: `After many centuries of trade, the standard terms for coins are the same in every land: crowns (the largest coin in size), marks and pennies. Crowns and marks can be minted of gold or silver, while pennies can be silver or copper, the last often called simply a copper. In different lands, however, these coins are of different sizes and weights. Even in one nation, coins of different sizes and weights have been minted by different rulers. Because of trade, the coins of many nations can be found almost anywhere, and for that reason, bankers, moneylenders and merchants all use scales to determine the value of any given coin. Even large numbers of coins are weighed.

The heaviest coins come from Andor and Tar Valon, and in those two places the relative values are: 10 copper pennies = 1 silver penny; 100 silver pennies = 1 silver mark; 10 silver marks = 1 silver crown; 10 silver crowns = 1 gold mark; 10 gold marks = 1 gold crown. By contrast, in Altara, where the larger coins contain less gold or silver, the relative values are: 10 copper pennies = 1 silver penny; 21 silver pennies = 1 silver mark; 20 silver marks = 1 silver crown; 20 silver crowns = 1 gold mark; 20 gold marks = 1 gold crown.`}},
	}

	entries["Hundred Companions, the"] = Entry{
		word:          "Hundred Companions, the",
		pronunciation: "",
		defenitions:   []Defenition{{defenition: "One hundred male Aes Sedai, among the most powerful of the Age of Legends, who, led by Lews Therin Telamon, launched the final stroke that ended the War of the Shadow by sealing the Dark One back into his prison. The Dark One’s counterstroke tainted saidin; the Hundred Companions went mad and began the Breaking of the World."}},
	}

	entries["Illian"] = Entry{
		word:          "Illian",
		pronunciation: "IHL-lee-ahn",
		defenitions:   []Defenition{{defenition: "A great port on the Sea of Storms, capital city of the nation of the same name. The sign of Illian is nine golden bees on a field of dark green."}},
	}

	entries["One Power, the"] = Entry{
		word:          "One Power, the",
		pronunciation: "",
		defenitions:   []Defenition{{defenition: "The power drawn from the True Source. The vast majority of people are completely unable to learn to channel the One Power. A very small number can be taught to channel, and an even tinier number have the ability inborn. For these few there is no need to be taught; they will touch the True Source and channel the Power whether they want to or not, perhaps without even realizing what they are doing. This inborn ability usually manifests itself in late adolescence or early adulthood. If control is not taught, or self-learned (extremely difficult, with a success rate of only one in four), death is certain. Since the time of Madness, no man has been able to channel the Power without eventually going completely, horribly mad; and then, even if he has learned some control, dying from a wasting sickness which causes the sufferer to rot alive—a sickness caused, as is the madness, by the Dark One’s taint on saidin. For a woman the death that comes without control of the Power is less horrible, but it is death just the same. Aes Sedai search for girls with the inborn ability as much to save their lives as to increase Aes Sedai numbers, and for men with it in order to stop the terrible things they inevitably do with the Power in their madness. See also channel; Time of Madness; True Source."}},
	}

	entries["saidar; saidin"] = Entry{
		word:          "saidar; saidin",
		pronunciation: "sah-ih-DAHR; sah-ih-DEEN",
		defenitions:   []Defenition{{defenition: "See True Source."}},
	}

	entries["Time of Madness"] = Entry{
		word:          "Time of Madness",
		pronunciation: "",
		defenitions:   []Defenition{{defenition: "See Breaking of the World, the."}},
	}

	entries["True Source"] = Entry{
		word:          "True Source",
		pronunciation: "",
		defenitions:   []Defenition{{defenition: "The driving force of the universe, which turns the Wheel of Time. It is divided into a male half (saidin) and a female half (saidar), which work at the same time with and against each other. Only a man can draw on saidin, only a woman on saidar. Since the beginning of the Time of Madness, saidin has been tainted by the Dark One’s touch. See also One Power."}},
	}

	return entries
}
