package e2e_test

import (
	"bytes"
	"context"
	"fmt"
	gloo_config_core "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/core/v3"
	gloo_type_matcher "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/type/matcher/v3"
	glootype "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/type/v3"
	"io"
	"net/http"
	"strings"

	"github.com/envoyproxy/go-control-plane/pkg/wellknown"

	csrf "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/filters/http/csrf/v3"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/cors"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/test/services"
	"github.com/solo-io/gloo/test/v1helpers"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

type perCsrfTestData struct {
	up            *gloov1.Upstream
	envoyInstance *services.EnvoyInstance
	envoyPort     uint32
	envoyAdminUrl string
}
type csrfTestData struct {
	testClients services.TestClients
	ctx         context.Context
	cancel      context.CancelFunc
	per         perCsrfTestData
}

var _ = FDescribe("csrf 2", func() {

	var td csrfTestData

	const (
		corsFilterString       = `"name": "` + wellknown.CORS + `"`
		corsActiveConfigString = `"cors":`
	)

	apiFilter := &gloo_config_core.RuntimeFractionalPercent{
		DefaultValue: &glootype.FractionalPercent{
			Numerator:   uint32(100),
			Denominator: glootype.FractionalPercent_HUNDRED,
		},
	}

	apiAdditionalOrigins := []*gloo_type_matcher.StringMatcher{
		{
			MatchPattern: &gloo_type_matcher.StringMatcher_SafeRegex{
				SafeRegex: &gloo_type_matcher.RegexMatcher{
					EngineType: &gloo_type_matcher.RegexMatcher_GoogleRe2{
						GoogleRe2: &gloo_type_matcher.RegexMatcher_GoogleRE2{},
					},
					Regex: "allowThisOne.solo.io",
				},
			},
			IgnoreCase: true,
		},
	}

	BeforeEach(func() {
		td.ctx, td.cancel = context.WithCancel(context.Background())
		td.testClients = services.RunGateway(td.ctx, true)
		td.per = perCsrfTestData{}
	})

	AfterEach(func() {
		td.cancel()
	})

	Context("with envoy", func() {

		BeforeEach(func() {
			var err error
			td.per.envoyInstance, err = envoyFactory.NewEnvoyInstance()
			Expect(err).NotTo(HaveOccurred())
			td.per.envoyAdminUrl = fmt.Sprintf("http://%s:%d/config_dump",
				td.per.envoyInstance.LocalAddr(),
				td.per.envoyInstance.AdminPort)

			err = td.per.envoyInstance.Run(td.testClients.GlooPort)
			Expect(err).NotTo(HaveOccurred())

			td.per.up = td.setupUpstream()
		})

		AfterEach(func() {
			if td.per.envoyInstance != nil {
				td.per.envoyInstance.Clean()
			}
		})

		It("should run with cors", func() {

			allowedOrigins := []string{"allowThisOne.solo.io"}
			//allowedOrigins := []string{"*"}
			allowedMethods := []string{"GET", "POST"}
			cors := &cors.CorsPolicy{
				AllowOrigin:      allowedOrigins,
				AllowOriginRegex: allowedOrigins,
				AllowMethods:     allowedMethods,
			}

			csrf := &csrf.CsrfPolicy{
				FilterEnabled:     apiFilter,
				AdditionalOrigins: apiAdditionalOrigins,
			}

			td.setupInitialProxy(cors, csrf)
			envoyConfig := td.per.getEnvoyConfig()

			By("Check config")
			print(envoyConfig)
			//Expect(envoyConfig).To(MatchRegexp(corsFilterString))
			//Expect(envoyConfig).To(MatchRegexp(corsActiveConfigString))
			//Expect(envoyConfig).To(MatchRegexp(allowedOrigins[0]))

			By("Request with allowed origin")
			mockOrigin := allowedOrigins[0]
			//h := td.per.getOptions(mockOrigin, "GET")
			h :=  td.per.getOptionsCsrf(mockOrigin, "GET")

			v, ok := h[requestACHMethods]
			Expect(ok).To(BeTrue())
			Expect(strings.Split(v[0], ",")).Should(ConsistOf(allowedMethods))
			v, ok = h[requestACHOrigin]
			Expect(ok).To(BeTrue())
			Expect(len(v)).To(Equal(1))
			Expect(v[0]).To(Equal(mockOrigin))

			By("Request with disallowed origin")
			mockOrigin = "http://example.com"
			h = td.per.getOptions(mockOrigin, "GET")
			v, ok = h[requestACHMethods]
			Expect(ok).To(BeFalse())

		})
		//It("should run without cors", func() {
		//	td.setupInitialProxy(nil, nil)
		//	envoyConfig := td.per.getEnvoyConfig()
		//
		//	Expect(envoyConfig).To(MatchRegexp(corsFilterString))
		//	Expect(envoyConfig).NotTo(MatchRegexp(corsActiveConfigString))
		//})
	})
})

func (td *csrfTestData) getGlooCsrfProxy(cors *cors.CorsPolicy, csrf * csrf.CsrfPolicy) (*gloov1.Proxy, error) {
	readProxy, err := td.testClients.ProxyClient.Read("default", "proxy", clients.ReadOpts{})
	if err != nil {
		return nil, err
	}
	return td.per.getGlooCsrfProxyWithVersion(readProxy.Metadata.ResourceVersion, cors, csrf), nil
}

func (ptd *perCsrfTestData) getGlooCsrfProxyWithVersion(resourceVersion string, cors *cors.CorsPolicy, csrf *csrf.CsrfPolicy) *gloov1.Proxy {
	return &gloov1.Proxy{
		Metadata: &core.Metadata{
			Name:            "proxy",
			Namespace:       "default",
			ResourceVersion: resourceVersion,
		},
		Listeners: []*gloov1.Listener{{
			Name:        "listener",
			BindAddress: "0.0.0.0",
			BindPort:    ptd.envoyPort,
			ListenerType: &gloov1.Listener_HttpListener{
				HttpListener: &gloov1.HttpListener{
					VirtualHosts: []*gloov1.VirtualHost{{
						Name:    "virt1",
						Domains: []string{"*"},
						Routes: []*gloov1.Route{{
							Action: &gloov1.Route_RouteAction{
								RouteAction: &gloov1.RouteAction{
									Destination: &gloov1.RouteAction_Single{
										Single: &gloov1.Destination{
											DestinationType: &gloov1.Destination_Upstream{
												Upstream: ptd.up.Metadata.Ref(),
											},
										},
									},
								},
							},
						}},
						Options: &gloov1.VirtualHostOptions{
							//Cors: cors,
							Csrf: csrf, // Bug
						},
					}},
				},
			},
		}},
	}
}

func (td *csrfTestData) setupProxy(proxy *gloov1.Proxy) error {
	proxyCli := td.testClients.ProxyClient
	_, err := proxyCli.Write(proxy, clients.WriteOpts{OverwriteExisting: true})
	return err
}

func (td *csrfTestData) setupInitialProxy(cors *cors.CorsPolicy, csrf *csrf.CsrfPolicy) {
	By("Setup proxy")
	td.per.envoyPort = defaults.HttpPort
	proxy := td.per.getGlooCsrfProxyWithVersion("", cors, csrf)
	err := td.setupProxy(proxy)
	// Call with retries to ensure proxy is available
	Eventually(func() error {
		proxy, err := td.getGlooCsrfProxy(cors, csrf)
		if err != nil {
			return err
		}
		return td.setupProxy(proxy)
	}, "10s", ".1s").Should(BeNil())
	Expect(err).NotTo(HaveOccurred())
	Eventually(func() error {
		_, err := http.Get(fmt.Sprintf("http://%s:%d/status/200", "localhost", td.per.envoyPort))
		if err != nil {
			return err
		}
		return nil
	}, "10s", ".1s").Should(BeNil())
}

func (td *csrfTestData) setupUpstream() *gloov1.Upstream {
	tu := v1helpers.NewTestHttpUpstream(td.ctx, td.per.envoyInstance.LocalAddr())
	// drain channel as we don't care about it
	go func() {
		for range tu.C {
		}
	}()
	up := tu.Upstream
	_, err := td.testClients.UpstreamClient.Write(up, clients.WriteOpts{OverwriteExisting: true})
	Expect(err).NotTo(HaveOccurred())
	return up
}

func (ptd *perCsrfTestData) getOptionsCsrf(origin, method string) http.Header {
	h := http.Header{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:%v", ptd.envoyPort), nil)
	if err != nil {
		print(err)
	}
	req.Header.Set("Origin", origin)
	req.Header.Set("Access-Control-Request-Method", method)
	req.Header.Set("Access-Control-Request-Headers", "X-Requested-With")

	resp, err := http.DefaultClient.Do(req)
	print(resp)

	return h
}

// To test this with curl:
// curl -H "Origin: http://example.com" \
//   -H "Access-Control-Request-Method: POST" \
//   -H "Access-Control-Request-Headers: X-Requested-With" \
//   -X OPTIONS --verbose localhost:11082
func (ptd *perCsrfTestData) getOptions(origin, method string) http.Header {
	h := http.Header{}
	Eventually(func() error {
		req, err := http.NewRequest("OPTIONS", fmt.Sprintf("http://localhost:%v", ptd.envoyPort), nil)
		if err != nil {
			return err
		}
		req.Header.Set("Origin", origin)
		req.Header.Set("Access-Control-Request-Method", method)
		req.Header.Set("Access-Control-Request-Headers", "X-Requested-With")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		h = resp.Header
		return nil
	}, "10s", ".1s").Should(BeNil())
	return h
}

func (ptd *perCsrfTestData) getEnvoyConfig() string {
	By("Get config")
	envoyConfig := ""
	Eventually(func() error {
		r, err := http.Get(ptd.envoyAdminUrl)
		if err != nil {
			return err
		}
		p := new(bytes.Buffer)
		if _, err := io.Copy(p, r.Body); err != nil {
			return err
		}
		defer r.Body.Close()
		envoyConfig = p.String()
		return nil
	}, "10s", ".1s").Should(BeNil())
	return envoyConfig
}
