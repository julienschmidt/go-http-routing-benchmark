// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"testing"
)

// http://developer.github.com/v3/
var githubAPI = []route{
	// OAuth Authorizations
	{"GET", "/authorizations"},
	{"GET", "/authorizations/:id"},
	{"POST", "/authorizations"},
	//{"PUT", "/authorizations/clients/:client_id"},
	//{"PATCH", "/authorizations/:id"},
	{"DELETE", "/authorizations/:id"},
	{"GET", "/applications/:client_id/tokens/:access_token"},
	{"DELETE", "/applications/:client_id/tokens"},
	{"DELETE", "/applications/:client_id/tokens/:access_token"},

	// Activity
	{"GET", "/events"},
	{"GET", "/repos/:owner/:repo/events"},
	{"GET", "/networks/:owner/:repo/events"},
	{"GET", "/orgs/:org/events"},
	{"GET", "/users/:user/received_events"},
	{"GET", "/users/:user/received_events/public"},
	{"GET", "/users/:user/events"},
	{"GET", "/users/:user/events/public"},
	{"GET", "/users/:user/events/orgs/:org"},
	{"GET", "/feeds"},
	{"GET", "/notifications"},
	{"GET", "/repos/:owner/:repo/notifications"},
	{"PUT", "/notifications"},
	{"PUT", "/repos/:owner/:repo/notifications"},
	{"GET", "/notifications/threads/:id"},
	//{"PATCH", "/notifications/threads/:id"},
	{"GET", "/notifications/threads/:id/subscription"},
	{"PUT", "/notifications/threads/:id/subscription"},
	{"DELETE", "/notifications/threads/:id/subscription"},
	{"GET", "/repos/:owner/:repo/stargazers"},
	{"GET", "/users/:user/starred"},
	{"GET", "/user/starred"},
	{"GET", "/user/starred/:owner/:repo"},
	{"PUT", "/user/starred/:owner/:repo"},
	{"DELETE", "/user/starred/:owner/:repo"},
	{"GET", "/repos/:owner/:repo/subscribers"},
	{"GET", "/users/:user/subscriptions"},
	{"GET", "/user/subscriptions"},
	{"GET", "/repos/:owner/:repo/subscription"},
	{"PUT", "/repos/:owner/:repo/subscription"},
	{"DELETE", "/repos/:owner/:repo/subscription"},
	{"GET", "/user/subscriptions/:owner/:repo"},
	{"PUT", "/user/subscriptions/:owner/:repo"},
	{"DELETE", "/user/subscriptions/:owner/:repo"},

	// Gists
	{"GET", "/users/:user/gists"},
	{"GET", "/gists"},
	//{"GET", "/gists/public"},
	//{"GET", "/gists/starred"},
	{"GET", "/gists/:id"},
	{"POST", "/gists"},
	//{"PATCH", "/gists/:id"},
	{"PUT", "/gists/:id/star"},
	{"DELETE", "/gists/:id/star"},
	{"GET", "/gists/:id/star"},
	{"POST", "/gists/:id/forks"},
	{"DELETE", "/gists/:id"},

	// Git Data
	{"GET", "/repos/:owner/:repo/git/blobs/:sha"},
	{"POST", "/repos/:owner/:repo/git/blobs"},
	{"GET", "/repos/:owner/:repo/git/commits/:sha"},
	{"POST", "/repos/:owner/:repo/git/commits"},
	//{"GET", "/repos/:owner/:repo/git/refs/*ref"},
	{"GET", "/repos/:owner/:repo/git/refs"},
	{"POST", "/repos/:owner/:repo/git/refs"},
	//{"PATCH", "/repos/:owner/:repo/git/refs/*ref"},
	//{"DELETE", "/repos/:owner/:repo/git/refs/*ref"},
	{"GET", "/repos/:owner/:repo/git/tags/:sha"},
	{"POST", "/repos/:owner/:repo/git/tags"},
	{"GET", "/repos/:owner/:repo/git/trees/:sha"},
	{"POST", "/repos/:owner/:repo/git/trees"},

	// Issues
	{"GET", "/issues"},
	{"GET", "/user/issues"},
	{"GET", "/orgs/:org/issues"},
	{"GET", "/repos/:owner/:repo/issues"},
	{"GET", "/repos/:owner/:repo/issues/:number"},
	{"POST", "/repos/:owner/:repo/issues"},
	//{"PATCH", "/repos/:owner/:repo/issues/:number"},
	{"GET", "/repos/:owner/:repo/assignees"},
	{"GET", "/repos/:owner/:repo/assignees/:assignee"},
	{"GET", "/repos/:owner/:repo/issues/:number/comments"},
	//{"GET", "/repos/:owner/:repo/issues/comments"},
	//{"GET", "/repos/:owner/:repo/issues/comments/:id"},
	{"POST", "/repos/:owner/:repo/issues/:number/comments"},
	//{"PATCH", "/repos/:owner/:repo/issues/comments/:id"},
	//{"DELETE", "/repos/:owner/:repo/issues/comments/:id"},
	{"GET", "/repos/:owner/:repo/issues/:number/events"},
	//{"GET", "/repos/:owner/:repo/issues/events"},
	//{"GET", "/repos/:owner/:repo/issues/events/:id"},
	{"GET", "/repos/:owner/:repo/labels"},
	{"GET", "/repos/:owner/:repo/labels/:name"},
	{"POST", "/repos/:owner/:repo/labels"},
	//{"PATCH", "/repos/:owner/:repo/labels/:name"},
	{"DELETE", "/repos/:owner/:repo/labels/:name"},
	{"GET", "/repos/:owner/:repo/issues/:number/labels"},
	{"POST", "/repos/:owner/:repo/issues/:number/labels"},
	{"DELETE", "/repos/:owner/:repo/issues/:number/labels/:name"},
	{"PUT", "/repos/:owner/:repo/issues/:number/labels"},
	{"DELETE", "/repos/:owner/:repo/issues/:number/labels"},
	{"GET", "/repos/:owner/:repo/milestones/:number/labels"},
	{"GET", "/repos/:owner/:repo/milestones"},
	{"GET", "/repos/:owner/:repo/milestones/:number"},
	{"POST", "/repos/:owner/:repo/milestones"},
	//{"PATCH", "/repos/:owner/:repo/milestones/:number"},
	{"DELETE", "/repos/:owner/:repo/milestones/:number"},

	// Miscellaneous
	{"GET", "/emojis"},
	{"GET", "/gitignore/templates"},
	{"GET", "/gitignore/templates/:name"},
	{"POST", "/markdown"},
	{"POST", "/markdown/raw"},
	{"GET", "/meta"},
	{"GET", "/rate_limit"},

	// Organizations
	{"GET", "/users/:user/orgs"},
	{"GET", "/user/orgs"},
	{"GET", "/orgs/:org"},
	//{"PATCH", "/orgs/:org"},
	{"GET", "/orgs/:org/members"},
	{"GET", "/orgs/:org/members/:user"},
	{"DELETE", "/orgs/:org/members/:user"},
	{"GET", "/orgs/:org/public_members"},
	{"GET", "/orgs/:org/public_members/:user"},
	{"PUT", "/orgs/:org/public_members/:user"},
	{"DELETE", "/orgs/:org/public_members/:user"},
	{"GET", "/orgs/:org/teams"},
	{"GET", "/teams/:id"},
	{"POST", "/orgs/:org/teams"},
	//{"PATCH", "/teams/:id"},
	{"DELETE", "/teams/:id"},
	{"GET", "/teams/:id/members"},
	{"GET", "/teams/:id/members/:user"},
	{"PUT", "/teams/:id/members/:user"},
	{"DELETE", "/teams/:id/members/:user"},
	{"GET", "/teams/:id/repos"},
	{"GET", "/teams/:id/repos/:owner/:repo"},
	{"PUT", "/teams/:id/repos/:owner/:repo"},
	{"DELETE", "/teams/:id/repos/:owner/:repo"},
	{"GET", "/user/teams"},

	// Pull Requests
	{"GET", "/repos/:owner/:repo/pulls"},
	{"GET", "/repos/:owner/:repo/pulls/:number"},
	{"POST", "/repos/:owner/:repo/pulls"},
	//{"PATCH", "/repos/:owner/:repo/pulls/:number"},
	{"GET", "/repos/:owner/:repo/pulls/:number/commits"},
	{"GET", "/repos/:owner/:repo/pulls/:number/files"},
	{"GET", "/repos/:owner/:repo/pulls/:number/merge"},
	{"PUT", "/repos/:owner/:repo/pulls/:number/merge"},
	{"GET", "/repos/:owner/:repo/pulls/:number/comments"},
	//{"GET", "/repos/:owner/:repo/pulls/comments"},
	//{"GET", "/repos/:owner/:repo/pulls/comments/:number"},
	{"PUT", "/repos/:owner/:repo/pulls/:number/comments"},
	//{"PATCH", "/repos/:owner/:repo/pulls/comments/:number"},
	//{"DELETE", "/repos/:owner/:repo/pulls/comments/:number"},

	// Repositories
	{"GET", "/user/repos"},
	{"GET", "/users/:user/repos"},
	{"GET", "/orgs/:org/repos"},
	{"GET", "/repositories"},
	{"POST", "/user/repos"},
	{"POST", "/orgs/:org/repos"},
	{"GET", "/repos/:owner/:repo"},
	//{"PATCH", "/repos/:owner/:repo"},
	{"GET", "/repos/:owner/:repo/contributors"},
	{"GET", "/repos/:owner/:repo/languages"},
	{"GET", "/repos/:owner/:repo/teams"},
	{"GET", "/repos/:owner/:repo/tags"},
	{"GET", "/repos/:owner/:repo/branches"},
	{"GET", "/repos/:owner/:repo/branches/:branch"},
	{"DELETE", "/repos/:owner/:repo"},
	{"GET", "/repos/:owner/:repo/collaborators"},
	{"GET", "/repos/:owner/:repo/collaborators/:user"},
	{"PUT", "/repos/:owner/:repo/collaborators/:user"},
	{"DELETE", "/repos/:owner/:repo/collaborators/:user"},
	{"GET", "/repos/:owner/:repo/comments"},
	{"GET", "/repos/:owner/:repo/commits/:sha/comments"},
	{"POST", "/repos/:owner/:repo/commits/:sha/comments"},
	{"GET", "/repos/:owner/:repo/comments/:id"},
	//{"PATCH", "/repos/:owner/:repo/comments/:id"},
	{"DELETE", "/repos/:owner/:repo/comments/:id"},
	{"GET", "/repos/:owner/:repo/commits"},
	{"GET", "/repos/:owner/:repo/commits/:sha"},
	{"GET", "/repos/:owner/:repo/readme"},
	//{"GET", "/repos/:owner/:repo/contents/*path"},
	//{"PUT", "/repos/:owner/:repo/contents/*path"},
	//{"DELETE", "/repos/:owner/:repo/contents/*path"},
	//{"GET", "/repos/:owner/:repo/:archive_format/:ref"},
	{"GET", "/repos/:owner/:repo/keys"},
	{"GET", "/repos/:owner/:repo/keys/:id"},
	{"POST", "/repos/:owner/:repo/keys"},
	//{"PATCH", "/repos/:owner/:repo/keys/:id"},
	{"DELETE", "/repos/:owner/:repo/keys/:id"},
	{"GET", "/repos/:owner/:repo/downloads"},
	{"GET", "/repos/:owner/:repo/downloads/:id"},
	{"DELETE", "/repos/:owner/:repo/downloads/:id"},
	{"GET", "/repos/:owner/:repo/forks"},
	{"POST", "/repos/:owner/:repo/forks"},
	{"GET", "/repos/:owner/:repo/hooks"},
	{"GET", "/repos/:owner/:repo/hooks/:id"},
	{"POST", "/repos/:owner/:repo/hooks"},
	//{"PATCH", "/repos/:owner/:repo/hooks/:id"},
	{"POST", "/repos/:owner/:repo/hooks/:id/tests"},
	{"DELETE", "/repos/:owner/:repo/hooks/:id"},
	{"POST", "/repos/:owner/:repo/merges"},
	{"GET", "/repos/:owner/:repo/releases"},
	{"GET", "/repos/:owner/:repo/releases/:id"},
	{"POST", "/repos/:owner/:repo/releases"},
	//{"PATCH", "/repos/:owner/:repo/releases/:id"},
	{"DELETE", "/repos/:owner/:repo/releases/:id"},
	{"GET", "/repos/:owner/:repo/releases/:id/assets"},
	{"GET", "/repos/:owner/:repo/stats/contributors"},
	{"GET", "/repos/:owner/:repo/stats/commit_activity"},
	{"GET", "/repos/:owner/:repo/stats/code_frequency"},
	{"GET", "/repos/:owner/:repo/stats/participation"},
	{"GET", "/repos/:owner/:repo/stats/punch_card"},
	{"GET", "/repos/:owner/:repo/statuses/:ref"},
	{"POST", "/repos/:owner/:repo/statuses/:ref"},

	// Search
	{"GET", "/search/repositories"},
	{"GET", "/search/code"},
	{"GET", "/search/issues"},
	{"GET", "/search/users"},
	{"GET", "/legacy/issues/search/:owner/:repository/:state/:keyword"},
	{"GET", "/legacy/repos/search/:keyword"},
	{"GET", "/legacy/user/search/:keyword"},
	{"GET", "/legacy/user/email/:email"},

	// Users
	{"GET", "/users/:user"},
	{"GET", "/user"},
	//{"PATCH", "/user"},
	{"GET", "/users"},
	{"GET", "/user/emails"},
	{"POST", "/user/emails"},
	{"DELETE", "/user/emails"},
	{"GET", "/users/:user/followers"},
	{"GET", "/user/followers"},
	{"GET", "/users/:user/following"},
	{"GET", "/user/following"},
	{"GET", "/user/following/:user"},
	{"GET", "/users/:user/following/:target_user"},
	{"PUT", "/user/following/:user"},
	{"DELETE", "/user/following/:user"},
	{"GET", "/users/:user/keys"},
	{"GET", "/user/keys"},
	{"GET", "/user/keys/:id"},
	{"POST", "/user/keys"},
	//{"PATCH", "/user/keys/:id"},
	{"DELETE", "/user/keys/:id"},
}

