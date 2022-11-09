import { defineStore } from "pinia";
import { ref } from "vue";

interface useServerStore {
  serverAddress: String;
}
export const useServerStore = defineStore("server", () => {
  const serverAddress = ref("http://localhost:3000/");

  return { serverAddress };
});
