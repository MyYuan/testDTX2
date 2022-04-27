#!/bin/bash

# Copyright (c) 2021 PaddlePaddle Authors. All Rights Reserved.
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Script to start Xdb and its dependent services with docker-compose.
# Usage: ./network_up.sh {start|stop|restart}

# Export system environment variables
export COMPOSE_PROJECT_NAME=blockchain
# Default System channel name of fabric network
export SYS_CHANNEL=byfn-sys-channel

# Directory of temporary files generated by script
# 脚本启动生成的临时文件目录
TMP_CONF_NAME="testdatatmp"

# Xdb contract name
# 合约名称
CONTRACT_NAME=xdb01
# Network type, support xchain and fabric
# 网络类型, 支持xchain和fabric, default xchain
BLOCKCHAIN_TYPE=${2-"xchain"}

# Configuration items required for xchain network startup
# User's blockchain account address, used invoke contract 
# 用户安装合约所使用的区块链账户地址
ADDRESS_PATH=../$TMP_CONF_NAME/blockchain/xchain/user
CONTRACT_ACCOUNT=1111111111111111
TRANSFER_AMOUNT=110009887797

# Configuration items required for fabric network startup
# Timeout duration - the duration the CLI should wait for a response from
# Another container before giving up
CLI_TIMEOUT=10
# Default for delay between commands
CLI_DELAY=3
# Use golang as the default language for chaincode
LANGUAGE=golang
# default image tag
IMAGETAG="1.4.8"
# Default channel_name
CHANNEL_NAME="mychannel"

function start() {
  # 1. Standardize Conf
  # 1. 标准化xchain网络配置文件
  standardizeConf

  # 2. Start blockchain network, support xchain or fabric
  # 2. 启动区块链网络, 支持xchain或fabric
  if [ $BLOCKCHAIN_TYPE = "fabric" ];then
    startFabric
  else
    startXchain
  fi
  
  # 3. Start decentralized storage network
  # 3. 启动去中心化存储网络
  startXdb
}

# startXchain start xchain network with docker compose
# 通过docker-compose启动区块链网络
function startXchain() {
  # 1. Start xchain network
  # 1. 启动区块链网络
  print_blue "==========> Xchain network start ..."
  docker-compose -f ../$TMP_CONF_NAME/blockchain/xchain/docker-compose.yml up -d
  sleep 6

  xchainContainers="xchain1.node.com xchain2.node.com xchain3.node.com"
  checkContainerStatus "$xchainContainers" "Xchain network"
  print_green "==========> Xchain network starts successfully !"

  # 2. Compile and install Xdb contract
  # 2. 编译和安装智能合约
  installXchainContract
  sleep 5
}

