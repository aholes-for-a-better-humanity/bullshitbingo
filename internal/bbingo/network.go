package bbingo

/*
The net/nats interactions are as follow :

- upon connection :
	- check if the stream "currentGame" exists, if not, create it.
	- fetch the latest message in "currentGame"
		- if the stream has a game start message, there is a game allowing to join, then join.
		- else, post a message to announce the start of this game
- once the name game is known :
	- subscribe to <gamename>.>
		- whohaswords : announces of
		-
*/

import (
	"strings"
	"time"

	"github.com/aholes-for-a-better-humanity/bullshitbingo/internal/misc"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

var (
	defaultNatsEndPoint       = "nats://127.0.0.1:4222" // demo.nats.io:4222 works just fine
	gameOpeningPeriodDuration = 15 * time.Minute        // duration to accept players into a same game
	gamePlayingPeriodDuration = 180 * time.Minute       // (expected) duration of a game after first player opened it
)

const (
	currentGameStreamName = "currentGame" // Announce of games that are opening
	gamesDataStreamName   = "gamesData"   // data of games while they are running
	pressedWordsTopic     = "pressed"
	hasWordsTopic         = "have"
	playerLeavingTopic    = "leaving"
)

// network provides communication capabilities for *Game
func (g *Game) network() error {
	var nc *nats.Conn
	var err error

	for {
		nc, err = nats.Connect(defaultNatsEndPoint)
		if err == nil {
			defer nc.Close()
			break
		}
		select {
		case <-g.ctx.Done():
			return g.ctx.Err()
		default:
			log.Error().Err(err).Msg("nats.Connect")
		}
	}

	jsc, err := nc.JetStream()
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	err = g.ensureGameStreamExists(jsc)
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	err = g.beInAGame(jsc) // g.gameWeAreIn is now reliable.
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	defer func() {
		// notify all players when leaving game (and adjust counts, quorum, etc.)
		// relies on the game life cycle not being shorted by Ebiten (works on desktop)
		sub := strings.Join([]string{"game", g.gameWeAreIn, playerLeavingTopic}, ".")
		_, _ = jsc.Publish(sub, []byte(g.nickname+"|")) // we could not react to any error, the context being Done.
	}()
	return g.networkMainLoop(jsc)

}

func (*Game) ensureGameStreamExists(jsc nats.JetStreamManager) error {
	_, err := jsc.AddStream(&nats.StreamConfig{
		Name:              currentGameStreamName,
		Description:       "Announce of games that are opening",
		Subjects:          []string{currentGameStreamName},
		Retention:         nats.LimitsPolicy,
		MaxConsumers:      0,
		MaxMsgs:           0,
		MaxBytes:          0,
		Discard:           nats.DiscardOld,
		MaxAge:            gameOpeningPeriodDuration,
		MaxMsgsPerSubject: 1,
		MaxMsgSize:        0,
		Storage:           nats.MemoryStorage,
		Duplicates:        2 * time.Minute,
	})
	if err != nil {
		log.Error().Err(err).Caller().Send()
	}
	_, err = jsc.AddStream(&nats.StreamConfig{
		Name:              gamesDataStreamName,
		Description:       "All things related to the games",
		Subjects:          []string{"game", "game.>"},
		Retention:         nats.LimitsPolicy,
		MaxMsgs:           10000,
		Discard:           nats.DiscardOld,
		MaxAge:            gamePlayingPeriodDuration,
		MaxMsgsPerSubject: 1000,
		MaxMsgSize:        0,
		Storage:           nats.MemoryStorage,
		Duplicates:        2 * time.Minute,
	})
	return err
}

// either join the current online game, or create a new one
func (g *Game) beInAGame(jsc nats.JetStreamContext) error {
	si, err := jsc.StreamInfo(currentGameStreamName)
	if err != nil {
		return err
	}
	if si.State.Msgs != 0 {
		rm, err := jsc.GetMsg(currentGameStreamName, si.State.LastSeq)
		if err != nil {
			return err
		}
		gameName := strings.ReplaceAll(string(rm.Data), " ", "-")
		log.Info().Msg(gameName)
		g.gameWeAreIn = gameName
		return nil
	}
	g.gameWeAreIn = strings.ReplaceAll(g.nickname, " ", "-")
	_, err = jsc.Publish(currentGameStreamName, []byte(g.gameWeAreIn))
	if err != nil {
		return err
	}
	return err
}

// networkMainLoop is the long-lived goroutine that subscribes and publishes all along the life of the program
func (g *Game) networkMainLoop(jsc nats.JetStreamContext) error {
	ch := make(chan *nats.Msg, 256)
	unsub := func(s *nats.Subscription) { _ = s.Unsubscribe() }
	if s, err := jsc.ChanSubscribe(strings.Join([]string{"game", g.gameWeAreIn}, "."), ch); err == nil {
		defer unsub(s)
	} else {
		log.Error().Err(err).Send()
		return err
	}
	if s, err := jsc.ChanSubscribe(strings.Join([]string{"game", g.gameWeAreIn, ">"}, "."), ch); err == nil {
		defer unsub(s)
	} else {
		log.Error().Err(err).Send()
		return err
	}
	g.eg.Go(func() error {
		for w := range g.ourWords {
			select {
			case g.toNetwork <- netMsg{topic: hasWordsTopic, content: w}:
			case <-g.ctx.Done():
				return g.ctx.Err()
			}
		}
		log.Info().Msg("sent all our words")
		return nil
	})
	for {
		select {
		case <-g.ctx.Done():
			return g.ctx.Err()
		case nm := <-g.toNetwork:
			sub := strings.Join([]string{"game", g.gameWeAreIn, nm.topic}, ".")
			if _, err := jsc.Publish(sub, []byte(g.nickname+"|"+nm.content)); err != nil {
				return err
			}
		case msg := <-ch:
			if err := msg.Ack(); err != nil {
				return err
			}
			subjTokens := strings.Split(msg.Subject, ".")
			topic := subjTokens[len(subjTokens)-1]
			datacontent := strings.Split(string(msg.Data), "|")
			sender, content := datacontent[0], datacontent[1]
			g.networkProcess(topic, sender, content)
		}
	}
}

func (g *Game) networkProcess(topic, sender, content string) {
	log := log.With().Str("topic", topic).Str("sender", sender).Str("content", content).Logger()
	log.Debug().Msg("recieved JS msg")
	if sender == g.nickname {
		log.Info().Msg("...")
		return
	}
	switch topic {
	case hasWordsTopic:
		if validState, ok := g.ourWords[content]; ok {
			validState.total++
			if g.whoHas[content] != nil {
				for _, name := range g.whoHas[content] {
					if name == sender {
						return
					}
				}
			}
			g.whoHas[content] = append(g.whoHas[content], sender)
			log.Debug().Strs("have_it", g.whoHas[content]).Int("number", validState.total).Send()
		} else {
			log.Error().Msg("donthavethat")
		}
	case pressedWordsTopic:
		if _, ok := g.ourWords[content]; ok {
			g.events <- gameEvent{sig: gameWordPressedByOther, word: content, sender: sender}
		}
	case playerLeavingTopic:
		log.Debug().Msg("player is leaving")
		for ourRealWord, vl := range g.ourWords {
			if misc.Has(g.whoPressed[ourRealWord], sender) {
				g.whoPressed[ourRealWord] = misc.RemoveFrom(g.whoPressed[ourRealWord], sender)
				g.ourWords[ourRealWord] = &validationLevel{
					Validated:     false,
					Self:          vl.Self,
					total:         vl.total - 1,
					OthersPressed: vl.OthersPressed - 1,
				}
				log.Debug().Msg("REVALIDATING")
				g.maybeValidate(ourRealWord)
				continue
			}
		}
	default:
		log.Error().Msg("missed")
	}
}
