<template>
  <div>
    <Table :data="audits" :title="title" />
  </div>
</template>

<script>
import Table from "@/components/Table.vue";

export default {
  name: "AuditAll",
  components: {
    Table,
  },
  props: {
      query: String
  },
  data() {
    return {
      audits: [],
      title: ""
    };
  },
  methods: {
    fetchAudits() {
      fetch("http://localhost:8090/api/fivem-audit/auditlog-v1")
        .then((response) => response.json())
        .then((data) => {
          console.log(this.query);
            switch (this.query) {
              case "all": {
                this.audits = data.Audits; 
                break;
              }
              case "player": {
                this.audits = this.filter(data.Audits, 'playerJoining', 'playerDropping');
                break;
              }
              case "chat": {
                this.audits = this.filter(data.Audits, 'chatMessage');
                break;
              }
            }
            this.title = `Auditlogs (${this.audits.length} audits)`
        });
    },
    filter(audits, ...logTypes) {
      if (!audits) {
        return;
      }
      return audits.filter(log => logTypes.includes(log.type))
    }
  },
  mounted() {
    this.fetchAudits();
  },
};
</script>