# Generate standard config.toml file
# 生成标准配置文件
function standardizeConf() {
  # Create a temporary file directory required for xdb network startup
  rm -rf ../$TMP_CONF_NAME && mkdir ../$TMP_CONF_NAME && cp -r ../testdata/* ../$TMP_CONF_NAME/
  # Replace user parameters into configuration file
  sampleConfigFiles=`find $(dirname ${PWD})/$TMP_CONF_NAME -name "config.toml.sample"`
  MNEMONIC=`cat $ADDRESS_PATH/mnemonic`
  for file in $sampleConfigFiles
  do
    conf=${file%.sample}
    eval "cat <<EOF
$(< $file)
EOF
"  > $conf
  done
}

# startFabric start fabric network with docker compose
# 启动fabric网络
function startFabric() {
  # 1. Start fabric network tool container, for certificate generation
  # 1. 启动fabric网络工具容器, 用于证书生成
  docker run -itd --name fabric-network-start-tool \
  -v $(dirname ${PWD})/$TMP_CONF_NAME/blockchain/fabric/conf:/home/conf/ \
  registry.baidubce.com/paddledtx/fabric-network-start-tool:1.0 
  docker exec fabric-network-start-tool sh -c "
  cd conf \
  && ../cryptogen generate --config=./crypto-config.yaml"
  if [ $? -ne 0 ]; then
    print_red "==========> ERROR !!!! Unable to generate the x509 certificates ..."
    exit 1
  fi
  
  # 2. Generate orderer genesis block, channel configuration transaction and
  # anchor peer update transactions
  # 2. 生成创世区块
  docker exec fabric-network-start-tool sh -c "
  cd conf && mkdir channel-artifacts \
  && ../configtxgen -profile TwoOrgsOrdererGenesis -channelID $SYS_CHANNEL -outputBlock ./channel-artifacts/genesis.block \
  && ../configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME \
  && ../configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP \
  && ../configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org2MSP
  "
  if [ $? -ne 0 ]; then
    print_red "==========> ERROR !!!! Unable to generate orderer genesis block ..."
    exit 1
  fi

  # 3. Copy the generated certificate and Genesis block file to the fabric directory
  # 3. 将生成的证书和创世区块文件拷贝到fabric目录
  docker cp fabric-network-start-tool:/home/conf/crypto-config ../$TMP_CONF_NAME/blockchain/fabric/
  docker cp fabric-network-start-tool:/home/conf/channel-artifacts ../$TMP_CONF_NAME/blockchain/fabric/
  # Clear tool build container
  # 清除工具生成容器
  clearToolContainer
  
  # 4. start fabric network
  # 4. 启动fabric网络
  # Generate dependent packages required for chain code installation
  generateVendor
  IMAGE_TAG=$IMAGETAG docker-compose -f ../$TMP_CONF_NAME/blockchain/fabric/docker-compose.yaml up -d 2>&1
  docker ps -a
  if [ $? -ne 0 ]; then
    print_red "==========> ERROR !!!! Unable to start fabric network ..."
    exit 1
  fi

  # 5. 通过fabric script安装链码
  # 5. Installation chain code by fabric script
  VERBOSE=false
  NO_CHAINCODE=false
  docker exec cli scripts/script.sh $CHANNEL_NAME $CLI_DELAY $LANGUAGE $CLI_TIMEOUT $VERBOSE $NO_CHAINCODE $CONTRACT_NAME
  if [ $? -ne 0 ]; then
    print_red "==========> ERROR !!!! Install chaincode error ..."
    exit 1
  fi
  sleep 3

  # Copy fabric sdk yaml file into xdb configuration directory
  # 拷贝sdk yaml文件到xdb配置目录下
  CONFFILE_LIST="dataowner storage1 storage2 storage3"
  for CONF_PATH in $CONFFILE_LIST
  do
    mkdir -p ../$TMP_CONF_NAME/xdb/$CONF_PATH/conf/fabric/
    cp -f ../$TMP_CONF_NAME/blockchain/fabric/conf/config.yaml ../$TMP_CONF_NAME/xdb/$CONF_PATH/conf/fabric/
    cp -rf ../$TMP_CONF_NAME/blockchain/fabric/crypto-config ../$TMP_CONF_NAME/xdb/$CONF_PATH/conf/fabric/
  done
}

# startXdb start Xdb network with docker compose
# 通过docker-compose启动去中心化存储网络
function startXdb() {
  print_blue "==========> Xdb network start ..."
  docker-compose -p xdb -f ../$TMP_CONF_NAME/xdb/docker-compose.yml up -d
  sleep 6

  xdbContainers="dataowner.node.com storage1.node.com storage2.node.com storage3.node.com"
  checkContainerStatus "$xdbContainers" "Xdb network"
  print_green "==========> Xdb network starts successfully !"
}

# checkContainerStatus check container status
# 判断容器启动状态
function checkContainerStatus() {
  for containerName in $1
  do
    exist=`docker inspect --format '{{.State.Running}}' ${containerName}`
    if [ "${exist}" != "true" ]; then
      print_red "==========> $2 start error ..."
      exit 1
    fi
  done
}

function generateVendor() {
  cp -f ../blockchain/fabric/chaincode/* ../$TMP_CONF_NAME/blockchain/fabric/chaincode/
  docker run -it --rm \
      -v $(dirname ${PWD}):/workspace \
      -v ~/.ssh:/root/.ssh \
      -w /workspace \
      -e GONOSUMDB=* \
      -e GOPROXY=https://goproxy.cn \
      -e GO111MODULE=on \
      golang:1.13.4 sh -c "cd ./$TMP_CONF_NAME/blockchain/fabric/chaincode/ && go mod vendor"
}

# compileXchainContract compiling xchain contract through golang: 1.13.4 container
# 通过golang:1.13.4容器编译xchain智能合约
function compileXchainContract() {
  docker run -it --rm \
      -v $(dirname ${PWD}):/workspace \
      -v ~/.ssh:/root/.ssh \
      -w /workspace \
      -e GONOSUMDB=* \
      -e GOPROXY=https://goproxy.cn \
      -e GO111MODULE=on \
      golang:1.13.4 sh -c "go build -o ./$TMP_CONF_NAME/blockchain/xchain/contract/$CONTRACT_NAME ./blockchain/xchain/contract"
  # Copy contract file to xchain1 container
  # 将本地合约编译结果拷贝到区块链节点1容器中
  docker cp ../$TMP_CONF_NAME/blockchain/xchain/contract/$CONTRACT_NAME xchain1.node.com:/home/work/xchain/$CONTRACT_NAME
}

# installXchainContract install xchain contract by xchain contractAccount
# 通过合约账户安装xchain智能合约
function installXchainContract() {
  print_blue "==========> Install $CONTRACT_NAME contract start ..."
  compileXchainContract

  address=`cat $ADDRESS_PATH/address`
  echo "user address $address ..."

  # Transfer token from miner's account to user's account
  # 从矿工账户给用户账户转移token
  transferAddressResult=`docker exec -it xchain1.node.com sh -c "
    ./xchain-cli transfer --to $address --amount $TRANSFER_AMOUNT --keys ./data/keys --host xchain1.node.com:37101 
  "`
  checkOperateResult "$transferAddressResult"
  # Get required fee to create the contract account
  # 获取创建合约账户所需要的fee
  accountFee=`docker exec -it xchain1.node.com sh -c "
    ./xchain-cli account new --account $CONTRACT_ACCOUNT --host xchain1.node.com:37101 --keys ./user
  " | grep "The gas you" | awk -F: '{print $2}' | awk 'BEGIN{RS="\r";ORS="";}{print $0}' | awk '$1=$1'`

  # Create xchain contract account
  # 创建合约账户
  contractAccountResult=`docker exec -it xchain1.node.com sh -c "
    ./xchain-cli account new --account $CONTRACT_ACCOUNT --host xchain1.node.com:37101 --fee $accountFee --keys ./user
  "`
  checkOperateResult "$contractAccountResult"

  print_blue "create contract account successfully, $contractAccountResult ..."

  # Transfer token from miner's account to contract account
  # 从矿工账户给用户创建的合约账户转移token
  transferAccountResult=`docker exec -it xchain1.node.com sh -c "
    ./xchain-cli transfer --to XC${CONTRACT_ACCOUNT}@xuper --amount $TRANSFER_AMOUNT --keys ./data/keys  --host xchain1.node.com:37101 
  "`
  checkOperateResult "$transferAccountResult"

  print_blue "transfer amount to contract account successfully, $transferAccountResult ..."

  # Ensure that the contract account is created successfully, and then install the contract
  # 确保合约账户创建成功后，再进行合约安装
  sleep 5
  # Get required fee for deploy contract
  # 获取安装合约所需要的fee
  contractFee=`docker exec -it xchain1.node.com sh -c "
    ./xchain-cli native deploy --account XC${CONTRACT_ACCOUNT}@xuper --host xchain1.node.com:37101 --runtime go -a '{\"creator\":\"XC${CONTRACT_ACCOUNT}@xuper\"}' --cname $CONTRACT_NAME ./$CONTRACT_NAME --keys ./user
  " | grep "The gas you" | awk -F: '{print $2}' | awk 'BEGIN{RS="\r";ORS="";}{print $0}' | awk '$1=$1'`

  # Install contract
  # 安装智能合约
  installResult=`docker exec -it xchain1.node.com sh -c "
    ./xchain-cli native deploy --account XC${CONTRACT_ACCOUNT}@xuper --fee $contractFee  --host xchain1.node.com:37101 --runtime go -a '{\"creator\":\"XC${CONTRACT_ACCOUNT}@xuper\"}' --cname $CONTRACT_NAME ./$CONTRACT_NAME --keys ./user
  "`
  checkOperateResult "$installResult"
  print_green "==========> Install $CONTRACT_NAME contract successfully ! "
}

function stop() {
  # Stop xdb network
  # 停止去中心化存储网络
  print_blue "==========> Stop xdb network ..."
  docker-compose -f ../$TMP_CONF_NAME/xdb/docker-compose.yml down --remove-orphans
  docker rm -f $(docker ps -a | grep "paddledtx/*" | awk "{print \$1}")

  # Stop blockchain network
  # 停止区块链网络
  if [ $BLOCKCHAIN_TYPE = "fabric" ];then
    print_blue "==========> Stop fabric network ..."
    IMAGE_TAG=$IMAGETAG docker-compose -f ../$TMP_CONF_NAME/blockchain/fabric/docker-compose.yaml down
    docker rm -f $(docker ps -a | grep "$CONTRACT_NAME/*" | awk "{print \$1}")
    clearToolContainer
    # Cleanup chaincode images
    # 清除链码容器镜像
    removeChaincodeImages
  else 
    print_blue "==========> Stop xchain network ..."
    docker-compose -f ../$TMP_CONF_NAME/blockchain/xchain/docker-compose.yml down
  fi
  # Delete temporary profiles by container
  # 通过容器方式删除临时配置文件, 避免因为文件权限问题导致的删除失败
  docker run -it --rm \
    -v $(dirname ${PWD}):/workspace \
    golang:1.13.4 sh -c "cd /workspace && rm -rf $TMP_CONF_NAME"
  docker volume rm $(docker volume ls -qf dangling=true)
  print_green "==========> Xdb network stopped !"
}

# Obtain CONTAINER_IDS and remove them
# TODO Might want to make this optional - could clear other containers
function clearToolContainer() {
  CONTAINERID=`docker ps -a | awk '($2 ~ /fabric-network-start-tool*/) {print $1}'`
  if [ -z "$CONTAINERID" -o "$CONTAINERID" == " " ]; then
    print_blue "==========> Fabric-network-start-tool containers available for deletion ..."
  else
    docker rm -f $CONTAINERID
  fi
}