var (
	githubAce             http.Handler
	githubAero            http.Handler
	githubBear            http.Handler
	githubBeego           http.Handler
	githubBone            http.Handler
	githubChi             http.Handler
	githubCloudyKitRouter http.Handler
	githubDenco           http.Handler
	githubEcho            http.Handler
	githubGin             http.Handler
	githubGocraftWeb      http.Handler
	githubGoji            http.Handler
	githubGojiv2          http.Handler
	githubGoJsonRest      http.Handler
	githubGoRestful       http.Handler
	githubGorillaMux      http.Handler
	githubGowwwRouter     http.Handler
	githubHttpRouter      http.Handler
	githubHttpTreeMux     http.Handler
	githubKocha           http.Handler
	githubLARS            http.Handler
	githubLenrouter       http.Handler
	githubMacaron         http.Handler
	githubMartini         http.Handler
	githubPat             http.Handler
	githubPossum          http.Handler
	githubR2router        http.Handler
	githubRevel           http.Handler
	githubRivet           http.Handler
	githubTango           http.Handler
	githubTigerTonic      http.Handler
	githubTraffic         http.Handler
	githubVulcan          http.Handler
	// githubZeus        http.Handler
)

func init() {
	println("#GithubAPI Routes:", len(githubAPI))

	calcMem("Ace", func() {
		githubAce = loadAce(githubAPI)
	})
	calcMem("Aero", func() {
		githubAero = loadAero(githubAPI)
	})
	calcMem("Bear", func() {
		githubBear = loadBear(githubAPI)
	})
	calcMem("Beego", func() {
		githubBeego = loadBeego(githubAPI)
	})
	calcMem("Bone", func() {
		githubBone = loadBone(githubAPI)
	})
	calcMem("Chi", func() {
		githubChi = loadChi(githubAPI)
	})
	calcMem("CloudyKitRouter", func() {
		githubCloudyKitRouter = loadCloudyKitRouter(githubAPI)
	})
	calcMem("Denco", func() {
		githubDenco = loadDenco(githubAPI)
	})
	calcMem("Echo", func() {
		githubEcho = loadEcho(githubAPI)
	})
	calcMem("Gin", func() {
		githubGin = loadGin(githubAPI)
	})
	calcMem("GocraftWeb", func() {
		githubGocraftWeb = loadGocraftWeb(githubAPI)
	})
	calcMem("Goji", func() {
		githubGoji = loadGoji(githubAPI)
	})
	calcMem("Gojiv2", func() {
		githubGojiv2 = loadGojiv2(githubAPI)
	})
	calcMem("GoJsonRest", func() {
		githubGoJsonRest = loadGoJsonRest(githubAPI)
	})
	calcMem("GoRestful", func() {
		githubGoRestful = loadGoRestful(githubAPI)
	})
	calcMem("GorillaMux", func() {
		githubGorillaMux = loadGorillaMux(githubAPI)
	})
	calcMem("GowwwRouter", func() {
		githubGowwwRouter = loadGowwwRouter(githubAPI)
	})
	calcMem("HttpRouter", func() {
		githubHttpRouter = loadHttpRouter(githubAPI)
	})
	calcMem("HttpTreeMux", func() {
		githubHttpTreeMux = loadHttpTreeMux(githubAPI)
	})
	calcMem("Kocha", func() {
		githubKocha = loadKocha(githubAPI)
	})
	calcMem("LARS", func() {
		githubLARS = loadLARS(githubAPI)
	})
	calcMem("Lenrouter", func() {
		githubLenrouter = loadLenrouter(githubAPI)
	})
	calcMem("Macaron", func() {
		githubMacaron = loadMacaron(githubAPI)
	})
	calcMem("Martini", func() {
		githubMartini = loadMartini(githubAPI)
	})
	calcMem("Pat", func() {
		githubPat = loadPat(githubAPI)
	})
	calcMem("Possum", func() {
		githubPossum = loadPossum(githubAPI)
	})
	calcMem("R2router", func() {
		githubR2router = loadR2router(githubAPI)
	})
	// calcMem("Revel", func() {
	// 	githubRevel = loadRevel(githubAPI)
	// })
	calcMem("Rivet", func() {
		githubRivet = loadRivet(githubAPI)
	})
	calcMem("Tango", func() {
		githubTango = loadTango(githubAPI)
	})
	calcMem("TigerTonic", func() {
		githubTigerTonic = loadTigerTonic(githubAPI)
	})
	calcMem("Traffic", func() {
		githubTraffic = loadTraffic(githubAPI)
	})
	calcMem("Vulcan", func() {
		githubVulcan = loadVulcan(githubAPI)
	})
	// calcMem("Zeus", func() {
	// 	githubZeus = loadZeus(githubAPI)
	// })

	println()
}

