export const getEnv = () => {
  return {
    serverURL: process.env.NEXT_PUBLIC_SERVER_URL,
    clientURL: process.env.NEXT_PUBLIC_FRONT_URL,
  };
};
