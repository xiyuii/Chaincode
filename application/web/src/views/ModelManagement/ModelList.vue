<template>
  <div class="model-list">
    <el-table :data="models" style="width: 100%">
      <el-table-column prop="name" label="模型名称" width="180"/>
      <el-table-column prop="description" label="描述"/>
      <el-table-column prop="createdAt" label="上传时间" width="180"/>
      <el-table-column fixed="right" label="操作" width="120">
        <template slot-scope="scope">
          <el-button @click="viewDetails(scope.row.id)" type="text" size="small">查看详情</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getModelList } from '@/api/model';

export default {
  data() {
    return {
      models: []
    };
  },
  created() {
    this.fetchModels();
  },
  methods: {
    async fetchModels() {
      try {
        const response = await getModelList();
        this.models = response.data;
      } catch (error) {
        console.error('Failed to fetch models:', error);
      }
    },
    viewDetails(id) {
      this.$router.push({ name: 'ModelDetails', params: { id } });
    }
  }
};
</script>

<style scoped>
.model-list {
  padding: 20px;
}
</style>
