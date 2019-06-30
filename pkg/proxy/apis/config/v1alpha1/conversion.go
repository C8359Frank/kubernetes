package v1alpha1

import (
	"k8s.io/apimachinery/pkg/conversion"
	configv1alpha1 "k8s.io/component-base/config/v1alpha1"
	v1alpha1 "k8s.io/kube-proxy/config/v1alpha1"
	"k8s.io/kubernetes/pkg/proxy/apis/config"
	"unsafe"
)

func Convert_v1alpha1_KubeProxyConfiguration_To_config_KubeProxyConfiguration(in *v1alpha1.KubeProxyConfiguration, out *config.KubeProxyConfiguration, s conversion.Scope) error {
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	out.BindAddress = []string{in.BindAddress}
	out.HealthzBindAddress = in.HealthzBindAddress
	out.MetricsBindAddress = in.MetricsBindAddress
	out.EnableProfiling = in.EnableProfiling
	out.ClusterCIDR = []string{in.ClusterCIDR}
	out.HostnameOverride = in.HostnameOverride
	if err := configv1alpha1.Convert_v1alpha1_ClientConnectionConfiguration_To_config_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_KubeProxyIPTablesConfiguration_To_config_KubeProxyIPTablesConfiguration(&in.IPTables, &out.IPTables, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_KubeProxyIPVSConfiguration_To_config_KubeProxyIPVSConfiguration(&in.IPVS, &out.IPVS, s); err != nil {
		return err
	}
	out.OOMScoreAdj = (*int32)(unsafe.Pointer(in.OOMScoreAdj))
	out.Mode = config.ProxyMode(in.Mode)
	out.PortRange = in.PortRange
	out.UDPIdleTimeout = in.UDPIdleTimeout
	if err := Convert_v1alpha1_KubeProxyConntrackConfiguration_To_config_KubeProxyConntrackConfiguration(&in.Conntrack, &out.Conntrack, s); err != nil {
		return err
	}
	out.ConfigSyncPeriod = in.ConfigSyncPeriod
	out.NodePortAddresses = *(*[]string)(unsafe.Pointer(&in.NodePortAddresses))
	if err := Convert_v1alpha1_KubeProxyWinkernelConfiguration_To_config_KubeProxyWinkernelConfiguration(&in.Winkernel, &out.Winkernel, s); err != nil {
		return err
	}
	return nil
}
