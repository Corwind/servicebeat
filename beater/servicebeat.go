package beater

import (
	"fmt"
	"time"

        "os/exec"
        "strings"
        //"bytes"
        "regexp"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/Corwind/servicebeat/config"
)

// Servicebeat configuration.
type Servicebeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of servicebeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Servicebeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts servicebeat.
func (bt *Servicebeat) Run(b *beat.Beat) error {
	logp.Info("servicebeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

        var result bool
        var output []byte
        var services common.MapStr

        var regex *regexp.Regexp

        var command []string

        //var stdout bytes.Buffer
        //var stderr bytes.Buffer

        regex, err = regexp.Compile(bt.config.Regex)
        
        command = strings.Fields(bt.config.Command)
        var command_slice = make([]string, len(command) + 1, len(command) + 1)

        for index, _ := range command {
            command_slice[index] = command[index]
        }

        if err != nil {
            return err
        }

	ticker := time.NewTicker(bt.config.Period)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

                var fields = common.MapStr{
                        "type":    b.Info.Name,
                }

                services = common.MapStr {}
                for _, element := range bt.config.Services {
                    command_slice[len(command)] = element

                    output, err = exec.Command(command_slice[0], command_slice[1:]...).Output()

                    if err != nil {
                        return err
                    }

                    result = regex.Match(output)
                    var result_int = 0
                    if result {
                        result_int = 1
                    }
                    services.Put(element, result_int)
                }

                fields.Put("services", services)

		event := beat.Event{
			Timestamp: time.Now(),
                        Fields: fields,
		}
		bt.client.Publish(event)
		logp.Info("Event sent")
	}
}

// Stop stops servicebeat.
func (bt *Servicebeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
