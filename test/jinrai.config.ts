import { defineConfig } from "jinrai";

export default defineConfig({
  url: "https://fld.ru",
  pages: [
    "/",
    "/products",
    "/products/{klapany}",
    "/product/{krany_sharovye_serii_105_prohodnoj_2h_hodovoj}",
    "/{CMC-8M-4N}",
  ],
  api: "http://nginx",
  meta: "http://nginx/Api/Meta/GetTags",
  outDir: "export",
});
