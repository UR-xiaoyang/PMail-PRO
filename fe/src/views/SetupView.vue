<template>
  <div id="main" class="setup-container">
    <el-card class="setup-card" shadow="hover">
      <div class="header">
        <img src="@/assets/logo.svg" alt="PMail Logo" class="logo" />
        <h1 class="title">PMail Pro Setup</h1>
      </div>

      <el-steps :active="active" align-center finish-status="success" class="steps-bar">
        <el-step :title="lang.welcome" :icon="HomeFilled" />
        <el-step :title="lang.setDatabase" :icon="Coin" />
        <el-step :title="lang.setAdminPassword" :icon="Lock" />
        <el-step :title="lang.SetDomail" :icon="Link" />
        <el-step :title="lang.setDNS" :icon="Setting" />
        <el-step :title="lang.setSSL" :icon="Key" />
        <el-step :title="lang.finish" :icon="Select" />
      </el-steps>

      <div class="content-wrapper">
        <transition name="fade-slide" mode="out-in">
          <div :key="active" class="step-content">
            
            <!-- Step 0: Welcome -->
            <div v-if="active === 0" class="step-pane welcome-pane">
              <div class="desc text-center">
                <h2>{{ lang.tks_pmail }}</h2>
                <p class="subtitle">{{ lang.guid_desc }}</p>
              </div>
            </div>

            <!-- Step 1: Database -->
            <div v-if="active === 1" class="step-pane">
              <div class="desc">
                <h2>{{ lang.select_db }}</h2>
                <p class="subtitle">{{ lang.db_desc }}</p>
              </div>
              <div class="form-container">
                <el-form label-width="120px" class="setup-form">
                  <el-form-item :label="lang.type">
                    <el-select :placeholder="lang.db_select_ph" v-model="dbSettings.type" @change="dbSettings.dsn = ''">
                      <el-option label="MySQL" value="mysql"/>
                      <el-option label="SQLite3" value="sqlite"/>
                      <el-option label="PostgreSQL" value="postgres"/>
                    </el-select>
                  </el-form-item>

                  <el-form-item :label="lang.mysql_dsn" v-if="dbSettings.type === 'mysql'">
                    <el-input :rows="2" type="textarea" v-model="dbSettings.dsn"
                              placeholder="root:12345@tcp(127.0.0.1:3306)/pmail?parseTime=True&loc=Local"></el-input>
                  </el-form-item>

                  <el-form-item :label="lang.pg_dsn" v-if="dbSettings.type === 'postgres'">
                    <el-input :rows="2" type="textarea" v-model="dbSettings.dsn"
                              placeholder="postgres://postgres:12345@127.0.0.1:5432/pmail?sslmode=disable"></el-input>
                  </el-form-item>

                  <el-form-item :label="lang.sqlite_db_path" v-if="dbSettings.type === 'sqlite'">
                    <el-input v-model="dbSettings.dsn" placeholder="./config/pmail.db"></el-input>
                  </el-form-item>
                </el-form>
              </div>
            </div>

            <!-- Step 2: Admin Password -->
            <div v-if="active === 2" class="step-pane">
              <div class="desc">
                <h2>{{ lang.setAdminPassword }}</h2>
              </div>
              <div class="form-container">
                <el-form label-width="120px" class="setup-form">
                  <el-form-item :label="lang.admin_account">
                    <el-input v-bind:disabled="adminSettings.hadSeted" placeholder="admin"
                              v-model="adminSettings.account"></el-input>
                  </el-form-item>

                  <el-form-item :label="lang.password">
                    <el-input type="password" v-bind:disabled="adminSettings.hadSeted" placeholder=""
                              v-model="adminSettings.password" show-password></el-input>
                  </el-form-item>

                  <el-form-item :label="lang.enter_again">
                    <el-input type="password" v-bind:disabled="adminSettings.hadSeted" placeholder=""
                              v-model="adminSettings.password2" show-password></el-input>
                  </el-form-item>
                </el-form>
              </div>
            </div>

            <!-- Step 3: Domain Settings -->
            <div v-if="active === 3" class="step-pane">
              <div class="desc">
                <h2>{{ lang.SetDomail }}</h2>
              </div>
              <div class="form-container">
                <el-form label-width="140px" class="setup-form">
                  <el-form-item :label="lang.smtp_domain">
                    <el-input placeholder="domain.com" v-model="domainSettings.smtp_domain">
                      <template #prepend>smtp.</template>
                    </el-input>
                  </el-form-item>

                  <el-form-item :label="lang.web_domain">
                    <el-input placeholder="pmail.domain.com" v-model="domainSettings.web_domain"></el-input>
                  </el-form-item>

                  <el-form-item :label="lang.multi_domain_setting">
                    <div class="multi-domain-header">
                      <span>{{ lang.multi_domain_setting_desc }}</span>
                      <el-button @click="addDomain" size="small" type="success" :icon="Plus" circle class="add-btn"></el-button>
                    </div>
                    <div v-for="(item, i) in domainSettings.multi_domain" :key="i" class="domain-input-group">
                       <el-input :placeholder="'domain' + (i+1) + '.com'" v-model="domainSettings.multi_domain[i]"></el-input>
                    </div>
                  </el-form-item>

                  <el-divider content-position="center">Port Settings</el-divider>

                  <div class="port-settings-grid">
                    <el-form-item label="HTTP Port">
                        <el-input type="number" placeholder="80" v-model="domainSettings.http_port"></el-input>
                    </el-form-item>
                    <el-form-item label="HTTPS Port">
                        <el-input type="number" placeholder="443" v-model="domainSettings.https_port"></el-input>
                    </el-form-item>
                    <el-form-item label="SMTP Port">
                        <el-input type="number" placeholder="25" v-model="domainSettings.smtp_port"></el-input>
                    </el-form-item>
                    <el-form-item label="IMAP Port">
                        <el-input type="number" placeholder="993" v-model="domainSettings.imap_port"></el-input>
                    </el-form-item>
                    <el-form-item label="POP3 Port">
                        <el-input type="number" placeholder="110" v-model="domainSettings.pop3_port"></el-input>
                    </el-form-item>
                    <el-form-item label="SMTPS Port">
                        <el-input type="number" placeholder="465" v-model="domainSettings.smtps_port"></el-input>
                    </el-form-item>
                    <el-form-item label="IMAPS Port">
                        <el-input type="number" placeholder="993" v-model="domainSettings.imaps_port"></el-input>
                    </el-form-item>
                    <el-form-item label="POP3S Port">
                        <el-input type="number" placeholder="995" v-model="domainSettings.pop3s_port"></el-input>
                    </el-form-item>
                  </div>
                </el-form>
              </div>
            </div>

            <!-- Step 4: DNS Settings -->
            <div v-if="active === 4" class="step-pane full-width">
              <div class="desc">
                <h2>{{ lang.setDNS }}</h2>
                <p class="subtitle">{{ lang.dns_desc }}</p>
              </div>
              <div class="dns-container" v-loading="dnsLoading" element-loading-text="Loading..." element-loading-background="rgba(255, 255, 255, 0.8)">
                <div v-for="(info, domain) in dnsInfos" :key="domain" class="dns-group">
                  <h3>{{ domain }}</h3>
                  <el-table :data="info" border style="width: 100%" stripe>
                    <el-table-column prop="host" label="HOSTNAME" width="180">
                        <template #default="scope">
                        <div class="table-cell-content">
                            <el-tooltip :content="lang.dns_root_desc" placement="top"
                                        v-if="scope.row.host === '' || scope.row.host === '@' ">
                            <span class="highlight-text">{{ scope.row.host || '@' }}</span>
                            </el-tooltip>
                            <span v-else>{{ scope.row.host }}</span>
                        </div>
                        </template>
                    </el-table-column>
                    <el-table-column prop="type" label="TYPE" width="100">
                         <template #default="scope">
                            <el-tag>{{ scope.row.type }}</el-tag>
                         </template>
                    </el-table-column>
                    <el-table-column prop="value" label="VALUE">
                        <template #default="scope">
                        <div class="clickable-value" @click="copyToClipboard(scope.row.value)">
                            <el-tooltip :content="scope.row.tips || 'Click to copy'" placement="top">
                                <span>{{ scope.row.value }}</span>
                            </el-tooltip>
                            <el-icon class="copy-icon"><CopyDocument /></el-icon>
                        </div>
                        </template>
                    </el-table-column>
                    <el-table-column prop="ttl" label="TTL" width="80"/>
                  </el-table>
                </div>
              </div>
            </div>

            <!-- Step 5: SSL Settings -->
            <div v-if="active === 5" class="step-pane">
               <el-alert :closable="false" title="Warning!" type="error" center show-icon
                    v-if="active === 5 && sslSettings.type === 0 && port !== 80" :description="lang.autoSSLWarn" class="mb-4"/>

               <el-alert :closable="false" title="SSL Error" type="error" center show-icon
                    v-if="sslError" :description="sslError" class="mb-4"/>

              <div class="desc">
                <h2>{{ lang.setSSL }}</h2>
              </div>
              <div class="form-container">
                <el-form label-width="140px" class="setup-form">
                  <el-form-item :label="lang.type">
                    <el-select :placeholder="lang.ssl_auto" v-model="sslSettings.type" :disabled="dnsChecking">
                      <el-option :label="lang.ssl_auto" value="0"/>
                      <el-option :label="lang.ssl_manuallyf" value="1"/>
                      <el-option :label="lang.ssl_none" value="3"/>
                    </el-select>
                  </el-form-item>

                  <el-alert :closable="false" title="Warning!" type="error" center show-icon
                    v-if="sslSettings.type === '3'" :description="lang.noSSLWarn" class="mb-4" style="margin-bottom: 20px;"/>

                  <el-form-item :label="lang.ssl_challenge_type" v-if="sslSettings.type === '0'">
                    <el-select :placeholder="lang.ssl_auto_http" v-model="sslSettings.challenge"
                               :disabled="dnsChecking">
                      <el-option :label="lang.ssl_auto_http" value="http"/>
                      <el-option :label="lang.ssl_auto_dns" value="dns"/>
                    </el-select>
                    <el-tooltip class="box-item" effect="dark" :content="lang.challenge_typ_desc"
                                placement="top-start">
                       <el-icon class="help-icon"><QuestionFilled /></el-icon>
                    </el-tooltip>
                  </el-form-item>

                  <el-form-item :label="lang.ssl_key_path" v-if="sslSettings.type === '1'">
                    <el-upload
                        class="upload-demo"
                        action=""
                        :auto-upload="false"
                        :on-change="handleKeyChange"
                        :limit="1"
                        :file-list="keyFileList"
                        accept=".key,.pem"
                        >
                        <el-button type="primary">Select Key File</el-button>
                    </el-upload>
                  </el-form-item>

                  <el-form-item :label="lang.ssl_crt_path" v-if="sslSettings.type === '1'">
                    <el-upload
                        class="upload-demo"
                        action=""
                        :auto-upload="false"
                        :on-change="handleCrtChange"
                        :limit="1"
                        :file-list="crtFileList"
                        accept=".crt,.pem,.cer"
                        >
                        <el-button type="primary">Select Crt File</el-button>
                    </el-upload>
                  </el-form-item>
                </el-form>
              </div>

               <div v-if="dnsChecking" class="dns-check-section">
                  <label class="section-label">{{ lang.dns_desc }}</label>
                  <el-table :data="sslSettings.paramsList" border v-loading="sslSettings.paramsList.length === 0" style="margin-top: 10px;">
                    <el-table-column prop="host" label="HOSTNAME" width="180"/>
                    <el-table-column prop="type" label="TYPE" width="100"/>
                    <el-table-column prop="value" label="VALUE">
                      <template #default="scope">
                        <div class="clickable-value" @click="copyToClipboard(scope.row.value)">
                          <el-tooltip :content="scope.row.tips || 'Click to copy'" placement="top">
                            <span>{{ scope.row.value }}</span>
                          </el-tooltip>
                           <el-icon class="copy-icon"><CopyDocument /></el-icon>
                        </div>
                      </template>
                    </el-table-column>
                    <el-table-column prop="ttl" label="TTL" width="80"/>
                  </el-table>
                </div>
            </div>

            <!-- Step 6: Finish -->
            <div v-if="active === 6" class="step-pane finish-pane">
              <div class="desc text-center" style="text-align: center; padding: 40px 0;">
                 <el-icon :size="80" color="#67C23A" style="margin-bottom: 20px;"><Select /></el-icon>
                 <h2>{{ lang.congratulations }}</h2>
                 <p class="subtitle">{{ lang.config_complete }}</p>
              </div>
            </div>

          </div>
        </transition>
      </div>

      <div class="footer-actions">
         <el-button 
            v-if="active !== 6"
            type="primary" 
            size="large"
            :loading="fullscreenLoading"
            class="next-btn"
            @click="next">
            {{ (active === 5 && dnsChecking) ? (sslError ? lang.retry : lang.check) : lang.next }}
            <el-icon class="el-icon--right"><ArrowRight /></el-icon>
        </el-button>
         <el-button 
            v-else
            type="success" 
            size="large"
            class="next-btn"
            @click="finishSetup">
            {{ lang.enter_mailbox }}
            <el-icon class="el-icon--right"><Select /></el-icon>
        </el-button>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import {reactive, ref, watch} from 'vue'
