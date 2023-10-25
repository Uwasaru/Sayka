export const font = fetch(
  new URL("../../assets/NotoSansJP-Bold.otf", import.meta.url)
).then((res) => res.arrayBuffer());
