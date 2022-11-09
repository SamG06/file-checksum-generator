import { defineStore } from "pinia";
import { ref } from "vue";

export const useServerStore = defineStore("server", () => {
  const serverAddress = ref("http://localhost:3000/");

  return { serverAddress };
});