# Delete any images that were generated as a part of this setup
# specifically the following images are often left behind:
# 清除链码容器镜像
function removeChaincodeImages() {
  DOCKER_IMAGE_IDS=$(docker images | grep "$CONTRACT_NAME/*" | awk "{print \$1}")
  if [ -z "$DOCKER_IMAGE_IDS" -o "$DOCKER_IMAGE_IDS" == " " ]; then
    print_blue "==========> No images available for deletion ..."
  else
    docker rmi -f $DOCKER_IMAGE_IDS
  fi
}

# Judge whether the command execution result is correct
# 判断命令执行结果
function checkOperateResult() {
  ERR_MESSAGE=`echo "$1" | grep "Error\|Fail\|err\|Failed"`
  if [ "$ERR_MESSAGE" ]; then
    print_red "==========> ERROR !!!! start xdb network failed: $1"
    exit 1
  fi
}

RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m'

function print_blue() {
  printf "${BLUE}%s${NC}\n" "$1"
}

function print_green() {
  printf "${GREEN}%s${NC}\n" "$1"
}

function print_red() {
  printf "${RED}%s${NC}\n" "$1"
}

action=$1
case $action in
start)
  start $@
  ;;
stop)
  stop $@
  ;;
restart)
  stop $@
  start $@
  ;;
*)
  echo "Usage: $0 {start|stop|restart}"
  exit 1
  ;;
esac
