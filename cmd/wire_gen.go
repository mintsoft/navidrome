// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire gen -tags "netgo"
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"context"
	"github.com/google/wire"
	"github.com/navidrome/navidrome/core"
	"github.com/navidrome/navidrome/core/agents"
	"github.com/navidrome/navidrome/core/agents/lastfm"
	"github.com/navidrome/navidrome/core/agents/listenbrainz"
	"github.com/navidrome/navidrome/core/artwork"
	"github.com/navidrome/navidrome/core/external"
	"github.com/navidrome/navidrome/core/ffmpeg"
	"github.com/navidrome/navidrome/core/metrics"
	"github.com/navidrome/navidrome/core/playback"
	"github.com/navidrome/navidrome/core/scrobbler"
	"github.com/navidrome/navidrome/db"
	"github.com/navidrome/navidrome/dlna"
	"github.com/navidrome/navidrome/model"
	"github.com/navidrome/navidrome/persistence"
	"github.com/navidrome/navidrome/scanner"
	"github.com/navidrome/navidrome/server"
	"github.com/navidrome/navidrome/server/events"
	"github.com/navidrome/navidrome/server/nativeapi"
	"github.com/navidrome/navidrome/server/public"
	"github.com/navidrome/navidrome/server/subsonic"
)

import (
	_ "github.com/navidrome/navidrome/adapters/taglib"
)

// Injectors from wire_injectors.go:

func CreateDataStore() model.DataStore {
	sqlDB := db.Db()
	dataStore := persistence.New(sqlDB)
	return dataStore
}

func CreateServer() *server.Server {
	sqlDB := db.Db()
	dataStore := persistence.New(sqlDB)
	broker := events.GetBroker()
	insights := metrics.GetInstance(dataStore)
	serverServer := server.New(dataStore, broker, insights)
	return serverServer
}

func CreateDLNAServer() *dlna.DLNAServer {
	sqlDB := db.Db()
	dataStore := persistence.New(sqlDB)
	broker := events.GetBroker()
	fFmpeg := ffmpeg.New()
	transcodingCache := core.GetTranscodingCache()
	mediaStreamer := core.NewMediaStreamer(dataStore, fFmpeg, transcodingCache)
	fileCache := artwork.GetImageCache()
	agentsAgents := agents.GetAgents(dataStore)
	externalMetadata := core.NewExternalMetadata(dataStore, agentsAgents)
	artworkArtwork := artwork.NewArtwork(dataStore, fileCache, fFmpeg, externalMetadata)
	dlnaServer := dlna.New(dataStore, broker, mediaStreamer, artworkArtwork)
	return dlnaServer
}

func CreateNativeAPIRouter() *nativeapi.Router {
	sqlDB := db.Db()
	dataStore := persistence.New(sqlDB)
	share := core.NewShare(dataStore)
	playlists := core.NewPlaylists(dataStore)
	insights := metrics.GetInstance(dataStore)
	router := nativeapi.New(dataStore, share, playlists, insights)
	return router
}

func CreateSubsonicAPIRouter(ctx context.Context) *subsonic.Router {
	sqlDB := db.Db()
	dataStore := persistence.New(sqlDB)
	fileCache := artwork.GetImageCache()
	fFmpeg := ffmpeg.New()
	agentsAgents := agents.GetAgents(dataStore)
	provider := external.NewProvider(dataStore, agentsAgents)
	artworkArtwork := artwork.NewArtwork(dataStore, fileCache, fFmpeg, provider)
	transcodingCache := core.GetTranscodingCache()
	mediaStreamer := core.NewMediaStreamer(dataStore, fFmpeg, transcodingCache)
	share := core.NewShare(dataStore)
	archiver := core.NewArchiver(mediaStreamer, dataStore, share)
	players := core.NewPlayers(dataStore)
	cacheWarmer := artwork.NewCacheWarmer(artworkArtwork, fileCache)
	broker := events.GetBroker()
	playlists := core.NewPlaylists(dataStore)
	metricsMetrics := metrics.NewPrometheusInstance(dataStore)
	scannerScanner := scanner.New(ctx, dataStore, cacheWarmer, broker, playlists, metricsMetrics)
	playTracker := scrobbler.GetPlayTracker(dataStore, broker)
	playbackServer := playback.GetInstance(dataStore)
	router := subsonic.New(dataStore, artworkArtwork, mediaStreamer, archiver, players, provider, scannerScanner, broker, playlists, playTracker, share, playbackServer)
	return router
}

