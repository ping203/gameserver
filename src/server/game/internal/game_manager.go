package internal

import (
	"errors"
	"time"

	"server/gamelogic"
	"server/gamelogic/game-poke"

	"github.com/name5566/leaf/module"
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
	users map[uint64]*user
}

func newGameManager() *gameManager {
	return &gameManager{
		games: map[uint32]gamelogic.Game{},
		users: map[uint64]*user{},
	}
}

func (gm *gameManager) newID() uint32 {
	gm.idMkr++
	return gm.idMkr
}

func (gm *gameManager) newGame() {
	id := gm.newID()
	g := poke.NewGame(newGameService(), cfgMgr, id)

	gm.setGameByID(id, g)
}

func (gm *gameManager) usersJoin(gameID uint32) {
	delete(gm.games, gameID)
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

func (gm *gameManager) findUser(userid uint64) (*user, bool) {
	//int a
	user, exist := gm.users[userid]
	return user, exist
}

func (gm *gameManager) removeUser(user gamelogic.User) {
	delete(gm.users, user.ID())
}

func (gm *gameManager) userJoin(user *user, gameID uint32) error {
	g, exist := gm.getGameByID(gameID)
	if !exist {
		return errCanNotFindRoom
	}

	err := g.UserJoin(user)
	if err != nil {
		return err
	}

	user.setGameID(gameID)
	return nil
}

func (gm *gameManager) userQuit(user *user, gameID uint32) error {
	g, exist := gm.getGameByID(gameID)
	if !exist {
		return errCanNotFindRoom
	}

	err := g.UserQuit(user)
	if err != nil {
		return err
	}

	user.clearGameID()

	if g.IsEmpty() {
		gm.deleteGame(gameID)
	}

	return nil
}

type gameService struct {
	*module.Skeleton
}

func (p *gameService) Post(f func()) {
	p.Skeleton.Post(f)
}
func (p *gameService) AfterPost(t time.Duration, f func()) {
	p.Skeleton.AfterFunc(t, f)
}

func newGameService() *gameService {
	return &gameService{
		Skeleton: skeleton,
	}
}
