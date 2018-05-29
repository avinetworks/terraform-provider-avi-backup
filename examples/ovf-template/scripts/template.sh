#!/usr/bin/env bash

set -e

TEMPLATE_URL="/home/vinay/git/terraform/src/github.com/avinetworks/terraform-provider-avi/examples/ovf-template/scripts/controller.ova"
TEMPLATE_DC="$(grep "datacenter" "${TF_DIR}/terraform.tfvars" | awk '{print $3}' | tr -d \'\"\' )"
TEMPLATE_DS="$(grep "datastore_name" "${TF_DIR}/terraform.tfvars" | awk '{print $3}' | tr -d \'\"\' )"
TEMPLATE_POOL="$(grep "resource_pool" "${TF_DIR}/terraform.tfvars" | awk '{print $3}' | tr -d \'\"\' )"
TEMPLATE_NAME="$(grep "template_name" "${TF_DIR}/terraform.tfvars" | awk '{print $3}' | tr -d \'\"\' )"
TEMPLATE_NETWORK_NAME="$(grep "network_name" "${TF_DIR}/terraform.tfvars" | awk '{print $3}' | tr -d \'\"\' )"
TEMPLATE_CONTROLLER_IP="$(grep "controller_ip" "${TF_DIR}/terraform.tfvars" | awk '{print $3}' | tr -d \'\"\' )"

echo "Uploading template file ${TEMPLATE_NAME} on datastore ${TEMPLATE_DS}"

#EXISTING_TEMPLATE="$(govc find -type m -name "${TEMPLATE_NAME}" | head -n 1)"
#if [ -n "${EXISTING_TEMPLATE}" ]; then
#  govc object.rename "${EXISTING_TEMPLATE}" "${TEMPLATE_NAME}_archived_$(date +%s)"
#fi

echo "ovftool --powerOn --acceptAllEulas --skipManifestCheck --allowExtraConfig --X:waitForIp -n="${TEMPLATE_NAME}" -ds="${TEMPLATE_DS}" --net:"Management"="${TEMPLATE_NETWORK_NAME}" --prop:"avi.mgmt-ip.CONTROLLER"="${TEMPLATE_CONTROLLER_IP}" "${TEMPLATE_URL}" vi://root:vmware@10.10.2.10/10GTest/host/Arista"

ovftool --overwrite --acceptAllEulas --skipManifestCheck --allowExtraConfig --X:waitForIp --X:injectOvfEnv -n="${TEMPLATE_NAME}" -ds="${TEMPLATE_DS}" --prop:"avi.mgmt-ip.CONTROLLER"="${TEMPLATE_CONTROLLER_IP}" --prop:"avi.mgmt-mask.CONTROLLER"="255.255.254.0" --prop:"avi.default-gw.CONTROLLER"="10.10.28.1" --net:"Management"="${TEMPLATE_NETWORK_NAME}" "${TEMPLATE_URL}" vi://root:vmware@10.10.2.10/10GTest/host/Arista 
#govc import.ova -dc="${TEMPLATE_DC}" -ds="${TEMPLATE_DS}" -pool="${TEMPLATE_POOL}" -name="${TEMPLATE_NAME}" "${TEMPLATE_URL}"
#govc snapshot.create -dc="${TEMPLATE_DC}" -vm="${TEMPLATE_NAME}" clone-root
#govc vm.markastemplate -dc="${TEMPLATE_DC}" "${TEMPLATE_NAME}"
