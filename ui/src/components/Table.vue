<template>
  <div class="table">
    <h1 class="title">{{ title }}</h1>
    <p>{{ description }}</p>
    <b-table striped paginated :data="data" :columns="columns"></b-table>
  </div>
</template>

<script>
export default {
  name: "Table",
  props: {
    data: Array,
    title: String,
    description: String
  },
  data() {
    return {
      columns: [
        {
          field: "auditId",
          label: "Audit ID",
          width: '25%'
        },
        {
          field: "audit",
          label: "Audit",
        },
        {
          field: "username",
          label: "Username",
        },
        {
          field: "userId",
          label: "User ID",
          numeric: true,
        },
        {
          field: "type",
          label: "Type",
        },
        {
          field: "timestamp",
          label: "Timestamp"
        },
      ],
    };
  }
};
</script>

<style scoped>
.table {
    padding-left: 3em;
    padding-right: 3em;
}
</style>