<template>
  <section>
    <h1>Helpdesk</h1>
    <br>
    <el-row>
    <el-button type="primary" @click="contactHelpdesk">Contact Helpdesk</el-button>
    </el-row>
  </section>
</template>

<script>

export default {
  name: 'helpdesk',
  components: {},
  methods: {
    contactHelpdesk() {
      window.location.href = 'mailto:Helpdesk@lar.net.au';
    },
  },
};
</script>

<style scoped>

</style>
