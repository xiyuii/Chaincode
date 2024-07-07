<template>
  <div>
    <h2>Upload AI Model</h2>
    <input type="file" @change="handleFileUpload">
    <button @click="uploadModel">Upload</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      file: null
    }
  },
  methods: {
    handleFileUpload(event) {
      this.file = event.target.files[0]
    },
    async uploadModel() {
      // 调用API上传文件
    }
  }
}
</script>