import {ElMessage} from 'element-plus'
import lang from '../i18n/i18n';
import axios from 'axios'
import {
  Plus, 
  HomeFilled, 
  Coin, 
  Lock, 
  Link, 
  Setting, 
  Key, 
  ArrowRight, 
  CopyDocument, 
  QuestionFilled,
  Select
} from '@element-plus/icons-vue'
import {http} from "@/utils/axios";

const waitDesc = ref(lang.wait_desc);

const adminSettings = reactive({
  "account": "admin",
  "password": "",
  "password2": "",
  "hadSeted": false
})

const dbSettings = reactive({
  "type": "sqlite",
  "dsn": "./config/pmail.db",
  "lable": ""
})

const domainSettings = reactive({
  "web_domain": "",
  "smtp_domain": "",
  "multi_domain": [],
  "smtp_port": 0,
  "imap_port": 0,
  "pop3_port": 0,
  "smtps_port": 0,
  "imaps_port": 0,
  "pop3s_port": 0,
  "http_port": 0,
  "https_port": 0
})

const sslSettings = reactive({
  "type": "0",
  "challenge": "http",
  "key_path": "./config/ssl/private.key",
  "crt_path": "./config/ssl/public.crt",
  "paramsList": [],
})


const sslError = ref('')
const active = ref(0)
const fullscreenLoading = ref(false)
const dnsChecking = ref(false)

