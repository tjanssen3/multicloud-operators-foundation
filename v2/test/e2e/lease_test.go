package e2e

import (
	"context"
	"fmt"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	addonv1alpha1 "github.com/open-cluster-management/api/addon/v1alpha1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	podNamespace = "open-cluster-management-agent"
)

var _ = ginkgo.Describe("Testing Lease", func() {
	ginkgo.Context("Get Lease", func() {
		ginkgo.It("should get/update lease successfully in cluster", func() {
			var firstLeaseTime *metav1.MicroTime
			// Creat managedclusteraddon apis
			var addon = &addonv1alpha1.ManagedClusterAddOn{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "work-manager",
					Namespace: managedClusterName,
				},
				Spec: addonv1alpha1.ManagedClusterAddOnSpec{
					InstallNamespace: podNamespace,
				},
			}
			_, err := addonClient.AddonV1alpha1().ManagedClusterAddOns(managedClusterName).Create(context.Background(), addon, metav1.CreateOptions{})
			gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
			gomega.Eventually(func() error {
				lease, err := kubeClient.CoordinationV1().Leases(podNamespace).Get(context.Background(), "work-manager", metav1.GetOptions{})
				if err != nil {
					return err
				}
				firstLeaseTime = lease.Spec.RenewTime
				return nil
			}, eventuallyTimeout, eventuallyInterval).ShouldNot(gomega.HaveOccurred())
			gomega.Eventually(func() error {
				updatedLease, err := kubeClient.CoordinationV1().Leases(podNamespace).Get(context.Background(), "work-manager", metav1.GetOptions{})
				if err != nil {
					return err
				}
				updatedLeaseTime := updatedLease.Spec.RenewTime
				if updatedLeaseTime.Equal(firstLeaseTime) {
					return fmt.Errorf("lease should be updated")
				}
				return nil
			}, eventuallyTimeout, eventuallyInterval).ShouldNot(gomega.HaveOccurred())
			// Ensure the addon status is correct
			gomega.Eventually(func() bool {
				addon, err := addonClient.AddonV1alpha1().ManagedClusterAddOns(managedClusterName).Get(context.Background(), "work-manager", metav1.GetOptions{})
				if err != nil {
					return false
				}
				return meta.IsStatusConditionTrue(addon.Status.Conditions, "Available")
			}, eventuallyTimeout, eventuallyInterval).Should(gomega.BeTrue())
		})
	})
})
