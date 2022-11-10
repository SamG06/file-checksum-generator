<template>
  <main>
    <h1>Hello Checksum Generator</h1>
    <p>Upload a file to generate checksums</p>
    <form
      ref="fileForm"
      enctype="multipart/form-data"
      @submit.prevent="firstRequest"
      type="multi"
    >
      <input type="file" name="file" id="file" />
      <button type="submit">Submit File</button>
    </form>
    <p>{{ serverResponse }}</p>
  </main>
</template>

<script setup lang="ts">
import { useServerStore } from "@/stores/server";
import { storeToRefs } from "pinia";
import { ref } from "vue";

const serverResponse = ref(
  "This is going to show what the server responds with"
);
const { serverAddress } = storeToRefs(useServerStore());

const fileForm = ref<HTMLFormElement | null>(null);

const firstRequest = async () => {
  if (!fileForm.value) return;

  const fileData = new FormData(fileForm.value);

  const response = await fetch(serverAddress.value, {
    method: "post",
    body: fileData,
  });

  const data = await response.text();
  serverResponse.value = data;
};
</script>

<style scoped>
main {
  width: 100%;
  height: 100vh;
  display: flex;
  flex-flow: column;
  align-items: center;
  justify-content: center;
}
h1 {
  color: white;
}

div {
  background: rgb(14, 56, 97);
  width: 300px;
  height: 20px;
}
</style>