const dnsInfos = ref({})
const dnsLoading = ref(false)

watch(dnsInfos, (newVal) => {
  console.log("DEBUG: dnsInfos updated:", newVal)
}, { deep: true })

const port = ref(80)


const addDomain = () => {
  domainSettings.multi_domain.push([])
}

const setPassword = () => {
  if (adminSettings.hadSeted) {
    active.value++;
    getDomainConfig();
    return;
  }

  if (adminSettings.password !== adminSettings.password2) {
    ElMessage.error(lang.err_pwd_diff)
  } else {
    http.post("/api/setup", {
      "action": "set",
      "step": "password",
      "account": adminSettings.account,
      "password": adminSettings.password
    }).then((res) => {
      if (res.errorNo !== 0) {
        ElMessage.error(res.errorMsg)
      } else {
        active.value++;
        getDomainConfig();
      }
    })
  }
}

const getPassword = () => {
  http.post("/api/setup", {"action": "get", "step": "password"}).then((res) => {
    if (res.errorNo !== 0) {
      ElMessage.error(res.errorMsg)
    } else {
      adminSettings.hadSeted = res.data !== ""
      if (adminSettings.hadSeted) {
        adminSettings.account = res.data
        adminSettings.password = "*******"
        adminSettings.password2 = "*******"
      }

    }
  })
}


