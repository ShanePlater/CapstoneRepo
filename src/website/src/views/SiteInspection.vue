<template>
  <section>
    <div v-if="title === 'New Site Inspection Record'">
      <h1>{{ title }}</h1>
      <br>
      <el-row>
        <el-col :span="12">
          <el-form ref="form" :model="form" label-width="12.5em" label-position="left">


           <!-- CLIENT INFORMATION  --> 
          <h2 style="font-size:20px"> Enter Site Inspection Record </h2>
          
          <!-- client information -->


            <el-form-item label="Project Number">
              <el-input v-model="form.ProjectNumber"></el-input>
            </el-form-item>

            <el-form-item label="Inspection Number">
              <el-input v-model="form.InspectionNumber"></el-input>
            </el-form-item>

            <el-form-item label="Inspected By">
              <el-input v-model="form.InspectedBy"></el-input>
            </el-form-item>       

            <el-form-item label="Inspection Date">
              <el-date-picker v-model="form.InspectionDate" type="date" placeholder="Pick a date">
              </el-date-picker>
            </el-form-item>

            <el-form-item label="Details">
              <el-input v-model="form.Details" type="textarea"></el-input>
            </el-form-item>
            

            <!-- shane fix this -->
            <el-form-item>
              <el-button type="primary" @click="redirecting">Add New Site Inspection</el-button>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
    </div>
    <div v-else>
      <el-row v-loading="isSearching">
        <search :table="clients"></search>
      </el-row>
    </div>
  </section>
</template>


<style scoped>

</style>
