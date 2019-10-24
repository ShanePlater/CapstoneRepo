<template>
  <section>
    <el-button class="login" type="text" @click="LoginDialogVisible = true" v-if="state.token === ''">Sign in</el-button>
    <p class="login" v-else>Hello {{ state.name }}</p>
    <el-dialog title="Sign in" :visible.sync="LoginDialogVisible" size="tiny">
      <el-form :model="form" :rules="rules" ref="form">
        <el-form-item label="Username" prop="name">
          <el-input v-model="form.name" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="Password" prop="password">
          <el-input v-model="form.password" auto-complete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="LoginDialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="getMe">Sign in</el-button>
      </div>
    </el-dialog>
  </section>
</template>

<script>
import api from '@/api.conf';

export default {
  name: 'login',
  data() {
    return {
      LoginDialogVisible: false,
      form: {
        name: '',
        password: '',
      },
      state: {
        name: '',
        token: '',
      },
      rules: {
        name: [
          { required: true, message: 'Please enter your username', trigger: 'blur' },
        ],
        password: [
          { required: true, message: 'Please enter your password', trigger: 'blur' },
        ],
      },
    };
  },
  created() {
    if (this.getCookie('token') !== '') {
      this.state = {
        name: this.getCookie('name'),
        token: this.getCookie('token'),
      };
    }
  },
  methods: {
    getMe() {
      fetch(api.getMe, {
        method: 'post',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          Username: this.form.name,
          Password: this.form.password,
        }),
      }).then((response) => {
        response.json().then((data) => {
          this.state = {
            name: this.form.name,
            token: data.Token,
          };
          this.setCookie('name', this.form.name, 365);
          this.setCookie('token', data.Token, 365);
          this.LoginDialogVisible = false;
          this.form.name = '';
          this.form.password = '';
        });
      });
    },
    setCookie(cname, cvalue, exdays) {
      const d = new Date();
      d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
      document.cookie = `${cname}=${cvalue};expires=${d.toGMTString()};path=/`;
    },
    getCookie(cname) {
      const name = `${cname}=`;
      let res = '';
      decodeURIComponent(document.cookie).split(';').forEach((ca) => {
        let a = ca;
        while (a.charAt(0) === ' ') {
          a = a.substring(1);
        }
        if (a.indexOf(name) === 0) {
          res = a.substring(name.length, a.length);
        }
      });
      return res;
    },
  },
};
</script>

<style scoped>
.login {
  float: right;
  color: #324057;
}
</style>
