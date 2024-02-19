package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"null.hexolan.dev/dev/pkg/database"
	"null.hexolan.dev/dev/pkg/messagerelay"
)

// Kubernetes for relay leader election
// https://kubernetes.io/docs/concepts/architecture/leases/
// https://github.com/kubernetes/client-go/blob/master/examples/leader-election/main.go
//
// https://carlosbecker.com/posts/k8s-leader-election/

// also keep in mind permissions of service account:
// "leaderelection.go:332] error retrieving resource lock default/testsvc-outbox-relay: leases.coordination.k8s.io "testsvc-outbox-relay" is forbidden: User "system:serviceaccount:default:default" cannot get resource "leases" in API group "coordination.k8s.io" in the namespace "default"
//
// https://stackoverflow.com/questions/47973570/kubernetes-log-user-systemserviceaccountdefaultdefault-cannot-get-services

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// following code is a useful example 
	// for graceful cancel of the contexts 
	//
	// SOURCE: https://github.com/kubernetes/client-go/blob/master/examples/leader-election/main.go (Apache-2.0 - 2018: The Kubernetes Authors)
	// listen for interrupts or the Linux SIGTERM signal and cancel
	// our context, which the leader election code will observe and
	// step down
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		log.Info().Msg("received termination. signalling shutdown")
		cancel()
	}()
	// END SOURCE : https://github.com/kubernetes/client-go/blob/master/examples/leader-election/main.go (Apache-2.0 - 2018: The Kubernetes Authors)

	// check if in a kubernetes cluster (for if leader election is required)
	if cfg, err := rest.InClusterConfig(); err != nil {
		log.Info().Err(err).Bool("in-cluster", false).Msg("running standalone")
		start(ctx)
	} else {
		log.Info().Bool("in-cluster", true).Msg("running clustered (with leader election)")
		startWithElectioner(ctx, cfg)
	}
}

func startWithElectioner(ctx context.Context, cfg *rest.Config) {
	// todo: better unique relay identifiers??
	// uuid + timestamp
	// or used exposed container environment variable (pod id from kubernetes?)
	relayIdentifier := uuid.New().String() 

	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Panic().Err(err).Msg("failed to create kubernetes clientset (for leader election)")
	}

	// Lease lock object
	lock := &resourcelock.LeaseLock{
		LeaseMeta: metav1.ObjectMeta{
			Name: "testsvc-outbox-relay", // the lease lock resource name
			Namespace: "default", // the lease lock resource namespace
		},
		Client: clientset.CoordinationV1(),
		LockConfig: resourcelock.ResourceLockConfig{
			Identity: relayIdentifier,
		},
	}

	// Start loop
	leaderelection.RunOrDie(
		ctx, 
		leaderelection.LeaderElectionConfig{
			Lock: lock,
			ReleaseOnCancel: true,
			LeaseDuration: 20 * time.Second,
			RenewDeadline: 15 * time.Second,
			RetryPeriod: 3 * time.Second,
			Callbacks: leaderelection.LeaderCallbacks{
				OnStartedLeading: func(ctx context.Context) {
					log.Info().Msg("started leading")
					start(ctx)
				},
				OnStoppedLeading: func() {
					log.Info().Msg("leader lost // stopped leading")
					os.Exit(0)
				},
				OnNewLeader: func(identity string) {
					if identity == relayIdentifier {
						// this instance is leading now
						return
					}

					log.Info().Str("leader-identity", identity).Msg("new leader has been elected")
				},
			},
		},
	)
}

func start(ctx context.Context) {
	log.Info().Msg("relay has started")
	// log.Info().Msg("TODO: no relay method has been implemented yet - try homegrown log tailing / db polling (currently trying Debezium)")

	// todo: ensuring stopped when context is cancelled
	pCl, err := database.NewPostgresConn("postgresql://postgres:postgres@test-service-postgres:5432/postgres?sslmode=disable")
	if err != nil {
		log.Panic().Err(err).Msg("")
	}
	messagerelay.PostgresLogTailing(ctx, pCl)

	/*
	for {
		select {
			case <-ctx.Done():
				return
			default:
		}
		// do something. forever. until context is cancelled
	}
	*/
}

// with Debezium:
//
// testing without Zookeeper for Kafka
// testing with using NATS instead of Kafka
// https://www.iamninad.com/posts/docker-compose-for-your-next-debezium-and-postgres-project/