module.exports = {
  plugins: {
    "postcss-preset-env": {
      stage: 2,
    },
    ...(process.env.NODE_ENV === "production"
      ? { cssnano: { preset: "default" } }
      : {}),
  },
};
