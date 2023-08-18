export const formatDateToJapanese = (isoDate: string): string => {
  const date = new Date(isoDate);

  const year = date.getFullYear();
  const month = date.getMonth() + 1;
  const day = date.getDate();
  const hours = date.getHours();
  const minutes = date.getMinutes();
  //   const seconds = date.getSeconds();

  return `${year}年${month}月${day}日 ${hours}時${minutes}分`;
};
