<template>
  <div class="model-management-container">
    <el-input v-model="modelName" placeholder="请输入模型名称" style="width: 300px;margin-right: 15px;" />
    <el-button type="primary" plain @click="uploadModel">上传模型</el-button>
    <el-button type="success" plain @click="fetchModels">获取所有模型</el-button>
    <el-table :data="modelData" style="width: 100%">
      <el-table-column label="模型名称" prop="name" />
      <el-table-column label="上传时间" prop="uploadTime" />
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button type="primary" plain @click="downloadModel(scope.row)">下载</el-button>
          <el-button type="danger" plain @click="deleteModel(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { uploadModel, fetchAllModels, downloadModel, deleteModel } from '@/api/model';

export default {
  name: 'ModelManagement',
  data() {
    return {
      modelName: '',
      modelData: []
    };
  },
  methods: {
    uploadModel() {
      const formData = new FormData();
      formData.append('modelName', this.modelName);
      uploadModel(formData).then(response => {
        this.$message.success('模型上传成功');
        this.fetchModels();
      }).catch(error => {
        this.$message.error('模型上传失败');
      });
    },
    fetchModels() {
      fetchAllModels().then(response => {
        this.modelData = response.data;
      }).catch(error => {
        this.$message.error('获取模型列表失败');
      });
    },
    downloadModel(model) {
      downloadModel(model.id).then(response => {
        this.$message.success('模型下载成功');
      }).catch(error => {
        this.$message.error('模型下载失败');
      });
    },
    deleteModel(model) {
      deleteModel(model.id).then(response => {
        this.$message.success('模型删除成功');
        this.fetchModels();
      }).catch(error => {
        this.$message.error('模型删除失败');
      });
    }
  },
  created() {
    this.fetchModels();
  }
};
</script>

<style scoped>
.model-management-container {
  margin: 30px;
}
</style>
