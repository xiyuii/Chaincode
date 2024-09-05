<template>
  <div class="user-profile">
    <el-form :model="user" label-width="120px">
      <el-form-item label="用户名">
        <el-input v-model="user.username" disabled></el-input>
      </el-form-item>
      <el-form-item label="邮箱">
        <el-input v-model="user.email"></el-input>
      </el-form-item>
      <el-form-item label="角色">
        <el-select v-model="user.role" placeholder="请选择角色">
          <el-option label="普通用户" value="user"></el-option>
          <el-option label="管理员" value="admin"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="saveProfile">保存</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { getUserProfile, updateUserProfile } from '@/api/user';

export default {
  data() {
    return {
      user: {
        username: '',
        email: '',
        role: ''
      }
    };
  },
  created() {
    this.fetchUserProfile();
  },
  methods: {
    async fetchUserProfile() {
      try {
        const response = await getUserProfile(this.$route.params.id);
        this.user = response.data;
      } catch (error) {
        console.error('Failed to fetch user profile:', error);
      }
    },
    async saveProfile() {
      try {
        await updateUserProfile(this.$route.params.id, this.user);
        this.$message.success('用户资料更新成功');
      } catch (error) {
        console.error('Failed to update user profile:', error);
        this.$message.error('用户资料更新失败');
      }
    }
  }
};
</script>

<style scoped>
.user-profile {
  padding: 20px;
}
</style>