const getDbConfig = () => {
  http.post("/api/setup", {"action": "get", "step": "database"}).then((res) => {
    if (res.errorNo !== 0) {
      ElMessage.error(res.errorMsg)
    } else {
      dbSettings.type = res.data.db_type;
      dbSettings.dsn = res.data.db_dsn;
    }
  })
}

const getDomainConfig = () => {
  http.post("/api/setup", {"action": "get", "step": "domain"}).then((res) => {
    if (res.errorNo !== 0) {
      ElMessage.error(res.errorMsg)
    } else {
      domainSettings.web_domain = res.data.web_domain;
      domainSettings.smtp_domain = res.data.smtp_domain;
      domainSettings.multi_domain = res.data.domains;
      domainSettings.smtp_port = res.data.smtp_port;
      domainSettings.imap_port = res.data.imap_port;
      domainSettings.pop3_port = res.data.pop3_port;
      domainSettings.smtps_port = res.data.smtps_port;
      domainSettings.imaps_port = res.data.imaps_port;
      domainSettings.pop3s_port = res.data.pop3s_port;
      domainSettings.http_port = res.data.http_port;
      domainSettings.https_port = res.data.https_port;
    }
  })
}

const setDbConfig = () => {
  // 切换数据库类型为sqlite时，数据库路径为空，则使用默认路径
  if (dbSettings.type === "sqlite" && !dbSettings.dsn) dbSettings.dsn = "./config/pmail.db";
  else if (!dbSettings.dsn) ElMessage({
    title: "Error",
    message: lang.err_db_dsn_empty,
    type: "error",
  });
  http.post("/api/setup", {
    "action": "set",
    "step": "database",
    "db_type": dbSettings.type,
    "db_dsn": dbSettings.dsn
  }).then((res) => {
    if (res.errorNo !== 0) {
      ElMessage.error(res.errorMsg)
    } else {
      active.value++;
      getPassword();
    }
  })
}

