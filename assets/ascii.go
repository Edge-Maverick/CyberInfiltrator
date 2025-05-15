package assets

import (
	"strings"
)

// ASCII art for the game

// HackSimLogo returns the ASCII art logo for HackSim
func HackSimLogo() string {
	return `
  _    _          _____ _  _______ _____ __  __ 
 | |  | |   /\   / ____| |/ / ____|_   _|  \/  |
 | |__| |  /  \ | |    | ' / (___   | | | \  / |
 |  __  | / /\ \| |    |  < \___ \  | | | |\/| |
 | |  | |/ ____ \ |____| . \____) |_| |_| |  | |
 |_|  |_/_/    \_\_____|_|\_\_____/|_____|_|  |_|
                                                
`
}

// TerminalHeader returns an ASCII art terminal header
func TerminalHeader() string {
	return `
┌──────────────────────────────────────────────────┐
│              TERMINAL ACCESS POINT               │
└──────────────────────────────────────────────────┘
`
}

// CompletedMission returns ASCII art for mission completion
func CompletedMission() string {
	return `
    ____  _  _  ____  ____  ____  ____  ____  / ) 
   (_  _)/ )( \(  __)/ ___)(  __)(  __)/ ___)(  ) 
     )(  ) __ ( ) _) \___ \ ) _)  ) _) \___ \ )(  
    (__) \_)(_/(____)(____/(____)(____)(____/(__) 
                                               
         Mission accomplished, hacker.
`
}

// FailedMission returns ASCII art for mission failure
func FailedMission() string {
	return `
    ____  ____  ____  __    __  __  ____  ____  / )
   (    \(  __)(_  _)(  )  (  )(  )(  _ \(  __)(  )
    ) D ( ) _)   )(   )(    )( /  \ )   / ) _)  )( 
   (____/(____) (__) (__)  (__)\_)(_(__\_)(____)(__)
                                                  
          Security systems detected you.
`
}

// ProgressBar generates an ASCII progress bar
func ProgressBar(percent float64, width int) string {
	filled := int(float64(width) * percent)
	
	if filled > width {
		filled = width
	}
	
	bar := "["
	bar += strings.Repeat("█", filled)
	bar += strings.Repeat("░", width-filled)
	bar += "]"
	
	return bar
}

// ComputerIcon returns a simple ASCII computer icon
func ComputerIcon() string {
	return `
  .----.
  |____| 
  |    | 
 .|____| 
 ||    | 
 ||____| 
.||____| 
'------'
`
}

// LockIcon returns an ASCII lock icon
func LockIcon() string {
	return `
   .-.
  (   )
 ,'-. |
('"""\|
'---._)
`
}

// HackerIcon returns an ASCII hacker icon
func HackerIcon() string {
	return `
     ,_     _,
     |\\___//|
     | >   < |
     | >_  <_|
     |      /
     \`'^'^'/
      '-...-'
`
}

// ServerIcon returns an ASCII server rack icon
func ServerIcon() string {
	return `
  .---.
 /___/|
 |   ||
 |___||
 |___||
 |___||
 |___||
 |___||
 '---'|
      |
      '
`
}

// DataIcon returns an ASCII data/file icon
func DataIcon() string {
	return `
    _____
   /|    |\
  | |    | |
  | |    | |
   \|____|/
     /  \
    / /\ \
    \ \/ /
     \  /
      \/
`
}