// Static
func BenchmarkAce_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubAce, req)
}
func BenchmarkAero_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubAero, req)
}
func BenchmarkBear_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubBear, req)
}
func BenchmarkBeego_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubBeego, req)
}
func BenchmarkBone_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubBone, req)
}
func BenchmarkCloudyKitRouter_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubCloudyKitRouter, req)
}
func BenchmarkChi_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubChi, req)
}
func BenchmarkDenco_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubDenco, req)
}
func BenchmarkEcho_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubEcho, req)
}
func BenchmarkGin_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubGin, req)
}
func BenchmarkGocraftWeb_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubGocraftWeb, req)
}
func BenchmarkGoji_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubGoji, req)
}
func BenchmarkGojiv2_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubGojiv2, req)
}
func BenchmarkGoRestful_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubGoRestful, req)
}
func BenchmarkGoJsonRest_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubGoJsonRest, req)
}
func BenchmarkGorillaMux_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubGorillaMux, req)
}
func BenchmarkGowwwRouter_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubGowwwRouter, req)
}
func BenchmarkHttpRouter_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubHttpRouter, req)
}
func BenchmarkHttpTreeMux_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubHttpTreeMux, req)
}
func BenchmarkKocha_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubKocha, req)
}
func BenchmarkLARS_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubLARS, req)
}
func BenchmarkLenrouter_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubLenrouter, req)
}
func BenchmarkMacaron_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubMacaron, req)
}
func BenchmarkMartini_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubMartini, req)
}
func BenchmarkPat_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubPat, req)
}
func BenchmarkPossum_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubPossum, req)
}
func BenchmarkR2router_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubR2router, req)
}

