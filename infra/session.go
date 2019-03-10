package infra

import (
	"crypto/tls"
	"errors"
	"log"
	"net"

	"gopkg.in/mgo.v2"
)

// Session mongodb session
type Session struct {
	session *mgo.Session
	DbName  string
}

// NewSession create a mongodb session
func NewSession(config *MongoConfiguration) (*Session, error) {
	if config.Host == "" || config.DbName == "" {
		log.Fatal("db connections string or dbname is empty")
		return nil, errors.New("db connections string or dbname is empty")
	}

	dialInfo, err := mgo.ParseURL(config.Host)
	tlsConfig := &tls.Config{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.Dial(config.Host)
	if err != nil {
		return nil, err
	}
	return &Session{session, config.DbName}, err
}

// Copy create a copy for the session
func (s *Session) Copy() *Session {
	return &Session{s.session.Copy(), s.DbName}
}

// GetCollection get a mongodb collections
func (s *Session) GetCollection(db string, col string) *mgo.Collection {
	return s.session.DB(db).C(col)
}

// Close close session
func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}

// DropDatabase drop mongodb database *TEST PRUPORSE*
func (s *Session) DropDatabase(db string) error {
	if s.session != nil {
		return s.session.DB(db).DropDatabase()
	}
	return nil
}
