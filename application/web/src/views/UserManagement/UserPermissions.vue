<template>
  <div class="user-permissions">
    <el-form :model="permissions" label-width="120px">
      <el-form-item label="角色">
        <el-select v-model="permissions.role" placeholder="请选择角色">
          <el-option label="普通用户" value="user"></el-option>
          <el-option label="管理员" value="admin"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="权限">
        <el-checkbox-group v-model="permissions.access">
          <el-checkbox label="模型上传"></el-checkbox>
          <el-checkbox label="模型下载"></el-checkbox>
          <el-checkbox label="模型管理"></el-checkbox>
        </el-checkbox-group>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="savePermissions">保存</el-button>
      </el-item>
    </el-form>
  </div>
</template>

<script>
import { getUserPermissions, updateUserPermissions } from '@/api/permissions';

export default {
  data() {
    return {
      permissions: {
        role: '',
        access: []
      }
    };
  },
  created() {
    this.fetchUserPermissions();
  },
  methods: {
    async fetchUserPermissions() {
      try {
        const response = await getUserPermissions(this.$route.params.id);
        this.permissions = response.data;
      } catch (error) {
        console.error('Failed to fetch user permissions:', error);
      }
    },
    async savePermissions() {
      try {
        await updateUserPermissions(this.$route.params.id, this.permissions);
        this.$message.success('用户权限更新成功');
      } catch (error) {
        console.error('Failed to update user permissions:', error);
        this.$message.error('用户权限更新失败');
      }
    }
  }
};
</script>

<style scoped>
.user-permissions {
  padding: 20px;
}
</style>