// func BenchmarkRevel_GithubStatic(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/user/repos", nil)
// 	benchRequest(b, githubRevel, req)
// }
func BenchmarkRivet_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubRivet, req)
}
func BenchmarkTango_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubTango, req)
}
func BenchmarkTigerTonic_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubTigerTonic, req)
}
func BenchmarkTraffic_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubTraffic, req)
}
func BenchmarkVulcan_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubVulcan, req)
}

// func BenchmarkZeus_GithubStatic(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/user/repos", nil)
// 	benchRequest(b, githubZeus, req)
// }

// Param
func BenchmarkAce_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubAce, req)
}
func BenchmarkAero_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubAero, req)
}
func BenchmarkBear_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubBear, req)
}
func BenchmarkBeego_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubBeego, req)
}
func BenchmarkBone_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubBone, req)
}
func BenchmarkChi_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubChi, req)
}
func BenchmarkCloudyKitRouter_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubCloudyKitRouter, req)
}
func BenchmarkDenco_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubDenco, req)
}
func BenchmarkEcho_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubEcho, req)
}
func BenchmarkGin_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGin, req)
}
func BenchmarkGocraftWeb_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGocraftWeb, req)
}
func BenchmarkGoji_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGoji, req)
}
func BenchmarkGojiv2_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGojiv2, req)
}
func BenchmarkGoJsonRest_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGoJsonRest, req)
}
func BenchmarkGoRestful_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGoRestful, req)
}
func BenchmarkGorillaMux_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGorillaMux, req)
}
func BenchmarkGowwwRouter_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGowwwRouter, req)
}
func BenchmarkHttpRouter_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubHttpRouter, req)
}
func BenchmarkHttpTreeMux_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubHttpTreeMux, req)
}
func BenchmarkKocha_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubKocha, req)
}
func BenchmarkLARS_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubLARS, req)
}
func BenchmarkLenrouter_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubLenrouter, req)
}
func BenchmarkMacaron_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubMacaron, req)
}
func BenchmarkMartini_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubMartini, req)
}
func BenchmarkPat_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubPat, req)
}
func BenchmarkPossum_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubPossum, req)
}
func BenchmarkR2router_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubR2router, req)
}