const getDNSConfig = () => {
  dnsLoading.value = true
  http.post("/api/setup", {"action": "get", "step": "dns"}).then((res) => {
    dnsLoading.value = false
    if (res.errorNo !== 0) {
      ElMessage.error(res.errorMsg)
    } else {
      dnsInfos.value = res.data
    }
  }).catch(err => {
    dnsLoading.value = false
  })
}


const getSSLConfig = () => {
  http.post("/api/setup", {"action": "get", "step": "ssl"}).then((res) => {
    if (res.errorNo !== 0) {
      ElMessage.error(res.errorMsg)
    } else {
      sslSettings.type = String(res.data.type)
      if (sslSettings.type === "2") {
        sslSettings.type = "0"
        sslSettings.challenge = "dns"
      }


      port.value = res.data.port
    }
  })
}


const keyFileList = ref([])
const crtFileList = ref([])

const handleKeyChange = (file) => {
    keyFileList.value = [file]
}

const handleCrtChange = (file) => {
    crtFileList.value = [file]
}

const setSSLConfig = () => {
  sslError.value = ''
  fullscreenLoading.value = true;

  let sslType = sslSettings.type;
  if (sslType === "0" && sslSettings.challenge === "dns") {
    sslType = "2"
  }

  if (sslType === "1") {
      // Manual upload
      if (keyFileList.value.length === 0 || crtFileList.value.length === 0) {
          fullscreenLoading.value = false;
          ElMessage.error("Please select both key and crt files")
          return
      }
      
      const formData = new FormData()
      formData.append("action", "set")
      formData.append("step", "ssl")
      formData.append("ssl_type", "1")
      formData.append("key_file", keyFileList.value[0].raw)
      formData.append("crt_file", crtFileList.value[0].raw)
      
      http.post("/api/setup", formData).then((res) => {
        if (res.errorNo !== 0) {
            fullscreenLoading.value = false;
            ElMessage.error(res.errorMsg)
        } else {
             checkStatus();
        }
      })
      return
  }


  http.post("/api/setup", {
    "action": "set",
    "step": "ssl",
    "ssl_type": sslType,
    "key_path": sslSettings.key_path,
    "crt_path": sslSettings.crt_path
  }).then((res) => {
    if (res.errorNo !== 0) {
      fullscreenLoading.value = false;
      ElMessage.error(res.errorMsg)
    } else {
      if (sslType == 2) {
        fullscreenLoading.value = false;
        dnsChecking.value = true;
        getSSLDNSParams();
      } else {
        checkStatus();
      }
    }
  })
}


