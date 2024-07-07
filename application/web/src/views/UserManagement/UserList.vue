<template>
  <div class="user-list">
    <el-table :data="users" style="width: 100%">
      <el-table-column prop="username" label="用户名" width="180"/>
      <el-table-column prop="email" label="邮箱"/>
      <el-table-column prop="role" label="角色" width="180"/>
      <el-table-column fixed="right" label="操作" width="120">
        <template slot-scope="scope">
          <el-button @click="editUser(scope.row.id)" type="text" size="small">编辑</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getUserList } from '@/api/user';

export default {
  data() {
    return {
      users: []
    };
  },
  created() {
    this.fetchUsers();
  },
  methods: {
    async fetchUsers() {
      try {
        const response = await getUserList();
        this.users = response.data;
      } catch (error) {
        console.error('Failed to fetch users:', error);
      }
    },
    editUser(id) {
      this.$router.push({ name: 'UserProfile', params: { id } });
    }
  }
};
</script>

<style scoped>
.user-list {
  padding: 20px;
}
</style>