// func BenchmarkRevel_GithubParam(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
// 	benchRequest(b, githubRevel, req)
// }
func BenchmarkRivet_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubRivet, req)
}
func BenchmarkTango_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubTango, req)
}
func BenchmarkTigerTonic_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubTigerTonic, req)
}
func BenchmarkTraffic_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubTraffic, req)
}
func BenchmarkVulcan_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubVulcan, req)
}

// func BenchmarkZeus_GithubParam(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
// 	benchRequest(b, githubZeus, req)
// }

// All routes
func BenchmarkAce_GithubAll(b *testing.B) {
	benchRoutes(b, githubAce, githubAPI)
}
func BenchmarkAero_GithubAll(b *testing.B) {
	benchRoutes(b, githubAero, githubAPI)
}
func BenchmarkBear_GithubAll(b *testing.B) {
	benchRoutes(b, githubBear, githubAPI)
}
func BenchmarkBeego_GithubAll(b *testing.B) {
	benchRoutes(b, githubBeego, githubAPI)
}
func BenchmarkBone_GithubAll(b *testing.B) {
	benchRoutes(b, githubBone, githubAPI)
}
func BenchmarkChi_GithubAll(b *testing.B) {
	benchRoutes(b, githubChi, githubAPI)
}
func BenchmarkCloudyKitRouter_GithubAll(b *testing.B) {
	benchRoutes(b, githubCloudyKitRouter, githubAPI)
}
func BenchmarkDenco_GithubAll(b *testing.B) {
	benchRoutes(b, githubDenco, githubAPI)
}
func BenchmarkEcho_GithubAll(b *testing.B) {
	benchRoutes(b, githubEcho, githubAPI)
}
func BenchmarkGin_GithubAll(b *testing.B) {
	benchRoutes(b, githubGin, githubAPI)
}
func BenchmarkGocraftWeb_GithubAll(b *testing.B) {
	benchRoutes(b, githubGocraftWeb, githubAPI)
}
func BenchmarkGoji_GithubAll(b *testing.B) {
	benchRoutes(b, githubGoji, githubAPI)
}
func BenchmarkGojiv2_GithubAll(b *testing.B) {
	benchRoutes(b, githubGojiv2, githubAPI)
}
func BenchmarkGoJsonRest_GithubAll(b *testing.B) {
	benchRoutes(b, githubGoJsonRest, githubAPI)
}
func BenchmarkGoRestful_GithubAll(b *testing.B) {
	benchRoutes(b, githubGoRestful, githubAPI)
}
func BenchmarkGorillaMux_GithubAll(b *testing.B) {
	benchRoutes(b, githubGorillaMux, githubAPI)
}
func BenchmarkGowwwRouter_GithubAll(b *testing.B) {
	benchRoutes(b, githubGowwwRouter, githubAPI)
}
func BenchmarkHttpRouter_GithubAll(b *testing.B) {
	benchRoutes(b, githubHttpRouter, githubAPI)
}
func BenchmarkHttpTreeMux_GithubAll(b *testing.B) {
	benchRoutes(b, githubHttpTreeMux, githubAPI)
}
func BenchmarkKocha_GithubAll(b *testing.B) {
	benchRoutes(b, githubKocha, githubAPI)
}
func BenchmarkLARS_GithubAll(b *testing.B) {
	benchRoutes(b, githubLARS, githubAPI)
}
func BenchmarkLenrouter_GithubAll(b *testing.B) {
	benchRoutes(b, githubLenrouter, githubAPI)
}
func BenchmarkMacaron_GithubAll(b *testing.B) {
	benchRoutes(b, githubMacaron, githubAPI)
}
func BenchmarkMartini_GithubAll(b *testing.B) {
	benchRoutes(b, githubMartini, githubAPI)
}
func BenchmarkPat_GithubAll(b *testing.B) {
	benchRoutes(b, githubPat, githubAPI)
}
func BenchmarkPossum_GithubAll(b *testing.B) {
	benchRoutes(b, githubPossum, githubAPI)
}
func BenchmarkR2router_GithubAll(b *testing.B) {
	benchRoutes(b, githubR2router, githubAPI)
}

// func BenchmarkRevel_GithubAll(b *testing.B) {
// 	benchRoutes(b, githubRevel, githubAPI)
// }
func BenchmarkRivet_GithubAll(b *testing.B) {
	benchRoutes(b, githubRivet, githubAPI)
}
func BenchmarkTango_GithubAll(b *testing.B) {
	benchRoutes(b, githubTango, githubAPI)
}
func BenchmarkTigerTonic_GithubAll(b *testing.B) {
	benchRoutes(b, githubTigerTonic, githubAPI)
}
func BenchmarkTraffic_GithubAll(b *testing.B) {
	benchRoutes(b, githubTraffic, githubAPI)
}
func BenchmarkVulcan_GithubAll(b *testing.B) {
	benchRoutes(b, githubVulcan, githubAPI)
}

// func BenchmarkZeus_GithubAll(b *testing.B) {
// 	benchRoutes(b, githubZeus, githubAPI)
// }
