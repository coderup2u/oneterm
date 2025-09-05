package sshsrv

import (
	"context"
	"fmt"
	"time"

	"github.com/gliderlabs/ssh"
	gossh "golang.org/x/crypto/ssh"

	"github.com/veops/oneterm/internal/acl"
	"github.com/veops/oneterm/pkg/config"
	"github.com/veops/oneterm/pkg/utils"
)

var (
	ctx, cancel = context.WithCancel(context.Background())
	server      *ssh.Server
)

func init() {
	server = &ssh.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Cfg.Ssh.Host, config.Cfg.Ssh.Port),
		Handler: handler,
		PasswordHandler: func(ctx ssh.Context, password string) bool {
			sess, err := acl.LoginByPassword(ctx, ctx.User(), password, utils.IpFromNetAddr(ctx.RemoteAddr()))
			ctx.SetValue("session", sess)
			return err == nil
		},
		PublicKeyHandler: func(ctx ssh.Context, key ssh.PublicKey) bool {
			sess, err := acl.LoginByPublicKey(ctx, ctx.User(), string(gossh.MarshalAuthorizedKey(key)), utils.IpFromNetAddr(ctx.RemoteAddr()))
			ctx.SetValue("session", sess)
			return err == nil
		},
		HostSigners: []ssh.Signer{},
	}
}

func RunSsh() error {
	// Initialize host signer after database is ready
	server.HostSigners = []ssh.Signer{signer()}
	return server.ListenAndServe()
}

func StopSsh() {
	if server != nil {
		// Use a fresh context for shutdown with timeout
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()
		
		if err := server.Shutdown(shutdownCtx); err != nil {
			// If graceful shutdown fails, force close
			server.Close()
		}
	}
	cancel()
}
