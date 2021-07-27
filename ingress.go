/* Video ingress component.

   - Receives RTMP connections from users.
   - Validates stream keys.
   - Listens for connections from egress components.
*/
package main

import (
	"github.com/aler9/gortsplib/pkg/headers"
	"github.com/aler9/rtsp-simple-server/internal/hlsserver"
	"github.com/aler9/rtsp-simple-server/internal/logger"
	"github.com/aler9/rtsp-simple-server/internal/pathman"
	"github.com/aler9/rtsp-simple-server/internal/rtmpserver"
	"github.com/aler9/rtsp-simple-server/internal/stats"
	"github.com/aler9/rtsp-simple-server/internal/conf"
	"os"
	"sync/atomic"
	"time"
)

import (
	"context"
)

type program struct {
	ctx        context.Context
	ctxCancel  func()
	stats      *stats.Stats
	logger     *logger.Logger
	pathMan    *pathman.PathManager
	rtmpServer *rtmpserver.Server
	hlsServer  *hlsserver.Server

	// out
	done chan struct{}
}

func newProgram(args []string) (*program, bool) {
	ctx, ctxCancel := context.WithCancel(context.Background())

	p := &program{
		ctx:       ctx,
		ctxCancel: ctxCancel,
		done:      make(chan struct{}),
	}

	err := p.createResources()
	if err != nil {
		p.Log(logger.Info, "ERR: %s", err)
		p.closeResources()
		return nil, false
	}

	go p.run()

	return p, true
}

func (p *program) close() {
	p.ctxCancel()
	<-p.done
}

func (p *program) Log(level logger.Level, format string, args ...interface{}) {
	countPublishers := atomic.LoadInt64(p.stats.CountPublishers)
	countReaders := atomic.LoadInt64(p.stats.CountReaders)

	p.logger.Log(level, "[%d/%d] "+format, append([]interface{}{
		countPublishers, countReaders,
	}, args...)...)
}

func (p *program) run() {
	defer close(p.done)

outer:
	for {
		select {
		case <-p.ctx.Done():
			break outer
		}
	}

	p.ctxCancel()

	p.closeResources()
}

const (
	ConfReadTimeout =  10 * time.Second
	ConfWriteTimeout = 10 * time.Second
	ConfReadBufferCount = 512
	ConfReadBufferSize = 2048
	ConfRTSPSAddress = ":8555"
	ConfRTMPAddress = ":1935"
	ConfHLSAddress = ":8888"
	ConfHLSSegmentCount = 5
	ConfHLSSegmentDuration = 1 * time.Second
	ConfHLSAllowOrigin = "*"
)
var (
	ConfAuthMethodsParsed = []headers.AuthMethod{headers.AuthBasic, headers.AuthDigest}
	ConfPaths = map[string]*conf.PathConf{
		"~^.*$": {},
	}
	ConfLogDestinations = map[logger.Destination]struct{} {
		logger.DestinationStdout: {},
	}
)

func (p *program) createResources() error {
	var err error

	p.stats = stats.New()
	p.logger, err = logger.New(
		logger.Debug,
		ConfLogDestinations,
		"rtsp-simple-server.log")
	if err != nil {
		return err
	}

	for name, pconf := range ConfPaths {
		if pconf == nil {
			ConfPaths[name] = &conf.PathConf{}
			pconf = ConfPaths[name]
		}

		err := pconf.FillAndCheck(name)
		if err != nil {
			return err
		}
	}

	p.pathMan = pathman.New(
		p.ctx,
		ConfRTSPSAddress,
		ConfReadTimeout,
		ConfWriteTimeout,
		ConfReadBufferCount,
		ConfReadBufferSize,
		ConfAuthMethodsParsed,
		ConfPaths,
		p.stats,
		p) // this is worrying

	p.rtmpServer, err = rtmpserver.New(
		p.ctx,
		ConfRTMPAddress,
		ConfReadTimeout,
		ConfWriteTimeout,
		ConfReadBufferCount,
		ConfRTSPSAddress,
		"",
		false,
		p.stats,
		p.pathMan,
		p)
	if err != nil {
			return err
		}

	p.hlsServer, err = hlsserver.New(
		p.ctx,
		ConfHLSAddress,
		ConfHLSSegmentCount,
		ConfHLSSegmentDuration,
		ConfHLSAllowOrigin,
		ConfReadBufferCount,
		p.stats,
		p.pathMan,
		p)
	if err != nil {
			return err
		}

	return nil
}

func (p *program) closeResources() {
	if p.pathMan != nil {
		p.pathMan.Close()
		p.pathMan = nil
	}

	if p.hlsServer != nil {
		p.hlsServer.Close()
		p.hlsServer = nil
	}

	if p.rtmpServer != nil {
		p.rtmpServer.Close()
		p.rtmpServer = nil
	}

	if p.logger != nil {
		p.logger.Close()
		p.logger = nil
	}

	if p.stats != nil {
		p.stats.Close()
	}
}

func main() {
	p, ok := newProgram(os.Args[1:])
	if !ok {
		os.Exit(1)
	}
	<-p.done
}
