package store_test

import (
	"github.com/cloudfoundry/gunk/workpool"
	. "github.com/cloudfoundry/hm9000/store"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry/hm9000/config"
	"github.com/cloudfoundry/hm9000/models"
	"github.com/cloudfoundry/hm9000/testhelpers/appfixture"
	"github.com/cloudfoundry/hm9000/testhelpers/fakelogger"
	"github.com/cloudfoundry/storeadapter"
	"github.com/cloudfoundry/storeadapter/etcdstoreadapter"
)

var _ = Describe("Actual State", func() {
	var (
		store        Store
		storeAdapter storeadapter.StoreAdapter
		conf         *config.Config
		dea          appfixture.DeaFixture
		otherDea     appfixture.DeaFixture
	)

	BeforeEach(func() {
		var err error
		conf, err = config.DefaultConfig()
		Expect(err).NotTo(HaveOccurred())
		wpool, err := workpool.NewWorkPool(conf.StoreMaxConcurrentRequests)
		Expect(err).NotTo(HaveOccurred())
		storeAdapter, err = etcdstoreadapter.New(
			&etcdstoreadapter.ETCDOptions{ClusterUrls: etcdRunner.NodeURLS()},
			wpool,
		)
		Expect(err).NotTo(HaveOccurred())
		err = storeAdapter.Connect()
		Expect(err).NotTo(HaveOccurred())
		conf.StoreHeartbeatCacheRefreshIntervalInMilliseconds = 100
		store = NewStore(conf, storeAdapter, fakelogger.NewFakeLogger())

		dea = appfixture.NewDeaFixture()
		otherDea = appfixture.NewDeaFixture()
	})

	AfterEach(func() {
		storeAdapter.Disconnect()
	})

	Describe("#EnsureCacheIsReady", func() {

		BeforeEach(func() {
			store.SyncHeartbeats(dea.HeartbeatWith(
				dea.GetApp(0).InstanceAtIndex(1).Heartbeat(),
				dea.GetApp(1).InstanceAtIndex(3).Heartbeat(),
			))
		})

		It("populates the local cache from the store", func() {
			err := store.EnsureCacheIsReady()
			Expect(err).NotTo(HaveOccurred())

			results := store.GetCachedInstanceHeartbeats()
			Expect(results).To(ContainElement(dea.GetApp(0).InstanceAtIndex(1).Heartbeat()))
			Expect(results).To(ContainElement(dea.GetApp(1).InstanceAtIndex(3).Heartbeat()))
		})
	})

	Describe("Saving actual state", func() {
		JustBeforeEach(func() {
			store.SyncHeartbeats(dea.HeartbeatWith(
				dea.GetApp(0).InstanceAtIndex(1).Heartbeat(),
				dea.GetApp(1).InstanceAtIndex(3).Heartbeat(),
			))
		})

		Context("when there are no heartbeats stored for the DEA in question", func() {
			BeforeEach(func() {
				results, err := store.GetInstanceHeartbeats()
				Expect(err).NotTo(HaveOccurred())
				Expect(results).To(HaveLen(0))

				results = store.GetCachedInstanceHeartbeats()
				Expect(results).To(HaveLen(0))
			})

			It("saves the instance heartbeats for the passed-in heartbeat", func() {
				results, err := store.GetInstanceHeartbeats()
				Expect(err).NotTo(HaveOccurred())
				Expect(results).To(HaveLen(2))
				Expect(results).To(ContainElement(dea.GetApp(0).InstanceAtIndex(1).Heartbeat()))
				Expect(results).To(ContainElement(dea.GetApp(1).InstanceAtIndex(3).Heartbeat()))
			})

			It("stores the instance heartbeats in the cache", func() {
				results := store.GetCachedInstanceHeartbeats()
				Expect(results).To(HaveLen(2))
				Expect(results).To(ContainElement(dea.GetApp(0).InstanceAtIndex(1).Heartbeat()))
				Expect(results).To(ContainElement(dea.GetApp(1).InstanceAtIndex(3).Heartbeat()))
			})

		})

		Context("when there are already instance heartbeats stored for the DEA in question", func() {
			var modifiedHeartbeat models.InstanceHeartbeat
			JustBeforeEach(func() {
				modifiedHeartbeat = dea.GetApp(1).InstanceAtIndex(3).Heartbeat()
				modifiedHeartbeat.State = models.InstanceStateEvacuating
				store.SyncHeartbeats(dea.HeartbeatWith(
					modifiedHeartbeat,
					dea.GetApp(2).InstanceAtIndex(2).Heartbeat(),
				))
			})

			It("syncs the heartbeats (add new ones, adjust ones that have changed state, and delete old ones)", func() {
				results, err := store.GetInstanceHeartbeats()
				Expect(err).NotTo(HaveOccurred())
				Expect(results).To(HaveLen(2))
				Expect(results).To(ContainElement(modifiedHeartbeat))
				Expect(results).To(ContainElement(dea.GetApp(2).InstanceAtIndex(2).Heartbeat()))
			})
		})

		Context("when saving heartbeats from multiple DEAs at once", func() {
			var modifiedHeartbeat models.InstanceHeartbeat
			var yetAnotherDea appfixture.DeaFixture

			JustBeforeEach(func() {
				yetAnotherDea = appfixture.NewDeaFixture()

				store.SyncHeartbeats(dea.HeartbeatWith(
					dea.GetApp(0).InstanceAtIndex(1).Heartbeat(),
					dea.GetApp(1).InstanceAtIndex(3).Heartbeat(),
				), otherDea.HeartbeatWith(
					otherDea.GetApp(3).InstanceAtIndex(0).Heartbeat(),
					otherDea.GetApp(2).InstanceAtIndex(1).Heartbeat(),
				), yetAnotherDea.HeartbeatWith(
					yetAnotherDea.GetApp(0).InstanceAtIndex(0).Heartbeat(),
				))

				modifiedHeartbeat = dea.GetApp(1).InstanceAtIndex(3).Heartbeat()
				modifiedHeartbeat.State = models.InstanceStateEvacuating
				store.SyncHeartbeats(dea.HeartbeatWith(
					modifiedHeartbeat,
					dea.GetApp(2).InstanceAtIndex(2).Heartbeat(),
				), otherDea.HeartbeatWith(
					otherDea.GetApp(2).InstanceAtIndex(1).Heartbeat(),
					otherDea.GetApp(3).InstanceAtIndex(2).Heartbeat(),
				))
			})

			It("collates the hearbeats", func() {
				results, err := store.GetInstanceHeartbeats()
				Expect(err).NotTo(HaveOccurred())
				Expect(results).To(HaveLen(5))
				Expect(results).To(ContainElement(modifiedHeartbeat))
				Expect(results).To(ContainElement(dea.GetApp(2).InstanceAtIndex(2).Heartbeat()))
				Expect(results).To(ContainElement(otherDea.GetApp(2).InstanceAtIndex(1).Heartbeat()))
				Expect(results).To(ContainElement(otherDea.GetApp(3).InstanceAtIndex(2).Heartbeat()))
				Expect(results).To(ContainElement(yetAnotherDea.GetApp(0).InstanceAtIndex(0).Heartbeat()))
			})
		})

		Context("when one of the keys fails to delete", func() {
			It("should soldier on", func() {
				store.SyncHeartbeats(dea.HeartbeatWith(
					dea.GetApp(0).InstanceAtIndex(1).Heartbeat(),
					dea.GetApp(1).InstanceAtIndex(3).Heartbeat(),
				))

				done := make(chan error, 2)

				go func() {
					done <- store.SyncHeartbeats(dea.HeartbeatWith(
						dea.GetApp(0).InstanceAtIndex(1).Heartbeat(),
					))
				}()

				go func() {
					done <- store.SyncHeartbeats(dea.HeartbeatWith(
						dea.GetApp(0).InstanceAtIndex(1).Heartbeat(),
					))
				}()

				err1 := <-done
				err2 := <-done
				Expect(err1).NotTo(HaveOccurred())
				Expect(err2).NotTo(HaveOccurred())
			})
		})

		Context("when the in-memory cache no longer matches the store", func() {
			It("recovers by updating the store when it receives the next heartbeat", func() {
				corruptedHeartbeat := dea.GetApp(0).InstanceAtIndex(1).Heartbeat()
				storeAdapter.Delete("/hm/v1/apps/actual/" + store.AppKey(corruptedHeartbeat.AppGuid, corruptedHeartbeat.AppVersion) + "/" + corruptedHeartbeat.InstanceGuid)

				results, err := store.GetInstanceHeartbeats()
				Expect(err).NotTo(HaveOccurred())
				Expect(results).To(HaveLen(1))

				results = store.GetCachedInstanceHeartbeats()
				Expect(results).To(HaveLen(2))

				store.SyncHeartbeats(dea.HeartbeatWith(
					dea.GetApp(0).InstanceAtIndex(1).Heartbeat(),
					dea.GetApp(1).InstanceAtIndex(3).Heartbeat(),
				))

				results, err = store.GetInstanceHeartbeats()
				Expect(err).NotTo(HaveOccurred())
				Expect(results).To(HaveLen(2))
			})
		})
	})

	Describe("Fetching all actual state", func() {
		Context("when there is none saved", func() {
			It("comes back empty", func() {
				results, err := store.GetInstanceHeartbeats()
				Expect(err).NotTo(HaveOccurred())
				Expect(results).To(BeEmpty())
			})
		})

		Context("when there is actual state saved", func() {
			var heartbeatOnDea, heartbeatOnOtherDea models.InstanceHeartbeat

			BeforeEach(func() {
				appOnBothDeas := appfixture.NewAppFixture()

				heartbeatOnDea = appOnBothDeas.InstanceAtIndex(0).Heartbeat()
				heartbeatOnDea.DeaGuid = dea.DeaGuid

				heartbeatOnOtherDea = appOnBothDeas.InstanceAtIndex(1).Heartbeat()
				heartbeatOnOtherDea.DeaGuid = otherDea.DeaGuid

				store.SyncHeartbeats(dea.HeartbeatWith(
					dea.GetApp(0).InstanceAtIndex(1).Heartbeat(),
					dea.GetApp(1).InstanceAtIndex(3).Heartbeat(),
					heartbeatOnDea,
				))

				store.SyncHeartbeats(otherDea.HeartbeatWith(
					otherDea.GetApp(0).InstanceAtIndex(1).Heartbeat(),
					otherDea.GetApp(1).InstanceAtIndex(0).Heartbeat(),
					heartbeatOnOtherDea,
				))
			})

			Context("when the DEA heartbeats have not expired", func() {
				It("returns the instance heartbeats", func() {
					results, err := store.GetInstanceHeartbeats()
					Expect(err).NotTo(HaveOccurred())
					Expect(results).To(HaveLen(6))
					Expect(results).To(ContainElement(dea.GetApp(0).InstanceAtIndex(1).Heartbeat()))
					Expect(results).To(ContainElement(dea.GetApp(1).InstanceAtIndex(3).Heartbeat()))
					Expect(results).To(ContainElement(heartbeatOnDea))
					Expect(results).To(ContainElement(otherDea.GetApp(0).InstanceAtIndex(1).Heartbeat()))
					Expect(results).To(ContainElement(otherDea.GetApp(1).InstanceAtIndex(0).Heartbeat()))
					Expect(results).To(ContainElement(heartbeatOnOtherDea))
				})
			})

			Context("when a DEA heartbeat has expired", func() {
				BeforeEach(func() {
					storeAdapter.Delete("/hm/v1/dea-presence/" + dea.DeaGuid)
				})

				It("does not return any expired instance heartbeats", func() {
					results, err := store.GetInstanceHeartbeats()
					Expect(err).NotTo(HaveOccurred())
					Expect(results).To(HaveLen(3))
					Expect(results).To(ContainElement(otherDea.GetApp(0).InstanceAtIndex(1).Heartbeat()))
					Expect(results).To(ContainElement(otherDea.GetApp(1).InstanceAtIndex(0).Heartbeat()))
					Expect(results).To(ContainElement(heartbeatOnOtherDea))

					//we fetch twice to ensure that nothing is incorrectly deleted
					results, err = store.GetInstanceHeartbeats()
					Expect(err).NotTo(HaveOccurred())
					Expect(results).To(HaveLen(3))
					Expect(results).To(ContainElement(otherDea.GetApp(0).InstanceAtIndex(1).Heartbeat()))
					Expect(results).To(ContainElement(otherDea.GetApp(1).InstanceAtIndex(0).Heartbeat()))
					Expect(results).To(ContainElement(heartbeatOnOtherDea))
				})

				It("removes expired instance heartbeats from the store", func() {
					_, err := storeAdapter.Get("/hm/v1/apps/actual/" + store.AppKey(dea.GetApp(0).AppGuid, dea.GetApp(0).AppVersion) + "/" + dea.GetApp(0).InstanceAtIndex(1).Heartbeat().StoreKey())
					Expect(err).NotTo(HaveOccurred())
					_, err = storeAdapter.Get("/hm/v1/apps/actual/" + store.AppKey(dea.GetApp(1).AppGuid, dea.GetApp(1).AppVersion) + "/" + dea.GetApp(1).InstanceAtIndex(3).Heartbeat().StoreKey())
					Expect(err).NotTo(HaveOccurred())

					_, err = store.GetInstanceHeartbeats()
					Expect(err).NotTo(HaveOccurred())

					_, err = storeAdapter.Get("/hm/v1/apps/actual/" + store.AppKey(dea.GetApp(0).AppGuid, dea.GetApp(0).AppVersion) + "/" + dea.GetApp(0).InstanceAtIndex(1).Heartbeat().StoreKey())
					Expect(err).To(Equal(storeadapter.ErrorKeyNotFound))
					_, err = storeAdapter.Get("/hm/v1/apps/actual/" + store.AppKey(dea.GetApp(1).AppGuid, dea.GetApp(1).AppVersion) + "/" + dea.GetApp(1).InstanceAtIndex(3).Heartbeat().StoreKey())
					Expect(err).To(Equal(storeadapter.ErrorKeyNotFound))
				})

				Context("if it fails to remove them", func() {
					It("should soldier on", func() {
						resultChan := make(chan []models.InstanceHeartbeat, 2)
						errChan := make(chan error, 2)
						go func() {
							results, err := store.GetInstanceHeartbeats()
							resultChan <- results
							errChan <- err
						}()

						go func() {
							results, err := store.GetInstanceHeartbeats()
							resultChan <- results
							errChan <- err
						}()

						Expect(<-resultChan).To(HaveLen(3))
						Expect(<-resultChan).To(HaveLen(3))
						Expect(<-errChan).NotTo(HaveOccurred())
						Expect(<-errChan).NotTo(HaveOccurred())
					})
				})
			})
		})
	})

	Describe("Fetching actual state for a specific app guid & version", func() {
		var app appfixture.AppFixture
		BeforeEach(func() {
			app = appfixture.NewAppFixture()
		})

		Context("when there is none saved", func() {
			It("comes back empty", func() {
				results, err := store.GetInstanceHeartbeatsForApp(app.AppGuid, app.AppVersion)
				Expect(err).NotTo(HaveOccurred())
				Expect(results).To(BeEmpty())
			})
		})

		Context("when there is actual state saved", func() {
			var heartbeatA, heartbeatB models.InstanceHeartbeat

			BeforeEach(func() {
				heartbeatA = app.InstanceAtIndex(0).Heartbeat()
				heartbeatA.DeaGuid = "A"

				store.SyncHeartbeats(&models.Heartbeat{
					DeaGuid: "A",
					InstanceHeartbeats: []models.InstanceHeartbeat{
						heartbeatA,
					},
				})

				heartbeatB = app.InstanceAtIndex(1).Heartbeat()
				heartbeatB.DeaGuid = "B"

				store.SyncHeartbeats(&models.Heartbeat{
					DeaGuid: "B",
					InstanceHeartbeats: []models.InstanceHeartbeat{
						heartbeatB,
					},
				})
			})

			Context("when the corresponding DEA heartbeat has not expired", func() {
				It("returns the instance heartbeats", func() {
					results, err := store.GetInstanceHeartbeatsForApp(app.AppGuid, app.AppVersion)
					Expect(err).NotTo(HaveOccurred())
					Expect(results).To(HaveLen(2))
					Expect(results).To(ContainElement(heartbeatA))
					Expect(results).To(ContainElement(heartbeatB))
				})
			})

			Context("when an associated DEA heartbeat has expired", func() {
				BeforeEach(func() {
					storeAdapter.Delete("/hm/v1/dea-presence/A")
				})

				It("does not return any expired instance heartbeats", func() {
					results, err := store.GetInstanceHeartbeatsForApp(app.AppGuid, app.AppVersion)
					Expect(err).NotTo(HaveOccurred())
					Expect(results).To(HaveLen(1))
					Expect(results).To(ContainElement(heartbeatB))
				})

				It("removes expired instance heartbeats from the store", func() {
					_, err := storeAdapter.Get("/hm/v1/apps/actual/" + store.AppKey(app.AppGuid, app.AppVersion) + "/" + heartbeatA.StoreKey())
					Expect(err).NotTo(HaveOccurred())

					_, err = store.GetInstanceHeartbeatsForApp(app.AppGuid, app.AppVersion)
					Expect(err).NotTo(HaveOccurred())

					_, err = storeAdapter.Get("/hm/v1/apps/actual/" + store.AppKey(app.AppGuid, app.AppVersion) + "/" + heartbeatA.StoreKey())
					Expect(err).To(Equal(storeadapter.ErrorKeyNotFound))
				})

				Context("if it fails to remove them", func() {
					It("should soldier on", func() {
						resultChan := make(chan []models.InstanceHeartbeat, 2)
						errChan := make(chan error, 2)
						go func() {
							results, err := store.GetInstanceHeartbeatsForApp(app.AppGuid, app.AppVersion)
							resultChan <- results
							errChan <- err
						}()

						go func() {
							results, err := store.GetInstanceHeartbeatsForApp(app.AppGuid, app.AppVersion)
							resultChan <- results
							errChan <- err
						}()

						Expect(<-resultChan).To(HaveLen(1))
						Expect(<-resultChan).To(HaveLen(1))
						Expect(<-errChan).NotTo(HaveOccurred())
						Expect(<-errChan).NotTo(HaveOccurred())
					})
				})
			})

			Context("when all the DEA heartbeats have expired", func() {
				BeforeEach(func() {
					storeAdapter.Delete("/hm/v1/dea-presence/A", "/hm/v1/dea-presence/B")
				})

				It("does not return any instance heartbeats", func() {
					results, err := store.GetInstanceHeartbeatsForApp(app.AppGuid, app.AppVersion)
					Expect(err).NotTo(HaveOccurred())
					Expect(results).NotTo(BeNil())
					Expect(results).To(HaveLen(0))
				})
			})
		})
	})
})
