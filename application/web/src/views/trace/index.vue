<template>
  <div class="trace-container">
    <el-input v-model="input" placeholder="请输入溯源码查询" style="width: 300px;margin-right: 15px;" />
    <el-button type="primary" plain @click="ModelInfo"> 查询 </el-button>
    <el-button type="success" plain @click="AllModelInfo"> 获取所有AI模型信息 </el-button>
    <el-table
      :data="tracedata"
      style="width: 100%"
    >
      <el-table-column type="expand">
        <template slot-scope="props">
          <el-form label-position="left" inline class="demo-table-expand">
            <div><span class="trace-text" style="color: #67C23A;">AI模型研发信息</span></div>
            <el-form-item label="AI模型名称：">
              <span>{{ props.row.developer_input.aiModelName }}</span>
            </el-form-item>
            <el-form-item label="研发批次：">
              <span>{{ props.row.developer_input.developmentBatch }}</span>
            </el-form-item>
            <el-form-item label="大模型发布时间：">
              <span>{{ props.row.developer_input.publishTime }}</span>
            </el-form-item>
            <el-form-item label="训练参数：">
              <span>{{ props.row.developer_input.trainingParams }}</span>
            </el-form-item>
            <el-form-item label="研发机构或组织：">
              <span>{{ props.row.developer_input.researchOrg }}</span>
            </el-form-item>
            <el-form-item label="区块链交易ID：">
              <span>{{ props.row.developer_input.de_txid }}</span>
            </el-form-item>
            <el-form-item label="区块链交易时间：">
              <span>{{ props.row.developer_input.de_timestamp }}</span>
            </el-form-item>
            <div><span class="trace-text" style="color: #409EFF;">AI模型发布信息</span></div>
            <el-form-item label="发布组织名称：">
              <span>{{ props.row.publisher_input.fac_productName }}</span>
            </el-form-item>
            <el-form-item label="发布平台：">
              <span>{{ props.row.publisher_input.platform }}</span>
            </el-form-item>
            <el-form-item label="发布信息：">
              <span>{{ props.row.publisher_input.publishInfo }}</span>
            </el-form-item>
            <el-form-item label="曾发布其他模型记录：">
              <span>{{ props.row.publisher_input.otherModelsHistory }}</span>
            </el-form-item>
            <el-form-item label="联系方式：">
              <span>{{ props.row.publisher_input.contactInfo }}</span>
            </el-form-item>
            <el-form-item label="区块链交易ID：">
              <span>{{ props.row.publisher_input.pu_txid }}</span>
            </el-form-item>
            <el-form-item label="区块链交易时间：">
              <span>{{ props.row.publisher_input.pu_timestamp }}</span>
            </el-form-item>
            <div><span class="trace-text" style="color: #E6A23C;">AI模型共享信息</span></div>
            <el-form-item label="共享时间：">
              <span>{{ props.row.sharer_input.sharingTime }}</span>
            </el-form-item>
            <el-form-item label="使用时间：">
              <span>{{ props.row.sharer_input.usageTime }}</span>
            </el-form-item>
            <el-form-item label="使用信息：">
              <span>{{ props.row.sharer_input.phoneNumber }}</span>
            </el-form-item>
            <el-form-item label="使用单位或组织">
              <span>{{ props.row.sharer_input.usingOrganization }}</span>
            </el-form-item>
            <el-form-item label="使用单位联系方式">
              <span>{{ props.row.sharer_input.contactInfo }}</span>
            </el-form-item>
            <el-form-item label="区块链交易ID：">
              <span>{{ props.row.sharer_input.sh_txid }}</span>
            </el-form-item>
            <el-form-item label="区块链交易时间：">
              <span>{{ props.row.sharer_input.sh_timestamp }}</span>
            </el-form-item>
            <div><span class="trace-text" style="color: #909399;">AI模型使用反馈信息</span></div>
            <el-form-item label="反馈时间：">
              <span>{{ props.row.user_input.feedbackTime }}</span>
            </el-form-item>
            <el-form-item label="销售时间：">
              <span>{{ props.row.user_input.salesTime }}</span>
            </el-form-item>
            <el-form-item label="反馈用户名称：">
              <span>{{ props.row.user_input.userName }}</span>
            </el-form-item>
            <el-form-item label="反馈用户位置：">
              <span>{{ props.row.user_input.userLocation }}</span>
            </el-form-item>
            <el-form-item label="反馈用户联系方式：">
              <span>{{ props.row.user_input.userContactInfo }}</span>
            </el-form-item>
            <el-form-item label="区块链交易ID：">
              <span>{{ props.row.user_input.u_txid }}</span>
            </el-form-item>
            <el-form-item label="区块链交易时间：">
              <span>{{ props.row.user_input.u_timestamp }}</span>
            </el-form-item>
          </el-form>
        </template>
      </el-table-column>
      <el-table-column
        label="溯源码"
        prop="traceability_code"
      />
      <el-table-column
        label="AI名称"
        prop="developer_input.aiModelName"
      />
      <el-table-column
        label="AI模型上传时间"
        prop="developer_input.publishTime"
      />
    </el-table>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { getModelInfo, getModelList, getAllModelInfo, getModelHistory } from '@/api/trace'

export default {
  name: 'Trace',
  data() {
    return {
      tracedata: [],
      loading: false,
      input: ''
    }
  },
  computed: {
    ...mapGetters([
      'name',
      'userType'
    ])
  },
  created() {
    getModelList().then(res => {
      this.tracedata = JSON.parse(res.data)
    })
  },
  methods: {
    AllModelInfo() {
      getAllModelInfo().then(res => {
        this.tracedata = JSON.parse(res.data)
      })
    },
    ModelHistory() {
      getModelHistory().then(res => {
        // console.log(res)
      })
    },
    ModelInfo() {
      var formData = new FormData()
      formData.append('traceability_code', this.input)
      getModelInfo(formData).then(res => {
        if (res.code === 200) {
          // console.log(res)
          this.tracedata = []
          this.tracedata[0] = JSON.parse(res.data)
          return
        } else {
          this.$message.error(res.message)
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>

.demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
.trace {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
