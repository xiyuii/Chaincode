<template>
  <div class="upload-container">
    <h2 class="upload-title">Upload AI Model</h2>
    <div class="upload-input">
      <input type="file" id="file-upload" @change="handleFileUpload" />
      <label for="file-upload" class="custom-file-label">Choose File</label>
    </div>
    <button @click="uploadModel" class="upload-button">Upload</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      file: null
    };
  },
  methods: {
    handleFileUpload(event) {
      this.file = event.target.files[0];
    },
    async uploadModel() {
      // 调用API上传文件
    }
  }
};
</script>

<style scoped>
.upload-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  background: linear-gradient(135deg, #d16ba5, #86a8e7, #5ffbf1);
}

.upload-title {
  font-family: 'Nunito', sans-serif;
  font-size: 24px;
  color: #333;
  margin-bottom: 20px;
}

.upload-input {
  position: relative;
  margin-bottom: 20px;
}

#file-upload {
  opacity: 0;
  position: absolute;
  z-index: -1;
}

.custom-file-label {
  display: inline-block;
  padding: 10px 20px;
  background-color: #7f53ac;
  color: white;
  border-radius: 8px;
  cursor: pointer;
  font-family: 'Nunito', sans-serif;
  transition: background 0.3s ease;
}

.custom-file-label:hover {
  background-color: #647dee;
}

.upload-button {
  padding: 12px 24px;
  background-color: #7f53ac;
  background-image: linear-gradient(135deg, #7f53ac, #647dee);
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  cursor: pointer;
  transition: background 0.3s ease;
  font-family: 'Nunito', sans-serif;
}

.upload-button:hover {
  background-image: linear-gradient(135deg, #647dee, #7f53ac);
}
</style>