const checkSSLStatus = () => {
  fullscreenLoading.value = true;
  http.post("/api/setup", {"action": "check_ssl", "step": "ssl"}).then((res) => {
    if (res.data.status === "success") {
      checkStatus();
    } else if (res.data.status === "running") {
      fullscreenLoading.value = false;
      ElMessage.info(lang.wait_desc);
    } else if (res.data.status === "error") {
      sslError.value = res.data.error;
      fullscreenLoading.value = false;
      ElMessage.error(res.data.error);
    } else {
      setSSLConfig();
    }
  })
}


const checkStatus = () => {
  axios.post("/api/ping", {}).then((res) => {
    if (res.data.errorNo !== 0) {
      setTimeout(function () {
        checkStatus()
      }, 1000);
    } else {
      fullscreenLoading.value = false;
      active.value++;
    }
  }).catch(() => {
    setTimeout(function () {
      checkStatus()
    }, 1000);
  })
}


const setDomainConfig = () => {
  http.post("/api/setup", {
    "action": "set",
    "step": "domain",
    "web_domain": domainSettings.web_domain,
    "smtp_domain": domainSettings.smtp_domain,
    "multi_domain": domainSettings.multi_domain.join(","),
    "smtp_port": domainSettings.smtp_port,
    "imap_port": domainSettings.imap_port,
    "pop3_port": domainSettings.pop3_port,
    "smtps_port": domainSettings.smtps_port,
    "imaps_port": domainSettings.imaps_port,
    "pop3s_port": domainSettings.pop3s_port,
    "http_port": domainSettings.http_port,
    "https_port": domainSettings.https_port
  }).then((res) => {
    if (res.errorNo !== 0) {
      ElMessage.error(res.errorMsg)
    } else {
      active.value++;
      getDNSConfig();
    }
  })
}

const getSSLDNSParams = () => {
  http.post("/api/setup", {"action": "getParams", "step": "ssl"}).then((res) => {
    if (res.errorNo !== 0) {
      ElMessage.error(res.errorMsg)
    } else {
      sslSettings.paramsList = res.data
      console.log(sslSettings.paramsList)
    }
  })

  if (sslSettings.paramsList.length === 0) {
    setTimeout(function () {
      getSSLDNSParams()
    }, 1000);
  }


}

const copyToClipboard = (text) => {
  if (!text) return;
  navigator.clipboard.writeText(text).then(() => {
    ElMessage({
      message: "Copied to clipboard",
      type: "success",
    });
  });
};

