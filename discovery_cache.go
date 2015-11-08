package myopenid

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yohcop/openid-go"
)

type MysqlDiscoveryCache struct {
	opEndpoint string
	opLocalID  string
	claimedID  string
}

func (m *MysqlDiscoveryCache) OpEndpoint() string {
	return m.opEndpoint
}

func (m *MysqlDiscoveryCache) OpLocalID() string {
	return m.opLocalID
}

func (m *MysqlDiscoveryCache) ClaimedID() string {
	return m.claimedID
}

func (m MysqlDiscoveryCache) Put(id string, info openid.DiscoveredInfo) {
	_, err := db.Exec("INSERT INTO discoverycache SET opendpoint=?, oplocalid=?, claimedid=?, discoverid=?", info.OpEndpoint(), info.OpLocalID(), info.ClaimedID(), id)
	if err != nil {
		log.Printf("\nError 1: %s", err.Error())
	}
}

func (m MysqlDiscoveryCache) Get(id string) openid.DiscoveredInfo {
	info := new(MysqlDiscoveryCache)
	err := db.QueryRow("SELECT opendpoint, oplocalid, claimedid FROM discoverycache	WHERE discoverid=?", id).Scan(&info.opEndpoint, &info.opLocalID, &info.claimedID)
	switch {
	case err == sql.ErrNoRows:
		return nil
	case err != nil:
		log.Printf("\nError 2: %s", err.Error())
	}
	return info
}
