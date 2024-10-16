/*
SPDX-License-Identifier: Apache-2.0

Copyright Contributors to the Submariner project.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kubeproxy

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/submariner-io/admiral/pkg/log"
	"github.com/submariner-io/submariner/pkg/iptables"
	"github.com/submariner-io/submariner/pkg/routeagent_driver/constants"
	iptcommon "github.com/submariner-io/submariner/pkg/routeagent_driver/iptables"
	"k8s.io/klog"
)

func (kp *SyncHandler) createIPTableChains() error {
	ipt, err := iptables.New()
	if err != nil {
		return errors.Wrap(err, "error initializing iptables")
	}

	if err := iptcommon.InitSubmarinerPostRoutingChain(ipt); err != nil {
		return errors.Wrap(err, "error initializing POST routing chain")
	}

	klog.V(log.DEBUG).Infof("Install/ensure SUBMARINER-INPUT chain exists")

	if err = iptables.CreateChainIfNotExists(ipt, "filter", "SUBMARINER-INPUT"); err != nil {
		return errors.Wrap(err, "unable to create SUBMARINER-INPUT chain in iptables")
	}

	forwardToSubInputRuleSpec := []string{"-p", "udp", "-m", "udp", "-j", "SUBMARINER-INPUT"}
	if err = ipt.AppendUnique("filter", "INPUT", forwardToSubInputRuleSpec...); err != nil {
		return errors.Wrapf(err, "unable to append iptables rule %q", strings.Join(forwardToSubInputRuleSpec, " "))
	}

	klog.V(log.DEBUG).Infof("Allow VxLAN incoming traffic in SUBMARINER-INPUT Chain")

	ruleSpec := []string{"-p", "udp", "-m", "udp", "--dport", strconv.Itoa(VxLANPort), "-j", "ACCEPT"}

	if err = ipt.AppendUnique("filter", "SUBMARINER-INPUT", ruleSpec...); err != nil {
		return errors.Wrapf(err, "unable to append iptables rule %q", strings.Join(ruleSpec, " "))
	}

	klog.V(log.DEBUG).Infof("Insert rule to allow traffic over %s interface in FORWARDing Chain", VxLANIface)

	ruleSpec = []string{"-o", VxLANIface, "-j", "ACCEPT"}

	if err = iptables.PrependUnique(ipt, "filter", "FORWARD", ruleSpec); err != nil {
		return errors.Wrap(err, "unable to insert iptable rule in filter table to allow vxlan traffic")
	}

	if kp.cniIface != nil {
		// Program rules to support communication from HostNetwork to remoteCluster
		sourceAddress := strconv.Itoa(VxLANVTepNetworkPrefix) + ".0.0.0/8"
		ruleSpec = []string{"-s", sourceAddress, "-o", VxLANIface, "-j", "SNAT", "--to", kp.cniIface.IPAddress}
		klog.V(log.DEBUG).Infof("Installing rule for host network to remote cluster communication: %s", strings.Join(ruleSpec, " "))

		if err = ipt.AppendUnique("nat", constants.SmPostRoutingChain, ruleSpec...); err != nil {
			return errors.Wrapf(err, "error appending iptables rule %q", strings.Join(ruleSpec, " "))
		}
	}

	return nil
}

func (kp *SyncHandler) updateIptableRulesForInterClusterTraffic(inputCidrBlocks []string, operation Operation) {
	for _, inputCidrBlock := range inputCidrBlocks {
		if operation == Add {
			err := kp.programIptableRulesForInterClusterTraffic(inputCidrBlock)
			if err != nil {
				klog.Errorf("Failed to program iptable rule. %v", err)
			}
		} else if operation == Delete {
			// TODO: Handle this use-case
			klog.Warning("Handle the delete use-case")
		}
	}
}

func (kp *SyncHandler) programIptableRulesForInterClusterTraffic(remoteCidrBlock string) error {
	ipt, err := iptables.New()
	if err != nil {
		return errors.Wrap(err, "error initializing iptables")
	}

	for _, localClusterCidr := range kp.localClusterCidr {
		ruleSpec := []string{"-s", localClusterCidr, "-d", remoteCidrBlock, "-j", "ACCEPT"}
		klog.V(log.DEBUG).Infof("Installing iptables rule for outgoing traffic: %s", strings.Join(ruleSpec, " "))

		if err = ipt.AppendUnique("nat", constants.SmPostRoutingChain, ruleSpec...); err != nil {
			return errors.Wrapf(err, "error appending iptables rule %q", strings.Join(ruleSpec, " "))
		}

		// TODO: revisit, we only have to program rules to allow traffic from the podCidr
		ruleSpec = []string{"-s", remoteCidrBlock, "-d", localClusterCidr, "-j", "ACCEPT"}
		klog.V(log.DEBUG).Infof("Installing iptables rule for incoming traffic: %s", strings.Join(ruleSpec, " "))

		if err = ipt.AppendUnique("nat", constants.SmPostRoutingChain, ruleSpec...); err != nil {
			return errors.Wrapf(err, "error appending iptables rule %q", strings.Join(ruleSpec, " "))
		}
	}

	return nil
}