func CreatePublicRouter() *public.Router {
	sqlDB := db.Db()
	dataStore := persistence.New(sqlDB)
	fileCache := artwork.GetImageCache()
	fFmpeg := ffmpeg.New()
	agentsAgents := agents.GetAgents(dataStore)
	provider := external.NewProvider(dataStore, agentsAgents)
	artworkArtwork := artwork.NewArtwork(dataStore, fileCache, fFmpeg, provider)
	transcodingCache := core.GetTranscodingCache()
	mediaStreamer := core.NewMediaStreamer(dataStore, fFmpeg, transcodingCache)
	share := core.NewShare(dataStore)
	archiver := core.NewArchiver(mediaStreamer, dataStore, share)
	router := public.New(dataStore, artworkArtwork, mediaStreamer, share, archiver)
	return router
}

func CreateLastFMRouter() *lastfm.Router {
	sqlDB := db.Db()
	dataStore := persistence.New(sqlDB)
	router := lastfm.NewRouter(dataStore)
	return router
}

func CreateListenBrainzRouter() *listenbrainz.Router {
	sqlDB := db.Db()
	dataStore := persistence.New(sqlDB)
	router := listenbrainz.NewRouter(dataStore)
	return router
}

func CreateInsights() metrics.Insights {
	sqlDB := db.Db()
	dataStore := persistence.New(sqlDB)
	insights := metrics.GetInstance(dataStore)
	return insights
}

func CreatePrometheus() metrics.Metrics {
	sqlDB := db.Db()
	dataStore := persistence.New(sqlDB)
	metricsMetrics := metrics.NewPrometheusInstance(dataStore)
	return metricsMetrics
}

func CreateScanner(ctx context.Context) scanner.Scanner {
	sqlDB := db.Db()
	dataStore := persistence.New(sqlDB)
	fileCache := artwork.GetImageCache()
	fFmpeg := ffmpeg.New()
	agentsAgents := agents.GetAgents(dataStore)
	provider := external.NewProvider(dataStore, agentsAgents)
	artworkArtwork := artwork.NewArtwork(dataStore, fileCache, fFmpeg, provider)
	cacheWarmer := artwork.NewCacheWarmer(artworkArtwork, fileCache)
	broker := events.GetBroker()
	playlists := core.NewPlaylists(dataStore)
	metricsMetrics := metrics.NewPrometheusInstance(dataStore)
	scannerScanner := scanner.New(ctx, dataStore, cacheWarmer, broker, playlists, metricsMetrics)
	return scannerScanner
}

func CreateScanWatcher(ctx context.Context) scanner.Watcher {
	sqlDB := db.Db()
	dataStore := persistence.New(sqlDB)
	fileCache := artwork.GetImageCache()
	fFmpeg := ffmpeg.New()
	agentsAgents := agents.GetAgents(dataStore)
	provider := external.NewProvider(dataStore, agentsAgents)
	artworkArtwork := artwork.NewArtwork(dataStore, fileCache, fFmpeg, provider)
	cacheWarmer := artwork.NewCacheWarmer(artworkArtwork, fileCache)
	broker := events.GetBroker()
	playlists := core.NewPlaylists(dataStore)
	metricsMetrics := metrics.NewPrometheusInstance(dataStore)
	scannerScanner := scanner.New(ctx, dataStore, cacheWarmer, broker, playlists, metricsMetrics)
	watcher := scanner.NewWatcher(dataStore, scannerScanner)
	return watcher
}

func GetPlaybackServer() playback.PlaybackServer {
	sqlDB := db.Db()
	dataStore := persistence.New(sqlDB)
	playbackServer := playback.GetInstance(dataStore)
	return playbackServer
}

// wire_injectors.go:

var allProviders = wire.NewSet(core.Set, artwork.Set, server.New, dlna.New, subsonic.New, nativeapi.New, public.New, persistence.New, lastfm.NewRouter, listenbrainz.NewRouter, events.GetBroker, scanner.New, scanner.NewWatcher, metrics.NewPrometheusInstance, db.Db)
