/*
 * Copyright (c) 2024 OceanBase.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ob

import (
	"github.com/oceanbase/obshell/agent/constant"
	"github.com/oceanbase/obshell/agent/engine/task"
	"github.com/oceanbase/obshell/agent/service/agent"
	"github.com/oceanbase/obshell/agent/service/obcluster"
	taskservice "github.com/oceanbase/obshell/agent/service/task"
)

const (

	// start check item
	CHECK_PORT = "checkPort"
	CHECK_DIR  = "checkDir"
	CHECK_MEM  = "checkMem"
	CHECK_DISK = "checkDisk"
	CHECK_NET  = "checkNet"
	CHECK_AIO  = "checkAio"
	CHECK_ULIM = "checkUlimit"

	// task context key
	PARAM_CONFIG     = "config"
	PARAM_DIRS       = "dirs"
	PARAM_PORTS      = "ports"
	PARAM_ROOT_PWD   = "rootPwd"
	PARAM_REMOTE_ID  = "remoteId"
	PARAM_ZONE_RS    = "zoneRS"
	PARAM_ZONE_ORDER = "zones"
	PARAM_UNRS       = "unRs" // PARAM_UNRS is a map, key is zone, value is servers that not in rs
	PARAM_DELETE_ALL = "deleteAll"
	// for scale out
	PARAM_IS_NEW_ZONE    = "isNewZone"
	PARAM_AGENT_INFO     = "agentInfo"
	PARAM_SCALE_OUT_UUID = "scaleOutUUID"

	PARAM_EXPECTED_STAGE         = "expectedStage"
	PARAM_MAIN_DAG_ID            = "mainDagId"
	PARAM_MAIN_AGENT             = "mainAgent"
	PARAM_SCOPE                  = "scope"
	PARAM_ALL_AGENTS             = "allAgents"
	PARAM_FORCE_PASS_DAG         = "forcePassDag"
	PARAM_START_OWN_OBSVR        = "startOwnObsvr"
	PARAM_EXPECT_MAIN_NEXT_STAGE = "expectedMainNextStage"
	PARAM_URI                    = "uri"
	// scale out
	PARAM_EXPECT_DEPLOY_NEXT_STAGE   = "expectedDeployNextStage"
	PARAM_EXPECT_START_NEXT_STAGE    = "expectedStartNextStage"
	PARAM_WAIT_DEPLOY_RETRY_STAGE    = "waitDeployRetryStage"
	PARAM_WAIT_START_RETRY_STAGE     = "waitStartRetryStage"
	PARAM_EXPECT_ROLLBACK_NEXT_STAGE = "expectedRollbackNextStage"
	PARAM_COORDINATE_DAG_ID          = "coordinateDagId"
	PARAM_COORDINATE_AGENT           = "coordinateAgent"
	PARAM_JOIN_MASTER_INFO           = "joinMasterInfo"

	// for upgrade
	PARAM_VERSION                = "version"
	PARAM_BUILD_NUMBER           = "buildNumber"
	PARAM_DISTRIBUTION           = "distribution"
	PARAM_RELEASE_DISTRIBUTION   = "releaseDistribution"
	PARAM_UPGRADE_DIR            = "upgradeDir"
	PARAM_TASK_TIME              = "taskTime"
	PARAM_ONLY_FOR_AGENT         = "onlyForAgent"
	PARAM_AGENT_UPGRADE_ROUTE    = "agentUpgradeRoute"
	PARAM_UPGRADE_CHECK_TASK_DIR = "upgradeCheckTaskDir"
	PARAM_SCRIPT_FILE            = "scriptFile"
	PARAM_OB_PARAMETERS          = "obParameters"
	PARAM_UPGRADE_ROUTE          = "upgradeRoute"
	PARAM_UPGRADE_ROUTE_INDEX    = "upgradeRouteIndex"
	PARAM_ZONE                   = "zone"
	PARAM_CLUSTER_NAME           = "cluster"
	PARAM_CLUSTER_ID             = "cluster_id"

	PARAM_TARGET_AGENT_BUILD_VERSION = "targetAgentBuildVersion"

	DATA_ALL_AGENT_DAG_MAP = "allAgentDagMap"
	DATA_SKIP_START_TASK   = "skipStartTask"

	DATA_SUB_DAG_NEED_EXEC_CMD = "subDagNeedExecCmd"
	DATA_SUB_DAG_INFO          = "subDagInfo"
	DATA_SUB_DAG_SUCCEED       = "subDagSucceed"
	DATA_STOPPED_ZONES         = "stoppedZones"
	DATA_STOPPED_SERVERS       = "stoppedServers"

	// for upgrade
	DATA_BACKUP_DIR = "backupDir"

	// scope
	SCOPE_GLOBAL = "GLOBAL"
	SCOPE_ZONE   = "ZONE"
	SCOPE_SERVER = "SERVER"

	// remote request retry times
	DEFAULT_REMOTE_REQUEST_RETRY_TIMES = 30

	// task name
	TASK_NAME_INTEGRATE_CONFIG                   = "Integrate config"
	TASK_NAME_DEPLOY                             = "Create observer workdir"
	TASK_NAME_DESTROY                            = "Destroy observer workdir"
	TASK_NAME_START                              = "Start observer"
	TASK_NAME_STOP                               = "Stop observer"
	TASK_NAME_BOOTSTRAP                          = "Cluster boostrap"
	TASK_NAME_MIGRATE_TABLE                      = "Migrate table"
	TASK_NAME_MODIFY_PWD                         = "Modify password"
	TASK_NAME_MIGRATE_DATA                       = "Migrate data"
	TASK_NAME_UPDATE_AGENT                       = "Update all agents"
	TASK_NAME_INITIALIZE_DATA                    = "Initialize data"
	TASK_NAME_AGENT_SYNC                         = "Synchronize agent from cluster"
	TASK_NAME_UPDATE_CONFIG                      = "Update cluster config"
	TASK_NAME_UPDATE_OB_CONFIG                   = "Update observer config"
	TASK_NAME_GET_CONN_FOR_EMERGENCY_START       = "Get connection of observer"
	TASK_CONVERT_FOLLOWER_TO_CLUSTER             = "Convert follower to cluster agent"
	TASK_CONVERT_MASTER_TO_CLUSTER               = "Convert master to cluster agent"
	TASK_START_PREPARATIONS                      = "Start preparations"
	TASK_CHECK_OB_PROC_AND_CONIFG                = "Check ob process and config"
	TASK_EXEC_START_OBSERVER_SQL                 = "Execute start observer sql"
	TASK_WAIT_FOR_TASK_TO_END                    = "Wait for task to end"
	TASK_CHECK_PYTHON_ENV                        = "Check the python environment"
	TASK_BACKUP_FOR_UPGRADE                      = "Backup for upgrade"
	TASK_INSTALL_NEW_OBSHELL                     = "Install new obshell"
	TASK_RESTART_OBSHELL                         = "Restart obshell"
	TASK_CREATE_UPGRADE_DIR                      = "Create upgrade dir"
	TASK_REMOVE_UPGRADE_CHECK_TASK_DIR           = "Remove upgrade check task dir"
	TASK_UPGRADE_POST_TABLE_MAINTAIN             = "Upgrade post table maintain"
	TASK_EXEC_UPGRADE_SCRIPT                     = "Execute upgrade script"
	TASK_EXEC_UPGRADE_CHECKER_SCRIPT             = "Execute upgrade checker script"
	TASK_EXEC_UPGRADE_PRE_SCRIPT                 = "Execute upgrade pre script"
	TASK_EXEC_UPGRADE_POST_SCRIPT                = "Execute upgrade post script"
	TASK_EXEC_UPGRADE_HEALTH_CHECKER_SCRIPT      = "Execute upgrade health checker script"
	TASK_EXEC_UPGRADE_ZONE_HEALTH_CHECKER_SCRIPT = "Execute upgrade zone health checker script"
	TASK_BACKUP_PARAMETERS                       = "Backup parameters"
	TASK_RESTORE_PARAMETERS                      = "Restore parameters"
	TASK_CHECK_ALL_REQUIRED_PKGS                 = "Check all required packages"
	TASK_GET_ALL_REQUIRED_PKGS                   = "Download all required packages"
	TASK_INSTALL_ALL_REQUIRED_PKGS               = "Unpack all required packages"
	TASK_REINSTALL_AND_RESTART_OBSERVER          = "Reinstall and restart observer"
	TASK_NAME_BE_SCALING                         = "Be scaling agent"
	TASK_NAME_WAIT_DEPLOY_RETRY                  = "Wait deploy retry"
	TASK_NAME_WAIT_START_RETRY                   = "Wait start retry"
	TASK_NAME_WATCH_DAG                          = "Watch cluster scale_out dag"
	TASK_NAME_SYNC_FROM_OB                       = "Sync from observer"
	TASK_NAME_INTEGRATE_SERVER_CONFIG            = "Integrate server config"
	TASK_NAME_WAIT_SCALING_READY                 = "Wait scale out ready"
	TASK_NAME_CREATE_LOCAL_SCALE_OUT_DAG         = "Create local scale out dag"
	TASK_NAME_WAIT_REMOTE_DEPLOY_FINISH          = "Wait remote deploy finish"
	TASK_NAME_WAIT_REMOTE_START_FINISH           = "Wait remote start finish"
	TASK_NAME_PREV_CHECK                         = "Prev check for scale_out"
	TASK_NAME_ADD_NEW_ZONE                       = "Add new zone  for scale_out"
	TASK_NAME_START_NEW_ZONE                     = "Start new zone  for scale_out"
	TASK_NAME_ADD_SERVER                         = "Add server for scale_out"
	TASK_NAME_ADD_AGENT                          = "Add agent  for scale_out"
	TASK_NAME_FINISH                             = "Check cluster scale_out whether finished"
	TASK_NAME_MINOR_FREEZE                       = "Minor freeze before stop server"

	// dag name
	DAG_EMERGENCY_START           = "Start local observer"
	DAG_EMERGENCY_STOP            = "Stop local observer"
	DAG_INIT_CLUSTER              = "Initialize cluster"
	DAG_START_OBSERVER            = "Start observer"
	DAG_START_OB                  = "Start OB"
	DAG_STOP_OBSERVER             = "Stop observer"
	DAG_STOP_OB                   = "Stop OB"
	DAG_TAKE_OVER                 = "Take over"
	DAG_CHECK_AND_UPGRADE_OBSHELL = "Check and upgrade obshell"
	DAG_UPGRADE_OBSHELL           = "Upgrade obshell"
	DAG_UPGRADE_CHECK_OBSHELL     = "Upgrade check obshell"
	DAG_UPGRADE_CHECK_OB          = "Upgrade check OB"
	DAG_OB_STOP_SVC_UPGRADE       = "OB stop service upgrade"
	DAG_OB_ROLLING_UPGRADE        = "OB rolling upgrade"
	DAG_NAME_LOCAL_SCALE_OUT      = "Local scale out"
	DAG_NAME_CLUSTER_SCALE_OUT    = "Cluster scale out"

	// rpc retry times
	MAX_RETRY_RPC_TIMES = 3
	RPC_RETRY_INTERVAL  = 1

	// stop ob retry times
	STOP_OB_MAX_RETRY_TIME     = 15
	STOP_OB_MAX_RETRY_INTERVAL = 5

	// additional key
	ADDL_KEY_SUB_DAGS    = "sub_dags"
	ADDL_KEY_MAIN_DAG_ID = "main_dag_id"
)

var (
	// denied configs in ObServerConfigParams
	DeniedConfig = []string{
		constant.CONFIG_CLUSTER_ID,
		constant.CONFIG_CLUSTER_NAME,
		constant.CONFIG_ROOT_PASSWORD,
		constant.CONFIG_RS_LIST,
		constant.CONFIG_ZONE,
		constant.CONFIG_LOCAL_IP,
		constant.CONFIG_DEV_NAME,
	}

	// the default port map for ob
	defaultPortMap = map[string]int{
		constant.CONFIG_RPC_PORT:   constant.DEFAULT_RPC_PORT,
		constant.CONFIG_MYSQL_PORT: constant.DEFAULT_MYSQL_PORT,
	}
	// the list of all directory configuration items required by OB,
	// already arranged according to their hierarchical relationship.
	allDirOrder = []string{
		constant.CONFIG_HOME_PATH,
		constant.CONFIG_DATA_DIR,
		constant.CONFIG_REDO_DIR,
		constant.CONFIG_CLOG_DIR,
		constant.CONFIG_SLOG_DIR,
	}
	// a subset of allDirOrder, excluding the homePath.
	storeDirOrder = allDirOrder[1:]
	// a map that represents the parent directory key for each directory in storeDirOrder.
	parentDirKeys = map[string]string{
		constant.CONFIG_DATA_DIR: constant.CONFIG_HOME_PATH,
		constant.CONFIG_REDO_DIR: constant.CONFIG_DATA_DIR,
		constant.CONFIG_CLOG_DIR: constant.CONFIG_REDO_DIR,
		constant.CONFIG_SLOG_DIR: constant.CONFIG_DATA_DIR,
	}
	// a map that represents the actual directory name for each directory configuration item key.
	dirMap = map[string]string{
		constant.CONFIG_DATA_DIR: constant.OB_DIR_STORE,
		constant.CONFIG_REDO_DIR: "",
		constant.CONFIG_CLOG_DIR: constant.OB_DIR_CLOG,
		constant.CONFIG_SLOG_DIR: constant.OB_DIR_SLOG,
	}
	// unconfigurable dir list
	unconfigurableDirList = []string{
		constant.OB_DIR_ETC,
		constant.OB_DIR_LOG,
		constant.OB_DIR_SSTABLE,
	}
	// clear dir map. the value is prefix
	clearDirMap = map[string]string{
		constant.OB_DIR_ETC: constant.OB_CONFIG_FILE,
		constant.OB_DIR_LOG: "",
	}

	// start
	requiredConfigItems = []string{
		constant.CONFIG_CLUSTER_NAME,
		constant.CONFIG_CLUSTER_ID,
		constant.CONFIG_ZONE,
		constant.CONFIG_MYSQL_PORT,
		constant.CONFIG_RPC_PORT,
	}
	startOptionsMap = map[string]string{
		constant.CONFIG_CLUSTER_NAME: "-n",
		constant.CONFIG_CLUSTER_ID:   "-c",
		constant.CONFIG_DATA_DIR:     "-d",
		constant.CONFIG_ZONE:         "-z",
		constant.CONFIG_DEV_NAME:     "-i",
		constant.CONFIG_LOCAL_IP:     "-I",
		constant.CONFIG_MYSQL_PORT:   "-p",
		constant.CONFIG_RPC_PORT:     "-P",
		constant.CONFIG_LOG_LEVEL:    "-l",
		constant.CONFIG_RS_LIST:      "-r",
	}
	nonStartItems = append(allDirOrder, constant.CONFIG_ROOT_PWD)

	agentService       = agent.AgentService{}
	observerService    = obcluster.ObserverService{}
	obclusterService   = obcluster.ObclusterService{}
	localTaskService   = taskservice.NewLocalTaskService()
	clusterTaskService = taskservice.NewClusterTaskService()
	taskService        = taskservice.NewClusterTaskService()
)

func RegisterObInitTask() {
	task.RegisterTaskType(UpdateOBClusterConfigTask{})
	task.RegisterTaskType(UpdateOBServerConfigTask{})
	task.RegisterTaskType(IntegrateObConfigTask{})
	task.RegisterTaskType(DeployTask{})
	task.RegisterTaskType(DestroyTask{})
	task.RegisterTaskType(StartObserverTask{})
	task.RegisterTaskType(StopObserverTask{})
	task.RegisterTaskType(ClusterBoostrapTask{})
	task.RegisterTaskType(MigrateTableTask{})
	task.RegisterTaskType(ModifyPwdTask{})
	task.RegisterTaskType(MigrateDataTask{})
	task.RegisterTaskType(ConvertFollowerToClusterAgentTask{})
	task.RegisterTaskType(AgentSyncTask{})
	task.RegisterTaskType(ConvertMasterToClusterAgentTask{})
}

func RegisterObStartTask() {
	task.RegisterTaskType(CreateSubStartDagTask{})
	task.RegisterTaskType(CheckSubStartDagReadyTask{})
	task.RegisterTaskType(RetrySubStartDagTask{})
	task.RegisterTaskType(WaitSubStartDagFinishTask{})
	task.RegisterTaskType(StartZoneTask{})
	task.RegisterTaskType(PassSubStartDagTask{})
	task.RegisterTaskType(CheckDagStageTask{})
	task.RegisterTaskType(CheckObserverForStartTask{})
	task.RegisterTaskType(AlterStartServerTask{})
	task.RegisterTaskType(WaitPassOperatorTask{})
	task.RegisterTaskType(GetConnForEStartTask{})
}

func RegisterObStopTask() {
	task.RegisterTaskType(CreateSubStopDagTask{})
	task.RegisterTaskType(CheckSubStopDagReadyTask{})
	task.RegisterTaskType(RetrySubStopDagTask{})
	task.RegisterTaskType(WaitSubStopDagFinishTask{})
	task.RegisterTaskType(PassSubStopDagTask{})
	task.RegisterTaskType(CheckDagStageTask{})
	task.RegisterTaskType(ExecStopSqlTask{})
	task.RegisterTaskType(WaitPassOperatorTask{})
	task.RegisterTaskType(MinorFreezeTask{})
}

func RegisterObScaleOutTask() {
	// for Scaling Agent
	task.RegisterTaskType(AgentBeScalingOutTask{})
	task.RegisterTaskType(WaitDeployRetryTask{})
	task.RegisterTaskType(WaitStartRetryTask{})
	task.RegisterTaskType(WatchDagTask{})
	task.RegisterTaskType(SyncFromOB{})

	// for Cluster Agent
	task.RegisterTaskType(IntegrateSingleObConfigTask{})
	task.RegisterTaskType(CreateLocalScaleOutDagTask{})
	task.RegisterTaskType(WaitScalingReadyTask{})
	task.RegisterTaskType(WaitRemoteDeployTaskFinish{})
	task.RegisterTaskType(WaitRemoteStartTaskFinish{})
	task.RegisterTaskType(PrevCheckTask{})
	task.RegisterTaskType(AddNewZoneTask{})
	task.RegisterTaskType(StartNewZoneTask{})
	task.RegisterTaskType(AddServerTask{})
	task.RegisterTaskType(AddAgentTask{})
	task.RegisterTaskType(FinishTask{})
}

func RegisterUpgradeTask() {
	// upgrade check
	task.RegisterTaskType(CreateUpgradeDirTask{})
	task.RegisterTaskType(GetAllRequiredPkgsTask{})
	task.RegisterTaskType(CheckAllRequiredPkgsTask{})
	task.RegisterTaskType(InstallAllRequiredPkgsTask{})
	task.RegisterTaskType(RemoveUpgradeCheckDirTask{})

	// agent upgrade
	task.RegisterTaskType(BackupAgentForUpgradeTask{})
	task.RegisterTaskType(InstallNewAgentTask{})
	task.RegisterTaskType(RestartAgentTask{})
	task.RegisterTaskType(UpgradePostTableMaintainTask{})

	// ob upgrade
	task.RegisterTaskType(CheckEnvTask{})
	task.RegisterTaskType(BackupParametersTask{})
	task.RegisterTaskType(ExecScriptTask{})
	task.RegisterTaskType(StopZoneTask{})
	task.RegisterTaskType(ReinstallAndRestartObTask{})
	task.RegisterTaskType(StartOneZoneTask{})
	task.RegisterTaskType(RestoreParametersTask{})
}
