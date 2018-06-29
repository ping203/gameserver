package internal

import (
	"errors"
	"fmt"
	"time"

	"server/gamelogic"
	"server/gamelogic/game-poke"

	"github.com/name5566/leaf/module"
	"github.com/sirupsen/logrus"
)

var errCanNotFindUser = errors.New("errCanNotFindUser")
var errCanNotFindPlayer = errors.New("errCanNotFindPlayer")
var errCanNotFindRoom = errors.New("errCanNotFindRoom")
var errUserAlreadyExist = errors.New("errUserAlreadyExist")
var errUserInvalid = errors.New("errUserInvalid")

// gameManager 桌子 user 管理器
type gameManager struct {
	idMkr uint32
	games map[uint32]gamelogic.Game //桌子
}

func newGameManager() *gameManager {
	return &gameManager{
		games: map[uint32]gamelogic.Game{},
	}
}

func (gm *gameManager) newID() uint32 {
	gm.idMkr++
	return gm.idMkr
}

func (gm *gameManager) newGame() uint32 {
	id := gm.newID()
	g := poke.NewGame(newGameService(), cfgMgr, id)

	gm.setGameByID(id, g)
	return id
}

func (gm *gameManager) startGameWithUsers(users ...gamelogic.User) error {
	id := gm.newGame()
	for _, v := range users {
		err := gm.userJoin(v, id)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"users": users,
				"err":   err,
			}).Error("startGameWithUsers")
			gm.deleteGame(id)
		}
	}
	return nil
}

func (gm *gameManager) deleteGame(gameID uint32) {
	delete(gm.games, gameID)
}

func (gm *gameManager) setGameByID(gameID uint32, game gamelogic.Game) {
	gm.games[gameID] = game
}

func (gm *gameManager) getGameByID(gameID uint32) (gamelogic.Game, bool) {
	g, exist := gm.games[gameID]
	if !exist {
		return nil, false
	}
	return g, exist
}

func (gm *gameManager) userJoin(user gamelogic.User, gameID uint32) error {
	g, exist := gm.getGameByID(gameID)
	if !exist {
		return errCanNotFindRoom
	}

	err := g.UserJoin(user)
	if err != nil {
		return err
	}

	return nil
}

func (gm *gameManager) userQuit(user gamelogic.User, gameID uint32) error {
	g, exist := gm.getGameByID(gameID)
	if !exist {
		return errCanNotFindRoom
	}

	err := g.UserQuit(user)
	if err != nil {
		return err
	}

	return nil
}

type gameService struct {
	*module.Skeleton
}

func (p *gameService) Post(f func()) {
	p.Skeleton.Post(f)
}
func (p *gameService) AfterPost(t time.Duration, f func()) func() {
	timer := p.Skeleton.AfterFunc(t, f)
	return timer.Stop
}

func (p *gameService) GameOver(gameID uint32) {
	_, exist := gameMgr.getGameByID(gameID)
	if !exist {
		logrus.Error("GameOver %v", gameID)
		return
	}

	delete(gameMgr.games, gameID)
	logrus.Debug(fmt.Sprintf("Game %v 游戏结束, 当前进行的游戏:%v", gameID, len(gameMgr.games)))
}

func newGameService() *gameService {
	return &gameService{
		Skeleton: skeleton,
	}
}
