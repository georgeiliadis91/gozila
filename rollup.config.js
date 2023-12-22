import commonjs from "@rollup/plugin-commonjs";
import resolve from "@rollup/plugin-node-resolve";
import babel from "@rollup/plugin-babel";
import livereload from "rollup-plugin-livereload";

export default {
  input: "src/main.js",
  output: {
    file: "assets/js/main.js",
    format: "es",
  },
  plugins: [
    resolve({
      browser: true,
    }),
    commonjs(),
    babel({
      babelHelpers: "bundled",
      exclude: "node_modules/**",
    }),
    livereload({
      watch: "src",
    }),
  ],
};
