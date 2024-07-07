<template>
  <div class="model-details">
    <el-card>
      <h2>{{ model.name }}</h2>
      <p><strong>描述：</strong>{{ model.description }}</p>
      <p><strong>上传时间：</strong>{{ model.createdAt }}</p>
      <el-button @click="downloadModel" type="primary">下载模型</el-button>
    </el-card>
  </div>
</template>

<script>
import { getModelDetails } from '@/api/model';

export default {
  data() {
    return {
      model: {}
    };
  },
  created() {
    this.fetchModelDetails();
  },
  methods: {
    async fetchModelDetails() {
      const { id } = this.$route.params;
      try {
        const response = await getModelDetails(id);
        this.model = response.data;
      } catch (error) {
        console.error('Failed to fetch model details:', error);
      }
    },
    downloadModel() {
      // 模型下载逻辑
      console.log('下载模型:', this.model.id);
    }
  }
};
</script>

<style scoped>
.model-details {
  padding: 20px;
}
</style>
