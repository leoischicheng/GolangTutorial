package greetings
import(
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init(){
	rand.Seed(time.Now().UnixNano());
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name, return an error with an message 
    if name == "" {
    	return "", errors.New("Empty Name")
    }

    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	//A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
    	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