const finishSetup = () => {
    if (sslSettings.type === "3") {
        window.location.href = "http://" + domainSettings.web_domain;
    } else {
        window.location.href = "https://" + domainSettings.web_domain;
    }
}

const next = () => {
  switch (active.value) {
    case 0:
      active.value++
      getDbConfig();
      break
    case 1:
      setDbConfig();
      break;
    case 2:
      setPassword();
      break;
    case 3:
      setDomainConfig();
      break;
    case 4:
      getSSLConfig();
      active.value++
      break
    case 5:
      if (dnsChecking.value) {
        if (sslError.value) {
          setSSLConfig();
        } else {
          checkSSLStatus();
        }
      } else {
        setSSLConfig();
      }
      break
  }

}
</script>


<style scoped>
.setup-container {
  width: 100%;
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 20px;
}

.setup-card {
  width: 100%;
  max-width: 900px;
  border-radius: 12px;
  overflow: hidden;
  background-color: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
}

.header {
  text-align: center;
  margin-bottom: 30px;
}

.logo {
  height: 60px;
  margin-bottom: 10px;
}

.title {
  font-size: 24px;
  color: #2c3e50;
  font-weight: 600;
  margin: 0;
}

.steps-bar {
  margin-bottom: 40px;
  padding: 0 20px;
}

.content-wrapper {
  min-height: 300px;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 0 20px;
}

.step-content {
  width: 100%;
  max-width: 600px;
}

.step-pane {
  animation: fadeIn 0.5s ease;
}

.welcome-pane {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  padding: 40px 0;
}

.desc {
  text-align: center;
  margin-bottom: 30px;
}

.desc h2 {
  font-size: 22px;
  margin-bottom: 10px;
  color: #303133;
}

.subtitle {
  color: #606266;
  font-size: 14px;
  line-height: 1.6;
}

.form-container {
  display: flex;
  justify-content: center;
}

.setup-form {
  width: 100%;
}

.full-width {
  max-width: 100%;
}

.dns-container {
  width: 100%;
  min-height: 200px;
}

.dns-group {
  margin-bottom: 30px;
}

.dns-group h3 {
  font-size: 18px;
  margin-bottom: 15px;
  color: #409EFF;
  border-bottom: 2px solid #ebeef5;
  padding-bottom: 8px;
}

.table-cell-content {
  display: flex;
  align-items: center;
}

.clickable-value {
  display: flex;
  align-items: center;
  cursor: pointer;
  color: #409EFF;
}

.clickable-value:hover {
  text-decoration: underline;
}

.copy-icon {
  margin-left: 5px;
  font-size: 14px;
}

.help-icon {
  margin-left: 8px;
  font-size: 16px;
  color: #909399;
  cursor: help;
}

.highlight-text {
  font-weight: bold;
}

.multi-domain-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 10px;
}

.domain-input-group {
  margin-bottom: 10px;
}

.port-settings-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 10px;
}

.footer-actions {
  display: flex;
  justify-content: center;
  margin-top: 40px;
  padding-top: 20px;
  border-top: 1px solid #ebeef5;
}

.next-btn {
  width: 200px;
  font-weight: bold;
  letter-spacing: 1px;
}

.mb-4 {
    margin-bottom: 1rem;
}

/* Animations */
.fade-slide-enter-active,
.fade-slide-leave-active {
  transition: opacity 0.3s ease, transform 0.3s ease;
}

.fade-slide-enter-from {
  opacity: 0;
  transform: translateX(20px);
}

.fade-slide-leave-to {
  opacity: 0;
  transform: translateX(-20px);
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}

@media (max-width: 768px) {
    .setup-card {
        padding: 10px;
        width: 100%;
    }
    
    .port-settings-grid {
        grid-template-columns: 1fr;
    }

    .steps-bar {
        display: none; /* Hide steps on very small screens if needed, or adjust */
    }
}
</style>
