package infra

import (
	"crypto/tls"
	"net"

	"gopkg.in/mgo.v2"
)

type Session struct {
	session *mgo.Session
}

func NewSession(url string) (*Session, error) {
	dialInfo, err := mgo.ParseURL(url)
	tlsConfig := &tls.Config{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, err
	}
	return &Session{session}, err
}

func (s *Session) Copy() *Session {
	return &Session{s.session.Copy()}
}

func (s *Session) GetCollection(db string, col string) *mgo.Collection {
	return s.session.DB(db).C(col)
}

func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}
