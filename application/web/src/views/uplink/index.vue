<template>
  <div class="uplink-container">
    <div style="color:#909399;margin-bottom: 30px">
      当前用户：{{ name }};
      用户角色: {{ userType }}
    </div>
    <div>
      <el-form ref="form" :model="tracedata" label-width="80px" size="mini" style="">
        <el-form-item v-show="userType!='研发者'&userType!='消费者'" label="溯源码:" style="width: 300px" label-width="120px">
          <el-input v-model="tracedata.traceability_code" />
        </el-form-item>

        <div v-show="userType=='研发者'">
          <el-form-item label="AI模型名称:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Developer_input.De_AIModelName" />
          </el-form-item>
          <el-form-item label="研发批次:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Developer_input.De_DevelopmentBatch" />
          </el-form-item>
          <el-form-item label="大模型发布时间:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Developer_input.De_publishTime" />
          </el-form-item>
          <el-form-item label="训练参数:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Developer_input.De_TrainingParams" />
          </el-form-item>
          <el-form-item label="研发机构或组织:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Developer_input.De_ResearchOrg" />
          </el-form-item>
        </div>
        <div v-show="userType=='发布者'">
          <el-form-item label="发布组织名称:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Publisher_input.Pu_OrganizationName" />
          </el-form-item>
          <el-form-item label="发布平台:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Publisher_input.Pu_Platform" />
          </el-form-item>
          <el-form-item label="发布信息:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Publisher_input.Pu_PublishInfo" />
          </el-form-item>
          <el-form-item label="曾发布其他模型记录:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Publisher_input.Pu_OtherModelsHistory" />
          </el-form-item>
          <el-form-item label="联系方式:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Publisher_input.Pu_ContactInfo" />
          </el-form-item>
        </div>
        <div v-show="userType=='共享者'">
          <el-form-item label="共享时间:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Sharer_input.Sh_SharingTime" />
          </el-form-item>
          <el-form-item label="使用时间:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Sharer_input.Sh_UsageTime" />
          </el-form-item>
          <el-form-item label="使用信息:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Sharer_input.Sh_PhoneNumber" />
          </el-form-item>
          <el-form-item label="使用单位或组织:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Sharer_input.Sh_UsingOrganization" />
          </el-form-item>
          <el-form-item label="使用单位联系方式:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Sharer_input.Sh_ContactInfo" />
          </el-form-item>
        </div>
        <div v-show="userType=='反馈者'">
          <el-form-item label="反馈时间:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.User_input.U_FeedbackTime" />
          </el-form-item>
          <el-form-item label销售时间:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.User_input.U_SalesTime" />
          </el-form-item>
          <el-form-item label="反馈用户名称:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.User_input.U_UserName" />
          </el-form-item>
          <el-form-item label="反馈用户位置:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.User_input.U_UserLocation" />
          </el-form-item>
          <el-form-item label="反馈用户联系方式:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.User_input.U_UserContactInfo" />
          </el-form-item>
        </div>
      </el-form>
      <span slot="footer" style="color: gray;" class="dialog-footer">
        <el-button v-show="userType != '消费者'" type="primary" plain style="margin-left: 220px;" @click="submittracedata()">提 交</el-button>
      </span>
      <span v-show="userType == '消费者'" slot="footer" style="color: gray;" class="dialog-footer">
        消费者没有权限录入！请使用溯源功能!
      </span>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { uplink } from '@/api/trace'

export default {
  name: 'Uplink',
  data() {
    return {
      tracedata: {
        traceability_code: '',
        Developer_input: {
          De_AIModelName: '',
          De_DevelopmentBatch: '',
          De_PublishTime: '',
          De_TrainingParams: '',
          De_ResearchOrg: ''
          
        },
        Publisher_input: {
          Pu_OrganizationName: '',
          Pu_Platform: '',
          Pu_PublishInfo: '',
          Pu_OtherModelsHistory: '',
          Pu_ContactInfo: ''
        },
        Sharer_input: {
          Sh_SharingTime: '',
          Sh_UsageTime: '',
          Sh_PhoneNumber: '',
          Sh_UsingOrganization: '',
          Sh_ContactInfo: ''
        },
        User_input: {
          U_FeedbackTime: '',
          U_SalesTime: '',
          U_UserName: '',
          U_UserLocation: '',
          U_UserContactInfo: ''
        }
      },
      loading: false
    }
  },
  computed: {
    ...mapGetters([
      'name',
      'userType'
    ])
  },
  methods: {
    submittracedata() {
      console.log(this.tracedata)
      const loading = this.$loading({
        lock: true,
        text: '数据上链中...',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      var formData = new FormData()
      formData.append('traceability_code', this.tracedata.traceability_code)
      // 根据不同的用户给arg1、arg2、arg3..赋值,
      switch (this.userType) {
        case '研发者':
          formData.append('arg1', this.tracedata.Developer_input.De_AIModelName)
          formData.append('arg2', this.tracedata.Developer_input.De_PublishTime)
          formData.append('arg3', this.tracedata.Developer_input.De_DevelopmentBatch)
          formData.append('arg4', this.tracedata.Developer_input.De_TrainingParams)
          formData.append('arg5', this.tracedata.Developer_input.De_ResearchOrg)
          break
        case '发布者':
          formData.append('arg1', this.tracedata.Publisher_input.Pu_OrganizationName)
          formData.append('arg2', this.tracedata.Publisher_input.Pu_Platform)
          formData.append('arg3', this.tracedata.Publisher_input.Pu_PublishInfo)
          formData.append('arg4', this.tracedata.Publisher_input.Pu_OtherModelsHistory)
          formData.append('arg5', this.tracedata.Publisher_input.Pu_ContactInfo)
          break
        case '共享者':
          formData.append('arg1', this.tracedata.Sharer_input.Sh_SharingTime)
          formData.append('arg2', this.tracedata.Sharer_input.Sh_UsageTime)
          formData.append('arg3', this.tracedata.Sharer_input.Sh_PhoneNumber)
          formData.append('arg4', this.tracedata.Sharer_input.Sh_UsingOrganization)
          formData.append('arg5', this.tracedata.Sharer_input.Sh_ContactInfo)
          break
        case '反馈者':
          formData.append('arg1', this.tracedata.User_input.U_FeedbackTime)
          formData.append('arg2', this.tracedata.User_input.U_SalesTime)
          formData.append('arg3', this.tracedata.User_input.U_UserName)
          formData.append('arg4', this.tracedata.User_input.U_UserLocation)
          formData.append('arg5', this.tracedata.User_input.U_UserContactInfo)
          break
      }
      uplink(formData).then(res => {
        if (res.code === 200) {
          loading.close()
          this.$message({
            message: '上链成功，交易ID：' + res.txid + '\n溯源码：' + res.traceability_code,
            type: 'success'
          })
        } else {
          loading.close()
          this.$message({
            message: '上链失败',
            type: 'error'
          })
        }
      }).catch(err => {
        loading.close()
        console.log(err)
      })
    }
  }
}

</script>

<style lang="scss" scoped>
.uplink {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
